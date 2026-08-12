import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { PassportMethodDto } from "./PassportMethodDto";
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
 * Action to communicate with the action passportMethodGet
 */
export type PassportMethodGetActionOptions = {
  queryKey?: unknown[];
  params: PassportMethodGetActionPathParameter;
  qs?: URLSearchParams;
};
export type PassportMethodGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<PassportMethodDto>, unknown[]>,
  "queryKey"
> &
  PassportMethodGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => PassportMethodDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const usePassportMethodGetActionQuery = (
  options: PassportMethodGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return PassportMethodGetAction.Fetch(
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
    queryKey: [PassportMethodGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type PassportMethodGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PassportMethodGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PassportMethodDto;
  }>;
export const usePassportMethodGetAction = (
  options: PassportMethodGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return PassportMethodGetAction.Fetch(
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
 * Path parameters for PassportMethodGetAction
 */
export type PassportMethodGetActionPathParameter = {
  uniqueId: string;
};
/**
 * PassportMethodGetAction
 */
export class PassportMethodGetAction {
  //
  static URL = "/passportMethod/:uniqueId";
  static NewUrl = (
    params: PassportMethodGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PassportMethodGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: PassportMethodGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<PassportMethodDto>, unknown, unknown>(
      overrideUrl ?? PassportMethodGetAction.NewUrl(params, qs),
      {
        method: PassportMethodGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PassportMethodGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => PassportMethodDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new PassportMethodDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new PassportMethodDto(item));
    const res = await PassportMethodGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<PassportMethodDto>();
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
    name: "passportMethodGet",
    cliName: "get",
    cliShort: "passportMethod-g",
    url: "/passportMethod/:uniqueId string",
    method: "get",
    description: 'Looks up a single "passportMethod" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "PassportMethodDto",
    },
  };
}
