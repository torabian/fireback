import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for walletProviderConfigOptionalDto
 **/
export class WalletProviderConfigOptionalDto {
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
   * Which GatewayAdapter this configures - must match a registered adapter's Code(), e.g. "stripe", "przelewy24", "blik", "zarinpal". Free-form (not a closed oneof) so a new provider package doesn't require a schema change here.
   * @type {string}
   **/
  #providerType?: string | null = undefined;
  /**
   * Which GatewayAdapter this configures - must match a registered adapter's Code(), e.g. "stripe", "przelewy24", "blik", "zarinpal". Free-form (not a closed oneof) so a new provider package doesn't require a schema change here.
   * @returns {string}
   **/
  get providerType() {
    return this.#providerType;
  }
  /**
   * Which GatewayAdapter this configures - must match a registered adapter's Code(), e.g. "stripe", "przelewy24", "blik", "zarinpal". Free-form (not a closed oneof) so a new provider package doesn't require a schema change here.
   * @type {string}
   **/
  set providerType(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#providerType = correctType ? value : String(value);
  }
  setProviderType(value: string | null | undefined) {
    this.providerType = value;
    return this;
  }
  /**
   * Which region this config applies to - an ISO-3166-ish code (e.g. "PL", "IR", "US") or "global", meaning every region. Combined with providerType, must be unique: you can't have two rows configuring the same provider for the same region.
   * @type {string}
   **/
  #region?: string | null = "global";
  /**
   * Which region this config applies to - an ISO-3166-ish code (e.g. "PL", "IR", "US") or "global", meaning every region. Combined with providerType, must be unique: you can't have two rows configuring the same provider for the same region.
   * @returns {string}
   **/
  get region() {
    return this.#region;
  }
  /**
   * Which region this config applies to - an ISO-3166-ish code (e.g. "PL", "IR", "US") or "global", meaning every region. Combined with providerType, must be unique: you can't have two rows configuring the same provider for the same region.
   * @type {string}
   **/
  set region(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#region = correctType ? value : String(value);
  }
  setRegion(value: string | null | undefined) {
    this.region = value;
    return this;
  }
  /**
   * Whether this provider+region combination is currently usable. Toggled through the regular update action - there is no dedicated enable/disable endpoint.
   * @type {boolean}
   **/
  #isEnabled?: boolean | null = false;
  /**
   * Whether this provider+region combination is currently usable. Toggled through the regular update action - there is no dedicated enable/disable endpoint.
   * @returns {boolean}
   **/
  get isEnabled() {
    return this.#isEnabled;
  }
  /**
   * Whether this provider+region combination is currently usable. Toggled through the regular update action - there is no dedicated enable/disable endpoint.
   * @type {boolean}
   **/
  set isEnabled(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#isEnabled = correctType ? value : Boolean(value);
  }
  setIsEnabled(value: boolean | null | undefined) {
    this.isEnabled = value;
    return this;
  }
  /**
   * Provider-specific, non-secret settings (e.g. preferred/allowed currencies, routing hints, sandbox overrides). Must never hold raw API keys/secrets - those stay in each provider package's own environment variables (see modules/wallet/providers/*'s doc comments), matching walletGateway.config's same convention.
   * @type {JSON}
   **/
  #config!: JSON;
  /**
   * Provider-specific, non-secret settings (e.g. preferred/allowed currencies, routing hints, sandbox overrides). Must never hold raw API keys/secrets - those stay in each provider package's own environment variables (see modules/wallet/providers/*'s doc comments), matching walletGateway.config's same convention.
   * @returns {JSON}
   **/
  get config() {
    return this.#config;
  }
  /**
   * Provider-specific, non-secret settings (e.g. preferred/allowed currencies, routing hints, sandbox overrides). Must never hold raw API keys/secrets - those stay in each provider package's own environment variables (see modules/wallet/providers/*'s doc comments), matching walletGateway.config's same convention.
   * @type {JSON}
   **/
  set config(value: JSON) {
    this.#config = value;
  }
  setConfig(value: JSON) {
    this.config = value;
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
    const d = data as Partial<WalletProviderConfigOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.providerType !== undefined) {
      this.providerType = d.providerType;
    }
    if (d.region !== undefined) {
      this.region = d.region;
    }
    if (d.isEnabled !== undefined) {
      this.isEnabled = d.isEnabled;
    }
    if (d.config !== undefined) {
      this.config = d.config;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      providerType: this.#providerType,
      region: this.#region,
      isEnabled: this.#isEnabled,
      config: this.#config,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      providerType: "providerType",
      region: "region",
      isEnabled: "isEnabled",
      config: "config",
    };
  }
  /**
   * Creates an instance of WalletProviderConfigOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WalletProviderConfigOptionalDtoType) {
    return new WalletProviderConfigOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WalletProviderConfigOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<WalletProviderConfigOptionalDtoType>,
  ) {
    return new WalletProviderConfigOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WalletProviderConfigOptionalDtoType>,
  ): InstanceType<typeof WalletProviderConfigOptionalDto> {
    return new WalletProviderConfigOptionalDto({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof WalletProviderConfigOptionalDto> {
    return new WalletProviderConfigOptionalDto(this.toJSON());
  }
}
export abstract class WalletProviderConfigOptionalDtoFactory {
  abstract create(data: unknown): WalletProviderConfigOptionalDto;
}
/**
 * The base type definition for walletProviderConfigOptionalDto
 **/
export type WalletProviderConfigOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * Which GatewayAdapter this configures - must match a registered adapter's Code(), e.g. "stripe", "przelewy24", "blik", "zarinpal". Free-form (not a closed oneof) so a new provider package doesn't require a schema change here.
   * @type {string}
   **/
  providerType?: string;
  /**
   * Which region this config applies to - an ISO-3166-ish code (e.g. "PL", "IR", "US") or "global", meaning every region. Combined with providerType, must be unique: you can't have two rows configuring the same provider for the same region.
   * @type {string}
   **/
  region?: string;
  /**
   * Whether this provider+region combination is currently usable. Toggled through the regular update action - there is no dedicated enable/disable endpoint.
   * @type {boolean}
   **/
  isEnabled?: boolean;
  /**
   * Provider-specific, non-secret settings (e.g. preferred/allowed currencies, routing hints, sandbox overrides). Must never hold raw API keys/secrets - those stay in each provider package's own environment variables (see modules/wallet/providers/*'s doc comments), matching walletGateway.config's same convention.
   * @type {JSON}
   **/
  config: JSON;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WalletProviderConfigOptionalDtoType {}
