import { GResponse } from "../sdk/envelopes/index";
import { WorkspaceInviteDto } from "./WorkspaceInviteDto";
import { WorkspaceInviteOptionalDto } from "./WorkspaceInviteOptionalDto";
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
 * Action to communicate with the action workspaceInviteUpdate
 */
export type WorkspaceInviteUpdateActionOptions = {
  queryKey?: unknown[];
  params: WorkspaceInviteUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type WorkspaceInviteUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceInviteUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceInviteDto;
  }>;
export const useWorkspaceInviteUpdateAction = (
  options: WorkspaceInviteUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WorkspaceInviteOptionalDto) => {
    setCompleteState(false);
    return WorkspaceInviteUpdateAction.Fetch(
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
 * Path parameters for WorkspaceInviteUpdateAction
 */
export type WorkspaceInviteUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * WorkspaceInviteUpdateAction
 */
export class WorkspaceInviteUpdateAction {
  //
  static URL = "/workspaceInvite/:uniqueId";
  static NewUrl = (
    params: WorkspaceInviteUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WorkspaceInviteUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: WorkspaceInviteUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WorkspaceInviteOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WorkspaceInviteDto>,
      WorkspaceInviteOptionalDto,
      unknown
    >(
      overrideUrl ?? WorkspaceInviteUpdateAction.NewUrl(params, qs),
      {
        method: WorkspaceInviteUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WorkspaceInviteUpdateActionPathParameter,
    init?: TypedRequestInit<WorkspaceInviteOptionalDto, unknown>,
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
    const res = await WorkspaceInviteUpdateAction.Fetch$(
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
    name: "workspaceInviteUpdate",
    cliShort: "workspaceInvite-u",
    url: "/workspaceInvite/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "workspaceInvite" row by uniqueId.',
    in: {
      dto: "WorkspaceInviteOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WorkspaceInviteDto",
    },
  };
}
