import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for workspaceInviteDto
 **/
export class WorkspaceInviteDto {
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
   * A long hash to get the user into the confirm or signup page without sending the email or phone number, for example if an administrator wants to copy the link.
   * @type {string}
   **/
  #publicKey: string = "";
  /**
   * A long hash to get the user into the confirm or signup page without sending the email or phone number, for example if an administrator wants to copy the link.
   * @returns {string}
   **/
  get publicKey() {
    return this.#publicKey;
  }
  /**
   * A long hash to get the user into the confirm or signup page without sending the email or phone number, for example if an administrator wants to copy the link.
   * @type {string}
   **/
  set publicKey(value: string) {
    this.#publicKey = String(value);
  }
  setPublicKey(value: string) {
    this.publicKey = value;
    return this;
  }
  /**
   * The content that user will receive to understand the reason of the letter.
   * @type {string}
   **/
  #coverLetter: string = "";
  /**
   * The content that user will receive to understand the reason of the letter.
   * @returns {string}
   **/
  get coverLetter() {
    return this.#coverLetter;
  }
  /**
   * The content that user will receive to understand the reason of the letter.
   * @type {string}
   **/
  set coverLetter(value: string) {
    this.#coverLetter = String(value);
  }
  setCoverLetter(value: string) {
    this.coverLetter = value;
    return this;
  }
  /**
   * If the invited person has a different language, then you can define that so the interface for him will be automatically translated.
   * @type {string}
   **/
  #targetUserLocale: string = "";
  /**
   * If the invited person has a different language, then you can define that so the interface for him will be automatically translated.
   * @returns {string}
   **/
  get targetUserLocale() {
    return this.#targetUserLocale;
  }
  /**
   * If the invited person has a different language, then you can define that so the interface for him will be automatically translated.
   * @type {string}
   **/
  set targetUserLocale(value: string) {
    this.#targetUserLocale = String(value);
  }
  setTargetUserLocale(value: string) {
    this.targetUserLocale = value;
    return this;
  }
  /**
   * The email address of the person which is invited.
   * @type {string}
   **/
  #email: string = "";
  /**
   * The email address of the person which is invited.
   * @returns {string}
   **/
  get email() {
    return this.#email;
  }
  /**
   * The email address of the person which is invited.
   * @type {string}
   **/
  set email(value: string) {
    this.#email = String(value);
  }
  setEmail(value: string) {
    this.email = value;
    return this;
  }
  /**
   * The phone number of the person which is invited.
   * @type {string}
   **/
  #phonenumber: string = "";
  /**
   * The phone number of the person which is invited.
   * @returns {string}
   **/
  get phonenumber() {
    return this.#phonenumber;
  }
  /**
   * The phone number of the person which is invited.
   * @type {string}
   **/
  set phonenumber(value: string) {
    this.#phonenumber = String(value);
  }
  setPhonenumber(value: string) {
    this.phonenumber = value;
    return this;
  }
  /**
   * The unique-id of the workspace which user is being invited to.
   * @type {string}
   **/
  #workspaceId?: string | null = undefined;
  /**
   * The unique-id of the workspace which user is being invited to.
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   * The unique-id of the workspace which user is being invited to.
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
   * First name of the person which is invited
   * @type {string}
   **/
  #firstName: string = "";
  /**
   * First name of the person which is invited
   * @returns {string}
   **/
  get firstName() {
    return this.#firstName;
  }
  /**
   * First name of the person which is invited
   * @type {string}
   **/
  set firstName(value: string) {
    this.#firstName = String(value);
  }
  setFirstName(value: string) {
    this.firstName = value;
    return this;
  }
  /**
   * Last name of the person which is invited.
   * @type {string}
   **/
  #lastName: string = "";
  /**
   * Last name of the person which is invited.
   * @returns {string}
   **/
  get lastName() {
    return this.#lastName;
  }
  /**
   * Last name of the person which is invited.
   * @type {string}
   **/
  set lastName(value: string) {
    this.#lastName = String(value);
  }
  setLastName(value: string) {
    this.lastName = value;
    return this;
  }
  /**
   * If forced, the email address cannot be changed by the user which has been invited.
   * @type {boolean}
   **/
  #forceEmailAddress?: boolean | null = undefined;
  /**
   * If forced, the email address cannot be changed by the user which has been invited.
   * @returns {boolean}
   **/
  get forceEmailAddress() {
    return this.#forceEmailAddress;
  }
  /**
   * If forced, the email address cannot be changed by the user which has been invited.
   * @type {boolean}
   **/
  set forceEmailAddress(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#forceEmailAddress = correctType ? value : Boolean(value);
  }
  setForceEmailAddress(value: boolean | null | undefined) {
    this.forceEmailAddress = value;
    return this;
  }
  /**
   * If forced, user cannot change the phone number and needs to complete signup.
   * @type {boolean}
   **/
  #forcePhoneNumber?: boolean | null = undefined;
  /**
   * If forced, user cannot change the phone number and needs to complete signup.
   * @returns {boolean}
   **/
  get forcePhoneNumber() {
    return this.#forcePhoneNumber;
  }
  /**
   * If forced, user cannot change the phone number and needs to complete signup.
   * @type {boolean}
   **/
  set forcePhoneNumber(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#forcePhoneNumber = correctType ? value : Boolean(value);
  }
  setForcePhoneNumber(value: boolean | null | undefined) {
    this.forcePhoneNumber = value;
    return this;
  }
  /**
   * The role which invitee get if they accept the request.
   * @type {string}
   **/
  #roleId?: string | null = undefined;
  /**
   * The role which invitee get if they accept the request.
   * @returns {string}
   **/
  get roleId() {
    return this.#roleId;
  }
  /**
   * The role which invitee get if they accept the request.
   * @type {string}
   **/
  set roleId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#roleId = correctType ? value : String(value);
  }
  setRoleId(value: string | null | undefined) {
    this.roleId = value;
    return this;
  }
  /**
   * The unique-id of the user which created/owns the record.
   * @type {string}
   **/
  #userId?: string | null = undefined;
  /**
   * The unique-id of the user which created/owns the record.
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * The unique-id of the user which created/owns the record.
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
    const d = data as Partial<WorkspaceInviteDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.publicKey !== undefined) {
      this.publicKey = d.publicKey;
    }
    if (d.coverLetter !== undefined) {
      this.coverLetter = d.coverLetter;
    }
    if (d.targetUserLocale !== undefined) {
      this.targetUserLocale = d.targetUserLocale;
    }
    if (d.email !== undefined) {
      this.email = d.email;
    }
    if (d.phonenumber !== undefined) {
      this.phonenumber = d.phonenumber;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
    }
    if (d.firstName !== undefined) {
      this.firstName = d.firstName;
    }
    if (d.lastName !== undefined) {
      this.lastName = d.lastName;
    }
    if (d.forceEmailAddress !== undefined) {
      this.forceEmailAddress = d.forceEmailAddress;
    }
    if (d.forcePhoneNumber !== undefined) {
      this.forcePhoneNumber = d.forcePhoneNumber;
    }
    if (d.roleId !== undefined) {
      this.roleId = d.roleId;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
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
      publicKey: this.#publicKey,
      coverLetter: this.#coverLetter,
      targetUserLocale: this.#targetUserLocale,
      email: this.#email,
      phonenumber: this.#phonenumber,
      workspaceId: this.#workspaceId,
      firstName: this.#firstName,
      lastName: this.#lastName,
      forceEmailAddress: this.#forceEmailAddress,
      forcePhoneNumber: this.#forcePhoneNumber,
      roleId: this.#roleId,
      userId: this.#userId,
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
      publicKey: "publicKey",
      coverLetter: "coverLetter",
      targetUserLocale: "targetUserLocale",
      email: "email",
      phonenumber: "phonenumber",
      workspaceId: "workspaceId",
      firstName: "firstName",
      lastName: "lastName",
      forceEmailAddress: "forceEmailAddress",
      forcePhoneNumber: "forcePhoneNumber",
      roleId: "roleId",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of WorkspaceInviteDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WorkspaceInviteDtoType) {
    return new WorkspaceInviteDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WorkspaceInviteDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WorkspaceInviteDtoType>) {
    return new WorkspaceInviteDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WorkspaceInviteDtoType>,
  ): InstanceType<typeof WorkspaceInviteDto> {
    return new WorkspaceInviteDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WorkspaceInviteDto> {
    return new WorkspaceInviteDto(this.toJSON());
  }
}
export abstract class WorkspaceInviteDtoFactory {
  abstract create(data: unknown): WorkspaceInviteDto;
}
/**
 * The base type definition for workspaceInviteDto
 **/
export type WorkspaceInviteDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * A long hash to get the user into the confirm or signup page without sending the email or phone number, for example if an administrator wants to copy the link.
   * @type {string}
   **/
  publicKey: string;
  /**
   * The content that user will receive to understand the reason of the letter.
   * @type {string}
   **/
  coverLetter: string;
  /**
   * If the invited person has a different language, then you can define that so the interface for him will be automatically translated.
   * @type {string}
   **/
  targetUserLocale: string;
  /**
   * The email address of the person which is invited.
   * @type {string}
   **/
  email: string;
  /**
   * The phone number of the person which is invited.
   * @type {string}
   **/
  phonenumber: string;
  /**
   * The unique-id of the workspace which user is being invited to.
   * @type {string}
   **/
  workspaceId?: string;
  /**
   * First name of the person which is invited
   * @type {string}
   **/
  firstName: string;
  /**
   * Last name of the person which is invited.
   * @type {string}
   **/
  lastName: string;
  /**
   * If forced, the email address cannot be changed by the user which has been invited.
   * @type {boolean}
   **/
  forceEmailAddress?: boolean;
  /**
   * If forced, user cannot change the phone number and needs to complete signup.
   * @type {boolean}
   **/
  forcePhoneNumber?: boolean;
  /**
   * The role which invitee get if they accept the request.
   * @type {string}
   **/
  roleId?: string;
  /**
   * The unique-id of the user which created/owns the record.
   * @type {string}
   **/
  userId?: string;
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
export namespace WorkspaceInviteDtoType {}
