import { EmailSenderDto } from "./EmailSenderDto";
import { EmailSenderOptionalDto } from "./EmailSenderOptionalDto";
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
 * Action to communicate with the action emailSenderUpdate
 */
export type EmailSenderUpdateActionOptions = {
  queryKey?: unknown[];
  params: EmailSenderUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type EmailSenderUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  EmailSenderUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => EmailSenderDto;
  }>;
export const useEmailSenderUpdateAction = (
  options: EmailSenderUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: EmailSenderOptionalDto) => {
    setCompleteState(false);
    return EmailSenderUpdateAction.Fetch(
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
 * Path parameters for EmailSenderUpdateAction
 */
export type EmailSenderUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * EmailSenderUpdateAction
 */
export class EmailSenderUpdateAction {
  //
  static URL = "/emailSender/:uniqueId";
  static NewUrl = (
    params: EmailSenderUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(EmailSenderUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: EmailSenderUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<EmailSenderOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<EmailSenderDto>, EmailSenderOptionalDto, unknown>(
      overrideUrl ?? EmailSenderUpdateAction.NewUrl(params, qs),
      {
        method: EmailSenderUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: EmailSenderUpdateActionPathParameter,
    init?: TypedRequestInit<EmailSenderOptionalDto, unknown>,
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
    const res = await EmailSenderUpdateAction.Fetch$(
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
    name: "emailSenderUpdate",
    cliShort: "emailSender-u",
    url: "/emailSender/:uniqueId string",
    method: "patch",
    description: 'Applies a partial update to a "emailSender" row by uniqueId.',
    in: {
      dto: "EmailSenderOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "EmailSenderDto",
    },
  };
}
