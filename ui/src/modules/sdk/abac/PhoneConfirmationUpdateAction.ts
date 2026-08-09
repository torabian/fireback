import { GResponse } from "../sdk/envelopes/index";
import { PhoneConfirmationDto } from "./PhoneConfirmationDto";
import { PhoneConfirmationOptionalDto } from "./PhoneConfirmationOptionalDto";
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
 * Action to communicate with the action phoneConfirmationUpdate
 */
export type PhoneConfirmationUpdateActionOptions = {
  queryKey?: unknown[];
  params: PhoneConfirmationUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type PhoneConfirmationUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PhoneConfirmationUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PhoneConfirmationDto;
  }>;
export const usePhoneConfirmationUpdateAction = (
  options: PhoneConfirmationUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PhoneConfirmationOptionalDto) => {
    setCompleteState(false);
    return PhoneConfirmationUpdateAction.Fetch(
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
 * Path parameters for PhoneConfirmationUpdateAction
 */
export type PhoneConfirmationUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * PhoneConfirmationUpdateAction
 */
export class PhoneConfirmationUpdateAction {
  //
  static URL = "/phoneConfirmation/:uniqueId";
  static NewUrl = (
    params: PhoneConfirmationUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(PhoneConfirmationUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: PhoneConfirmationUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PhoneConfirmationOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<PhoneConfirmationDto>,
      PhoneConfirmationOptionalDto,
      unknown
    >(
      overrideUrl ?? PhoneConfirmationUpdateAction.NewUrl(params, qs),
      {
        method: PhoneConfirmationUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: PhoneConfirmationUpdateActionPathParameter,
    init?: TypedRequestInit<PhoneConfirmationOptionalDto, unknown>,
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
    const res = await PhoneConfirmationUpdateAction.Fetch$(
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
    name: "phoneConfirmationUpdate",
    cliName: "update",
    cliShort: "phoneConfirmation-u",
    url: "/phoneConfirmation/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "phoneConfirmation" row by uniqueId.',
    in: {
      dto: "PhoneConfirmationOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PhoneConfirmationDto",
    },
  };
}
