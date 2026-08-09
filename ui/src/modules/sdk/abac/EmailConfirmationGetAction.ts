import { EmailConfirmationDto } from "./EmailConfirmationDto";
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
 * Action to communicate with the action emailConfirmationGet
 */
export type EmailConfirmationGetActionOptions = {
  queryKey?: unknown[];
  params: EmailConfirmationGetActionPathParameter;
  qs?: URLSearchParams;
};
export type EmailConfirmationGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<EmailConfirmationDto>, unknown[]>,
  "queryKey"
> &
  EmailConfirmationGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => EmailConfirmationDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useEmailConfirmationGetActionQuery = (
  options: EmailConfirmationGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return EmailConfirmationGetAction.Fetch(
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
    queryKey: [EmailConfirmationGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type EmailConfirmationGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  EmailConfirmationGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => EmailConfirmationDto;
  }>;
export const useEmailConfirmationGetAction = (
  options: EmailConfirmationGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return EmailConfirmationGetAction.Fetch(
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
 * Path parameters for EmailConfirmationGetAction
 */
export type EmailConfirmationGetActionPathParameter = {
  uniqueId: string;
};
/**
 * EmailConfirmationGetAction
 */
export class EmailConfirmationGetAction {
  //
  static URL = "/emailConfirmation/:uniqueId";
  static NewUrl = (
    params: EmailConfirmationGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(EmailConfirmationGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: EmailConfirmationGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<EmailConfirmationDto>, unknown, unknown>(
      overrideUrl ?? EmailConfirmationGetAction.NewUrl(params, qs),
      {
        method: EmailConfirmationGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: EmailConfirmationGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => EmailConfirmationDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new EmailConfirmationDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new EmailConfirmationDto(item));
    const res = await EmailConfirmationGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<EmailConfirmationDto>();
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
    name: "emailConfirmationGet",
    cliName: "get",
    cliShort: "emailConfirmation-g",
    url: "/emailConfirmation/:uniqueId string",
    method: "get",
    description: 'Looks up a single "emailConfirmation" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "EmailConfirmationDto",
    },
  };
}
