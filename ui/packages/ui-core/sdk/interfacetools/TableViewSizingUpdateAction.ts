import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { TableViewSizingDto } from "./TableViewSizingDto";
import { TableViewSizingOptionalDto } from "./TableViewSizingOptionalDto";
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
 * Action to communicate with the action tableViewSizingUpdate
 */
export type TableViewSizingUpdateActionOptions = {
  queryKey?: unknown[];
  params: TableViewSizingUpdateActionPathParameter;
  qs?: URLSearchParams;
};
export type TableViewSizingUpdateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  TableViewSizingUpdateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => TableViewSizingDto;
  }>;
export const useTableViewSizingUpdateAction = (
  options: TableViewSizingUpdateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: TableViewSizingOptionalDto) => {
    setCompleteState(false);
    return TableViewSizingUpdateAction.Fetch(
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
 * Path parameters for TableViewSizingUpdateAction
 */
export type TableViewSizingUpdateActionPathParameter = {
  uniqueId: string;
};
/**
 * TableViewSizingUpdateAction
 */
export class TableViewSizingUpdateAction {
  //
  static URL = "/tableViewSizing/:uniqueId";
  static NewUrl = (
    params: TableViewSizingUpdateActionPathParameter,
    qs?: URLSearchParams,
  ) => buildUrl(TableViewSizingUpdateAction.URL, params, qs);
  static Method = "PATCH";
  static Fetch$ = async (
    params: TableViewSizingUpdateActionPathParameter,
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<TableViewSizingOptionalDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<TableViewSizingDto>,
      TableViewSizingOptionalDto,
      unknown
    >(
      overrideUrl ?? TableViewSizingUpdateAction.NewUrl(params, qs),
      {
        method: TableViewSizingUpdateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    params: TableViewSizingUpdateActionPathParameter,
    init?: TypedRequestInit<TableViewSizingOptionalDto, unknown>,
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
    const res = await TableViewSizingUpdateAction.Fetch$(
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
    name: "tableViewSizingUpdate",
    cliName: "update",
    cliShort: "tableViewSizing-u",
    url: "/tableViewSizing/:uniqueId string",
    method: "patch",
    description:
      'Applies a partial update to a "tableViewSizing" row by uniqueId.',
    in: {
      dto: "TableViewSizingOptionalDto",
    },
    out: {
      envelope: "GResponse",
      dto: "TableViewSizingDto",
    },
  };
}
