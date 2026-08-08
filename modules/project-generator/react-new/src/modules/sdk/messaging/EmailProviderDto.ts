import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for emailProviderDto
 **/
export class EmailProviderDto {
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
   * Type of the service, or communication which actually is being used under the hood for providing the service, such as third party or printing right away for terminal or logs.
   * @type {"sendgrid" | "mailgun" | "postmark" | "resend" | "smtp" | "terminal"}
   **/
  #type!: "sendgrid" | "mailgun" | "postmark" | "resend" | "smtp" | "terminal";
  /**
   * Type of the service, or communication which actually is being used under the hood for providing the service, such as third party or printing right away for terminal or logs.
   * @returns {"sendgrid" | "mailgun" | "postmark" | "resend" | "smtp" | "terminal"}
   **/
  get type() {
    return this.#type;
  }
  /**
   * Type of the service, or communication which actually is being used under the hood for providing the service, such as third party or printing right away for terminal or logs.
   * @type {"sendgrid" | "mailgun" | "postmark" | "resend" | "smtp" | "terminal"}
   **/
  set type(
    value: "sendgrid" | "mailgun" | "postmark" | "resend" | "smtp" | "terminal",
  ) {
    this.#type = value;
  }
  setType(
    value: "sendgrid" | "mailgun" | "postmark" | "resend" | "smtp" | "terminal",
  ) {
    this.type = value;
    return this;
  }
  /**
   * Give the email provider configuration a name, which makes it easier later to query.
   * @type {string}
   **/
  #title: string = "";
  /**
   * Give the email provider configuration a name, which makes it easier later to query.
   * @returns {string}
   **/
  get title() {
    return this.#title;
  }
  /**
   * Give the email provider configuration a name, which makes it easier later to query.
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
   * JSON field which contains api keys, or other kind of configuration based on the type of the email provider.
   * @type {JSON}
   **/
  #config!: JSON;
  /**
   * JSON field which contains api keys, or other kind of configuration based on the type of the email provider.
   * @returns {JSON}
   **/
  get config() {
    return this.#config;
  }
  /**
   * JSON field which contains api keys, or other kind of configuration based on the type of the email provider.
   * @type {JSON}
   **/
  set config(value: JSON) {
    this.#config = value;
  }
  setConfig(value: JSON) {
    this.config = value;
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
    const d = data as Partial<EmailProviderDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.type !== undefined) {
      this.type = d.type;
    }
    if (d.title !== undefined) {
      this.title = d.title;
    }
    if (d.config !== undefined) {
      this.config = d.config;
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
      type: this.#type,
      title: this.#title,
      config: this.#config,
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
      type: "type",
      title: "title",
      config: "config",
      workspaceId: "workspaceId",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of EmailProviderDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: EmailProviderDtoType) {
    return new EmailProviderDto(possibleDtoObject);
  }
  /**
   * Creates an instance of EmailProviderDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<EmailProviderDtoType>) {
    return new EmailProviderDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<EmailProviderDtoType>,
  ): InstanceType<typeof EmailProviderDto> {
    return new EmailProviderDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof EmailProviderDto> {
    return new EmailProviderDto(this.toJSON());
  }
}
export abstract class EmailProviderDtoFactory {
  abstract create(data: unknown): EmailProviderDto;
}
/**
 * The base type definition for emailProviderDto
 **/
export type EmailProviderDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * Type of the service, or communication which actually is being used under the hood for providing the service, such as third party or printing right away for terminal or logs.
   * @type {"sendgrid" | "mailgun" | "postmark" | "resend" | "smtp" | "terminal"}
   **/
  type: "sendgrid" | "mailgun" | "postmark" | "resend" | "smtp" | "terminal";
  /**
   * Give the email provider configuration a name, which makes it easier later to query.
   * @type {string}
   **/
  title: string;
  /**
   * JSON field which contains api keys, or other kind of configuration based on the type of the email provider.
   * @type {JSON}
   **/
  config: JSON;
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
export namespace EmailProviderDtoType {}
