// @ts-nocheck
export type TypedRequestInit<TBody = unknown, THeaders = unknown> = Omit<
  RequestInit,
  "body" | "headers"
> & {
  body?: TBody;
  headers?: THeaders;
};
export class TypedResponse<T> extends Response {
  override json(): Promise<T> {
    return super.json();
  }
  result:
    | T
    | undefined
    | ReadableStream<Uint8Array<ArrayBuffer>>
    | null
    | string;
}
export async function fetchx<
  TResponse = unknown,
  TBody = unknown,
  THeaders = unknown,
>(
  input: RequestInfo | URL,
  init?: TypedRequestInit<TBody, THeaders>,
  ctx?: FetchxContext,
): Promise<TypedResponse<TResponse>> {
  let url = input.toString();
  let reqInit: TypedRequestInit<TBody, THeaders> = init || {};
  let res: TypedResponse<TResponse>;
  let fetchFn = fetch;
  if (ctx) {
    [url, reqInit] = await ctx.apply(url, reqInit);
    if (ctx.fetchOverrideFn) {
      fetchFn = ctx.fetchOverrideFn;
    }
  }
  res = (await fetchFn(
    url,
    reqInit as RequestInit,
  )) as TypedResponse<TResponse>;
  if (ctx) {
    res = await ctx.handle(res);
  }
  return res;
}
type DtoFactory<T> = { new (data: any): T } | ((data: any) => T);
function isConstructor<T>(fn: DtoFactory<T>): fn is { new (data: any): T } {
  return (
    typeof fn === "function" && fn.prototype && fn.prototype.constructor === fn
  );
}
export async function handleFetchResponse<T>(
  res: TypedResponse<T>,
  dto?: DtoFactory<T>,
  onMessage?: (msg: any) => void,
  signal?: AbortSignal | null,
): Promise<{ done: Promise<void>; response: TypedResponse<T> }> {
  const ct = res.headers.get("content-type") || "";
  const cd = res.headers.get("content-disposition") || "";
  if (ct.includes("text/event-stream")) {
    return SSEFetch(res, onMessage, signal);
  }
  if (
    cd.includes("attachment") ||
    (!ct.includes("json") && !ct.startsWith("text/"))
  ) {
    (res as any).result = res.body;
  } else if (ct.includes("application/json")) {
    const json = await res.json();
    if (dto) {
      if (isConstructor(dto)) {
        (res as any).result = new dto(json); // ✅ class constructor
      } else {
        (res as any).result = dto(json); // ✅ factory function
      }
    } else {
      (res as any).result = json;
    }
  } else {
    (res as any).result = await res.text();
  }
  return { done: Promise.resolve(), response: res as any };
}
export const SSEFetch = <T = string>(
  res: TypedResponse<T>,
  onMessage?: (ev: MessageEvent) => void,
  signal?: AbortSignal | null,
): { response: TypedResponse<T>; done: Promise<void> } => {
  if (!res.body) throw new Error("SSE requires readable body");
  const reader = res.body.getReader();
  const decoder = new TextDecoder();
  let buffer = "";
  const done = new Promise<void>((resolve, reject) => {
    function readChunk() {
      reader
        .read()
        .then(({ done: finished, value }) => {
          if (signal?.aborted) {
            reader.cancel();
            return resolve(); // resolve on abort
          }
          if (finished) return resolve(); // normal end
          buffer += decoder.decode(value, { stream: true });
          const parts = buffer.split("\n\n");
          buffer = parts.pop() || "";
          for (const part of parts) {
            let data = "";
            let event = "message";
            part.split("\n").forEach((line) => {
              if (line.startsWith("data:")) data += line.slice(5).trim();
              else if (line.startsWith("event:")) event = line.slice(6).trim();
            });
            if (data) {
              if (data === "[DONE]") return resolve();
              onMessage?.(new MessageEvent(event, { data }));
            }
          }
          readChunk();
        })
        .catch((err) => {
          if (err.name === "AbortError") resolve();
          else reject(err);
        });
    }
    readChunk();
  });
  return { response: res, done };
};
export class FetchxContext {
  constructor(
    public baseUrl: string = "",
    public defaultHeaders: Record<string, string> = {},
    public requestInterceptor?: (
      url: string,
      init: TypedRequestInit<any, any>,
    ) =>
      | Promise<[string, TypedRequestInit<any, any>]>
      | [string, TypedRequestInit<any, any>],
    public responseInterceptor?: <T>(
      res: TypedResponse<T>,
    ) => Promise<TypedResponse<T>>,
    /**
     * Overrides the browser fetch function, for different purposes. It would recieve the same first 2 arguments as fetch,
     * as well as third one of fetchx context. If you pass the fetch itself to override, it should have no effect.
     */
    public fetchOverrideFn: (
      input: RequestInfo | URL,
      init?: TypedRequestInit,
    ) => Promise<Response> = null,
  ) {}
  async apply<T>(
    url: string,
    init: TypedRequestInit<any, any>,
  ): Promise<[string, TypedRequestInit<any, any>]> {
    // prefix baseUrl
    if (!/^https?:\/\//.test(url)) {
      url = this.baseUrl + url;
    }
    // merge default headers
    (init.headers as unknown) = {
      ...this.defaultHeaders,
      ...((init.headers as object) || {}),
    };
    // call request interceptor if present
    if (this.requestInterceptor) {
      return this.requestInterceptor(url, init);
    }
    return [url, init];
  }
  async handle<T>(res: TypedResponse<T>): Promise<TypedResponse<T>> {
    if (this.responseInterceptor) {
      return this.responseInterceptor(res);
    }
    return res;
  }
  clone(overrides?: Partial<FetchxContext>): FetchxContext {
    return new FetchxContext(
      overrides?.baseUrl ?? this.baseUrl,
      { ...this.defaultHeaders, ...(overrides?.defaultHeaders || {}) },
      overrides?.requestInterceptor ?? this.requestInterceptor,
      overrides?.responseInterceptor ?? this.responseInterceptor,
    );
  }
}
