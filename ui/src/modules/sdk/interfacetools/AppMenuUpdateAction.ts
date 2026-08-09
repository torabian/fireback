import { AppMenuDto } from "./AppMenuDto";
import { AppMenuOptionalDto } from "./AppMenuOptionalDto";
import { GResponse } from "../sdk/envelopes/index";
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
 * Action to communicate with the action appMenuUpdate
 */
export type AppMenuUpdateActionOptions = {
  queryKey?: unknown[];
  params: AppMenuUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type AppMenuUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  AppMenuUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => AppMenuDto;
  }>;
export const useAppMenuUpdateAction = (
  options: AppMenuUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: AppMenuOptionalDto) => {
    setCompleteState(false);
    return AppMenuUpdateAction.Fetch(
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
 * Path parameters for AppMenuUpdateAction
 */
export type AppMenuUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * AppMenuUpdateAction
 */
export class AppMenuUpdateAction {
  //
  static URL = "/appMenu/:uniqueId";
  static NewUrl = (
    params: AppMenuUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(AppMenuUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: AppMenuUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<AppMenuOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<AppMenuDto>, AppMenuOptionalDto, unknown>(
      overrideUrl ?? AppMenuUpdateAction.NewUrl(params, qs),
      {
        method: AppMenuUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: AppMenuUpdateActionPathParameter,
    init?: TypedRequestInit<AppMenuOptionalDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => AppMenuDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new AppMenuDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new AppMenuDto(item));
    const res = await AppMenuUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<AppMenuDto>();
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
    name: "appMenuUpdate",
    cliName: "update",
    cliShort: "appMenu-u",
    url: "/appMenu/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "appMenu" row by uniqueId.',
    in: {
      dto: "AppMenuOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "AppMenuDto",
    },
  };
}
