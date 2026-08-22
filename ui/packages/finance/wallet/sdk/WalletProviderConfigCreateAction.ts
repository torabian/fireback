import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletProviderConfigDto } from "./WalletProviderConfigDto";
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
 * Action to communicate with the action walletProviderConfigCreate
 */
export type WalletProviderConfigCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type WalletProviderConfigCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletProviderConfigCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletProviderConfigDto;
  }>;
export const useWalletProviderConfigCreateAction = (
  options?: WalletProviderConfigCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WalletProviderConfigDto) => {
    setCompleteState(false);
    return WalletProviderConfigCreateAction.Fetch(
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
 * WalletProviderConfigCreateAction
 */
export class WalletProviderConfigCreateAction {
  //
  static URL = "/walletProviderConfig";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WalletProviderConfigCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WalletProviderConfigDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WalletProviderConfigDto>,
      WalletProviderConfigDto,
      unknown
    >(
      overrideUrl ?? WalletProviderConfigCreateAction.NewUrl(qs),
      {
        method: WalletProviderConfigCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<WalletProviderConfigDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletProviderConfigDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletProviderConfigDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletProviderConfigDto(item));
    const res = await WalletProviderConfigCreateAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletProviderConfigDto>();
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
    name: "walletProviderConfigCreate",
    cliName: "create",
    cliShort: "c",
    url: "/walletProviderConfig",
    method: "post",
    description: 'Creates a new "walletProviderConfig" row.',
    in: {
      dto: "WalletProviderConfigDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WalletProviderConfigDto",
    },
  };
}
