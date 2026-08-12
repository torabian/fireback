// WebSocketWasm — a drop-in replacement for the native WebSocket, backed by the
// Go-WASM reactor (emigo.WasmReactor). Swap `new WebSocket(url)` for
// `new WebSocketWasm(url)` and the rest of your code is identical: readyState,
// onopen/onmessage/onclose/onerror, addEventListener, send(), close().
//
// It is NOT a network client — there is no handshake and no socket. The
// constructor's URL is only used for its pathname (to pick the reactor handler)
// and query string; the host is ignored.

// The Go-WASM reactor installs these globals on `window`. Declare them so the
// bridge calls below type-check.
declare global {
  interface Window {
    wasmWsOpen?: (
      path: string,
      query: string,
      onMessage: (data: string) => void,
      onClose: (reason: string) => void,
    ) => number;
    wasmWsSend: (id: number, data: string) => void;
    wasmWsClose: (id: number) => void;
  }
}

export class WebSocketWasm extends EventTarget {
  static readonly CONNECTING = 0;
  static readonly OPEN = 1;
  static readonly CLOSING = 2;
  static readonly CLOSED = 3;

  url: string;
  readyState: number;
  onopen: ((ev: Event) => void) | null;
  onmessage: ((ev: MessageEvent) => void) | null;
  onerror: ((ev: Event) => void) | null;
  onclose: ((ev: CloseEvent) => void) | null;
  private _id: number;

  constructor(url: string, _protocols?: string | string[]) {
    super();
    this.url = url;
    this.readyState = WebSocketWasm.CONNECTING;
    this.onopen = null;
    this.onmessage = null;
    this.onerror = null;
    this.onclose = null;
    this._id = -1;

    const u = new URL(url, "ws://wasm.local"); // base lets bare paths parse
    const path = u.pathname;
    const query = u.search.replace(/^\?/, "");

    // Defer so listeners attached right after `new` still catch 'open', exactly
    // like the native async connect.
    queueMicrotask(() => this._open(path, query));
  }

  private _open(path: string, query: string): void {
    const onMessage = (data: string): void => {
      const ev = new MessageEvent("message", { data });
      this.onmessage?.(ev);
      this.dispatchEvent(ev);
    };
    const onClose = (reason: string): void => {
      if (this.readyState === WebSocketWasm.CLOSED) return;
      this.readyState = WebSocketWasm.CLOSED;
      const ev = new CloseEvent("close", {
        reason,
        code: 1000,
        wasClean: true,
      });
      this.onclose?.(ev);
      this.dispatchEvent(ev);
    };

    if (typeof window.wasmWsOpen !== "function") {
      return this._fail("reactor not ready (window.wasmWsOpen missing)");
    }

    this._id = window.wasmWsOpen(path, query, onMessage, onClose);
    if (this._id < 0) {
      return this._fail(`no reactive handler registered for "${path}"`);
    }

    this.readyState = WebSocketWasm.OPEN;
    const ev = new Event("open");
    this.onopen?.(ev);
    this.dispatchEvent(ev);
  }

  private _fail(message: string): void {
    this.readyState = WebSocketWasm.CLOSED;
    const err: Event & { message?: string } = new Event("error");
    err.message = message;
    this.onerror?.(err);
    this.dispatchEvent(err);
    const close = new CloseEvent("close", {
      reason: message,
      code: 1006,
      wasClean: false,
    });
    this.onclose?.(close);
    this.dispatchEvent(close);
  }

  send(data: unknown): void {
    if (this.readyState !== WebSocketWasm.OPEN) {
      throw new DOMException("WebSocketWasm is not open", "InvalidStateError");
    }
    // Strings pass through; everything else is stringified for now. Real binary
    // (ArrayBuffer/Blob) would carry a type flag across the bridge — future work.
    window.wasmWsSend(this._id, typeof data === "string" ? data : String(data));
  }

  close(_code?: number, _reason?: string): void {
    if (
      this.readyState === WebSocketWasm.CLOSING ||
      this.readyState === WebSocketWasm.CLOSED
    ) {
      return;
    }
    this.readyState = WebSocketWasm.CLOSING;
    if (this._id >= 0) window.wasmWsClose(this._id);
  }
}
