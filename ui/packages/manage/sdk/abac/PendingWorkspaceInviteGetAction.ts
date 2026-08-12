import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { PendingWorkspaceInviteDto } from "./PendingWorkspaceInviteDto";
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
 * Action to communicate with the action pendingWorkspaceInviteGet
 */
export type PendingWorkspaceInviteGetActionOptions = {
  queryKey?: unknown[];
  params: PendingWorkspaceInviteGetActionPathParameter;
  qs?: URLSearchParams;
};
export type PendingWorkspaceInviteGetActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<PendingWorkspaceInviteDto>,
    unknown[]
  >,
  "queryKey"
> &
  PendingWorkspaceInviteGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => PendingWorkspaceInviteDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const usePendingWorkspaceInviteGetActionQuery = (
  options: PendingWorkspaceInviteGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return PendingWorkspaceInviteGetAction.Fetch(
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
    queryKey: [
      PendingWorkspaceInviteGetAction.NewUrl(options.params, options?.qs),
    ],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type PendingWorkspaceInviteGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PendingWorkspaceInviteGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PendingWorkspaceInviteDto;
  }>;
export const usePendingWorkspaceInviteGetAction = (
  options: PendingWorkspaceInviteGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return PendingWorkspaceInviteGetAction.Fetch(
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
 * Path parameters for PendingWorkspaceInviteGetAction
 */
export type PendingWorkspaceInviteGetActionPathParameter = {
  uniqueId: string;
};
/**
 * PendingWorkspaceInviteGetAction
 */
export class PendingWorkspaceInviteGetAction {
  //
  static URL = "/pendingWorkspaceInvite/:uniqueId";
  static NewUrl = (
    params: PendingWorkspaceInviteGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PendingWorkspaceInviteGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: PendingWorkspaceInviteGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<PendingWorkspaceInviteDto>, unknown, unknown>(
      overrideUrl ?? PendingWorkspaceInviteGetAction.NewUrl(params, qs),
      {
        method: PendingWorkspaceInviteGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PendingWorkspaceInviteGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => PendingWorkspaceInviteDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new PendingWorkspaceInviteDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new PendingWorkspaceInviteDto(item));
    const res = await PendingWorkspaceInviteGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<PendingWorkspaceInviteDto>();
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
    name: "pendingWorkspaceInviteGet",
    cliName: "get",
    cliShort: "pendingWorkspaceInvite-g",
    url: "/pendingWorkspaceInvite/:uniqueId string",
    method: "get",
    description: 'Looks up a single "pendingWorkspaceInvite" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "PendingWorkspaceInviteDto",
    },
  };
}
