import { EmailSenderDto } from "./EmailSenderDto";
import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
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
 * Action to communicate with the action emailSenderGet
 */
export type EmailSenderGetActionOptions = {
  queryKey?: unknown[];
  params: EmailSenderGetActionPathParameter;
  qs?: URLSearchParams;
};
export type EmailSenderGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<EmailSenderDto>, unknown[]>,
  "queryKey"
> &
  EmailSenderGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => EmailSenderDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useEmailSenderGetActionQuery = (
  options: EmailSenderGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return EmailSenderGetAction.Fetch(
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
    queryKey: [EmailSenderGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type EmailSenderGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  EmailSenderGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => EmailSenderDto;
  }>;
export const useEmailSenderGetAction = (
  options: EmailSenderGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return EmailSenderGetAction.Fetch(
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
 * Path parameters for EmailSenderGetAction
 */
export type EmailSenderGetActionPathParameter = {
  uniqueId: string;
};
/**
 * EmailSenderGetAction
 */
export class EmailSenderGetAction {
  //
  static URL = "/emailSender/:uniqueId";
  static NewUrl = (
    params: EmailSenderGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(EmailSenderGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: EmailSenderGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<EmailSenderDto>, unknown, unknown>(
      overrideUrl ?? EmailSenderGetAction.NewUrl(params, qs),
      {
        method: EmailSenderGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: EmailSenderGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => EmailSenderDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new EmailSenderDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new EmailSenderDto(item));
    const res = await EmailSenderGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<EmailSenderDto>();
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
    name: "emailSenderGet",
    cliName: "get",
    cliShort: "emailSender-g",
    url: "/emailSender/:uniqueId string",
    method: "get",
    description: 'Looks up a single "emailSender" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "EmailSenderDto",
    },
  };
}
