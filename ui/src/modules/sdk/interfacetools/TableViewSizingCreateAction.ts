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
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action tableViewSizingCreate
 */
export type TableViewSizingCreateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type TableViewSizingCreateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  TableViewSizingCreateActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => TableViewSizingDto;
  }>;
export const useTableViewSizingCreateAction = (
  options?: TableViewSizingCreateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: TableViewSizingDto) => {
    setCompleteState(false);
    return TableViewSizingCreateAction.Fetch(
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
 * TableViewSizingCreateAction
 */
export class TableViewSizingCreateAction {
  //
  static URL = "/tableViewSizing";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(TableViewSizingCreateAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<TableViewSizingDto, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<TableViewSizingDto>, TableViewSizingDto, unknown>(
      overrideUrl ?? TableViewSizingCreateAction.NewUrl(qs),
      {
        method: TableViewSizingCreateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<TableViewSizingDto, unknown>,
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
    const res = await TableViewSizingCreateAction.Fetch$(
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
    name: "tableViewSizingCreate",
    cliName: "create",
    cliShort: "tableViewSizing-c",
    url: "/tableViewSizing",
    method: "post",
    description: 'Creates a new "tableViewSizing" row.',
    in: {
      dto: "TableViewSizingDto",
    },
    out: {
      envelope: "GResponse",
      dto: "TableViewSizingDto",
    },
  };
}
