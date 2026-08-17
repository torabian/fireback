import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { buildUrl } from "@fireback/js-remote-ctx/common/buildUrl";
import {
  fetchx,
  handleFetchResponse,
  type FetchxContext,
  type PartialDeep,
  type TypedRequestInit,
  type TypedResponse,
} from "@fireback/js-remote-ctx/common/fetchx";
import { type UseMutationOptions, useMutation } from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action SendPassportResetEmail
 */
export type SendPassportResetEmailActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type SendPassportResetEmailActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  SendPassportResetEmailActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => SendPassportResetEmailActionRes;
  }>;
export const useSendPassportResetEmailAction = (
  options?: SendPassportResetEmailActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: SendPassportResetEmailActionReq) => {
    setCompleteState(false);
    return SendPassportResetEmailAction.Fetch(
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
 * SendPassportResetEmailAction
 */
export class SendPassportResetEmailAction {
  //
  static URL = "/passport/send-reset-email";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(SendPassportResetEmailAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<SendPassportResetEmailActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<SendPassportResetEmailActionRes>,
      SendPassportResetEmailActionReq,
      unknown
    >(
      overrideUrl ?? SendPassportResetEmailAction.NewUrl(qs),
      {
        method: SendPassportResetEmailAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<SendPassportResetEmailActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?:
        | ((item: unknown) => SendPassportResetEmailActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new SendPassportResetEmailActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new SendPassportResetEmailActionRes(item));
    const res = await SendPassportResetEmailAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<SendPassportResetEmailActionRes>();
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
    name: "SendPassportResetEmail",
    cliName: "send-reset-email",
    url: "/passport/send-reset-email",
    method: "post",
    description:
      "Root-only. Sends the same OTP/recovery email or SMS ClassicPassportRequestOtp sends for self-service login, but triggered by an administrator on the user's behalf and addressed by passport uniqueId instead of by value - lets an admin hand a locked-out user a way to reset their own password rather than the admin setting it directly (see SetPassportPassword for that alternative).",
    in: {
      fields: [
        {
          name: "uniqueId",
          description: "The passport uniqueId to send the reset email/SMS to.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
      ],
    },
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "sent",
          type: "bool",
        },
      ],
    },
  };
}
/**
 * The base class definition for sendPassportResetEmailActionReq
 **/
export class SendPassportResetEmailActionReq {
  /**
   * The passport uniqueId to send the reset email/SMS to.
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * The passport uniqueId to send the reset email/SMS to.
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * The passport uniqueId to send the reset email/SMS to.
   * @type {string}
   **/
  set uniqueId(value: string) {
    this.#uniqueId = String(value);
  }
  setUniqueId(value: string) {
    this.uniqueId = value;
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
    const d = data as Partial<SendPassportResetEmailActionReq>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
    };
  }
  /**
   * Creates an instance of SendPassportResetEmailActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: SendPassportResetEmailActionReqType) {
    return new SendPassportResetEmailActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of SendPassportResetEmailActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<SendPassportResetEmailActionReqType>,
  ) {
    return new SendPassportResetEmailActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<SendPassportResetEmailActionReqType>,
  ): InstanceType<typeof SendPassportResetEmailActionReq> {
    return new SendPassportResetEmailActionReq({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof SendPassportResetEmailActionReq> {
    return new SendPassportResetEmailActionReq(this.toJSON());
  }
}
export abstract class SendPassportResetEmailActionReqFactory {
  abstract create(data: unknown): SendPassportResetEmailActionReq;
}
/**
 * The base type definition for sendPassportResetEmailActionReq
 **/
export type SendPassportResetEmailActionReqType = {
  /**
   * The passport uniqueId to send the reset email/SMS to.
   * @type {string}
   **/
  uniqueId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SendPassportResetEmailActionReqType {}
/**
 * The base class definition for sendPassportResetEmailActionRes
 **/
export class SendPassportResetEmailActionRes {
  /**
   *
   * @type {boolean}
   **/
  #sent!: boolean;
  /**
   *
   * @returns {boolean}
   **/
  get sent() {
    return this.#sent;
  }
  /**
   *
   * @type {boolean}
   **/
  set sent(value: boolean) {
    this.#sent = Boolean(value);
  }
  setSent(value: boolean) {
    this.sent = value;
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
    const d = data as Partial<SendPassportResetEmailActionRes>;
    if (d.sent !== undefined) {
      this.sent = d.sent;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      sent: this.#sent,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      sent: "sent",
    };
  }
  /**
   * Creates an instance of SendPassportResetEmailActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: SendPassportResetEmailActionResType) {
    return new SendPassportResetEmailActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of SendPassportResetEmailActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<SendPassportResetEmailActionResType>,
  ) {
    return new SendPassportResetEmailActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<SendPassportResetEmailActionResType>,
  ): InstanceType<typeof SendPassportResetEmailActionRes> {
    return new SendPassportResetEmailActionRes({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof SendPassportResetEmailActionRes> {
    return new SendPassportResetEmailActionRes(this.toJSON());
  }
}
export abstract class SendPassportResetEmailActionResFactory {
  abstract create(data: unknown): SendPassportResetEmailActionRes;
}
/**
 * The base type definition for sendPassportResetEmailActionRes
 **/
export type SendPassportResetEmailActionResType = {
  /**
   *
   * @type {boolean}
   **/
  sent: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SendPassportResetEmailActionResType {}
