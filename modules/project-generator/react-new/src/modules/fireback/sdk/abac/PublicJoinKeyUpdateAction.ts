import { GResponse } from "../sdk/envelopes/index";
import { PublicJoinKeyDto } from "./PublicJoinKeyDto";
import { PublicJoinKeyOptionalDto } from "./PublicJoinKeyOptionalDto";
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
 * Action to communicate with the action publicJoinKeyUpdate
 */
export type PublicJoinKeyUpdateActionOptions = {
  queryKey?: unknown[];
  params: PublicJoinKeyUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type PublicJoinKeyUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PublicJoinKeyUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PublicJoinKeyDto;
  }>;
export const usePublicJoinKeyUpdateAction = (
  options: PublicJoinKeyUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PublicJoinKeyOptionalDto) => {
    setCompleteState(false);
    return PublicJoinKeyUpdateAction.Fetch(
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
 * Path parameters for PublicJoinKeyUpdateAction
 */
export type PublicJoinKeyUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * PublicJoinKeyUpdateAction
 */
export class PublicJoinKeyUpdateAction {
  //
  static URL = "/publicJoinKey/:uniqueId";
  static NewUrl = (
    params: PublicJoinKeyUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PublicJoinKeyUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: PublicJoinKeyUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PublicJoinKeyOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<PublicJoinKeyDto>,
      PublicJoinKeyOptionalDto,
      unknown
    >(
      overrideUrl ?? PublicJoinKeyUpdateAction.NewUrl(params, qs),
      {
        method: PublicJoinKeyUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PublicJoinKeyUpdateActionPathParameter,
    init?: TypedRequestInit<PublicJoinKeyOptionalDto, unknown>,
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
    const res = await PublicJoinKeyUpdateAction.Fetch$(
      params,
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
    name: "publicJoinKeyUpdate",
    cliShort: "publicJoinKey-u",
    url: "/publicJoinKey/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "publicJoinKey" row by uniqueId.',
    in: {
      dto: "PublicJoinKeyOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PublicJoinKeyDto",
    },
  };
}
