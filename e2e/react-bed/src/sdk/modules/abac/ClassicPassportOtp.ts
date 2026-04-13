import { GResponse } from "../../sdk/envelopes/index";
import { UserSessionDto } from "./UserSessionDto";
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
 * Action to communicate with the action ClassicPassportOtp
 */
export type ClassicPassportOtpActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type ClassicPassportOtpActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  ClassicPassportOtpActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => ClassicPassportOtpActionRes;
  }>;
export const useClassicPassportOtpAction = (
  options?: ClassicPassportOtpActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: ClassicPassportOtpActionReq) => {
    setCompleteState(false);
    return ClassicPassportOtpAction.Fetch(
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
 * ClassicPassportOtpAction
 */
export class ClassicPassportOtpAction {
  //
  static URL = "/workspace/passport/otp";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(ClassicPassportOtpAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<ClassicPassportOtpActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<ClassicPassportOtpActionRes>,
      ClassicPassportOtpActionReq,
      unknown
    >(
      overrideUrl ?? ClassicPassportOtpAction.NewUrl(qs),
      {
        method: ClassicPassportOtpAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<ClassicPassportOtpActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => ClassicPassportOtpActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new ClassicPassportOtpActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new ClassicPassportOtpActionRes(item));
    const res = await ClassicPassportOtpAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<ClassicPassportOtpActionRes>();
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
    name: "ClassicPassportOtp",
    cliName: "otp",
    url: "/workspace/passport/otp",
    method: "post",
    description:
      "Authenticate the user publicly for classic methods using communication service, such as sms, call, or email. You need to call classicPassportRequestOtp beforehand to send a otp code, and then validate it with this API. Also checkClassicPassport action might already sent the otp, so make sure you don't send it twice.",
    in: {
      fields: [
        {
          name: "value",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "otp",
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
          name: "session",
          description:
            "Upon successful authentication, there will be a session dto generated, which is a ground information of authorized user and can be stored in front-end.",
          type: "one?",
          target: "UserSessionDto",
        },
        {
          name: "totpUrl",
          description:
            "If time based otp is available, we add it response to make it easier for ui.",
          type: "string",
        },
        {
          name: "sessionSecret",
          description:
            "The session secret will be used to call complete user registration api.",
          type: "string",
        },
        {
          name: "continueWithCreation",
          description:
            "If return true, means the OTP is correct and user needs to be created before continue the authentication process.",
          type: "bool",
        },
      ],
    },
  };
}
/**
 * The base class definition for classicPassportOtpActionReq
 **/
export class ClassicPassportOtpActionReq {
  /**
   *
   * @type {string}
   **/
  #value: string = "";
  /**
   *
   * @returns {string}
   **/
  get value() {
    return this.#value;
  }
  /**
   *
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
   *
   * @type {string}
   **/
  #otp: string = "";
  /**
   *
   * @returns {string}
   **/
  get otp() {
    return this.#otp;
  }
  /**
   *
   * @type {string}
   **/
  set otp(value: string) {
    this.#otp = String(value);
  }
  setOtp(value: string) {
    this.otp = value;
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
    const d = data as Partial<ClassicPassportOtpActionReq>;
    if (d.value !== undefined) {
      this.value = d.value;
    }
    if (d.otp !== undefined) {
      this.otp = d.otp;
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
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      value: "value",
      otp: "otp",
    };
  }
  /**
   * Creates an instance of ClassicPassportOtpActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ClassicPassportOtpActionReqType) {
    return new ClassicPassportOtpActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of ClassicPassportOtpActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ClassicPassportOtpActionReqType>) {
    return new ClassicPassportOtpActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ClassicPassportOtpActionReqType>,
  ): InstanceType<typeof ClassicPassportOtpActionReq> {
    return new ClassicPassportOtpActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ClassicPassportOtpActionReq> {
    return new ClassicPassportOtpActionReq(this.toJSON());
  }
}
export abstract class ClassicPassportOtpActionReqFactory {
  abstract create(data: unknown): ClassicPassportOtpActionReq;
}
/**
 * The base type definition for classicPassportOtpActionReq
 **/
export type ClassicPassportOtpActionReqType = {
  /**
   *
   * @type {string}
   **/
  value: string;
  /**
   *
   * @type {string}
   **/
  otp: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ClassicPassportOtpActionReqType {}
/**
 * The base class definition for classicPassportOtpActionRes
 **/
export class ClassicPassportOtpActionRes {
  /**
   * Upon successful authentication, there will be a session dto generated, which is a ground information of authorized user and can be stored in front-end.
   * @type {UserSessionDto}
   **/
  #session?: UserSessionDto | null = undefined;
  /**
   * Upon successful authentication, there will be a session dto generated, which is a ground information of authorized user and can be stored in front-end.
   * @returns {UserSessionDto}
   **/
  get session() {
    return this.#session;
  }
  /**
   * Upon successful authentication, there will be a session dto generated, which is a ground information of authorized user and can be stored in front-end.
   * @type {UserSessionDto}
   **/
  set session(value: UserSessionDto | null | undefined) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof UserSessionDto) {
      this.#session = value;
    } else {
      this.#session = new UserSessionDto(value);
    }
  }
  setSession(value: UserSessionDto | null | undefined) {
    this.session = value;
    return this;
  }
  /**
   * If time based otp is available, we add it response to make it easier for ui.
   * @type {string}
   **/
  #totpUrl: string = "";
  /**
   * If time based otp is available, we add it response to make it easier for ui.
   * @returns {string}
   **/
  get totpUrl() {
    return this.#totpUrl;
  }
  /**
   * If time based otp is available, we add it response to make it easier for ui.
   * @type {string}
   **/
  set totpUrl(value: string) {
    this.#totpUrl = String(value);
  }
  setTotpUrl(value: string) {
    this.totpUrl = value;
    return this;
  }
  /**
   * The session secret will be used to call complete user registration api.
   * @type {string}
   **/
  #sessionSecret: string = "";
  /**
   * The session secret will be used to call complete user registration api.
   * @returns {string}
   **/
  get sessionSecret() {
    return this.#sessionSecret;
  }
  /**
   * The session secret will be used to call complete user registration api.
   * @type {string}
   **/
  set sessionSecret(value: string) {
    this.#sessionSecret = String(value);
  }
  setSessionSecret(value: string) {
    this.sessionSecret = value;
    return this;
  }
  /**
   * If return true, means the OTP is correct and user needs to be created before continue the authentication process.
   * @type {boolean}
   **/
  #continueWithCreation!: boolean;
  /**
   * If return true, means the OTP is correct and user needs to be created before continue the authentication process.
   * @returns {boolean}
   **/
  get continueWithCreation() {
    return this.#continueWithCreation;
  }
  /**
   * If return true, means the OTP is correct and user needs to be created before continue the authentication process.
   * @type {boolean}
   **/
  set continueWithCreation(value: boolean) {
    this.#continueWithCreation = Boolean(value);
  }
  setContinueWithCreation(value: boolean) {
    this.continueWithCreation = value;
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
    const d = data as Partial<ClassicPassportOtpActionRes>;
    if (d.session !== undefined) {
      this.session = d.session;
    }
    if (d.totpUrl !== undefined) {
      this.totpUrl = d.totpUrl;
    }
    if (d.sessionSecret !== undefined) {
      this.sessionSecret = d.sessionSecret;
    }
    if (d.continueWithCreation !== undefined) {
      this.continueWithCreation = d.continueWithCreation;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      session: this.#session,
      totpUrl: this.#totpUrl,
      sessionSecret: this.#sessionSecret,
      continueWithCreation: this.#continueWithCreation,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      session: "session",
      totpUrl: "totpUrl",
      sessionSecret: "sessionSecret",
      continueWithCreation: "continueWithCreation",
    };
  }
  /**
   * Creates an instance of ClassicPassportOtpActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ClassicPassportOtpActionResType) {
    return new ClassicPassportOtpActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of ClassicPassportOtpActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ClassicPassportOtpActionResType>) {
    return new ClassicPassportOtpActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ClassicPassportOtpActionResType>,
  ): InstanceType<typeof ClassicPassportOtpActionRes> {
    return new ClassicPassportOtpActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ClassicPassportOtpActionRes> {
    return new ClassicPassportOtpActionRes(this.toJSON());
  }
}
export abstract class ClassicPassportOtpActionResFactory {
  abstract create(data: unknown): ClassicPassportOtpActionRes;
}
/**
 * The base type definition for classicPassportOtpActionRes
 **/
export type ClassicPassportOtpActionResType = {
  /**
   * Upon successful authentication, there will be a session dto generated, which is a ground information of authorized user and can be stored in front-end.
   * @type {UserSessionDto}
   **/
  session?: UserSessionDto;
  /**
   * If time based otp is available, we add it response to make it easier for ui.
   * @type {string}
   **/
  totpUrl: string;
  /**
   * The session secret will be used to call complete user registration api.
   * @type {string}
   **/
  sessionSecret: string;
  /**
   * If return true, means the OTP is correct and user needs to be created before continue the authentication process.
   * @type {boolean}
   **/
  continueWithCreation: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ClassicPassportOtpActionResType {}
