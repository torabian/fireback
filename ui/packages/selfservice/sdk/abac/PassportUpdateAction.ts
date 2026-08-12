import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { PassportDto } from "./PassportDto";
import { PassportOptionalDto } from "./PassportOptionalDto";
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
 * Action to communicate with the action passportUpdate
 */
export type PassportUpdateActionOptions = {
  queryKey?: unknown[];
  params: PassportUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type PassportUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PassportUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PassportDto;
  }>;
export const usePassportUpdateAction = (
  options: PassportUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PassportOptionalDto) => {
    setCompleteState(false);
    return PassportUpdateAction.Fetch(
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
 * Path parameters for PassportUpdateAction
 */
export type PassportUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * PassportUpdateAction
 */
export class PassportUpdateAction {
  //
  static URL = "/passport/:uniqueId";
  static NewUrl = (
    params: PassportUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PassportUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: PassportUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PassportOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<PassportDto>, PassportOptionalDto, unknown>(
      overrideUrl ?? PassportUpdateAction.NewUrl(params, qs),
      {
        method: PassportUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PassportUpdateActionPathParameter,
    init?: TypedRequestInit<PassportOptionalDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => PassportDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new PassportDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new PassportDto(item));
    const res = await PassportUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<PassportDto>();
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
    name: "passportUpdate",
    cliName: "update",
    cliShort: "passport-u",
    url: "/passport/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "passport" row by uniqueId.',
    in: {
      dto: "PassportOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PassportDto",
    },
  };
}
