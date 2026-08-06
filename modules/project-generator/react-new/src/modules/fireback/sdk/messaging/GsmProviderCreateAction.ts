import { GResponse } from "../sdk/envelopes/index";
import { GsmProviderDto } from "./GsmProviderDto";
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
 * Action to communicate with the action gsmProviderCreate
 */
export type GsmProviderCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type GsmProviderCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  GsmProviderCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => GsmProviderDto;
  }>;
export const useGsmProviderCreateAction = (
  options?: GsmProviderCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: GsmProviderDto) => {
    setCompleteState(false);
    return GsmProviderCreateAction.Fetch(
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
 * GsmProviderCreateAction
 */
export class GsmProviderCreateAction {
  //
  static URL = "/gsmProvider";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(GsmProviderCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<GsmProviderDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<GsmProviderDto>, GsmProviderDto, unknown>(
      overrideUrl ?? GsmProviderCreateAction.NewUrl(qs),
      {
        method: GsmProviderCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<GsmProviderDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => GsmProviderDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new GsmProviderDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new GsmProviderDto(item));
    const res = await GsmProviderCreateAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<GsmProviderDto>();
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
    name: "gsmProviderCreate",
    cliShort: "gsmProvider-c",
    url: "/gsmProvider",
    method: "post",
    description: 'Creates a new "gsmProvider" row.',
    in: {
      dto: "GsmProviderDto",
    },
    out: {
      envelope: "GResponse",
      dto: "GsmProviderDto",
    },
  };
}
