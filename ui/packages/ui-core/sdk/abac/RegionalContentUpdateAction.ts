import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { RegionalContentDto } from "./RegionalContentDto";
import { RegionalContentOptionalDto } from "./RegionalContentOptionalDto";
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
 * Action to communicate with the action regionalContentUpdate
 */
export type RegionalContentUpdateActionOptions = {
  queryKey?: unknown[];
  params: RegionalContentUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type RegionalContentUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  RegionalContentUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => RegionalContentDto;
  }>;
export const useRegionalContentUpdateAction = (
  options: RegionalContentUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: RegionalContentOptionalDto) => {
    setCompleteState(false);
    return RegionalContentUpdateAction.Fetch(
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
 * Path parameters for RegionalContentUpdateAction
 */
export type RegionalContentUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * RegionalContentUpdateAction
 */
export class RegionalContentUpdateAction {
  //
  static URL = "/regionalContent/:uniqueId";
  static NewUrl = (
    params: RegionalContentUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(RegionalContentUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: RegionalContentUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<RegionalContentOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<RegionalContentDto>,
      RegionalContentOptionalDto,
      unknown
    >(
      overrideUrl ?? RegionalContentUpdateAction.NewUrl(params, qs),
      {
        method: RegionalContentUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: RegionalContentUpdateActionPathParameter,
    init?: TypedRequestInit<RegionalContentOptionalDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => RegionalContentDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new RegionalContentDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new RegionalContentDto(item));
    const res = await RegionalContentUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<RegionalContentDto>();
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
    name: "regionalContentUpdate",
    cliName: "update",
    cliShort: "regionalContent-u",
    url: "/regionalContent/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "regionalContent" row by uniqueId.',
    in: {
      dto: "RegionalContentOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "RegionalContentDto",
    },
  };
}
