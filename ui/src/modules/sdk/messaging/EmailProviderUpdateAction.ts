import { EmailProviderDto } from "./EmailProviderDto";
import { EmailProviderOptionalDto } from "./EmailProviderOptionalDto";
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
 * Action to communicate with the action emailProviderUpdate
 */
export type EmailProviderUpdateActionOptions = {
  queryKey?: unknown[];
  params: EmailProviderUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type EmailProviderUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  EmailProviderUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => EmailProviderDto;
  }>;
export const useEmailProviderUpdateAction = (
  options: EmailProviderUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: EmailProviderOptionalDto) => {
    setCompleteState(false);
    return EmailProviderUpdateAction.Fetch(
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
 * Path parameters for EmailProviderUpdateAction
 */
export type EmailProviderUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * EmailProviderUpdateAction
 */
export class EmailProviderUpdateAction {
  //
  static URL = "/emailProvider/:uniqueId";
  static NewUrl = (
    params: EmailProviderUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(EmailProviderUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: EmailProviderUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<EmailProviderOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<EmailProviderDto>,
      EmailProviderOptionalDto,
      unknown
    >(
      overrideUrl ?? EmailProviderUpdateAction.NewUrl(params, qs),
      {
        method: EmailProviderUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: EmailProviderUpdateActionPathParameter,
    init?: TypedRequestInit<EmailProviderOptionalDto, unknown>,
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
    const res = await EmailProviderUpdateAction.Fetch$(
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
    name: "emailProviderUpdate",
    cliName: "update",
    cliShort: "emailProvider-u",
    url: "/emailProvider/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "emailProvider" row by uniqueId.',
    in: {
      dto: "EmailProviderOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "EmailProviderDto",
    },
  };
}
