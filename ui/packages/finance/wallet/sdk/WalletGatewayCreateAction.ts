import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletGatewayDto } from "./WalletGatewayDto";
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
 * Action to communicate with the action walletGatewayCreate
 */
export type WalletGatewayCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type WalletGatewayCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletGatewayCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletGatewayDto;
  }>;
export const useWalletGatewayCreateAction = (
  options?: WalletGatewayCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WalletGatewayDto) => {
    setCompleteState(false);
    return WalletGatewayCreateAction.Fetch(
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
 * WalletGatewayCreateAction
 */
export class WalletGatewayCreateAction {
  //
  static URL = "/walletGateway";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WalletGatewayCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WalletGatewayDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletGatewayDto>, WalletGatewayDto, unknown>(
      overrideUrl ?? WalletGatewayCreateAction.NewUrl(qs),
      {
        method: WalletGatewayCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<WalletGatewayDto, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletGatewayDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletGatewayDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletGatewayDto(item));
    const res = await WalletGatewayCreateAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletGatewayDto>();
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
    name: "walletGatewayCreate",
    cliName: "create",
    cliShort: "c",
    url: "/walletGateway",
    method: "post",
    description: 'Creates a new "walletGateway" row.',
    in: {
      dto: "WalletGatewayDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WalletGatewayDto",
    },
  };
}
