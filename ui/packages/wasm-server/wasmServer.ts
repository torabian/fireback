// wasmServer — downloads a fireback server compiled to wasm
// (cmd/fireback-wasm, `make wasm`) and boots it in the browser, then exposes
// a fetch-compatible bridge to it. TS port of the loader inlined in the emi
// in-browser-server example's browser/main.js, extracted into a reusable
// function per fireback's own request.
//
// Usage:
//
//   import { startWasmServer, wasmFetchOverride } from "@fireback/wasm-server/wasmServer";
//   import { FetchxContext } from "@fireback/js-remote-ctx/common/fetchx";
//
//   await startWasmServer(); // downloads + boots wasm_exec.js, the .wasm, and pglite
//   const ctx = new FetchxContext("", {}, undefined, undefined, wasmFetchOverride());
//   // every fetchx(url, init, ctx) call now lands on the in-browser Go server
//
// Most apps won't call this directly — see WithWasmServer.tsx for the
// component that gates rendering on boot completing, driven by the
// VITE_USE_WASM_SERVER build variable.
import { installPgliteBridge } from "./pgliteBridge";
import type { TypedRequestInit } from "@fireback/js-remote-ctx/common/fetchx";

declare global {
  interface Window {
    // Installed by wasm_exec.js.
    Go?: new () => GoInstance;
    // Installed by emigo.LiftWasmServer (cmd/fireback-wasm/main.go), once
    // the Go module's main() has run.
    handleWasmRequest?: (
      method: string,
      url: string,
      body: string,
      headersJSON: string,
    ) => Promise<string>;
  }
}

interface GoInstance {
  importObject: WebAssembly.Imports;
  run(instance: WebAssembly.Instance): Promise<void>;
}

/** Byte progress of the .wasm download. `total` is null when the server didn't send a Content-Length (progress is still meaningful — just show bytes loaded, not a percentage). */
export interface WasmDownloadProgress {
  loaded: number;
  total: number | null;
}

type ProgressListener = (progress: WasmDownloadProgress) => void;

const progressListeners = new Set<ProgressListener>();
let lastProgress: WasmDownloadProgress | null = null;

function emitProgress(progress: WasmDownloadProgress): void {
  lastProgress = progress;
  for (const listener of progressListeners) listener(progress);
}

/**
 * Subscribe to .wasm download progress. Independent of who actually calls
 * startWasmServer — a late subscriber (e.g. a second component mounting
 * useWasmServer after boot already started) is replayed the most recent
 * progress immediately, then gets live updates same as everyone else.
 * Returns an unsubscribe function.
 */
export function onWasmDownloadProgress(listener: ProgressListener): () => void {
  progressListeners.add(listener);
  if (lastProgress) listener(lastProgress);
  return () => {
    progressListeners.delete(listener);
  };
}

export interface WasmServerOptions {
  /** Where to fetch the compiled server from. Default "/fireback.wasm". */
  wasmUrl?: string;
  /** Where to fetch Go's wasm runtime glue from. Default "/wasm_exec.js" — `make wasm` (or `cp $(go env GOROOT)/lib/wasm/wasm_exec.js ui/public/`) puts it in ui/public, which Vite serves at the site root unmodified. */
  wasmExecUrl?: string;
  /**
   * Storage location for the in-browser Postgres. Default persists across
   * reloads via IndexedDB — see installPgliteBridge.
   */
  pgliteDataDir?: string;
  /**
   * Skip installing the pglite-backed window.queryDatabase bridge, e.g. if
   * the caller wired up its own (or the wasm binary doesn't touch a
   * database at all). Default false.
   */
  skipDatabaseBridge?: boolean;
}

let bootPromise: Promise<void> | null = null;

/**
 * Fetches wasm_exec.js (if window.Go isn't already defined) and the compiled
 * server binary, instantiates and starts it, wires up the pglite database
 * bridge it expects, and resolves once window.handleWasmRequest is live.
 * Safe to call more than once — later calls just await the first boot.
 */
export function startWasmServer(opts: WasmServerOptions = {}): Promise<void> {
  if (!bootPromise) {
    bootPromise = bootWasmServer(opts)
      .catch((err) => {
        // Let a failed boot be retried instead of permanently wedging every
        // future caller on the same rejected promise.
        bootPromise = null;
        throw err;
      })
      .then();
  }
  return bootPromise;
}

async function bootWasmServer(opts: WasmServerOptions): Promise<void> {
  const wasmUrl = opts.wasmUrl ?? "/fireback.wasm";
  const wasmExecUrl = opts.wasmExecUrl ?? "/wasm_exec.js";

  // Go's main() (cmd/fireback-wasm/main.go) reads window.queryDatabase
  // synchronously at startup and bails if it's missing, so this has to be
  // in place before go.run() below.
  if (!opts.skipDatabaseBridge) {
    await installPgliteBridge(opts.pgliteDataDir);
  }

  await loadWasmExec(wasmExecUrl);

  const go = new window.Go!();
  const { instance } = await WebAssembly.instantiateStreaming(
    fetchWasmWithProgress(wasmUrl),
    go.importObject,
  );

  // Deliberately not awaited: main() blocks on `select {}` forever so the
  // exposed callback stays callable. Awaiting here would hang startup.
  void go.run(instance);

  await waitFor(() => typeof window.handleWasmRequest === "function");
}

// Wraps fetch(url) so every chunk of the (typically tens-of-MB) .wasm
// download is reported via onWasmDownloadProgress, while still handing
// WebAssembly.instantiateStreaming a real streamed Response — the fetch
// itself, and streaming compilation off it, keep running exactly as before,
// this just taps the byte stream as it passes through.
function fetchWasmWithProgress(url: string): Promise<Response> {
  return fetch(url).then((res) => {
    if (!res.ok || !res.body) {
      emitProgress({ loaded: 0, total: null });
      return res;
    }

    const totalHeader = res.headers.get("content-length");
    const total = totalHeader ? Number(totalHeader) : null;
    let loaded = 0;

    const reader = res.body.getReader();
    const stream = new ReadableStream<Uint8Array>({
      async pull(controller) {
        const { done, value } = await reader.read();
        if (done) {
          controller.close();
          return;
        }
        loaded += value.byteLength;
        emitProgress({ loaded, total });
        controller.enqueue(value);
      },
      cancel(reason) {
        reader.cancel(reason);
      },
    });

    // Rebuild a Response around the tapped stream. Passing the original
    // res.headers through keeps Content-Type: application/wasm intact,
    // which instantiateStreaming needs to accept it as a streaming compile
    // rather than falling back to buffering the whole thing first.
    return new Response(stream, {
      status: res.status,
      statusText: res.statusText,
      headers: res.headers,
    });
  });
}

function loadWasmExec(url: string): Promise<void> {
  if (typeof window.Go === "function") return Promise.resolve();
  return new Promise((resolve, reject) => {
    const script = document.createElement("script");
    script.src = url;
    script.onload = () => resolve();
    script.onerror = () =>
      reject(new Error(`wasmServer: failed to load ${url}`));
    document.head.appendChild(script);
  });
}

function waitFor(predicate: () => boolean, intervalMs = 20): Promise<void> {
  return new Promise((resolve) => {
    const tick = () => {
      if (predicate()) return resolve();
      setTimeout(tick, intervalMs);
    };
    tick();
  });
}

/**
 * A drop-in fetch replacement backed by the in-browser Go server (see
 * emigo.LiftWasmServer in cmd/fireback-wasm/main.go): it turns the call into
 * a real *http.Request, runs it through the server's mux, and hands back a
 * genuine Response — callers, including generated SDK actions, can't tell
 * it isn't a network round trip.
 *
 * Requires startWasmServer() to have resolved first.
 */
export async function wasmFetch(
  url: string,
  init: TypedRequestInit = {},
): Promise<Response> {
  if (typeof window.handleWasmRequest !== "function") {
    throw new Error(
      "wasmServer: window.handleWasmRequest is not defined yet — call and await startWasmServer() first",
    );
  }

  const body =
    typeof init.body === "string"
      ? init.body
      : init.body !== undefined
        ? JSON.stringify(init.body)
        : "";

  const raw = await window.handleWasmRequest(
    init.method || "GET",
    url,
    body,
    JSON.stringify(init.headers || {}),
  );
  const {
    status,
    headers,
    body: resBody,
  } = JSON.parse(raw) as {
    status: number;
    headers: Record<string, string[]>;
    body: string;
  };

  const h = new Headers();
  for (const [k, vs] of Object.entries(headers || {})) {
    for (const v of vs) h.append(k, v);
  }
  return new Response(resBody, { status, headers: h });
}

/**
 * fetchOverrideFn for FetchxContext (see @fireback/js-remote-ctx's
 * common/fetchx.ts) that routes through wasmFetch. Plug it into a
 * FetchxContext and every fetchx() call against that context lands on the
 * in-browser server instead of the network:
 *
 *   const ctx = new FetchxContext("", {}, undefined, undefined, wasmFetchOverride());
 */
export function wasmFetchOverride() {
  return (
    input: RequestInfo | URL,
    init?: TypedRequestInit,
  ): Promise<Response> => wasmFetch(input.toString(), init);
}
