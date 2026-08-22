import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { URLSearchParamsX } from "@fireback/js-remote-ctx/common/URLSearchParamsX";
import { WebPushConfigOptionalDto } from "./WebPushConfigOptionalDto";
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
 * Action to communicate with the action webPushConfigBrowse
 */
export type WebPushConfigBrowseActionOptions = {
  queryKey?: unknown[];
  qs?: WebPushConfigBrowseActionQueryParams;
};
export type WebPushConfigBrowseActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<WebPushConfigOptionalDto>,
    unknown[]
  >,
  "queryKey"
> &
  WebPushConfigBrowseActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WebPushConfigOptionalDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWebPushConfigBrowseActionQuery = (
  options: WebPushConfigBrowseActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WebPushConfigBrowseAction.Fetch(
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
    queryKey: [WebPushConfigBrowseAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WebPushConfigBrowseActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WebPushConfigBrowseActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WebPushConfigOptionalDto;
  }>;
export const useWebPushConfigBrowseAction = (
  options?: WebPushConfigBrowseActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WebPushConfigBrowseAction.Fetch(
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
 * WebPushConfigBrowseAction
 */
export class WebPushConfigBrowseAction {
  //
  static URL = "/webPushConfig/browse";
  static NewUrl = (qs?: WebPushConfigBrowseActionQueryParams) =>
    buildUrl(WebPushConfigBrowseAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: WebPushConfigBrowseActionQueryParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WebPushConfigOptionalDto>, unknown, unknown>(
      overrideUrl ?? WebPushConfigBrowseAction.NewUrl(qs),
      {
        method: WebPushConfigBrowseAction.Method,
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
      creatorFn?: ((item: unknown) => WebPushConfigOptionalDto) | undefined;
      qs?: WebPushConfigBrowseActionQueryParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WebPushConfigOptionalDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WebPushConfigOptionalDto(item));
    const res = await WebPushConfigBrowseAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WebPushConfigOptionalDto>();
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
    name: "webPushConfigBrowse",
    cliName: "browse",
    cliShort: "b",
    url: "/webPushConfig/browse",
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
      'Returns "webPushConfig" rows matching a filter, sorted/paged (see emigorm.ApplyQueryFilter/ApplyQueryScope).',
    out: {
      envelope: "GResponse",
      dto: "WebPushConfigOptionalDto",
    },
  };
}
/**
 * WebPushConfigBrowseActionQueryParams class
 * Auto-generated from EmiAction
 */
export class WebPushConfigBrowseActionQueryParams extends URLSearchParamsX {
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
