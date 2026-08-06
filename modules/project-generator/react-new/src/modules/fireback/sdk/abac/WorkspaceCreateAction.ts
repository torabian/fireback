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
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action workspaceCreate
 */
export type WorkspaceCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type WorkspaceCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceDto;
  }>;
export const useWorkspaceCreateAction = (
  options?: WorkspaceCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WorkspaceDto) => {
    setCompleteState(false);
    return WorkspaceCreateAction.Fetch(
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
 * WorkspaceCreateAction
 */
export class WorkspaceCreateAction {
  //
  static URL = "/workspace";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WorkspaceCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WorkspaceDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WorkspaceDto>, WorkspaceDto, unknown>(
      overrideUrl ?? WorkspaceCreateAction.NewUrl(qs),
      {
        method: WorkspaceCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<WorkspaceDto, unknown>,
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
    const res = await WorkspaceCreateAction.Fetch$(qs, ctx, init, overrideUrl);
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
    name: "workspaceCreate",
    cliShort: "workspace-c",
    url: "/workspace",
    method: "post",
    description: 'Creates a new "workspace" row.',
    in: {
      dto: "WorkspaceDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WorkspaceDto",
    },
  };
}
