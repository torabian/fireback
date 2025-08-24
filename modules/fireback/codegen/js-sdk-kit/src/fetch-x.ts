export type TypedRequestInit<TBody = unknown, THeaders = unknown> = Omit<
  RequestInit,
  "body" | "headers"
> & {
  body?: TBody;
  headers?: THeaders;
};

class TypedResponse<T> extends Response {
  json(): Promise<T> {
    return super.json();
  }
}

export function fetchx<
  TResponse = undefined,
  TBody = unknown,
  THeaders = unknown
>(
  input: RequestInfo | URL,
  init?: TypedRequestInit<TBody, THeaders>
): Promise<TypedResponse<TResponse>> {
  return fetch(input, init as RequestInit) as Promise<TypedResponse<TResponse>>;
}
