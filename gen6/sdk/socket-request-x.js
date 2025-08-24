import { BaseRequest } from "./base-request-x";
export class SocketRequestX extends BaseRequest {
    constructor() {
        super(...arguments);
        this.ws = null; // exposed
        this.subscribers = [];
    }
    subscribe(callback) {
        this.subscribers.push(callback);
        return () => {
            this.subscribers = this.subscribers.filter((cb) => cb !== callback);
        };
    }
    emit(data) {
        this.subscribers.forEach((cb) => cb(data));
    }
    /** Connect WS, but return the instance so user can attach custom handlers */
    connect() {
        if (!this.url)
            throw new Error("WebSocket URL not set");
        if (this.ws)
            this.ws.close();
        this.ws = new WebSocket(this.url);
        // default onmessage
        this.ws.onmessage = (event) => {
            try {
                const data = this.processMessage(event, event.data);
                this.emit(data);
            }
            catch (err) {
                console.error("Failed to parse WS message", err);
            }
        };
        return this; // allow user to attach extra handlers if they want
    }
    send(message) {
        if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
            throw new Error("WebSocket is not connected or ready");
        }
        this.ws.send(message); // raw bytes or string
    }
    disconnect() {
        var _a;
        (_a = this.ws) === null || _a === void 0 ? void 0 : _a.close();
        this.ws = null;
    }
    cleanup() {
        super.cleanup();
        this.disconnect();
    }
}
