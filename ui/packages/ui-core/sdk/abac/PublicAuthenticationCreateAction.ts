import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { PublicAuthenticationDto } from "./PublicAuthenticationDto";
import { buildUrl } from "@fireback/js-remote-ctx/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
  type TypedResponse,
} from "@fireback/js-remote-ctx/common/fetchx";
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action publicAuthenticationCreate
 */
export type PublicAuthenticationCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type PublicAuthenticationCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PublicAuthenticationCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PublicAuthenticationDto;
  }>;
export const usePublicAuthenticationCreateAction = (
  options?: PublicAuthenticationCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PublicAuthenticationDto) => {
    setCompleteState(false);
    return PublicAuthenticationCreateAction.Fetch(
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
 * PublicAuthenticationCreateAction
 */
export class PublicAuthenticationCreateAction {
  //
  static URL = "/publicAuthentication";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(PublicAuthenticationCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PublicAuthenticationDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<PublicAuthenticationDto>,
      PublicAuthenticationDto,
      unknown
    >(
      overrideUrl ?? PublicAuthenticationCreateAction.NewUrl(qs),
      {
        method: PublicAuthenticationCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<PublicAuthenticationDto, unknown>,
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
    const res = await PublicAuthenticationCreateAction.Fetch$(
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
    name: "publicAuthenticationCreate",
    cliName: "create",
    cliShort: "publicAuthentication-c",
    url: "/publicAuthentication",
    method: "post",
    description: 'Creates a new "publicAuthentication" row.',
    in: {
      dto: "PublicAuthenticationDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PublicAuthenticationDto",
    },
  };
}
