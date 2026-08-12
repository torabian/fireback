import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { RegionalContentDto } from "./RegionalContentDto";
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
 * Action to communicate with the action regionalContentGet
 */
export type RegionalContentGetActionOptions = {
  queryKey?: unknown[];
  params: RegionalContentGetActionPathParameter;
  qs?: URLSearchParams;
};
export type RegionalContentGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<RegionalContentDto>, unknown[]>,
  "queryKey"
> &
  RegionalContentGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => RegionalContentDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useRegionalContentGetActionQuery = (
  options: RegionalContentGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return RegionalContentGetAction.Fetch(
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
    queryKey: [RegionalContentGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type RegionalContentGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  RegionalContentGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => RegionalContentDto;
  }>;
export const useRegionalContentGetAction = (
  options: RegionalContentGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return RegionalContentGetAction.Fetch(
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
 * Path parameters for RegionalContentGetAction
 */
export type RegionalContentGetActionPathParameter = {
  uniqueId: string;
};
/**
 * RegionalContentGetAction
 */
export class RegionalContentGetAction {
  //
  static URL = "/regionalContent/:uniqueId";
  static NewUrl = (
    params: RegionalContentGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(RegionalContentGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: RegionalContentGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<RegionalContentDto>, unknown, unknown>(
      overrideUrl ?? RegionalContentGetAction.NewUrl(params, qs),
      {
        method: RegionalContentGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: RegionalContentGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => RegionalContentDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new RegionalContentDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new RegionalContentDto(item));
    const res = await RegionalContentGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<RegionalContentDto>();
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
    name: "regionalContentGet",
    cliName: "get",
    cliShort: "regionalContent-g",
    url: "/regionalContent/:uniqueId string",
    method: "get",
    description: 'Looks up a single "regionalContent" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "RegionalContentDto",
    },
  };
}
