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
 * Action to communicate with the action OauthAuthenticate
 */
export type OauthAuthenticateActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type OauthAuthenticateActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  OauthAuthenticateActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => OauthAuthenticateActionRes;
  }>;
export const useOauthAuthenticateAction = (
  options?: OauthAuthenticateActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: OauthAuthenticateActionReq) => {
    setCompleteState(false);
    return OauthAuthenticateAction.Fetch(
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
 * OauthAuthenticateAction
 */
export class OauthAuthenticateAction {
  //
  static URL = "/passport/via-oauth";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(OauthAuthenticateAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<OauthAuthenticateActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<OauthAuthenticateActionRes>,
      OauthAuthenticateActionReq,
      unknown
    >(
      overrideUrl ?? OauthAuthenticateAction.NewUrl(qs),
      {
        method: OauthAuthenticateAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<OauthAuthenticateActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => OauthAuthenticateActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new OauthAuthenticateActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new OauthAuthenticateActionRes(item));
    const res = await OauthAuthenticateAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<OauthAuthenticateActionRes>();
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
    name: "OauthAuthenticate",
    url: "/passport/via-oauth",
    method: "post",
    description:
      "When a token is got from a oauth service such as google, we send the token here to authenticate the user. To me seems this doesn't need to have 2FA or anything, so we return the session directly, or maybe there needs to be next step.",
    in: {
      fields: [
        {
          name: "token",
          description:
            "The token that Auth2 provider returned to the front-end, which will be used to validate the backend",
          type: "string",
        },
        {
          name: "service",
          description:
            "The service name, such as 'google' which later backend will use to authorize the token and create the user.",
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
      ],
    },
  };
}
/**
 * The base class definition for oauthAuthenticateActionReq
 **/
export class OauthAuthenticateActionReq {
  /**
   * The token that Auth2 provider returned to the front-end, which will be used to validate the backend
   * @type {string}
   **/
  #token: string = "";
  /**
   * The token that Auth2 provider returned to the front-end, which will be used to validate the backend
   * @returns {string}
   **/
  get token() {
    return this.#token;
  }
  /**
   * The token that Auth2 provider returned to the front-end, which will be used to validate the backend
   * @type {string}
   **/
  set token(value: string) {
    this.#token = String(value);
  }
  setToken(value: string) {
    this.token = value;
    return this;
  }
  /**
   * The service name, such as 'google' which later backend will use to authorize the token and create the user.
   * @type {string}
   **/
  #service: string = "";
  /**
   * The service name, such as 'google' which later backend will use to authorize the token and create the user.
   * @returns {string}
   **/
  get service() {
    return this.#service;
  }
  /**
   * The service name, such as 'google' which later backend will use to authorize the token and create the user.
   * @type {string}
   **/
  set service(value: string) {
    this.#service = String(value);
  }
  setService(value: string) {
    this.service = value;
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
    const d = data as Partial<OauthAuthenticateActionReq>;
    if (d.token !== undefined) {
      this.token = d.token;
    }
    if (d.service !== undefined) {
      this.service = d.service;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      token: this.#token,
      service: this.#service,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      token: "token",
      service: "service",
    };
  }
  /**
   * Creates an instance of OauthAuthenticateActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: OauthAuthenticateActionReqType) {
    return new OauthAuthenticateActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of OauthAuthenticateActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<OauthAuthenticateActionReqType>) {
    return new OauthAuthenticateActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<OauthAuthenticateActionReqType>,
  ): InstanceType<typeof OauthAuthenticateActionReq> {
    return new OauthAuthenticateActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof OauthAuthenticateActionReq> {
    return new OauthAuthenticateActionReq(this.toJSON());
  }
}
export abstract class OauthAuthenticateActionReqFactory {
  abstract create(data: unknown): OauthAuthenticateActionReq;
}
/**
 * The base type definition for oauthAuthenticateActionReq
 **/
export type OauthAuthenticateActionReqType = {
  /**
   * The token that Auth2 provider returned to the front-end, which will be used to validate the backend
   * @type {string}
   **/
  token: string;
  /**
   * The service name, such as 'google' which later backend will use to authorize the token and create the user.
   * @type {string}
   **/
  service: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace OauthAuthenticateActionReqType {}
/**
 * The base class definition for oauthAuthenticateActionRes
 **/
export class OauthAuthenticateActionRes {
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
    const d = data as Partial<OauthAuthenticateActionRes>;
    if (d.session !== undefined) {
      this.session = d.session;
    }
    if (d.next !== undefined) {
      this.next = d.next;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<OauthAuthenticateActionRes>;
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
    };
  }
  /**
   * Creates an instance of OauthAuthenticateActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: OauthAuthenticateActionResType) {
    return new OauthAuthenticateActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of OauthAuthenticateActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<OauthAuthenticateActionResType>) {
    return new OauthAuthenticateActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<OauthAuthenticateActionResType>,
  ): InstanceType<typeof OauthAuthenticateActionRes> {
    return new OauthAuthenticateActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof OauthAuthenticateActionRes> {
    return new OauthAuthenticateActionRes(this.toJSON());
  }
}
export abstract class OauthAuthenticateActionResFactory {
  abstract create(data: unknown): OauthAuthenticateActionRes;
}
/**
 * The base type definition for oauthAuthenticateActionRes
 **/
export type OauthAuthenticateActionResType = {
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
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace OauthAuthenticateActionResType {}
