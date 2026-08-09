import { GResponse } from "../sdk/envelopes/index";
import { RoleDto } from "./RoleDto";
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
 * Action to communicate with the action roleCreate
 */
export type RoleCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type RoleCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  RoleCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => RoleDto;
  }>;
export const useRoleCreateAction = (
  options?: RoleCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: RoleDto) => {
    setCompleteState(false);
    return RoleCreateAction.Fetch(
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
 * RoleCreateAction
 */
export class RoleCreateAction {
  //
  static URL = "/role";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(RoleCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<RoleDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<RoleDto>, RoleDto, unknown>(
      overrideUrl ?? RoleCreateAction.NewUrl(qs),
      {
        method: RoleCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<RoleDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => RoleDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new RoleDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new RoleDto(item));
    const res = await RoleCreateAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<RoleDto>();
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
    name: "roleCreate",
    cliName: "create",
    cliShort: "role-c",
    url: "/role",
    method: "post",
    description: 'Creates a new "role" row.',
    in: {
      dto: "RoleDto",
    },
    out: {
      envelope: "GResponse",
      dto: "RoleDto",
    },
  };
}
