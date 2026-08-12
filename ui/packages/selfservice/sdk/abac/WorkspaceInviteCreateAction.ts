import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WorkspaceInviteDto } from "./WorkspaceInviteDto";
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
 * Action to communicate with the action workspaceInviteCreate
 */
export type WorkspaceInviteCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type WorkspaceInviteCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceInviteCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceInviteDto;
  }>;
export const useWorkspaceInviteCreateAction = (
  options?: WorkspaceInviteCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WorkspaceInviteDto) => {
    setCompleteState(false);
    return WorkspaceInviteCreateAction.Fetch(
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
 * WorkspaceInviteCreateAction
 */
export class WorkspaceInviteCreateAction {
  //
  static URL = "/workspaceInvite";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WorkspaceInviteCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WorkspaceInviteDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WorkspaceInviteDto>, WorkspaceInviteDto, unknown>(
      overrideUrl ?? WorkspaceInviteCreateAction.NewUrl(qs),
      {
        method: WorkspaceInviteCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<WorkspaceInviteDto, unknown>,
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
    const res = await WorkspaceInviteCreateAction.Fetch$(
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
    name: "workspaceInviteCreate",
    cliName: "create",
    cliShort: "workspaceInvite-c",
    url: "/workspaceInvite",
    method: "post",
    description: 'Creates a new "workspaceInvite" row.',
    in: {
      dto: "WorkspaceInviteDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WorkspaceInviteDto",
    },
  };
}
