import { GResponse } from "../sdk/envelopes/index";
import { PreferenceDto } from "./PreferenceDto";
import { PreferenceOptionalDto } from "./PreferenceOptionalDto";
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
 * Action to communicate with the action preferenceUpdate
 */
export type PreferenceUpdateActionOptions = {
  queryKey?: unknown[];
  params: PreferenceUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type PreferenceUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PreferenceUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PreferenceDto;
  }>;
export const usePreferenceUpdateAction = (
  options: PreferenceUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PreferenceOptionalDto) => {
    setCompleteState(false);
    return PreferenceUpdateAction.Fetch(
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
 * Path parameters for PreferenceUpdateAction
 */
export type PreferenceUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * PreferenceUpdateAction
 */
export class PreferenceUpdateAction {
  //
  static URL = "/preference/:uniqueId";
  static NewUrl = (
    params: PreferenceUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PreferenceUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: PreferenceUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PreferenceOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<PreferenceDto>, PreferenceOptionalDto, unknown>(
      overrideUrl ?? PreferenceUpdateAction.NewUrl(params, qs),
      {
        method: PreferenceUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PreferenceUpdateActionPathParameter,
    init?: TypedRequestInit<PreferenceOptionalDto, unknown>,
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
    const res = await PreferenceUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
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
    name: "preferenceUpdate",
    cliShort: "preference-u",
    url: "/preference/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "preference" row by uniqueId.',
    in: {
      dto: "PreferenceOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PreferenceDto",
    },
  };
}
