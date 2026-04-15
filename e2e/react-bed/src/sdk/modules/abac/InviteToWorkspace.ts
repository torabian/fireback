import { WorkspaceInvitationDto } from "./WorkspaceInvitationDto";
import { buildUrl } from "../../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
  type TypedResponse,
} from "../../sdk/common/fetchx";
import { type UseMutationOptions, useMutation } from "react-query";
import { useFetchxContext } from "../../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action InviteToWorkspace
 */
export type InviteToWorkspaceActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type InviteToWorkspaceActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  InviteToWorkspaceActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WorkspaceInvitationDto;
  }>;
export const useInviteToWorkspaceAction = (
  options?: InviteToWorkspaceActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WorkspaceInvitationDto) => {
    setCompleteState(false);
    return InviteToWorkspaceAction.Fetch(
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
 * InviteToWorkspaceAction
 */
export class InviteToWorkspaceAction {
  //
  static URL = "/workspace/invite";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(InviteToWorkspaceAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<WorkspaceInvitationDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<WorkspaceInvitationDto, WorkspaceInvitationDto, unknown>(
      overrideUrl ?? InviteToWorkspaceAction.NewUrl(qs),
      {
        method: InviteToWorkspaceAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<WorkspaceInvitationDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WorkspaceInvitationDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WorkspaceInvitationDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WorkspaceInvitationDto(item));
    const res = await InviteToWorkspaceAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "InviteToWorkspace",
    cliName: "invite",
    url: "/workspace/invite",
    method: "post",
    description:
      "Invite a new person (either a user, with passport or without passport)",
    in: {
      dto: "WorkspaceInvitationDto",
    },
    out: {
      dto: "WorkspaceInvitationDto",
    },
  };
}
