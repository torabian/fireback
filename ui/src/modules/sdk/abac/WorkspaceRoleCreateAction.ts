import { GResponse } from "../sdk/envelopes/index";
import { WorkspaceRoleDto } from "./WorkspaceRoleDto";
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
 * Action to communicate with the action workspaceRoleCreate
 */
export type WorkspaceRoleCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type WorkspaceRoleCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceRoleCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceRoleDto;
  }>;
export const useWorkspaceRoleCreateAction = (
  options?: WorkspaceRoleCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WorkspaceRoleDto) => {
    setCompleteState(false);
    return WorkspaceRoleCreateAction.Fetch(
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
 * WorkspaceRoleCreateAction
 */
export class WorkspaceRoleCreateAction {
  //
  static URL = "/workspaceRole";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WorkspaceRoleCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WorkspaceRoleDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WorkspaceRoleDto>, WorkspaceRoleDto, unknown>(
      overrideUrl ?? WorkspaceRoleCreateAction.NewUrl(qs),
      {
        method: WorkspaceRoleCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<WorkspaceRoleDto, unknown>,
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
    const res = await WorkspaceRoleCreateAction.Fetch$(
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
    name: "workspaceRoleCreate",
    cliShort: "workspaceRole-c",
    url: "/workspaceRole",
    method: "post",
    description: 'Creates a new "workspaceRole" row.',
    in: {
      dto: "WorkspaceRoleDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WorkspaceRoleDto",
    },
  };
}
