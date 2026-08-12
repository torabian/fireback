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
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action webPushConfigCreate
 */
export type WebPushConfigCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type WebPushConfigCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WebPushConfigCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WebPushConfigDto;
  }>;
export const useWebPushConfigCreateAction = (
  options?: WebPushConfigCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WebPushConfigDto) => {
    setCompleteState(false);
    return WebPushConfigCreateAction.Fetch(
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
 * WebPushConfigCreateAction
 */
export class WebPushConfigCreateAction {
  //
  static URL = "/webPushConfig";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WebPushConfigCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WebPushConfigDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WebPushConfigDto>, WebPushConfigDto, unknown>(
      overrideUrl ?? WebPushConfigCreateAction.NewUrl(qs),
      {
        method: WebPushConfigCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<WebPushConfigDto, unknown>,
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
    const res = await WebPushConfigCreateAction.Fetch$(
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
    name: "webPushConfigCreate",
    cliName: "create",
    cliShort: "webPushConfig-c",
    url: "/webPushConfig",
    method: "post",
    description: 'Creates a new "webPushConfig" row.',
    in: {
      dto: "WebPushConfigDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WebPushConfigDto",
    },
  };
}
