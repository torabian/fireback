import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { MArray } from "@fireback/js-remote-ctx/common/operators";
import { buildUrl } from "@fireback/js-remote-ctx/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
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
import { withPrefix } from "@fireback/js-remote-ctx/common/withPrefix";
/**
 * Action to communicate with the action Whoami
 */
export type WhoamiActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type WhoamiActionQueryOptions = Omit<
  UseQueryOptions<unknown, unknown, GResponse<WhoamiActionRes>, unknown[]>,
  "queryKey"
> &
  WhoamiActionOptions &
  Partial<{
    creatorFn: (item: unknown) => WhoamiActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useWhoamiActionQuery = (options: WhoamiActionQueryOptions) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return WhoamiAction.Fetch(
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
    queryKey: [WhoamiAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type WhoamiActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  WhoamiActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WhoamiActionRes;
  }>;
export const useWhoamiAction = (options?: WhoamiActionMutationOptions) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return WhoamiAction.Fetch(
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
 * WhoamiAction
 */
export class WhoamiAction {
  //
  static URL = "/whoami";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(WhoamiAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WhoamiActionRes>, unknown, unknown>(
      overrideUrl ?? WhoamiAction.NewUrl(qs),
      {
        method: WhoamiAction.Method,
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
      creatorFn?: ((item: unknown) => WhoamiActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WhoamiActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WhoamiActionRes(item));
    const res = await WhoamiAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WhoamiActionRes>();
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
    name: "Whoami",
    cliName: "whoami",
    url: "/whoami",
    method: "get",
    description:
      "Returns information about the currently authenticated user - their userId, and every workspace they belong to along with their role(s) and capabilities in each (the same UserRoleWorkspace shape QueryUserRoleWorkspaces returns).",
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "userId",
          type: "string",
        },
        {
          name: "workspaces",
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
              description: "Workspace level capabilities which are available",
              type: "slice",
              primitive: "string",
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
      ],
    },
  };
}
/**
 * The base class definition for whoamiActionRes
 **/
export class WhoamiActionRes {
  /**
   *
   * @type {string}
   **/
  #userId: string = "";
  /**
   *
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   *
   * @type {string}
   **/
  set userId(value: string) {
    this.#userId = String(value);
  }
  setUserId(value: string) {
    this.userId = value;
    return this;
  }
  /**
   *
   * @type {WhoamiActionRes.Workspaces}
   **/
  #workspaces: MArray<InstanceType<typeof WhoamiActionRes.Workspaces>> =
    MArray.of([]);
  /**
   *
   * @returns {WhoamiActionRes.Workspaces}
   **/
  get workspaces() {
    return this.#workspaces;
  }
  /**
   *
   * @type {WhoamiActionRes.Workspaces}
   **/
  set workspaces(
    value:
      | MArray<InstanceType<typeof WhoamiActionRes.Workspaces>>
      | InstanceType<typeof WhoamiActionRes.Workspaces>[],
  ) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof WhoamiActionRes.Workspaces) {
        this.#workspaces = MArray.of(value);
      } else {
        this.#workspaces = MArray.of(
          value.map((item) => new WhoamiActionRes.Workspaces(item)),
        );
      }
      return;
    }
    // If the instance is already an MArray, we assume it's all good.
    if (value instanceof MArray) {
      this.#workspaces = value;
      return;
    }
    // If the value is not array, and is not a MArray, we need to be consider,
    // it might be eligible to be casted into MArray.
    const { ok, value: mcastValue } = MArray.cast<unknown>(value);
    if (ok) {
      this.#workspaces = mcastValue as any;
      return;
    }
    console.warn(
      "Cannot assing value to workspaces, because it needs MArray instance or an Array.",
    );
  }
  setWorkspaces(
    value:
      | MArray<InstanceType<typeof WhoamiActionRes.Workspaces>>
      | InstanceType<typeof WhoamiActionRes.Workspaces>[],
  ) {
    this.workspaces = value;
    return this;
  }
  /**
   * The base class definition for workspaces
   **/
  static Workspaces = class Workspaces {
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
     * @type {WhoamiActionRes.Workspaces.Roles}
     **/
    #roles: MArray<InstanceType<typeof WhoamiActionRes.Workspaces.Roles>> =
      MArray.of([]);
    /**
     *
     * @returns {WhoamiActionRes.Workspaces.Roles}
     **/
    get roles() {
      return this.#roles;
    }
    /**
     *
     * @type {WhoamiActionRes.Workspaces.Roles}
     **/
    set roles(
      value:
        | MArray<InstanceType<typeof WhoamiActionRes.Workspaces.Roles>>
        | InstanceType<typeof WhoamiActionRes.Workspaces.Roles>[],
    ) {
      // When the passed value is already an array, we check if we need to
      // cast the inner items into class instance.
      if (Array.isArray(value)) {
        if (
          value.length > 0 &&
          value[0] instanceof WhoamiActionRes.Workspaces.Roles
        ) {
          this.#roles = MArray.of(value);
        } else {
          this.#roles = MArray.of(
            value.map((item) => new WhoamiActionRes.Workspaces.Roles(item)),
          );
        }
        return;
      }
      // If the instance is already an MArray, we assume it's all good.
      if (value instanceof MArray) {
        this.#roles = value;
        return;
      }
      // If the value is not array, and is not a MArray, we need to be consider,
      // it might be eligible to be casted into MArray.
      const { ok, value: mcastValue } = MArray.cast<unknown>(value);
      if (ok) {
        this.#roles = mcastValue as any;
        return;
      }
      console.warn(
        "Cannot assing value to roles, because it needs MArray instance or an Array.",
      );
    }
    setRoles(
      value:
        | MArray<InstanceType<typeof WhoamiActionRes.Workspaces.Roles>>
        | InstanceType<typeof WhoamiActionRes.Workspaces.Roles>[],
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
            return "workspaces.roles.capabilities[:i]";
          },
        };
      }
      /**
       * Creates an instance of WhoamiActionRes.Workspaces.Roles, and possibleDtoObject
       * needs to satisfy the type requirement fully, otherwise typescript compile would
       * be complaining.
       **/
      static from(
        possibleDtoObject: WhoamiActionResType.WorkspacesType.RolesType,
      ) {
        return new WhoamiActionRes.Workspaces.Roles(possibleDtoObject);
      }
      /**
       * Creates an instance of WhoamiActionRes.Workspaces.Roles, and partialDtoObject
       * needs to satisfy the type, but partially, and rest of the content would
       * be constructed according to data types and nullability.
       **/
      static with(
        partialDtoObject: PartialDeep<WhoamiActionResType.WorkspacesType.RolesType>,
      ) {
        return new WhoamiActionRes.Workspaces.Roles(partialDtoObject);
      }
      copyWith(
        partial: PartialDeep<WhoamiActionResType.WorkspacesType.RolesType>,
      ): InstanceType<typeof WhoamiActionRes.Workspaces.Roles> {
        return new WhoamiActionRes.Workspaces.Roles({
          ...this.toJSON(),
          ...partial,
        });
      }
      clone(): InstanceType<typeof WhoamiActionRes.Workspaces.Roles> {
        return new WhoamiActionRes.Workspaces.Roles(this.toJSON());
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
      const d = data as Partial<Workspaces>;
      if (d.name !== undefined) {
        this.name = d.name;
      }
      if (d.uniqueId !== undefined) {
        this.uniqueId = d.uniqueId;
      }
      if (d.capabilities !== undefined) {
        this.capabilities = d.capabilities;
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
        uniqueId: this.#uniqueId,
        capabilities: this.#capabilities,
        roles: this.#roles,
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
          return "workspaces.capabilities[:i]";
        },
        roles$: "roles",
        get roles() {
          return withPrefix(
            "workspaces.roles[:i]",
            WhoamiActionRes.Workspaces.Roles.Fields,
          );
        },
      };
    }
    /**
     * Creates an instance of WhoamiActionRes.Workspaces, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: WhoamiActionResType.WorkspacesType) {
      return new WhoamiActionRes.Workspaces(possibleDtoObject);
    }
    /**
     * Creates an instance of WhoamiActionRes.Workspaces, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<WhoamiActionResType.WorkspacesType>,
    ) {
      return new WhoamiActionRes.Workspaces(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<WhoamiActionResType.WorkspacesType>,
    ): InstanceType<typeof WhoamiActionRes.Workspaces> {
      return new WhoamiActionRes.Workspaces({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof WhoamiActionRes.Workspaces> {
      return new WhoamiActionRes.Workspaces(this.toJSON());
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
    const d = data as Partial<WhoamiActionRes>;
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.workspaces !== undefined) {
      this.workspaces = d.workspaces;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      userId: this.#userId,
      workspaces: this.#workspaces,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      userId: "userId",
      workspaces$: "workspaces",
      get workspaces() {
        return withPrefix("workspaces[:i]", WhoamiActionRes.Workspaces.Fields);
      },
    };
  }
  /**
   * Creates an instance of WhoamiActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WhoamiActionResType) {
    return new WhoamiActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of WhoamiActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WhoamiActionResType>) {
    return new WhoamiActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WhoamiActionResType>,
  ): InstanceType<typeof WhoamiActionRes> {
    return new WhoamiActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WhoamiActionRes> {
    return new WhoamiActionRes(this.toJSON());
  }
}
export abstract class WhoamiActionResFactory {
  abstract create(data: unknown): WhoamiActionRes;
}
/**
 * The base type definition for whoamiActionRes
 **/
export type WhoamiActionResType = {
  /**
   *
   * @type {string}
   **/
  userId: string;
  /**
   *
   * @type {WhoamiActionResType.WorkspacesType[]}
   **/
  workspaces: WhoamiActionResType.WorkspacesType[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WhoamiActionResType {
  /**
   * The base type definition for workspacesType
   **/
  export type WorkspacesType = {
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
     * Workspace level capabilities which are available
     * @type {string[]}
     **/
    capabilities: string[];
    /**
     *
     * @type {WhoamiActionResType.WorkspacesType.RolesType[]}
     **/
    roles: WhoamiActionResType.WorkspacesType.RolesType[];
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace WorkspacesType {
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
}
