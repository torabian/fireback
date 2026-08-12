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
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action emailSenderCreate
 */
export type EmailSenderCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type EmailSenderCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  EmailSenderCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => EmailSenderDto;
  }>;
export const useEmailSenderCreateAction = (
  options?: EmailSenderCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: EmailSenderDto) => {
    setCompleteState(false);
    return EmailSenderCreateAction.Fetch(
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
 * EmailSenderCreateAction
 */
export class EmailSenderCreateAction {
  //
  static URL = "/emailSender";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(EmailSenderCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<EmailSenderDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<EmailSenderDto>, EmailSenderDto, unknown>(
      overrideUrl ?? EmailSenderCreateAction.NewUrl(qs),
      {
        method: EmailSenderCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<EmailSenderDto, unknown>,
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
    const res = await EmailSenderCreateAction.Fetch$(
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
    name: "emailSenderCreate",
    cliName: "create",
    cliShort: "emailSender-c",
    url: "/emailSender",
    method: "post",
    description: 'Creates a new "emailSender" row.',
    in: {
      dto: "EmailSenderDto",
    },
    out: {
      envelope: "GResponse",
      dto: "EmailSenderDto",
    },
  };
}
