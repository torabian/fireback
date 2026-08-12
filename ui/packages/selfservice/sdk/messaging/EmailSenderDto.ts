import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for emailSenderDto
 **/
export class EmailSenderDto {
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
   *
   * @type {string}
   **/
  #fromName: string = "";
  /**
   *
   * @returns {string}
   **/
  get fromName() {
    return this.#fromName;
  }
  /**
   *
   * @type {string}
   **/
  set fromName(value: string) {
    this.#fromName = String(value);
  }
  setFromName(value: string) {
    this.fromName = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #fromEmailAddress: string = "";
  /**
   *
   * @returns {string}
   **/
  get fromEmailAddress() {
    return this.#fromEmailAddress;
  }
  /**
   *
   * @type {string}
   **/
  set fromEmailAddress(value: string) {
    this.#fromEmailAddress = String(value);
  }
  setFromEmailAddress(value: string) {
    this.fromEmailAddress = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #replyTo: string = "";
  /**
   *
   * @returns {string}
   **/
  get replyTo() {
    return this.#replyTo;
  }
  /**
   *
   * @type {string}
   **/
  set replyTo(value: string) {
    this.#replyTo = String(value);
  }
  setReplyTo(value: string) {
    this.replyTo = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #nickName: string = "";
  /**
   *
   * @returns {string}
   **/
  get nickName() {
    return this.#nickName;
  }
  /**
   *
   * @type {string}
   **/
  set nickName(value: string) {
    this.#nickName = String(value);
  }
  setNickName(value: string) {
    this.nickName = value;
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
    const d = data as Partial<EmailSenderDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.fromName !== undefined) {
      this.fromName = d.fromName;
    }
    if (d.fromEmailAddress !== undefined) {
      this.fromEmailAddress = d.fromEmailAddress;
    }
    if (d.replyTo !== undefined) {
      this.replyTo = d.replyTo;
    }
    if (d.nickName !== undefined) {
      this.nickName = d.nickName;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
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
      fromName: this.#fromName,
      fromEmailAddress: this.#fromEmailAddress,
      replyTo: this.#replyTo,
      nickName: this.#nickName,
      workspaceId: this.#workspaceId,
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
      fromName: "fromName",
      fromEmailAddress: "fromEmailAddress",
      replyTo: "replyTo",
      nickName: "nickName",
      workspaceId: "workspaceId",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of EmailSenderDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: EmailSenderDtoType) {
    return new EmailSenderDto(possibleDtoObject);
  }
  /**
   * Creates an instance of EmailSenderDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<EmailSenderDtoType>) {
    return new EmailSenderDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<EmailSenderDtoType>,
  ): InstanceType<typeof EmailSenderDto> {
    return new EmailSenderDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof EmailSenderDto> {
    return new EmailSenderDto(this.toJSON());
  }
}
export abstract class EmailSenderDtoFactory {
  abstract create(data: unknown): EmailSenderDto;
}
/**
 * The base type definition for emailSenderDto
 **/
export type EmailSenderDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   *
   * @type {string}
   **/
  fromName: string;
  /**
   *
   * @type {string}
   **/
  fromEmailAddress: string;
  /**
   *
   * @type {string}
   **/
  replyTo: string;
  /**
   *
   * @type {string}
   **/
  nickName: string;
  /**
   * The unique-id of the workspace which content belongs to.
   * @type {string}
   **/
  workspaceId?: string;
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
export namespace EmailSenderDtoType {}
