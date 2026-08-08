import { GResponse } from "../sdk/envelopes/index";
import { WorkspaceConfigDto } from "./WorkspaceConfigDto";
import { WorkspaceConfigOptionalDto } from "./WorkspaceConfigOptionalDto";
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
 * Action to communicate with the action workspaceConfigUpdate
 */
export type WorkspaceConfigUpdateActionOptions = {
  queryKey?: unknown[];
  params: WorkspaceConfigUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type WorkspaceConfigUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceConfigUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceConfigDto;
  }>;
export const useWorkspaceConfigUpdateAction = (
  options: WorkspaceConfigUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WorkspaceConfigOptionalDto) => {
    setCompleteState(false);
    return WorkspaceConfigUpdateAction.Fetch(
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
 * Path parameters for WorkspaceConfigUpdateAction
 */
export type WorkspaceConfigUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * WorkspaceConfigUpdateAction
 */
export class WorkspaceConfigUpdateAction {
  //
  static URL = "/workspaceConfig/:uniqueId";
  static NewUrl = (
    params: WorkspaceConfigUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WorkspaceConfigUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: WorkspaceConfigUpdateActionPathParameter,
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
      overrideUrl ?? WorkspaceConfigUpdateAction.NewUrl(params, qs),
      {
        method: WorkspaceConfigUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WorkspaceConfigUpdateActionPathParameter,
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
    const res = await WorkspaceConfigUpdateAction.Fetch$(
      params,
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
    name: "workspaceConfigUpdate",
    cliShort: "workspaceConfig-u",
    url: "/workspaceConfig/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "workspaceConfig" row by uniqueId.',
    in: {
      dto: "WorkspaceConfigOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WorkspaceConfigDto",
    },
  };
}
