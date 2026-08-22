import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletProviderConfigDto } from "./WalletProviderConfigDto";
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
 * Action to communicate with the action walletProviderConfigGet
 */
export type WalletProviderConfigGetActionOptions = {
  queryKey?: unknown[];
  params: WalletProviderConfigGetActionPathParameter;
  qs?: URLSearchParams;
};
export type WalletProviderConfigGetActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<WalletProviderConfigDto>,
    unknown[]
  >,
  "queryKey"
> &
  WalletProviderConfigGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WalletProviderConfigDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWalletProviderConfigGetActionQuery = (
  options: WalletProviderConfigGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WalletProviderConfigGetAction.Fetch(
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
    queryKey: [
      WalletProviderConfigGetAction.NewUrl(options.params, options?.qs),
    ],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WalletProviderConfigGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletProviderConfigGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletProviderConfigDto;
  }>;
export const useWalletProviderConfigGetAction = (
  options: WalletProviderConfigGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WalletProviderConfigGetAction.Fetch(
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
 * Path parameters for WalletProviderConfigGetAction
 */
export type WalletProviderConfigGetActionPathParameter = {
  uniqueId: string;
};
/**
 * WalletProviderConfigGetAction
 */
export class WalletProviderConfigGetAction {
  //
  static URL = "/walletProviderConfig/:uniqueId";
  static NewUrl = (
    params: WalletProviderConfigGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WalletProviderConfigGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: WalletProviderConfigGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletProviderConfigDto>, unknown, unknown>(
      overrideUrl ?? WalletProviderConfigGetAction.NewUrl(params, qs),
      {
        method: WalletProviderConfigGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WalletProviderConfigGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletProviderConfigDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletProviderConfigDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletProviderConfigDto(item));
    const res = await WalletProviderConfigGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletProviderConfigDto>();
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
    name: "walletProviderConfigGet",
    cliName: "get",
    cliShort: "g",
    url: "/walletProviderConfig/:uniqueId string",
    method: "get",
    description: 'Looks up a single "walletProviderConfig" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "WalletProviderConfigDto",
    },
  };
}
