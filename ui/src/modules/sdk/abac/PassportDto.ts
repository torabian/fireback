import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for passportDto
 **/
export class PassportDto {
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
   * When user creates account via oauth services such as google, it's essential to set the provider and do not allow passwordless logins if it's not via that specific provider.
   * @type {string}
   **/
  #thirdPartyVerifier: string = "";
  /**
   * When user creates account via oauth services such as google, it's essential to set the provider and do not allow passwordless logins if it's not via that specific provider.
   * @returns {string}
   **/
  get thirdPartyVerifier() {
    return this.#thirdPartyVerifier;
  }
  /**
   * When user creates account via oauth services such as google, it's essential to set the provider and do not allow passwordless logins if it's not via that specific provider.
   * @type {string}
   **/
  set thirdPartyVerifier(value: string) {
    this.#thirdPartyVerifier = String(value);
  }
  setThirdPartyVerifier(value: string) {
    this.thirdPartyVerifier = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #type: string = "";
  /**
   *
   * @returns {string}
   **/
  get type() {
    return this.#type;
  }
  /**
   *
   * @type {string}
   **/
  set type(value: string) {
    this.#type = String(value);
  }
  setType(value: string) {
    this.type = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #userId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   *
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
   *
   * @type {string}
   **/
  #value: string = "";
  /**
   *
   * @returns {string}
   **/
  get value() {
    return this.#value;
  }
  /**
   *
   * @type {string}
   **/
  set value(value: string) {
    this.#value = String(value);
  }
  setValue(value: string) {
    this.value = value;
    return this;
  }
  /**
   * Store the secret of 2FA using time based dual factor authentication here for this specific passport. If set, during authorization will be asked.
   * @type {string}
   **/
  #totpSecret: string = "";
  /**
   * Store the secret of 2FA using time based dual factor authentication here for this specific passport. If set, during authorization will be asked.
   * @returns {string}
   **/
  get totpSecret() {
    return this.#totpSecret;
  }
  /**
   * Store the secret of 2FA using time based dual factor authentication here for this specific passport. If set, during authorization will be asked.
   * @type {string}
   **/
  set totpSecret(value: string) {
    this.#totpSecret = String(value);
  }
  setTotpSecret(value: string) {
    this.totpSecret = value;
    return this;
  }
  /**
   * Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
   * @type {boolean}
   **/
  #totpConfirmed?: boolean | null = undefined;
  /**
   * Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
   * @returns {boolean}
   **/
  get totpConfirmed() {
    return this.#totpConfirmed;
  }
  /**
   * Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
   * @type {boolean}
   **/
  set totpConfirmed(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#totpConfirmed = correctType ? value : Boolean(value);
  }
  setTotpConfirmed(value: boolean | null | undefined) {
    this.totpConfirmed = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #password: string = "";
  /**
   *
   * @returns {string}
   **/
  get password() {
    return this.#password;
  }
  /**
   *
   * @type {string}
   **/
  set password(value: string) {
    this.#password = String(value);
  }
  setPassword(value: string) {
    this.password = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #confirmed?: boolean | null = undefined;
  /**
   *
   * @returns {boolean}
   **/
  get confirmed() {
    return this.#confirmed;
  }
  /**
   *
   * @type {boolean}
   **/
  set confirmed(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#confirmed = correctType ? value : Boolean(value);
  }
  setConfirmed(value: boolean | null | undefined) {
    this.confirmed = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #accessToken: string = "";
  /**
   *
   * @returns {string}
   **/
  get accessToken() {
    return this.#accessToken;
  }
  /**
   *
   * @type {string}
   **/
  set accessToken(value: string) {
    this.#accessToken = String(value);
  }
  setAccessToken(value: string) {
    this.accessToken = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #workspaceId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   *
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
   *
   * @type {PlainTime}
   **/
  #createdAt!: PlainTime;
  /**
   *
   * @returns {PlainTime}
   **/
  get createdAt() {
    return this.#createdAt;
  }
  /**
   *
   * @type {PlainTime}
   **/
  set createdAt(value: PlainTime) {
    this.#createdAt = value;
  }
  setCreatedAt(value: PlainTime) {
    this.createdAt = value;
    return this;
  }
  /**
   *
   * @type {PlainTime}
   **/
  #updatedAt!: PlainTime;
  /**
   *
   * @returns {PlainTime}
   **/
  get updatedAt() {
    return this.#updatedAt;
  }
  /**
   *
   * @type {PlainTime}
   **/
  set updatedAt(value: PlainTime) {
    this.#updatedAt = value;
  }
  setUpdatedAt(value: PlainTime) {
    this.updatedAt = value;
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
    const d = data as Partial<PassportDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.thirdPartyVerifier !== undefined) {
      this.thirdPartyVerifier = d.thirdPartyVerifier;
    }
    if (d.type !== undefined) {
      this.type = d.type;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.value !== undefined) {
      this.value = d.value;
    }
    if (d.totpSecret !== undefined) {
      this.totpSecret = d.totpSecret;
    }
    if (d.totpConfirmed !== undefined) {
      this.totpConfirmed = d.totpConfirmed;
    }
    if (d.password !== undefined) {
      this.password = d.password;
    }
    if (d.confirmed !== undefined) {
      this.confirmed = d.confirmed;
    }
    if (d.accessToken !== undefined) {
      this.accessToken = d.accessToken;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
    }
    if (d.createdAt !== undefined) {
      this.createdAt = d.createdAt;
    }
    if (d.updatedAt !== undefined) {
      this.updatedAt = d.updatedAt;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      thirdPartyVerifier: this.#thirdPartyVerifier,
      type: this.#type,
      userId: this.#userId,
      value: this.#value,
      totpSecret: this.#totpSecret,
      totpConfirmed: this.#totpConfirmed,
      password: this.#password,
      confirmed: this.#confirmed,
      accessToken: this.#accessToken,
      workspaceId: this.#workspaceId,
      createdAt: this.#createdAt,
      updatedAt: this.#updatedAt,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      thirdPartyVerifier: "thirdPartyVerifier",
      type: "type",
      userId: "userId",
      value: "value",
      totpSecret: "totpSecret",
      totpConfirmed: "totpConfirmed",
      password: "password",
      confirmed: "confirmed",
      accessToken: "accessToken",
      workspaceId: "workspaceId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of PassportDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: PassportDtoType) {
    return new PassportDto(possibleDtoObject);
  }
  /**
   * Creates an instance of PassportDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<PassportDtoType>) {
    return new PassportDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<PassportDtoType>,
  ): InstanceType<typeof PassportDto> {
    return new PassportDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof PassportDto> {
    return new PassportDto(this.toJSON());
  }
}
export abstract class PassportDtoFactory {
  abstract create(data: unknown): PassportDto;
}
/**
 * The base type definition for passportDto
 **/
export type PassportDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * When user creates account via oauth services such as google, it's essential to set the provider and do not allow passwordless logins if it's not via that specific provider.
   * @type {string}
   **/
  thirdPartyVerifier: string;
  /**
   *
   * @type {string}
   **/
  type: string;
  /**
   *
   * @type {string}
   **/
  userId?: string;
  /**
   *
   * @type {string}
   **/
  value: string;
  /**
   * Store the secret of 2FA using time based dual factor authentication here for this specific passport. If set, during authorization will be asked.
   * @type {string}
   **/
  totpSecret: string;
  /**
   * Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
   * @type {boolean}
   **/
  totpConfirmed?: boolean;
  /**
   *
   * @type {string}
   **/
  password: string;
  /**
   *
   * @type {boolean}
   **/
  confirmed?: boolean;
  /**
   *
   * @type {string}
   **/
  accessToken: string;
  /**
   *
   * @type {string}
   **/
  workspaceId?: string;
  /**
   *
   * @type {PlainTime}
   **/
  createdAt: PlainTime;
  /**
   *
   * @type {PlainTime}
   **/
  updatedAt: PlainTime;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PassportDtoType {}
