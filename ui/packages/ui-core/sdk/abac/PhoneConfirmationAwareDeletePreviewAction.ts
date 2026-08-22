import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { MArray } from "@fireback/js-remote-ctx/common/operators";
import { URLSearchParamsX } from "@fireback/js-remote-ctx/common/URLSearchParamsX";
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
 * Action to communicate with the action phoneConfirmationAwareDeletePreview
 */
export type PhoneConfirmationAwareDeletePreviewActionOptions = {
  queryKey?: unknown[];
  qs?: PhoneConfirmationAwareDeletePreviewActionQueryParams;
};
export type PhoneConfirmationAwareDeletePreviewActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<PhoneConfirmationAwareDeletePreviewActionRes>,
    unknown[]
  >,
  "queryKey"
> &
  PhoneConfirmationAwareDeletePreviewActionOptions &
  Partial<{
    creatorFn: (item: unknown) => PhoneConfirmationAwareDeletePreviewActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const usePhoneConfirmationAwareDeletePreviewActionQuery = (
  options: PhoneConfirmationAwareDeletePreviewActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return PhoneConfirmationAwareDeletePreviewAction.Fetch(
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
    queryKey: [PhoneConfirmationAwareDeletePreviewAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type PhoneConfirmationAwareDeletePreviewActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PhoneConfirmationAwareDeletePreviewActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => PhoneConfirmationAwareDeletePreviewActionRes;
  }>;
export const usePhoneConfirmationAwareDeletePreviewAction = (
  options?: PhoneConfirmationAwareDeletePreviewActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return PhoneConfirmationAwareDeletePreviewAction.Fetch(
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
 * PhoneConfirmationAwareDeletePreviewAction
 */
export class PhoneConfirmationAwareDeletePreviewAction {
  //
  static URL = "/phoneConfirmation/delete-preview";
  static NewUrl = (qs?: PhoneConfirmationAwareDeletePreviewActionQueryParams) =>
    buildUrl(PhoneConfirmationAwareDeletePreviewAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: PhoneConfirmationAwareDeletePreviewActionQueryParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<PhoneConfirmationAwareDeletePreviewActionRes>,
      unknown,
      unknown
    >(
      overrideUrl ?? PhoneConfirmationAwareDeletePreviewAction.NewUrl(qs),
      {
        method: PhoneConfirmationAwareDeletePreviewAction.Method,
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
        | ((item: unknown) => PhoneConfirmationAwareDeletePreviewActionRes)
        | undefined;
      qs?: PhoneConfirmationAwareDeletePreviewActionQueryParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) =>
        new PhoneConfirmationAwareDeletePreviewActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn ||
      ((item) => new PhoneConfirmationAwareDeletePreviewActionRes(item));
    const res = await PhoneConfirmationAwareDeletePreviewAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp =
          new GResponse<PhoneConfirmationAwareDeletePreviewActionRes>();
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
    name: "phoneConfirmationAwareDeletePreview",
    cliName: "delete-preview",
    cliShort: "dp",
    url: "/phoneConfirmation/delete-preview",
    method: "get",
    qs: [
      {
        name: "uniqueIds",
        type: "slice",
        primitive: "string",
      },
    ],
    description:
      'Reports what deleting the given "phoneConfirmation" uniqueIds would affect, without deleting anything.',
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
 * The base class definition for phoneConfirmationAwareDeletePreviewActionRes
 **/
export class PhoneConfirmationAwareDeletePreviewActionRes {
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
   * @type {PhoneConfirmationAwareDeletePreviewActionRes.Affected}
   **/
  #affected: MArray<
    InstanceType<typeof PhoneConfirmationAwareDeletePreviewActionRes.Affected>
  > = MArray.of([]);
  /**
   *
   * @returns {PhoneConfirmationAwareDeletePreviewActionRes.Affected}
   **/
  get affected() {
    return this.#affected;
  }
  /**
   *
   * @type {PhoneConfirmationAwareDeletePreviewActionRes.Affected}
   **/
  set affected(
    value:
      | MArray<
          InstanceType<
            typeof PhoneConfirmationAwareDeletePreviewActionRes.Affected
          >
        >
      | InstanceType<
          typeof PhoneConfirmationAwareDeletePreviewActionRes.Affected
        >[],
  ) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (
        value.length > 0 &&
        value[0] instanceof
          PhoneConfirmationAwareDeletePreviewActionRes.Affected
      ) {
        this.#affected = MArray.of(value);
      } else {
        this.#affected = MArray.of(
          value.map(
            (item) =>
              new PhoneConfirmationAwareDeletePreviewActionRes.Affected(item),
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
            typeof PhoneConfirmationAwareDeletePreviewActionRes.Affected
          >
        >
      | InstanceType<
          typeof PhoneConfirmationAwareDeletePreviewActionRes.Affected
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
     * Creates an instance of PhoneConfirmationAwareDeletePreviewActionRes.Affected, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(
      possibleDtoObject: PhoneConfirmationAwareDeletePreviewActionResType.AffectedType,
    ) {
      return new PhoneConfirmationAwareDeletePreviewActionRes.Affected(
        possibleDtoObject,
      );
    }
    /**
     * Creates an instance of PhoneConfirmationAwareDeletePreviewActionRes.Affected, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<PhoneConfirmationAwareDeletePreviewActionResType.AffectedType>,
    ) {
      return new PhoneConfirmationAwareDeletePreviewActionRes.Affected(
        partialDtoObject,
      );
    }
    copyWith(
      partial: PartialDeep<PhoneConfirmationAwareDeletePreviewActionResType.AffectedType>,
    ): InstanceType<
      typeof PhoneConfirmationAwareDeletePreviewActionRes.Affected
    > {
      return new PhoneConfirmationAwareDeletePreviewActionRes.Affected({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<
      typeof PhoneConfirmationAwareDeletePreviewActionRes.Affected
    > {
      return new PhoneConfirmationAwareDeletePreviewActionRes.Affected(
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
    const d = data as Partial<PhoneConfirmationAwareDeletePreviewActionRes>;
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
          PhoneConfirmationAwareDeletePreviewActionRes.Affected.Fields,
        );
      },
    };
  }
  /**
   * Creates an instance of PhoneConfirmationAwareDeletePreviewActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(
    possibleDtoObject: PhoneConfirmationAwareDeletePreviewActionResType,
  ) {
    return new PhoneConfirmationAwareDeletePreviewActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of PhoneConfirmationAwareDeletePreviewActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<PhoneConfirmationAwareDeletePreviewActionResType>,
  ) {
    return new PhoneConfirmationAwareDeletePreviewActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<PhoneConfirmationAwareDeletePreviewActionResType>,
  ): InstanceType<typeof PhoneConfirmationAwareDeletePreviewActionRes> {
    return new PhoneConfirmationAwareDeletePreviewActionRes({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof PhoneConfirmationAwareDeletePreviewActionRes> {
    return new PhoneConfirmationAwareDeletePreviewActionRes(this.toJSON());
  }
}
export abstract class PhoneConfirmationAwareDeletePreviewActionResFactory {
  abstract create(data: unknown): PhoneConfirmationAwareDeletePreviewActionRes;
}
/**
 * The base type definition for phoneConfirmationAwareDeletePreviewActionRes
 **/
export type PhoneConfirmationAwareDeletePreviewActionResType = {
  /**
   *
   * @type {string}
   **/
  message: string;
  /**
   *
   * @type {PhoneConfirmationAwareDeletePreviewActionResType.AffectedType[]}
   **/
  affected: PhoneConfirmationAwareDeletePreviewActionResType.AffectedType[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PhoneConfirmationAwareDeletePreviewActionResType {
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
 * PhoneConfirmationAwareDeletePreviewActionQueryParams class
 * Auto-generated from EmiAction
 */
export class PhoneConfirmationAwareDeletePreviewActionQueryParams extends URLSearchParamsX {
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
