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
 * Action to communicate with the action ChangePassword
 */
export type ChangePasswordActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type ChangePasswordActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  ChangePasswordActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => ChangePasswordActionRes;
  }>;
export const useChangePasswordAction = (
  options?: ChangePasswordActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: ChangePasswordActionReq) => {
    setCompleteState(false);
    return ChangePasswordAction.Fetch(
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
 * ChangePasswordAction
 */
export class ChangePasswordAction {
  //
  static URL = "/passport/change-password";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(ChangePasswordAction.URL, undefined, qs);
  static Method = "post";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<ChangePasswordActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<ChangePasswordActionRes>,
      ChangePasswordActionReq,
      unknown
    >(
      overrideUrl ?? ChangePasswordAction.NewUrl(qs),
      {
        method: ChangePasswordAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<ChangePasswordActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => ChangePasswordActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new ChangePasswordActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new ChangePasswordActionRes(item));
    const res = await ChangePasswordAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<ChangePasswordActionRes>();
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
    name: "ChangePassword",
    cliName: "cp",
    url: "/passport/change-password",
    method: "post",
    description:
      "Change the password for a given passport of the user. User needs to be authenticated in order to be able to change the password for a given account.",
    in: {
      fields: [
        {
          name: "password",
          description: "New password meeting the security requirements.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "uniqueId",
          description:
            "The passport uniqueId (not the email or phone number) which password would be applied to. Don't confuse with value.",
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
          name: "changed",
          type: "bool",
        },
      ],
    },
  };
}
/**
 * The base class definition for changePasswordActionReq
 **/
export class ChangePasswordActionReq {
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
  /**
   * The passport uniqueId (not the email or phone number) which password would be applied to. Don't confuse with value.
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * The passport uniqueId (not the email or phone number) which password would be applied to. Don't confuse with value.
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * The passport uniqueId (not the email or phone number) which password would be applied to. Don't confuse with value.
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
    const d = data as Partial<ChangePasswordActionReq>;
    if (d.password !== undefined) {
      this.password = d.password;
    }
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
      password: this.#password,
      uniqueId: this.#uniqueId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      password: "password",
      uniqueId: "uniqueId",
    };
  }
  /**
   * Creates an instance of ChangePasswordActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ChangePasswordActionReqType) {
    return new ChangePasswordActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of ChangePasswordActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ChangePasswordActionReqType>) {
    return new ChangePasswordActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ChangePasswordActionReqType>,
  ): InstanceType<typeof ChangePasswordActionReq> {
    return new ChangePasswordActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ChangePasswordActionReq> {
    return new ChangePasswordActionReq(this.toJSON());
  }
}
export abstract class ChangePasswordActionReqFactory {
  abstract create(data: unknown): ChangePasswordActionReq;
}
/**
 * The base type definition for changePasswordActionReq
 **/
export type ChangePasswordActionReqType = {
  /**
   * New password meeting the security requirements.
   * @type {string}
   **/
  password: string;
  /**
   * The passport uniqueId (not the email or phone number) which password would be applied to. Don't confuse with value.
   * @type {string}
   **/
  uniqueId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ChangePasswordActionReqType {}
/**
 * The base class definition for changePasswordActionRes
 **/
export class ChangePasswordActionRes {
  /**
   *
   * @type {boolean}
   **/
  #changed!: boolean;
  /**
   *
   * @returns {boolean}
   **/
  get changed() {
    return this.#changed;
  }
  /**
   *
   * @type {boolean}
   **/
  set changed(value: boolean) {
    this.#changed = Boolean(value);
  }
  setChanged(value: boolean) {
    this.changed = value;
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
    const d = data as Partial<ChangePasswordActionRes>;
    if (d.changed !== undefined) {
      this.changed = d.changed;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      changed: this.#changed,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      changed: "changed",
    };
  }
  /**
   * Creates an instance of ChangePasswordActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ChangePasswordActionResType) {
    return new ChangePasswordActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of ChangePasswordActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ChangePasswordActionResType>) {
    return new ChangePasswordActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ChangePasswordActionResType>,
  ): InstanceType<typeof ChangePasswordActionRes> {
    return new ChangePasswordActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ChangePasswordActionRes> {
    return new ChangePasswordActionRes(this.toJSON());
  }
}
export abstract class ChangePasswordActionResFactory {
  abstract create(data: unknown): ChangePasswordActionRes;
}
/**
 * The base type definition for changePasswordActionRes
 **/
export type ChangePasswordActionResType = {
  /**
   *
   * @type {boolean}
   **/
  changed: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ChangePasswordActionResType {}
