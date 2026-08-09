import { CapabilityDto } from "./CapabilityDto";
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
 * Action to communicate with the action capabilityCreate
 */
export type CapabilityCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type CapabilityCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  CapabilityCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => CapabilityDto;
  }>;
export const useCapabilityCreateAction = (
  options?: CapabilityCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: CapabilityDto) => {
    setCompleteState(false);
    return CapabilityCreateAction.Fetch(
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
 * CapabilityCreateAction
 */
export class CapabilityCreateAction {
  //
  static URL = "/capability";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(CapabilityCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<CapabilityDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<CapabilityDto>, CapabilityDto, unknown>(
      overrideUrl ?? CapabilityCreateAction.NewUrl(qs),
      {
        method: CapabilityCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<CapabilityDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => CapabilityDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new CapabilityDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new CapabilityDto(item));
    const res = await CapabilityCreateAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<CapabilityDto>();
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
    name: "capabilityCreate",
    cliName: "create",
    cliShort: "capability-c",
    url: "/capability",
    method: "post",
    description: 'Creates a new "capability" row.',
    in: {
      dto: "CapabilityDto",
    },
    out: {
      envelope: "GResponse",
      dto: "CapabilityDto",
    },
  };
}
