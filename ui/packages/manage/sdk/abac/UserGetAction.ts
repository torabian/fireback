import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { UserDto } from "./UserDto";
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
 * Action to communicate with the action userGet
 */
export type UserGetActionOptions = {
  queryKey?: unknown[];
  params: UserGetActionPathParameter;
  qs?: URLSearchParams;
};
export type UserGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<UserDto>, unknown[]>,
  "queryKey"
> &
  UserGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => UserDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useUserGetActionQuery = (options: UserGetActionQueryOptions) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return UserGetAction.Fetch(
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
    queryKey: [UserGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type UserGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  UserGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => UserDto;
  }>;
export const useUserGetAction = (options: UserGetActionMutationOptions) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return UserGetAction.Fetch(
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
 * Path parameters for UserGetAction
 */
export type UserGetActionPathParameter = {
  uniqueId: string;
};
/**
 * UserGetAction
 */
export class UserGetAction {
  //
  static URL = "/user/:uniqueId";
  static NewUrl = (params: UserGetActionPathParameter, qs?: URLSearchParams) =>
    buildUrl(UserGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: UserGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<UserDto>, unknown, unknown>(
      overrideUrl ?? UserGetAction.NewUrl(params, qs),
      {
        method: UserGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: UserGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => UserDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new UserDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new UserDto(item));
    const res = await UserGetAction.Fetch$(params, qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<UserDto>();
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
    name: "userGet",
    cliName: "get",
    cliShort: "g",
    url: "/user/:uniqueId string",
    method: "get",
    description: 'Looks up a single "user" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "UserDto",
    },
  };
}
