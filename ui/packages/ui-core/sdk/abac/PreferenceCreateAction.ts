import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { PreferenceDto } from "./PreferenceDto";
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
 * Action to communicate with the action preferenceCreate
 */
export type PreferenceCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type PreferenceCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PreferenceCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PreferenceDto;
  }>;
export const usePreferenceCreateAction = (
  options?: PreferenceCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PreferenceDto) => {
    setCompleteState(false);
    return PreferenceCreateAction.Fetch(
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
 * PreferenceCreateAction
 */
export class PreferenceCreateAction {
  //
  static URL = "/preference";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(PreferenceCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PreferenceDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<PreferenceDto>, PreferenceDto, unknown>(
      overrideUrl ?? PreferenceCreateAction.NewUrl(qs),
      {
        method: PreferenceCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<PreferenceDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => PreferenceDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new PreferenceDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new PreferenceDto(item));
    const res = await PreferenceCreateAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<PreferenceDto>();
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
    name: "preferenceCreate",
    cliName: "create",
    cliShort: "c",
    url: "/preference",
    method: "post",
    description: 'Creates a new "preference" row.',
    in: {
      dto: "PreferenceDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PreferenceDto",
    },
  };
}
