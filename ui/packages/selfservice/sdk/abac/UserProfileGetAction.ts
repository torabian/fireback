import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { UserProfileDto } from "./UserProfileDto";
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
 * Action to communicate with the action userProfileGet
 */
export type UserProfileGetActionOptions = {
  queryKey?: unknown[];
  params: UserProfileGetActionPathParameter;
  qs?: URLSearchParams;
};
export type UserProfileGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<UserProfileDto>, unknown[]>,
  "queryKey"
> &
  UserProfileGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => UserProfileDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useUserProfileGetActionQuery = (
  options: UserProfileGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return UserProfileGetAction.Fetch(
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
    queryKey: [UserProfileGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type UserProfileGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  UserProfileGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => UserProfileDto;
  }>;
export const useUserProfileGetAction = (
  options: UserProfileGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return UserProfileGetAction.Fetch(
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
 * Path parameters for UserProfileGetAction
 */
export type UserProfileGetActionPathParameter = {
  uniqueId: string;
};
/**
 * UserProfileGetAction
 */
export class UserProfileGetAction {
  //
  static URL = "/userProfile/:uniqueId";
  static NewUrl = (
    params: UserProfileGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(UserProfileGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: UserProfileGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<UserProfileDto>, unknown, unknown>(
      overrideUrl ?? UserProfileGetAction.NewUrl(params, qs),
      {
        method: UserProfileGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: UserProfileGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => UserProfileDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new UserProfileDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new UserProfileDto(item));
    const res = await UserProfileGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<UserProfileDto>();
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
    name: "userProfileGet",
    cliName: "get",
    cliShort: "g",
    url: "/userProfile/:uniqueId string",
    method: "get",
    description: 'Looks up a single "userProfile" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "UserProfileDto",
    },
  };
}
