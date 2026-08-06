import { EmailProviderDto } from "./EmailProviderDto";
import { GResponse } from "../sdk/envelopes/index";
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
 * Action to communicate with the action emailProviderGet
 */
export type EmailProviderGetActionOptions = {
  queryKey?: unknown[];
  params: EmailProviderGetActionPathParameter;
  qs?: URLSearchParams;
};
export type EmailProviderGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<EmailProviderDto>, unknown[]>,
  "queryKey"
> &
  EmailProviderGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => EmailProviderDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useEmailProviderGetActionQuery = (
  options: EmailProviderGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return EmailProviderGetAction.Fetch(
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
    queryKey: [EmailProviderGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type EmailProviderGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  EmailProviderGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => EmailProviderDto;
  }>;
export const useEmailProviderGetAction = (
  options: EmailProviderGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return EmailProviderGetAction.Fetch(
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
 * Path parameters for EmailProviderGetAction
 */
export type EmailProviderGetActionPathParameter = {
  uniqueId: string;
};
/**
 * EmailProviderGetAction
 */
export class EmailProviderGetAction {
  //
  static URL = "/emailProvider/:uniqueId";
  static NewUrl = (
    params: EmailProviderGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(EmailProviderGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: EmailProviderGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<EmailProviderDto>, unknown, unknown>(
      overrideUrl ?? EmailProviderGetAction.NewUrl(params, qs),
      {
        method: EmailProviderGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: EmailProviderGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => EmailProviderDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new EmailProviderDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new EmailProviderDto(item));
    const res = await EmailProviderGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<EmailProviderDto>();
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
    name: "emailProviderGet",
    cliShort: "emailProvider-g",
    url: "/emailProvider/:uniqueId string",
    method: "get",
    description: 'Looks up a single "emailProvider" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "EmailProviderDto",
    },
  };
}
