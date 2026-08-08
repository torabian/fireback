import { GResponse } from "../sdk/envelopes/index";
import { UserWorkspaceDto } from "./UserWorkspaceDto";
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
 * Action to communicate with the action userWorkspaceCreate
 */
export type UserWorkspaceCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type UserWorkspaceCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  UserWorkspaceCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => UserWorkspaceDto;
  }>;
export const useUserWorkspaceCreateAction = (
  options?: UserWorkspaceCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: UserWorkspaceDto) => {
    setCompleteState(false);
    return UserWorkspaceCreateAction.Fetch(
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
 * UserWorkspaceCreateAction
 */
export class UserWorkspaceCreateAction {
  //
  static URL = "/userWorkspace";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(UserWorkspaceCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<UserWorkspaceDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<UserWorkspaceDto>, UserWorkspaceDto, unknown>(
      overrideUrl ?? UserWorkspaceCreateAction.NewUrl(qs),
      {
        method: UserWorkspaceCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<UserWorkspaceDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => UserWorkspaceDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new UserWorkspaceDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new UserWorkspaceDto(item));
    const res = await UserWorkspaceCreateAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<UserWorkspaceDto>();
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
    name: "userWorkspaceCreate",
    cliShort: "userWorkspace-c",
    url: "/userWorkspace",
    method: "post",
    description: 'Creates a new "userWorkspace" row.',
    in: {
      dto: "UserWorkspaceDto",
    },
    out: {
      envelope: "GResponse",
      dto: "UserWorkspaceDto",
    },
  };
}
