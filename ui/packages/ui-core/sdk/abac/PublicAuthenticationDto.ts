import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for publicAuthenticationDto
 **/
export class PublicAuthenticationDto {
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
  #totpSecret: string = "";
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
  set totpSecret(value: string) {
    this.#totpSecret = String(value);
  }
  setTotpSecret(value: string) {
    this.totpSecret = value;
    return this;
  }
  /**
   * The url which will be converted into QR code on client side to scan
   * @type {string}
   **/
  #totpLink: string = "";
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
  set totpLink(value: string) {
    this.#totpLink = String(value);
  }
  setTotpLink(value: string) {
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
  #sessionSecret: string = "";
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
  set sessionSecret(value: string) {
    this.#sessionSecret = String(value);
  }
  setSessionSecret(value: string) {
    this.sessionSecret = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #passportValue: string = "";
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
  set passportValue(value: string) {
    this.#passportValue = String(value);
  }
  setPassportValue(value: string) {
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
  #status: string = "";
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
  set status(value: string) {
    this.#status = String(value);
  }
  setStatus(value: string) {
    this.status = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #blockedUntil: number = 0;
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
  set blockedUntil(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#blockedUntil = parsedValue;
    }
  }
  setBlockedUntil(value: number) {
    this.blockedUntil = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #otp: string = "";
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
  set otp(value: string) {
    this.#otp = String(value);
  }
  setOtp(value: string) {
    this.otp = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #recoveryAbsoluteUrl: string = "";
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
  set recoveryAbsoluteUrl(value: string) {
    this.#recoveryAbsoluteUrl = String(value);
  }
  setRecoveryAbsoluteUrl(value: string) {
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
    const d = data as Partial<PublicAuthenticationDto>;
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
   * Creates an instance of PublicAuthenticationDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: PublicAuthenticationDtoType) {
    return new PublicAuthenticationDto(possibleDtoObject);
  }
  /**
   * Creates an instance of PublicAuthenticationDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<PublicAuthenticationDtoType>) {
    return new PublicAuthenticationDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<PublicAuthenticationDtoType>,
  ): InstanceType<typeof PublicAuthenticationDto> {
    return new PublicAuthenticationDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof PublicAuthenticationDto> {
    return new PublicAuthenticationDto(this.toJSON());
  }
}
export abstract class PublicAuthenticationDtoFactory {
  abstract create(data: unknown): PublicAuthenticationDto;
}
/**
 * The base type definition for publicAuthenticationDto
 **/
export type PublicAuthenticationDtoType = {
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
  totpSecret: string;
  /**
   * The url which will be converted into QR code on client side to scan
   * @type {string}
   **/
  totpLink: string;
  /**
   * The unique-id of the passport this record belongs to.
   * @type {string}
   **/
  passportId?: string;
  /**
   * This is a long hash generated and will be used to authenticate user after he confirmed the otp to finish the signup process and add more information before creating an account
   * @type {string}
   **/
  sessionSecret: string;
  /**
   *
   * @type {string}
   **/
  passportValue: string;
  /**
   *
   * @type {boolean}
   **/
  isInCreationProcess?: boolean;
  /**
   *
   * @type {string}
   **/
  status: string;
  /**
   *
   * @type {number}
   **/
  blockedUntil: number;
  /**
   *
   * @type {string}
   **/
  otp: string;
  /**
   *
   * @type {string}
   **/
  recoveryAbsoluteUrl: string;
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
export namespace PublicAuthenticationDtoType {}
