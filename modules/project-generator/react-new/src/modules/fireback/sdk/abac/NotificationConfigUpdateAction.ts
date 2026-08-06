import { GResponse } from "../sdk/envelopes/index";
import { NotificationConfigDto } from "./NotificationConfigDto";
import { NotificationConfigOptionalDto } from "./NotificationConfigOptionalDto";
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
 * Action to communicate with the action notificationConfigUpdate
 */
export type NotificationConfigUpdateActionOptions = {
  queryKey?: unknown[];
  params: NotificationConfigUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type NotificationConfigUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  NotificationConfigUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => NotificationConfigDto;
  }>;
export const useNotificationConfigUpdateAction = (
  options: NotificationConfigUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: NotificationConfigOptionalDto) => {
    setCompleteState(false);
    return NotificationConfigUpdateAction.Fetch(
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
 * Path parameters for NotificationConfigUpdateAction
 */
export type NotificationConfigUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * NotificationConfigUpdateAction
 */
export class NotificationConfigUpdateAction {
  //
  static URL = "/notificationConfig/:uniqueId";
  static NewUrl = (
    params: NotificationConfigUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(NotificationConfigUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: NotificationConfigUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<NotificationConfigOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<NotificationConfigDto>,
      NotificationConfigOptionalDto,
      unknown
    >(
      overrideUrl ?? NotificationConfigUpdateAction.NewUrl(params, qs),
      {
        method: NotificationConfigUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: NotificationConfigUpdateActionPathParameter,
    init?: TypedRequestInit<NotificationConfigOptionalDto, unknown>,
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
    const res = await NotificationConfigUpdateAction.Fetch$(
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
    name: "notificationConfigUpdate",
    cliShort: "notificationConfig-u",
    url: "/notificationConfig/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "notificationConfig" row by uniqueId.',
    in: {
      dto: "NotificationConfigOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "NotificationConfigDto",
    },
  };
}
