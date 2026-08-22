import { MOne } from "@fireback/js-remote-ctx/common/operators";
import { WalletDto } from "./WalletDto";
import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for walletTransactionDto
 **/
export class WalletTransactionDto {
  /**
   *
   * @type {string}
   **/
  #uniqueId?: string | null = undefined;
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
  set uniqueId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#uniqueId = correctType ? value : String(value);
  }
  setUniqueId(value: string | null | undefined) {
    this.uniqueId = value;
    return this;
  }
  /**
   * The wallet this ledger entry belongs to.
   * @type {WalletDto}
   **/
  #wallet?: MOne<WalletDto> | null = undefined;
  /**
   * The wallet this ledger entry belongs to.
   * @returns {WalletDto}
   **/
  get wallet() {
    return this.#wallet;
  }
  /**
   * The wallet this ledger entry belongs to.
   * @type {WalletDto}
   **/
  set wallet(
    value: MOne<WalletDto> | InstanceType<typeof WalletDto> | null | undefined,
  ) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof MOne) {
      this.#wallet = value;
    } else if (value instanceof WalletDto) {
      this.#wallet = MOne.of(value);
    } else {
      this.#wallet = MOne.of(new WalletDto(value));
    }
  }
  setWallet(
    value: MOne<WalletDto> | InstanceType<typeof WalletDto> | null | undefined,
  ) {
    this.wallet = value;
    return this;
  }
  /**
   * "credit" increases the wallet's balance, "debit" decreases it. amount is always a positive minor-units string; direction carries the sign.
   * @type {string}
   **/
  #direction: string = "";
  /**
   * "credit" increases the wallet's balance, "debit" decreases it. amount is always a positive minor-units string; direction carries the sign.
   * @returns {string}
   **/
  get direction() {
    return this.#direction;
  }
  /**
   * "credit" increases the wallet's balance, "debit" decreases it. amount is always a positive minor-units string; direction carries the sign.
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
   * Magnitude of this change, as a positive decimal string of integer minor-units at the wallet's currency precision.
   * @type {string}
   **/
  #amount: string = "";
  /**
   * Magnitude of this change, as a positive decimal string of integer minor-units at the wallet's currency precision.
   * @returns {string}
   **/
  get amount() {
    return this.#amount;
  }
  /**
   * Magnitude of this change, as a positive decimal string of integer minor-units at the wallet's currency precision.
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
   * Snapshot of the wallet's balance immediately after this entry was applied, as a minor-units string. Kept for audit/debugging even though it's derivable by replaying the ledger.
   * @type {string}
   **/
  #balanceAfter: string = "";
  /**
   * Snapshot of the wallet's balance immediately after this entry was applied, as a minor-units string. Kept for audit/debugging even though it's derivable by replaying the ledger.
   * @returns {string}
   **/
  get balanceAfter() {
    return this.#balanceAfter;
  }
  /**
   * Snapshot of the wallet's balance immediately after this entry was applied, as a minor-units string. Kept for audit/debugging even though it's derivable by replaying the ledger.
   * @type {string}
   **/
  set balanceAfter(value: string) {
    this.#balanceAfter = String(value);
  }
  setBalanceAfter(value: string) {
    this.balanceAfter = value;
    return this;
  }
  /**
   * What kind of event produced this ledger entry.
   * @type {string}
   **/
  #reason: string = "";
  /**
   * What kind of event produced this ledger entry.
   * @returns {string}
   **/
  get reason() {
    return this.#reason;
  }
  /**
   * What kind of event produced this ledger entry.
   * @type {string}
   **/
  set reason(value: string) {
    this.#reason = String(value);
  }
  setReason(value: string) {
    this.reason = value;
    return this;
  }
  /**
   * Free-form identifier of the calling module/feature that caused this entry, e.g. "course-purchase" - lets other nima modules tag their own purchases without editing this module's yaml.
   * @type {string}
   **/
  #referenceType?: string | null = undefined;
  /**
   * Free-form identifier of the calling module/feature that caused this entry, e.g. "course-purchase" - lets other nima modules tag their own purchases without editing this module's yaml.
   * @returns {string}
   **/
  get referenceType() {
    return this.#referenceType;
  }
  /**
   * Free-form identifier of the calling module/feature that caused this entry, e.g. "course-purchase" - lets other nima modules tag their own purchases without editing this module's yaml.
   * @type {string}
   **/
  set referenceType(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#referenceType = correctType ? value : String(value);
  }
  setReferenceType(value: string | null | undefined) {
    this.referenceType = value;
    return this;
  }
  /**
   * Id within referenceType this entry relates to (e.g. the order id in the calling module).
   * @type {string}
   **/
  #referenceId?: string | null = undefined;
  /**
   * Id within referenceType this entry relates to (e.g. the order id in the calling module).
   * @returns {string}
   **/
  get referenceId() {
    return this.#referenceId;
  }
  /**
   * Id within referenceType this entry relates to (e.g. the order id in the calling module).
   * @type {string}
   **/
  set referenceId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#referenceId = correctType ? value : String(value);
  }
  setReferenceId(value: string | null | undefined) {
    this.referenceId = value;
    return this;
  }
  /**
   * Caller-supplied key that makes the operation that produced this entry safe to retry - a second call with the same key returns the existing row instead of double-applying the change. Unique across all wallet transactions.
   * @type {string}
   **/
  #idempotencyKey: string = "";
  /**
   * Caller-supplied key that makes the operation that produced this entry safe to retry - a second call with the same key returns the existing row instead of double-applying the change. Unique across all wallet transactions.
   * @returns {string}
   **/
  get idempotencyKey() {
    return this.#idempotencyKey;
  }
  /**
   * Caller-supplied key that makes the operation that produced this entry safe to retry - a second call with the same key returns the existing row instead of double-applying the change. Unique across all wallet transactions.
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
   * Optional human-readable note, e.g. why an adjustment was made.
   * @type {string}
   **/
  #note?: string | null = undefined;
  /**
   * Optional human-readable note, e.g. why an adjustment was made.
   * @returns {string}
   **/
  get note() {
    return this.#note;
  }
  /**
   * Optional human-readable note, e.g. why an adjustment was made.
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
  /**
   * Unique id of the user who triggered this entry, or a fixed marker such as "system" or the gateway code, for entries with no human actor.
   * @type {string}
   **/
  #createdBy?: string | null = undefined;
  /**
   * Unique id of the user who triggered this entry, or a fixed marker such as "system" or the gateway code, for entries with no human actor.
   * @returns {string}
   **/
  get createdBy() {
    return this.#createdBy;
  }
  /**
   * Unique id of the user who triggered this entry, or a fixed marker such as "system" or the gateway code, for entries with no human actor.
   * @type {string}
   **/
  set createdBy(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#createdBy = correctType ? value : String(value);
  }
  setCreatedBy(value: string | null | undefined) {
    this.createdBy = value;
    return this;
  }
  /**
   * When this ledger entry was recorded.
   * @type {XDate}
   **/
  #createdAt!: XDate;
  /**
   * When this ledger entry was recorded.
   * @returns {XDate}
   **/
  get createdAt() {
    return this.#createdAt;
  }
  /**
   * When this ledger entry was recorded.
   * @type {XDate}
   **/
  set createdAt(value: XDate) {
    this.#createdAt = value;
  }
  setCreatedAt(value: XDate) {
    this.createdAt = value;
    return this;
  }
  /**
   * Extra structured, reason-specific data (e.g. gateway fee breakdown) not worth a dedicated column.
   * @type {JSON}
   **/
  #metadata!: JSON;
  /**
   * Extra structured, reason-specific data (e.g. gateway fee breakdown) not worth a dedicated column.
   * @returns {JSON}
   **/
  get metadata() {
    return this.#metadata;
  }
  /**
   * Extra structured, reason-specific data (e.g. gateway fee breakdown) not worth a dedicated column.
   * @type {JSON}
   **/
  set metadata(value: JSON) {
    this.#metadata = value;
  }
  setMetadata(value: JSON) {
    this.metadata = value;
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
    const d = data as Partial<WalletTransactionDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.wallet !== undefined) {
      this.wallet = d.wallet;
    }
    if (d.direction !== undefined) {
      this.direction = d.direction;
    }
    if (d.amount !== undefined) {
      this.amount = d.amount;
    }
    if (d.balanceAfter !== undefined) {
      this.balanceAfter = d.balanceAfter;
    }
    if (d.reason !== undefined) {
      this.reason = d.reason;
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
    if (d.createdBy !== undefined) {
      this.createdBy = d.createdBy;
    }
    if (d.createdAt !== undefined) {
      this.createdAt = d.createdAt;
    }
    if (d.metadata !== undefined) {
      this.metadata = d.metadata;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      wallet: this.#wallet,
      direction: this.#direction,
      amount: this.#amount,
      balanceAfter: this.#balanceAfter,
      reason: this.#reason,
      referenceType: this.#referenceType,
      referenceId: this.#referenceId,
      idempotencyKey: this.#idempotencyKey,
      note: this.#note,
      createdBy: this.#createdBy,
      createdAt: this.#createdAt,
      metadata: this.#metadata,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      wallet: "wallet",
      direction: "direction",
      amount: "amount",
      balanceAfter: "balanceAfter",
      reason: "reason",
      referenceType: "referenceType",
      referenceId: "referenceId",
      idempotencyKey: "idempotencyKey",
      note: "note",
      createdBy: "createdBy",
      createdAt: "createdAt",
      metadata: "metadata",
    };
  }
  /**
   * Creates an instance of WalletTransactionDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WalletTransactionDtoType) {
    return new WalletTransactionDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WalletTransactionDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WalletTransactionDtoType>) {
    return new WalletTransactionDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WalletTransactionDtoType>,
  ): InstanceType<typeof WalletTransactionDto> {
    return new WalletTransactionDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WalletTransactionDto> {
    return new WalletTransactionDto(this.toJSON());
  }
}
export abstract class WalletTransactionDtoFactory {
  abstract create(data: unknown): WalletTransactionDto;
}
/**
 * The base type definition for walletTransactionDto
 **/
export type WalletTransactionDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * The wallet this ledger entry belongs to.
   * @type {WalletDto}
   **/
  wallet?: WalletDto;
  /**
   * "credit" increases the wallet's balance, "debit" decreases it. amount is always a positive minor-units string; direction carries the sign.
   * @type {string}
   **/
  direction: string;
  /**
   * Magnitude of this change, as a positive decimal string of integer minor-units at the wallet's currency precision.
   * @type {string}
   **/
  amount: string;
  /**
   * Snapshot of the wallet's balance immediately after this entry was applied, as a minor-units string. Kept for audit/debugging even though it's derivable by replaying the ledger.
   * @type {string}
   **/
  balanceAfter: string;
  /**
   * What kind of event produced this ledger entry.
   * @type {string}
   **/
  reason: string;
  /**
   * Free-form identifier of the calling module/feature that caused this entry, e.g. "course-purchase" - lets other nima modules tag their own purchases without editing this module's yaml.
   * @type {string}
   **/
  referenceType?: string;
  /**
   * Id within referenceType this entry relates to (e.g. the order id in the calling module).
   * @type {string}
   **/
  referenceId?: string;
  /**
   * Caller-supplied key that makes the operation that produced this entry safe to retry - a second call with the same key returns the existing row instead of double-applying the change. Unique across all wallet transactions.
   * @type {string}
   **/
  idempotencyKey: string;
  /**
   * Optional human-readable note, e.g. why an adjustment was made.
   * @type {string}
   **/
  note?: string;
  /**
   * Unique id of the user who triggered this entry, or a fixed marker such as "system" or the gateway code, for entries with no human actor.
   * @type {string}
   **/
  createdBy?: string;
  /**
   * When this ledger entry was recorded.
   * @type {XDate}
   **/
  createdAt: XDate;
  /**
   * Extra structured, reason-specific data (e.g. gateway fee breakdown) not worth a dedicated column.
   * @type {JSON}
   **/
  metadata: JSON;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WalletTransactionDtoType {}
