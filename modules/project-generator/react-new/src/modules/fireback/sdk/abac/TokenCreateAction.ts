import { GResponse } from "../sdk/envelopes/index";
import { TokenDto } from "./TokenDto";
import { buildUrl } from "../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
  type TypedResponse,
} from "../sdk/common/fetchx";
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action tokenCreate
 */
export type TokenCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type TokenCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  TokenCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => TokenDto;
  }>;
export const useTokenCreateAction = (
  options?: TokenCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: TokenDto) => {
    setCompleteState(false);
    return TokenCreateAction.Fetch(
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
 * TokenCreateAction
 */
export class TokenCreateAction {
  //
  static URL = "/token";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(TokenCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<TokenDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<TokenDto>, TokenDto, unknown>(
      overrideUrl ?? TokenCreateAction.NewUrl(qs),
      {
        method: TokenCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<TokenDto, unknown>,
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
    const res = await TokenCreateAction.Fetch$(qs, ctx, init, overrideUrl);
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
    name: "tokenCreate",
    cliShort: "token-c",
    url: "/token",
    method: "post",
    description: 'Creates a new "token" row.',
    in: {
      dto: "TokenDto",
    },
    out: {
      envelope: "GResponse",
      dto: "TokenDto",
    },
  };
}
