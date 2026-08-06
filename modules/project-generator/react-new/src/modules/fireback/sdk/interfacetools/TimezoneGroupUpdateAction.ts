import { GResponse } from "../sdk/envelopes/index";
import { TimezoneGroupDto } from "./TimezoneGroupDto";
import { TimezoneGroupOptionalDto } from "./TimezoneGroupOptionalDto";
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
 * Action to communicate with the action timezoneGroupUpdate
 */
export type TimezoneGroupUpdateActionOptions = {
  queryKey?: unknown[];
  params: TimezoneGroupUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type TimezoneGroupUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  TimezoneGroupUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => TimezoneGroupDto;
  }>;
export const useTimezoneGroupUpdateAction = (
  options: TimezoneGroupUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: TimezoneGroupOptionalDto) => {
    setCompleteState(false);
    return TimezoneGroupUpdateAction.Fetch(
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
 * Path parameters for TimezoneGroupUpdateAction
 */
export type TimezoneGroupUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * TimezoneGroupUpdateAction
 */
export class TimezoneGroupUpdateAction {
  //
  static URL = "/timezoneGroup/:uniqueId";
  static NewUrl = (
    params: TimezoneGroupUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(TimezoneGroupUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: TimezoneGroupUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<TimezoneGroupOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<TimezoneGroupDto>,
      TimezoneGroupOptionalDto,
      unknown
    >(
      overrideUrl ?? TimezoneGroupUpdateAction.NewUrl(params, qs),
      {
        method: TimezoneGroupUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: TimezoneGroupUpdateActionPathParameter,
    init?: TypedRequestInit<TimezoneGroupOptionalDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => TimezoneGroupDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new TimezoneGroupDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new TimezoneGroupDto(item));
    const res = await TimezoneGroupUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<TimezoneGroupDto>();
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
    name: "timezoneGroupUpdate",
    cliShort: "timezoneGroup-u",
    url: "/timezoneGroup/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "timezoneGroup" row by uniqueId.',
    in: {
      dto: "TimezoneGroupOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "TimezoneGroupDto",
    },
  };
}
