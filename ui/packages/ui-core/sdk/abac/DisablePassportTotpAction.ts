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
 * Action to communicate with the action DisablePassportTotp
 */
export type DisablePassportTotpActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type DisablePassportTotpActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  DisablePassportTotpActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => DisablePassportTotpActionRes;
  }>;
export const useDisablePassportTotpAction = (
  options?: DisablePassportTotpActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: DisablePassportTotpActionReq) => {
    setCompleteState(false);
    return DisablePassportTotpAction.Fetch(
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
 * DisablePassportTotpAction
 */
export class DisablePassportTotpAction {
  //
  static URL = "/passport/totp/disable";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(DisablePassportTotpAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<DisablePassportTotpActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<DisablePassportTotpActionRes>,
      DisablePassportTotpActionReq,
      unknown
    >(
      overrideUrl ?? DisablePassportTotpAction.NewUrl(qs),
      {
        method: DisablePassportTotpAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<DisablePassportTotpActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => DisablePassportTotpActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new DisablePassportTotpActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new DisablePassportTotpActionRes(item));
    const res = await DisablePassportTotpAction.Fetch$(
      qs,
      ctx,
      init,
      overrideUrl,
    );
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<DisablePassportTotpActionRes>();
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
    name: "DisablePassportTotp",
    cliName: "disable-totp",
    url: "/passport/totp/disable",
    method: "post",
    description:
      'Root-only. Force-disables TOTP (2FA) on a passport by clearing its secret and confirmation flag - for unblocking a user who lost their authenticator device. There is no matching "enable" action, since turning TOTP on requires the user\'s own authenticator app to scan a fresh secret (see ConfirmClassicPassportTotpAction), which an administrator cannot do on their behalf.',
    in: {
      fields: [
        {
          name: "uniqueId",
          description: "The passport uniqueId to disable TOTP for.",
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
          name: "disabled",
          type: "bool",
        },
      ],
    },
  };
}
/**
 * The base class definition for disablePassportTotpActionReq
 **/
export class DisablePassportTotpActionReq {
  /**
   * The passport uniqueId to disable TOTP for.
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * The passport uniqueId to disable TOTP for.
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * The passport uniqueId to disable TOTP for.
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
    const d = data as Partial<DisablePassportTotpActionReq>;
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
      uniqueId: this.#uniqueId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
    };
  }
  /**
   * Creates an instance of DisablePassportTotpActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: DisablePassportTotpActionReqType) {
    return new DisablePassportTotpActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of DisablePassportTotpActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<DisablePassportTotpActionReqType>) {
    return new DisablePassportTotpActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<DisablePassportTotpActionReqType>,
  ): InstanceType<typeof DisablePassportTotpActionReq> {
    return new DisablePassportTotpActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof DisablePassportTotpActionReq> {
    return new DisablePassportTotpActionReq(this.toJSON());
  }
}
export abstract class DisablePassportTotpActionReqFactory {
  abstract create(data: unknown): DisablePassportTotpActionReq;
}
/**
 * The base type definition for disablePassportTotpActionReq
 **/
export type DisablePassportTotpActionReqType = {
  /**
   * The passport uniqueId to disable TOTP for.
   * @type {string}
   **/
  uniqueId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DisablePassportTotpActionReqType {}
/**
 * The base class definition for disablePassportTotpActionRes
 **/
export class DisablePassportTotpActionRes {
  /**
   *
   * @type {boolean}
   **/
  #disabled!: boolean;
  /**
   *
   * @returns {boolean}
   **/
  get disabled() {
    return this.#disabled;
  }
  /**
   *
   * @type {boolean}
   **/
  set disabled(value: boolean) {
    this.#disabled = Boolean(value);
  }
  setDisabled(value: boolean) {
    this.disabled = value;
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
    const d = data as Partial<DisablePassportTotpActionRes>;
    if (d.disabled !== undefined) {
      this.disabled = d.disabled;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      disabled: this.#disabled,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      disabled: "disabled",
    };
  }
  /**
   * Creates an instance of DisablePassportTotpActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: DisablePassportTotpActionResType) {
    return new DisablePassportTotpActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of DisablePassportTotpActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<DisablePassportTotpActionResType>) {
    return new DisablePassportTotpActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<DisablePassportTotpActionResType>,
  ): InstanceType<typeof DisablePassportTotpActionRes> {
    return new DisablePassportTotpActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof DisablePassportTotpActionRes> {
    return new DisablePassportTotpActionRes(this.toJSON());
  }
}
export abstract class DisablePassportTotpActionResFactory {
  abstract create(data: unknown): DisablePassportTotpActionRes;
}
/**
 * The base type definition for disablePassportTotpActionRes
 **/
export type DisablePassportTotpActionResType = {
  /**
   *
   * @type {boolean}
   **/
  disabled: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DisablePassportTotpActionResType {}
