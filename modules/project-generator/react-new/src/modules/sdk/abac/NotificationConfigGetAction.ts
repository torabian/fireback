import { GResponse } from "../sdk/envelopes/index";
import { NotificationConfigDto } from "./NotificationConfigDto";
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
 * Action to communicate with the action notificationConfigGet
 */
export type NotificationConfigGetActionOptions = {
  queryKey?: unknown[];
  params: NotificationConfigGetActionPathParameter;
  qs?: URLSearchParams;
};
export type NotificationConfigGetActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<NotificationConfigDto>,
    unknown[]
  >,
  "queryKey"
> &
  NotificationConfigGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => NotificationConfigDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useNotificationConfigGetActionQuery = (
  options: NotificationConfigGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return NotificationConfigGetAction.Fetch(
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
    queryKey: [NotificationConfigGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type NotificationConfigGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  NotificationConfigGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => NotificationConfigDto;
  }>;
export const useNotificationConfigGetAction = (
  options: NotificationConfigGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return NotificationConfigGetAction.Fetch(
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
 * Path parameters for NotificationConfigGetAction
 */
export type NotificationConfigGetActionPathParameter = {
  uniqueId: string;
};
/**
 * NotificationConfigGetAction
 */
export class NotificationConfigGetAction {
  //
  static URL = "/notificationConfig/:uniqueId";
  static NewUrl = (
    params: NotificationConfigGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(NotificationConfigGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: NotificationConfigGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<NotificationConfigDto>, unknown, unknown>(
      overrideUrl ?? NotificationConfigGetAction.NewUrl(params, qs),
      {
        method: NotificationConfigGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: NotificationConfigGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => NotificationConfigDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new NotificationConfigDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new NotificationConfigDto(item));
    const res = await NotificationConfigGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<NotificationConfigDto>();
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
    name: "notificationConfigGet",
    cliShort: "notificationConfig-g",
    url: "/notificationConfig/:uniqueId string",
    method: "get",
    description: 'Looks up a single "notificationConfig" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "NotificationConfigDto",
    },
  };
}
