import { BaseRequest } from "./base-request-x";
export type Subscriber<T> = (value: T) => void;
export declare class SocketRequestX<TMessage = unknown, TError = unknown> extends BaseRequest<TMessage, TError> {
    ws: WebSocket | null;
    private subscribers;
    subscribe(callback: Subscriber<TMessage>): () => void;
    private emit;
    /** Connect WS, but return the instance so user can attach custom handlers */
    connect(): this;
    send(message: string | ArrayBuffer | Uint8Array): void;
    disconnect(): void;
    cleanup(): void;
}
