import { GResponse } from "../sdk/envelopes/index";
import { TokenDto } from "./TokenDto";
import { TokenOptionalDto } from "./TokenOptionalDto";
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
 * Action to communicate with the action tokenUpdate
 */
export type TokenUpdateActionOptions = {
  queryKey?: unknown[];
  params: TokenUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type TokenUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  TokenUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => TokenDto;
  }>;
export const useTokenUpdateAction = (
  options: TokenUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: TokenOptionalDto) => {
    setCompleteState(false);
    return TokenUpdateAction.Fetch(
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
 * Path parameters for TokenUpdateAction
 */
export type TokenUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * TokenUpdateAction
 */
export class TokenUpdateAction {
  //
  static URL = "/token/:uniqueId";
  static NewUrl = (
    params: TokenUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(TokenUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: TokenUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<TokenOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<TokenDto>, TokenOptionalDto, unknown>(
      overrideUrl ?? TokenUpdateAction.NewUrl(params, qs),
      {
        method: TokenUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: TokenUpdateActionPathParameter,
    init?: TypedRequestInit<TokenOptionalDto, unknown>,
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
    const res = await TokenUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
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
    name: "tokenUpdate",
    cliShort: "token-u",
    url: "/token/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "token" row by uniqueId.',
    in: {
      dto: "TokenOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "TokenDto",
    },
  };
}
