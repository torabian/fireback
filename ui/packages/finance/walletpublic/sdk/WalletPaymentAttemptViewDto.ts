import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for walletPaymentAttemptViewDto
 **/
export class WalletPaymentAttemptViewDto {
  /**
   * Unique id of this attempt.
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * Unique id of this attempt.
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * Unique id of this attempt.
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
   * "topup", "purchase", or "withdrawal".
   * @type {string}
   **/
  #purpose: string = "";
  /**
   * "topup", "purchase", or "withdrawal".
   * @returns {string}
   **/
  get purpose() {
    return this.#purpose;
  }
  /**
   * "topup", "purchase", or "withdrawal".
   * @type {string}
   **/
  set purpose(value: string) {
    this.#purpose = String(value);
  }
  setPurpose(value: string) {
    this.purpose = value;
    return this;
  }
  /**
   * Requested amount as a minor-units string.
   * @type {string}
   **/
  #amount: string = "";
  /**
   * Requested amount as a minor-units string.
   * @returns {string}
   **/
  get amount() {
    return this.#amount;
  }
  /**
   * Requested amount as a minor-units string.
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
   * Currency code for amount.
   * @type {string}
   **/
  #currency: string = "";
  /**
   * Currency code for amount.
   * @returns {string}
   **/
  get currency() {
    return this.#currency;
  }
  /**
   * Currency code for amount.
   * @type {string}
   **/
  set currency(value: string) {
    this.#currency = String(value);
  }
  setCurrency(value: string) {
    this.currency = value;
    return this;
  }
  /**
   * Current lifecycle state of this attempt.
   * @type {string}
   **/
  #status: string = "";
  /**
   * Current lifecycle state of this attempt.
   * @returns {string}
   **/
  get status() {
    return this.#status;
  }
  /**
   * Current lifecycle state of this attempt.
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
   * Code of the gateway this attempt is routed through.
   * @type {string}
   **/
  #gatewayCode: string = "";
  /**
   * Code of the gateway this attempt is routed through.
   * @returns {string}
   **/
  get gatewayCode() {
    return this.#gatewayCode;
  }
  /**
   * Code of the gateway this attempt is routed through.
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
   * Human-readable reason, populated when status is "failed".
   * @type {string}
   **/
  #failureReason?: string | null = undefined;
  /**
   * Human-readable reason, populated when status is "failed".
   * @returns {string}
   **/
  get failureReason() {
    return this.#failureReason;
  }
  /**
   * Human-readable reason, populated when status is "failed".
   * @type {string}
   **/
  set failureReason(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#failureReason = correctType ? value : String(value);
  }
  setFailureReason(value: string | null | undefined) {
    this.failureReason = value;
    return this;
  }
  /**
   * When this attempt was created.
   * @type {XDate}
   **/
  #createdAt!: XDate;
  /**
   * When this attempt was created.
   * @returns {XDate}
   **/
  get createdAt() {
    return this.#createdAt;
  }
  /**
   * When this attempt was created.
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
    const d = data as Partial<WalletPaymentAttemptViewDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.purpose !== undefined) {
      this.purpose = d.purpose;
    }
    if (d.amount !== undefined) {
      this.amount = d.amount;
    }
    if (d.currency !== undefined) {
      this.currency = d.currency;
    }
    if (d.status !== undefined) {
      this.status = d.status;
    }
    if (d.gatewayCode !== undefined) {
      this.gatewayCode = d.gatewayCode;
    }
    if (d.failureReason !== undefined) {
      this.failureReason = d.failureReason;
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
      purpose: this.#purpose,
      amount: this.#amount,
      currency: this.#currency,
      status: this.#status,
      gatewayCode: this.#gatewayCode,
      failureReason: this.#failureReason,
      createdAt: this.#createdAt,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      purpose: "purpose",
      amount: "amount",
      currency: "currency",
      status: "status",
      gatewayCode: "gatewayCode",
      failureReason: "failureReason",
      createdAt: "createdAt",
    };
  }
  /**
   * Creates an instance of WalletPaymentAttemptViewDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WalletPaymentAttemptViewDtoType) {
    return new WalletPaymentAttemptViewDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WalletPaymentAttemptViewDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WalletPaymentAttemptViewDtoType>) {
    return new WalletPaymentAttemptViewDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WalletPaymentAttemptViewDtoType>,
  ): InstanceType<typeof WalletPaymentAttemptViewDto> {
    return new WalletPaymentAttemptViewDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WalletPaymentAttemptViewDto> {
    return new WalletPaymentAttemptViewDto(this.toJSON());
  }
}
export abstract class WalletPaymentAttemptViewDtoFactory {
  abstract create(data: unknown): WalletPaymentAttemptViewDto;
}
/**
 * The base type definition for walletPaymentAttemptViewDto
 **/
export type WalletPaymentAttemptViewDtoType = {
  /**
   * Unique id of this attempt.
   * @type {string}
   **/
  uniqueId: string;
  /**
   * "topup", "purchase", or "withdrawal".
   * @type {string}
   **/
  purpose: string;
  /**
   * Requested amount as a minor-units string.
   * @type {string}
   **/
  amount: string;
  /**
   * Currency code for amount.
   * @type {string}
   **/
  currency: string;
  /**
   * Current lifecycle state of this attempt.
   * @type {string}
   **/
  status: string;
  /**
   * Code of the gateway this attempt is routed through.
   * @type {string}
   **/
  gatewayCode: string;
  /**
   * Human-readable reason, populated when status is "failed".
   * @type {string}
   **/
  failureReason?: string;
  /**
   * When this attempt was created.
   * @type {XDate}
   **/
  createdAt: XDate;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WalletPaymentAttemptViewDtoType {}
