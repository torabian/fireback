import { CapabilityInfoDto } from "./CapabilityInfoDto";
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
 * Action to communicate with the action CapabilitiesTree
 */
export type CapabilitiesTreeActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type CapabilitiesTreeActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<CapabilitiesTreeActionRes>,
    unknown[]
  >,
  "queryKey"
> &
  CapabilitiesTreeActionOptions &
  Partial<{
    creatorFn: (item: unknown) => CapabilitiesTreeActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext;
  };
export const useCapabilitiesTreeActionQuery = (
  options: CapabilitiesTreeActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return CapabilitiesTreeAction.Fetch(
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
    queryKey: [CapabilitiesTreeAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type CapabilitiesTreeActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  CapabilitiesTreeActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => CapabilitiesTreeActionRes;
  }>;
export const useCapabilitiesTreeAction = (
  options?: CapabilitiesTreeActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return CapabilitiesTreeAction.Fetch(
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
 * CapabilitiesTreeAction
 */
export class CapabilitiesTreeAction {
  //
  static URL = "/capabilitiesTree";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(CapabilitiesTreeAction.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<CapabilitiesTreeActionRes>, unknown, unknown>(
      overrideUrl ?? CapabilitiesTreeAction.NewUrl(qs),
      {
        method: CapabilitiesTreeAction.Method,
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
      creatorFn?: ((item: unknown) => CapabilitiesTreeActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new CapabilitiesTreeActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new CapabilitiesTreeActionRes(item));
    const res = await CapabilitiesTreeAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<CapabilitiesTreeActionRes>();
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
    name: "CapabilitiesTree",
    cliName: "treex",
    url: "/capabilitiesTree",
    method: "get",
    description:
      "dLists all of the capabilities in database as a array of string as root access",
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "capabilities",
          type: "collection",
          target: "CapabilityInfoDto",
        },
        {
          name: "nested",
          type: "collection",
          target: "CapabilityInfoDto",
        },
      ],
    },
  };
}
/**
 * The base class definition for capabilitiesTreeActionRes
 **/
export class CapabilitiesTreeActionRes {
  /**
   *
   * @type {CapabilityInfoDto[]}
   **/
  #capabilities: CapabilityInfoDto[] = [];
  /**
   *
   * @returns {CapabilityInfoDto[]}
   **/
  get capabilities() {
    return this.#capabilities;
  }
  /**
   *
   * @type {CapabilityInfoDto[]}
   **/
  set capabilities(value: CapabilityInfoDto[]) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof CapabilityInfoDto) {
      this.#capabilities = value;
    } else {
      this.#capabilities = value.map((item) => new CapabilityInfoDto(item));
    }
  }
  setCapabilities(value: CapabilityInfoDto[]) {
    this.capabilities = value;
    return this;
  }
  /**
   *
   * @type {CapabilityInfoDto[]}
   **/
  #nested: CapabilityInfoDto[] = [];
  /**
   *
   * @returns {CapabilityInfoDto[]}
   **/
  get nested() {
    return this.#nested;
  }
  /**
   *
   * @type {CapabilityInfoDto[]}
   **/
  set nested(value: CapabilityInfoDto[]) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof CapabilityInfoDto) {
      this.#nested = value;
    } else {
      this.#nested = value.map((item) => new CapabilityInfoDto(item));
    }
  }
  setNested(value: CapabilityInfoDto[]) {
    this.nested = value;
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
    const d = data as Partial<CapabilitiesTreeActionRes>;
    if (d.capabilities !== undefined) {
      this.capabilities = d.capabilities;
    }
    if (d.nested !== undefined) {
      this.nested = d.nested;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      capabilities: this.#capabilities,
      nested: this.#nested,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      capabilities$: "capabilities",
      get capabilities() {
        return withPrefix("capabilities[:i]", CapabilityInfoDto.Fields);
      },
      nested$: "nested",
      get nested() {
        return withPrefix("nested[:i]", CapabilityInfoDto.Fields);
      },
    };
  }
  /**
   * Creates an instance of CapabilitiesTreeActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CapabilitiesTreeActionResType) {
    return new CapabilitiesTreeActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of CapabilitiesTreeActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<CapabilitiesTreeActionResType>) {
    return new CapabilitiesTreeActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CapabilitiesTreeActionResType>,
  ): InstanceType<typeof CapabilitiesTreeActionRes> {
    return new CapabilitiesTreeActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CapabilitiesTreeActionRes> {
    return new CapabilitiesTreeActionRes(this.toJSON());
  }
}
export abstract class CapabilitiesTreeActionResFactory {
  abstract create(data: unknown): CapabilitiesTreeActionRes;
}
/**
 * The base type definition for capabilitiesTreeActionRes
 **/
export type CapabilitiesTreeActionResType = {
  /**
   *
   * @type {CapabilityInfoDto[]}
   **/
  capabilities: CapabilityInfoDto[];
  /**
   *
   * @type {CapabilityInfoDto[]}
   **/
  nested: CapabilityInfoDto[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CapabilitiesTreeActionResType {}
