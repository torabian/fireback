import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WorkspaceTypeDto } from "./WorkspaceTypeDto";
import { WorkspaceTypeOptionalDto } from "./WorkspaceTypeOptionalDto";
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
 * Action to communicate with the action workspaceTypeUpdate
 */
export type WorkspaceTypeUpdateActionOptions = {
  queryKey?: unknown[];
  params: WorkspaceTypeUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type WorkspaceTypeUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceTypeUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceTypeDto;
  }>;
export const useWorkspaceTypeUpdateAction = (
  options: WorkspaceTypeUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WorkspaceTypeOptionalDto) => {
    setCompleteState(false);
    return WorkspaceTypeUpdateAction.Fetch(
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
 * Path parameters for WorkspaceTypeUpdateAction
 */
export type WorkspaceTypeUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * WorkspaceTypeUpdateAction
 */
export class WorkspaceTypeUpdateAction {
  //
  static URL = "/workspaceType/:uniqueId";
  static NewUrl = (
    params: WorkspaceTypeUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WorkspaceTypeUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: WorkspaceTypeUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WorkspaceTypeOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WorkspaceTypeDto>,
      WorkspaceTypeOptionalDto,
      unknown
    >(
      overrideUrl ?? WorkspaceTypeUpdateAction.NewUrl(params, qs),
      {
        method: WorkspaceTypeUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WorkspaceTypeUpdateActionPathParameter,
    init?: TypedRequestInit<WorkspaceTypeOptionalDto, unknown>,
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
    const res = await WorkspaceTypeUpdateAction.Fetch$(
      params,
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
    name: "workspaceTypeUpdate",
    cliName: "update",
    cliShort: "workspaceType-u",
    url: "/workspaceType/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "workspaceType" row by uniqueId.',
    in: {
      dto: "WorkspaceTypeOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WorkspaceTypeDto",
    },
  };
}
