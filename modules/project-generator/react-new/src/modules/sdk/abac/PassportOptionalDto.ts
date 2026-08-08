import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for passportOptionalDto
 **/
export class PassportOptionalDto {
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
  #thirdPartyVerifier?: string | null = undefined;
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
  set thirdPartyVerifier(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#thirdPartyVerifier = correctType ? value : String(value);
  }
  setThirdPartyVerifier(value: string | null | undefined) {
    this.thirdPartyVerifier = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #type?: string | null = undefined;
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
  set type(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#type = correctType ? value : String(value);
  }
  setType(value: string | null | undefined) {
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
  #value?: string | null = undefined;
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
  set value(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#value = correctType ? value : String(value);
  }
  setValue(value: string | null | undefined) {
    this.value = value;
    return this;
  }
  /**
   * Store the secret of 2FA using time based dual factor authentication here for this specific passport. If set, during authorization will be asked.
   * @type {string}
   **/
  #totpSecret?: string | null = undefined;
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
  set totpSecret(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#totpSecret = correctType ? value : String(value);
  }
  setTotpSecret(value: string | null | undefined) {
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
  #password?: string | null = undefined;
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
  set password(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#password = correctType ? value : String(value);
  }
  setPassword(value: string | null | undefined) {
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
  #accessToken?: string | null = undefined;
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
  set accessToken(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#accessToken = correctType ? value : String(value);
  }
  setAccessToken(value: string | null | undefined) {
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
    const d = data as Partial<PassportOptionalDto>;
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
   * Creates an instance of PassportOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: PassportOptionalDtoType) {
    return new PassportOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of PassportOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<PassportOptionalDtoType>) {
    return new PassportOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<PassportOptionalDtoType>,
  ): InstanceType<typeof PassportOptionalDto> {
    return new PassportOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof PassportOptionalDto> {
    return new PassportOptionalDto(this.toJSON());
  }
}
export abstract class PassportOptionalDtoFactory {
  abstract create(data: unknown): PassportOptionalDto;
}
/**
 * The base type definition for passportOptionalDto
 **/
export type PassportOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * When user creates account via oauth services such as google, it's essential to set the provider and do not allow passwordless logins if it's not via that specific provider.
   * @type {string}
   **/
  thirdPartyVerifier?: string;
  /**
   *
   * @type {string}
   **/
  type?: string;
  /**
   *
   * @type {string}
   **/
  userId?: string;
  /**
   *
   * @type {string}
   **/
  value?: string;
  /**
   * Store the secret of 2FA using time based dual factor authentication here for this specific passport. If set, during authorization will be asked.
   * @type {string}
   **/
  totpSecret?: string;
  /**
   * Regardless of the secret, user needs to confirm his secret. There is an extra action to confirm user totp, could be used after signup or prior to login.
   * @type {boolean}
   **/
  totpConfirmed?: boolean;
  /**
   *
   * @type {string}
   **/
  password?: string;
  /**
   *
   * @type {boolean}
   **/
  confirmed?: boolean;
  /**
   *
   * @type {string}
   **/
  accessToken?: string;
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
export namespace PassportOptionalDtoType {}
