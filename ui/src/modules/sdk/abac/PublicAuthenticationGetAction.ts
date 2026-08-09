import { GResponse } from "../sdk/envelopes/index";
import { PublicAuthenticationDto } from "./PublicAuthenticationDto";
import { buildUrl } from "../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
  type TypedResponse,
} from "../sdk/common/fetchx";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "@tanstack/react-query";
import { useFetchxContext } from "../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action publicAuthenticationGet
 */
export type PublicAuthenticationGetActionOptions = {
  queryKey?: unknown[];
  params: PublicAuthenticationGetActionPathParameter;
  qs?: URLSearchParams;
};
export type PublicAuthenticationGetActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<PublicAuthenticationDto>,
    unknown[]
  >,
  "queryKey"
> &
  PublicAuthenticationGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => PublicAuthenticationDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const usePublicAuthenticationGetActionQuery = (
  options: PublicAuthenticationGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return PublicAuthenticationGetAction.Fetch(
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
      PublicAuthenticationGetAction.NewUrl(options.params, options?.qs),
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
export type PublicAuthenticationGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PublicAuthenticationGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PublicAuthenticationDto;
  }>;
export const usePublicAuthenticationGetAction = (
  options: PublicAuthenticationGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return PublicAuthenticationGetAction.Fetch(
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
 * Path parameters for PublicAuthenticationGetAction
 */
export type PublicAuthenticationGetActionPathParameter = {
  uniqueId: string;
};
/**
 * PublicAuthenticationGetAction
 */
export class PublicAuthenticationGetAction {
  //
  static URL = "/publicAuthentication/:uniqueId";
  static NewUrl = (
    params: PublicAuthenticationGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PublicAuthenticationGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: PublicAuthenticationGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<PublicAuthenticationDto>, unknown, unknown>(
      overrideUrl ?? PublicAuthenticationGetAction.NewUrl(params, qs),
      {
        method: PublicAuthenticationGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PublicAuthenticationGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
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
    const res = await PublicAuthenticationGetAction.Fetch$(
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
    name: "publicAuthenticationGet",
    cliName: "get",
    cliShort: "publicAuthentication-g",
    url: "/publicAuthentication/:uniqueId string",
    method: "get",
    description: 'Looks up a single "publicAuthentication" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "PublicAuthenticationDto",
    },
  };
}
