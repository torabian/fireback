import { GResponse } from "../sdk/envelopes/index";
import { PublicAuthenticationDto } from "./PublicAuthenticationDto";
import { PublicAuthenticationOptionalDto } from "./PublicAuthenticationOptionalDto";
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
 * Action to communicate with the action publicAuthenticationUpdate
 */
export type PublicAuthenticationUpdateActionOptions = {
  queryKey?: unknown[];
  params: PublicAuthenticationUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type PublicAuthenticationUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PublicAuthenticationUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PublicAuthenticationDto;
  }>;
export const usePublicAuthenticationUpdateAction = (
  options: PublicAuthenticationUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PublicAuthenticationOptionalDto) => {
    setCompleteState(false);
    return PublicAuthenticationUpdateAction.Fetch(
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
 * Path parameters for PublicAuthenticationUpdateAction
 */
export type PublicAuthenticationUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * PublicAuthenticationUpdateAction
 */
export class PublicAuthenticationUpdateAction {
  //
  static URL = "/publicAuthentication/:uniqueId";
  static NewUrl = (
    params: PublicAuthenticationUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PublicAuthenticationUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: PublicAuthenticationUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PublicAuthenticationOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<PublicAuthenticationDto>,
      PublicAuthenticationOptionalDto,
      unknown
    >(
      overrideUrl ?? PublicAuthenticationUpdateAction.NewUrl(params, qs),
      {
        method: PublicAuthenticationUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PublicAuthenticationUpdateActionPathParameter,
    init?: TypedRequestInit<PublicAuthenticationOptionalDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => PublicAuthenticationDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new PublicAuthenticationDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new PublicAuthenticationDto(item));
    const res = await PublicAuthenticationUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<PublicAuthenticationDto>();
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
    name: "publicAuthenticationUpdate",
    cliShort: "publicAuthentication-u",
    url: "/publicAuthentication/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "publicAuthentication" row by uniqueId.',
    in: {
      dto: "PublicAuthenticationOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PublicAuthenticationDto",
    },
  };
}
