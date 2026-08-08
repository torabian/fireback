import { GResponse } from "../sdk/envelopes/index";
import { GsmProviderDto } from "./GsmProviderDto";
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
 * Action to communicate with the action gsmProviderGet
 */
export type GsmProviderGetActionOptions = {
  queryKey?: unknown[];
  params: GsmProviderGetActionPathParameter;
  qs?: URLSearchParams;
};
export type GsmProviderGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<GsmProviderDto>, unknown[]>,
  "queryKey"
> &
  GsmProviderGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => GsmProviderDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useGsmProviderGetActionQuery = (
  options: GsmProviderGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return GsmProviderGetAction.Fetch(
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
    queryKey: [GsmProviderGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type GsmProviderGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  GsmProviderGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => GsmProviderDto;
  }>;
export const useGsmProviderGetAction = (
  options: GsmProviderGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return GsmProviderGetAction.Fetch(
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
 * Path parameters for GsmProviderGetAction
 */
export type GsmProviderGetActionPathParameter = {
  uniqueId: string;
};
/**
 * GsmProviderGetAction
 */
export class GsmProviderGetAction {
  //
  static URL = "/gsmProvider/:uniqueId";
  static NewUrl = (
    params: GsmProviderGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(GsmProviderGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: GsmProviderGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<GsmProviderDto>, unknown, unknown>(
      overrideUrl ?? GsmProviderGetAction.NewUrl(params, qs),
      {
        method: GsmProviderGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: GsmProviderGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => GsmProviderDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new GsmProviderDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new GsmProviderDto(item));
    const res = await GsmProviderGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<GsmProviderDto>();
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
    name: "gsmProviderGet",
    cliShort: "gsmProvider-g",
    url: "/gsmProvider/:uniqueId string",
    method: "get",
    description: 'Looks up a single "gsmProvider" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "GsmProviderDto",
    },
  };
}
