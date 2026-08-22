import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletCurrencyDto } from "./WalletCurrencyDto";
import { WalletCurrencyOptionalDto } from "./WalletCurrencyOptionalDto";
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
 * Action to communicate with the action walletCurrencyUpdate
 */
export type WalletCurrencyUpdateActionOptions = {
  queryKey?: unknown[];
  params: WalletCurrencyUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type WalletCurrencyUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletCurrencyUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletCurrencyDto;
  }>;
export const useWalletCurrencyUpdateAction = (
  options: WalletCurrencyUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WalletCurrencyOptionalDto) => {
    setCompleteState(false);
    return WalletCurrencyUpdateAction.Fetch(
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
 * Path parameters for WalletCurrencyUpdateAction
 */
export type WalletCurrencyUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * WalletCurrencyUpdateAction
 */
export class WalletCurrencyUpdateAction {
  //
  static URL = "/walletCurrency/:uniqueId";
  static NewUrl = (
    params: WalletCurrencyUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WalletCurrencyUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: WalletCurrencyUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WalletCurrencyOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WalletCurrencyDto>,
      WalletCurrencyOptionalDto,
      unknown
    >(
      overrideUrl ?? WalletCurrencyUpdateAction.NewUrl(params, qs),
      {
        method: WalletCurrencyUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WalletCurrencyUpdateActionPathParameter,
    init?: TypedRequestInit<WalletCurrencyOptionalDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletCurrencyDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletCurrencyDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletCurrencyDto(item));
    const res = await WalletCurrencyUpdateAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletCurrencyDto>();
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
    name: "walletCurrencyUpdate",
    cliName: "update",
    cliShort: "u",
    url: "/walletCurrency/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "walletCurrency" row by uniqueId.',
    in: {
      dto: "WalletCurrencyOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WalletCurrencyDto",
    },
  };
}
