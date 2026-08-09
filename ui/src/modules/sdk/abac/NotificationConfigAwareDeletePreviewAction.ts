import { GResponse } from "../sdk/envelopes/index";
import { MArray } from "../sdk/common/operators";
import { URLSearchParamsX } from "../sdk/common/URLSearchParamsX";
import { buildUrl } from "../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
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
import { withPrefix } from "../sdk/common/withPrefix";
/**
 * Action to communicate with the action notificationConfigAwareDeletePreview
 */
export type NotificationConfigAwareDeletePreviewActionOptions = {
  queryKey?: unknown[];
  qs?: NotificationConfigAwareDeletePreviewActionQueryParams;
};
export type NotificationConfigAwareDeletePreviewActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<NotificationConfigAwareDeletePreviewActionRes>,
    unknown[]
  >,
  "queryKey"
> &
  NotificationConfigAwareDeletePreviewActionOptions &
  Partial<{
    creatorFn: (item: unknown) => NotificationConfigAwareDeletePreviewActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useNotificationConfigAwareDeletePreviewActionQuery = (
  options: NotificationConfigAwareDeletePreviewActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return NotificationConfigAwareDeletePreviewAction.Fetch(
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
    queryKey: [NotificationConfigAwareDeletePreviewAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type NotificationConfigAwareDeletePreviewActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  NotificationConfigAwareDeletePreviewActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => NotificationConfigAwareDeletePreviewActionRes;
  }>;
export const useNotificationConfigAwareDeletePreviewAction = (
  options?: NotificationConfigAwareDeletePreviewActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return NotificationConfigAwareDeletePreviewAction.Fetch(
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
 * NotificationConfigAwareDeletePreviewAction
 */
export class NotificationConfigAwareDeletePreviewAction {
  //
  static URL = "/notificationConfig/delete-preview";
  static NewUrl = (
    qs?: NotificationConfigAwareDeletePreviewActionQueryParams,
  ) => buildUrl(NotificationConfigAwareDeletePreviewAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: NotificationConfigAwareDeletePreviewActionQueryParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<NotificationConfigAwareDeletePreviewActionRes>,
      unknown,
      unknown
    >(
      overrideUrl ?? NotificationConfigAwareDeletePreviewAction.NewUrl(qs),
      {
        method: NotificationConfigAwareDeletePreviewAction.Method,
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
        | ((item: unknown) => NotificationConfigAwareDeletePreviewActionRes)
        | undefined;
      qs?: NotificationConfigAwareDeletePreviewActionQueryParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) =>
        new NotificationConfigAwareDeletePreviewActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn ||
      ((item) => new NotificationConfigAwareDeletePreviewActionRes(item));
    const res = await NotificationConfigAwareDeletePreviewAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp =
          new GResponse<NotificationConfigAwareDeletePreviewActionRes>();
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
    name: "notificationConfigAwareDeletePreview",
    cliName: "delete-preview",
    cliShort: "notificationConfig-dp",
    url: "/notificationConfig/delete-preview",
    method: "get",
    qs: [
      {
        name: "uniqueIds",
        type: "slice",
        primitive: "string",
      },
    ],
    description:
      'Reports what deleting the given "notificationConfig" uniqueIds would affect, without deleting anything.',
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "message",
          type: "string",
        },
        {
          name: "affected",
          type: "array",
          fields: [
            {
              name: "relation",
              type: "string",
            },
            {
              name: "count",
              type: "int64",
            },
          ],
        },
      ],
    },
  };
}
/**
 * The base class definition for notificationConfigAwareDeletePreviewActionRes
 **/
export class NotificationConfigAwareDeletePreviewActionRes {
  /**
   *
   * @type {string}
   **/
  #message: string = "";
  /**
   *
   * @returns {string}
   **/
  get message() {
    return this.#message;
  }
  /**
   *
   * @type {string}
   **/
  set message(value: string) {
    this.#message = String(value);
  }
  setMessage(value: string) {
    this.message = value;
    return this;
  }
  /**
   *
   * @type {NotificationConfigAwareDeletePreviewActionRes.Affected}
   **/
  #affected: MArray<
    InstanceType<typeof NotificationConfigAwareDeletePreviewActionRes.Affected>
  > = MArray.of([]);
  /**
   *
   * @returns {NotificationConfigAwareDeletePreviewActionRes.Affected}
   **/
  get affected() {
    return this.#affected;
  }
  /**
   *
   * @type {NotificationConfigAwareDeletePreviewActionRes.Affected}
   **/
  set affected(
    value:
      | MArray<
          InstanceType<
            typeof NotificationConfigAwareDeletePreviewActionRes.Affected
          >
        >
      | InstanceType<
          typeof NotificationConfigAwareDeletePreviewActionRes.Affected
        >[],
  ) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (
        value.length > 0 &&
        value[0] instanceof
          NotificationConfigAwareDeletePreviewActionRes.Affected
      ) {
        this.#affected = MArray.of(value);
      } else {
        this.#affected = MArray.of(
          value.map(
            (item) =>
              new NotificationConfigAwareDeletePreviewActionRes.Affected(item),
          ),
        );
      }
      return;
    }
    // If the instance is already an MArray, we assume it's all good.
    if (value instanceof MArray) {
      this.#affected = value;
      return;
    }
    // If the value is not array, and is not a MArray, we need to be consider,
    // it might be eligible to be casted into MArray.
    const { ok, value: mcastValue } = MArray.cast<unknown>(value);
    if (ok) {
      this.#affected = mcastValue as any;
      return;
    }
    console.warn(
      "Cannot assing value to affected, because it needs MArray instance or an Array.",
    );
  }
  setAffected(
    value:
      | MArray<
          InstanceType<
            typeof NotificationConfigAwareDeletePreviewActionRes.Affected
          >
        >
      | InstanceType<
          typeof NotificationConfigAwareDeletePreviewActionRes.Affected
        >[],
  ) {
    this.affected = value;
    return this;
  }
  /**
   * The base class definition for affected
   **/
  static Affected = class Affected {
    /**
     *
     * @type {string}
     **/
    #relation: string = "";
    /**
     *
     * @returns {string}
     **/
    get relation() {
      return this.#relation;
    }
    /**
     *
     * @type {string}
     **/
    set relation(value: string) {
      this.#relation = String(value);
    }
    setRelation(value: string) {
      this.relation = value;
      return this;
    }
    /**
     *
     * @type {number}
     **/
    #count: number = 0;
    /**
     *
     * @returns {number}
     **/
    get count() {
      return this.#count;
    }
    /**
     *
     * @type {number}
     **/
    set count(value: number) {
      const correctType = typeof value === "number";
      const parsedValue = correctType ? value : Number(value);
      if (!Number.isNaN(parsedValue)) {
        this.#count = parsedValue;
      }
    }
    setCount(value: number) {
      this.count = value;
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
      const d = data as Partial<Affected>;
      if (d.relation !== undefined) {
        this.relation = d.relation;
      }
      if (d.count !== undefined) {
        this.count = d.count;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        relation: this.#relation,
        count: this.#count,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        relation: "relation",
        count: "count",
      };
    }
    /**
     * Creates an instance of NotificationConfigAwareDeletePreviewActionRes.Affected, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(
      possibleDtoObject: NotificationConfigAwareDeletePreviewActionResType.AffectedType,
    ) {
      return new NotificationConfigAwareDeletePreviewActionRes.Affected(
        possibleDtoObject,
      );
    }
    /**
     * Creates an instance of NotificationConfigAwareDeletePreviewActionRes.Affected, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<NotificationConfigAwareDeletePreviewActionResType.AffectedType>,
    ) {
      return new NotificationConfigAwareDeletePreviewActionRes.Affected(
        partialDtoObject,
      );
    }
    copyWith(
      partial: PartialDeep<NotificationConfigAwareDeletePreviewActionResType.AffectedType>,
    ): InstanceType<
      typeof NotificationConfigAwareDeletePreviewActionRes.Affected
    > {
      return new NotificationConfigAwareDeletePreviewActionRes.Affected({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<
      typeof NotificationConfigAwareDeletePreviewActionRes.Affected
    > {
      return new NotificationConfigAwareDeletePreviewActionRes.Affected(
        this.toJSON(),
      );
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
    const d = data as Partial<NotificationConfigAwareDeletePreviewActionRes>;
    if (d.message !== undefined) {
      this.message = d.message;
    }
    if (d.affected !== undefined) {
      this.affected = d.affected;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      message: this.#message,
      affected: this.#affected,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      message: "message",
      affected$: "affected",
      get affected() {
        return withPrefix(
          "affected[:i]",
          NotificationConfigAwareDeletePreviewActionRes.Affected.Fields,
        );
      },
    };
  }
  /**
   * Creates an instance of NotificationConfigAwareDeletePreviewActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(
    possibleDtoObject: NotificationConfigAwareDeletePreviewActionResType,
  ) {
    return new NotificationConfigAwareDeletePreviewActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of NotificationConfigAwareDeletePreviewActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<NotificationConfigAwareDeletePreviewActionResType>,
  ) {
    return new NotificationConfigAwareDeletePreviewActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<NotificationConfigAwareDeletePreviewActionResType>,
  ): InstanceType<typeof NotificationConfigAwareDeletePreviewActionRes> {
    return new NotificationConfigAwareDeletePreviewActionRes({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof NotificationConfigAwareDeletePreviewActionRes> {
    return new NotificationConfigAwareDeletePreviewActionRes(this.toJSON());
  }
}
export abstract class NotificationConfigAwareDeletePreviewActionResFactory {
  abstract create(data: unknown): NotificationConfigAwareDeletePreviewActionRes;
}
/**
 * The base type definition for notificationConfigAwareDeletePreviewActionRes
 **/
export type NotificationConfigAwareDeletePreviewActionResType = {
  /**
   *
   * @type {string}
   **/
  message: string;
  /**
   *
   * @type {NotificationConfigAwareDeletePreviewActionResType.AffectedType[]}
   **/
  affected: NotificationConfigAwareDeletePreviewActionResType.AffectedType[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace NotificationConfigAwareDeletePreviewActionResType {
  /**
   * The base type definition for affectedType
   **/
  export type AffectedType = {
    /**
     *
     * @type {string}
     **/
    relation: string;
    /**
     *
     * @type {number}
     **/
    count: number;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace AffectedType {}
}
/**
 * NotificationConfigAwareDeletePreviewActionQueryParams class
 * Auto-generated from EmiAction
 */
export class NotificationConfigAwareDeletePreviewActionQueryParams extends URLSearchParamsX {
  /**
   *
   * @returns { any }
   */
  getUniqueIds() {
    return this.getTyped("uniqueIds", "any");
  }
  /**
   *
   * @param { any } value
   */
  setUniqueIds(value: any) {
    this.set("uniqueIds", value);
    return this;
  }
}
