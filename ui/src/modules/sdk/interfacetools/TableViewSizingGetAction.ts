import { GResponse } from "../sdk/envelopes/index";
import { TableViewSizingDto } from "./TableViewSizingDto";
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
 * Action to communicate with the action tableViewSizingGet
 */
export type TableViewSizingGetActionOptions = {
  queryKey?: unknown[];
  params: TableViewSizingGetActionPathParameter;
  qs?: URLSearchParams;
};
export type TableViewSizingGetActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<TableViewSizingDto>, unknown[]>,
  "queryKey"
> &
  TableViewSizingGetActionOptions &
  Partial<{
    creatorFn: (item: unknown) => TableViewSizingDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useTableViewSizingGetActionQuery = (
  options: TableViewSizingGetActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return TableViewSizingGetAction.Fetch(
      options.params,
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
    queryKey: [TableViewSizingGetAction.NewUrl(options.params, options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type TableViewSizingGetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  TableViewSizingGetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => TableViewSizingDto;
  }>;
export const useTableViewSizingGetAction = (
  options: TableViewSizingGetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return TableViewSizingGetAction.Fetch(
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
 * Path parameters for TableViewSizingGetAction
 */
export type TableViewSizingGetActionPathParameter = {
  uniqueId: string;
};
/**
 * TableViewSizingGetAction
 */
export class TableViewSizingGetAction {
  //
  static URL = "/tableViewSizing/:uniqueId";
  static NewUrl = (
    params: TableViewSizingGetActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(TableViewSizingGetAction.URL, params, qs);
  static Method = "GET";
  static Fetch$ = async (
    params: TableViewSizingGetActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<TableViewSizingDto>, unknown, unknown>(
      overrideUrl ?? TableViewSizingGetAction.NewUrl(params, qs),
      {
        method: TableViewSizingGetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: TableViewSizingGetActionPathParameter,
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => TableViewSizingDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new TableViewSizingDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new TableViewSizingDto(item));
    const res = await TableViewSizingGetAction.Fetch$(
      params,
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<TableViewSizingDto>();
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
    name: "tableViewSizingGet",
    cliName: "get",
    cliShort: "tableViewSizing-g",
    url: "/tableViewSizing/:uniqueId string",
    method: "get",
    description: 'Looks up a single "tableViewSizing" row by uniqueId.',
    out: {
      envelope: "GResponse",
      dto: "TableViewSizingDto",
    },
  };
}
