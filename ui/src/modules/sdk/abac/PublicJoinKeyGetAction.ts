import { GResponse } from "../sdk/envelopes/index";
import { PublicJoinKeyDto } from "./PublicJoinKeyDto";
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
 * Action to communicate with the action publicJoinKeyGet
 */
export type PublicJoinKeyGetActionOptions = {
  queryKey?: unknown[];
  params: PublicJoinKeyGetActionPathParameter;
  qs?: URLSearchParams;
};
export type PublicJoinKeyGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<PublicJoinKeyDto>, unknown[]>,
  "queryKey"
> &
  PublicJoinKeyGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => PublicJoinKeyDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const usePublicJoinKeyGetActionQuery = (
  options: PublicJoinKeyGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return PublicJoinKeyGetAction.Fetch(
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
    queryKey: [PublicJoinKeyGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type PublicJoinKeyGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PublicJoinKeyGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PublicJoinKeyDto;
  }>;
export const usePublicJoinKeyGetAction = (
  options: PublicJoinKeyGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return PublicJoinKeyGetAction.Fetch(
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
 * Path parameters for PublicJoinKeyGetAction
 */
export type PublicJoinKeyGetActionPathParameter = {
  uniqueId: string;
};
/**
 * PublicJoinKeyGetAction
 */
export class PublicJoinKeyGetAction {
  //
  static URL = "/publicJoinKey/:uniqueId";
  static NewUrl = (
    params: PublicJoinKeyGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PublicJoinKeyGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: PublicJoinKeyGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<PublicJoinKeyDto>, unknown, unknown>(
      overrideUrl ?? PublicJoinKeyGetAction.NewUrl(params, qs),
      {
        method: PublicJoinKeyGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PublicJoinKeyGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => PublicJoinKeyDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new PublicJoinKeyDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new PublicJoinKeyDto(item));
    const res = await PublicJoinKeyGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<PublicJoinKeyDto>();
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
    name: "publicJoinKeyGet",
    cliName: "get",
    cliShort: "publicJoinKey-g",
    url: "/publicJoinKey/:uniqueId string",
    method: "get",
    description: 'Looks up a single "publicJoinKey" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "PublicJoinKeyDto",
    },
  };
}
