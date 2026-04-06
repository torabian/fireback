import { buildUrl } from "../../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
  type TypedResponse,
} from "../../sdk/common/fetchx";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "react-query";
import { useFetchxContext } from "../../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action ProductSingleScreen
 */
export type ProductSingleScreenActionOptions = {
  queryKey?: unknown[];
  params: ProductSingleScreenActionPathParameter;
  qs?: URLSearchParams;
};
export type ProductSingleScreenActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, unknown[]>,
  "queryKey"
> &
  ProductSingleScreenActionOptions & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext;
  };
export const useProductSingleScreenActionQuery = (
  options: ProductSingleScreenActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return ProductSingleScreenAction.Fetch(
      options.params,
      {
        headers: options?.headers,
      },
      {
        creatorFn: options?.creatorFn,
        qs: options?.qs,
        ctx,
        onMessage: options?.onMessage,
        overrideUrl: options?.overrideUrl,
      },
    ).then((x) => {
      x.done.then(() => {
        setCompleteState(true);
      });
      setResponse(x.response);
      return x.response.result;
    });
  };
  const result = useQuery({
    queryKey: [ProductSingleScreenAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type ProductSingleScreenActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  ProductSingleScreenActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  };
export const useProductSingleScreenAction = (
  options: ProductSingleScreenActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return ProductSingleScreenAction.Fetch(
      options.params,
      {
        body,
        headers: options?.headers,
      },
      {
        creatorFn: options?.creatorFn,
        qs: options?.qs,
        ctx,
        onMessage: options?.onMessage,
        overrideUrl: options?.overrideUrl,
      },
    ).then((x) => {
      x.done.then(() => {
        setCompleteState(true);
      });
      setResponse(x.response);
      return x.response.result;
    });
  };
  const result = useMutation({
    mutationFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
/**
 * Path parameters for ProductSingleScreenAction
 */
export type ProductSingleScreenActionPathParameter = {
  id: string;
};
/**
 * ProductSingleScreenAction
 */
export class ProductSingleScreenAction {
  //
  static URL = "/product/:id";
  static NewUrl = (
    params: ProductSingleScreenActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(ProductSingleScreenAction.URL, params, qs);
  static Method = "get";
  static Fetch$ = async (
    params: ProductSingleScreenActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<unknown, unknown, unknown>(
      overrideUrl ?? ProductSingleScreenAction.NewUrl(params, qs),
      {
        method: ProductSingleScreenAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: ProductSingleScreenActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {},
  ) => {
    const res = await ProductSingleScreenAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(res, undefined, onMessage, init?.signal);
  };
  static Definition = {
    name: "ProductSingleScreen",
    url: "/product/:id",
    method: "get",
    description: "When a user opens a product single screen",
    out: {
      primitive: "string",
    },
  };
}
