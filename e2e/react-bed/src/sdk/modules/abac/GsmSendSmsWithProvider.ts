import { buildUrl } from "../../sdk/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
  type TypedRequestInit,
  type TypedResponse,
} from "../../sdk/common/fetchx";
import { type UseMutationOptions, useMutation } from "react-query";
import { useFetchxContext } from "../../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action GsmSendSmsWithProvider
 */
export type GsmSendSmsWithProviderActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type GsmSendSmsWithProviderActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  GsmSendSmsWithProviderActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => GsmSendSmsWithProviderActionRes;
  }>;
export const useGsmSendSmsWithProviderAction = (
  options?: GsmSendSmsWithProviderActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: GsmSendSmsWithProviderActionReq) => {
    setCompleteState(false);
    return GsmSendSmsWithProviderAction.Fetch(
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
 * GsmSendSmsWithProviderAction
 */
export class GsmSendSmsWithProviderAction {
  //
  static URL = "/gsmProvider/send/sms";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(GsmSendSmsWithProviderAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<GsmSendSmsWithProviderActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GsmSendSmsWithProviderActionRes,
      GsmSendSmsWithProviderActionReq,
      unknown
    >(
      overrideUrl ?? GsmSendSmsWithProviderAction.NewUrl(qs),
      {
        method: GsmSendSmsWithProviderAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<GsmSendSmsWithProviderActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?:
        | ((item: unknown) => GsmSendSmsWithProviderActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new GsmSendSmsWithProviderActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new GsmSendSmsWithProviderActionRes(item));
    const res = await GsmSendSmsWithProviderAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "GsmSendSmsWithProvider",
    cliName: "smsp",
    url: "/gsmProvider/send/sms",
    method: "post",
    description: "Send a text message using an specific gsm provider",
    in: {
      fields: [
        {
          name: "gsmProviderId",
          type: "string",
        },
        {
          name: "toNumber",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "body",
          type: "string",
          tags: {
            validate: "required",
          },
        },
      ],
    },
    out: {
      fields: [
        {
          name: "queueId",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for gsmSendSmsWithProviderActionReq
 **/
export class GsmSendSmsWithProviderActionReq {
  /**
   *
   * @type {string}
   **/
  #gsmProviderId: string = "";
  /**
   *
   * @returns {string}
   **/
  get gsmProviderId() {
    return this.#gsmProviderId;
  }
  /**
   *
   * @type {string}
   **/
  set gsmProviderId(value: string) {
    this.#gsmProviderId = String(value);
  }
  setGsmProviderId(value: string) {
    this.gsmProviderId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #toNumber: string = "";
  /**
   *
   * @returns {string}
   **/
  get toNumber() {
    return this.#toNumber;
  }
  /**
   *
   * @type {string}
   **/
  set toNumber(value: string) {
    this.#toNumber = String(value);
  }
  setToNumber(value: string) {
    this.toNumber = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #body: string = "";
  /**
   *
   * @returns {string}
   **/
  get body() {
    return this.#body;
  }
  /**
   *
   * @type {string}
   **/
  set body(value: string) {
    this.#body = String(value);
  }
  setBody(value: string) {
    this.body = value;
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
    const d = data as Partial<GsmSendSmsWithProviderActionReq>;
    if (d.gsmProviderId !== undefined) {
      this.gsmProviderId = d.gsmProviderId;
    }
    if (d.toNumber !== undefined) {
      this.toNumber = d.toNumber;
    }
    if (d.body !== undefined) {
      this.body = d.body;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      gsmProviderId: this.#gsmProviderId,
      toNumber: this.#toNumber,
      body: this.#body,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      gsmProviderId: "gsmProviderId",
      toNumber: "toNumber",
      body: "body",
    };
  }
  /**
   * Creates an instance of GsmSendSmsWithProviderActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: GsmSendSmsWithProviderActionReqType) {
    return new GsmSendSmsWithProviderActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of GsmSendSmsWithProviderActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<GsmSendSmsWithProviderActionReqType>,
  ) {
    return new GsmSendSmsWithProviderActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<GsmSendSmsWithProviderActionReqType>,
  ): InstanceType<typeof GsmSendSmsWithProviderActionReq> {
    return new GsmSendSmsWithProviderActionReq({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof GsmSendSmsWithProviderActionReq> {
    return new GsmSendSmsWithProviderActionReq(this.toJSON());
  }
}
export abstract class GsmSendSmsWithProviderActionReqFactory {
  abstract create(data: unknown): GsmSendSmsWithProviderActionReq;
}
/**
 * The base type definition for gsmSendSmsWithProviderActionReq
 **/
export type GsmSendSmsWithProviderActionReqType = {
  /**
   *
   * @type {string}
   **/
  gsmProviderId: string;
  /**
   *
   * @type {string}
   **/
  toNumber: string;
  /**
   *
   * @type {string}
   **/
  body: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GsmSendSmsWithProviderActionReqType {}
/**
 * The base class definition for gsmSendSmsWithProviderActionRes
 **/
export class GsmSendSmsWithProviderActionRes {
  /**
   *
   * @type {string}
   **/
  #queueId: string = "";
  /**
   *
   * @returns {string}
   **/
  get queueId() {
    return this.#queueId;
  }
  /**
   *
   * @type {string}
   **/
  set queueId(value: string) {
    this.#queueId = String(value);
  }
  setQueueId(value: string) {
    this.queueId = value;
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
    const d = data as Partial<GsmSendSmsWithProviderActionRes>;
    if (d.queueId !== undefined) {
      this.queueId = d.queueId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      queueId: this.#queueId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      queueId: "queueId",
    };
  }
  /**
   * Creates an instance of GsmSendSmsWithProviderActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: GsmSendSmsWithProviderActionResType) {
    return new GsmSendSmsWithProviderActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of GsmSendSmsWithProviderActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<GsmSendSmsWithProviderActionResType>,
  ) {
    return new GsmSendSmsWithProviderActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<GsmSendSmsWithProviderActionResType>,
  ): InstanceType<typeof GsmSendSmsWithProviderActionRes> {
    return new GsmSendSmsWithProviderActionRes({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof GsmSendSmsWithProviderActionRes> {
    return new GsmSendSmsWithProviderActionRes(this.toJSON());
  }
}
export abstract class GsmSendSmsWithProviderActionResFactory {
  abstract create(data: unknown): GsmSendSmsWithProviderActionRes;
}
/**
 * The base type definition for gsmSendSmsWithProviderActionRes
 **/
export type GsmSendSmsWithProviderActionResType = {
  /**
   *
   * @type {string}
   **/
  queueId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GsmSendSmsWithProviderActionResType {}
