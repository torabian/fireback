import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletGatewayDto } from "./WalletGatewayDto";
import { WalletGatewayOptionalDto } from "./WalletGatewayOptionalDto";
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
 * Action to communicate with the action walletGatewayUpdate
 */
export type WalletGatewayUpdateActionOptions = {
  queryKey?: unknown[];
  params: WalletGatewayUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type WalletGatewayUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletGatewayUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletGatewayDto;
  }>;
export const useWalletGatewayUpdateAction = (
  options: WalletGatewayUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WalletGatewayOptionalDto) => {
    setCompleteState(false);
    return WalletGatewayUpdateAction.Fetch(
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
 * Path parameters for WalletGatewayUpdateAction
 */
export type WalletGatewayUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * WalletGatewayUpdateAction
 */
export class WalletGatewayUpdateAction {
  //
  static URL = "/walletGateway/:uniqueId";
  static NewUrl = (
    params: WalletGatewayUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(WalletGatewayUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: WalletGatewayUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WalletGatewayOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WalletGatewayDto>,
      WalletGatewayOptionalDto,
      unknown
    >(
      overrideUrl ?? WalletGatewayUpdateAction.NewUrl(params, qs),
      {
        method: WalletGatewayUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: WalletGatewayUpdateActionPathParameter,
    init?: TypedRequestInit<WalletGatewayOptionalDto, unknown>,
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
    const res = await WalletGatewayUpdateAction.Fetch$(
      params,
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
    name: "walletGatewayUpdate",
    cliName: "update",
    cliShort: "u",
    url: "/walletGateway/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "walletGateway" row by uniqueId.',
    in: {
      dto: "WalletGatewayOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "WalletGatewayDto",
    },
  };
}
