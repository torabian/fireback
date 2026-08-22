import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for walletCurrencyDto
 **/
export class WalletCurrencyDto {
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
   * Currency code, e.g. "USD", "EUR", "BTC", "ETH". Unique across all currencies.
   * @type {string}
   **/
  #code: string = "";
  /**
   * Currency code, e.g. "USD", "EUR", "BTC", "ETH". Unique across all currencies.
   * @returns {string}
   **/
  get code() {
    return this.#code;
  }
  /**
   * Currency code, e.g. "USD", "EUR", "BTC", "ETH". Unique across all currencies.
   * @type {string}
   **/
  set code(value: string) {
    this.#code = String(value);
  }
  setCode(value: string) {
    this.code = value;
    return this;
  }
  /**
   * Display name, e.g. "US Dollar" or "Bitcoin".
   * @type {string}
   **/
  #name: string = "";
  /**
   * Display name, e.g. "US Dollar" or "Bitcoin".
   * @returns {string}
   **/
  get name() {
    return this.#name;
  }
  /**
   * Display name, e.g. "US Dollar" or "Bitcoin".
   * @type {string}
   **/
  set name(value: string) {
    this.#name = String(value);
  }
  setName(value: string) {
    this.name = value;
    return this;
  }
  /**
   * Whether this is a fiat or crypto currency.
   * @type {string}
   **/
  #kind: string = "";
  /**
   * Whether this is a fiat or crypto currency.
   * @returns {string}
   **/
  get kind() {
    return this.#kind;
  }
  /**
   * Whether this is a fiat or crypto currency.
   * @type {string}
   **/
  set kind(value: string) {
    this.#kind = String(value);
  }
  setKind(value: string) {
    this.kind = value;
    return this;
  }
  /**
   * Number of decimal places a minor-unit amount string represents for this currency (e.g. 2 for USD so "10050" means $100.50, 8 for BTC, 18 for ETH). Every wallet/transaction/attempt amount in this currency must be interpreted at this precision.
   * @type {number}
   **/
  #decimals: number = 0;
  /**
   * Number of decimal places a minor-unit amount string represents for this currency (e.g. 2 for USD so "10050" means $100.50, 8 for BTC, 18 for ETH). Every wallet/transaction/attempt amount in this currency must be interpreted at this precision.
   * @returns {number}
   **/
  get decimals() {
    return this.#decimals;
  }
  /**
   * Number of decimal places a minor-unit amount string represents for this currency (e.g. 2 for USD so "10050" means $100.50, 8 for BTC, 18 for ETH). Every wallet/transaction/attempt amount in this currency must be interpreted at this precision.
   * @type {number}
   **/
  set decimals(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#decimals = parsedValue;
    }
  }
  setDecimals(value: number) {
    this.decimals = value;
    return this;
  }
  /**
   * Optional display symbol, e.g. "$" or "₿".
   * @type {string}
   **/
  #symbol?: string | null = undefined;
  /**
   * Optional display symbol, e.g. "$" or "₿".
   * @returns {string}
   **/
  get symbol() {
    return this.#symbol;
  }
  /**
   * Optional display symbol, e.g. "$" or "₿".
   * @type {string}
   **/
  set symbol(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#symbol = correctType ? value : String(value);
  }
  setSymbol(value: string | null | undefined) {
    this.symbol = value;
    return this;
  }
  /**
   * Whether wallets/topups can currently be created in this currency. Existing wallets in a deactivated currency keep working; only new creation is blocked.
   * @type {boolean}
   **/
  #isActive: boolean = true;
  /**
   * Whether wallets/topups can currently be created in this currency. Existing wallets in a deactivated currency keep working; only new creation is blocked.
   * @returns {boolean}
   **/
  get isActive() {
    return this.#isActive;
  }
  /**
   * Whether wallets/topups can currently be created in this currency. Existing wallets in a deactivated currency keep working; only new creation is blocked.
   * @type {boolean}
   **/
  set isActive(value: boolean) {
    this.#isActive = Boolean(value);
  }
  setIsActive(value: boolean) {
    this.isActive = value;
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
    const d = data as Partial<WalletCurrencyDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.code !== undefined) {
      this.code = d.code;
    }
    if (d.name !== undefined) {
      this.name = d.name;
    }
    if (d.kind !== undefined) {
      this.kind = d.kind;
    }
    if (d.decimals !== undefined) {
      this.decimals = d.decimals;
    }
    if (d.symbol !== undefined) {
      this.symbol = d.symbol;
    }
    if (d.isActive !== undefined) {
      this.isActive = d.isActive;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      code: this.#code,
      name: this.#name,
      kind: this.#kind,
      decimals: this.#decimals,
      symbol: this.#symbol,
      isActive: this.#isActive,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      code: "code",
      name: "name",
      kind: "kind",
      decimals: "decimals",
      symbol: "symbol",
      isActive: "isActive",
    };
  }
  /**
   * Creates an instance of WalletCurrencyDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WalletCurrencyDtoType) {
    return new WalletCurrencyDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WalletCurrencyDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WalletCurrencyDtoType>) {
    return new WalletCurrencyDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WalletCurrencyDtoType>,
  ): InstanceType<typeof WalletCurrencyDto> {
    return new WalletCurrencyDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WalletCurrencyDto> {
    return new WalletCurrencyDto(this.toJSON());
  }
}
export abstract class WalletCurrencyDtoFactory {
  abstract create(data: unknown): WalletCurrencyDto;
}
/**
 * The base type definition for walletCurrencyDto
 **/
export type WalletCurrencyDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * Currency code, e.g. "USD", "EUR", "BTC", "ETH". Unique across all currencies.
   * @type {string}
   **/
  code: string;
  /**
   * Display name, e.g. "US Dollar" or "Bitcoin".
   * @type {string}
   **/
  name: string;
  /**
   * Whether this is a fiat or crypto currency.
   * @type {string}
   **/
  kind: string;
  /**
   * Number of decimal places a minor-unit amount string represents for this currency (e.g. 2 for USD so "10050" means $100.50, 8 for BTC, 18 for ETH). Every wallet/transaction/attempt amount in this currency must be interpreted at this precision.
   * @type {number}
   **/
  decimals: number;
  /**
   * Optional display symbol, e.g. "$" or "₿".
   * @type {string}
   **/
  symbol?: string;
  /**
   * Whether wallets/topups can currently be created in this currency. Existing wallets in a deactivated currency keep working; only new creation is blocked.
   * @type {boolean}
   **/
  isActive: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WalletCurrencyDtoType {}
