import { EmailProviderEntity } from "./EmailProviderEntity";
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
import { withPrefix } from "../../sdk/common/withPrefix";
/**
 * Action to communicate with the action SendEmailWithProvider
 */
export type SendEmailWithProviderActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type SendEmailWithProviderActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  SendEmailWithProviderActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => SendEmailWithProviderActionRes;
  }>;
export const useSendEmailWithProviderAction = (
  options?: SendEmailWithProviderActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: SendEmailWithProviderActionReq) => {
    setCompleteState(false);
    return SendEmailWithProviderAction.Fetch(
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
 * SendEmailWithProviderAction
 */
export class SendEmailWithProviderAction {
  //
  static URL = "/emailProvider/send";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(SendEmailWithProviderAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<SendEmailWithProviderActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      SendEmailWithProviderActionRes,
      SendEmailWithProviderActionReq,
      unknown
    >(
      overrideUrl ?? SendEmailWithProviderAction.NewUrl(qs),
      {
        method: SendEmailWithProviderAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<SendEmailWithProviderActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?:
        | ((item: unknown) => SendEmailWithProviderActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new SendEmailWithProviderActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new SendEmailWithProviderActionRes(item));
    const res = await SendEmailWithProviderAction.Fetch$(
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
    name: "SendEmailWithProvider",
    cliName: "emailp",
    url: "/emailProvider/send",
    method: "post",
    description: "Send a text message using an specific gsm provider",
    in: {
      fields: [
        {
          name: "emailProvider",
          type: "one",
          target: "EmailProviderEntity",
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
 * The base class definition for sendEmailWithProviderActionReq
 **/
export class SendEmailWithProviderActionReq {
  /**
   *
   * @type {EmailProviderEntity}
   **/
  #emailProvider!: EmailProviderEntity;
  /**
   *
   * @returns {EmailProviderEntity}
   **/
  get emailProvider() {
    return this.#emailProvider;
  }
  /**
   *
   * @type {EmailProviderEntity}
   **/
  set emailProvider(value: EmailProviderEntity) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof EmailProviderEntity) {
      this.#emailProvider = value;
    } else {
      this.#emailProvider = new EmailProviderEntity(value);
    }
  }
  setEmailProvider(value: EmailProviderEntity) {
    this.emailProvider = value;
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
      this.#lateInitFields();
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
    const d = data as Partial<SendEmailWithProviderActionReq>;
    if (d.emailProvider !== undefined) {
      this.emailProvider = d.emailProvider;
    }
    if (d.toAddress !== undefined) {
      this.toAddress = d.toAddress;
    }
    if (d.body !== undefined) {
      this.body = d.body;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<SendEmailWithProviderActionReq>;
    if (!(d.emailProvider instanceof EmailProviderEntity)) {
      this.emailProvider = new EmailProviderEntity(d.emailProvider || {});
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      emailProvider: this.#emailProvider,
      toAddress: this.#toAddress,
      body: this.#body,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      emailProvider$: "emailProvider",
      get emailProvider() {
        return withPrefix("emailProvider", EmailProviderEntity.Fields);
      },
      toAddress: "toAddress",
      body: "body",
    };
  }
  /**
   * Creates an instance of SendEmailWithProviderActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: SendEmailWithProviderActionReqType) {
    return new SendEmailWithProviderActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of SendEmailWithProviderActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<SendEmailWithProviderActionReqType>,
  ) {
    return new SendEmailWithProviderActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<SendEmailWithProviderActionReqType>,
  ): InstanceType<typeof SendEmailWithProviderActionReq> {
    return new SendEmailWithProviderActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof SendEmailWithProviderActionReq> {
    return new SendEmailWithProviderActionReq(this.toJSON());
  }
}
export abstract class SendEmailWithProviderActionReqFactory {
  abstract create(data: unknown): SendEmailWithProviderActionReq;
}
/**
 * The base type definition for sendEmailWithProviderActionReq
 **/
export type SendEmailWithProviderActionReqType = {
  /**
   *
   * @type {EmailProviderEntity}
   **/
  emailProvider: EmailProviderEntity;
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
export namespace SendEmailWithProviderActionReqType {}
/**
 * The base class definition for sendEmailWithProviderActionRes
 **/
export class SendEmailWithProviderActionRes {
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
    const d = data as Partial<SendEmailWithProviderActionRes>;
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
   * Creates an instance of SendEmailWithProviderActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: SendEmailWithProviderActionResType) {
    return new SendEmailWithProviderActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of SendEmailWithProviderActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<SendEmailWithProviderActionResType>,
  ) {
    return new SendEmailWithProviderActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<SendEmailWithProviderActionResType>,
  ): InstanceType<typeof SendEmailWithProviderActionRes> {
    return new SendEmailWithProviderActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof SendEmailWithProviderActionRes> {
    return new SendEmailWithProviderActionRes(this.toJSON());
  }
}
export abstract class SendEmailWithProviderActionResFactory {
  abstract create(data: unknown): SendEmailWithProviderActionRes;
}
/**
 * The base type definition for sendEmailWithProviderActionRes
 **/
export type SendEmailWithProviderActionResType = {
  /**
   *
   * @type {string}
   **/
  queueId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SendEmailWithProviderActionResType {}
