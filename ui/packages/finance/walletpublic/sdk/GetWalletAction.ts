import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { URLSearchParamsX } from "@fireback/js-remote-ctx/common/URLSearchParamsX";
import { WalletViewDto } from "./WalletViewDto";
import { buildUrl } from "@fireback/js-remote-ctx/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
  type TypedResponse,
} from "@fireback/js-remote-ctx/common/fetchx";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action getWallet
 */
export type GetWalletActionOptions = {
  queryKey?: unknown[];
  qs?: GetWalletActionQueryParams;
};
export type GetWalletActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<WalletViewDto>, unknown[]>,
  "queryKey"
> &
  GetWalletActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WalletViewDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useGetWalletActionQuery = (
  options: GetWalletActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return GetWalletAction.Fetch(
      {
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
  const result = useQuery({
    queryKey: [GetWalletAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type GetWalletActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  GetWalletActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletViewDto;
  }>;
export const useGetWalletAction = (
  options?: GetWalletActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return GetWalletAction.Fetch(
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
 * GetWalletAction
 */
export class GetWalletAction {
  //
  static URL = "/wallet/get";
  static NewUrl = (qs?: GetWalletActionQueryParams) =>
    buildUrl(GetWalletAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: GetWalletActionQueryParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletViewDto>, unknown, unknown>(
      overrideUrl ?? GetWalletAction.NewUrl(qs),
      {
        method: GetWalletAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletViewDto) | undefined;
      qs?: GetWalletActionQueryParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletViewDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletViewDto(item));
    const res = await GetWalletAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletViewDto>();
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
    name: "getWallet",
    cliShort: "get",
    url: "/wallet/get",
    method: "get",
    qs: [
      {
        name: "walletId",
        type: "string",
      },
    ],
    description:
      "Reads one wallet by id, scoped to the caller - the caller must be the owning user or a member of the owning workspace.",
    out: {
      envelope: "GResponse",
      dto: "WalletViewDto",
    },
  };
}
/**
 * GetWalletActionQueryParams class
 * Auto-generated from EmiAction
 */
export class GetWalletActionQueryParams extends URLSearchParamsX {
  /**
   *
   * @returns { string | null }
   */
  getWalletId() {
    return this.getTyped("walletId", "string | null");
  }
  /**
   *
   * @param { string | null } value
   */
  setWalletId(value: string | null) {
    this.set("walletId", value);
    return this;
  }
}
