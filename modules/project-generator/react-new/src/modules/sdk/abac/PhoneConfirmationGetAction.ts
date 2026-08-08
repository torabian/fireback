import { GResponse } from "../sdk/envelopes/index";
import { PhoneConfirmationDto } from "./PhoneConfirmationDto";
import { buildUrl } from "../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
  type TypedResponse,
} from "../sdk/common/fetchx";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "@tanstack/react-query";
import { useFetchxContext } from "../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action phoneConfirmationGet
 */
export type PhoneConfirmationGetActionOptions = {
  queryKey?: unknown[];
  params: PhoneConfirmationGetActionPathParameter;
  qs?: URLSearchParams;
};
export type PhoneConfirmationGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<PhoneConfirmationDto>, unknown[]>,
  "queryKey"
> &
  PhoneConfirmationGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => PhoneConfirmationDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const usePhoneConfirmationGetActionQuery = (
  options: PhoneConfirmationGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return PhoneConfirmationGetAction.Fetch(
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
    queryKey: [PhoneConfirmationGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type PhoneConfirmationGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PhoneConfirmationGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PhoneConfirmationDto;
  }>;
export const usePhoneConfirmationGetAction = (
  options: PhoneConfirmationGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return PhoneConfirmationGetAction.Fetch(
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
 * Path parameters for PhoneConfirmationGetAction
 */
export type PhoneConfirmationGetActionPathParameter = {
  uniqueId: string;
};
/**
 * PhoneConfirmationGetAction
 */
export class PhoneConfirmationGetAction {
  //
  static URL = "/phoneConfirmation/:uniqueId";
  static NewUrl = (
    params: PhoneConfirmationGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PhoneConfirmationGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: PhoneConfirmationGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<PhoneConfirmationDto>, unknown, unknown>(
      overrideUrl ?? PhoneConfirmationGetAction.NewUrl(params, qs),
      {
        method: PhoneConfirmationGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PhoneConfirmationGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => PhoneConfirmationDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new PhoneConfirmationDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new PhoneConfirmationDto(item));
    const res = await PhoneConfirmationGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<PhoneConfirmationDto>();
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
    name: "phoneConfirmationGet",
    cliShort: "phoneConfirmation-g",
    url: "/phoneConfirmation/:uniqueId string",
    method: "get",
    description: 'Looks up a single "phoneConfirmation" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "PhoneConfirmationDto",
    },
  };
}
