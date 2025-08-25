export type TypedRequestInit<TBody = unknown, THeaders = unknown> = Omit<RequestInit, "body" | "headers"> & {
    body?: TBody;
    headers?: THeaders;
};
declare class TypedResponse<T> extends Response {
    json(): Promise<T>;
}
export declare function fetchx<TResponse = undefined, TBody = unknown, THeaders = unknown>(input: RequestInfo | URL, init?: TypedRequestInit<TBody, THeaders>): Promise<TypedResponse<TResponse>>;
export {};
