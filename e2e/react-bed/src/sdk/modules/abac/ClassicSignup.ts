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
import { withPrefix } from "../../sdk/common/withPrefix";
/**
 * Action to communicate with the action ClassicSignup
 */
export type ClassicSignupActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type ClassicSignupActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  ClassicSignupActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => ClassicSignupActionRes;
  }>;
export const useClassicSignupAction = (
  options?: ClassicSignupActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: ClassicSignupActionReq) => {
    setCompleteState(false);
    return ClassicSignupAction.Fetch(
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
 * ClassicSignupAction
 */
export class ClassicSignupAction {
  //
  static URL = "/passports/signup/classic";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(ClassicSignupAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<ClassicSignupActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<ClassicSignupActionRes>,
      ClassicSignupActionReq,
      unknown
    >(
      overrideUrl ?? ClassicSignupAction.NewUrl(qs),
      {
        method: ClassicSignupAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<ClassicSignupActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => ClassicSignupActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new ClassicSignupActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new ClassicSignupActionRes(item));
    const res = await ClassicSignupAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<ClassicSignupActionRes>();
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
    name: "ClassicSignup",
    cliName: "up",
    url: "/passports/signup/classic",
    method: "post",
    description:
      "Signup a user into system via public access (aka website visitors) using either email or phone number.",
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
          name: "sessionSecret",
          description:
            "Required when the account creation requires recaptcha, or otp approval first. If such requirements are there, you first need to follow the otp apis, get the session secret and pass it here to complete the setup.",
          type: "string",
        },
        {
          name: "type",
          type: "enum",
          of: [
            {
              k: "phonenumber",
            },
            {
              k: "email",
            },
          ],
          tags: {
            validate: "required",
          },
        },
        {
          name: "password",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "firstName",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "lastName",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "inviteId",
          type: "string?",
        },
        {
          name: "publicJoinKeyId",
          type: "string?",
        },
        {
          name: "workspaceTypeId",
          type: "string?",
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
            "Returns the user session in case that signup is completely successful.",
          type: "one",
          target: "UserSessionDto",
        },
        {
          name: "totpUrl",
          description:
            "If time based otp is available, we add it response to make it easier for ui.",
          type: "string",
        },
        {
          name: "continueToTotp",
          description:
            "Returns true and session will be empty if, the totp is required by the installation. In such scenario, you need to forward user to setup totp screen.",
          type: "bool",
        },
        {
          name: "forcedTotp",
          description:
            "Determines if user must complete totp in order to continue based on workspace or installation",
          type: "bool",
        },
      ],
    },
  };
}
/**
 * The base class definition for classicSignupActionReq
 **/
export class ClassicSignupActionReq {
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
   * Required when the account creation requires recaptcha, or otp approval first. If such requirements are there, you first need to follow the otp apis, get the session secret and pass it here to complete the setup.
   * @type {string}
   **/
  #sessionSecret: string = "";
  /**
   * Required when the account creation requires recaptcha, or otp approval first. If such requirements are there, you first need to follow the otp apis, get the session secret and pass it here to complete the setup.
   * @returns {string}
   **/
  get sessionSecret() {
    return this.#sessionSecret;
  }
  /**
   * Required when the account creation requires recaptcha, or otp approval first. If such requirements are there, you first need to follow the otp apis, get the session secret and pass it here to complete the setup.
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
   *
   * @type {"phonenumber" | "email"}
   **/
  #type!: "phonenumber" | "email";
  /**
   *
   * @returns {"phonenumber" | "email"}
   **/
  get type() {
    return this.#type;
  }
  /**
   *
   * @type {"phonenumber" | "email"}
   **/
  set type(value: "phonenumber" | "email") {
    this.#type = value;
  }
  setType(value: "phonenumber" | "email") {
    this.type = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #password: string = "";
  /**
   *
   * @returns {string}
   **/
  get password() {
    return this.#password;
  }
  /**
   *
   * @type {string}
   **/
  set password(value: string) {
    this.#password = String(value);
  }
  setPassword(value: string) {
    this.password = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #firstName: string = "";
  /**
   *
   * @returns {string}
   **/
  get firstName() {
    return this.#firstName;
  }
  /**
   *
   * @type {string}
   **/
  set firstName(value: string) {
    this.#firstName = String(value);
  }
  setFirstName(value: string) {
    this.firstName = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #lastName: string = "";
  /**
   *
   * @returns {string}
   **/
  get lastName() {
    return this.#lastName;
  }
  /**
   *
   * @type {string}
   **/
  set lastName(value: string) {
    this.#lastName = String(value);
  }
  setLastName(value: string) {
    this.lastName = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get inviteId() {
    return this.#inviteId;
  }
  /**
   *
   * @type {string}
   **/
  set inviteId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#inviteId = correctType ? value : String(value);
  }
  setInviteId(value: string | null | undefined) {
    this.inviteId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #publicJoinKeyId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get publicJoinKeyId() {
    return this.#publicJoinKeyId;
  }
  /**
   *
   * @type {string}
   **/
  set publicJoinKeyId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#publicJoinKeyId = correctType ? value : String(value);
  }
  setPublicJoinKeyId(value: string | null | undefined) {
    this.publicJoinKeyId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #workspaceTypeId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get workspaceTypeId() {
    return this.#workspaceTypeId;
  }
  /**
   *
   * @type {string}
   **/
  set workspaceTypeId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#workspaceTypeId = correctType ? value : String(value);
  }
  setWorkspaceTypeId(value: string | null | undefined) {
    this.workspaceTypeId = value;
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
    const d = data as Partial<ClassicSignupActionReq>;
    if (d.value !== undefined) {
      this.value = d.value;
    }
    if (d.sessionSecret !== undefined) {
      this.sessionSecret = d.sessionSecret;
    }
    if (d.type !== undefined) {
      this.type = d.type;
    }
    if (d.password !== undefined) {
      this.password = d.password;
    }
    if (d.firstName !== undefined) {
      this.firstName = d.firstName;
    }
    if (d.lastName !== undefined) {
      this.lastName = d.lastName;
    }
    if (d.inviteId !== undefined) {
      this.inviteId = d.inviteId;
    }
    if (d.publicJoinKeyId !== undefined) {
      this.publicJoinKeyId = d.publicJoinKeyId;
    }
    if (d.workspaceTypeId !== undefined) {
      this.workspaceTypeId = d.workspaceTypeId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      value: this.#value,
      sessionSecret: this.#sessionSecret,
      type: this.#type,
      password: this.#password,
      firstName: this.#firstName,
      lastName: this.#lastName,
      inviteId: this.#inviteId,
      publicJoinKeyId: this.#publicJoinKeyId,
      workspaceTypeId: this.#workspaceTypeId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      value: "value",
      sessionSecret: "sessionSecret",
      type: "type",
      password: "password",
      firstName: "firstName",
      lastName: "lastName",
      inviteId: "inviteId",
      publicJoinKeyId: "publicJoinKeyId",
      workspaceTypeId: "workspaceTypeId",
    };
  }
  /**
   * Creates an instance of ClassicSignupActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ClassicSignupActionReqType) {
    return new ClassicSignupActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of ClassicSignupActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ClassicSignupActionReqType>) {
    return new ClassicSignupActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ClassicSignupActionReqType>,
  ): InstanceType<typeof ClassicSignupActionReq> {
    return new ClassicSignupActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ClassicSignupActionReq> {
    return new ClassicSignupActionReq(this.toJSON());
  }
}
export abstract class ClassicSignupActionReqFactory {
  abstract create(data: unknown): ClassicSignupActionReq;
}
/**
 * The base type definition for classicSignupActionReq
 **/
export type ClassicSignupActionReqType = {
  /**
   *
   * @type {string}
   **/
  value: string;
  /**
   * Required when the account creation requires recaptcha, or otp approval first. If such requirements are there, you first need to follow the otp apis, get the session secret and pass it here to complete the setup.
   * @type {string}
   **/
  sessionSecret: string;
  /**
   *
   * @type {"phonenumber" | "email"}
   **/
  type: "phonenumber" | "email";
  /**
   *
   * @type {string}
   **/
  password: string;
  /**
   *
   * @type {string}
   **/
  firstName: string;
  /**
   *
   * @type {string}
   **/
  lastName: string;
  /**
   *
   * @type {string}
   **/
  inviteId?: string;
  /**
   *
   * @type {string}
   **/
  publicJoinKeyId?: string;
  /**
   *
   * @type {string}
   **/
  workspaceTypeId?: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ClassicSignupActionReqType {}
/**
 * The base class definition for classicSignupActionRes
 **/
export class ClassicSignupActionRes {
  /**
   * Returns the user session in case that signup is completely successful.
   * @type {UserSessionDto}
   **/
  #session!: UserSessionDto;
  /**
   * Returns the user session in case that signup is completely successful.
   * @returns {UserSessionDto}
   **/
  get session() {
    return this.#session;
  }
  /**
   * Returns the user session in case that signup is completely successful.
   * @type {UserSessionDto}
   **/
  set session(value: UserSessionDto) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof UserSessionDto) {
      this.#session = value;
    } else {
      this.#session = new UserSessionDto(value);
    }
  }
  setSession(value: UserSessionDto) {
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
   * Returns true and session will be empty if, the totp is required by the installation. In such scenario, you need to forward user to setup totp screen.
   * @type {boolean}
   **/
  #continueToTotp!: boolean;
  /**
   * Returns true and session will be empty if, the totp is required by the installation. In such scenario, you need to forward user to setup totp screen.
   * @returns {boolean}
   **/
  get continueToTotp() {
    return this.#continueToTotp;
  }
  /**
   * Returns true and session will be empty if, the totp is required by the installation. In such scenario, you need to forward user to setup totp screen.
   * @type {boolean}
   **/
  set continueToTotp(value: boolean) {
    this.#continueToTotp = Boolean(value);
  }
  setContinueToTotp(value: boolean) {
    this.continueToTotp = value;
    return this;
  }
  /**
   * Determines if user must complete totp in order to continue based on workspace or installation
   * @type {boolean}
   **/
  #forcedTotp!: boolean;
  /**
   * Determines if user must complete totp in order to continue based on workspace or installation
   * @returns {boolean}
   **/
  get forcedTotp() {
    return this.#forcedTotp;
  }
  /**
   * Determines if user must complete totp in order to continue based on workspace or installation
   * @type {boolean}
   **/
  set forcedTotp(value: boolean) {
    this.#forcedTotp = Boolean(value);
  }
  setForcedTotp(value: boolean) {
    this.forcedTotp = value;
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
    const d = data as Partial<ClassicSignupActionRes>;
    if (d.session !== undefined) {
      this.session = d.session;
    }
    if (d.totpUrl !== undefined) {
      this.totpUrl = d.totpUrl;
    }
    if (d.continueToTotp !== undefined) {
      this.continueToTotp = d.continueToTotp;
    }
    if (d.forcedTotp !== undefined) {
      this.forcedTotp = d.forcedTotp;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<ClassicSignupActionRes>;
    if (!(d.session instanceof UserSessionDto)) {
      this.session = new UserSessionDto(d.session || {});
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
      continueToTotp: this.#continueToTotp,
      forcedTotp: this.#forcedTotp,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      session$: "session",
      get session() {
        return withPrefix("session", UserSessionDto.Fields);
      },
      totpUrl: "totpUrl",
      continueToTotp: "continueToTotp",
      forcedTotp: "forcedTotp",
    };
  }
  /**
   * Creates an instance of ClassicSignupActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ClassicSignupActionResType) {
    return new ClassicSignupActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of ClassicSignupActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ClassicSignupActionResType>) {
    return new ClassicSignupActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ClassicSignupActionResType>,
  ): InstanceType<typeof ClassicSignupActionRes> {
    return new ClassicSignupActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ClassicSignupActionRes> {
    return new ClassicSignupActionRes(this.toJSON());
  }
}
export abstract class ClassicSignupActionResFactory {
  abstract create(data: unknown): ClassicSignupActionRes;
}
/**
 * The base type definition for classicSignupActionRes
 **/
export type ClassicSignupActionResType = {
  /**
   * Returns the user session in case that signup is completely successful.
   * @type {UserSessionDto}
   **/
  session: UserSessionDto;
  /**
   * If time based otp is available, we add it response to make it easier for ui.
   * @type {string}
   **/
  totpUrl: string;
  /**
   * Returns true and session will be empty if, the totp is required by the installation. In such scenario, you need to forward user to setup totp screen.
   * @type {boolean}
   **/
  continueToTotp: boolean;
  /**
   * Determines if user must complete totp in order to continue based on workspace or installation
   * @type {boolean}
   **/
  forcedTotp: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ClassicSignupActionResType {}
