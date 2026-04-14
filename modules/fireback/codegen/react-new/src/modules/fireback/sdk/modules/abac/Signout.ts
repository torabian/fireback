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
 * Action to communicate with the action Signout
 */
export type SignoutActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type SignoutActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  SignoutActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => SignoutActionRes;
  }>;
export const useSignoutAction = (options?: SignoutActionMutationOptions) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: SignoutActionReq) => {
    setCompleteState(false);
    return SignoutAction.Fetch(
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
 * SignoutAction
 */
export class SignoutAction {
  //
  static URL = "/passport/signout";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(SignoutAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<SignoutActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<SignoutActionRes>, SignoutActionReq, unknown>(
      overrideUrl ?? SignoutAction.NewUrl(qs),
      {
        method: SignoutAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<SignoutActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => SignoutActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new SignoutActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new SignoutActionRes(item));
    const res = await SignoutAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<SignoutActionRes>();
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
    name: "Signout",
    url: "/passport/signout",
    method: "post",
    description:
      "Signout the user, clears cookies or does anything else if needed.",
    in: {
      fields: [
        {
          name: "clear",
          type: "bool?",
        },
      ],
    },
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "okay",
          type: "bool",
        },
      ],
    },
  };
}
/**
 * The base class definition for signoutActionReq
 **/
export class SignoutActionReq {
  /**
   *
   * @type {boolean}
   **/
  #clear?: boolean | null = undefined;
  /**
   *
   * @returns {boolean}
   **/
  get clear() {
    return this.#clear;
  }
  /**
   *
   * @type {boolean}
   **/
  set clear(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#clear = correctType ? value : Boolean(value);
  }
  setClear(value: boolean | null | undefined) {
    this.clear = value;
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
    const d = data as Partial<SignoutActionReq>;
    if (d.clear !== undefined) {
      this.clear = d.clear;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      clear: this.#clear,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      clear: "clear",
    };
  }
  /**
   * Creates an instance of SignoutActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: SignoutActionReqType) {
    return new SignoutActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of SignoutActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<SignoutActionReqType>) {
    return new SignoutActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<SignoutActionReqType>,
  ): InstanceType<typeof SignoutActionReq> {
    return new SignoutActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof SignoutActionReq> {
    return new SignoutActionReq(this.toJSON());
  }
}
export abstract class SignoutActionReqFactory {
  abstract create(data: unknown): SignoutActionReq;
}
/**
 * The base type definition for signoutActionReq
 **/
export type SignoutActionReqType = {
  /**
   *
   * @type {boolean}
   **/
  clear?: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SignoutActionReqType {}
/**
 * The base class definition for signoutActionRes
 **/
export class SignoutActionRes {
  /**
   *
   * @type {boolean}
   **/
  #okay!: boolean;
  /**
   *
   * @returns {boolean}
   **/
  get okay() {
    return this.#okay;
  }
  /**
   *
   * @type {boolean}
   **/
  set okay(value: boolean) {
    this.#okay = Boolean(value);
  }
  setOkay(value: boolean) {
    this.okay = value;
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
    const d = data as Partial<SignoutActionRes>;
    if (d.okay !== undefined) {
      this.okay = d.okay;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      okay: this.#okay,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      okay: "okay",
    };
  }
  /**
   * Creates an instance of SignoutActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: SignoutActionResType) {
    return new SignoutActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of SignoutActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<SignoutActionResType>) {
    return new SignoutActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<SignoutActionResType>,
  ): InstanceType<typeof SignoutActionRes> {
    return new SignoutActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof SignoutActionRes> {
    return new SignoutActionRes(this.toJSON());
  }
}
export abstract class SignoutActionResFactory {
  abstract create(data: unknown): SignoutActionRes;
}
/**
 * The base type definition for signoutActionRes
 **/
export type SignoutActionResType = {
  /**
   *
   * @type {boolean}
   **/
  okay: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SignoutActionResType {}
