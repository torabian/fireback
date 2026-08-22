import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WebPushConfigDto } from "./WebPushConfigDto";
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
 * Action to communicate with the action webPushConfigGet
 */
export type WebPushConfigGetActionOptions = {
  queryKey?: unknown[];
  params: WebPushConfigGetActionPathParameter;
  qs?: URLSearchParams;
};
export type WebPushConfigGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<WebPushConfigDto>, unknown[]>,
  "queryKey"
> &
  WebPushConfigGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WebPushConfigDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWebPushConfigGetActionQuery = (
  options: WebPushConfigGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WebPushConfigGetAction.Fetch(
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
    queryKey: [WebPushConfigGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WebPushConfigGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WebPushConfigGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WebPushConfigDto;
  }>;
export const useWebPushConfigGetAction = (
  options: WebPushConfigGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WebPushConfigGetAction.Fetch(
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
 * Path parameters for WebPushConfigGetAction
 */
export type WebPushConfigGetActionPathParameter = {
  uniqueId: string;
};
/**
 * WebPushConfigGetAction
 */
export class WebPushConfigGetAction {
  //
  static URL = "/webPushConfig/:uniqueId";
  static NewUrl = (
    params: WebPushConfigGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WebPushConfigGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: WebPushConfigGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WebPushConfigDto>, unknown, unknown>(
      overrideUrl ?? WebPushConfigGetAction.NewUrl(params, qs),
      {
        method: WebPushConfigGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WebPushConfigGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WebPushConfigDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WebPushConfigDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WebPushConfigDto(item));
    const res = await WebPushConfigGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WebPushConfigDto>();
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
    name: "webPushConfigGet",
    cliName: "get",
    cliShort: "g",
    url: "/webPushConfig/:uniqueId string",
    method: "get",
    description: 'Looks up a single "webPushConfig" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "WebPushConfigDto",
    },
  };
}
