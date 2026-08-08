import { GResponse } from "../sdk/envelopes/index";
import { WorkspaceConfigDto } from "./WorkspaceConfigDto";
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
 * Action to communicate with the action workspaceConfigGet
 */
export type WorkspaceConfigGetActionOptions = {
  queryKey?: unknown[];
  params: WorkspaceConfigGetActionPathParameter;
  qs?: URLSearchParams;
};
export type WorkspaceConfigGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<WorkspaceConfigDto>, unknown[]>,
  "queryKey"
> &
  WorkspaceConfigGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WorkspaceConfigDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWorkspaceConfigGetActionQuery = (
  options: WorkspaceConfigGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WorkspaceConfigGetAction.Fetch(
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
    queryKey: [WorkspaceConfigGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WorkspaceConfigGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceConfigGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceConfigDto;
  }>;
export const useWorkspaceConfigGetAction = (
  options: WorkspaceConfigGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WorkspaceConfigGetAction.Fetch(
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
 * Path parameters for WorkspaceConfigGetAction
 */
export type WorkspaceConfigGetActionPathParameter = {
  uniqueId: string;
};
/**
 * WorkspaceConfigGetAction
 */
export class WorkspaceConfigGetAction {
  //
  static URL = "/workspaceConfig/:uniqueId";
  static NewUrl = (
    params: WorkspaceConfigGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WorkspaceConfigGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: WorkspaceConfigGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WorkspaceConfigDto>, unknown, unknown>(
      overrideUrl ?? WorkspaceConfigGetAction.NewUrl(params, qs),
      {
        method: WorkspaceConfigGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WorkspaceConfigGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WorkspaceConfigDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WorkspaceConfigDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WorkspaceConfigDto(item));
    const res = await WorkspaceConfigGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WorkspaceConfigDto>();
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
    name: "workspaceConfigGet",
    cliShort: "workspaceConfig-g",
    url: "/workspaceConfig/:uniqueId string",
    method: "get",
    description: 'Looks up a single "workspaceConfig" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "WorkspaceConfigDto",
    },
  };
}
