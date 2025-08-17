import { BaseRequest, HttpMethod } from "./base-request-x";
/**
 * @template TQueryString - Type of the query string object sent with the request
 * @template THeaders - Type of the headers object
 * @template TBody - Type of the request body
 * @template TResponse - Type of the expected response data
 * @template TError - Type of the expected error object
 *
 * @description Provides common HTTP requests (GET, POST, PUT, PATCH, DELETE)
 * with optional typed response and error handling. Extends BaseRequest.
 * Supports query strings, headers, body, and typed response/error classes.
 */
export declare class WebRequestX<TQueryString = unknown, THeaders = unknown, TBody = unknown, TResponse = unknown, TError = unknown> extends BaseRequest<TResponse, TError> {
    /** HTTP method to use for the request */
    private method;
    /** Optional query string object */
    queryString: TQueryString | null;
    /** Optional headers object */
    headers: THeaders | null;
    /** Optional request body */
    body: TBody | null;
    constructor(SuccessClass?: {
        new (): any;
    }, ErrorClass?: {
        new (): TError;
    });
    /**
     * @description Sets the HTTP method (GET, POST, PATCH, etc.)
     * @param method - The HTTP method
     */
    setMethod(method: HttpMethod): this;
    /**
     * @description Returns the currently set HTTP method
     */
    getMethod(): HttpMethod | undefined;
    /**
     * @description Builds the full URL including query string
     * @param overrideUrl - Optional URL to override the instance URL
     */
    getComputedUrl(overrideUrl?: string | null): string;
    /**
     * @description Builds fetch options including method, headers, body, and signal
     * @param overrideFetchParams - Optional override for fetch options
     */
    getFetchOptions(overrideFetchParams?: RequestInit | null): RequestInit;
    /**
     * @description Executes the HTTP request and parses the response
     * @param overrideUrl - Optional URL override
     * @param overrideFetchParams - Optional fetch params override
     * @returns Promise with raw Response and typed data
     */
    exec(overrideUrl?: string | null, overrideFetchParams?: RequestInit | null): Promise<{
        raw: Response;
        data: TResponse;
    }>;
}
