import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for publicAuthenticationOptionalDto
 **/
export class PublicAuthenticationOptionalDto {
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
   * The unique-id of the user which this record belongs to.
   * @type {string}
   **/
  #userId?: string | null = undefined;
  /**
   * The unique-id of the user which this record belongs to.
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * The unique-id of the user which this record belongs to.
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
   * If the application requires totp dual factor upon account creation, we create a secret here and pass the link
   * @type {string}
   **/
  #totpSecret?: string | null = undefined;
  /**
   * If the application requires totp dual factor upon account creation, we create a secret here and pass the link
   * @returns {string}
   **/
  get totpSecret() {
    return this.#totpSecret;
  }
  /**
   * If the application requires totp dual factor upon account creation, we create a secret here and pass the link
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
   * The url which will be converted into QR code on client side to scan
   * @type {string}
   **/
  #totpLink?: string | null = undefined;
  /**
   * The url which will be converted into QR code on client side to scan
   * @returns {string}
   **/
  get totpLink() {
    return this.#totpLink;
  }
  /**
   * The url which will be converted into QR code on client side to scan
   * @type {string}
   **/
  set totpLink(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#totpLink = correctType ? value : String(value);
  }
  setTotpLink(value: string | null | undefined) {
    this.totpLink = value;
    return this;
  }
  /**
   * The unique-id of the passport this record belongs to.
   * @type {string}
   **/
  #passportId?: string | null = undefined;
  /**
   * The unique-id of the passport this record belongs to.
   * @returns {string}
   **/
  get passportId() {
    return this.#passportId;
  }
  /**
   * The unique-id of the passport this record belongs to.
   * @type {string}
   **/
  set passportId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#passportId = correctType ? value : String(value);
  }
  setPassportId(value: string | null | undefined) {
    this.passportId = value;
    return this;
  }
  /**
   * This is a long hash generated and will be used to authenticate user after he confirmed the otp to finish the signup process and add more information before creating an account
   * @type {string}
   **/
  #sessionSecret?: string | null = undefined;
  /**
   * This is a long hash generated and will be used to authenticate user after he confirmed the otp to finish the signup process and add more information before creating an account
   * @returns {string}
   **/
  get sessionSecret() {
    return this.#sessionSecret;
  }
  /**
   * This is a long hash generated and will be used to authenticate user after he confirmed the otp to finish the signup process and add more information before creating an account
   * @type {string}
   **/
  set sessionSecret(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#sessionSecret = correctType ? value : String(value);
  }
  setSessionSecret(value: string | null | undefined) {
    this.sessionSecret = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #passportValue?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get passportValue() {
    return this.#passportValue;
  }
  /**
   *
   * @type {string}
   **/
  set passportValue(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#passportValue = correctType ? value : String(value);
  }
  setPassportValue(value: string | null | undefined) {
    this.passportValue = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #isInCreationProcess?: boolean | null = undefined;
  /**
   *
   * @returns {boolean}
   **/
  get isInCreationProcess() {
    return this.#isInCreationProcess;
  }
  /**
   *
   * @type {boolean}
   **/
  set isInCreationProcess(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#isInCreationProcess = correctType ? value : Boolean(value);
  }
  setIsInCreationProcess(value: boolean | null | undefined) {
    this.isInCreationProcess = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #status?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get status() {
    return this.#status;
  }
  /**
   *
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
   *
   * @type {number}
   **/
  #blockedUntil?: number | null = undefined;
  /**
   *
   * @returns {number}
   **/
  get blockedUntil() {
    return this.#blockedUntil;
  }
  /**
   *
   * @type {number}
   **/
  set blockedUntil(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#blockedUntil = parsedValue;
    }
  }
  setBlockedUntil(value: number | null | undefined) {
    this.blockedUntil = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #otp?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get otp() {
    return this.#otp;
  }
  /**
   *
   * @type {string}
   **/
  set otp(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#otp = correctType ? value : String(value);
  }
  setOtp(value: string | null | undefined) {
    this.otp = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #recoveryAbsoluteUrl?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get recoveryAbsoluteUrl() {
    return this.#recoveryAbsoluteUrl;
  }
  /**
   *
   * @type {string}
   **/
  set recoveryAbsoluteUrl(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#recoveryAbsoluteUrl = correctType ? value : String(value);
  }
  setRecoveryAbsoluteUrl(value: string | null | undefined) {
    this.recoveryAbsoluteUrl = value;
    return this;
  }
  /**
   * The unique-id of the workspace which content belongs to.
   * @type {string}
   **/
  #workspaceId?: string | null = undefined;
  /**
   * The unique-id of the workspace which content belongs to.
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   * The unique-id of the workspace which content belongs to.
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
    const d = data as Partial<PublicAuthenticationOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.totpSecret !== undefined) {
      this.totpSecret = d.totpSecret;
    }
    if (d.totpLink !== undefined) {
      this.totpLink = d.totpLink;
    }
    if (d.passportId !== undefined) {
      this.passportId = d.passportId;
    }
    if (d.sessionSecret !== undefined) {
      this.sessionSecret = d.sessionSecret;
    }
    if (d.passportValue !== undefined) {
      this.passportValue = d.passportValue;
    }
    if (d.isInCreationProcess !== undefined) {
      this.isInCreationProcess = d.isInCreationProcess;
    }
    if (d.status !== undefined) {
      this.status = d.status;
    }
    if (d.blockedUntil !== undefined) {
      this.blockedUntil = d.blockedUntil;
    }
    if (d.otp !== undefined) {
      this.otp = d.otp;
    }
    if (d.recoveryAbsoluteUrl !== undefined) {
      this.recoveryAbsoluteUrl = d.recoveryAbsoluteUrl;
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
      userId: this.#userId,
      totpSecret: this.#totpSecret,
      totpLink: this.#totpLink,
      passportId: this.#passportId,
      sessionSecret: this.#sessionSecret,
      passportValue: this.#passportValue,
      isInCreationProcess: this.#isInCreationProcess,
      status: this.#status,
      blockedUntil: this.#blockedUntil,
      otp: this.#otp,
      recoveryAbsoluteUrl: this.#recoveryAbsoluteUrl,
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
      userId: "userId",
      totpSecret: "totpSecret",
      totpLink: "totpLink",
      passportId: "passportId",
      sessionSecret: "sessionSecret",
      passportValue: "passportValue",
      isInCreationProcess: "isInCreationProcess",
      status: "status",
      blockedUntil: "blockedUntil",
      otp: "otp",
      recoveryAbsoluteUrl: "recoveryAbsoluteUrl",
      workspaceId: "workspaceId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of PublicAuthenticationOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: PublicAuthenticationOptionalDtoType) {
    return new PublicAuthenticationOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of PublicAuthenticationOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<PublicAuthenticationOptionalDtoType>,
  ) {
    return new PublicAuthenticationOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<PublicAuthenticationOptionalDtoType>,
  ): InstanceType<typeof PublicAuthenticationOptionalDto> {
    return new PublicAuthenticationOptionalDto({
      ...this.toJSON(),
      ...partial,
    });
  }
  clone(): InstanceType<typeof PublicAuthenticationOptionalDto> {
    return new PublicAuthenticationOptionalDto(this.toJSON());
  }
}
export abstract class PublicAuthenticationOptionalDtoFactory {
  abstract create(data: unknown): PublicAuthenticationOptionalDto;
}
/**
 * The base type definition for publicAuthenticationOptionalDto
 **/
export type PublicAuthenticationOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * The unique-id of the user which this record belongs to.
   * @type {string}
   **/
  userId?: string;
  /**
   * If the application requires totp dual factor upon account creation, we create a secret here and pass the link
   * @type {string}
   **/
  totpSecret?: string;
  /**
   * The url which will be converted into QR code on client side to scan
   * @type {string}
   **/
  totpLink?: string;
  /**
   * The unique-id of the passport this record belongs to.
   * @type {string}
   **/
  passportId?: string;
  /**
   * This is a long hash generated and will be used to authenticate user after he confirmed the otp to finish the signup process and add more information before creating an account
   * @type {string}
   **/
  sessionSecret?: string;
  /**
   *
   * @type {string}
   **/
  passportValue?: string;
  /**
   *
   * @type {boolean}
   **/
  isInCreationProcess?: boolean;
  /**
   *
   * @type {string}
   **/
  status?: string;
  /**
   *
   * @type {number}
   **/
  blockedUntil?: number;
  /**
   *
   * @type {string}
   **/
  otp?: string;
  /**
   *
   * @type {string}
   **/
  recoveryAbsoluteUrl?: string;
  /**
   * The unique-id of the workspace which content belongs to.
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
export namespace PublicAuthenticationOptionalDtoType {}
