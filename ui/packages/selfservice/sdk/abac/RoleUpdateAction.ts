import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { RoleDto } from "./RoleDto";
import { RoleOptionalDto } from "./RoleOptionalDto";
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
 * Action to communicate with the action roleUpdate
 */
export type RoleUpdateActionOptions = {
  queryKey?: unknown[];
  params: RoleUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type RoleUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  RoleUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => RoleDto;
  }>;
export const useRoleUpdateAction = (
  options: RoleUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: RoleOptionalDto) => {
    setCompleteState(false);
    return RoleUpdateAction.Fetch(
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
 * Path parameters for RoleUpdateAction
 */
export type RoleUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * RoleUpdateAction
 */
export class RoleUpdateAction {
  //
  static URL = "/role/:uniqueId";
  static NewUrl = (
    params: RoleUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(RoleUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: RoleUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<RoleOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<RoleDto>, RoleOptionalDto, unknown>(
      overrideUrl ?? RoleUpdateAction.NewUrl(params, qs),
      {
        method: RoleUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: RoleUpdateActionPathParameter,
    init?: TypedRequestInit<RoleOptionalDto, unknown>,
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
    const res = await RoleUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
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
    name: "roleUpdate",
    cliName: "update",
    cliShort: "u",
    url: "/role/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "role" row by uniqueId.',
    in: {
      dto: "RoleOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "RoleDto",
    },
  };
}
