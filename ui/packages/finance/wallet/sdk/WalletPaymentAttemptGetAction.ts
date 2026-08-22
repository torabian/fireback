import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletPaymentAttemptDto } from "./WalletPaymentAttemptDto";
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
 * Action to communicate with the action walletPaymentAttemptGet
 */
export type WalletPaymentAttemptGetActionOptions = {
  queryKey?: unknown[];
  params: WalletPaymentAttemptGetActionPathParameter;
  qs?: URLSearchParams;
};
export type WalletPaymentAttemptGetActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<WalletPaymentAttemptDto>,
    unknown[]
  >,
  "queryKey"
> &
  WalletPaymentAttemptGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WalletPaymentAttemptDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWalletPaymentAttemptGetActionQuery = (
  options: WalletPaymentAttemptGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WalletPaymentAttemptGetAction.Fetch(
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
      WalletPaymentAttemptGetAction.NewUrl(options.params, options?.qs),
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
export type WalletPaymentAttemptGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletPaymentAttemptGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletPaymentAttemptDto;
  }>;
export const useWalletPaymentAttemptGetAction = (
  options: WalletPaymentAttemptGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WalletPaymentAttemptGetAction.Fetch(
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
 * Path parameters for WalletPaymentAttemptGetAction
 */
export type WalletPaymentAttemptGetActionPathParameter = {
  uniqueId: string;
};
/**
 * WalletPaymentAttemptGetAction
 */
export class WalletPaymentAttemptGetAction {
  //
  static URL = "/walletPaymentAttempt/:uniqueId";
  static NewUrl = (
    params: WalletPaymentAttemptGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WalletPaymentAttemptGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: WalletPaymentAttemptGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletPaymentAttemptDto>, unknown, unknown>(
      overrideUrl ?? WalletPaymentAttemptGetAction.NewUrl(params, qs),
      {
        method: WalletPaymentAttemptGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WalletPaymentAttemptGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletPaymentAttemptDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletPaymentAttemptDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletPaymentAttemptDto(item));
    const res = await WalletPaymentAttemptGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletPaymentAttemptDto>();
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
    name: "walletPaymentAttemptGet",
    cliName: "get",
    cliShort: "g",
    url: "/walletPaymentAttempt/:uniqueId string",
    method: "get",
    description: 'Looks up a single "walletPaymentAttempt" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "WalletPaymentAttemptDto",
    },
  };
}
