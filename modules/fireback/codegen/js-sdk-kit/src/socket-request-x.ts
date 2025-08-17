import { BaseRequest } from "./base-request-x";

export type Subscriber<T> = (value: T) => void;

export class SocketRequestX<
  TMessage = unknown,
  TError = unknown
> extends BaseRequest<TMessage, TError> {
  public ws: WebSocket | null = null; // exposed

  private subscribers: Subscriber<TMessage>[] = [];

  subscribe(callback: Subscriber<TMessage>) {
    this.subscribers.push(callback);
    return () => {
      this.subscribers = this.subscribers.filter((cb) => cb !== callback);
    };
  }

  private emit(data: TMessage) {
    this.subscribers.forEach((cb) => cb(data));
  }

  /** Connect WS, but return the instance so user can attach custom handlers */
  connect() {
    if (!this.url) throw new Error("WebSocket URL not set");

    if (this.ws) this.ws.close();

    this.ws = new WebSocket(this.url);

    // default onmessage
    this.ws.onmessage = (event) => {
      try {
        const data = this.processMessage(event as any, event.data);
        this.emit(data);
      } catch (err) {
        console.error("Failed to parse WS message", err);
      }
    };

    return this; // allow user to attach extra handlers if they want
  }

  send(message: string | ArrayBuffer | Uint8Array) {
    if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
      throw new Error("WebSocket is not connected or ready");
    }
    this.ws.send(message); // raw bytes or string
  }

  disconnect() {
    this.ws?.close();
    this.ws = null;
  }

  override cleanup() {
    super.cleanup();
    this.disconnect();
  }
}
