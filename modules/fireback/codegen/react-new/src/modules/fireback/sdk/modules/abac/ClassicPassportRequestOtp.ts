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
import { type UseMutationOptions, useMutation } from "react-query";
import { useFetchxContext } from "../../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action ClassicPassportRequestOtp
 */
export type ClassicPassportRequestOtpActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type ClassicPassportRequestOtpActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  ClassicPassportRequestOtpActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => ClassicPassportRequestOtpActionRes;
  }>;
export const useClassicPassportRequestOtpAction = (
  options?: ClassicPassportRequestOtpActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: ClassicPassportRequestOtpActionReq) => {
    setCompleteState(false);
    return ClassicPassportRequestOtpAction.Fetch(
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
 * ClassicPassportRequestOtpAction
 */
export class ClassicPassportRequestOtpAction {
  //
  static URL = "/workspace/passport/request-otp";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(ClassicPassportRequestOtpAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<ClassicPassportRequestOtpActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<ClassicPassportRequestOtpActionRes>,
      ClassicPassportRequestOtpActionReq,
      unknown
    >(
      overrideUrl ?? ClassicPassportRequestOtpAction.NewUrl(qs),
      {
        method: ClassicPassportRequestOtpAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<ClassicPassportRequestOtpActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?:
        | ((item: unknown) => ClassicPassportRequestOtpActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new ClassicPassportRequestOtpActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new ClassicPassportRequestOtpActionRes(item));
    const res = await ClassicPassportRequestOtpAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<ClassicPassportRequestOtpActionRes>();
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
    name: "ClassicPassportRequestOtp",
    cliName: "otp-request",
    url: "/workspace/passport/request-otp",
    method: "post",
    description:
      "Triggers an otp request, and will send an sms or email to the passport. This endpoint is not used for login, but rather makes a request at initial step. Later you can call classicPassportOtp to get in.",
    in: {
      fields: [
        {
          name: "value",
          description:
            "Passport value (email, phone number) which would be receiving the otp code.",
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
          name: "suspendUntil",
          type: "int64",
        },
        {
          name: "validUntil",
          type: "int64",
        },
        {
          name: "blockedUntil",
          type: "int64",
        },
        {
          name: "secondsToUnblock",
          description: "The amount of time left to unblock for next request",
          type: "int64",
        },
      ],
    },
  };
}
/**
 * The base class definition for classicPassportRequestOtpActionReq
 **/
export class ClassicPassportRequestOtpActionReq {
  /**
   * Passport value (email, phone number) which would be receiving the otp code.
   * @type {string}
   **/
  #value: string = "";
  /**
   * Passport value (email, phone number) which would be receiving the otp code.
   * @returns {string}
   **/
  get value() {
    return this.#value;
  }
  /**
   * Passport value (email, phone number) which would be receiving the otp code.
   * @type {string}
   **/
  set value(value: string) {
    this.#value = String(value);
  }
  setValue(value: string) {
    this.value = value;
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
    const d = data as Partial<ClassicPassportRequestOtpActionReq>;
    if (d.value !== undefined) {
      this.value = d.value;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      value: this.#value,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      value: "value",
    };
  }
  /**
   * Creates an instance of ClassicPassportRequestOtpActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ClassicPassportRequestOtpActionReqType) {
    return new ClassicPassportRequestOtpActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of ClassicPassportRequestOtpActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<ClassicPassportRequestOtpActionReqType>,
  ) {
    return new ClassicPassportRequestOtpActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ClassicPassportRequestOtpActionReqType>,
  ): InstanceType<typeof ClassicPassportRequestOtpActionReq> {
    return new ClassicPassportRequestOtpActionReq({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof ClassicPassportRequestOtpActionReq> {
    return new ClassicPassportRequestOtpActionReq(this.toJSON());
  }
}
export abstract class ClassicPassportRequestOtpActionReqFactory {
  abstract create(data: unknown): ClassicPassportRequestOtpActionReq;
}
/**
 * The base type definition for classicPassportRequestOtpActionReq
 **/
export type ClassicPassportRequestOtpActionReqType = {
  /**
   * Passport value (email, phone number) which would be receiving the otp code.
   * @type {string}
   **/
  value: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ClassicPassportRequestOtpActionReqType {}
/**
 * The base class definition for classicPassportRequestOtpActionRes
 **/
export class ClassicPassportRequestOtpActionRes {
  /**
   *
   * @type {number}
   **/
  #suspendUntil: number = 0;
  /**
   *
   * @returns {number}
   **/
  get suspendUntil() {
    return this.#suspendUntil;
  }
  /**
   *
   * @type {number}
   **/
  set suspendUntil(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#suspendUntil = parsedValue;
    }
  }
  setSuspendUntil(value: number) {
    this.suspendUntil = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #validUntil: number = 0;
  /**
   *
   * @returns {number}
   **/
  get validUntil() {
    return this.#validUntil;
  }
  /**
   *
   * @type {number}
   **/
  set validUntil(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#validUntil = parsedValue;
    }
  }
  setValidUntil(value: number) {
    this.validUntil = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #blockedUntil: number = 0;
  /**
   *
   * @returns {number}
   **/
  get blockedUntil() {
    return this.#blockedUntil;
  }
  /**
   *
   * @type {number}
   **/
  set blockedUntil(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#blockedUntil = parsedValue;
    }
  }
  setBlockedUntil(value: number) {
    this.blockedUntil = value;
    return this;
  }
  /**
   * The amount of time left to unblock for next request
   * @type {number}
   **/
  #secondsToUnblock: number = 0;
  /**
   * The amount of time left to unblock for next request
   * @returns {number}
   **/
  get secondsToUnblock() {
    return this.#secondsToUnblock;
  }
  /**
   * The amount of time left to unblock for next request
   * @type {number}
   **/
  set secondsToUnblock(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#secondsToUnblock = parsedValue;
    }
  }
  setSecondsToUnblock(value: number) {
    this.secondsToUnblock = value;
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
    const d = data as Partial<ClassicPassportRequestOtpActionRes>;
    if (d.suspendUntil !== undefined) {
      this.suspendUntil = d.suspendUntil;
    }
    if (d.validUntil !== undefined) {
      this.validUntil = d.validUntil;
    }
    if (d.blockedUntil !== undefined) {
      this.blockedUntil = d.blockedUntil;
    }
    if (d.secondsToUnblock !== undefined) {
      this.secondsToUnblock = d.secondsToUnblock;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      suspendUntil: this.#suspendUntil,
      validUntil: this.#validUntil,
      blockedUntil: this.#blockedUntil,
      secondsToUnblock: this.#secondsToUnblock,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      suspendUntil: "suspendUntil",
      validUntil: "validUntil",
      blockedUntil: "blockedUntil",
      secondsToUnblock: "secondsToUnblock",
    };
  }
  /**
   * Creates an instance of ClassicPassportRequestOtpActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ClassicPassportRequestOtpActionResType) {
    return new ClassicPassportRequestOtpActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of ClassicPassportRequestOtpActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<ClassicPassportRequestOtpActionResType>,
  ) {
    return new ClassicPassportRequestOtpActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ClassicPassportRequestOtpActionResType>,
  ): InstanceType<typeof ClassicPassportRequestOtpActionRes> {
    return new ClassicPassportRequestOtpActionRes({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof ClassicPassportRequestOtpActionRes> {
    return new ClassicPassportRequestOtpActionRes(this.toJSON());
  }
}
export abstract class ClassicPassportRequestOtpActionResFactory {
  abstract create(data: unknown): ClassicPassportRequestOtpActionRes;
}
/**
 * The base type definition for classicPassportRequestOtpActionRes
 **/
export type ClassicPassportRequestOtpActionResType = {
  /**
   *
   * @type {number}
   **/
  suspendUntil: number;
  /**
   *
   * @type {number}
   **/
  validUntil: number;
  /**
   *
   * @type {number}
   **/
  blockedUntil: number;
  /**
   * The amount of time left to unblock for next request
   * @type {number}
   **/
  secondsToUnblock: number;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ClassicPassportRequestOtpActionResType {}
