import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletGatewayDto } from "./WalletGatewayDto";
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
 * Action to communicate with the action walletGatewayGet
 */
export type WalletGatewayGetActionOptions = {
  queryKey?: unknown[];
  params: WalletGatewayGetActionPathParameter;
  qs?: URLSearchParams;
};
export type WalletGatewayGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<WalletGatewayDto>, unknown[]>,
  "queryKey"
> &
  WalletGatewayGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WalletGatewayDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWalletGatewayGetActionQuery = (
  options: WalletGatewayGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WalletGatewayGetAction.Fetch(
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
    queryKey: [WalletGatewayGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WalletGatewayGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletGatewayGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletGatewayDto;
  }>;
export const useWalletGatewayGetAction = (
  options: WalletGatewayGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WalletGatewayGetAction.Fetch(
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
 * Path parameters for WalletGatewayGetAction
 */
export type WalletGatewayGetActionPathParameter = {
  uniqueId: string;
};
/**
 * WalletGatewayGetAction
 */
export class WalletGatewayGetAction {
  //
  static URL = "/walletGateway/:uniqueId";
  static NewUrl = (
    params: WalletGatewayGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WalletGatewayGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: WalletGatewayGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletGatewayDto>, unknown, unknown>(
      overrideUrl ?? WalletGatewayGetAction.NewUrl(params, qs),
      {
        method: WalletGatewayGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WalletGatewayGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletGatewayDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletGatewayDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletGatewayDto(item));
    const res = await WalletGatewayGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletGatewayDto>();
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
    name: "walletGatewayGet",
    cliName: "get",
    cliShort: "g",
    url: "/walletGateway/:uniqueId string",
    method: "get",
    description: 'Looks up a single "walletGateway" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "WalletGatewayDto",
    },
  };
}
