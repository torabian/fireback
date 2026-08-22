import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { PassportMethodDto } from "./PassportMethodDto";
import { PassportMethodOptionalDto } from "./PassportMethodOptionalDto";
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
 * Action to communicate with the action passportMethodUpdate
 */
export type PassportMethodUpdateActionOptions = {
  queryKey?: unknown[];
  params: PassportMethodUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type PassportMethodUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PassportMethodUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PassportMethodDto;
  }>;
export const usePassportMethodUpdateAction = (
  options: PassportMethodUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PassportMethodOptionalDto) => {
    setCompleteState(false);
    return PassportMethodUpdateAction.Fetch(
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
 * Path parameters for PassportMethodUpdateAction
 */
export type PassportMethodUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * PassportMethodUpdateAction
 */
export class PassportMethodUpdateAction {
  //
  static URL = "/passportMethod/:uniqueId";
  static NewUrl = (
    params: PassportMethodUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PassportMethodUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: PassportMethodUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PassportMethodOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<PassportMethodDto>,
      PassportMethodOptionalDto,
      unknown
    >(
      overrideUrl ?? PassportMethodUpdateAction.NewUrl(params, qs),
      {
        method: PassportMethodUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PassportMethodUpdateActionPathParameter,
    init?: TypedRequestInit<PassportMethodOptionalDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => PassportMethodDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new PassportMethodDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new PassportMethodDto(item));
    const res = await PassportMethodUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<PassportMethodDto>();
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
    name: "passportMethodUpdate",
    cliName: "update",
    cliShort: "u",
    url: "/passportMethod/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "passportMethod" row by uniqueId.',
    in: {
      dto: "PassportMethodOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PassportMethodDto",
    },
  };
}
