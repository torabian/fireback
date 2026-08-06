import { AppMenuOptionalDto } from "./AppMenuOptionalDto";
import { GResponse } from "../sdk/envelopes/index";
import { URLSearchParamsX } from "../sdk/common/URLSearchParamsX";
import { buildUrl } from "../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type TypedRequestInit,
  type TypedResponse,
} from "../sdk/common/fetchx";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "@tanstack/react-query";
import { useFetchxContext } from "../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action CteAppMenus
 */
export type CteAppMenusActionOptions = {
  queryKey?: unknown[];
  qs?: CteAppMenusActionQueryParams;
};
export type CteAppMenusActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<AppMenuOptionalDto>, unknown[]>,
  "queryKey"
> &
  CteAppMenusActionOptions &
  Partial<{
    creatorFn: (item: unknown) => AppMenuOptionalDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useCteAppMenusActionQuery = (
  options: CteAppMenusActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return CteAppMenusAction.Fetch(
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
    queryKey: [CteAppMenusAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type CteAppMenusActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  CteAppMenusActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => AppMenuOptionalDto;
  }>;
export const useCteAppMenusAction = (
  options?: CteAppMenusActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return CteAppMenusAction.Fetch(
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
 * CteAppMenusAction
 */
export class CteAppMenusAction {
  //
  static URL = "/cte-app-menus";
  static NewUrl = (qs?: CteAppMenusActionQueryParams) =>
    buildUrl(CteAppMenusAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: CteAppMenusActionQueryParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<AppMenuOptionalDto>, unknown, unknown>(
      overrideUrl ?? CteAppMenusAction.NewUrl(qs),
      {
        method: CteAppMenusAction.Method,
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
      creatorFn?: ((item: unknown) => AppMenuOptionalDto) | undefined;
      qs?: CteAppMenusActionQueryParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new AppMenuOptionalDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new AppMenuOptionalDto(item));
    const res = await CteAppMenusAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<AppMenuOptionalDto>();
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
    name: "CteAppMenus",
    cliName: "cte-app-menus",
    url: "/cte-app-menus",
    method: "get",
    qs: [
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
      'Returns "appMenu" rows assembled into a tree (each item\'s children nested under it via a recursive query) instead of Browse\'s flat list - restores the old Module3-generated /cte-app-menus route. See AppMenuActions.CteQuery / AppMenuTreeNode. Note: Emi has no self-referencing dto, so the declared response type here is the flat AppMenuOptionalDto - the actual JSON payload additionally nests a "children" array of the same shape under each item.',
    out: {
      envelope: "GResponse",
      dto: "AppMenuOptionalDto",
    },
  };
}
/**
 * CteAppMenusActionQueryParams class
 * Auto-generated from EmiAction
 */
export class CteAppMenusActionQueryParams extends URLSearchParamsX {
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
