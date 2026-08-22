import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for walletTransactionViewDto
 **/
export class WalletTransactionViewDto {
  /**
   * Unique id of this ledger entry.
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * Unique id of this ledger entry.
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * Unique id of this ledger entry.
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
   * "credit" or "debit".
   * @type {string}
   **/
  #direction: string = "";
  /**
   * "credit" or "debit".
   * @returns {string}
   **/
  get direction() {
    return this.#direction;
  }
  /**
   * "credit" or "debit".
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
   * Magnitude of the change, as a positive minor-units string.
   * @type {string}
   **/
  #amount: string = "";
  /**
   * Magnitude of the change, as a positive minor-units string.
   * @returns {string}
   **/
  get amount() {
    return this.#amount;
  }
  /**
   * Magnitude of the change, as a positive minor-units string.
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
   * Wallet balance immediately after this entry, as a minor-units string.
   * @type {string}
   **/
  #balanceAfter: string = "";
  /**
   * Wallet balance immediately after this entry, as a minor-units string.
   * @returns {string}
   **/
  get balanceAfter() {
    return this.#balanceAfter;
  }
  /**
   * Wallet balance immediately after this entry, as a minor-units string.
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
   * What kind of event produced this entry.
   * @type {string}
   **/
  #reason: string = "";
  /**
   * What kind of event produced this entry.
   * @returns {string}
   **/
  get reason() {
    return this.#reason;
  }
  /**
   * What kind of event produced this entry.
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
   * Free-form name of the module/feature that caused this entry.
   * @type {string}
   **/
  #referenceType?: string | null = undefined;
  /**
   * Free-form name of the module/feature that caused this entry.
   * @returns {string}
   **/
  get referenceType() {
    return this.#referenceType;
  }
  /**
   * Free-form name of the module/feature that caused this entry.
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
   * Id within referenceType this entry relates to.
   * @type {string}
   **/
  #referenceId?: string | null = undefined;
  /**
   * Id within referenceType this entry relates to.
   * @returns {string}
   **/
  get referenceId() {
    return this.#referenceId;
  }
  /**
   * Id within referenceType this entry relates to.
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
  /**
   * When this entry was recorded.
   * @type {XDate}
   **/
  #createdAt!: XDate;
  /**
   * When this entry was recorded.
   * @returns {XDate}
   **/
  get createdAt() {
    return this.#createdAt;
  }
  /**
   * When this entry was recorded.
   * @type {XDate}
   **/
  set createdAt(value: XDate) {
    this.#createdAt = value;
  }
  setCreatedAt(value: XDate) {
    this.createdAt = value;
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
    const d = data as Partial<WalletTransactionViewDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
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
    if (d.note !== undefined) {
      this.note = d.note;
    }
    if (d.createdAt !== undefined) {
      this.createdAt = d.createdAt;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      direction: this.#direction,
      amount: this.#amount,
      balanceAfter: this.#balanceAfter,
      reason: this.#reason,
      referenceType: this.#referenceType,
      referenceId: this.#referenceId,
      note: this.#note,
      createdAt: this.#createdAt,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      direction: "direction",
      amount: "amount",
      balanceAfter: "balanceAfter",
      reason: "reason",
      referenceType: "referenceType",
      referenceId: "referenceId",
      note: "note",
      createdAt: "createdAt",
    };
  }
  /**
   * Creates an instance of WalletTransactionViewDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WalletTransactionViewDtoType) {
    return new WalletTransactionViewDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WalletTransactionViewDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WalletTransactionViewDtoType>) {
    return new WalletTransactionViewDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WalletTransactionViewDtoType>,
  ): InstanceType<typeof WalletTransactionViewDto> {
    return new WalletTransactionViewDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WalletTransactionViewDto> {
    return new WalletTransactionViewDto(this.toJSON());
  }
}
export abstract class WalletTransactionViewDtoFactory {
  abstract create(data: unknown): WalletTransactionViewDto;
}
/**
 * The base type definition for walletTransactionViewDto
 **/
export type WalletTransactionViewDtoType = {
  /**
   * Unique id of this ledger entry.
   * @type {string}
   **/
  uniqueId: string;
  /**
   * "credit" or "debit".
   * @type {string}
   **/
  direction: string;
  /**
   * Magnitude of the change, as a positive minor-units string.
   * @type {string}
   **/
  amount: string;
  /**
   * Wallet balance immediately after this entry, as a minor-units string.
   * @type {string}
   **/
  balanceAfter: string;
  /**
   * What kind of event produced this entry.
   * @type {string}
   **/
  reason: string;
  /**
   * Free-form name of the module/feature that caused this entry.
   * @type {string}
   **/
  referenceType?: string;
  /**
   * Id within referenceType this entry relates to.
   * @type {string}
   **/
  referenceId?: string;
  /**
   * Optional human-readable note.
   * @type {string}
   **/
  note?: string;
  /**
   * When this entry was recorded.
   * @type {XDate}
   **/
  createdAt: XDate;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WalletTransactionViewDtoType {}
