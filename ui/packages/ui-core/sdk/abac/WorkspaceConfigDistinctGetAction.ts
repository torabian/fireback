import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WorkspaceConfigDto } from "./WorkspaceConfigDto";
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
 * Action to communicate with the action WorkspaceConfigDistinctGet
 */
export type WorkspaceConfigDistinctGetActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type WorkspaceConfigDistinctGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<WorkspaceConfigDto>, unknown[]>,
  "queryKey"
> &
  WorkspaceConfigDistinctGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WorkspaceConfigDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWorkspaceConfigDistinctGetActionQuery = (
  options: WorkspaceConfigDistinctGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WorkspaceConfigDistinctGetAction.Fetch(
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
    queryKey: [WorkspaceConfigDistinctGetAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WorkspaceConfigDistinctGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceConfigDistinctGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceConfigDto;
  }>;
export const useWorkspaceConfigDistinctGetAction = (
  options?: WorkspaceConfigDistinctGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WorkspaceConfigDistinctGetAction.Fetch(
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
 * WorkspaceConfigDistinctGetAction
 */
export class WorkspaceConfigDistinctGetAction {
  //
  static URL = "/workspace-config/distinct";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WorkspaceConfigDistinctGetAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WorkspaceConfigDto>, unknown, unknown>(
      overrideUrl ?? WorkspaceConfigDistinctGetAction.NewUrl(qs),
      {
        method: WorkspaceConfigDistinctGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
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
    const res = await WorkspaceConfigDistinctGetAction.Fetch$(
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
    name: "WorkspaceConfigDistinctGet",
    cliName: "workspace-config-distinct-get",
    url: "/workspace-config/distinct",
    method: "get",
    description:
      "Returns the single WorkspaceConfig row for the caller's current workspace, looked up by workspace instead of uniqueId - if none exists yet, returns an empty WorkspaceConfig instead of a 404. Restores the old Module3-generated GET /workspace-config/distinct route.",
    out: {
      envelope: "GResponse",
      dto: "WorkspaceConfigDto",
    },
  };
}
