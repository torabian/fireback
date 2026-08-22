import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { URLSearchParamsX } from "@fireback/js-remote-ctx/common/URLSearchParamsX";
import { WalletTransactionViewDto } from "./WalletTransactionViewDto";
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
 * Action to communicate with the action walletHistory
 */
export type WalletHistoryActionOptions = {
  queryKey?: unknown[];
  qs?: WalletHistoryActionQueryParams;
};
export type WalletHistoryActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<WalletTransactionViewDto>,
    unknown[]
  >,
  "queryKey"
> &
  WalletHistoryActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WalletTransactionViewDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWalletHistoryActionQuery = (
  options: WalletHistoryActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WalletHistoryAction.Fetch(
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
    queryKey: [WalletHistoryAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WalletHistoryActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WalletHistoryActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletTransactionViewDto;
  }>;
export const useWalletHistoryAction = (
  options?: WalletHistoryActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WalletHistoryAction.Fetch(
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
 * WalletHistoryAction
 */
export class WalletHistoryAction {
  //
  static URL = "/wallet/transactions";
  static NewUrl = (qs?: WalletHistoryActionQueryParams) =>
    buildUrl(WalletHistoryAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: WalletHistoryActionQueryParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletTransactionViewDto>, unknown, unknown>(
      overrideUrl ?? WalletHistoryAction.NewUrl(qs),
      {
        method: WalletHistoryAction.Method,
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
      creatorFn?: ((item: unknown) => WalletTransactionViewDto) | undefined;
      qs?: WalletHistoryActionQueryParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletTransactionViewDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletTransactionViewDto(item));
    const res = await WalletHistoryAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletTransactionViewDto>();
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
    name: "walletHistory",
    cliShort: "history",
    url: "/wallet/transactions",
    method: "get",
    qs: [
      {
        name: "walletId",
        type: "string",
      },
      {
        name: "filter",
        type: "string",
      },
      {
        name: "sort",
        type: "string",
      },
      {
        name: "startIndex",
        type: "int",
      },
      {
        name: "itemsPerPage",
        type: "int",
      },
      {
        name: "cursor",
        type: "string",
      },
    ],
    description:
      "Paginated balance-history ledger for one wallet, scoped to the caller the same way getWallet is.",
    out: {
      envelope: "GResponse",
      dto: "WalletTransactionViewDto",
    },
  };
}
/**
 * WalletHistoryActionQueryParams class
 * Auto-generated from EmiAction
 */
export class WalletHistoryActionQueryParams extends URLSearchParamsX {
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
  /**
   *
   * @returns { string | null }
   */
  getFilter() {
    return this.getTyped("filter", "string | null");
  }
  /**
   *
   * @param { string | null } value
   */
  setFilter(value: string | null) {
    this.set("filter", value);
    return this;
  }
  /**
   *
   * @returns { string | null }
   */
  getSort() {
    return this.getTyped("sort", "string | null");
  }
  /**
   *
   * @param { string | null } value
   */
  setSort(value: string | null) {
    this.set("sort", value);
    return this;
  }
  /**
   *
   * @returns { number | null }
   */
  getStartIndex() {
    return this.getTyped("startIndex", "number | null");
  }
  /**
   *
   * @param { number | null } value
   */
  setStartIndex(value: number | null) {
    this.set("startIndex", value);
    return this;
  }
  /**
   *
   * @returns { number | null }
   */
  getItemsPerPage() {
    return this.getTyped("itemsPerPage", "number | null");
  }
  /**
   *
   * @param { number | null } value
   */
  setItemsPerPage(value: number | null) {
    this.set("itemsPerPage", value);
    return this;
  }
  /**
   *
   * @returns { string | null }
   */
  getCursor() {
    return this.getTyped("cursor", "string | null");
  }
  /**
   *
   * @param { string | null } value
   */
  setCursor(value: string | null) {
    this.set("cursor", value);
    return this;
  }
}
