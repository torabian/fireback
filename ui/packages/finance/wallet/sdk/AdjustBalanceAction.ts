import { GResponse } from "@fireback/js-remote-ctx/envelopes/index";
import { WalletTransactionDto } from "./WalletTransactionDto";
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
 * Action to communicate with the action adjustBalance
 */
export type AdjustBalanceActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type AdjustBalanceActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  AdjustBalanceActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletTransactionDto;
  }>;
export const useAdjustBalanceAction = (
  options?: AdjustBalanceActionMutationOptions,
) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: AdjustBalanceActionReq) => {
    setCompleteState(false);
    return AdjustBalanceAction.Fetch(
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
 * AdjustBalanceAction
 */
export class AdjustBalanceAction {
  //
  static URL = "/wallet/adjust";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(AdjustBalanceAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<AdjustBalanceActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<
      GResponse<WalletTransactionDto>,
      AdjustBalanceActionReq,
      unknown
    >(
      overrideUrl ?? AdjustBalanceAction.NewUrl(qs),
      {
        method: AdjustBalanceAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<AdjustBalanceActionReq, unknown>,
    {
      creatorFn,
      qs,
      ctx,
      onMessage,
      overrideUrl,
    }: {
      creatorFn?: ((item: unknown) => WalletTransactionDto) | undefined;
      qs?: URLSearchParams;
      ctx?: FetchxContext | null;
      onMessage?: (ev: MessageEvent) => void;
      overrideUrl?: string;
    } = {
      creatorFn: (item) => new WalletTransactionDto(item),
    },
  ) => {
    creatorFn = creatorFn || ((item) => new WalletTransactionDto(item));
    const res = await AdjustBalanceAction.Fetch$(qs, ctx, init, overrideUrl);
    return handleFetchResponse(
      res,
      (data) => {
        const resp = new GResponse<WalletTransactionDto>();
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
    name: "adjustBalance",
    cliShort: "adjust",
    url: "/wallet/adjust",
    method: "post",
    description:
      'Root-only manual balance correction (support/ops use). Runs through the same locked-transaction path as purchase, and always requires a note for audit. reason on the resulting ledger entry is always "adjustment".',
    in: {
      fields: [
        {
          name: "walletId",
          description: "Unique id of the wallet to adjust.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "direction",
          description:
            "Whether this adjustment increases or decreases the balance.",
          type: "string",
          tags: {
            validate: "required,oneof=credit debit",
          },
        },
        {
          name: "amount",
          description:
            "Magnitude of the adjustment, as a positive minor-units string.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "note",
          description:
            "Required human explanation of why this adjustment was made.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "idempotencyKey",
          description:
            "Makes this adjustment safe to retry without double-applying.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
      ],
    },
    out: {
      envelope: "GResponse",
      dto: "WalletTransactionDto",
    },
  };
}
/**
 * The base class definition for adjustBalanceActionReq
 **/
export class AdjustBalanceActionReq {
  /**
   * Unique id of the wallet to adjust.
   * @type {string}
   **/
  #walletId: string = "";
  /**
   * Unique id of the wallet to adjust.
   * @returns {string}
   **/
  get walletId() {
    return this.#walletId;
  }
  /**
   * Unique id of the wallet to adjust.
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
   * Whether this adjustment increases or decreases the balance.
   * @type {string}
   **/
  #direction: string = "";
  /**
   * Whether this adjustment increases or decreases the balance.
   * @returns {string}
   **/
  get direction() {
    return this.#direction;
  }
  /**
   * Whether this adjustment increases or decreases the balance.
   * @type {string}
   **/
  set direction(value: string) {
    this.#direction = String(value);
  }
  setDirection(value: string) {
    this.direction = value;
    return this;
  }
  /**
   * Magnitude of the adjustment, as a positive minor-units string.
   * @type {string}
   **/
  #amount: string = "";
  /**
   * Magnitude of the adjustment, as a positive minor-units string.
   * @returns {string}
   **/
  get amount() {
    return this.#amount;
  }
  /**
   * Magnitude of the adjustment, as a positive minor-units string.
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
   * Required human explanation of why this adjustment was made.
   * @type {string}
   **/
  #note: string = "";
  /**
   * Required human explanation of why this adjustment was made.
   * @returns {string}
   **/
  get note() {
    return this.#note;
  }
  /**
   * Required human explanation of why this adjustment was made.
   * @type {string}
   **/
  set note(value: string) {
    this.#note = String(value);
  }
  setNote(value: string) {
    this.note = value;
    return this;
  }
  /**
   * Makes this adjustment safe to retry without double-applying.
   * @type {string}
   **/
  #idempotencyKey: string = "";
  /**
   * Makes this adjustment safe to retry without double-applying.
   * @returns {string}
   **/
  get idempotencyKey() {
    return this.#idempotencyKey;
  }
  /**
   * Makes this adjustment safe to retry without double-applying.
   * @type {string}
   **/
  set idempotencyKey(value: string) {
    this.#idempotencyKey = String(value);
  }
  setIdempotencyKey(value: string) {
    this.idempotencyKey = value;
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
    const d = data as Partial<AdjustBalanceActionReq>;
    if (d.walletId !== undefined) {
      this.walletId = d.walletId;
    }
    if (d.direction !== undefined) {
      this.direction = d.direction;
    }
    if (d.amount !== undefined) {
      this.amount = d.amount;
    }
    if (d.note !== undefined) {
      this.note = d.note;
    }
    if (d.idempotencyKey !== undefined) {
      this.idempotencyKey = d.idempotencyKey;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      walletId: this.#walletId,
      direction: this.#direction,
      amount: this.#amount,
      note: this.#note,
      idempotencyKey: this.#idempotencyKey,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      walletId: "walletId",
      direction: "direction",
      amount: "amount",
      note: "note",
      idempotencyKey: "idempotencyKey",
    };
  }
  /**
   * Creates an instance of AdjustBalanceActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: AdjustBalanceActionReqType) {
    return new AdjustBalanceActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of AdjustBalanceActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<AdjustBalanceActionReqType>) {
    return new AdjustBalanceActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<AdjustBalanceActionReqType>,
  ): InstanceType<typeof AdjustBalanceActionReq> {
    return new AdjustBalanceActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof AdjustBalanceActionReq> {
    return new AdjustBalanceActionReq(this.toJSON());
  }
}
export abstract class AdjustBalanceActionReqFactory {
  abstract create(data: unknown): AdjustBalanceActionReq;
}
/**
 * The base type definition for adjustBalanceActionReq
 **/
export type AdjustBalanceActionReqType = {
  /**
   * Unique id of the wallet to adjust.
   * @type {string}
   **/
  walletId: string;
  /**
   * Whether this adjustment increases or decreases the balance.
   * @type {string}
   **/
  direction: string;
  /**
   * Magnitude of the adjustment, as a positive minor-units string.
   * @type {string}
   **/
  amount: string;
  /**
   * Required human explanation of why this adjustment was made.
   * @type {string}
   **/
  note: string;
  /**
   * Makes this adjustment safe to retry without double-applying.
   * @type {string}
   **/
  idempotencyKey: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AdjustBalanceActionReqType {}
