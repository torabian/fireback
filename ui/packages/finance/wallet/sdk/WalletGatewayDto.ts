import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for walletGatewayDto
 **/
export class WalletGatewayDto {
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
   * Unique code identifying this gateway's GatewayAdapter implementation, e.g. "stripe", "manual", "onchain_eth".
   * @type {string}
   **/
  #code: string = "";
  /**
   * Unique code identifying this gateway's GatewayAdapter implementation, e.g. "stripe", "manual", "onchain_eth".
   * @returns {string}
   **/
  get code() {
    return this.#code;
  }
  /**
   * Unique code identifying this gateway's GatewayAdapter implementation, e.g. "stripe", "manual", "onchain_eth".
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
   * Display name shown to wallet owners choosing a topup method.
   * @type {string}
   **/
  #name: string = "";
  /**
   * Display name shown to wallet owners choosing a topup method.
   * @returns {string}
   **/
  get name() {
    return this.#name;
  }
  /**
   * Display name shown to wallet owners choosing a topup method.
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
   * Whether this gateway settles in fiat or crypto.
   * @type {string}
   **/
  #kind: string = "";
  /**
   * Whether this gateway settles in fiat or crypto.
   * @returns {string}
   **/
  get kind() {
    return this.#kind;
  }
  /**
   * Whether this gateway settles in fiat or crypto.
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
   * Whether wallet owners can currently start a topup through this gateway.
   * @type {boolean}
   **/
  #isActive: boolean = true;
  /**
   * Whether wallet owners can currently start a topup through this gateway.
   * @returns {boolean}
   **/
  get isActive() {
    return this.#isActive;
  }
  /**
   * Whether wallet owners can currently start a topup through this gateway.
   * @type {boolean}
   **/
  set isActive(value: boolean) {
    this.#isActive = Boolean(value);
  }
  setIsActive(value: boolean) {
    this.isActive = value;
    return this;
  }
  /**
   * Provider configuration. Must only hold references to secrets (e.g. a secrets-manager key name), never raw API keys/webhook secrets in plaintext.
   * @type {JSON}
   **/
  #config!: JSON;
  /**
   * Provider configuration. Must only hold references to secrets (e.g. a secrets-manager key name), never raw API keys/webhook secrets in plaintext.
   * @returns {JSON}
   **/
  get config() {
    return this.#config;
  }
  /**
   * Provider configuration. Must only hold references to secrets (e.g. a secrets-manager key name), never raw API keys/webhook secrets in plaintext.
   * @type {JSON}
   **/
  set config(value: JSON) {
    this.#config = value;
  }
  setConfig(value: JSON) {
    this.config = value;
    return this;
  }
  /**
   * JSON array of walletCurrency codes this gateway can top up (e.g. ["USD","EUR"]).
   * @type {JSON}
   **/
  #supportedCurrencies!: JSON;
  /**
   * JSON array of walletCurrency codes this gateway can top up (e.g. ["USD","EUR"]).
   * @returns {JSON}
   **/
  get supportedCurrencies() {
    return this.#supportedCurrencies;
  }
  /**
   * JSON array of walletCurrency codes this gateway can top up (e.g. ["USD","EUR"]).
   * @type {JSON}
   **/
  set supportedCurrencies(value: JSON) {
    this.#supportedCurrencies = value;
  }
  setSupportedCurrencies(value: JSON) {
    this.supportedCurrencies = value;
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
    const d = data as Partial<WalletGatewayDto>;
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
    if (d.isActive !== undefined) {
      this.isActive = d.isActive;
    }
    if (d.config !== undefined) {
      this.config = d.config;
    }
    if (d.supportedCurrencies !== undefined) {
      this.supportedCurrencies = d.supportedCurrencies;
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
      isActive: this.#isActive,
      config: this.#config,
      supportedCurrencies: this.#supportedCurrencies,
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
      isActive: "isActive",
      config: "config",
      supportedCurrencies: "supportedCurrencies",
    };
  }
  /**
   * Creates an instance of WalletGatewayDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WalletGatewayDtoType) {
    return new WalletGatewayDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WalletGatewayDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WalletGatewayDtoType>) {
    return new WalletGatewayDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WalletGatewayDtoType>,
  ): InstanceType<typeof WalletGatewayDto> {
    return new WalletGatewayDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WalletGatewayDto> {
    return new WalletGatewayDto(this.toJSON());
  }
}
export abstract class WalletGatewayDtoFactory {
  abstract create(data: unknown): WalletGatewayDto;
}
/**
 * The base type definition for walletGatewayDto
 **/
export type WalletGatewayDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * Unique code identifying this gateway's GatewayAdapter implementation, e.g. "stripe", "manual", "onchain_eth".
   * @type {string}
   **/
  code: string;
  /**
   * Display name shown to wallet owners choosing a topup method.
   * @type {string}
   **/
  name: string;
  /**
   * Whether this gateway settles in fiat or crypto.
   * @type {string}
   **/
  kind: string;
  /**
   * Whether wallet owners can currently start a topup through this gateway.
   * @type {boolean}
   **/
  isActive: boolean;
  /**
   * Provider configuration. Must only hold references to secrets (e.g. a secrets-manager key name), never raw API keys/webhook secrets in plaintext.
   * @type {JSON}
   **/
  config: JSON;
  /**
   * JSON array of walletCurrency codes this gateway can top up (e.g. ["USD","EUR"]).
   * @type {JSON}
   **/
  supportedCurrencies: JSON;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WalletGatewayDtoType {}
