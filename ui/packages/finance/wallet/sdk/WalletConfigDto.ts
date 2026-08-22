import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for walletConfigDto
 **/
export class WalletConfigDto {
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
   * Maximum number of wallets a single user may own in total, across all currencies. 0 means unlimited.
   * @type {number}
   **/
  #maxWalletsPerUser: number = 5;
  /**
   * Maximum number of wallets a single user may own in total, across all currencies. 0 means unlimited.
   * @returns {number}
   **/
  get maxWalletsPerUser() {
    return this.#maxWalletsPerUser;
  }
  /**
   * Maximum number of wallets a single user may own in total, across all currencies. 0 means unlimited.
   * @type {number}
   **/
  set maxWalletsPerUser(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#maxWalletsPerUser = parsedValue;
    }
  }
  setMaxWalletsPerUser(value: number) {
    this.maxWalletsPerUser = value;
    return this;
  }
  /**
   * Maximum number of wallets a single workspace may own in total, across all currencies. 0 means unlimited.
   * @type {number}
   **/
  #maxWalletsPerWorkspace: number = 5;
  /**
   * Maximum number of wallets a single workspace may own in total, across all currencies. 0 means unlimited.
   * @returns {number}
   **/
  get maxWalletsPerWorkspace() {
    return this.#maxWalletsPerWorkspace;
  }
  /**
   * Maximum number of wallets a single workspace may own in total, across all currencies. 0 means unlimited.
   * @type {number}
   **/
  set maxWalletsPerWorkspace(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#maxWalletsPerWorkspace = parsedValue;
    }
  }
  setMaxWalletsPerWorkspace(value: number) {
    this.maxWalletsPerWorkspace = value;
    return this;
  }
  /**
   * Maximum wallets a user may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerUser.
   * @type {number}
   **/
  #maxWalletsPerUserPerCurrency?: number | null = undefined;
  /**
   * Maximum wallets a user may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerUser.
   * @returns {number}
   **/
  get maxWalletsPerUserPerCurrency() {
    return this.#maxWalletsPerUserPerCurrency;
  }
  /**
   * Maximum wallets a user may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerUser.
   * @type {number}
   **/
  set maxWalletsPerUserPerCurrency(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#maxWalletsPerUserPerCurrency = parsedValue;
    }
  }
  setMaxWalletsPerUserPerCurrency(value: number | null | undefined) {
    this.maxWalletsPerUserPerCurrency = value;
    return this;
  }
  /**
   * Maximum wallets a workspace may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerWorkspace.
   * @type {number}
   **/
  #maxWalletsPerWorkspacePerCurrency?: number | null = undefined;
  /**
   * Maximum wallets a workspace may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerWorkspace.
   * @returns {number}
   **/
  get maxWalletsPerWorkspacePerCurrency() {
    return this.#maxWalletsPerWorkspacePerCurrency;
  }
  /**
   * Maximum wallets a workspace may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerWorkspace.
   * @type {number}
   **/
  set maxWalletsPerWorkspacePerCurrency(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#maxWalletsPerWorkspacePerCurrency = parsedValue;
    }
  }
  setMaxWalletsPerWorkspacePerCurrency(value: number | null | undefined) {
    this.maxWalletsPerWorkspacePerCurrency = value;
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
    const d = data as Partial<WalletConfigDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.maxWalletsPerUser !== undefined) {
      this.maxWalletsPerUser = d.maxWalletsPerUser;
    }
    if (d.maxWalletsPerWorkspace !== undefined) {
      this.maxWalletsPerWorkspace = d.maxWalletsPerWorkspace;
    }
    if (d.maxWalletsPerUserPerCurrency !== undefined) {
      this.maxWalletsPerUserPerCurrency = d.maxWalletsPerUserPerCurrency;
    }
    if (d.maxWalletsPerWorkspacePerCurrency !== undefined) {
      this.maxWalletsPerWorkspacePerCurrency =
        d.maxWalletsPerWorkspacePerCurrency;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      maxWalletsPerUser: this.#maxWalletsPerUser,
      maxWalletsPerWorkspace: this.#maxWalletsPerWorkspace,
      maxWalletsPerUserPerCurrency: this.#maxWalletsPerUserPerCurrency,
      maxWalletsPerWorkspacePerCurrency:
        this.#maxWalletsPerWorkspacePerCurrency,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      maxWalletsPerUser: "maxWalletsPerUser",
      maxWalletsPerWorkspace: "maxWalletsPerWorkspace",
      maxWalletsPerUserPerCurrency: "maxWalletsPerUserPerCurrency",
      maxWalletsPerWorkspacePerCurrency: "maxWalletsPerWorkspacePerCurrency",
    };
  }
  /**
   * Creates an instance of WalletConfigDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WalletConfigDtoType) {
    return new WalletConfigDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WalletConfigDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WalletConfigDtoType>) {
    return new WalletConfigDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WalletConfigDtoType>,
  ): InstanceType<typeof WalletConfigDto> {
    return new WalletConfigDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WalletConfigDto> {
    return new WalletConfigDto(this.toJSON());
  }
}
export abstract class WalletConfigDtoFactory {
  abstract create(data: unknown): WalletConfigDto;
}
/**
 * The base type definition for walletConfigDto
 **/
export type WalletConfigDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * Maximum number of wallets a single user may own in total, across all currencies. 0 means unlimited.
   * @type {number}
   **/
  maxWalletsPerUser: number;
  /**
   * Maximum number of wallets a single workspace may own in total, across all currencies. 0 means unlimited.
   * @type {number}
   **/
  maxWalletsPerWorkspace: number;
  /**
   * Maximum wallets a user may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerUser.
   * @type {number}
   **/
  maxWalletsPerUserPerCurrency?: number;
  /**
   * Maximum wallets a workspace may own in a single currency. Empty means no per-currency limit beyond maxWalletsPerWorkspace.
   * @type {number}
   **/
  maxWalletsPerWorkspacePerCurrency?: number;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WalletConfigDtoType {}
