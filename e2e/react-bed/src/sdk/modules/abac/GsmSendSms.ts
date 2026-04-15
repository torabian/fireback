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
 * Action to communicate with the action GsmSendSms
 */
export type GsmSendSmsActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type GsmSendSmsActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  GsmSendSmsActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => GsmSendSmsActionRes;
  }>;
export const useGsmSendSmsAction = (
  options?: GsmSendSmsActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: GsmSendSmsActionReq) => {
    setCompleteState(false);
    return GsmSendSmsAction.Fetch(
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
 * GsmSendSmsAction
 */
export class GsmSendSmsAction {
  //
  static URL = "/gsm/send/sms";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(GsmSendSmsAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<GsmSendSmsActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GsmSendSmsActionRes, GsmSendSmsActionReq, unknown>(
      overrideUrl ?? GsmSendSmsAction.NewUrl(qs),
      {
        method: GsmSendSmsAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<GsmSendSmsActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => GsmSendSmsActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new GsmSendSmsActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new GsmSendSmsActionRes(item));
    const res = await GsmSendSmsAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "GsmSendSms",
    cliName: "sms",
    url: "/gsm/send/sms",
    method: "post",
    description:
      "Send a text message using default root notification configuration",
    in: {
      fields: [
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
 * The base class definition for gsmSendSmsActionReq
 **/
export class GsmSendSmsActionReq {
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
    const d = data as Partial<GsmSendSmsActionReq>;
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
      toNumber: this.#toNumber,
      body: this.#body,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      toNumber: "toNumber",
      body: "body",
    };
  }
  /**
   * Creates an instance of GsmSendSmsActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: GsmSendSmsActionReqType) {
    return new GsmSendSmsActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of GsmSendSmsActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<GsmSendSmsActionReqType>) {
    return new GsmSendSmsActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<GsmSendSmsActionReqType>,
  ): InstanceType<typeof GsmSendSmsActionReq> {
    return new GsmSendSmsActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof GsmSendSmsActionReq> {
    return new GsmSendSmsActionReq(this.toJSON());
  }
}
export abstract class GsmSendSmsActionReqFactory {
  abstract create(data: unknown): GsmSendSmsActionReq;
}
/**
 * The base type definition for gsmSendSmsActionReq
 **/
export type GsmSendSmsActionReqType = {
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
export namespace GsmSendSmsActionReqType {}
/**
 * The base class definition for gsmSendSmsActionRes
 **/
export class GsmSendSmsActionRes {
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
    const d = data as Partial<GsmSendSmsActionRes>;
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
   * Creates an instance of GsmSendSmsActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: GsmSendSmsActionResType) {
    return new GsmSendSmsActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of GsmSendSmsActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<GsmSendSmsActionResType>) {
    return new GsmSendSmsActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<GsmSendSmsActionResType>,
  ): InstanceType<typeof GsmSendSmsActionRes> {
    return new GsmSendSmsActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof GsmSendSmsActionRes> {
    return new GsmSendSmsActionRes(this.toJSON());
  }
}
export abstract class GsmSendSmsActionResFactory {
  abstract create(data: unknown): GsmSendSmsActionRes;
}
/**
 * The base type definition for gsmSendSmsActionRes
 **/
export type GsmSendSmsActionResType = {
  /**
   *
   * @type {string}
   **/
  queueId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GsmSendSmsActionResType {}
