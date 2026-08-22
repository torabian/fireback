import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { NotificationDto } from "./NotificationDto";
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
 * Action to communicate with the action notificationCreate
 */
export type NotificationCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type NotificationCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  NotificationCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => NotificationDto;
  }>;
export const useNotificationCreateAction = (
  options?: NotificationCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: NotificationDto) => {
    setCompleteState(false);
    return NotificationCreateAction.Fetch(
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
 * NotificationCreateAction
 */
export class NotificationCreateAction {
  //
  static URL = "/notification";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(NotificationCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<NotificationDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<NotificationDto>, NotificationDto, unknown>(
      overrideUrl ?? NotificationCreateAction.NewUrl(qs),
      {
        method: NotificationCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<NotificationDto, unknown>,
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
    const res = await NotificationCreateAction.Fetch$(
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
    name: "notificationCreate",
    cliName: "create",
    cliShort: "c",
    url: "/notification",
    method: "post",
    description: 'Creates a new "notification" row.',
    in: {
      dto: "NotificationDto",
    },
    out: {
      envelope: "GResponse",
      dto: "NotificationDto",
    },
  };
}
