import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { RegionalContentDto } from "./RegionalContentDto";
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
 * Action to communicate with the action regionalContentCreate
 */
export type RegionalContentCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type RegionalContentCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  RegionalContentCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => RegionalContentDto;
  }>;
export const useRegionalContentCreateAction = (
  options?: RegionalContentCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: RegionalContentDto) => {
    setCompleteState(false);
    return RegionalContentCreateAction.Fetch(
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
 * RegionalContentCreateAction
 */
export class RegionalContentCreateAction {
  //
  static URL = "/regionalContent";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(RegionalContentCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<RegionalContentDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<RegionalContentDto>, RegionalContentDto, unknown>(
      overrideUrl ?? RegionalContentCreateAction.NewUrl(qs),
      {
        method: RegionalContentCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<RegionalContentDto, unknown>,
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
    const res = await RegionalContentCreateAction.Fetch$(
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
    name: "regionalContentCreate",
    cliName: "create",
    cliShort: "regionalContent-c",
    url: "/regionalContent",
    method: "post",
    description: 'Creates a new "regionalContent" row.',
    in: {
      dto: "RegionalContentDto",
    },
    out: {
      envelope: "GResponse",
      dto: "RegionalContentDto",
    },
  };
}
