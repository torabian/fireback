import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { NotificationDto } from "./NotificationDto";
import { NotificationOptionalDto } from "./NotificationOptionalDto";
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
 * Action to communicate with the action notificationUpdate
 */
export type NotificationUpdateActionOptions = {
  queryKey?: unknown[];
  params: NotificationUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type NotificationUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  NotificationUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => NotificationDto;
  }>;
export const useNotificationUpdateAction = (
  options: NotificationUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: NotificationOptionalDto) => {
    setCompleteState(false);
    return NotificationUpdateAction.Fetch(
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
 * Path parameters for NotificationUpdateAction
 */
export type NotificationUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * NotificationUpdateAction
 */
export class NotificationUpdateAction {
  //
  static URL = "/notification/:uniqueId";
  static NewUrl = (
    params: NotificationUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(NotificationUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: NotificationUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<NotificationOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<NotificationDto>, NotificationOptionalDto, unknown>(
      overrideUrl ?? NotificationUpdateAction.NewUrl(params, qs),
      {
        method: NotificationUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: NotificationUpdateActionPathParameter,
    init?: TypedRequestInit<NotificationOptionalDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => NotificationDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new NotificationDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new NotificationDto(item));
    const res = await NotificationUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<NotificationDto>();
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
    name: "notificationUpdate",
    cliName: "update",
    cliShort: "u",
    url: "/notification/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "notification" row by uniqueId.',
    in: {
      dto: "NotificationOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "NotificationDto",
    },
  };
}
