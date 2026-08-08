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
 * Action to communicate with the action emailConfirmationAwareDeletePreview
 */
export type EmailConfirmationAwareDeletePreviewActionOptions = {
  queryKey?: unknown[];
  qs?: EmailConfirmationAwareDeletePreviewActionQueryParams;
};
export type EmailConfirmationAwareDeletePreviewActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<EmailConfirmationAwareDeletePreviewActionRes>,
    unknown[]
  >,
  "queryKey"
> &
  EmailConfirmationAwareDeletePreviewActionOptions &
  Partial<{
    creatorFn: (item: unknown) => EmailConfirmationAwareDeletePreviewActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext | null;
  };
export const useEmailConfirmationAwareDeletePreviewActionQuery = (
  options: EmailConfirmationAwareDeletePreviewActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return EmailConfirmationAwareDeletePreviewAction.Fetch(
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
    queryKey: [EmailConfirmationAwareDeletePreviewAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type EmailConfirmationAwareDeletePreviewActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  EmailConfirmationAwareDeletePreviewActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => EmailConfirmationAwareDeletePreviewActionRes;
  }>;
export const useEmailConfirmationAwareDeletePreviewAction = (
  options?: EmailConfirmationAwareDeletePreviewActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return EmailConfirmationAwareDeletePreviewAction.Fetch(
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
 * EmailConfirmationAwareDeletePreviewAction
 */
export class EmailConfirmationAwareDeletePreviewAction {
  //
  static URL = "/emailConfirmation/delete-preview";
  static NewUrl = (qs?: EmailConfirmationAwareDeletePreviewActionQueryParams) =>
    buildUrl(EmailConfirmationAwareDeletePreviewAction.URL, undefined, qs);
  static Method = "GET";
  static Fetch$ = async (
    qs?: EmailConfirmationAwareDeletePreviewActionQueryParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<EmailConfirmationAwareDeletePreviewActionRes>,
      unknown,
      unknown
    >(
      overrideUrl ?? EmailConfirmationAwareDeletePreviewAction.NewUrl(qs),
      {
        method: EmailConfirmationAwareDeletePreviewAction.Method,
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
        | ((item: unknown) => EmailConfirmationAwareDeletePreviewActionRes)
        | undefined;
      qs?: EmailConfirmationAwareDeletePreviewActionQueryParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) =>
        new EmailConfirmationAwareDeletePreviewActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn ||
      ((item) => new EmailConfirmationAwareDeletePreviewActionRes(item));
    const res = await EmailConfirmationAwareDeletePreviewAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp =
          new GResponse<EmailConfirmationAwareDeletePreviewActionRes>();
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
    name: "emailConfirmationAwareDeletePreview",
    cliShort: "emailConfirmation-dp",
    url: "/emailConfirmation/delete-preview",
    method: "get",
    qs: [
      {
        name: "uniqueIds",
        type: "slice",
        primitive: "string",
      },
    ],
    description:
      'Reports what deleting the given "emailConfirmation" uniqueIds would affect, without deleting anything.',
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
 * The base class definition for emailConfirmationAwareDeletePreviewActionRes
 **/
export class EmailConfirmationAwareDeletePreviewActionRes {
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
   * @type {EmailConfirmationAwareDeletePreviewActionRes.Affected}
   **/
  #affected: MArray<
    InstanceType<typeof EmailConfirmationAwareDeletePreviewActionRes.Affected>
  > = MArray.of([]);
  /**
   *
   * @returns {EmailConfirmationAwareDeletePreviewActionRes.Affected}
   **/
  get affected() {
    return this.#affected;
  }
  /**
   *
   * @type {EmailConfirmationAwareDeletePreviewActionRes.Affected}
   **/
  set affected(
    value:
      | MArray<
          InstanceType<
            typeof EmailConfirmationAwareDeletePreviewActionRes.Affected
          >
        >
      | InstanceType<
          typeof EmailConfirmationAwareDeletePreviewActionRes.Affected
        >[],
  ) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (
        value.length > 0 &&
        value[0] instanceof
          EmailConfirmationAwareDeletePreviewActionRes.Affected
      ) {
        this.#affected = MArray.of(value);
      } else {
        this.#affected = MArray.of(
          value.map(
            (item) =>
              new EmailConfirmationAwareDeletePreviewActionRes.Affected(item),
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
            typeof EmailConfirmationAwareDeletePreviewActionRes.Affected
          >
        >
      | InstanceType<
          typeof EmailConfirmationAwareDeletePreviewActionRes.Affected
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
     * Creates an instance of EmailConfirmationAwareDeletePreviewActionRes.Affected, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(
      possibleDtoObject: EmailConfirmationAwareDeletePreviewActionResType.AffectedType,
    ) {
      return new EmailConfirmationAwareDeletePreviewActionRes.Affected(
        possibleDtoObject,
      );
    }
    /**
     * Creates an instance of EmailConfirmationAwareDeletePreviewActionRes.Affected, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<EmailConfirmationAwareDeletePreviewActionResType.AffectedType>,
    ) {
      return new EmailConfirmationAwareDeletePreviewActionRes.Affected(
        partialDtoObject,
      );
    }
    copyWith(
      partial: PartialDeep<EmailConfirmationAwareDeletePreviewActionResType.AffectedType>,
    ): InstanceType<
      typeof EmailConfirmationAwareDeletePreviewActionRes.Affected
    > {
      return new EmailConfirmationAwareDeletePreviewActionRes.Affected({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<
      typeof EmailConfirmationAwareDeletePreviewActionRes.Affected
    > {
      return new EmailConfirmationAwareDeletePreviewActionRes.Affected(
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
    const d = data as Partial<EmailConfirmationAwareDeletePreviewActionRes>;
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
          EmailConfirmationAwareDeletePreviewActionRes.Affected.Fields,
        );
      },
    };
  }
  /**
   * Creates an instance of EmailConfirmationAwareDeletePreviewActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(
    possibleDtoObject: EmailConfirmationAwareDeletePreviewActionResType,
  ) {
    return new EmailConfirmationAwareDeletePreviewActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of EmailConfirmationAwareDeletePreviewActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<EmailConfirmationAwareDeletePreviewActionResType>,
  ) {
    return new EmailConfirmationAwareDeletePreviewActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<EmailConfirmationAwareDeletePreviewActionResType>,
  ): InstanceType<typeof EmailConfirmationAwareDeletePreviewActionRes> {
    return new EmailConfirmationAwareDeletePreviewActionRes({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof EmailConfirmationAwareDeletePreviewActionRes> {
    return new EmailConfirmationAwareDeletePreviewActionRes(this.toJSON());
  }
}
export abstract class EmailConfirmationAwareDeletePreviewActionResFactory {
  abstract create(data: unknown): EmailConfirmationAwareDeletePreviewActionRes;
}
/**
 * The base type definition for emailConfirmationAwareDeletePreviewActionRes
 **/
export type EmailConfirmationAwareDeletePreviewActionResType = {
  /**
   *
   * @type {string}
   **/
  message: string;
  /**
   *
   * @type {EmailConfirmationAwareDeletePreviewActionResType.AffectedType[]}
   **/
  affected: EmailConfirmationAwareDeletePreviewActionResType.AffectedType[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace EmailConfirmationAwareDeletePreviewActionResType {
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
 * EmailConfirmationAwareDeletePreviewActionQueryParams class
 * Auto-generated from EmiAction
 */
export class EmailConfirmationAwareDeletePreviewActionQueryParams extends URLSearchParamsX {
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
