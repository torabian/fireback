import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for messagingConfigOptionalDto
 **/
export class MessagingConfigOptionalDto {
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
   * The unique-id of the email provider service, which will be used to send the messages using it's service.
   * @type {string}
   **/
  #generalEmailProviderId?: string | null = undefined;
  /**
   * The unique-id of the email provider service, which will be used to send the messages using it's service.
   * @returns {string}
   **/
  get generalEmailProviderId() {
    return this.#generalEmailProviderId;
  }
  /**
   * The unique-id of the email provider service, which will be used to send the messages using it's service.
   * @type {string}
   **/
  set generalEmailProviderId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#generalEmailProviderId = correctType ? value : String(value);
  }
  setGeneralEmailProviderId(value: string | null | undefined) {
    this.generalEmailProviderId = value;
    return this;
  }
  /**
   * The unique-id of the general service which would be used to send text messages (sms).
   * @type {string}
   **/
  #generalGsmProviderId?: string | null = undefined;
  /**
   * The unique-id of the general service which would be used to send text messages (sms).
   * @returns {string}
   **/
  get generalGsmProviderId() {
    return this.#generalGsmProviderId;
  }
  /**
   * The unique-id of the general service which would be used to send text messages (sms).
   * @type {string}
   **/
  set generalGsmProviderId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#generalGsmProviderId = correctType ? value : String(value);
  }
  setGeneralGsmProviderId(value: string | null | undefined) {
    this.generalGsmProviderId = value;
    return this;
  }
  /**
   * The unique-id of the template used as default when a user is inviting a third-party into their own workspace.
   * @type {string}
   **/
  #inviteToWorkspaceContentId?: string | null = undefined;
  /**
   * The unique-id of the template used as default when a user is inviting a third-party into their own workspace.
   * @returns {string}
   **/
  get inviteToWorkspaceContentId() {
    return this.#inviteToWorkspaceContentId;
  }
  /**
   * The unique-id of the template used as default when a user is inviting a third-party into their own workspace.
   * @type {string}
   **/
  set inviteToWorkspaceContentId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#inviteToWorkspaceContentId = correctType ? value : String(value);
  }
  setInviteToWorkspaceContentId(value: string | null | undefined) {
    this.inviteToWorkspaceContentId = value;
    return this;
  }
  /**
   * The unique-id of the template used to fill the message for email one-time-password requests.
   * @type {string}
   **/
  #emailOtpContentId?: string | null = undefined;
  /**
   * The unique-id of the template used to fill the message for email one-time-password requests.
   * @returns {string}
   **/
  get emailOtpContentId() {
    return this.#emailOtpContentId;
  }
  /**
   * The unique-id of the template used to fill the message for email one-time-password requests.
   * @type {string}
   **/
  set emailOtpContentId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#emailOtpContentId = correctType ? value : String(value);
  }
  setEmailOtpContentId(value: string | null | undefined) {
    this.emailOtpContentId = value;
    return this;
  }
  /**
   * The unique-id of the template used for OTP text messages, including the one time password code.
   * @type {string}
   **/
  #smsOtpContentId?: string | null = undefined;
  /**
   * The unique-id of the template used for OTP text messages, including the one time password code.
   * @returns {string}
   **/
  get smsOtpContentId() {
    return this.#smsOtpContentId;
  }
  /**
   * The unique-id of the template used for OTP text messages, including the one time password code.
   * @type {string}
   **/
  set smsOtpContentId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#smsOtpContentId = correctType ? value : String(value);
  }
  setSmsOtpContentId(value: string | null | undefined) {
    this.smsOtpContentId = value;
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
    const d = data as Partial<MessagingConfigOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.generalEmailProviderId !== undefined) {
      this.generalEmailProviderId = d.generalEmailProviderId;
    }
    if (d.generalGsmProviderId !== undefined) {
      this.generalGsmProviderId = d.generalGsmProviderId;
    }
    if (d.inviteToWorkspaceContentId !== undefined) {
      this.inviteToWorkspaceContentId = d.inviteToWorkspaceContentId;
    }
    if (d.emailOtpContentId !== undefined) {
      this.emailOtpContentId = d.emailOtpContentId;
    }
    if (d.smsOtpContentId !== undefined) {
      this.smsOtpContentId = d.smsOtpContentId;
    }
    if (d.createdAt !== undefined) {
      this.createdAt = d.createdAt;
    }
    if (d.updatedAt !== undefined) {
      this.updatedAt = d.updatedAt;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      generalEmailProviderId: this.#generalEmailProviderId,
      generalGsmProviderId: this.#generalGsmProviderId,
      inviteToWorkspaceContentId: this.#inviteToWorkspaceContentId,
      emailOtpContentId: this.#emailOtpContentId,
      smsOtpContentId: this.#smsOtpContentId,
      createdAt: this.#createdAt,
      updatedAt: this.#updatedAt,
      workspaceId: this.#workspaceId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      generalEmailProviderId: "generalEmailProviderId",
      generalGsmProviderId: "generalGsmProviderId",
      inviteToWorkspaceContentId: "inviteToWorkspaceContentId",
      emailOtpContentId: "emailOtpContentId",
      smsOtpContentId: "smsOtpContentId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
      workspaceId: "workspaceId",
    };
  }
  /**
   * Creates an instance of MessagingConfigOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: MessagingConfigOptionalDtoType) {
    return new MessagingConfigOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of MessagingConfigOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<MessagingConfigOptionalDtoType>) {
    return new MessagingConfigOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<MessagingConfigOptionalDtoType>,
  ): InstanceType<typeof MessagingConfigOptionalDto> {
    return new MessagingConfigOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof MessagingConfigOptionalDto> {
    return new MessagingConfigOptionalDto(this.toJSON());
  }
}
export abstract class MessagingConfigOptionalDtoFactory {
  abstract create(data: unknown): MessagingConfigOptionalDto;
}
/**
 * The base type definition for messagingConfigOptionalDto
 **/
export type MessagingConfigOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * The unique-id of the email provider service, which will be used to send the messages using it's service.
   * @type {string}
   **/
  generalEmailProviderId?: string;
  /**
   * The unique-id of the general service which would be used to send text messages (sms).
   * @type {string}
   **/
  generalGsmProviderId?: string;
  /**
   * The unique-id of the template used as default when a user is inviting a third-party into their own workspace.
   * @type {string}
   **/
  inviteToWorkspaceContentId?: string;
  /**
   * The unique-id of the template used to fill the message for email one-time-password requests.
   * @type {string}
   **/
  emailOtpContentId?: string;
  /**
   * The unique-id of the template used for OTP text messages, including the one time password code.
   * @type {string}
   **/
  smsOtpContentId?: string;
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
  /**
   *
   * @type {string}
   **/
  workspaceId?: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace MessagingConfigOptionalDtoType {}
