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
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action preferenceGet
 */
export type PreferenceGetActionOptions = {
  queryKey?: unknown[];
  params: PreferenceGetActionPathParameter;
  qs?: URLSearchParams;
};
export type PreferenceGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<PreferenceDto>, unknown[]>,
  "queryKey"
> &
  PreferenceGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => PreferenceDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const usePreferenceGetActionQuery = (
  options: PreferenceGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return PreferenceGetAction.Fetch(
      options.params,
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
    queryKey: [PreferenceGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type PreferenceGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PreferenceGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PreferenceDto;
  }>;
export const usePreferenceGetAction = (
  options: PreferenceGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return PreferenceGetAction.Fetch(
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
 * Path parameters for PreferenceGetAction
 */
export type PreferenceGetActionPathParameter = {
  uniqueId: string;
};
/**
 * PreferenceGetAction
 */
export class PreferenceGetAction {
  //
  static URL = "/preference/:uniqueId";
  static NewUrl = (
    params: PreferenceGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PreferenceGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: PreferenceGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<PreferenceDto>, unknown, unknown>(
      overrideUrl ?? PreferenceGetAction.NewUrl(params, qs),
      {
        method: PreferenceGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PreferenceGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
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
    const res = await PreferenceGetAction.Fetch$(
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
    name: "preferenceGet",
    cliName: "get",
    cliShort: "g",
    url: "/preference/:uniqueId string",
    method: "get",
    description: 'Looks up a single "preference" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "PreferenceDto",
    },
  };
}
