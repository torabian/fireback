import { AppMenuDto } from "./AppMenuDto";
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
 * Action to communicate with the action appMenuCreate
 */
export type AppMenuCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type AppMenuCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  AppMenuCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => AppMenuDto;
  }>;
export const useAppMenuCreateAction = (
  options?: AppMenuCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: AppMenuDto) => {
    setCompleteState(false);
    return AppMenuCreateAction.Fetch(
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
 * AppMenuCreateAction
 */
export class AppMenuCreateAction {
  //
  static URL = "/appMenu";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(AppMenuCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<AppMenuDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<AppMenuDto>, AppMenuDto, unknown>(
      overrideUrl ?? AppMenuCreateAction.NewUrl(qs),
      {
        method: AppMenuCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<AppMenuDto, unknown>,
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
    const res = await AppMenuCreateAction.Fetch$(qs, ctx, init, overrideUrl);
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
    name: "appMenuCreate",
    cliName: "create",
    cliShort: "appMenu-c",
    url: "/appMenu",
    method: "post",
    description: 'Creates a new "appMenu" row.',
    in: {
      dto: "AppMenuDto",
    },
    out: {
      envelope: "GResponse",
      dto: "AppMenuDto",
    },
  };
}
