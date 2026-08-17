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
 * Action to communicate with the action SetPassportPassword
 */
export type SetPassportPasswordActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type SetPassportPasswordActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  SetPassportPasswordActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => SetPassportPasswordActionRes;
  }>;
export const useSetPassportPasswordAction = (
  options?: SetPassportPasswordActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: SetPassportPasswordActionReq) => {
    setCompleteState(false);
    return SetPassportPasswordAction.Fetch(
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
 * SetPassportPasswordAction
 */
export class SetPassportPasswordAction {
  //
  static URL = "/passport/set-password";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(SetPassportPasswordAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<SetPassportPasswordActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<SetPassportPasswordActionRes>,
      SetPassportPasswordActionReq,
      unknown
    >(
      overrideUrl ?? SetPassportPasswordAction.NewUrl(qs),
      {
        method: SetPassportPasswordAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<SetPassportPasswordActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => SetPassportPasswordActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new SetPassportPasswordActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new SetPassportPasswordActionRes(item));
    const res = await SetPassportPasswordAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<SetPassportPasswordActionRes>();
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
    name: "SetPassportPassword",
    cliName: "set-password",
    url: "/passport/set-password",
    method: "post",
    description:
      "Root-only. Directly sets a new password on any passport, without knowing (or being) the current one - unlike ChangePassword, which only ever touches the calling user's own passport. Meant for an administrator resetting a locked-out user's password on their behalf.",
    in: {
      fields: [
        {
          name: "uniqueId",
          description: "The passport uniqueId whose password will be replaced.",
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
          name: "changed",
          type: "bool",
        },
      ],
    },
  };
}
/**
 * The base class definition for setPassportPasswordActionReq
 **/
export class SetPassportPasswordActionReq {
  /**
   * The passport uniqueId whose password will be replaced.
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * The passport uniqueId whose password will be replaced.
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * The passport uniqueId whose password will be replaced.
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
    const d = data as Partial<SetPassportPasswordActionReq>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
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
      uniqueId: this.#uniqueId,
      password: this.#password,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      password: "password",
    };
  }
  /**
   * Creates an instance of SetPassportPasswordActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: SetPassportPasswordActionReqType) {
    return new SetPassportPasswordActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of SetPassportPasswordActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<SetPassportPasswordActionReqType>) {
    return new SetPassportPasswordActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<SetPassportPasswordActionReqType>,
  ): InstanceType<typeof SetPassportPasswordActionReq> {
    return new SetPassportPasswordActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof SetPassportPasswordActionReq> {
    return new SetPassportPasswordActionReq(this.toJSON());
  }
}
export abstract class SetPassportPasswordActionReqFactory {
  abstract create(data: unknown): SetPassportPasswordActionReq;
}
/**
 * The base type definition for setPassportPasswordActionReq
 **/
export type SetPassportPasswordActionReqType = {
  /**
   * The passport uniqueId whose password will be replaced.
   * @type {string}
   **/
  uniqueId: string;
  /**
   * New password meeting the security requirements.
   * @type {string}
   **/
  password: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SetPassportPasswordActionReqType {}
/**
 * The base class definition for setPassportPasswordActionRes
 **/
export class SetPassportPasswordActionRes {
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
    const d = data as Partial<SetPassportPasswordActionRes>;
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
   * Creates an instance of SetPassportPasswordActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: SetPassportPasswordActionResType) {
    return new SetPassportPasswordActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of SetPassportPasswordActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<SetPassportPasswordActionResType>) {
    return new SetPassportPasswordActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<SetPassportPasswordActionResType>,
  ): InstanceType<typeof SetPassportPasswordActionRes> {
    return new SetPassportPasswordActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof SetPassportPasswordActionRes> {
    return new SetPassportPasswordActionRes(this.toJSON());
  }
}
export abstract class SetPassportPasswordActionResFactory {
  abstract create(data: unknown): SetPassportPasswordActionRes;
}
/**
 * The base type definition for setPassportPasswordActionRes
 **/
export type SetPassportPasswordActionResType = {
  /**
   *
   * @type {boolean}
   **/
  changed: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace SetPassportPasswordActionResType {}
