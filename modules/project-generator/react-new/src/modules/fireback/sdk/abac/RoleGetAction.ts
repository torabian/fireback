import { GResponse } from "../sdk/envelopes/index";
import { RoleDto } from "./RoleDto";
import { buildUrl } from "../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
  type TypedResponse,
} from "../sdk/common/fetchx";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "@tanstack/react-query";
import { useFetchxContext } from "../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action roleGet
 */
export type RoleGetActionOptions = {
  queryKey?: unknown[];
  params: RoleGetActionPathParameter;
  qs?: URLSearchParams;
};
export type RoleGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<RoleDto>, unknown[]>,
  "queryKey"
> &
  RoleGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => RoleDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useRoleGetActionQuery = (options: RoleGetActionQueryOptions) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return RoleGetAction.Fetch(
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
    queryKey: [RoleGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type RoleGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  RoleGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => RoleDto;
  }>;
export const useRoleGetAction = (options: RoleGetActionMutationOptions) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return RoleGetAction.Fetch(
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
 * Path parameters for RoleGetAction
 */
export type RoleGetActionPathParameter = {
  uniqueId: string;
};
/**
 * RoleGetAction
 */
export class RoleGetAction {
  //
  static URL = "/role/:uniqueId";
  static NewUrl = (params: RoleGetActionPathParameter, qs?: URLSearchParams) =>
    buildUrl(RoleGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: RoleGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<RoleDto>, unknown, unknown>(
      overrideUrl ?? RoleGetAction.NewUrl(params, qs),
      {
        method: RoleGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: RoleGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => RoleDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new RoleDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new RoleDto(item));
    const res = await RoleGetAction.Fetch$(params, qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<RoleDto>();
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
    name: "roleGet",
    cliShort: "role-g",
    url: "/role/:uniqueId string",
    method: "get",
    description: 'Looks up a single "role" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "RoleDto",
    },
  };
}
