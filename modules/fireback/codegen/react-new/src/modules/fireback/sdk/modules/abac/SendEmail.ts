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
 * Action to communicate with the action SendEmail
 */
export type SendEmailActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type SendEmailActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  SendEmailActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => SendEmailActionRes;
  }>;
export const useSendEmailAction = (
  options?: SendEmailActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: SendEmailActionReq) => {
    setCompleteState(false);
    return SendEmailAction.Fetch(
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
 * SendEmailAction
 */
export class SendEmailAction {
  //
  static URL = "/email/send";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(SendEmailAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<SendEmailActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<SendEmailActionRes, SendEmailActionReq, unknown>(
      overrideUrl ?? SendEmailAction.NewUrl(qs),
      {
        method: SendEmailAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<SendEmailActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => SendEmailActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new SendEmailActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new SendEmailActionRes(item));
    const res = await SendEmailAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "SendEmail",
    cliName: "email",
    url: "/email/send",
    method: "post",
    description: "Send a email using default root notification configuration",
    in: {
      fields: [
        {
          name: "providerId",
          description:
            "Sending a test email requires to be sent through an specific email provider.",
          type: "string",
        },
        {
          name: "toAddress",
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
 * The base class definition for sendEmailActionReq
 **/
export class SendEmailActionReq {
  /**
   * Sending a test email requires to be sent through an specific email provider.
   * @type {string}
   **/
  #providerId: string = "";
  /**
   * Sending a test email requires to be sent through an specific email provider.
   * @returns {string}
   **/
  get providerId() {
    return this.#providerId;
  }
  /**
   * Sending a test email requires to be sent through an specific email provider.
   * @type {string}
   **/
  set providerId(value: string) {
    this.#providerId = String(value);
  }
  setProviderId(value: string) {
    this.providerId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #toAddress: string = "";
  /**
   *
   * @returns {string}
   **/
  get toAddress() {
    return this.#toAddress;
  }
  /**
   *
   * @type {string}
   **/
  set toAddress(value: string) {
    this.#toAddress = String(value);
  }
  setToAddress(value: string) {
    this.toAddress = value;
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
    const d = data as Partial<SendEmailActionReq>;
    if (d.providerId !== undefined) {
      this.providerId = d.providerId;
    }
    if (d.toAddress !== undefined) {
      this.toAddress = d.toAddress;
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
      providerId: this.#providerId,
      toAddress: this.#toAddress,
      body: this.#body,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      providerId: "providerId",
      toAddress: "toAddress",
      body: "body",
    };
  }
  /**
   * Creates an instance of SendEmailActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: SendEmailActionReqType) {
    return new SendEmailActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of SendEmailActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<SendEmailActionReqType>) {
    return new SendEmailActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<SendEmailActionReqType>,
  ): InstanceType<typeof SendEmailActionReq> {
    return new SendEmailActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof SendEmailActionReq> {
    return new SendEmailActionReq(this.toJSON());
  }
}
export abstract class SendEmailActionReqFactory {
  abstract create(data: unknown): SendEmailActionReq;
}
/**
 * The base type definition for sendEmailActionReq
 **/
export type SendEmailActionReqType = {
  /**
   * Sending a test email requires to be sent through an specific email provider.
   * @type {string}
   **/
  providerId: string;
  /**
   *
   * @type {string}
   **/
  toAddress: string;
  /**
   *
   * @type {string}
   **/
  body: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SendEmailActionReqType {}
/**
 * The base class definition for sendEmailActionRes
 **/
export class SendEmailActionRes {
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
    const d = data as Partial<SendEmailActionRes>;
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
   * Creates an instance of SendEmailActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: SendEmailActionResType) {
    return new SendEmailActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of SendEmailActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<SendEmailActionResType>) {
    return new SendEmailActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<SendEmailActionResType>,
  ): InstanceType<typeof SendEmailActionRes> {
    return new SendEmailActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof SendEmailActionRes> {
    return new SendEmailActionRes(this.toJSON());
  }
}
export abstract class SendEmailActionResFactory {
  abstract create(data: unknown): SendEmailActionRes;
}
/**
 * The base type definition for sendEmailActionRes
 **/
export type SendEmailActionResType = {
  /**
   *
   * @type {string}
   **/
  queueId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SendEmailActionResType {}
