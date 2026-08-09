import { GResponse } from "../sdk/envelopes/index";
import { UserWorkspaceDto } from "./UserWorkspaceDto";
import { UserWorkspaceOptionalDto } from "./UserWorkspaceOptionalDto";
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
 * Action to communicate with the action userWorkspaceUpdate
 */
export type UserWorkspaceUpdateActionOptions = {
  queryKey?: unknown[];
  params: UserWorkspaceUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type UserWorkspaceUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  UserWorkspaceUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => UserWorkspaceDto;
  }>;
export const useUserWorkspaceUpdateAction = (
  options: UserWorkspaceUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: UserWorkspaceOptionalDto) => {
    setCompleteState(false);
    return UserWorkspaceUpdateAction.Fetch(
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
 * Path parameters for UserWorkspaceUpdateAction
 */
export type UserWorkspaceUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * UserWorkspaceUpdateAction
 */
export class UserWorkspaceUpdateAction {
  //
  static URL = "/userWorkspace/:uniqueId";
  static NewUrl = (
    params: UserWorkspaceUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(UserWorkspaceUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: UserWorkspaceUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<UserWorkspaceOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<UserWorkspaceDto>,
      UserWorkspaceOptionalDto,
      unknown
    >(
      overrideUrl ?? UserWorkspaceUpdateAction.NewUrl(params, qs),
      {
        method: UserWorkspaceUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: UserWorkspaceUpdateActionPathParameter,
    init?: TypedRequestInit<UserWorkspaceOptionalDto, unknown>,
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
    const res = await UserWorkspaceUpdateAction.Fetch$(
      params,
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
    name: "userWorkspaceUpdate",
    cliName: "update",
    cliShort: "userWorkspace-u",
    url: "/userWorkspace/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "userWorkspace" row by uniqueId.',
    in: {
      dto: "UserWorkspaceOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "UserWorkspaceDto",
    },
  };
}
