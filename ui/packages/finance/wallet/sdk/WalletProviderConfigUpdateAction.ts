import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletProviderConfigDto } from "./WalletProviderConfigDto";
import { WalletProviderConfigOptionalDto } from "./WalletProviderConfigOptionalDto";
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
 * Action to communicate with the action walletProviderConfigUpdate
 */
export type WalletProviderConfigUpdateActionOptions = {
  queryKey?: unknown[];
  params: WalletProviderConfigUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type WalletProviderConfigUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletProviderConfigUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletProviderConfigDto;
  }>;
export const useWalletProviderConfigUpdateAction = (
  options: WalletProviderConfigUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WalletProviderConfigOptionalDto) => {
    setCompleteState(false);
    return WalletProviderConfigUpdateAction.Fetch(
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
 * Path parameters for WalletProviderConfigUpdateAction
 */
export type WalletProviderConfigUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * WalletProviderConfigUpdateAction
 */
export class WalletProviderConfigUpdateAction {
  //
  static URL = "/walletProviderConfig/:uniqueId";
  static NewUrl = (
    params: WalletProviderConfigUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WalletProviderConfigUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: WalletProviderConfigUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WalletProviderConfigOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WalletProviderConfigDto>,
      WalletProviderConfigOptionalDto,
      unknown
    >(
      overrideUrl ?? WalletProviderConfigUpdateAction.NewUrl(params, qs),
      {
        method: WalletProviderConfigUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WalletProviderConfigUpdateActionPathParameter,
    init?: TypedRequestInit<WalletProviderConfigOptionalDto, unknown>,
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
    const res = await WalletProviderConfigUpdateAction.Fetch$(
      params,
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
    name: "walletProviderConfigUpdate",
    cliName: "update",
    cliShort: "u",
    url: "/walletProviderConfig/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "walletProviderConfig" row by uniqueId.',
    in: {
      dto: "WalletProviderConfigOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WalletProviderConfigDto",
    },
  };
}
