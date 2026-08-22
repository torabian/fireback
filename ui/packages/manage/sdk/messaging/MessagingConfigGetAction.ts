import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { MessagingConfigDto } from "./MessagingConfigDto";
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
 * Action to communicate with the action MessagingConfigGet
 */
export type MessagingConfigGetActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type MessagingConfigGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<MessagingConfigDto>, unknown[]>,
  "queryKey"
> &
  MessagingConfigGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => MessagingConfigDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useMessagingConfigGetActionQuery = (
  options: MessagingConfigGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return MessagingConfigGetAction.Fetch(
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
    queryKey: [MessagingConfigGetAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type MessagingConfigGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  MessagingConfigGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => MessagingConfigDto;
  }>;
export const useMessagingConfigGetAction = (
  options?: MessagingConfigGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return MessagingConfigGetAction.Fetch(
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
 * MessagingConfigGetAction
 */
export class MessagingConfigGetAction {
  //
  static URL = "/messaging/config";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(MessagingConfigGetAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<MessagingConfigDto>, unknown, unknown>(
      overrideUrl ?? MessagingConfigGetAction.NewUrl(qs),
      {
        method: MessagingConfigGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<unknown, unknown>,
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
    const res = await MessagingConfigGetAction.Fetch$(
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
    name: "MessagingConfigGet",
    cliName: "get",
    url: "/messaging/config",
    method: "get",
    description:
      "Returns the single, global MessagingConfig row (there is exactly one for the whole installation, not one per workspace) - if none exists yet, returns an empty MessagingConfig instead of a 404.",
    out: {
      envelope: "GResponse",
      dto: "MessagingConfigDto",
    },
  };
}
