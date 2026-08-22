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
 * Action to communicate with the action purchase
 */
export type PurchaseActionOptions = {
  queryKey?: unknown[];
  qs?: URLSearchParams;
};
export type PurchaseActionMutationOptions = Omit<
  UseMutationOptions<unknown, unknown, unknown, unknown>,
  "mutationFn"
> &
  PurchaseActionOptions & {
    ctx?: FetchxContext | null;
    onMessage?: (ev: MessageEvent) => void;
    overrideUrl?: string;
    headers?: Headers;
  } & Partial<{
    creatorFn: (item: unknown) => WalletTransactionDto;
  }>;
export const usePurchaseAction = (options?: PurchaseActionMutationOptions) => {
  const globalCtx = useFetchxContext();
  const ctx = options?.ctx ?? globalCtx ?? undefined;
  const [isCompleted, setCompleteState] = useState(false);
  const [response, setResponse] = useState<TypedResponse<unknown>>();
  const fn = (body: PurchaseActionReq) => {
    setCompleteState(false);
    return PurchaseAction.Fetch(
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
 * PurchaseAction
 */
export class PurchaseAction {
  //
  static URL = "/wallet/purchase";
  static NewUrl = (qs?: URLSearchParams) =>
    buildUrl(PurchaseAction.URL, undefined, qs);
  static Method = "POST";
  static Fetch$ = async (
    qs?: URLSearchParams,
    ctx?: FetchxContext | null,
    init?: TypedRequestInit<PurchaseActionReq, unknown>,
    overrideUrl?: string,
  ) => {
    return fetchx<GResponse<WalletTransactionDto>, PurchaseActionReq, unknown>(
      overrideUrl ?? PurchaseAction.NewUrl(qs),
      {
        method: PurchaseAction.Method,
        ...(init || {}),
      },
      ctx,
    );
  };
  static Fetch = async (
    init?: TypedRequestInit<PurchaseActionReq, unknown>,
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
    const res = await PurchaseAction.Fetch$(qs, ctx, init, overrideUrl);
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
    name: "purchase",
    cliShort: "purchase",
    url: "/wallet/purchase",
    method: "post",
    description:
      "Internal/service-only: debits a wallet for a purchase made by another module. This is the HTTP-callable twin of the in-process wallet.Purchase Go function (modules/wallet/Purchase.go) - both share the same locked-transaction implementation. Not exposed to wallet owners directly.",
    in: {
      fields: [
        {
          name: "walletId",
          description: "Unique id of the wallet to debit.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "amount",
          description: "Amount to debit, as a positive minor-units string.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "referenceType",
          description:
            'Free-form name of the calling module/feature, e.g. "course-purchase".',
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "referenceId",
          description: "Id within referenceType this purchase relates to.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "idempotencyKey",
          description:
            "Makes this purchase safe to retry without double-debiting.",
          type: "string",
          tags: {
            validate: "required",
          },
        },
        {
          name: "note",
          description: "Optional human-readable note.",
          type: "string?",
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
 * The base class definition for purchaseActionReq
 **/
export class PurchaseActionReq {
  /**
   * Unique id of the wallet to debit.
   * @type {string}
   **/
  #walletId: string = "";
  /**
   * Unique id of the wallet to debit.
   * @returns {string}
   **/
  get walletId() {
    return this.#walletId;
  }
  /**
   * Unique id of the wallet to debit.
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
   * Amount to debit, as a positive minor-units string.
   * @type {string}
   **/
  #amount: string = "";
  /**
   * Amount to debit, as a positive minor-units string.
   * @returns {string}
   **/
  get amount() {
    return this.#amount;
  }
  /**
   * Amount to debit, as a positive minor-units string.
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
   * Free-form name of the calling module/feature, e.g. "course-purchase".
   * @type {string}
   **/
  #referenceType: string = "";
  /**
   * Free-form name of the calling module/feature, e.g. "course-purchase".
   * @returns {string}
   **/
  get referenceType() {
    return this.#referenceType;
  }
  /**
   * Free-form name of the calling module/feature, e.g. "course-purchase".
   * @type {string}
   **/
  set referenceType(value: string) {
    this.#referenceType = String(value);
  }
  setReferenceType(value: string) {
    this.referenceType = value;
    return this;
  }
  /**
   * Id within referenceType this purchase relates to.
   * @type {string}
   **/
  #referenceId: string = "";
  /**
   * Id within referenceType this purchase relates to.
   * @returns {string}
   **/
  get referenceId() {
    return this.#referenceId;
  }
  /**
   * Id within referenceType this purchase relates to.
   * @type {string}
   **/
  set referenceId(value: string) {
    this.#referenceId = String(value);
  }
  setReferenceId(value: string) {
    this.referenceId = value;
    return this;
  }
  /**
   * Makes this purchase safe to retry without double-debiting.
   * @type {string}
   **/
  #idempotencyKey: string = "";
  /**
   * Makes this purchase safe to retry without double-debiting.
   * @returns {string}
   **/
  get idempotencyKey() {
    return this.#idempotencyKey;
  }
  /**
   * Makes this purchase safe to retry without double-debiting.
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
   * Optional human-readable note.
   * @type {string}
   **/
  #note?: string | null = undefined;
  /**
   * Optional human-readable note.
   * @returns {string}
   **/
  get note() {
    return this.#note;
  }
  /**
   * Optional human-readable note.
   * @type {string}
   **/
  set note(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#note = correctType ? value : String(value);
  }
  setNote(value: string | null | undefined) {
    this.note = value;
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
    const d = data as Partial<PurchaseActionReq>;
    if (d.walletId !== undefined) {
      this.walletId = d.walletId;
    }
    if (d.amount !== undefined) {
      this.amount = d.amount;
    }
    if (d.referenceType !== undefined) {
      this.referenceType = d.referenceType;
    }
    if (d.referenceId !== undefined) {
      this.referenceId = d.referenceId;
    }
    if (d.idempotencyKey !== undefined) {
      this.idempotencyKey = d.idempotencyKey;
    }
    if (d.note !== undefined) {
      this.note = d.note;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      walletId: this.#walletId,
      amount: this.#amount,
      referenceType: this.#referenceType,
      referenceId: this.#referenceId,
      idempotencyKey: this.#idempotencyKey,
      note: this.#note,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      walletId: "walletId",
      amount: "amount",
      referenceType: "referenceType",
      referenceId: "referenceId",
      idempotencyKey: "idempotencyKey",
      note: "note",
    };
  }
  /**
   * Creates an instance of PurchaseActionReq, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: PurchaseActionReqType) {
    return new PurchaseActionReq(possibleDtoObject);
  }
  /**
   * Creates an instance of PurchaseActionReq, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<PurchaseActionReqType>) {
    return new PurchaseActionReq(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<PurchaseActionReqType>,
  ): InstanceType<typeof PurchaseActionReq> {
    return new PurchaseActionReq({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof PurchaseActionReq> {
    return new PurchaseActionReq(this.toJSON());
  }
}
export abstract class PurchaseActionReqFactory {
  abstract create(data: unknown): PurchaseActionReq;
}
/**
 * The base type definition for purchaseActionReq
 **/
export type PurchaseActionReqType = {
  /**
   * Unique id of the wallet to debit.
   * @type {string}
   **/
  walletId: string;
  /**
   * Amount to debit, as a positive minor-units string.
   * @type {string}
   **/
  amount: string;
  /**
   * Free-form name of the calling module/feature, e.g. "course-purchase".
   * @type {string}
   **/
  referenceType: string;
  /**
   * Id within referenceType this purchase relates to.
   * @type {string}
   **/
  referenceId: string;
  /**
   * Makes this purchase safe to retry without double-debiting.
   * @type {string}
   **/
  idempotencyKey: string;
  /**
   * Optional human-readable note.
   * @type {string}
   **/
  note?: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PurchaseActionReqType {}
