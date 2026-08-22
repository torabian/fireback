import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { PublicJoinKeyDto } from "./PublicJoinKeyDto";
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
 * Action to communicate with the action publicJoinKeyCreate
 */
export type PublicJoinKeyCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type PublicJoinKeyCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PublicJoinKeyCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PublicJoinKeyDto;
  }>;
export const usePublicJoinKeyCreateAction = (
  options?: PublicJoinKeyCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PublicJoinKeyDto) => {
    setCompleteState(false);
    return PublicJoinKeyCreateAction.Fetch(
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
 * PublicJoinKeyCreateAction
 */
export class PublicJoinKeyCreateAction {
  //
  static URL = "/publicJoinKey";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(PublicJoinKeyCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PublicJoinKeyDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<PublicJoinKeyDto>, PublicJoinKeyDto, unknown>(
      overrideUrl ?? PublicJoinKeyCreateAction.NewUrl(qs),
      {
        method: PublicJoinKeyCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<PublicJoinKeyDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => PublicJoinKeyDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new PublicJoinKeyDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new PublicJoinKeyDto(item));
    const res = await PublicJoinKeyCreateAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<PublicJoinKeyDto>();
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
    name: "publicJoinKeyCreate",
    cliName: "create",
    cliShort: "c",
    url: "/publicJoinKey",
    method: "post",
    description: 'Creates a new "publicJoinKey" row.',
    in: {
      dto: "PublicJoinKeyDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PublicJoinKeyDto",
    },
  };
}
