import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WorkspaceRoleDto } from "./WorkspaceRoleDto";
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
 * Action to communicate with the action workspaceRoleGet
 */
export type WorkspaceRoleGetActionOptions = {
  queryKey?: unknown[];
  params: WorkspaceRoleGetActionPathParameter;
  qs?: URLSearchParams;
};
export type WorkspaceRoleGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<WorkspaceRoleDto>, unknown[]>,
  "queryKey"
> &
  WorkspaceRoleGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WorkspaceRoleDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWorkspaceRoleGetActionQuery = (
  options: WorkspaceRoleGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WorkspaceRoleGetAction.Fetch(
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
    queryKey: [WorkspaceRoleGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WorkspaceRoleGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceRoleGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceRoleDto;
  }>;
export const useWorkspaceRoleGetAction = (
  options: WorkspaceRoleGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WorkspaceRoleGetAction.Fetch(
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
 * Path parameters for WorkspaceRoleGetAction
 */
export type WorkspaceRoleGetActionPathParameter = {
  uniqueId: string;
};
/**
 * WorkspaceRoleGetAction
 */
export class WorkspaceRoleGetAction {
  //
  static URL = "/workspaceRole/:uniqueId";
  static NewUrl = (
    params: WorkspaceRoleGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WorkspaceRoleGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: WorkspaceRoleGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WorkspaceRoleDto>, unknown, unknown>(
      overrideUrl ?? WorkspaceRoleGetAction.NewUrl(params, qs),
      {
        method: WorkspaceRoleGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WorkspaceRoleGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WorkspaceRoleDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WorkspaceRoleDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WorkspaceRoleDto(item));
    const res = await WorkspaceRoleGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WorkspaceRoleDto>();
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
    name: "workspaceRoleGet",
    cliName: "get",
    cliShort: "workspaceRole-g",
    url: "/workspaceRole/:uniqueId string",
    method: "get",
    description: 'Looks up a single "workspaceRole" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "WorkspaceRoleDto",
    },
  };
}
