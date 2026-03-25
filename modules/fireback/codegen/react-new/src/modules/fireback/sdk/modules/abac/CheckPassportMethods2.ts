import {
  FetchxContext,
  fetchx,
  handleFetchResponse,
  type TypedRequestInit,
  type TypedResponse,
} from "../../sdk/common/fetchx";
import { GResponse } from "../../sdk/envelopes/index";
import { buildUrl } from "../../sdk/common/buildUrl";
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "react-query";
import { useFetchxContext } from "../../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action CheckPassportMethods2
 */
export type CheckPassportMethods2ActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type CheckPassportMethods2ActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<CheckPassportMethods2ActionRes>,
    unknown[]
  >,
  "queryKey"
> &
  CheckPassportMethods2ActionOptions &
  Partial<{
    creatorFn: (item: unknown) => CheckPassportMethods2ActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext;
  };
export const useCheckPassportMethods2ActionQuery = (
  options: CheckPassportMethods2ActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn: any = () => {
    setCompleteState(false);
    return CheckPassportMethods2Action.Fetch(
      {
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

  const result = useQuery({
    queryKey: [CheckPassportMethods2Action.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type CheckPassportMethods2ActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  CheckPassportMethods2ActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => CheckPassportMethods2ActionRes;
  }>;
export const useCheckPassportMethods2Action = (
  options?: CheckPassportMethods2ActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return CheckPassportMethods2Action.Fetch(
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
 * CheckPassportMethods2Action
 */
export class CheckPassportMethods2Action {
  //
  static URL = "/passports/available-methods2";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(CheckPassportMethods2Action.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<CheckPassportMethods2ActionRes>, unknown, unknown>(
      overrideUrl ?? CheckPassportMethods2Action.NewUrl(qs),
      {
        method: CheckPassportMethods2Action.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<unknown, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?:
        | ((item: unknown) => CheckPassportMethods2ActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new CheckPassportMethods2ActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new CheckPassportMethods2ActionRes(item));
    const res = await CheckPassportMethods2Action.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    const x = handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<CheckPassportMethods2ActionRes>();
        if (creatorFn) {
          resp.setCreator(creatorFn);
        }
        resp.inject(data);
        return resp;
      },
      onMessage,
      init?.signal,
    );

    return x;
  };
  static Definition = {
    name: "CheckPassportMethods2",
    url: "/passports/available-methods2",
    method: "get",
    description:
      "Publicly available information to create the authentication form, and show users how they can signin or signup to the system. Based on the PassportMethod entities, it will compute the available methods for the user, considering their region (IP for example)",
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "email",
          type: "bool",
          default: false,
        },
        {
          name: "phone",
          type: "bool",
          default: false,
        },
        {
          name: "google",
          type: "bool",
          default: false,
        },
        {
          name: "facebook",
          type: "bool",
          default: false,
        },
        {
          name: "googleOAuthClientKey",
          type: "string",
        },
        {
          name: "facebookAppId",
          type: "string",
        },
        {
          name: "enabledRecaptcha2",
          type: "bool",
          default: false,
        },
        {
          name: "recaptcha2ClientKey",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for checkPassportMethods2ActionRes
 **/
export class CheckPassportMethods2ActionRes {
  /**
   *
   * @type {boolean}
   **/
  #email: boolean = false;
  /**
   *
   * @returns {boolean}
   **/
  get email() {
    return this.#email;
  }
  /**
   *
   * @type {boolean}
   **/
  set email(value: boolean) {
    this.#email = Boolean(value);
  }
  setEmail(value: boolean) {
    this.email = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #phone: boolean = false;
  /**
   *
   * @returns {boolean}
   **/
  get phone() {
    return this.#phone;
  }
  /**
   *
   * @type {boolean}
   **/
  set phone(value: boolean) {
    this.#phone = Boolean(value);
  }
  setPhone(value: boolean) {
    this.phone = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #google: boolean = false;
  /**
   *
   * @returns {boolean}
   **/
  get google() {
    return this.#google;
  }
  /**
   *
   * @type {boolean}
   **/
  set google(value: boolean) {
    this.#google = Boolean(value);
  }
  setGoogle(value: boolean) {
    this.google = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #facebook: boolean = false;
  /**
   *
   * @returns {boolean}
   **/
  get facebook() {
    return this.#facebook;
  }
  /**
   *
   * @type {boolean}
   **/
  set facebook(value: boolean) {
    this.#facebook = Boolean(value);
  }
  setFacebook(value: boolean) {
    this.facebook = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #googleOAuthClientKey: string = "";
  /**
   *
   * @returns {string}
   **/
  get googleOAuthClientKey() {
    return this.#googleOAuthClientKey;
  }
  /**
   *
   * @type {string}
   **/
  set googleOAuthClientKey(value: string) {
    this.#googleOAuthClientKey = String(value);
  }
  setGoogleOAuthClientKey(value: string) {
    this.googleOAuthClientKey = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #facebookAppId: string = "";
  /**
   *
   * @returns {string}
   **/
  get facebookAppId() {
    return this.#facebookAppId;
  }
  /**
   *
   * @type {string}
   **/
  set facebookAppId(value: string) {
    this.#facebookAppId = String(value);
  }
  setFacebookAppId(value: string) {
    this.facebookAppId = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #enabledRecaptcha2: boolean = false;
  /**
   *
   * @returns {boolean}
   **/
  get enabledRecaptcha2() {
    return this.#enabledRecaptcha2;
  }
  /**
   *
   * @type {boolean}
   **/
  set enabledRecaptcha2(value: boolean) {
    this.#enabledRecaptcha2 = Boolean(value);
  }
  setEnabledRecaptcha2(value: boolean) {
    this.enabledRecaptcha2 = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #recaptcha2ClientKey: string = "";
  /**
   *
   * @returns {string}
   **/
  get recaptcha2ClientKey() {
    return this.#recaptcha2ClientKey;
  }
  /**
   *
   * @type {string}
   **/
  set recaptcha2ClientKey(value: string) {
    this.#recaptcha2ClientKey = String(value);
  }
  setRecaptcha2ClientKey(value: string) {
    this.recaptcha2ClientKey = value;
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
    const d = data as Partial<CheckPassportMethods2ActionRes>;
    if (d.email !== undefined) {
      this.email = d.email;
    }
    if (d.phone !== undefined) {
      this.phone = d.phone;
    }
    if (d.google !== undefined) {
      this.google = d.google;
    }
    if (d.facebook !== undefined) {
      this.facebook = d.facebook;
    }
    if (d.googleOAuthClientKey !== undefined) {
      this.googleOAuthClientKey = d.googleOAuthClientKey;
    }
    if (d.facebookAppId !== undefined) {
      this.facebookAppId = d.facebookAppId;
    }
    if (d.enabledRecaptcha2 !== undefined) {
      this.enabledRecaptcha2 = d.enabledRecaptcha2;
    }
    if (d.recaptcha2ClientKey !== undefined) {
      this.recaptcha2ClientKey = d.recaptcha2ClientKey;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      email: this.#email,
      phone: this.#phone,
      google: this.#google,
      facebook: this.#facebook,
      googleOAuthClientKey: this.#googleOAuthClientKey,
      facebookAppId: this.#facebookAppId,
      enabledRecaptcha2: this.#enabledRecaptcha2,
      recaptcha2ClientKey: this.#recaptcha2ClientKey,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      email: "email",
      phone: "phone",
      google: "google",
      facebook: "facebook",
      googleOAuthClientKey: "googleOAuthClientKey",
      facebookAppId: "facebookAppId",
      enabledRecaptcha2: "enabledRecaptcha2",
      recaptcha2ClientKey: "recaptcha2ClientKey",
    };
  }
  /**
   * Creates an instance of CheckPassportMethods2ActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CheckPassportMethods2ActionResType) {
    return new CheckPassportMethods2ActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of CheckPassportMethods2ActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<CheckPassportMethods2ActionResType>,
  ) {
    return new CheckPassportMethods2ActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CheckPassportMethods2ActionResType>,
  ): InstanceType<typeof CheckPassportMethods2ActionRes> {
    return new CheckPassportMethods2ActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CheckPassportMethods2ActionRes> {
    return new CheckPassportMethods2ActionRes(this.toJSON());
  }
}
export abstract class CheckPassportMethods2ActionResFactory {
  abstract create(data: unknown): CheckPassportMethods2ActionRes;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for checkPassportMethods2ActionRes
 **/
export type CheckPassportMethods2ActionResType = {
  /**
   *
   * @type {boolean}
   **/
  email: boolean;
  /**
   *
   * @type {boolean}
   **/
  phone: boolean;
  /**
   *
   * @type {boolean}
   **/
  google: boolean;
  /**
   *
   * @type {boolean}
   **/
  facebook: boolean;
  /**
   *
   * @type {string}
   **/
  googleOAuthClientKey: string;
  /**
   *
   * @type {string}
   **/
  facebookAppId: string;
  /**
   *
   * @type {boolean}
   **/
  enabledRecaptcha2: boolean;
  /**
   *
   * @type {string}
   **/
  recaptcha2ClientKey: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CheckPassportMethods2ActionResType {}
