import { useEffect, useState } from "react";
import { startWasmServer, wasmFetchOverride, type WasmServerOptions } from "./wasmServer";
import { FetchxContext } from "@fireback/js-remote-ctx/common/fetchx";

export interface UseWasmServerResult {
  /** true once window.handleWasmRequest is live and ctx is safe to use. */
  ready: boolean;
  /** Set if downloading/booting the wasm server failed. */
  error: Error | null;
  /**
   * A FetchxContext wired to route through the in-browser server. Pass it as
   * fetchx()'s third argument (or via FetchxProvider) the same way you'd use
   * a context pointed at a real API base URL — calls just no-op/reject until
   * `ready` flips true.
   */
  ctx: FetchxContext;
}

/**
 * Downloads and boots cmd/fireback-wasm's compiled server (see
 * wasmServer.ts) on mount, and returns a FetchxContext that routes through it
 * once ready. One boot is shared across every component that calls this hook
 * (see startWasmServer's memoization) — mounting it twice doesn't download or
 * start a second server.
 *
 * Most apps won't call this directly — see WithWasmServer.tsx, which wraps
 * it behind the VITE_USE_WASM_SERVER build variable.
 */
export function useWasmServer(opts?: WasmServerOptions): UseWasmServerResult {
  const [ready, setReady] = useState(false);
  const [error, setError] = useState<Error | null>(null);
  const [ctx] = useState(
    () => new FetchxContext("", {}, undefined, undefined, wasmFetchOverride()),
  );

  useEffect(() => {
    let cancelled = false;
    startWasmServer(opts)
      .then(() => {
        if (!cancelled) setReady(true);
      })
      .catch((err: Error) => {
        if (!cancelled) setError(err);
      });
    return () => {
      cancelled = true;
    };
    // opts is intentionally read once — startWasmServer boots a single
    // shared instance, changing options after boot wouldn't do anything.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return { ready, error, ctx };
}
