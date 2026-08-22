import { MOne } from "@fireback/js-remote-ctx/common/operators";
import { WalletDto } from "./WalletDto";
import { WalletGatewayDto } from "./WalletGatewayDto";
import { WalletTransactionDto } from "./WalletTransactionDto";
import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for walletPaymentAttemptDto
 **/
export class WalletPaymentAttemptDto {
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
   * The wallet this attempt would credit/debit if it succeeds.
   * @type {WalletDto}
   **/
  #wallet?: MOne<WalletDto> | null = undefined;
  /**
   * The wallet this attempt would credit/debit if it succeeds.
   * @returns {WalletDto}
   **/
  get wallet() {
    return this.#wallet;
  }
  /**
   * The wallet this attempt would credit/debit if it succeeds.
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
   * The gateway this attempt is routed through.
   * @type {WalletGatewayDto}
   **/
  #gateway?: MOne<WalletGatewayDto> | null = undefined;
  /**
   * The gateway this attempt is routed through.
   * @returns {WalletGatewayDto}
   **/
  get gateway() {
    return this.#gateway;
  }
  /**
   * The gateway this attempt is routed through.
   * @type {WalletGatewayDto}
   **/
  set gateway(
    value:
      | MOne<WalletGatewayDto>
      | InstanceType<typeof WalletGatewayDto>
      | null
      | undefined,
  ) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof MOne) {
      this.#gateway = value;
    } else if (value instanceof WalletGatewayDto) {
      this.#gateway = MOne.of(value);
    } else {
      this.#gateway = MOne.of(new WalletGatewayDto(value));
    }
  }
  setGateway(
    value:
      | MOne<WalletGatewayDto>
      | InstanceType<typeof WalletGatewayDto>
      | null
      | undefined,
  ) {
    this.gateway = value;
    return this;
  }
  /**
   * What this attempt is for.
   * @type {string}
   **/
  #purpose: string = "topup";
  /**
   * What this attempt is for.
   * @returns {string}
   **/
  get purpose() {
    return this.#purpose;
  }
  /**
   * What this attempt is for.
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
   * Requested amount as a minor-units string, in currency.
   * @type {string}
   **/
  #amount: string = "";
  /**
   * Requested amount as a minor-units string, in currency.
   * @returns {string}
   **/
  get amount() {
    return this.#amount;
  }
  /**
   * Requested amount as a minor-units string, in currency.
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
   * Currency code for amount - must match the wallet's currency.
   * @type {string}
   **/
  #currency: string = "";
  /**
   * Currency code for amount - must match the wallet's currency.
   * @returns {string}
   **/
  get currency() {
    return this.#currency;
  }
  /**
   * Currency code for amount - must match the wallet's currency.
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
  #status: string = "pending";
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
   * The gateway's own id for this attempt (e.g. a PaymentIntent id or a transaction hash), once known. Indexed for webhook lookups.
   * @type {string}
   **/
  #gatewayReference?: string | null = undefined;
  /**
   * The gateway's own id for this attempt (e.g. a PaymentIntent id or a transaction hash), once known. Indexed for webhook lookups.
   * @returns {string}
   **/
  get gatewayReference() {
    return this.#gatewayReference;
  }
  /**
   * The gateway's own id for this attempt (e.g. a PaymentIntent id or a transaction hash), once known. Indexed for webhook lookups.
   * @type {string}
   **/
  set gatewayReference(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#gatewayReference = correctType ? value : String(value);
  }
  setGatewayReference(value: string | null | undefined) {
    this.gatewayReference = value;
    return this;
  }
  /**
   * Caller-supplied key making topup-initiation safe to retry without creating duplicate attempts at the gateway.
   * @type {string}
   **/
  #idempotencyKey: string = "";
  /**
   * Caller-supplied key making topup-initiation safe to retry without creating duplicate attempts at the gateway.
   * @returns {string}
   **/
  get idempotencyKey() {
    return this.#idempotencyKey;
  }
  /**
   * Caller-supplied key making topup-initiation safe to retry without creating duplicate attempts at the gateway.
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
   * The raw request sent to the gateway when initiating this attempt.
   * @type {JSON}
   **/
  #rawRequest!: JSON;
  /**
   * The raw request sent to the gateway when initiating this attempt.
   * @returns {JSON}
   **/
  get rawRequest() {
    return this.#rawRequest;
  }
  /**
   * The raw request sent to the gateway when initiating this attempt.
   * @type {JSON}
   **/
  set rawRequest(value: JSON) {
    this.#rawRequest = value;
  }
  setRawRequest(value: JSON) {
    this.rawRequest = value;
    return this;
  }
  /**
   * The raw response/init payload received back from the gateway.
   * @type {JSON}
   **/
  #rawResponse!: JSON;
  /**
   * The raw response/init payload received back from the gateway.
   * @returns {JSON}
   **/
  get rawResponse() {
    return this.#rawResponse;
  }
  /**
   * The raw response/init payload received back from the gateway.
   * @type {JSON}
   **/
  set rawResponse(value: JSON) {
    this.#rawResponse = value;
  }
  setRawResponse(value: JSON) {
    this.rawResponse = value;
    return this;
  }
  /**
   * The ledger entry that was created once this attempt succeeded. Empty until then.
   * @type {WalletTransactionDto}
   **/
  #walletTransaction?: MOne<WalletTransactionDto> | null = undefined;
  /**
   * The ledger entry that was created once this attempt succeeded. Empty until then.
   * @returns {WalletTransactionDto}
   **/
  get walletTransaction() {
    return this.#walletTransaction;
  }
  /**
   * The ledger entry that was created once this attempt succeeded. Empty until then.
   * @type {WalletTransactionDto}
   **/
  set walletTransaction(
    value:
      | MOne<WalletTransactionDto>
      | InstanceType<typeof WalletTransactionDto>
      | null
      | undefined,
  ) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof MOne) {
      this.#walletTransaction = value;
    } else if (value instanceof WalletTransactionDto) {
      this.#walletTransaction = MOne.of(value);
    } else {
      this.#walletTransaction = MOne.of(new WalletTransactionDto(value));
    }
  }
  setWalletTransaction(
    value:
      | MOne<WalletTransactionDto>
      | InstanceType<typeof WalletTransactionDto>
      | null
      | undefined,
  ) {
    this.walletTransaction = value;
    return this;
  }
  /**
   * When this attempt expires if not completed (gateway-dependent). Empty if the gateway doesn't impose one.
   * @type {XDate}
   **/
  #expiresAt!: XDate;
  /**
   * When this attempt expires if not completed (gateway-dependent). Empty if the gateway doesn't impose one.
   * @returns {XDate}
   **/
  get expiresAt() {
    return this.#expiresAt;
  }
  /**
   * When this attempt expires if not completed (gateway-dependent). Empty if the gateway doesn't impose one.
   * @type {XDate}
   **/
  set expiresAt(value: XDate) {
    this.#expiresAt = value;
  }
  setExpiresAt(value: XDate) {
    this.expiresAt = value;
    return this;
  }
  /**
   * When this attempt reached a terminal status. Empty until then.
   * @type {XDate}
   **/
  #completedAt!: XDate;
  /**
   * When this attempt reached a terminal status. Empty until then.
   * @returns {XDate}
   **/
  get completedAt() {
    return this.#completedAt;
  }
  /**
   * When this attempt reached a terminal status. Empty until then.
   * @type {XDate}
   **/
  set completedAt(value: XDate) {
    this.#completedAt = value;
  }
  setCompletedAt(value: XDate) {
    this.completedAt = value;
    return this;
  }
  /**
   * Where to send the caller's browser back to once a redirect-based gateway (e.g. Przelewy24, ZarinPal) completes the payment. Not needed by gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow).
   * @type {string}
   **/
  #returnUrl?: string | null = undefined;
  /**
   * Where to send the caller's browser back to once a redirect-based gateway (e.g. Przelewy24, ZarinPal) completes the payment. Not needed by gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow).
   * @returns {string}
   **/
  get returnUrl() {
    return this.#returnUrl;
  }
  /**
   * Where to send the caller's browser back to once a redirect-based gateway (e.g. Przelewy24, ZarinPal) completes the payment. Not needed by gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow).
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
    const d = data as Partial<WalletPaymentAttemptDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.wallet !== undefined) {
      this.wallet = d.wallet;
    }
    if (d.gateway !== undefined) {
      this.gateway = d.gateway;
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
    if (d.gatewayReference !== undefined) {
      this.gatewayReference = d.gatewayReference;
    }
    if (d.idempotencyKey !== undefined) {
      this.idempotencyKey = d.idempotencyKey;
    }
    if (d.createdAt !== undefined) {
      this.createdAt = d.createdAt;
    }
    if (d.failureReason !== undefined) {
      this.failureReason = d.failureReason;
    }
    if (d.rawRequest !== undefined) {
      this.rawRequest = d.rawRequest;
    }
    if (d.rawResponse !== undefined) {
      this.rawResponse = d.rawResponse;
    }
    if (d.walletTransaction !== undefined) {
      this.walletTransaction = d.walletTransaction;
    }
    if (d.expiresAt !== undefined) {
      this.expiresAt = d.expiresAt;
    }
    if (d.completedAt !== undefined) {
      this.completedAt = d.completedAt;
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
      uniqueId: this.#uniqueId,
      wallet: this.#wallet,
      gateway: this.#gateway,
      purpose: this.#purpose,
      amount: this.#amount,
      currency: this.#currency,
      status: this.#status,
      gatewayReference: this.#gatewayReference,
      idempotencyKey: this.#idempotencyKey,
      createdAt: this.#createdAt,
      failureReason: this.#failureReason,
      rawRequest: this.#rawRequest,
      rawResponse: this.#rawResponse,
      walletTransaction: this.#walletTransaction,
      expiresAt: this.#expiresAt,
      completedAt: this.#completedAt,
      returnUrl: this.#returnUrl,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      wallet: "wallet",
      gateway: "gateway",
      purpose: "purpose",
      amount: "amount",
      currency: "currency",
      status: "status",
      gatewayReference: "gatewayReference",
      idempotencyKey: "idempotencyKey",
      createdAt: "createdAt",
      failureReason: "failureReason",
      rawRequest: "rawRequest",
      rawResponse: "rawResponse",
      walletTransaction: "walletTransaction",
      expiresAt: "expiresAt",
      completedAt: "completedAt",
      returnUrl: "returnUrl",
    };
  }
  /**
   * Creates an instance of WalletPaymentAttemptDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WalletPaymentAttemptDtoType) {
    return new WalletPaymentAttemptDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WalletPaymentAttemptDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WalletPaymentAttemptDtoType>) {
    return new WalletPaymentAttemptDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WalletPaymentAttemptDtoType>,
  ): InstanceType<typeof WalletPaymentAttemptDto> {
    return new WalletPaymentAttemptDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WalletPaymentAttemptDto> {
    return new WalletPaymentAttemptDto(this.toJSON());
  }
}
export abstract class WalletPaymentAttemptDtoFactory {
  abstract create(data: unknown): WalletPaymentAttemptDto;
}
/**
 * The base type definition for walletPaymentAttemptDto
 **/
export type WalletPaymentAttemptDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * The wallet this attempt would credit/debit if it succeeds.
   * @type {WalletDto}
   **/
  wallet?: WalletDto;
  /**
   * The gateway this attempt is routed through.
   * @type {WalletGatewayDto}
   **/
  gateway?: WalletGatewayDto;
  /**
   * What this attempt is for.
   * @type {string}
   **/
  purpose: string;
  /**
   * Requested amount as a minor-units string, in currency.
   * @type {string}
   **/
  amount: string;
  /**
   * Currency code for amount - must match the wallet's currency.
   * @type {string}
   **/
  currency: string;
  /**
   * Current lifecycle state of this attempt.
   * @type {string}
   **/
  status: string;
  /**
   * The gateway's own id for this attempt (e.g. a PaymentIntent id or a transaction hash), once known. Indexed for webhook lookups.
   * @type {string}
   **/
  gatewayReference?: string;
  /**
   * Caller-supplied key making topup-initiation safe to retry without creating duplicate attempts at the gateway.
   * @type {string}
   **/
  idempotencyKey: string;
  /**
   * When this attempt was created.
   * @type {XDate}
   **/
  createdAt: XDate;
  /**
   * Human-readable reason, populated when status is "failed".
   * @type {string}
   **/
  failureReason?: string;
  /**
   * The raw request sent to the gateway when initiating this attempt.
   * @type {JSON}
   **/
  rawRequest: JSON;
  /**
   * The raw response/init payload received back from the gateway.
   * @type {JSON}
   **/
  rawResponse: JSON;
  /**
   * The ledger entry that was created once this attempt succeeded. Empty until then.
   * @type {WalletTransactionDto}
   **/
  walletTransaction?: WalletTransactionDto;
  /**
   * When this attempt expires if not completed (gateway-dependent). Empty if the gateway doesn't impose one.
   * @type {XDate}
   **/
  expiresAt: XDate;
  /**
   * When this attempt reached a terminal status. Empty until then.
   * @type {XDate}
   **/
  completedAt: XDate;
  /**
   * Where to send the caller's browser back to once a redirect-based gateway (e.g. Przelewy24, ZarinPal) completes the payment. Not needed by gateways that never redirect the browser (e.g. Stripe's client-secret confirmation flow).
   * @type {string}
   **/
  returnUrl?: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WalletPaymentAttemptDtoType {}
