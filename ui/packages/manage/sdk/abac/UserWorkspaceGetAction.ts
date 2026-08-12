import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { UserWorkspaceDto } from "./UserWorkspaceDto";
import { buildUrl } from "@fireback/js-remote-ctx/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
  type TypedResponse,
} from "@fireback/js-remote-ctx/common/fetchx";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action userWorkspaceGet
 */
export type UserWorkspaceGetActionOptions = {
  queryKey?: unknown[];
  params: UserWorkspaceGetActionPathParameter;
  qs?: URLSearchParams;
};
export type UserWorkspaceGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<UserWorkspaceDto>, unknown[]>,
  "queryKey"
> &
  UserWorkspaceGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => UserWorkspaceDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useUserWorkspaceGetActionQuery = (
  options: UserWorkspaceGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return UserWorkspaceGetAction.Fetch(
      options.params,
      {
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
  const result = useQuery({
    queryKey: [UserWorkspaceGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type UserWorkspaceGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  UserWorkspaceGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => UserWorkspaceDto;
  }>;
export const useUserWorkspaceGetAction = (
  options: UserWorkspaceGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return UserWorkspaceGetAction.Fetch(
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
 * Path parameters for UserWorkspaceGetAction
 */
export type UserWorkspaceGetActionPathParameter = {
  uniqueId: string;
};
/**
 * UserWorkspaceGetAction
 */
export class UserWorkspaceGetAction {
  //
  static URL = "/userWorkspace/:uniqueId";
  static NewUrl = (
    params: UserWorkspaceGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(UserWorkspaceGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: UserWorkspaceGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<UserWorkspaceDto>, unknown, unknown>(
      overrideUrl ?? UserWorkspaceGetAction.NewUrl(params, qs),
      {
        method: UserWorkspaceGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: UserWorkspaceGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
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
    const res = await UserWorkspaceGetAction.Fetch$(
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
    name: "userWorkspaceGet",
    cliName: "get",
    cliShort: "userWorkspace-g",
    url: "/userWorkspace/:uniqueId string",
    method: "get",
    description: 'Looks up a single "userWorkspace" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "UserWorkspaceDto",
    },
  };
}
