import { GResponse } from "../sdk/envelopes/index";
import { WorkspaceInviteDto } from "./WorkspaceInviteDto";
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
 * Action to communicate with the action workspaceInviteGet
 */
export type WorkspaceInviteGetActionOptions = {
  queryKey?: unknown[];
  params: WorkspaceInviteGetActionPathParameter;
  qs?: URLSearchParams;
};
export type WorkspaceInviteGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<WorkspaceInviteDto>, unknown[]>,
  "queryKey"
> &
  WorkspaceInviteGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WorkspaceInviteDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWorkspaceInviteGetActionQuery = (
  options: WorkspaceInviteGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WorkspaceInviteGetAction.Fetch(
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
    queryKey: [WorkspaceInviteGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WorkspaceInviteGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceInviteGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceInviteDto;
  }>;
export const useWorkspaceInviteGetAction = (
  options: WorkspaceInviteGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WorkspaceInviteGetAction.Fetch(
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
 * Path parameters for WorkspaceInviteGetAction
 */
export type WorkspaceInviteGetActionPathParameter = {
  uniqueId: string;
};
/**
 * WorkspaceInviteGetAction
 */
export class WorkspaceInviteGetAction {
  //
  static URL = "/workspaceInvite/:uniqueId";
  static NewUrl = (
    params: WorkspaceInviteGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WorkspaceInviteGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: WorkspaceInviteGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WorkspaceInviteDto>, unknown, unknown>(
      overrideUrl ?? WorkspaceInviteGetAction.NewUrl(params, qs),
      {
        method: WorkspaceInviteGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WorkspaceInviteGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WorkspaceInviteDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WorkspaceInviteDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WorkspaceInviteDto(item));
    const res = await WorkspaceInviteGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WorkspaceInviteDto>();
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
    name: "workspaceInviteGet",
    cliShort: "workspaceInvite-g",
    url: "/workspaceInvite/:uniqueId string",
    method: "get",
    description: 'Looks up a single "workspaceInvite" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "WorkspaceInviteDto",
    },
  };
}
