import { stringify } from "qs";
import { BaseRequest } from "./base-request-x";
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
export class WebRequestX extends BaseRequest {
    constructor(SuccessClass, ErrorClass) {
        super(SuccessClass, ErrorClass);
        /** Optional query string object */
        this.queryString = null;
        /** Optional headers object */
        this.headers = null;
        /** Optional request body */
        this.body = null;
    }
    /**
     * @description Sets the HTTP method (GET, POST, PATCH, etc.)
     * @param method - The HTTP method
     */
    setMethod(method) {
        this.method = method;
        return this;
    }
    /**
     * @description Returns the currently set HTTP method
     */
    getMethod() {
        return this.method;
    }
    /**
     * @description Builds the full URL including query string
     * @param overrideUrl - Optional URL to override the instance URL
     */
    getComputedUrl(overrideUrl = null) {
        const url = overrideUrl || this.getUrl();
        if (!url)
            throw new Error("URL not set");
        return `${url}?${stringify(this.queryString || {})}`;
    }
    /**
     * @description Builds fetch options including method, headers, body, and signal
     * @param overrideFetchParams - Optional override for fetch options
     */
    getFetchOptions(overrideFetchParams = null) {
        var _a;
        const options = {
            method: this.method,
            headers: this.headers || {},
            signal: (_a = this.controller) === null || _a === void 0 ? void 0 : _a.signal,
        };
        if (this.body) {
            options.body = JSON.stringify(this.body);
        }
        return {
            ...options,
            ...(overrideFetchParams || {}),
        };
    }
    /**
     * @description Executes the HTTP request and parses the response
     * @param overrideUrl - Optional URL override
     * @param overrideFetchParams - Optional fetch params override
     * @returns Promise with raw Response and typed data
     */
    exec(overrideUrl = null, overrideFetchParams = null) {
        return fetch(this.getComputedUrl(overrideUrl), this.getFetchOptions(overrideFetchParams)).then(async (res) => {
            let parsed;
            const contentType = res.headers.get("content-type") || "";
            if (contentType.includes("json")) {
                parsed = await res.json();
            }
            else {
                parsed = await res.text();
            }
            if (!res.ok) {
                // wrap in your ErrorClass if defined
                console.log(7, parsed);
                const err = this.processError(res, parsed);
                throw err;
            }
            parsed = this.processMessage(res, parsed);
            return { raw: res, data: parsed };
        });
    }
}
