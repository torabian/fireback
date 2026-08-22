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
import { withPrefix } from "@fireback/js-remote-ctx/common/withPrefix";
/**
 * Action to communicate with the action topup
 */
export type TopupActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type TopupActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  TopupActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => TopupActionRes;
  }>;
export const useTopupAction = (options?: TopupActionMutationOptions) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: TopupActionReq) => {
    setCompleteState(false);
    return TopupAction.Fetch(
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
 * TopupAction
 */
export class TopupAction {
  //
  static URL = "/wallet/topup";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(TopupAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<TopupActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<TopupActionRes, TopupActionReq, unknown>(
      overrideUrl ?? TopupAction.NewUrl(qs),
      {
        method: TopupAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<TopupActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => TopupActionRes) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new TopupActionRes(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new TopupActionRes(item));
    const res = await TopupAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (item) => (creatorFn ? creatorFn(item) : item),
      onMessage,
      init?.signal,
    );
  };
  static Definition = {
    name: "topup",
    cliShort: "topup",
    url: "/wallet/topup",
    method: "post",
    description:
      "Starts a topup of walletId through gatewayCode: creates a pending walletPaymentAttempt and asks the gateway adapter to initiate payment, returning whatever the caller needs to complete it (redirect URL and/or a client secret, gateway-dependent). idempotencyKey makes retrying a timed-out call safe - it will not create a second attempt at the gateway.",
    in: {
      fields: [
        {
          name: "walletId",
          description: "Unique id of the wallet to top up.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "gatewayCode",
          description: "Code of the walletGateway to pay through.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "amount",
          description: "Amount to top up, as a positive minor-units string.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "idempotencyKey",
          description: "Makes this topup-initiation safe to retry.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "returnUrl",
          description:
            "Where to send the caller's browser back to once a redirect-based gateway (Przelewy24, ZarinPal, BLIK) completes the payment. Not needed for gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow) - effectively required for the others.",
          type: "string?",
        },
      ],
    },
    out: {
      fields: [
        {
          name: "attempt",
          description: "The created payment attempt.",
          type: "object",
          fields: [
            {
              name: "uniqueId",
              type: "string",
            },
            {
              name: "status",
              type: "string",
            },
            {
              name: "gatewayCode",
              type: "string",
            },
          ],
        },
        {
          name: "redirectUrl",
          description: "URL to send the owner to, for gateways that need one.",
          type: "string?",
        },
        {
          name: "clientSecret",
          description:
            "Client-side secret/token, for gateways that need one instead.",
          type: "string?",
        },
      ],
    },
  };
}
/**
 * The base class definition for topupActionReq
 **/
export class TopupActionReq {
  /**
   * Unique id of the wallet to top up.
   * @type {string}
   **/
  #walletId: string = "";
  /**
   * Unique id of the wallet to top up.
   * @returns {string}
   **/
  get walletId() {
    return this.#walletId;
  }
  /**
   * Unique id of the wallet to top up.
   * @type {string}
   **/
  set walletId(value: string) {
    this.#walletId = String(value);
  }
  setWalletId(value: string) {
    this.walletId = value;
    return this;
  }
  /**
   * Code of the walletGateway to pay through.
   * @type {string}
   **/
  #gatewayCode: string = "";
  /**
   * Code of the walletGateway to pay through.
   * @returns {string}
   **/
  get gatewayCode() {
    return this.#gatewayCode;
  }
  /**
   * Code of the walletGateway to pay through.
   * @type {string}
   **/
  set gatewayCode(value: string) {
    this.#gatewayCode = String(value);
  }
  setGatewayCode(value: string) {
    this.gatewayCode = value;
    return this;
  }
  /**
   * Amount to top up, as a positive minor-units string.
   * @type {string}
   **/
  #amount: string = "";
  /**
   * Amount to top up, as a positive minor-units string.
   * @returns {string}
   **/
  get amount() {
    return this.#amount;
  }
  /**
   * Amount to top up, as a positive minor-units string.
   * @type {string}
   **/
  set amount(value: string) {
    this.#amount = String(value);
  }
  setAmount(value: string) {
    this.amount = value;
    return this;
  }
  /**
   * Makes this topup-initiation safe to retry.
   * @type {string}
   **/
  #idempotencyKey: string = "";
  /**
   * Makes this topup-initiation safe to retry.
   * @returns {string}
   **/
  get idempotencyKey() {
    return this.#idempotencyKey;
  }
  /**
   * Makes this topup-initiation safe to retry.
   * @type {string}
   **/
  set idempotencyKey(value: string) {
    this.#idempotencyKey = String(value);
  }
  setIdempotencyKey(value: string) {
    this.idempotencyKey = value;
    return this;
  }
  /**
   * Where to send the caller's browser back to once a redirect-based gateway (Przelewy24, ZarinPal, BLIK) completes the payment. Not needed for gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow) - effectively required for the others.
   * @type {string}
   **/
  #returnUrl?: string | null = undefined;
  /**
   * Where to send the caller's browser back to once a redirect-based gateway (Przelewy24, ZarinPal, BLIK) completes the payment. Not needed for gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow) - effectively required for the others.
   * @returns {string}
   **/
  get returnUrl() {
    return this.#returnUrl;
  }
  /**
   * Where to send the caller's browser back to once a redirect-based gateway (Przelewy24, ZarinPal, BLIK) completes the payment. Not needed for gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow) - effectively required for the others.
   * @type {string}
   **/
  set returnUrl(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#returnUrl = correctType ? value : String(value);
  }
  setReturnUrl(value: string | null | undefined) {
    this.returnUrl = value;
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
    const d = data as Partial<TopupActionReq>;
    if (d.walletId !== undefined) {
      this.walletId = d.walletId;
    }
    if (d.gatewayCode !== undefined) {
      this.gatewayCode = d.gatewayCode;
    }
    if (d.amount !== undefined) {
      this.amount = d.amount;
    }
    if (d.idempotencyKey !== undefined) {
      this.idempotencyKey = d.idempotencyKey;
    }
    if (d.returnUrl !== undefined) {
      this.returnUrl = d.returnUrl;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      walletId: this.#walletId,
      gatewayCode: this.#gatewayCode,
      amount: this.#amount,
      idempotencyKey: this.#idempotencyKey,
      returnUrl: this.#returnUrl,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      walletId: "walletId",
      gatewayCode: "gatewayCode",
      amount: "amount",
      idempotencyKey: "idempotencyKey",
      returnUrl: "returnUrl",
    };
  }
  /**
   * Creates an instance of TopupActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: TopupActionReqType) {
    return new TopupActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of TopupActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<TopupActionReqType>) {
    return new TopupActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<TopupActionReqType>,
  ): InstanceType<typeof TopupActionReq> {
    return new TopupActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof TopupActionReq> {
    return new TopupActionReq(this.toJSON());
  }
}
export abstract class TopupActionReqFactory {
  abstract create(data: unknown): TopupActionReq;
}
/**
 * The base type definition for topupActionReq
 **/
export type TopupActionReqType = {
  /**
   * Unique id of the wallet to top up.
   * @type {string}
   **/
  walletId: string;
  /**
   * Code of the walletGateway to pay through.
   * @type {string}
   **/
  gatewayCode: string;
  /**
   * Amount to top up, as a positive minor-units string.
   * @type {string}
   **/
  amount: string;
  /**
   * Makes this topup-initiation safe to retry.
   * @type {string}
   **/
  idempotencyKey: string;
  /**
   * Where to send the caller's browser back to once a redirect-based gateway (Przelewy24, ZarinPal, BLIK) completes the payment. Not needed for gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow) - effectively required for the others.
   * @type {string}
   **/
  returnUrl?: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace TopupActionReqType {}
/**
 * The base class definition for topupActionRes
 **/
export class TopupActionRes {
  /**
   * The created payment attempt.
   * @type {TopupActionRes.Attempt}
   **/
  #attempt!: InstanceType<typeof TopupActionRes.Attempt>;
  /**
   * The created payment attempt.
   * @returns {TopupActionRes.Attempt}
   **/
  get attempt() {
    return this.#attempt;
  }
  /**
   * The created payment attempt.
   * @type {TopupActionRes.Attempt}
   **/
  set attempt(value: InstanceType<typeof TopupActionRes.Attempt>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof TopupActionRes.Attempt) {
      this.#attempt = value;
    } else {
      this.#attempt = new TopupActionRes.Attempt(value);
    }
  }
  setAttempt(value: InstanceType<typeof TopupActionRes.Attempt>) {
    this.attempt = value;
    return this;
  }
  /**
   * URL to send the owner to, for gateways that need one.
   * @type {string}
   **/
  #redirectUrl?: string | null = undefined;
  /**
   * URL to send the owner to, for gateways that need one.
   * @returns {string}
   **/
  get redirectUrl() {
    return this.#redirectUrl;
  }
  /**
   * URL to send the owner to, for gateways that need one.
   * @type {string}
   **/
  set redirectUrl(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#redirectUrl = correctType ? value : String(value);
  }
  setRedirectUrl(value: string | null | undefined) {
    this.redirectUrl = value;
    return this;
  }
  /**
   * Client-side secret/token, for gateways that need one instead.
   * @type {string}
   **/
  #clientSecret?: string | null = undefined;
  /**
   * Client-side secret/token, for gateways that need one instead.
   * @returns {string}
   **/
  get clientSecret() {
    return this.#clientSecret;
  }
  /**
   * Client-side secret/token, for gateways that need one instead.
   * @type {string}
   **/
  set clientSecret(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#clientSecret = correctType ? value : String(value);
  }
  setClientSecret(value: string | null | undefined) {
    this.clientSecret = value;
    return this;
  }
  /**
   * The base class definition for attempt
   **/
  static Attempt = class Attempt {
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
    #status: string = "";
    /**
     *
     * @returns {string}
     **/
    get status() {
      return this.#status;
    }
    /**
     *
     * @type {string}
     **/
    set status(value: string) {
      this.#status = String(value);
    }
    setStatus(value: string) {
      this.status = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #gatewayCode: string = "";
    /**
     *
     * @returns {string}
     **/
    get gatewayCode() {
      return this.#gatewayCode;
    }
    /**
     *
     * @type {string}
     **/
    set gatewayCode(value: string) {
      this.#gatewayCode = String(value);
    }
    setGatewayCode(value: string) {
      this.gatewayCode = value;
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
      const d = data as Partial<Attempt>;
      if (d.uniqueId !== undefined) {
        this.uniqueId = d.uniqueId;
      }
      if (d.status !== undefined) {
        this.status = d.status;
      }
      if (d.gatewayCode !== undefined) {
        this.gatewayCode = d.gatewayCode;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        uniqueId: this.#uniqueId,
        status: this.#status,
        gatewayCode: this.#gatewayCode,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        uniqueId: "uniqueId",
        status: "status",
        gatewayCode: "gatewayCode",
      };
    }
    /**
     * Creates an instance of TopupActionRes.Attempt, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: TopupActionResType.AttemptType) {
      return new TopupActionRes.Attempt(possibleDtoObject);
    }
    /**
     * Creates an instance of TopupActionRes.Attempt, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject: PartialDeep<TopupActionResType.AttemptType>) {
      return new TopupActionRes.Attempt(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<TopupActionResType.AttemptType>,
    ): InstanceType<typeof TopupActionRes.Attempt> {
      return new TopupActionRes.Attempt({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof TopupActionRes.Attempt> {
      return new TopupActionRes.Attempt(this.toJSON());
    }
  };
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
    const d = data as Partial<TopupActionRes>;
    if (d.attempt !== undefined) {
      this.attempt = d.attempt;
    }
    if (d.redirectUrl !== undefined) {
      this.redirectUrl = d.redirectUrl;
    }
    if (d.clientSecret !== undefined) {
      this.clientSecret = d.clientSecret;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<TopupActionRes>;
    if (!(d.attempt instanceof TopupActionRes.Attempt)) {
      this.attempt = new TopupActionRes.Attempt(d.attempt || {});
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      attempt: this.#attempt,
      redirectUrl: this.#redirectUrl,
      clientSecret: this.#clientSecret,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      attempt$: "attempt",
      get attempt() {
        return withPrefix("attempt", TopupActionRes.Attempt.Fields);
      },
      redirectUrl: "redirectUrl",
      clientSecret: "clientSecret",
    };
  }
  /**
   * Creates an instance of TopupActionRes, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: TopupActionResType) {
    return new TopupActionRes(possibleDtoObject);
  }
  /**
   * Creates an instance of TopupActionRes, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<TopupActionResType>) {
    return new TopupActionRes(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<TopupActionResType>,
  ): InstanceType<typeof TopupActionRes> {
    return new TopupActionRes({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof TopupActionRes> {
    return new TopupActionRes(this.toJSON());
  }
}
export abstract class TopupActionResFactory {
  abstract create(data: unknown): TopupActionRes;
}
/**
 * The base type definition for topupActionRes
 **/
export type TopupActionResType = {
  /**
   * The created payment attempt.
   * @type {TopupActionResType.AttemptType}
   **/
  attempt: TopupActionResType.AttemptType;
  /**
   * URL to send the owner to, for gateways that need one.
   * @type {string}
   **/
  redirectUrl?: string;
  /**
   * Client-side secret/token, for gateways that need one instead.
   * @type {string}
   **/
  clientSecret?: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace TopupActionResType {
  /**
   * The base type definition for attemptType
   **/
  export type AttemptType = {
    /**
     *
     * @type {string}
     **/
    uniqueId: string;
    /**
     *
     * @type {string}
     **/
    status: string;
    /**
     *
     * @type {string}
     **/
    gatewayCode: string;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace AttemptType {}
}
