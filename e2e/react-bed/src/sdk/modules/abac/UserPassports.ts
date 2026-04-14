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
import {
  type UseMutationOptions,
  type UseQueryOptions,
  useMutation,
  useQuery,
} from "react-query";
import { useFetchxContext } from "../../sdk/react/useFetchx";
import { useState } from "react";
/**
 * Action to communicate with the action UserPassports
 */
export type UserPassportsActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type UserPassportsActionQueryOptions = Omit<
  UseQueryOptions<
    unknown,
    unknown,
    GResponse<UserPassportsActionRes>,
    unknown[]
  >,
  "queryKey"
> &
  UserPassportsActionOptions &
  Partial<{
    creatorFn: (item: unknown) => UserPassportsActionRes;
  }> & {
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
    ctx?: FetchxContext;
  };
export const useUserPassportsActionQuery = (
  options: UserPassportsActionQueryOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = () => {
    setCompleteState(false);
    return UserPassportsAction.Fetch(
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
    queryKey: [UserPassportsAction.NewUrl(options?.qs)],
    queryFn: fn,
    ...(options || {}),
  });
  return {
    ...result,
    isCompleted,
    response,
  };
};
export type UserPassportsActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  UserPassportsActionOptions & {
    ctx?: FetchxContext;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => UserPassportsActionRes;
  }>;
export const useUserPassportsAction = (
  options?: UserPassportsActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: unknown) => {
    setCompleteState(false);
    return UserPassportsAction.Fetch(
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
 * UserPassportsAction
 */
export class UserPassportsAction {
  //
  static URL = "/user/passports";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(UserPassportsAction.URL, undefined, qs);
  static Method = "get";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext,
    init?: TypedRequestInit<unknown, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<UserPassportsActionRes>, unknown, unknown>(
      overrideUrl ?? UserPassportsAction.NewUrl(qs),
      {
        method: UserPassportsAction.Method,
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
      creatorFn?: ((item: unknown) => UserPassportsActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new UserPassportsActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new UserPassportsActionRes(item));
    const res = await UserPassportsAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<UserPassportsActionRes>();
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
    name: "UserPassports",
    url: "/user/passports",
    method: "get",
    description: "Returns list of passports belongs to an specific user.",
    out: {
      envelope: "GResponse",
      fields: [
        {
          name: "value",
          description:
            "The passport value, such as email address or phone number",
          type: "string",
        },
        {
          name: "uniqueId",
          description:
            "Unique identifier of the passport to operate some action on top of it",
          type: "string",
        },
        {
          name: "type",
          description: "The type of the passport, such as email, phone number",
          type: "string",
        },
        {
          name: "totpConfirmed",
          description:
            "Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.",
          type: "bool",
        },
      ],
    },
  };
}
/**
 * The base class definition for userPassportsActionRes
 **/
export class UserPassportsActionRes {
  /**
   * The passport value, such as email address or phone number
   * @type {string}
   **/
  #value: string = "";
  /**
   * The passport value, such as email address or phone number
   * @returns {string}
   **/
  get value() {
    return this.#value;
  }
  /**
   * The passport value, such as email address or phone number
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
   * Unique identifier of the passport to operate some action on top of it
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * Unique identifier of the passport to operate some action on top of it
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * Unique identifier of the passport to operate some action on top of it
   * @type {string}
   **/
  set uniqueId(value: string) {
    this.#uniqueId = String(value);
  }
  setUniqueId(value: string) {
    this.uniqueId = value;
    return this;
  }
  /**
   * The type of the passport, such as email, phone number
   * @type {string}
   **/
  #type: string = "";
  /**
   * The type of the passport, such as email, phone number
   * @returns {string}
   **/
  get type() {
    return this.#type;
  }
  /**
   * The type of the passport, such as email, phone number
   * @type {string}
   **/
  set type(value: string) {
    this.#type = String(value);
  }
  setType(value: string) {
    this.type = value;
    return this;
  }
  /**
   * Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
   * @type {boolean}
   **/
  #totpConfirmed!: boolean;
  /**
   * Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
   * @returns {boolean}
   **/
  get totpConfirmed() {
    return this.#totpConfirmed;
  }
  /**
   * Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
   * @type {boolean}
   **/
  set totpConfirmed(value: boolean) {
    this.#totpConfirmed = Boolean(value);
  }
  setTotpConfirmed(value: boolean) {
    this.totpConfirmed = value;
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
    const d = data as Partial<UserPassportsActionRes>;
    if (d.value !== undefined) {
      this.value = d.value;
    }
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.type !== undefined) {
      this.type = d.type;
    }
    if (d.totpConfirmed !== undefined) {
      this.totpConfirmed = d.totpConfirmed;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      value: this.#value,
      uniqueId: this.#uniqueId,
      type: this.#type,
      totpConfirmed: this.#totpConfirmed,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      value: "value",
      uniqueId: "uniqueId",
      type: "type",
      totpConfirmed: "totpConfirmed",
    };
  }
  /**
   * Creates an instance of UserPassportsActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: UserPassportsActionResType) {
    return new UserPassportsActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of UserPassportsActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<UserPassportsActionResType>) {
    return new UserPassportsActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<UserPassportsActionResType>,
  ): InstanceType<typeof UserPassportsActionRes> {
    return new UserPassportsActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof UserPassportsActionRes> {
    return new UserPassportsActionRes(this.toJSON());
  }
}
export abstract class UserPassportsActionResFactory {
  abstract create(data: unknown): UserPassportsActionRes;
}
/**
 * The base type definition for userPassportsActionRes
 **/
export type UserPassportsActionResType = {
  /**
   * The passport value, such as email address or phone number
   * @type {string}
   **/
  value: string;
  /**
   * Unique identifier of the passport to operate some action on top of it
   * @type {string}
   **/
  uniqueId: string;
  /**
   * The type of the passport, such as email, phone number
   * @type {string}
   **/
  type: string;
  /**
   * Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
   * @type {boolean}
   **/
  totpConfirmed: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace UserPassportsActionResType {}
