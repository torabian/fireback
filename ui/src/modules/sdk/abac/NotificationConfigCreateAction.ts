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
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action notificationConfigCreate
 */
export type NotificationConfigCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type NotificationConfigCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  NotificationConfigCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => NotificationConfigDto;
  }>;
export const useNotificationConfigCreateAction = (
  options?: NotificationConfigCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: NotificationConfigDto) => {
    setCompleteState(false);
    return NotificationConfigCreateAction.Fetch(
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
 * NotificationConfigCreateAction
 */
export class NotificationConfigCreateAction {
  //
  static URL = "/notificationConfig";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(NotificationConfigCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<NotificationConfigDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<NotificationConfigDto>,
      NotificationConfigDto,
      unknown
    >(
      overrideUrl ?? NotificationConfigCreateAction.NewUrl(qs),
      {
        method: NotificationConfigCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<NotificationConfigDto, unknown>,
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
    const res = await NotificationConfigCreateAction.Fetch$(
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
    name: "notificationConfigCreate",
    cliName: "create",
    cliShort: "notificationConfig-c",
    url: "/notificationConfig",
    method: "post",
    description: 'Creates a new "notificationConfig" row.',
    in: {
      dto: "NotificationConfigDto",
    },
    out: {
      envelope: "GResponse",
      dto: "NotificationConfigDto",
    },
  };
}
