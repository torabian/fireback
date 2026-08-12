import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for emailConfirmationDto
 **/
export class EmailConfirmationDto {
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
   * The unique-id of the user this confirmation belongs to.
   * @type {string}
   **/
  #userId?: string | null = undefined;
  /**
   * The unique-id of the user this confirmation belongs to.
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * The unique-id of the user this confirmation belongs to.
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
   * @type {string}
   **/
  #email: string = "";
  /**
   *
   * @returns {string}
   **/
  get email() {
    return this.#email;
  }
  /**
   *
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
   *
   * @type {string}
   **/
  #key: string = "";
  /**
   *
   * @returns {string}
   **/
  get key() {
    return this.#key;
  }
  /**
   *
   * @type {string}
   **/
  set key(value: string) {
    this.#key = String(value);
  }
  setKey(value: string) {
    this.key = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #expiresAt: string = "";
  /**
   *
   * @returns {string}
   **/
  get expiresAt() {
    return this.#expiresAt;
  }
  /**
   *
   * @type {string}
   **/
  set expiresAt(value: string) {
    this.#expiresAt = String(value);
  }
  setExpiresAt(value: string) {
    this.expiresAt = value;
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
    const d = data as Partial<EmailConfirmationDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.status !== undefined) {
      this.status = d.status;
    }
    if (d.email !== undefined) {
      this.email = d.email;
    }
    if (d.key !== undefined) {
      this.key = d.key;
    }
    if (d.expiresAt !== undefined) {
      this.expiresAt = d.expiresAt;
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
      status: this.#status,
      email: this.#email,
      key: this.#key,
      expiresAt: this.#expiresAt,
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
      status: "status",
      email: "email",
      key: "key",
      expiresAt: "expiresAt",
      workspaceId: "workspaceId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of EmailConfirmationDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: EmailConfirmationDtoType) {
    return new EmailConfirmationDto(possibleDtoObject);
  }
  /**
   * Creates an instance of EmailConfirmationDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<EmailConfirmationDtoType>) {
    return new EmailConfirmationDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<EmailConfirmationDtoType>,
  ): InstanceType<typeof EmailConfirmationDto> {
    return new EmailConfirmationDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof EmailConfirmationDto> {
    return new EmailConfirmationDto(this.toJSON());
  }
}
export abstract class EmailConfirmationDtoFactory {
  abstract create(data: unknown): EmailConfirmationDto;
}
/**
 * The base type definition for emailConfirmationDto
 **/
export type EmailConfirmationDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * The unique-id of the user this confirmation belongs to.
   * @type {string}
   **/
  userId?: string;
  /**
   *
   * @type {string}
   **/
  status: string;
  /**
   *
   * @type {string}
   **/
  email: string;
  /**
   *
   * @type {string}
   **/
  key: string;
  /**
   *
   * @type {string}
   **/
  expiresAt: string;
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
export namespace EmailConfirmationDtoType {}
