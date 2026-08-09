import { GResponse } from "../sdk/envelopes/index";
import { WorkspaceDto } from "./WorkspaceDto";
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
 * Action to communicate with the action workspaceGet
 */
export type WorkspaceGetActionOptions = {
  queryKey?: unknown[];
  params: WorkspaceGetActionPathParameter;
  qs?: URLSearchParams;
};
export type WorkspaceGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<WorkspaceDto>, unknown[]>,
  "queryKey"
> &
  WorkspaceGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WorkspaceDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWorkspaceGetActionQuery = (
  options: WorkspaceGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WorkspaceGetAction.Fetch(
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
    queryKey: [WorkspaceGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WorkspaceGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceDto;
  }>;
export const useWorkspaceGetAction = (
  options: WorkspaceGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WorkspaceGetAction.Fetch(
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
 * Path parameters for WorkspaceGetAction
 */
export type WorkspaceGetActionPathParameter = {
  uniqueId: string;
};
/**
 * WorkspaceGetAction
 */
export class WorkspaceGetAction {
  //
  static URL = "/workspace/:uniqueId";
  static NewUrl = (
    params: WorkspaceGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WorkspaceGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: WorkspaceGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WorkspaceDto>, unknown, unknown>(
      overrideUrl ?? WorkspaceGetAction.NewUrl(params, qs),
      {
        method: WorkspaceGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WorkspaceGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WorkspaceDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WorkspaceDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WorkspaceDto(item));
    const res = await WorkspaceGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WorkspaceDto>();
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
    name: "workspaceGet",
    cliName: "get",
    cliShort: "workspace-g",
    url: "/workspace/:uniqueId string",
    method: "get",
    description: 'Looks up a single "workspace" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "WorkspaceDto",
    },
  };
}
