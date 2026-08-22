import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { MessagingConfigDto } from "./MessagingConfigDto";
import { MessagingConfigOptionalDto } from "./MessagingConfigOptionalDto";
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
 * Action to communicate with the action MessagingConfigUpdate
 */
export type MessagingConfigUpdateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type MessagingConfigUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  MessagingConfigUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => MessagingConfigDto;
  }>;
export const useMessagingConfigUpdateAction = (
  options?: MessagingConfigUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: MessagingConfigOptionalDto) => {
    setCompleteState(false);
    return MessagingConfigUpdateAction.Fetch(
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
 * MessagingConfigUpdateAction
 */
export class MessagingConfigUpdateAction {
  //
  static URL = "/messaging/config";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(MessagingConfigUpdateAction.URL, undefined, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<MessagingConfigOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<MessagingConfigDto>,
      MessagingConfigOptionalDto,
      unknown
    >(
      overrideUrl ?? MessagingConfigUpdateAction.NewUrl(qs),
      {
        method: MessagingConfigUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<MessagingConfigOptionalDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => MessagingConfigDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new MessagingConfigDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new MessagingConfigDto(item));
    const res = await MessagingConfigUpdateAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<MessagingConfigDto>();
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
    name: "MessagingConfigUpdate",
    cliName: "update",
    url: "/messaging/config",
    method: "patch",
    description:
      "Creates or updates the single, global MessagingConfig row (there is exactly one for the whole installation, not one per workspace) - the row is created on first write if it doesn't exist yet.",
    in: {
      dto: "MessagingConfigOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "MessagingConfigDto",
    },
  };
}
