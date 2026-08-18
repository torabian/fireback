import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
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
 * Action to communicate with the action CreatePassportForUser
 */
export type CreatePassportForUserActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type CreatePassportForUserActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  CreatePassportForUserActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => CreatePassportForUserActionRes;
  }>;
export const useCreatePassportForUserAction = (
  options?: CreatePassportForUserActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: CreatePassportForUserActionReq) => {
    setCompleteState(false);
    return CreatePassportForUserAction.Fetch(
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
 * CreatePassportForUserAction
 */
export class CreatePassportForUserAction {
  //
  static URL = "/passport/create-for-user";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(CreatePassportForUserAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<CreatePassportForUserActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<CreatePassportForUserActionRes>,
      CreatePassportForUserActionReq,
      unknown
    >(
      overrideUrl ?? CreatePassportForUserAction.NewUrl(qs),
      {
        method: CreatePassportForUserAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<CreatePassportForUserActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?:
        | ((item: unknown) => CreatePassportForUserActionRes)
        | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new CreatePassportForUserActionRes(item),
    },
  ) => {
    creatorFn =
      creatorFn || ((item) => new CreatePassportForUserActionRes(item));
    const res = await CreatePassportForUserAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<CreatePassportForUserActionRes>();
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
    name: "CreatePassportForUser",
    cliName: "create-for-user",
    url: "/passport/create-for-user",
    method: "post",
    description:
      "Creates a new email or phone passport for an existing user with a password already set, and marks it confirmed - for an admin manually provisioning a login when the user can't complete the normal signup/verification flow themselves. Root only. The generic 'passport create' endpoint can't do this - PassportEntity's password field is deliberately write-only from JSON (json:\"-\"), by design, everywhere except here and SetPassportPassword.",
    in: {
      fields: [
        {
          name: "userId",
          description:
            "UniqueId of the existing user to create the passport for.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "type",
          description: "The passport method: 'email' or 'phone'.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "value",
          description:
            "The email address or phone number for this passport. Must be globally unique across every passport.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "password",
          description:
            "Plaintext password to set on the new passport (minimum 6 characters) - hashed before it's stored, never returned.",
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
          name: "uniqueId",
          type: "string",
        },
        {
          name: "userId",
          type: "string",
        },
        {
          name: "type",
          type: "string",
        },
        {
          name: "value",
          type: "string",
        },
      ],
    },
  };
}
/**
 * The base class definition for createPassportForUserActionReq
 **/
export class CreatePassportForUserActionReq {
  /**
   * UniqueId of the existing user to create the passport for.
   * @type {string}
   **/
  #userId: string = "";
  /**
   * UniqueId of the existing user to create the passport for.
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * UniqueId of the existing user to create the passport for.
   * @type {string}
   **/
  set userId(value: string) {
    this.#userId = String(value);
  }
  setUserId(value: string) {
    this.userId = value;
    return this;
  }
  /**
   * The passport method: 'email' or 'phone'.
   * @type {string}
   **/
  #type: string = "";
  /**
   * The passport method: 'email' or 'phone'.
   * @returns {string}
   **/
  get type() {
    return this.#type;
  }
  /**
   * The passport method: 'email' or 'phone'.
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
   * The email address or phone number for this passport. Must be globally unique across every passport.
   * @type {string}
   **/
  #value: string = "";
  /**
   * The email address or phone number for this passport. Must be globally unique across every passport.
   * @returns {string}
   **/
  get value() {
    return this.#value;
  }
  /**
   * The email address or phone number for this passport. Must be globally unique across every passport.
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
   * Plaintext password to set on the new passport (minimum 6 characters) - hashed before it's stored, never returned.
   * @type {string}
   **/
  #password: string = "";
  /**
   * Plaintext password to set on the new passport (minimum 6 characters) - hashed before it's stored, never returned.
   * @returns {string}
   **/
  get password() {
    return this.#password;
  }
  /**
   * Plaintext password to set on the new passport (minimum 6 characters) - hashed before it's stored, never returned.
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
    const d = data as Partial<CreatePassportForUserActionReq>;
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.type !== undefined) {
      this.type = d.type;
    }
    if (d.value !== undefined) {
      this.value = d.value;
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
      userId: this.#userId,
      type: this.#type,
      value: this.#value,
      password: this.#password,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      userId: "userId",
      type: "type",
      value: "value",
      password: "password",
    };
  }
  /**
   * Creates an instance of CreatePassportForUserActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CreatePassportForUserActionReqType) {
    return new CreatePassportForUserActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of CreatePassportForUserActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<CreatePassportForUserActionReqType>,
  ) {
    return new CreatePassportForUserActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CreatePassportForUserActionReqType>,
  ): InstanceType<typeof CreatePassportForUserActionReq> {
    return new CreatePassportForUserActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CreatePassportForUserActionReq> {
    return new CreatePassportForUserActionReq(this.toJSON());
  }
}
export abstract class CreatePassportForUserActionReqFactory {
  abstract create(data: unknown): CreatePassportForUserActionReq;
}
/**
 * The base type definition for createPassportForUserActionReq
 **/
export type CreatePassportForUserActionReqType = {
  /**
   * UniqueId of the existing user to create the passport for.
   * @type {string}
   **/
  userId: string;
  /**
   * The passport method: 'email' or 'phone'.
   * @type {string}
   **/
  type: string;
  /**
   * The email address or phone number for this passport. Must be globally unique across every passport.
   * @type {string}
   **/
  value: string;
  /**
   * Plaintext password to set on the new passport (minimum 6 characters) - hashed before it's stored, never returned.
   * @type {string}
   **/
  password: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CreatePassportForUserActionReqType {}
/**
 * The base class definition for createPassportForUserActionRes
 **/
export class CreatePassportForUserActionRes {
  /**
   *
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   *
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   *
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
   *
   * @type {string}
   **/
  #userId: string = "";
  /**
   *
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   *
   * @type {string}
   **/
  set userId(value: string) {
    this.#userId = String(value);
  }
  setUserId(value: string) {
    this.userId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #type: string = "";
  /**
   *
   * @returns {string}
   **/
  get type() {
    return this.#type;
  }
  /**
   *
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
    const d = data as Partial<CreatePassportForUserActionRes>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.type !== undefined) {
      this.type = d.type;
    }
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
      uniqueId: this.#uniqueId,
      userId: this.#userId,
      type: this.#type,
      value: this.#value,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      userId: "userId",
      type: "type",
      value: "value",
    };
  }
  /**
   * Creates an instance of CreatePassportForUserActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CreatePassportForUserActionResType) {
    return new CreatePassportForUserActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of CreatePassportForUserActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<CreatePassportForUserActionResType>,
  ) {
    return new CreatePassportForUserActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CreatePassportForUserActionResType>,
  ): InstanceType<typeof CreatePassportForUserActionRes> {
    return new CreatePassportForUserActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CreatePassportForUserActionRes> {
    return new CreatePassportForUserActionRes(this.toJSON());
  }
}
export abstract class CreatePassportForUserActionResFactory {
  abstract create(data: unknown): CreatePassportForUserActionRes;
}
/**
 * The base type definition for createPassportForUserActionRes
 **/
export type CreatePassportForUserActionResType = {
  /**
   *
   * @type {string}
   **/
  uniqueId: string;
  /**
   *
   * @type {string}
   **/
  userId: string;
  /**
   *
   * @type {string}
   **/
  type: string;
  /**
   *
   * @type {string}
   **/
  value: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CreatePassportForUserActionResType {}
