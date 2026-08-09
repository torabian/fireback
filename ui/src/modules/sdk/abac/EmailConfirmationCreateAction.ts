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
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action emailConfirmationCreate
 */
export type EmailConfirmationCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type EmailConfirmationCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  EmailConfirmationCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => EmailConfirmationDto;
  }>;
export const useEmailConfirmationCreateAction = (
  options?: EmailConfirmationCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: EmailConfirmationDto) => {
    setCompleteState(false);
    return EmailConfirmationCreateAction.Fetch(
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
 * EmailConfirmationCreateAction
 */
export class EmailConfirmationCreateAction {
  //
  static URL = "/emailConfirmation";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(EmailConfirmationCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<EmailConfirmationDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<EmailConfirmationDto>,
      EmailConfirmationDto,
      unknown
    >(
      overrideUrl ?? EmailConfirmationCreateAction.NewUrl(qs),
      {
        method: EmailConfirmationCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<EmailConfirmationDto, unknown>,
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
    const res = await EmailConfirmationCreateAction.Fetch$(
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
    name: "emailConfirmationCreate",
    cliName: "create",
    cliShort: "emailConfirmation-c",
    url: "/emailConfirmation",
    method: "post",
    description: 'Creates a new "emailConfirmation" row.',
    in: {
      dto: "EmailConfirmationDto",
    },
    out: {
      envelope: "GResponse",
      dto: "EmailConfirmationDto",
    },
  };
}
