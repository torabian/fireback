import { GResponse } from "../sdk/envelopes/index";
import { PendingWorkspaceInviteDto } from "./PendingWorkspaceInviteDto";
import { PendingWorkspaceInviteOptionalDto } from "./PendingWorkspaceInviteOptionalDto";
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
 * Action to communicate with the action pendingWorkspaceInviteUpdate
 */
export type PendingWorkspaceInviteUpdateActionOptions = {
  queryKey?: unknown[];
  params: PendingWorkspaceInviteUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type PendingWorkspaceInviteUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PendingWorkspaceInviteUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PendingWorkspaceInviteDto;
  }>;
export const usePendingWorkspaceInviteUpdateAction = (
  options: PendingWorkspaceInviteUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PendingWorkspaceInviteOptionalDto) => {
    setCompleteState(false);
    return PendingWorkspaceInviteUpdateAction.Fetch(
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
 * Path parameters for PendingWorkspaceInviteUpdateAction
 */
export type PendingWorkspaceInviteUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * PendingWorkspaceInviteUpdateAction
 */
export class PendingWorkspaceInviteUpdateAction {
  //
  static URL = "/pendingWorkspaceInvite/:uniqueId";
  static NewUrl = (
    params: PendingWorkspaceInviteUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PendingWorkspaceInviteUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: PendingWorkspaceInviteUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PendingWorkspaceInviteOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<PendingWorkspaceInviteDto>,
      PendingWorkspaceInviteOptionalDto,
      unknown
    >(
      overrideUrl ?? PendingWorkspaceInviteUpdateAction.NewUrl(params, qs),
      {
        method: PendingWorkspaceInviteUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PendingWorkspaceInviteUpdateActionPathParameter,
    init?: TypedRequestInit<PendingWorkspaceInviteOptionalDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => PendingWorkspaceInviteDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new PendingWorkspaceInviteDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new PendingWorkspaceInviteDto(item));
    const res = await PendingWorkspaceInviteUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<PendingWorkspaceInviteDto>();
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
    name: "pendingWorkspaceInviteUpdate",
    cliName: "update",
    cliShort: "pendingWorkspaceInvite-u",
    url: "/pendingWorkspaceInvite/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "pendingWorkspaceInvite" row by uniqueId.',
    in: {
      dto: "PendingWorkspaceInviteOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PendingWorkspaceInviteDto",
    },
  };
}
