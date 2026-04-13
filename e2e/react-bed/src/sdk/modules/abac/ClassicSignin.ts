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
 * Action to communicate with the action ClassicSignin
 */
export type ClassicSigninActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type ClassicSigninActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  ClassicSigninActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => ClassicSigninActionRes;
  }>;
export const useClassicSigninAction = (
  options?: ClassicSigninActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: ClassicSigninActionReq) => {
    setCompleteState(false);
    return ClassicSigninAction.Fetch(
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
 * ClassicSigninAction
 */
export class ClassicSigninAction {
  //
  static URL = "/passports/signin/classic";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(ClassicSigninAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<ClassicSigninActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<ClassicSigninActionRes>,
      ClassicSigninActionReq,
      unknown
    >(
      overrideUrl ?? ClassicSigninAction.NewUrl(qs),
      {
        method: ClassicSigninAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<ClassicSigninActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => ClassicSigninActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new ClassicSigninActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new ClassicSigninActionRes(item));
    const res = await ClassicSigninAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<ClassicSigninActionRes>();
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
    name: "ClassicSignin",
    cliName: "in",
    url: "/passports/signin/classic",
    method: "post",
    description:
      "Signin publicly to and account using class passports (email, password)",
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
          name: "password",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "totpCode",
          description:
            "Accepts login with totp code. If enabled, first login would return a success response with next[enter-totp] value and ui can understand that user needs to be navigated into the screen other screen.",
          type: "string",
        },
        {
          name: "sessionSecret",
          description:
            "Session secret when logging in to the application requires more steps to complete.",
          type: "string",
        },
      ],
    },
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "session",
          type: "one",
          target: "UserSessionDto",
        },
        {
          name: "next",
          description: "The next possible action which is suggested.",
          type: "slice",
          primitive: "string",
        },
        {
          name: "totpUrl",
          description:
            "In case the account doesn't have totp, but enforced by installation, this value will contain the link",
          type: "string",
        },
        {
          name: "sessionSecret",
          description:
            "Returns a secret session if the authentication requires more steps.",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for classicSigninActionReq
 **/
export class ClassicSigninActionReq {
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
   * Accepts login with totp code. If enabled, first login would return a success response with next[enter-totp] value and ui can understand that user needs to be navigated into the screen other screen.
   * @type {string}
   **/
  #totpCode: string = "";
  /**
   * Accepts login with totp code. If enabled, first login would return a success response with next[enter-totp] value and ui can understand that user needs to be navigated into the screen other screen.
   * @returns {string}
   **/
  get totpCode() {
    return this.#totpCode;
  }
  /**
   * Accepts login with totp code. If enabled, first login would return a success response with next[enter-totp] value and ui can understand that user needs to be navigated into the screen other screen.
   * @type {string}
   **/
  set totpCode(value: string) {
    this.#totpCode = String(value);
  }
  setTotpCode(value: string) {
    this.totpCode = value;
    return this;
  }
  /**
   * Session secret when logging in to the application requires more steps to complete.
   * @type {string}
   **/
  #sessionSecret: string = "";
  /**
   * Session secret when logging in to the application requires more steps to complete.
   * @returns {string}
   **/
  get sessionSecret() {
    return this.#sessionSecret;
  }
  /**
   * Session secret when logging in to the application requires more steps to complete.
   * @type {string}
   **/
  set sessionSecret(value: string) {
    this.#sessionSecret = String(value);
  }
  setSessionSecret(value: string) {
    this.sessionSecret = value;
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
    const d = data as Partial<ClassicSigninActionReq>;
    if (d.value !== undefined) {
      this.value = d.value;
    }
    if (d.password !== undefined) {
      this.password = d.password;
    }
    if (d.totpCode !== undefined) {
      this.totpCode = d.totpCode;
    }
    if (d.sessionSecret !== undefined) {
      this.sessionSecret = d.sessionSecret;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      value: this.#value,
      password: this.#password,
      totpCode: this.#totpCode,
      sessionSecret: this.#sessionSecret,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      value: "value",
      password: "password",
      totpCode: "totpCode",
      sessionSecret: "sessionSecret",
    };
  }
  /**
   * Creates an instance of ClassicSigninActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ClassicSigninActionReqType) {
    return new ClassicSigninActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of ClassicSigninActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ClassicSigninActionReqType>) {
    return new ClassicSigninActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ClassicSigninActionReqType>,
  ): InstanceType<typeof ClassicSigninActionReq> {
    return new ClassicSigninActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ClassicSigninActionReq> {
    return new ClassicSigninActionReq(this.toJSON());
  }
}
export abstract class ClassicSigninActionReqFactory {
  abstract create(data: unknown): ClassicSigninActionReq;
}
/**
 * The base type definition for classicSigninActionReq
 **/
export type ClassicSigninActionReqType = {
  /**
   *
   * @type {string}
   **/
  value: string;
  /**
   *
   * @type {string}
   **/
  password: string;
  /**
   * Accepts login with totp code. If enabled, first login would return a success response with next[enter-totp] value and ui can understand that user needs to be navigated into the screen other screen.
   * @type {string}
   **/
  totpCode: string;
  /**
   * Session secret when logging in to the application requires more steps to complete.
   * @type {string}
   **/
  sessionSecret: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ClassicSigninActionReqType {}
/**
 * The base class definition for classicSigninActionRes
 **/
export class ClassicSigninActionRes {
  /**
   *
   * @type {UserSessionDto}
   **/
  #session!: UserSessionDto;
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
   * The next possible action which is suggested.
   * @type {string[]}
   **/
  #next: string[] = [];
  /**
   * The next possible action which is suggested.
   * @returns {string[]}
   **/
  get next() {
    return this.#next;
  }
  /**
   * The next possible action which is suggested.
   * @type {string[]}
   **/
  set next(value: string[]) {
    this.#next = value;
  }
  setNext(value: string[]) {
    this.next = value;
    return this;
  }
  /**
   * In case the account doesn't have totp, but enforced by installation, this value will contain the link
   * @type {string}
   **/
  #totpUrl: string = "";
  /**
   * In case the account doesn't have totp, but enforced by installation, this value will contain the link
   * @returns {string}
   **/
  get totpUrl() {
    return this.#totpUrl;
  }
  /**
   * In case the account doesn't have totp, but enforced by installation, this value will contain the link
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
   * Returns a secret session if the authentication requires more steps.
   * @type {string}
   **/
  #sessionSecret: string = "";
  /**
   * Returns a secret session if the authentication requires more steps.
   * @returns {string}
   **/
  get sessionSecret() {
    return this.#sessionSecret;
  }
  /**
   * Returns a secret session if the authentication requires more steps.
   * @type {string}
   **/
  set sessionSecret(value: string) {
    this.#sessionSecret = String(value);
  }
  setSessionSecret(value: string) {
    this.sessionSecret = value;
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
    const d = data as Partial<ClassicSigninActionRes>;
    if (d.session !== undefined) {
      this.session = d.session;
    }
    if (d.next !== undefined) {
      this.next = d.next;
    }
    if (d.totpUrl !== undefined) {
      this.totpUrl = d.totpUrl;
    }
    if (d.sessionSecret !== undefined) {
      this.sessionSecret = d.sessionSecret;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<ClassicSigninActionRes>;
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
      next: this.#next,
      totpUrl: this.#totpUrl,
      sessionSecret: this.#sessionSecret,
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
      next$: "next",
      get next() {
        return "next[:i]";
      },
      totpUrl: "totpUrl",
      sessionSecret: "sessionSecret",
    };
  }
  /**
   * Creates an instance of ClassicSigninActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ClassicSigninActionResType) {
    return new ClassicSigninActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of ClassicSigninActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ClassicSigninActionResType>) {
    return new ClassicSigninActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ClassicSigninActionResType>,
  ): InstanceType<typeof ClassicSigninActionRes> {
    return new ClassicSigninActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ClassicSigninActionRes> {
    return new ClassicSigninActionRes(this.toJSON());
  }
}
export abstract class ClassicSigninActionResFactory {
  abstract create(data: unknown): ClassicSigninActionRes;
}
/**
 * The base type definition for classicSigninActionRes
 **/
export type ClassicSigninActionResType = {
  /**
   *
   * @type {UserSessionDto}
   **/
  session: UserSessionDto;
  /**
   * The next possible action which is suggested.
   * @type {string[]}
   **/
  next: string[];
  /**
   * In case the account doesn't have totp, but enforced by installation, this value will contain the link
   * @type {string}
   **/
  totpUrl: string;
  /**
   * Returns a secret session if the authentication requires more steps.
   * @type {string}
   **/
  sessionSecret: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ClassicSigninActionResType {}
