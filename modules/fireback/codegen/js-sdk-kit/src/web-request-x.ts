import { stringify } from "qs";
import { BaseRequest, HttpMethod, WebResult } from "./base-request-x";

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
export class WebRequestX<
  TQueryString = unknown,
  THeaders = unknown,
  TBody = unknown,
  TResponse = unknown,
  TError = unknown
> extends BaseRequest<TResponse, TError> {
  /** HTTP method to use for the request */
  private method: HttpMethod | undefined;

  /** Optional query string object */
  public queryString: TQueryString | null = null;

  /** Optional headers object */
  public headers: THeaders | null = null;

  /** Optional request body */
  public body: TBody | null = null;

  constructor(SuccessClass?: { new (): any }, ErrorClass?: { new (): TError }) {
    super(SuccessClass, ErrorClass);
  }

  /**
   * @description Sets the HTTP method (GET, POST, PATCH, etc.)
   * @param method - The HTTP method
   */
  setMethod(method: HttpMethod) {
    this.method = method;
    return this;
  }

  /**
   * @description Returns the currently set HTTP method
   */
  getMethod(): HttpMethod | undefined {
    return this.method;
  }

  /**
   * @description Builds the full URL including query string
   * @param overrideUrl - Optional URL to override the instance URL
   */
  getComputedUrl(overrideUrl: string | null = null): string {
    const url = overrideUrl || this.getUrl();
    if (!url) throw new Error("URL not set");
    return `${url}?${stringify(this.queryString || {})}`;
  }

  /**
   * @description Builds fetch options including method, headers, body, and signal
   * @param overrideFetchParams - Optional override for fetch options
   */
  getFetchOptions(overrideFetchParams: RequestInit | null = null): RequestInit {
    const options: RequestInit = {
      method: this.method,
      headers: this.headers || {},
      signal: this.controller?.signal,
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
  exec(
    overrideUrl: string | null = null,
    overrideFetchParams: RequestInit | null = null
  ): Promise<{ raw: Response; data: TResponse }> {
    return fetch(
      this.getComputedUrl(overrideUrl),
      this.getFetchOptions(overrideFetchParams)
    ).then(async (res) => {
      let parsed: any;
      const contentType = res.headers.get("content-type") || "";

      if (contentType.includes("json")) {
        parsed = await res.json();
      } else {
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

  // /**
  //  * @description Executes the HTTP request safely, returning a success or error object
  //  * @param overrideUrl - Optional URL override
  //  * @param overrideFetchParams - Optional fetch params override
  //  */
  // execSafe(
  //   overrideUrl: string | null = null,
  //   overrideFetchParams: RequestInit | null = null
  // ): Promise<WebResult<TResponse, TError>> {
  //   return this.exec(overrideUrl, overrideFetchParams)
  //     .then((res) => ({ kind: "success", value: res.data, raw: res.raw }))
  //     .catch((err) => ({
  //       kind: "error",
  //       error: this.processError(err),
  //     })) as Promise<WebResult<TResponse, TError>>;
  // }
}
