import { GResponse } from "../sdk/envelopes/index";
import { GsmProviderDto } from "./GsmProviderDto";
import { GsmProviderOptionalDto } from "./GsmProviderOptionalDto";
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
 * Action to communicate with the action gsmProviderUpdate
 */
export type GsmProviderUpdateActionOptions = {
  queryKey?: unknown[];
  params: GsmProviderUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type GsmProviderUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  GsmProviderUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => GsmProviderDto;
  }>;
export const useGsmProviderUpdateAction = (
  options: GsmProviderUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: GsmProviderOptionalDto) => {
    setCompleteState(false);
    return GsmProviderUpdateAction.Fetch(
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
 * Path parameters for GsmProviderUpdateAction
 */
export type GsmProviderUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * GsmProviderUpdateAction
 */
export class GsmProviderUpdateAction {
  //
  static URL = "/gsmProvider/:uniqueId";
  static NewUrl = (
    params: GsmProviderUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(GsmProviderUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: GsmProviderUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<GsmProviderOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<GsmProviderDto>, GsmProviderOptionalDto, unknown>(
      overrideUrl ?? GsmProviderUpdateAction.NewUrl(params, qs),
      {
        method: GsmProviderUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: GsmProviderUpdateActionPathParameter,
    init?: TypedRequestInit<GsmProviderOptionalDto, unknown>,
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
    const res = await GsmProviderUpdateAction.Fetch$(
      params,
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
    name: "gsmProviderUpdate",
    cliName: "update",
    cliShort: "gsmProvider-u",
    url: "/gsmProvider/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "gsmProvider" row by uniqueId.',
    in: {
      dto: "GsmProviderOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "GsmProviderDto",
    },
  };
}
