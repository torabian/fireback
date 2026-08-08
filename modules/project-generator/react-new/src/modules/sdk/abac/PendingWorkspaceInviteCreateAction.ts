import { GResponse } from "../sdk/envelopes/index";
import { PendingWorkspaceInviteDto } from "./PendingWorkspaceInviteDto";
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
 * Action to communicate with the action pendingWorkspaceInviteCreate
 */
export type PendingWorkspaceInviteCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type PendingWorkspaceInviteCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PendingWorkspaceInviteCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PendingWorkspaceInviteDto;
  }>;
export const usePendingWorkspaceInviteCreateAction = (
  options?: PendingWorkspaceInviteCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PendingWorkspaceInviteDto) => {
    setCompleteState(false);
    return PendingWorkspaceInviteCreateAction.Fetch(
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
 * PendingWorkspaceInviteCreateAction
 */
export class PendingWorkspaceInviteCreateAction {
  //
  static URL = "/pendingWorkspaceInvite";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(PendingWorkspaceInviteCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PendingWorkspaceInviteDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<PendingWorkspaceInviteDto>,
      PendingWorkspaceInviteDto,
      unknown
    >(
      overrideUrl ?? PendingWorkspaceInviteCreateAction.NewUrl(qs),
      {
        method: PendingWorkspaceInviteCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<PendingWorkspaceInviteDto, unknown>,
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
    const res = await PendingWorkspaceInviteCreateAction.Fetch$(
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
    name: "pendingWorkspaceInviteCreate",
    cliShort: "pendingWorkspaceInvite-c",
    url: "/pendingWorkspaceInvite",
    method: "post",
    description: 'Creates a new "pendingWorkspaceInvite" row.',
    in: {
      dto: "PendingWorkspaceInviteDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PendingWorkspaceInviteDto",
    },
  };
}
