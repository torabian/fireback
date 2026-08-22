import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletCurrencyDto } from "./WalletCurrencyDto";
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
 * Action to communicate with the action walletCurrencyCreate
 */
export type WalletCurrencyCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type WalletCurrencyCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletCurrencyCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletCurrencyDto;
  }>;
export const useWalletCurrencyCreateAction = (
  options?: WalletCurrencyCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WalletCurrencyDto) => {
    setCompleteState(false);
    return WalletCurrencyCreateAction.Fetch(
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
 * WalletCurrencyCreateAction
 */
export class WalletCurrencyCreateAction {
  //
  static URL = "/walletCurrency";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WalletCurrencyCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WalletCurrencyDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletCurrencyDto>, WalletCurrencyDto, unknown>(
      overrideUrl ?? WalletCurrencyCreateAction.NewUrl(qs),
      {
        method: WalletCurrencyCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<WalletCurrencyDto, unknown>,
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
    const res = await WalletCurrencyCreateAction.Fetch$(
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
    name: "walletCurrencyCreate",
    cliName: "create",
    cliShort: "c",
    url: "/walletCurrency",
    method: "post",
    description: 'Creates a new "walletCurrency" row.',
    in: {
      dto: "WalletCurrencyDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WalletCurrencyDto",
    },
  };
}
