import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { buildUrl } from "@fireback/js-remote-ctx/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
  type TypedRequestInit,
  type TypedResponse,
} from "@fireback/js-remote-ctx/common/fetchx";
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action QueryWorkspaceRoles
 */
export type QueryWorkspaceRolesActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type QueryWorkspaceRolesActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  QueryWorkspaceRolesActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => QueryWorkspaceRolesActionRes;
  }>;
export const useQueryWorkspaceRolesAction = (
  options?: QueryWorkspaceRolesActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: QueryWorkspaceRolesActionReq) => {
    setCompleteState(false);
    return QueryWorkspaceRolesAction.Fetch(
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
 * QueryWorkspaceRolesAction
 */
export class QueryWorkspaceRolesAction {
  //
  static URL = "/workspace/roles";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(QueryWorkspaceRolesAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<QueryWorkspaceRolesActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<QueryWorkspaceRolesActionRes>,
      QueryWorkspaceRolesActionReq,
      unknown
    >(
      overrideUrl ?? QueryWorkspaceRolesAction.NewUrl(qs),
      {
        method: QueryWorkspaceRolesAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<QueryWorkspaceRolesActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => QueryWorkspaceRolesActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new QueryWorkspaceRolesActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new QueryWorkspaceRolesActionRes(item));
    const res = await QueryWorkspaceRolesAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<QueryWorkspaceRolesActionRes>();
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
    name: "QueryWorkspaceRoles",
    cliName: "workspace-roles",
    url: "/workspace/roles",
    method: "post",
    description:
      "Lists the roles that actually belong to a given workspace - root only, and deliberately bypasses the normal per-workspace capability check the generic 'role browse'/GET /role/browse endpoint enforces (root has no real membership in most workspaces, so switching Workspace-Id to one it doesn't belong to would otherwise 401 there). Use this - not root's own role list - to populate a role picker for AddUserToWorkspace/ChangeUserWorkspaceRole.",
    in: {
      fields: [
        {
          name: "workspaceId",
          description: "UniqueId of the workspace whose roles to list.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
      ],
    },
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "uniqueId",
          type: "string",
        },
        {
          name: "name",
          type: "string",
        },
        {
          name: "capabilitiesListId",
          type: "slice",
          primitive: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for queryWorkspaceRolesActionReq
 **/
export class QueryWorkspaceRolesActionReq {
  /**
   * UniqueId of the workspace whose roles to list.
   * @type {string}
   **/
  #workspaceId: string = "";
  /**
   * UniqueId of the workspace whose roles to list.
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   * UniqueId of the workspace whose roles to list.
   * @type {string}
   **/
  set workspaceId(value: string) {
    this.#workspaceId = String(value);
  }
  setWorkspaceId(value: string) {
    this.workspaceId = value;
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
    const d = data as Partial<QueryWorkspaceRolesActionReq>;
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      workspaceId: this.#workspaceId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      workspaceId: "workspaceId",
    };
  }
  /**
   * Creates an instance of QueryWorkspaceRolesActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: QueryWorkspaceRolesActionReqType) {
    return new QueryWorkspaceRolesActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of QueryWorkspaceRolesActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<QueryWorkspaceRolesActionReqType>) {
    return new QueryWorkspaceRolesActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<QueryWorkspaceRolesActionReqType>,
  ): InstanceType<typeof QueryWorkspaceRolesActionReq> {
    return new QueryWorkspaceRolesActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof QueryWorkspaceRolesActionReq> {
    return new QueryWorkspaceRolesActionReq(this.toJSON());
  }
}
export abstract class QueryWorkspaceRolesActionReqFactory {
  abstract create(data: unknown): QueryWorkspaceRolesActionReq;
}
/**
 * The base type definition for queryWorkspaceRolesActionReq
 **/
export type QueryWorkspaceRolesActionReqType = {
  /**
   * UniqueId of the workspace whose roles to list.
   * @type {string}
   **/
  workspaceId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace QueryWorkspaceRolesActionReqType {}
/**
 * The base class definition for queryWorkspaceRolesActionRes
 **/
export class QueryWorkspaceRolesActionRes {
  /**
   *
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   *
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   *
   * @type {string}
   **/
  set uniqueId(value: string) {
    this.#uniqueId = String(value);
  }
  setUniqueId(value: string) {
    this.uniqueId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #name: string = "";
  /**
   *
   * @returns {string}
   **/
  get name() {
    return this.#name;
  }
  /**
   *
   * @type {string}
   **/
  set name(value: string) {
    this.#name = String(value);
  }
  setName(value: string) {
    this.name = value;
    return this;
  }
  /**
   *
   * @type {string[]}
   **/
  #capabilitiesListId: string[] = [];
  /**
   *
   * @returns {string[]}
   **/
  get capabilitiesListId() {
    return this.#capabilitiesListId;
  }
  /**
   *
   * @type {string[]}
   **/
  set capabilitiesListId(value: string[]) {
    this.#capabilitiesListId = value;
  }
  setCapabilitiesListId(value: string[]) {
    this.capabilitiesListId = value;
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
    const d = data as Partial<QueryWorkspaceRolesActionRes>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.name !== undefined) {
      this.name = d.name;
    }
    if (d.capabilitiesListId !== undefined) {
      this.capabilitiesListId = d.capabilitiesListId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      name: this.#name,
      capabilitiesListId: this.#capabilitiesListId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      name: "name",
      capabilitiesListId$: "capabilitiesListId",
      get capabilitiesListId() {
        return "capabilitiesListId[:i]";
      },
    };
  }
  /**
   * Creates an instance of QueryWorkspaceRolesActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: QueryWorkspaceRolesActionResType) {
    return new QueryWorkspaceRolesActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of QueryWorkspaceRolesActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<QueryWorkspaceRolesActionResType>) {
    return new QueryWorkspaceRolesActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<QueryWorkspaceRolesActionResType>,
  ): InstanceType<typeof QueryWorkspaceRolesActionRes> {
    return new QueryWorkspaceRolesActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof QueryWorkspaceRolesActionRes> {
    return new QueryWorkspaceRolesActionRes(this.toJSON());
  }
}
export abstract class QueryWorkspaceRolesActionResFactory {
  abstract create(data: unknown): QueryWorkspaceRolesActionRes;
}
/**
 * The base type definition for queryWorkspaceRolesActionRes
 **/
export type QueryWorkspaceRolesActionResType = {
  /**
   *
   * @type {string}
   **/
  uniqueId: string;
  /**
   *
   * @type {string}
   **/
  name: string;
  /**
   *
   * @type {string[]}
   **/
  capabilitiesListId: string[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace QueryWorkspaceRolesActionResType {}
