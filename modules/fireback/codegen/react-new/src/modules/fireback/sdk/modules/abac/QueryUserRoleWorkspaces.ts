import { GResponse } from "../../sdk/envelopes/index";
import { buildUrl } from "../../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
  type TypedRequestInit,
  type TypedResponse,
} from "../../sdk/common/fetchx";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "react-query";
import { useFetchxContext } from "../../sdk/react/useFetchx";
import { useState } from "react";
import { withPrefix } from "../../sdk/common/withPrefix";
/**
 * Action to communicate with the action QueryUserRoleWorkspaces
 */
export type QueryUserRoleWorkspacesActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type QueryUserRoleWorkspacesActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<QueryUserRoleWorkspacesActionRes>,
    unknown[]
  >,
  "queryKey"
> &
  QueryUserRoleWorkspacesActionOptions &
  Partial<{
    creatorFn: (item: unknown) => QueryUserRoleWorkspacesActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext;
  };
export const useQueryUserRoleWorkspacesActionQuery = (
  options: QueryUserRoleWorkspacesActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return QueryUserRoleWorkspacesAction.Fetch(
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
    queryKey: [QueryUserRoleWorkspacesAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type QueryUserRoleWorkspacesActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  QueryUserRoleWorkspacesActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => QueryUserRoleWorkspacesActionRes;
  }>;
export const useQueryUserRoleWorkspacesAction = (
  options?: QueryUserRoleWorkspacesActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return QueryUserRoleWorkspacesAction.Fetch(
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
 * QueryUserRoleWorkspacesAction
 */
export class QueryUserRoleWorkspacesAction {
  //
  static URL = "/urw/query";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(QueryUserRoleWorkspacesAction.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<QueryUserRoleWorkspacesActionRes>,
      unknown,
      unknown
    >(
      overrideUrl ?? QueryUserRoleWorkspacesAction.NewUrl(qs),
      {
        method: QueryUserRoleWorkspacesAction.Method,
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
        | ((item: unknown) => QueryUserRoleWorkspacesActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new QueryUserRoleWorkspacesActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new QueryUserRoleWorkspacesActionRes(item));
    const res = await QueryUserRoleWorkspacesAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<QueryUserRoleWorkspacesActionRes>();
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
    name: "QueryUserRoleWorkspaces",
    cliName: "urw",
    url: "/urw/query",
    method: "get",
    description:
      "Returns the workspaces that user belongs to, as well as his role in there, and the permissions for each role",
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "name",
          type: "string",
        },
        {
          name: "capabilities",
          description: "Workspace level capabilities which are available",
          type: "slice",
          primitive: "string",
        },
        {
          name: "uniqueId",
          type: "string",
        },
        {
          name: "roles",
          type: "array",
          fields: [
            {
              name: "name",
              type: "string",
            },
            {
              name: "uniqueId",
              type: "string",
            },
            {
              name: "capabilities",
              description:
                "Capabilities related to this role which are available",
              type: "slice",
              primitive: "string",
            },
          ],
        },
      ],
    },
  };
}
/**
 * The base class definition for queryUserRoleWorkspacesActionRes
 **/
export class QueryUserRoleWorkspacesActionRes {
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
   * Workspace level capabilities which are available
   * @type {string[]}
   **/
  #capabilities: string[] = [];
  /**
   * Workspace level capabilities which are available
   * @returns {string[]}
   **/
  get capabilities() {
    return this.#capabilities;
  }
  /**
   * Workspace level capabilities which are available
   * @type {string[]}
   **/
  set capabilities(value: string[]) {
    this.#capabilities = value;
  }
  setCapabilities(value: string[]) {
    this.capabilities = value;
    return this;
  }
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
   * @type {QueryUserRoleWorkspacesActionRes.Roles}
   **/
  #roles: InstanceType<typeof QueryUserRoleWorkspacesActionRes.Roles>[] = [];
  /**
   *
   * @returns {QueryUserRoleWorkspacesActionRes.Roles}
   **/
  get roles() {
    return this.#roles;
  }
  /**
   *
   * @type {QueryUserRoleWorkspacesActionRes.Roles}
   **/
  set roles(
    value: InstanceType<typeof QueryUserRoleWorkspacesActionRes.Roles>[],
  ) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (
      value.length > 0 &&
      value[0] instanceof QueryUserRoleWorkspacesActionRes.Roles
    ) {
      this.#roles = value;
    } else {
      this.#roles = value.map(
        (item) => new QueryUserRoleWorkspacesActionRes.Roles(item),
      );
    }
  }
  setRoles(
    value: InstanceType<typeof QueryUserRoleWorkspacesActionRes.Roles>[],
  ) {
    this.roles = value;
    return this;
  }
  /**
   * The base class definition for roles
   **/
  static Roles = class Roles {
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
     * Capabilities related to this role which are available
     * @type {string[]}
     **/
    #capabilities: string[] = [];
    /**
     * Capabilities related to this role which are available
     * @returns {string[]}
     **/
    get capabilities() {
      return this.#capabilities;
    }
    /**
     * Capabilities related to this role which are available
     * @type {string[]}
     **/
    set capabilities(value: string[]) {
      this.#capabilities = value;
    }
    setCapabilities(value: string[]) {
      this.capabilities = value;
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
      const d = data as Partial<Roles>;
      if (d.name !== undefined) {
        this.name = d.name;
      }
      if (d.uniqueId !== undefined) {
        this.uniqueId = d.uniqueId;
      }
      if (d.capabilities !== undefined) {
        this.capabilities = d.capabilities;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        name: this.#name,
        uniqueId: this.#uniqueId,
        capabilities: this.#capabilities,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        name: "name",
        uniqueId: "uniqueId",
        capabilities$: "capabilities",
        get capabilities() {
          return "roles.capabilities[:i]";
        },
      };
    }
    /**
     * Creates an instance of QueryUserRoleWorkspacesActionRes.Roles, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(
      possibleDtoObject: QueryUserRoleWorkspacesActionResType.RolesType,
    ) {
      return new QueryUserRoleWorkspacesActionRes.Roles(possibleDtoObject);
    }
    /**
     * Creates an instance of QueryUserRoleWorkspacesActionRes.Roles, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<QueryUserRoleWorkspacesActionResType.RolesType>,
    ) {
      return new QueryUserRoleWorkspacesActionRes.Roles(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<QueryUserRoleWorkspacesActionResType.RolesType>,
    ): InstanceType<typeof QueryUserRoleWorkspacesActionRes.Roles> {
      return new QueryUserRoleWorkspacesActionRes.Roles({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<typeof QueryUserRoleWorkspacesActionRes.Roles> {
      return new QueryUserRoleWorkspacesActionRes.Roles(this.toJSON());
    }
  };
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
    const d = data as Partial<QueryUserRoleWorkspacesActionRes>;
    if (d.name !== undefined) {
      this.name = d.name;
    }
    if (d.capabilities !== undefined) {
      this.capabilities = d.capabilities;
    }
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.roles !== undefined) {
      this.roles = d.roles;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      name: this.#name,
      capabilities: this.#capabilities,
      uniqueId: this.#uniqueId,
      roles: this.#roles,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      name: "name",
      capabilities$: "capabilities",
      get capabilities() {
        return "capabilities[:i]";
      },
      uniqueId: "uniqueId",
      roles$: "roles",
      get roles() {
        return withPrefix(
          "roles[:i]",
          QueryUserRoleWorkspacesActionRes.Roles.Fields,
        );
      },
    };
  }
  /**
   * Creates an instance of QueryUserRoleWorkspacesActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: QueryUserRoleWorkspacesActionResType) {
    return new QueryUserRoleWorkspacesActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of QueryUserRoleWorkspacesActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<QueryUserRoleWorkspacesActionResType>,
  ) {
    return new QueryUserRoleWorkspacesActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<QueryUserRoleWorkspacesActionResType>,
  ): InstanceType<typeof QueryUserRoleWorkspacesActionRes> {
    return new QueryUserRoleWorkspacesActionRes({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof QueryUserRoleWorkspacesActionRes> {
    return new QueryUserRoleWorkspacesActionRes(this.toJSON());
  }
}
export abstract class QueryUserRoleWorkspacesActionResFactory {
  abstract create(data: unknown): QueryUserRoleWorkspacesActionRes;
}
/**
 * The base type definition for queryUserRoleWorkspacesActionRes
 **/
export type QueryUserRoleWorkspacesActionResType = {
  /**
   *
   * @type {string}
   **/
  name: string;
  /**
   * Workspace level capabilities which are available
   * @type {string[]}
   **/
  capabilities: string[];
  /**
   *
   * @type {string}
   **/
  uniqueId: string;
  /**
   *
   * @type {QueryUserRoleWorkspacesActionResType.RolesType[]}
   **/
  roles: QueryUserRoleWorkspacesActionResType.RolesType[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace QueryUserRoleWorkspacesActionResType {
  /**
   * The base type definition for rolesType
   **/
  export type RolesType = {
    /**
     *
     * @type {string}
     **/
    name: string;
    /**
     *
     * @type {string}
     **/
    uniqueId: string;
    /**
     * Capabilities related to this role which are available
     * @type {string[]}
     **/
    capabilities: string[];
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace RolesType {}
}
