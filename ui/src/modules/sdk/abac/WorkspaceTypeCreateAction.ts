import { GResponse } from "../sdk/envelopes/index";
import { WorkspaceTypeDto } from "./WorkspaceTypeDto";
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
 * Action to communicate with the action workspaceTypeCreate
 */
export type WorkspaceTypeCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type WorkspaceTypeCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceTypeCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceTypeDto;
  }>;
export const useWorkspaceTypeCreateAction = (
  options?: WorkspaceTypeCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WorkspaceTypeDto) => {
    setCompleteState(false);
    return WorkspaceTypeCreateAction.Fetch(
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
 * WorkspaceTypeCreateAction
 */
export class WorkspaceTypeCreateAction {
  //
  static URL = "/workspaceType";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WorkspaceTypeCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WorkspaceTypeDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WorkspaceTypeDto>, WorkspaceTypeDto, unknown>(
      overrideUrl ?? WorkspaceTypeCreateAction.NewUrl(qs),
      {
        method: WorkspaceTypeCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<WorkspaceTypeDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WorkspaceTypeDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WorkspaceTypeDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WorkspaceTypeDto(item));
    const res = await WorkspaceTypeCreateAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WorkspaceTypeDto>();
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
    name: "workspaceTypeCreate",
    cliName: "create",
    cliShort: "workspaceType-c",
    url: "/workspaceType",
    method: "post",
    description: 'Creates a new "workspaceType" row.',
    in: {
      dto: "WorkspaceTypeDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WorkspaceTypeDto",
    },
  };
}
