import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { TokenDto } from "./TokenDto";
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
 * Action to communicate with the action tokenGet
 */
export type TokenGetActionOptions = {
  queryKey?: unknown[];
  params: TokenGetActionPathParameter;
  qs?: URLSearchParams;
};
export type TokenGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<TokenDto>, unknown[]>,
  "queryKey"
> &
  TokenGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => TokenDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useTokenGetActionQuery = (options: TokenGetActionQueryOptions) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return TokenGetAction.Fetch(
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
    queryKey: [TokenGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type TokenGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  TokenGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => TokenDto;
  }>;
export const useTokenGetAction = (options: TokenGetActionMutationOptions) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return TokenGetAction.Fetch(
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
 * Path parameters for TokenGetAction
 */
export type TokenGetActionPathParameter = {
  uniqueId: string;
};
/**
 * TokenGetAction
 */
export class TokenGetAction {
  //
  static URL = "/token/:uniqueId";
  static NewUrl = (params: TokenGetActionPathParameter, qs?: URLSearchParams) =>
    buildUrl(TokenGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: TokenGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<TokenDto>, unknown, unknown>(
      overrideUrl ?? TokenGetAction.NewUrl(params, qs),
      {
        method: TokenGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: TokenGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => TokenDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new TokenDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new TokenDto(item));
    const res = await TokenGetAction.Fetch$(params, qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<TokenDto>();
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
    name: "tokenGet",
    cliName: "get",
    cliShort: "token-g",
    url: "/token/:uniqueId string",
    method: "get",
    description: 'Looks up a single "token" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "TokenDto",
    },
  };
}
