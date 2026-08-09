import { GResponse } from "../sdk/envelopes/index";
import { WorkspaceDto } from "./WorkspaceDto";
import { WorkspaceOptionalDto } from "./WorkspaceOptionalDto";
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
 * Action to communicate with the action workspaceUpdate
 */
export type WorkspaceUpdateActionOptions = {
  queryKey?: unknown[];
  params: WorkspaceUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type WorkspaceUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceDto;
  }>;
export const useWorkspaceUpdateAction = (
  options: WorkspaceUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WorkspaceOptionalDto) => {
    setCompleteState(false);
    return WorkspaceUpdateAction.Fetch(
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
 * Path parameters for WorkspaceUpdateAction
 */
export type WorkspaceUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * WorkspaceUpdateAction
 */
export class WorkspaceUpdateAction {
  //
  static URL = "/workspace/:uniqueId";
  static NewUrl = (
    params: WorkspaceUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WorkspaceUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: WorkspaceUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WorkspaceOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WorkspaceDto>, WorkspaceOptionalDto, unknown>(
      overrideUrl ?? WorkspaceUpdateAction.NewUrl(params, qs),
      {
        method: WorkspaceUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WorkspaceUpdateActionPathParameter,
    init?: TypedRequestInit<WorkspaceOptionalDto, unknown>,
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
    const res = await WorkspaceUpdateAction.Fetch$(
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
    name: "workspaceUpdate",
    cliName: "update",
    cliShort: "workspace-u",
    url: "/workspace/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "workspace" row by uniqueId.',
    in: {
      dto: "WorkspaceOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WorkspaceDto",
    },
  };
}
