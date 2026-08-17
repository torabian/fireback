import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { MOne } from "@fireback/js-remote-ctx/common/operators";
import { UserSessionDto } from "./UserSessionDto";
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
 * Action to communicate with the action CompletePassportPasswordReset
 */
export type CompletePassportPasswordResetActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type CompletePassportPasswordResetActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  CompletePassportPasswordResetActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => CompletePassportPasswordResetActionRes;
  }>;
export const useCompletePassportPasswordResetAction = (
  options?: CompletePassportPasswordResetActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: CompletePassportPasswordResetActionReq) => {
    setCompleteState(false);
    return CompletePassportPasswordResetAction.Fetch(
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
 * CompletePassportPasswordResetAction
 */
export class CompletePassportPasswordResetAction {
  //
  static URL = "/passport/reset-password";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(CompletePassportPasswordResetAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<CompletePassportPasswordResetActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<CompletePassportPasswordResetActionRes>,
      CompletePassportPasswordResetActionReq,
      unknown
    >(
      overrideUrl ?? CompletePassportPasswordResetAction.NewUrl(qs),
      {
        method: CompletePassportPasswordResetAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<CompletePassportPasswordResetActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?:
        | ((item: unknown) => CompletePassportPasswordResetActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new CompletePassportPasswordResetActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new CompletePassportPasswordResetActionRes(item));
    const res = await CompletePassportPasswordResetAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<CompletePassportPasswordResetActionRes>();
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
    name: "CompletePassportPasswordReset",
    cliName: "reset-password",
    url: "/passport/reset-password",
    method: "post",
    description:
      "Public. Finishes the self-service reset flow started by ClassicPassportRequestOtp or SendPassportResetEmail - given the passport value and the OTP code that was emailed/texted to it, sets a new password and signs the user in, the same way a successful ClassicPassportOtp login would.",
    in: {
      fields: [
        {
          name: "value",
          description:
            "Passport value (email, phone number) the OTP was sent to.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "otp",
          description: "The OTP code emailed/texted to the passport value.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "password",
          description: "New password meeting the security requirements.",
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
          name: "reset",
          type: "bool",
        },
        {
          name: "session",
          type: "one?",
          target: "UserSessionDto",
        },
      ],
    },
  };
}
/**
 * The base class definition for completePassportPasswordResetActionReq
 **/
export class CompletePassportPasswordResetActionReq {
  /**
   * Passport value (email, phone number) the OTP was sent to.
   * @type {string}
   **/
  #value: string = "";
  /**
   * Passport value (email, phone number) the OTP was sent to.
   * @returns {string}
   **/
  get value() {
    return this.#value;
  }
  /**
   * Passport value (email, phone number) the OTP was sent to.
   * @type {string}
   **/
  set value(value: string) {
    this.#value = String(value);
  }
  setValue(value: string) {
    this.value = value;
    return this;
  }
  /**
   * The OTP code emailed/texted to the passport value.
   * @type {string}
   **/
  #otp: string = "";
  /**
   * The OTP code emailed/texted to the passport value.
   * @returns {string}
   **/
  get otp() {
    return this.#otp;
  }
  /**
   * The OTP code emailed/texted to the passport value.
   * @type {string}
   **/
  set otp(value: string) {
    this.#otp = String(value);
  }
  setOtp(value: string) {
    this.otp = value;
    return this;
  }
  /**
   * New password meeting the security requirements.
   * @type {string}
   **/
  #password: string = "";
  /**
   * New password meeting the security requirements.
   * @returns {string}
   **/
  get password() {
    return this.#password;
  }
  /**
   * New password meeting the security requirements.
   * @type {string}
   **/
  set password(value: string) {
    this.#password = String(value);
  }
  setPassword(value: string) {
    this.password = value;
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
    const d = data as Partial<CompletePassportPasswordResetActionReq>;
    if (d.value !== undefined) {
      this.value = d.value;
    }
    if (d.otp !== undefined) {
      this.otp = d.otp;
    }
    if (d.password !== undefined) {
      this.password = d.password;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      value: this.#value,
      otp: this.#otp,
      password: this.#password,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      value: "value",
      otp: "otp",
      password: "password",
    };
  }
  /**
   * Creates an instance of CompletePassportPasswordResetActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CompletePassportPasswordResetActionReqType) {
    return new CompletePassportPasswordResetActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of CompletePassportPasswordResetActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<CompletePassportPasswordResetActionReqType>,
  ) {
    return new CompletePassportPasswordResetActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CompletePassportPasswordResetActionReqType>,
  ): InstanceType<typeof CompletePassportPasswordResetActionReq> {
    return new CompletePassportPasswordResetActionReq({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof CompletePassportPasswordResetActionReq> {
    return new CompletePassportPasswordResetActionReq(this.toJSON());
  }
}
export abstract class CompletePassportPasswordResetActionReqFactory {
  abstract create(data: unknown): CompletePassportPasswordResetActionReq;
}
/**
 * The base type definition for completePassportPasswordResetActionReq
 **/
export type CompletePassportPasswordResetActionReqType = {
  /**
   * Passport value (email, phone number) the OTP was sent to.
   * @type {string}
   **/
  value: string;
  /**
   * The OTP code emailed/texted to the passport value.
   * @type {string}
   **/
  otp: string;
  /**
   * New password meeting the security requirements.
   * @type {string}
   **/
  password: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CompletePassportPasswordResetActionReqType {}
/**
 * The base class definition for completePassportPasswordResetActionRes
 **/
export class CompletePassportPasswordResetActionRes {
  /**
   *
   * @type {boolean}
   **/
  #reset!: boolean;
  /**
   *
   * @returns {boolean}
   **/
  get reset() {
    return this.#reset;
  }
  /**
   *
   * @type {boolean}
   **/
  set reset(value: boolean) {
    this.#reset = Boolean(value);
  }
  setReset(value: boolean) {
    this.reset = value;
    return this;
  }
  /**
   *
   * @type {UserSessionDto}
   **/
  #session?: MOne<UserSessionDto> | null = undefined;
  /**
   *
   * @returns {UserSessionDto}
   **/
  get session() {
    return this.#session;
  }
  /**
   *
   * @type {UserSessionDto}
   **/
  set session(
    value:
      | MOne<UserSessionDto>
      | InstanceType<typeof UserSessionDto>
      | null
      | undefined,
  ) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof MOne) {
      this.#session = value;
    } else if (value instanceof UserSessionDto) {
      this.#session = MOne.of(value);
    } else {
      this.#session = MOne.of(new UserSessionDto(value));
    }
  }
  setSession(
    value:
      | MOne<UserSessionDto>
      | InstanceType<typeof UserSessionDto>
      | null
      | undefined,
  ) {
    this.session = value;
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
    const d = data as Partial<CompletePassportPasswordResetActionRes>;
    if (d.reset !== undefined) {
      this.reset = d.reset;
    }
    if (d.session !== undefined) {
      this.session = d.session;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      reset: this.#reset,
      session: this.#session,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      reset: "reset",
      session: "session",
    };
  }
  /**
   * Creates an instance of CompletePassportPasswordResetActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CompletePassportPasswordResetActionResType) {
    return new CompletePassportPasswordResetActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of CompletePassportPasswordResetActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<CompletePassportPasswordResetActionResType>,
  ) {
    return new CompletePassportPasswordResetActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CompletePassportPasswordResetActionResType>,
  ): InstanceType<typeof CompletePassportPasswordResetActionRes> {
    return new CompletePassportPasswordResetActionRes({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof CompletePassportPasswordResetActionRes> {
    return new CompletePassportPasswordResetActionRes(this.toJSON());
  }
}
export abstract class CompletePassportPasswordResetActionResFactory {
  abstract create(data: unknown): CompletePassportPasswordResetActionRes;
}
/**
 * The base type definition for completePassportPasswordResetActionRes
 **/
export type CompletePassportPasswordResetActionResType = {
  /**
   *
   * @type {boolean}
   **/
  reset: boolean;
  /**
   *
   * @type {UserSessionDto}
   **/
  session?: UserSessionDto;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CompletePassportPasswordResetActionResType {}
