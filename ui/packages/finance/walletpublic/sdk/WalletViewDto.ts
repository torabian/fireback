import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for walletViewDto
 **/
export class WalletViewDto {
  /**
   * Unique id of the wallet.
   * @type {string}
   **/
  #uniqueId: string = "";
  /**
   * Unique id of the wallet.
   * @returns {string}
   **/
  get uniqueId() {
    return this.#uniqueId;
  }
  /**
   * Unique id of the wallet.
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
   * "user" or "workspace".
   * @type {string}
   **/
  #ownerType: string = "";
  /**
   * "user" or "workspace".
   * @returns {string}
   **/
  get ownerType() {
    return this.#ownerType;
  }
  /**
   * "user" or "workspace".
   * @type {string}
   **/
  set ownerType(value: string) {
    this.#ownerType = String(value);
  }
  setOwnerType(value: string) {
    this.ownerType = value;
    return this;
  }
  /**
   * Owning workspace id, set only when ownerType is "workspace".
   * @type {string}
   **/
  #workspaceId?: string | null = undefined;
  /**
   * Owning workspace id, set only when ownerType is "workspace".
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   * Owning workspace id, set only when ownerType is "workspace".
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
   * Currency code this wallet holds.
   * @type {string}
   **/
  #currency: string = "";
  /**
   * Currency code this wallet holds.
   * @returns {string}
   **/
  get currency() {
    return this.#currency;
  }
  /**
   * Currency code this wallet holds.
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
   * Current balance as a minor-units decimal string.
   * @type {string}
   **/
  #balance: string = "";
  /**
   * Current balance as a minor-units decimal string.
   * @returns {string}
   **/
  get balance() {
    return this.#balance;
  }
  /**
   * Current balance as a minor-units decimal string.
   * @type {string}
   **/
  set balance(value: string) {
    this.#balance = String(value);
  }
  setBalance(value: string) {
    this.balance = value;
    return this;
  }
  /**
   * "active", "frozen", or "closed".
   * @type {string}
   **/
  #status: string = "";
  /**
   * "active", "frozen", or "closed".
   * @returns {string}
   **/
  get status() {
    return this.#status;
  }
  /**
   * "active", "frozen", or "closed".
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
   * Owner-given nickname, if any.
   * @type {string}
   **/
  #label?: string | null = undefined;
  /**
   * Owner-given nickname, if any.
   * @returns {string}
   **/
  get label() {
    return this.#label;
  }
  /**
   * Owner-given nickname, if any.
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
   * Whether this is the owner's default wallet for its currency.
   * @type {boolean}
   **/
  #isDefault!: boolean;
  /**
   * Whether this is the owner's default wallet for its currency.
   * @returns {boolean}
   **/
  get isDefault() {
    return this.#isDefault;
  }
  /**
   * Whether this is the owner's default wallet for its currency.
   * @type {boolean}
   **/
  set isDefault(value: boolean) {
    this.#isDefault = Boolean(value);
  }
  setIsDefault(value: boolean) {
    this.isDefault = value;
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
    const d = data as Partial<WalletViewDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.ownerType !== undefined) {
      this.ownerType = d.ownerType;
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
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      ownerType: this.#ownerType,
      workspaceId: this.#workspaceId,
      currency: this.#currency,
      balance: this.#balance,
      status: this.#status,
      label: this.#label,
      isDefault: this.#isDefault,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      ownerType: "ownerType",
      workspaceId: "workspaceId",
      currency: "currency",
      balance: "balance",
      status: "status",
      label: "label",
      isDefault: "isDefault",
    };
  }
  /**
   * Creates an instance of WalletViewDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WalletViewDtoType) {
    return new WalletViewDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WalletViewDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WalletViewDtoType>) {
    return new WalletViewDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WalletViewDtoType>,
  ): InstanceType<typeof WalletViewDto> {
    return new WalletViewDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WalletViewDto> {
    return new WalletViewDto(this.toJSON());
  }
}
export abstract class WalletViewDtoFactory {
  abstract create(data: unknown): WalletViewDto;
}
/**
 * The base type definition for walletViewDto
 **/
export type WalletViewDtoType = {
  /**
   * Unique id of the wallet.
   * @type {string}
   **/
  uniqueId: string;
  /**
   * "user" or "workspace".
   * @type {string}
   **/
  ownerType: string;
  /**
   * Owning workspace id, set only when ownerType is "workspace".
   * @type {string}
   **/
  workspaceId?: string;
  /**
   * Currency code this wallet holds.
   * @type {string}
   **/
  currency: string;
  /**
   * Current balance as a minor-units decimal string.
   * @type {string}
   **/
  balance: string;
  /**
   * "active", "frozen", or "closed".
   * @type {string}
   **/
  status: string;
  /**
   * Owner-given nickname, if any.
   * @type {string}
   **/
  label?: string;
  /**
   * Whether this is the owner's default wallet for its currency.
   * @type {boolean}
   **/
  isDefault: boolean;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WalletViewDtoType {}
