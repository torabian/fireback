import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { PassportMethodDto } from "./PassportMethodDto";
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
 * Action to communicate with the action passportMethodCreate
 */
export type PassportMethodCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type PassportMethodCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PassportMethodCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PassportMethodDto;
  }>;
export const usePassportMethodCreateAction = (
  options?: PassportMethodCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PassportMethodDto) => {
    setCompleteState(false);
    return PassportMethodCreateAction.Fetch(
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
 * PassportMethodCreateAction
 */
export class PassportMethodCreateAction {
  //
  static URL = "/passportMethod";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(PassportMethodCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PassportMethodDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<PassportMethodDto>, PassportMethodDto, unknown>(
      overrideUrl ?? PassportMethodCreateAction.NewUrl(qs),
      {
        method: PassportMethodCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<PassportMethodDto, unknown>,
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
    const res = await PassportMethodCreateAction.Fetch$(
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
    name: "passportMethodCreate",
    cliName: "create",
    cliShort: "c",
    url: "/passportMethod",
    method: "post",
    description: 'Creates a new "passportMethod" row.',
    in: {
      dto: "PassportMethodDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PassportMethodDto",
    },
  };
}
