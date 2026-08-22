import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletTransactionDto } from "./WalletTransactionDto";
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
 * Action to communicate with the action walletTransactionGet
 */
export type WalletTransactionGetActionOptions = {
  queryKey?: unknown[];
  params: WalletTransactionGetActionPathParameter;
  qs?: URLSearchParams;
};
export type WalletTransactionGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<WalletTransactionDto>, unknown[]>,
  "queryKey"
> &
  WalletTransactionGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WalletTransactionDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWalletTransactionGetActionQuery = (
  options: WalletTransactionGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WalletTransactionGetAction.Fetch(
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
    queryKey: [WalletTransactionGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WalletTransactionGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletTransactionGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletTransactionDto;
  }>;
export const useWalletTransactionGetAction = (
  options: WalletTransactionGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WalletTransactionGetAction.Fetch(
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
 * Path parameters for WalletTransactionGetAction
 */
export type WalletTransactionGetActionPathParameter = {
  uniqueId: string;
};
/**
 * WalletTransactionGetAction
 */
export class WalletTransactionGetAction {
  //
  static URL = "/walletTransaction/:uniqueId";
  static NewUrl = (
    params: WalletTransactionGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WalletTransactionGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: WalletTransactionGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletTransactionDto>, unknown, unknown>(
      overrideUrl ?? WalletTransactionGetAction.NewUrl(params, qs),
      {
        method: WalletTransactionGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WalletTransactionGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletTransactionDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletTransactionDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletTransactionDto(item));
    const res = await WalletTransactionGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletTransactionDto>();
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
    name: "walletTransactionGet",
    cliName: "get",
    cliShort: "g",
    url: "/walletTransaction/:uniqueId string",
    method: "get",
    description: 'Looks up a single "walletTransaction" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "WalletTransactionDto",
    },
  };
}
