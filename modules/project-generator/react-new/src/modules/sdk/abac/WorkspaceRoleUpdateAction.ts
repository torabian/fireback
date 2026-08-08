import { GResponse } from "../sdk/envelopes/index";
import { WorkspaceRoleDto } from "./WorkspaceRoleDto";
import { WorkspaceRoleOptionalDto } from "./WorkspaceRoleOptionalDto";
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
 * Action to communicate with the action workspaceRoleUpdate
 */
export type WorkspaceRoleUpdateActionOptions = {
  queryKey?: unknown[];
  params: WorkspaceRoleUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type WorkspaceRoleUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceRoleUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceRoleDto;
  }>;
export const useWorkspaceRoleUpdateAction = (
  options: WorkspaceRoleUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WorkspaceRoleOptionalDto) => {
    setCompleteState(false);
    return WorkspaceRoleUpdateAction.Fetch(
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
 * Path parameters for WorkspaceRoleUpdateAction
 */
export type WorkspaceRoleUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * WorkspaceRoleUpdateAction
 */
export class WorkspaceRoleUpdateAction {
  //
  static URL = "/workspaceRole/:uniqueId";
  static NewUrl = (
    params: WorkspaceRoleUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WorkspaceRoleUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: WorkspaceRoleUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WorkspaceRoleOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WorkspaceRoleDto>,
      WorkspaceRoleOptionalDto,
      unknown
    >(
      overrideUrl ?? WorkspaceRoleUpdateAction.NewUrl(params, qs),
      {
        method: WorkspaceRoleUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WorkspaceRoleUpdateActionPathParameter,
    init?: TypedRequestInit<WorkspaceRoleOptionalDto, unknown>,
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
    const res = await WorkspaceRoleUpdateAction.Fetch$(
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
    name: "workspaceRoleUpdate",
    cliShort: "workspaceRole-u",
    url: "/workspaceRole/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "workspaceRole" row by uniqueId.',
    in: {
      dto: "WorkspaceRoleOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WorkspaceRoleDto",
    },
  };
}
