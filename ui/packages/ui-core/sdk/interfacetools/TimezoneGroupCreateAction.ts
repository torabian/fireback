import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { TimezoneGroupDto } from "./TimezoneGroupDto";
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
 * Action to communicate with the action timezoneGroupCreate
 */
export type TimezoneGroupCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type TimezoneGroupCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  TimezoneGroupCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => TimezoneGroupDto;
  }>;
export const useTimezoneGroupCreateAction = (
  options?: TimezoneGroupCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: TimezoneGroupDto) => {
    setCompleteState(false);
    return TimezoneGroupCreateAction.Fetch(
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
 * TimezoneGroupCreateAction
 */
export class TimezoneGroupCreateAction {
  //
  static URL = "/timezoneGroup";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(TimezoneGroupCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<TimezoneGroupDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<TimezoneGroupDto>, TimezoneGroupDto, unknown>(
      overrideUrl ?? TimezoneGroupCreateAction.NewUrl(qs),
      {
        method: TimezoneGroupCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<TimezoneGroupDto, unknown>,
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
    const res = await TimezoneGroupCreateAction.Fetch$(
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
    name: "timezoneGroupCreate",
    cliName: "create",
    cliShort: "c",
    url: "/timezoneGroup",
    method: "post",
    description: 'Creates a new "timezoneGroup" row.',
    in: {
      dto: "TimezoneGroupDto",
    },
    out: {
      envelope: "GResponse",
      dto: "TimezoneGroupDto",
    },
  };
}
