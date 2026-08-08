import { GResponse } from "../sdk/envelopes/index";
import { PhoneConfirmationDto } from "./PhoneConfirmationDto";
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
 * Action to communicate with the action phoneConfirmationCreate
 */
export type PhoneConfirmationCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type PhoneConfirmationCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PhoneConfirmationCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PhoneConfirmationDto;
  }>;
export const usePhoneConfirmationCreateAction = (
  options?: PhoneConfirmationCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PhoneConfirmationDto) => {
    setCompleteState(false);
    return PhoneConfirmationCreateAction.Fetch(
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
 * PhoneConfirmationCreateAction
 */
export class PhoneConfirmationCreateAction {
  //
  static URL = "/phoneConfirmation";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(PhoneConfirmationCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PhoneConfirmationDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<PhoneConfirmationDto>,
      PhoneConfirmationDto,
      unknown
    >(
      overrideUrl ?? PhoneConfirmationCreateAction.NewUrl(qs),
      {
        method: PhoneConfirmationCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<PhoneConfirmationDto, unknown>,
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
    const res = await PhoneConfirmationCreateAction.Fetch$(
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
    name: "phoneConfirmationCreate",
    cliShort: "phoneConfirmation-c",
    url: "/phoneConfirmation",
    method: "post",
    description: 'Creates a new "phoneConfirmation" row.',
    in: {
      dto: "PhoneConfirmationDto",
    },
    out: {
      envelope: "GResponse",
      dto: "PhoneConfirmationDto",
    },
  };
}
