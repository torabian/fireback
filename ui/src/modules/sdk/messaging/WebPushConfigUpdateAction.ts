import { GResponse } from "../sdk/envelopes/index";
import { WebPushConfigDto } from "./WebPushConfigDto";
import { WebPushConfigOptionalDto } from "./WebPushConfigOptionalDto";
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
 * Action to communicate with the action webPushConfigUpdate
 */
export type WebPushConfigUpdateActionOptions = {
  queryKey?: unknown[];
  params: WebPushConfigUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type WebPushConfigUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WebPushConfigUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WebPushConfigDto;
  }>;
export const useWebPushConfigUpdateAction = (
  options: WebPushConfigUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WebPushConfigOptionalDto) => {
    setCompleteState(false);
    return WebPushConfigUpdateAction.Fetch(
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
 * Path parameters for WebPushConfigUpdateAction
 */
export type WebPushConfigUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * WebPushConfigUpdateAction
 */
export class WebPushConfigUpdateAction {
  //
  static URL = "/webPushConfig/:uniqueId";
  static NewUrl = (
    params: WebPushConfigUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WebPushConfigUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: WebPushConfigUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WebPushConfigOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WebPushConfigDto>,
      WebPushConfigOptionalDto,
      unknown
    >(
      overrideUrl ?? WebPushConfigUpdateAction.NewUrl(params, qs),
      {
        method: WebPushConfigUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WebPushConfigUpdateActionPathParameter,
    init?: TypedRequestInit<WebPushConfigOptionalDto, unknown>,
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
    const res = await WebPushConfigUpdateAction.Fetch$(
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
    name: "webPushConfigUpdate",
    cliName: "update",
    cliShort: "webPushConfig-u",
    url: "/webPushConfig/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "webPushConfig" row by uniqueId.',
    in: {
      dto: "WebPushConfigOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WebPushConfigDto",
    },
  };
}
