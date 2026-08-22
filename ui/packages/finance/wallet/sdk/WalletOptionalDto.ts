import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for walletOptionalDto
 **/
export class WalletOptionalDto {
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
   * Who owns this wallet: "user" (userId is set, workspaceId empty) or "workspace" (workspaceId is set, userId empty). Exactly one of userId/workspaceId is populated, matching ownerType.
   * @type {string}
   **/
  #ownerType?: string | null = undefined;
  /**
   * Who owns this wallet: "user" (userId is set, workspaceId empty) or "workspace" (workspaceId is set, userId empty). Exactly one of userId/workspaceId is populated, matching ownerType.
   * @returns {string}
   **/
  get ownerType() {
    return this.#ownerType;
  }
  /**
   * Who owns this wallet: "user" (userId is set, workspaceId empty) or "workspace" (workspaceId is set, userId empty). Exactly one of userId/workspaceId is populated, matching ownerType.
   * @type {string}
   **/
  set ownerType(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#ownerType = correctType ? value : String(value);
  }
  setOwnerType(value: string | null | undefined) {
    this.ownerType = value;
    return this;
  }
  /**
   * Unique id of the owning user. Set only when ownerType is "user".
   * @type {string}
   **/
  #userId?: string | null = undefined;
  /**
   * Unique id of the owning user. Set only when ownerType is "user".
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * Unique id of the owning user. Set only when ownerType is "user".
   * @type {string}
   **/
  set userId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#userId = correctType ? value : String(value);
  }
  setUserId(value: string | null | undefined) {
    this.userId = value;
    return this;
  }
  /**
   * Unique id of the owning workspace. Set only when ownerType is "workspace".
   * @type {string}
   **/
  #workspaceId?: string | null = undefined;
  /**
   * Unique id of the owning workspace. Set only when ownerType is "workspace".
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   * Unique id of the owning workspace. Set only when ownerType is "workspace".
   * @type {string}
   **/
  set workspaceId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#workspaceId = correctType ? value : String(value);
  }
  setWorkspaceId(value: string | null | undefined) {
    this.workspaceId = value;
    return this;
  }
  /**
   * Currency code this wallet holds (must match an active walletCurrency.code, e.g. "USD" or "BTC"). Fixed for the wallet's lifetime - a currency change would invalidate its whole balance history, so a new wallet is created instead.
   * @type {string}
   **/
  #currency?: string | null = undefined;
  /**
   * Currency code this wallet holds (must match an active walletCurrency.code, e.g. "USD" or "BTC"). Fixed for the wallet's lifetime - a currency change would invalidate its whole balance history, so a new wallet is created instead.
   * @returns {string}
   **/
  get currency() {
    return this.#currency;
  }
  /**
   * Currency code this wallet holds (must match an active walletCurrency.code, e.g. "USD" or "BTC"). Fixed for the wallet's lifetime - a currency change would invalidate its whole balance history, so a new wallet is created instead.
   * @type {string}
   **/
  set currency(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#currency = correctType ? value : String(value);
  }
  setCurrency(value: string | null | undefined) {
    this.currency = value;
    return this;
  }
  /**
   * Current balance as a decimal string of integer minor-units at the wallet's currency's declared precision (see walletCurrency.decimals). Never read or written as a float. Only ever mutated inside a locked DB transaction by wallet.Purchase/adjustBalance, mirrored by a walletTransaction ledger row on every change.
   * @type {string}
   **/
  #balance?: string | null = "0";
  /**
   * Current balance as a decimal string of integer minor-units at the wallet's currency's declared precision (see walletCurrency.decimals). Never read or written as a float. Only ever mutated inside a locked DB transaction by wallet.Purchase/adjustBalance, mirrored by a walletTransaction ledger row on every change.
   * @returns {string}
   **/
  get balance() {
    return this.#balance;
  }
  /**
   * Current balance as a decimal string of integer minor-units at the wallet's currency's declared precision (see walletCurrency.decimals). Never read or written as a float. Only ever mutated inside a locked DB transaction by wallet.Purchase/adjustBalance, mirrored by a walletTransaction ledger row on every change.
   * @type {string}
   **/
  set balance(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#balance = correctType ? value : String(value);
  }
  setBalance(value: string | null | undefined) {
    this.balance = value;
    return this;
  }
  /**
   * "active" wallets can be topped up and purchased from. "frozen" blocks purchases/topups but keeps the wallet visible (e.g. pending dispute). "closed" is a terminal, non-reactivatable state.
   * @type {string}
   **/
  #status?: string | null = "active";
  /**
   * "active" wallets can be topped up and purchased from. "frozen" blocks purchases/topups but keeps the wallet visible (e.g. pending dispute). "closed" is a terminal, non-reactivatable state.
   * @returns {string}
   **/
  get status() {
    return this.#status;
  }
  /**
   * "active" wallets can be topped up and purchased from. "frozen" blocks purchases/topups but keeps the wallet visible (e.g. pending dispute). "closed" is a terminal, non-reactivatable state.
   * @type {string}
   **/
  set status(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#status = correctType ? value : String(value);
  }
  setStatus(value: string | null | undefined) {
    this.status = value;
    return this;
  }
  /**
   * Optional owner-given nickname for the wallet (e.g. "Main", "Savings") - purely cosmetic, shown in the owner-facing UI.
   * @type {string}
   **/
  #label?: string | null = undefined;
  /**
   * Optional owner-given nickname for the wallet (e.g. "Main", "Savings") - purely cosmetic, shown in the owner-facing UI.
   * @returns {string}
   **/
  get label() {
    return this.#label;
  }
  /**
   * Optional owner-given nickname for the wallet (e.g. "Main", "Savings") - purely cosmetic, shown in the owner-facing UI.
   * @type {string}
   **/
  set label(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#label = correctType ? value : String(value);
  }
  setLabel(value: string | null | undefined) {
    this.label = value;
    return this;
  }
  /**
   * Whether this is the owner's default wallet for its currency, used by internal purchase callers that don't specify a walletId explicitly. At most one wallet per (owner, currency) should have this set - enforced in Go, not at the DB level.
   * @type {boolean}
   **/
  #isDefault?: boolean | null = false;
  /**
   * Whether this is the owner's default wallet for its currency, used by internal purchase callers that don't specify a walletId explicitly. At most one wallet per (owner, currency) should have this set - enforced in Go, not at the DB level.
   * @returns {boolean}
   **/
  get isDefault() {
    return this.#isDefault;
  }
  /**
   * Whether this is the owner's default wallet for its currency, used by internal purchase callers that don't specify a walletId explicitly. At most one wallet per (owner, currency) should have this set - enforced in Go, not at the DB level.
   * @type {boolean}
   **/
  set isDefault(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#isDefault = correctType ? value : Boolean(value);
  }
  setIsDefault(value: boolean | null | undefined) {
    this.isDefault = value;
    return this;
  }
  /**
   * Incremented on every balance mutation. The actual concurrency guard is a DB row lock (SELECT ... FOR UPDATE) taken during purchase/adjustBalance; this field is a defense-in-depth optimistic check and an easy audit signal ("has this wallet ever moved").
   * @type {number}
   **/
  #version?: number | null = 0;
  /**
   * Incremented on every balance mutation. The actual concurrency guard is a DB row lock (SELECT ... FOR UPDATE) taken during purchase/adjustBalance; this field is a defense-in-depth optimistic check and an easy audit signal ("has this wallet ever moved").
   * @returns {number}
   **/
  get version() {
    return this.#version;
  }
  /**
   * Incremented on every balance mutation. The actual concurrency guard is a DB row lock (SELECT ... FOR UPDATE) taken during purchase/adjustBalance; this field is a defense-in-depth optimistic check and an easy audit signal ("has this wallet ever moved").
   * @type {number}
   **/
  set version(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#version = parsedValue;
    }
  }
  setVersion(value: number | null | undefined) {
    this.version = value;
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
    const d = data as Partial<WalletOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.ownerType !== undefined) {
      this.ownerType = d.ownerType;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
    }
    if (d.currency !== undefined) {
      this.currency = d.currency;
    }
    if (d.balance !== undefined) {
      this.balance = d.balance;
    }
    if (d.status !== undefined) {
      this.status = d.status;
    }
    if (d.label !== undefined) {
      this.label = d.label;
    }
    if (d.isDefault !== undefined) {
      this.isDefault = d.isDefault;
    }
    if (d.version !== undefined) {
      this.version = d.version;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      ownerType: this.#ownerType,
      userId: this.#userId,
      workspaceId: this.#workspaceId,
      currency: this.#currency,
      balance: this.#balance,
      status: this.#status,
      label: this.#label,
      isDefault: this.#isDefault,
      version: this.#version,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      ownerType: "ownerType",
      userId: "userId",
      workspaceId: "workspaceId",
      currency: "currency",
      balance: "balance",
      status: "status",
      label: "label",
      isDefault: "isDefault",
      version: "version",
    };
  }
  /**
   * Creates an instance of WalletOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WalletOptionalDtoType) {
    return new WalletOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WalletOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WalletOptionalDtoType>) {
    return new WalletOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WalletOptionalDtoType>,
  ): InstanceType<typeof WalletOptionalDto> {
    return new WalletOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WalletOptionalDto> {
    return new WalletOptionalDto(this.toJSON());
  }
}
export abstract class WalletOptionalDtoFactory {
  abstract create(data: unknown): WalletOptionalDto;
}
/**
 * The base type definition for walletOptionalDto
 **/
export type WalletOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * Who owns this wallet: "user" (userId is set, workspaceId empty) or "workspace" (workspaceId is set, userId empty). Exactly one of userId/workspaceId is populated, matching ownerType.
   * @type {string}
   **/
  ownerType?: string;
  /**
   * Unique id of the owning user. Set only when ownerType is "user".
   * @type {string}
   **/
  userId?: string;
  /**
   * Unique id of the owning workspace. Set only when ownerType is "workspace".
   * @type {string}
   **/
  workspaceId?: string;
  /**
   * Currency code this wallet holds (must match an active walletCurrency.code, e.g. "USD" or "BTC"). Fixed for the wallet's lifetime - a currency change would invalidate its whole balance history, so a new wallet is created instead.
   * @type {string}
   **/
  currency?: string;
  /**
   * Current balance as a decimal string of integer minor-units at the wallet's currency's declared precision (see walletCurrency.decimals). Never read or written as a float. Only ever mutated inside a locked DB transaction by wallet.Purchase/adjustBalance, mirrored by a walletTransaction ledger row on every change.
   * @type {string}
   **/
  balance?: string;
  /**
   * "active" wallets can be topped up and purchased from. "frozen" blocks purchases/topups but keeps the wallet visible (e.g. pending dispute). "closed" is a terminal, non-reactivatable state.
   * @type {string}
   **/
  status?: string;
  /**
   * Optional owner-given nickname for the wallet (e.g. "Main", "Savings") - purely cosmetic, shown in the owner-facing UI.
   * @type {string}
   **/
  label?: string;
  /**
   * Whether this is the owner's default wallet for its currency, used by internal purchase callers that don't specify a walletId explicitly. At most one wallet per (owner, currency) should have this set - enforced in Go, not at the DB level.
   * @type {boolean}
   **/
  isDefault?: boolean;
  /**
   * Incremented on every balance mutation. The actual concurrency guard is a DB row lock (SELECT ... FOR UPDATE) taken during purchase/adjustBalance; this field is a defense-in-depth optimistic check and an easy audit signal ("has this wallet ever moved").
   * @type {number}
   **/
  version?: number;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WalletOptionalDtoType {}
