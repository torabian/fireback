import { EmailConfirmationDto } from "./EmailConfirmationDto";
import { EmailConfirmationOptionalDto } from "./EmailConfirmationOptionalDto";
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
 * Action to communicate with the action emailConfirmationUpdate
 */
export type EmailConfirmationUpdateActionOptions = {
  queryKey?: unknown[];
  params: EmailConfirmationUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type EmailConfirmationUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  EmailConfirmationUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => EmailConfirmationDto;
  }>;
export const useEmailConfirmationUpdateAction = (
  options: EmailConfirmationUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: EmailConfirmationOptionalDto) => {
    setCompleteState(false);
    return EmailConfirmationUpdateAction.Fetch(
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
 * Path parameters for EmailConfirmationUpdateAction
 */
export type EmailConfirmationUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * EmailConfirmationUpdateAction
 */
export class EmailConfirmationUpdateAction {
  //
  static URL = "/emailConfirmation/:uniqueId";
  static NewUrl = (
    params: EmailConfirmationUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(EmailConfirmationUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: EmailConfirmationUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<EmailConfirmationOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<EmailConfirmationDto>,
      EmailConfirmationOptionalDto,
      unknown
    >(
      overrideUrl ?? EmailConfirmationUpdateAction.NewUrl(params, qs),
      {
        method: EmailConfirmationUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: EmailConfirmationUpdateActionPathParameter,
    init?: TypedRequestInit<EmailConfirmationOptionalDto, unknown>,
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
    const res = await EmailConfirmationUpdateAction.Fetch$(
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
    name: "emailConfirmationUpdate",
    cliShort: "emailConfirmation-u",
    url: "/emailConfirmation/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "emailConfirmation" row by uniqueId.',
    in: {
      dto: "EmailConfirmationOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "EmailConfirmationDto",
    },
  };
}
