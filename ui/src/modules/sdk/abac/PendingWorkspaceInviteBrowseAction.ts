import { GResponse } from "../sdk/envelopes/index";
import { PendingWorkspaceInviteOptionalDto } from "./PendingWorkspaceInviteOptionalDto";
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
 * Action to communicate with the action pendingWorkspaceInviteBrowse
 */
export type PendingWorkspaceInviteBrowseActionOptions = {
  queryKey?: unknown[];
  qs?: PendingWorkspaceInviteBrowseActionQueryParams;
};
export type PendingWorkspaceInviteBrowseActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<PendingWorkspaceInviteOptionalDto>,
    unknown[]
  >,
  "queryKey"
> &
  PendingWorkspaceInviteBrowseActionOptions &
  Partial<{
    creatorFn: (item: unknown) => PendingWorkspaceInviteOptionalDto;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const usePendingWorkspaceInviteBrowseActionQuery = (
  options: PendingWorkspaceInviteBrowseActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return PendingWorkspaceInviteBrowseAction.Fetch(
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
    queryKey: [PendingWorkspaceInviteBrowseAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type PendingWorkspaceInviteBrowseActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PendingWorkspaceInviteBrowseActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PendingWorkspaceInviteOptionalDto;
  }>;
export const usePendingWorkspaceInviteBrowseAction = (
  options?: PendingWorkspaceInviteBrowseActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return PendingWorkspaceInviteBrowseAction.Fetch(
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
 * PendingWorkspaceInviteBrowseAction
 */
export class PendingWorkspaceInviteBrowseAction {
  //
  static URL = "/pendingWorkspaceInvite/browse";
  static NewUrl = (qs?: PendingWorkspaceInviteBrowseActionQueryParams) =>
    buildUrl(PendingWorkspaceInviteBrowseAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: PendingWorkspaceInviteBrowseActionQueryParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<PendingWorkspaceInviteOptionalDto>,
      unknown,
      unknown
    >(
      overrideUrl ?? PendingWorkspaceInviteBrowseAction.NewUrl(qs),
      {
        method: PendingWorkspaceInviteBrowseAction.Method,
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
      creatorFn?:
        | ((item: unknown) => PendingWorkspaceInviteOptionalDto)
        | undefined;
      qs?: PendingWorkspaceInviteBrowseActionQueryParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new PendingWorkspaceInviteOptionalDto(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new PendingWorkspaceInviteOptionalDto(item));
    const res = await PendingWorkspaceInviteBrowseAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<PendingWorkspaceInviteOptionalDto>();
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
    name: "pendingWorkspaceInviteBrowse",
    cliShort: "pendingWorkspaceInvite-b",
    url: "/pendingWorkspaceInvite/browse",
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
      'Returns "pendingWorkspaceInvite" rows matching a filter, sorted/paged (see emigorm.ApplyQueryFilter/ApplyQueryScope).',
    out: {
      envelope: "GResponse",
      dto: "PendingWorkspaceInviteOptionalDto",
    },
  };
}
/**
 * PendingWorkspaceInviteBrowseActionQueryParams class
 * Auto-generated from EmiAction
 */
export class PendingWorkspaceInviteBrowseActionQueryParams extends URLSearchParamsX {
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
