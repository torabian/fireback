import { GResponse } from "../sdk/envelopes/index";
import { UserProfileDto } from "./UserProfileDto";
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
 * Action to communicate with the action userProfileCreate
 */
export type UserProfileCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type UserProfileCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  UserProfileCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => UserProfileDto;
  }>;
export const useUserProfileCreateAction = (
  options?: UserProfileCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: UserProfileDto) => {
    setCompleteState(false);
    return UserProfileCreateAction.Fetch(
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
 * UserProfileCreateAction
 */
export class UserProfileCreateAction {
  //
  static URL = "/userProfile";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(UserProfileCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<UserProfileDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<UserProfileDto>, UserProfileDto, unknown>(
      overrideUrl ?? UserProfileCreateAction.NewUrl(qs),
      {
        method: UserProfileCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<UserProfileDto, unknown>,
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
    const res = await UserProfileCreateAction.Fetch$(
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
    name: "userProfileCreate",
    cliName: "create",
    cliShort: "userProfile-c",
    url: "/userProfile",
    method: "post",
    description: 'Creates a new "userProfile" row.',
    in: {
      dto: "UserProfileDto",
    },
    out: {
      envelope: "GResponse",
      dto: "UserProfileDto",
    },
  };
}
