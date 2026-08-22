import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { UserProfileDto } from "./UserProfileDto";
import { UserProfileOptionalDto } from "./UserProfileOptionalDto";
import { buildUrl } from "@fireback/js-remote-ctx/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
  type TypedResponse,
} from "@fireback/js-remote-ctx/common/fetchx";
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action userProfileUpdate
 */
export type UserProfileUpdateActionOptions = {
  queryKey?: unknown[];
  params: UserProfileUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type UserProfileUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  UserProfileUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => UserProfileDto;
  }>;
export const useUserProfileUpdateAction = (
  options: UserProfileUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: UserProfileOptionalDto) => {
    setCompleteState(false);
    return UserProfileUpdateAction.Fetch(
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
 * Path parameters for UserProfileUpdateAction
 */
export type UserProfileUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * UserProfileUpdateAction
 */
export class UserProfileUpdateAction {
  //
  static URL = "/userProfile/:uniqueId";
  static NewUrl = (
    params: UserProfileUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(UserProfileUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: UserProfileUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<UserProfileOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<UserProfileDto>, UserProfileOptionalDto, unknown>(
      overrideUrl ?? UserProfileUpdateAction.NewUrl(params, qs),
      {
        method: UserProfileUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: UserProfileUpdateActionPathParameter,
    init?: TypedRequestInit<UserProfileOptionalDto, unknown>,
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
    const res = await UserProfileUpdateAction.Fetch$(
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
    name: "userProfileUpdate",
    cliName: "update",
    cliShort: "u",
    url: "/userProfile/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "userProfile" row by uniqueId.',
    in: {
      dto: "UserProfileOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "UserProfileDto",
    },
  };
}
