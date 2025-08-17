export type WebResult<T, E> = {
    kind: "success";
    value: T;
    raw: Response;
} | {
    kind: "error";
    error: E;
};
export declare enum HttpMethod {
    GET = "GET",
    POST = "POST",
    PUT = "PUT",
    PATCH = "PATCH",
    DELETE = "DELETE"
}
export declare class JsonMessage {
    static __jsonParsable: boolean;
}
export declare function JsonParsable<T extends {
    new (...args: any[]): {};
}>(constructor: T): T;
export declare function isJsonParsable(cls: any): boolean;
/**
 * @description Base class for network-based requests.
 * Provides common properties and methods for URL handling,
 * typed responses/messages, and cleanup.
 */
export declare abstract class BaseRequest<TMessage = unknown, TError = unknown> {
    /** The URL or endpoint for the request/connection */
    protected url: string;
    /** Optional typed class for response/message deserialization */
    protected MessageClass?: {
        new (raw?: string): TMessage;
    };
    /** Optional typed class for error deserialization */
    protected ErrorClass?: {
        new (): TError;
    };
    /** Optional abort or cleanup controller */
    protected controller?: AbortController;
    constructor(MessageClass?: {
        new (): TMessage;
    }, ErrorClass?: {
        new (): TError;
    });
    /**
     * @description Sets or replaces the URL
     * @param url - The endpoint or server URL
     */
    setUrl(url: string): this;
    /**
     * @description Returns the current URL
     */
    getUrl(): string;
    /**
     * @description Sets a controller for aborting or cleanup
     * @param controller - AbortController or custom controller
     */
    setController(controller: AbortController): this;
    /**
     * @description Cleans up resources (connections, streams, etc.)
     * Override in subclasses if needed
     */
    cleanup(): void;
    protected deserialize<T>(cls: new () => T, response: Response, raw: any): T;
    protected processMessage(response: Response, raw: any): TMessage;
    protected processError(response: Response, err: any): TError;
}
