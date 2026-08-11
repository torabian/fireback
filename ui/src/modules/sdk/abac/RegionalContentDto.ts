import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for regionalContentDto
 **/
export class RegionalContentDto {
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
   * The template body sent to the user - supports Go template syntax to insert dynamic values, such as {{.Otp}} for the one-time password.
   * @type {string}
   **/
  #content: string = "";
  /**
   * The template body sent to the user - supports Go template syntax to insert dynamic values, such as {{.Otp}} for the one-time password.
   * @returns {string}
   **/
  get content() {
    return this.#content;
  }
  /**
   * The template body sent to the user - supports Go template syntax to insert dynamic values, such as {{.Otp}} for the one-time password.
   * @type {string}
   **/
  set content(value: string) {
    this.#content = String(value);
  }
  setContent(value: string) {
    this.content = value;
    return this;
  }
  /**
   * Region or locale this content applies to, for example any, us, eu, or asia/*. Use any unless you need to target a specific region.
   * @type {string}
   **/
  #region: string = "";
  /**
   * Region or locale this content applies to, for example any, us, eu, or asia/*. Use any unless you need to target a specific region.
   * @returns {string}
   **/
  get region() {
    return this.#region;
  }
  /**
   * Region or locale this content applies to, for example any, us, eu, or asia/*. Use any unless you need to target a specific region.
   * @type {string}
   **/
  set region(value: string) {
    this.#region = String(value);
  }
  setRegion(value: string) {
    this.region = value;
    return this;
  }
  /**
   * Optional subject line - only used for email-type content.
   * @type {string}
   **/
  #title: string = "";
  /**
   * Optional subject line - only used for email-type content.
   * @returns {string}
   **/
  get title() {
    return this.#title;
  }
  /**
   * Optional subject line - only used for email-type content.
   * @type {string}
   **/
  set title(value: string) {
    this.#title = String(value);
  }
  setTitle(value: string) {
    this.title = value;
    return this;
  }
  /**
   * Language code this content is written in, for example en, fa, or pl. Falls back to English if nothing matches a user's language.
   * @type {string}
   **/
  #languageId: string = "";
  /**
   * Language code this content is written in, for example en, fa, or pl. Falls back to English if nothing matches a user's language.
   * @returns {string}
   **/
  get languageId() {
    return this.#languageId;
  }
  /**
   * Language code this content is written in, for example en, fa, or pl. Falls back to English if nothing matches a user's language.
   * @type {string}
   **/
  set languageId(value: string) {
    this.#languageId = String(value);
  }
  setLanguageId(value: string) {
    this.languageId = value;
    return this;
  }
  /**
   * Which kind of message this content is used for.
   * @type {"SMS_OTP" | "EMAIL_OTP" | "INVITE_TO_WORKSPACE"}
   **/
  #keyGroup!: "SMS_OTP" | "EMAIL_OTP" | "INVITE_TO_WORKSPACE";
  /**
   * Which kind of message this content is used for.
   * @returns {"SMS_OTP" | "EMAIL_OTP" | "INVITE_TO_WORKSPACE"}
   **/
  get keyGroup() {
    return this.#keyGroup;
  }
  /**
   * Which kind of message this content is used for.
   * @type {"SMS_OTP" | "EMAIL_OTP" | "INVITE_TO_WORKSPACE"}
   **/
  set keyGroup(value: "SMS_OTP" | "EMAIL_OTP" | "INVITE_TO_WORKSPACE") {
    this.#keyGroup = value;
  }
  setKeyGroup(value: "SMS_OTP" | "EMAIL_OTP" | "INVITE_TO_WORKSPACE") {
    this.keyGroup = value;
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
    const d = data as Partial<RegionalContentDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.content !== undefined) {
      this.content = d.content;
    }
    if (d.region !== undefined) {
      this.region = d.region;
    }
    if (d.title !== undefined) {
      this.title = d.title;
    }
    if (d.languageId !== undefined) {
      this.languageId = d.languageId;
    }
    if (d.keyGroup !== undefined) {
      this.keyGroup = d.keyGroup;
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
      content: this.#content,
      region: this.#region,
      title: this.#title,
      languageId: this.#languageId,
      keyGroup: this.#keyGroup,
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
      content: "content",
      region: "region",
      title: "title",
      languageId: "languageId",
      keyGroup: "keyGroup",
      workspaceId: "workspaceId",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of RegionalContentDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: RegionalContentDtoType) {
    return new RegionalContentDto(possibleDtoObject);
  }
  /**
   * Creates an instance of RegionalContentDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<RegionalContentDtoType>) {
    return new RegionalContentDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<RegionalContentDtoType>,
  ): InstanceType<typeof RegionalContentDto> {
    return new RegionalContentDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof RegionalContentDto> {
    return new RegionalContentDto(this.toJSON());
  }
}
export abstract class RegionalContentDtoFactory {
  abstract create(data: unknown): RegionalContentDto;
}
/**
 * The base type definition for regionalContentDto
 **/
export type RegionalContentDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * The template body sent to the user - supports Go template syntax to insert dynamic values, such as {{.Otp}} for the one-time password.
   * @type {string}
   **/
  content: string;
  /**
   * Region or locale this content applies to, for example any, us, eu, or asia/*. Use any unless you need to target a specific region.
   * @type {string}
   **/
  region: string;
  /**
   * Optional subject line - only used for email-type content.
   * @type {string}
   **/
  title: string;
  /**
   * Language code this content is written in, for example en, fa, or pl. Falls back to English if nothing matches a user's language.
   * @type {string}
   **/
  languageId: string;
  /**
   * Which kind of message this content is used for.
   * @type {"SMS_OTP" | "EMAIL_OTP" | "INVITE_TO_WORKSPACE"}
   **/
  keyGroup: "SMS_OTP" | "EMAIL_OTP" | "INVITE_TO_WORKSPACE";
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
export namespace RegionalContentDtoType {}
