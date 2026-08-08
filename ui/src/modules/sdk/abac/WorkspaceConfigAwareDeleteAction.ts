import { buildUrl } from "../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
  type TypedRequestInit,
  type TypedResponse,
} from "../sdk/common/fetchx";
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action workspaceConfigAwareDelete
 */
export type WorkspaceConfigAwareDeleteActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type WorkspaceConfigAwareDeleteActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WorkspaceConfigAwareDeleteActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  };
export const useWorkspaceConfigAwareDeleteAction = (
  options?: WorkspaceConfigAwareDeleteActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: WorkspaceConfigAwareDeleteActionReq) => {
    setCompleteState(false);
    return WorkspaceConfigAwareDeleteAction.Fetch(
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
 * WorkspaceConfigAwareDeleteAction
 */
export class WorkspaceConfigAwareDeleteAction {
  //
  static URL = "/workspaceConfig/delete";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WorkspaceConfigAwareDeleteAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<WorkspaceConfigAwareDeleteActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<unknown, WorkspaceConfigAwareDeleteActionReq, unknown>(
      overrideUrl ?? WorkspaceConfigAwareDeleteAction.NewUrl(qs),
      {
        method: WorkspaceConfigAwareDeleteAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<WorkspaceConfigAwareDeleteActionReq, unknown>,
    {
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {},
  ) => {
    const res = await WorkspaceConfigAwareDeleteAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(res, undefined, onMessage, init?.signal);
  };
  static Definition = {
    name: "workspaceConfigAwareDelete",
    cliShort: "workspaceConfig-d",
    url: "/workspaceConfig/delete",
    method: "post",
    description:
      'Deletes the given "workspaceConfig" uniqueIds, along with everything workspaceConfigAwareDeletePreview reports.',
    in: {
      fields: [
        {
          name: "uniqueIds",
          type: "slice",
          primitive: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for workspaceConfigAwareDeleteActionReq
 **/
export class WorkspaceConfigAwareDeleteActionReq {
  /**
   *
   * @type {string[]}
   **/
  #uniqueIds: string[] = [];
  /**
   *
   * @returns {string[]}
   **/
  get uniqueIds() {
    return this.#uniqueIds;
  }
  /**
   *
   * @type {string[]}
   **/
  set uniqueIds(value: string[]) {
    this.#uniqueIds = value;
  }
  setUniqueIds(value: string[]) {
    this.uniqueIds = value;
    return this;
  }
  constructor(data: unknown = undefined) {
    if (data === null || data === undefined) {
      return;
    }
    if (typeof data === "string") {
      this.applyFromObject(JSON.parse(data));
    } else if (this.#isJsonAppliable(data)) {
      this.applyFromObject(data);
    } else {
      throw new Error(
        "Instance cannot be created on an unknown value, check the content being passed. got: " +
          typeof data,
      );
    }
  }
  #isJsonAppliable(obj: unknown) {
    const g = globalThis as unknown as { Buffer: any; Blob: any };
    const isBuffer =
      typeof g.Buffer !== "undefined" &&
      typeof g.Buffer.isBuffer === "function" &&
      g.Buffer.isBuffer(obj);
    const isBlob = typeof g.Blob !== "undefined" && obj instanceof g.Blob;
    return (
      obj &&
      typeof obj === "object" &&
      !Array.isArray(obj) &&
      !isBuffer &&
      !(obj instanceof ArrayBuffer) &&
      !isBlob
    );
  }
  /**
   * casts the fields of a javascript object into the class properties one by one
   **/
  applyFromObject(data = {}) {
    const d = data as Partial<WorkspaceConfigAwareDeleteActionReq>;
    if (d.uniqueIds !== undefined) {
      this.uniqueIds = d.uniqueIds;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueIds: this.#uniqueIds,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueIds$: "uniqueIds",
      get uniqueIds() {
        return "uniqueIds[:i]";
      },
    };
  }
  /**
   * Creates an instance of WorkspaceConfigAwareDeleteActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WorkspaceConfigAwareDeleteActionReqType) {
    return new WorkspaceConfigAwareDeleteActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of WorkspaceConfigAwareDeleteActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<WorkspaceConfigAwareDeleteActionReqType>,
  ) {
    return new WorkspaceConfigAwareDeleteActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WorkspaceConfigAwareDeleteActionReqType>,
  ): InstanceType<typeof WorkspaceConfigAwareDeleteActionReq> {
    return new WorkspaceConfigAwareDeleteActionReq({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof WorkspaceConfigAwareDeleteActionReq> {
    return new WorkspaceConfigAwareDeleteActionReq(this.toJSON());
  }
}
export abstract class WorkspaceConfigAwareDeleteActionReqFactory {
  abstract create(data: unknown): WorkspaceConfigAwareDeleteActionReq;
}
/**
 * The base type definition for workspaceConfigAwareDeleteActionReq
 **/
export type WorkspaceConfigAwareDeleteActionReqType = {
  /**
   *
   * @type {string[]}
   **/
  uniqueIds: string[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WorkspaceConfigAwareDeleteActionReqType {}
