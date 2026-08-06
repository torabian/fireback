import { GResponse } from "../sdk/envelopes/index";
import { UserDto } from "./UserDto";
import { UserOptionalDto } from "./UserOptionalDto";
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
 * Action to communicate with the action userUpdate
 */
export type UserUpdateActionOptions = {
  queryKey?: unknown[];
  params: UserUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type UserUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  UserUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => UserDto;
  }>;
export const useUserUpdateAction = (
  options: UserUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: UserOptionalDto) => {
    setCompleteState(false);
    return UserUpdateAction.Fetch(
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
 * Path parameters for UserUpdateAction
 */
export type UserUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * UserUpdateAction
 */
export class UserUpdateAction {
  //
  static URL = "/user/:uniqueId";
  static NewUrl = (
    params: UserUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(UserUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: UserUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<UserOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<UserDto>, UserOptionalDto, unknown>(
      overrideUrl ?? UserUpdateAction.NewUrl(params, qs),
      {
        method: UserUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: UserUpdateActionPathParameter,
    init?: TypedRequestInit<UserOptionalDto, unknown>,
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
    const res = await UserUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
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
    name: "userUpdate",
    cliShort: "user-u",
    url: "/user/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "user" row by uniqueId.',
    in: {
      dto: "UserOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "UserDto",
    },
  };
}
