import { GResponse } from "../sdk/envelopes/index";
import { TimezoneGroupDto } from "./TimezoneGroupDto";
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
 * Action to communicate with the action timezoneGroupGet
 */
export type TimezoneGroupGetActionOptions = {
  queryKey?: unknown[];
  params: TimezoneGroupGetActionPathParameter;
  qs?: URLSearchParams;
};
export type TimezoneGroupGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<TimezoneGroupDto>, unknown[]>,
  "queryKey"
> &
  TimezoneGroupGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => TimezoneGroupDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useTimezoneGroupGetActionQuery = (
  options: TimezoneGroupGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return TimezoneGroupGetAction.Fetch(
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
    queryKey: [TimezoneGroupGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type TimezoneGroupGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  TimezoneGroupGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => TimezoneGroupDto;
  }>;
export const useTimezoneGroupGetAction = (
  options: TimezoneGroupGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return TimezoneGroupGetAction.Fetch(
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
 * Path parameters for TimezoneGroupGetAction
 */
export type TimezoneGroupGetActionPathParameter = {
  uniqueId: string;
};
/**
 * TimezoneGroupGetAction
 */
export class TimezoneGroupGetAction {
  //
  static URL = "/timezoneGroup/:uniqueId";
  static NewUrl = (
    params: TimezoneGroupGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(TimezoneGroupGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: TimezoneGroupGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<TimezoneGroupDto>, unknown, unknown>(
      overrideUrl ?? TimezoneGroupGetAction.NewUrl(params, qs),
      {
        method: TimezoneGroupGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: TimezoneGroupGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => TimezoneGroupDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new TimezoneGroupDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new TimezoneGroupDto(item));
    const res = await TimezoneGroupGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<TimezoneGroupDto>();
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
    name: "timezoneGroupGet",
    cliShort: "timezoneGroup-g",
    url: "/timezoneGroup/:uniqueId string",
    method: "get",
    description: 'Looks up a single "timezoneGroup" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "TimezoneGroupDto",
    },
  };
}
