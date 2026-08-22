import { EmailProviderDto } from "./EmailProviderDto";
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
 * Action to communicate with the action emailProviderCreate
 */
export type EmailProviderCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type EmailProviderCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  EmailProviderCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => EmailProviderDto;
  }>;
export const useEmailProviderCreateAction = (
  options?: EmailProviderCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: EmailProviderDto) => {
    setCompleteState(false);
    return EmailProviderCreateAction.Fetch(
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
 * EmailProviderCreateAction
 */
export class EmailProviderCreateAction {
  //
  static URL = "/emailProvider";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(EmailProviderCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<EmailProviderDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<EmailProviderDto>, EmailProviderDto, unknown>(
      overrideUrl ?? EmailProviderCreateAction.NewUrl(qs),
      {
        method: EmailProviderCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<EmailProviderDto, unknown>,
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
    const res = await EmailProviderCreateAction.Fetch$(
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
    name: "emailProviderCreate",
    cliName: "create",
    cliShort: "c",
    url: "/emailProvider",
    method: "post",
    description: 'Creates a new "emailProvider" row.',
    in: {
      dto: "EmailProviderDto",
    },
    out: {
      envelope: "GResponse",
      dto: "EmailProviderDto",
    },
  };
}
