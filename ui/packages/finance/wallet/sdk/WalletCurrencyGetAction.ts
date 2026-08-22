import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletCurrencyDto } from "./WalletCurrencyDto";
import { buildUrl } from "@fireback/js-remote-ctx/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
  type TypedResponse,
} from "@fireback/js-remote-ctx/common/fetchx";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action walletCurrencyGet
 */
export type WalletCurrencyGetActionOptions = {
  queryKey?: unknown[];
  params: WalletCurrencyGetActionPathParameter;
  qs?: URLSearchParams;
};
export type WalletCurrencyGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<WalletCurrencyDto>, unknown[]>,
  "queryKey"
> &
  WalletCurrencyGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WalletCurrencyDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWalletCurrencyGetActionQuery = (
  options: WalletCurrencyGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WalletCurrencyGetAction.Fetch(
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
    queryKey: [WalletCurrencyGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WalletCurrencyGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletCurrencyGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletCurrencyDto;
  }>;
export const useWalletCurrencyGetAction = (
  options: WalletCurrencyGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WalletCurrencyGetAction.Fetch(
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
 * Path parameters for WalletCurrencyGetAction
 */
export type WalletCurrencyGetActionPathParameter = {
  uniqueId: string;
};
/**
 * WalletCurrencyGetAction
 */
export class WalletCurrencyGetAction {
  //
  static URL = "/walletCurrency/:uniqueId";
  static NewUrl = (
    params: WalletCurrencyGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WalletCurrencyGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: WalletCurrencyGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletCurrencyDto>, unknown, unknown>(
      overrideUrl ?? WalletCurrencyGetAction.NewUrl(params, qs),
      {
        method: WalletCurrencyGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WalletCurrencyGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletCurrencyDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletCurrencyDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletCurrencyDto(item));
    const res = await WalletCurrencyGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletCurrencyDto>();
        if (creatorFn) {
          resp.setCreator(creatorFn);
        }
        resp.inject(data);
        return resp;
      },
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "walletCurrencyGet",
    cliName: "get",
    cliShort: "g",
    url: "/walletCurrency/:uniqueId string",
    method: "get",
    description: 'Looks up a single "walletCurrency" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "WalletCurrencyDto",
    },
  };
}
