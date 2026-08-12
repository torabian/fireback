import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WorkspaceConfigDto } from "./WorkspaceConfigDto";
import { WorkspaceConfigOptionalDto } from "./WorkspaceConfigOptionalDto";
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
 * Action to communicate with the action WorkspaceConfigDistinctUpdate
 */
export type WorkspaceConfigDistinctUpdateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type WorkspaceConfigDistinctUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceConfigDistinctUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceConfigDto;
  }>;
export const useWorkspaceConfigDistinctUpdateAction = (
  options?: WorkspaceConfigDistinctUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WorkspaceConfigOptionalDto) => {
    setCompleteState(false);
    return WorkspaceConfigDistinctUpdateAction.Fetch(
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
 * WorkspaceConfigDistinctUpdateAction
 */
export class WorkspaceConfigDistinctUpdateAction {
  //
  static URL = "/workspace-config/distinct";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WorkspaceConfigDistinctUpdateAction.URL, undefined, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WorkspaceConfigOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WorkspaceConfigDto>,
      WorkspaceConfigOptionalDto,
      unknown
    >(
      overrideUrl ?? WorkspaceConfigDistinctUpdateAction.NewUrl(qs),
      {
        method: WorkspaceConfigDistinctUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<WorkspaceConfigOptionalDto, unknown>,
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
    const res = await WorkspaceConfigDistinctUpdateAction.Fetch$(
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
    name: "WorkspaceConfigDistinctUpdate",
    cliName: "workspace-config-distinct-update",
    url: "/workspace-config/distinct",
    method: "patch",
    description:
      "Creates or updates the single WorkspaceConfig row for the caller's current workspace, resolved/written by workspace instead of a uniqueId path param. Restores the old Module3-generated PATCH /workspace-config/distinct route.",
    in: {
      dto: "WorkspaceConfigOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WorkspaceConfigDto",
    },
  };
}
