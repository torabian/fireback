import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletEventDto } from "./WalletEventDto";
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
 * Action to communicate with the action walletEventGet
 */
export type WalletEventGetActionOptions = {
  queryKey?: unknown[];
  params: WalletEventGetActionPathParameter;
  qs?: URLSearchParams;
};
export type WalletEventGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<WalletEventDto>, unknown[]>,
  "queryKey"
> &
  WalletEventGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WalletEventDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWalletEventGetActionQuery = (
  options: WalletEventGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WalletEventGetAction.Fetch(
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
    queryKey: [WalletEventGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WalletEventGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletEventGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletEventDto;
  }>;
export const useWalletEventGetAction = (
  options: WalletEventGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WalletEventGetAction.Fetch(
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
 * Path parameters for WalletEventGetAction
 */
export type WalletEventGetActionPathParameter = {
  uniqueId: string;
};
/**
 * WalletEventGetAction
 */
export class WalletEventGetAction {
  //
  static URL = "/walletEvent/:uniqueId";
  static NewUrl = (
    params: WalletEventGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WalletEventGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: WalletEventGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletEventDto>, unknown, unknown>(
      overrideUrl ?? WalletEventGetAction.NewUrl(params, qs),
      {
        method: WalletEventGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WalletEventGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletEventDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletEventDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletEventDto(item));
    const res = await WalletEventGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletEventDto>();
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
    name: "walletEventGet",
    cliName: "get",
    cliShort: "g",
    url: "/walletEvent/:uniqueId string",
    method: "get",
    description: 'Looks up a single "walletEvent" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "WalletEventDto",
    },
  };
}
