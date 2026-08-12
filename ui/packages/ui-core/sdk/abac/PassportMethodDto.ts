import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for passportMethodDto
 **/
export class PassportMethodDto {
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
   * @type {"email" | "phone" | "google" | "facebook"}
   **/
  #type!: "email" | "phone" | "google" | "facebook";
  /**
   *
   * @returns {"email" | "phone" | "google" | "facebook"}
   **/
  get type() {
    return this.#type;
  }
  /**
   *
   * @type {"email" | "phone" | "google" | "facebook"}
   **/
  set type(value: "email" | "phone" | "google" | "facebook") {
    this.#type = value;
  }
  setType(value: "email" | "phone" | "google" | "facebook") {
    this.type = value;
    return this;
  }
  /**
   * The region which would be using this method of passports for authentication. In Fireback open-source, only 'global' is available.
   * @type {"global"}
   **/
  #region: "global" = "global";
  /**
   * The region which would be using this method of passports for authentication. In Fireback open-source, only 'global' is available.
   * @returns {"global"}
   **/
  get region() {
    return this.#region;
  }
  /**
   * The region which would be using this method of passports for authentication. In Fireback open-source, only 'global' is available.
   * @type {"global"}
   **/
  set region(value: "global") {
    this.#region = value;
  }
  setRegion(value: "global") {
    this.region = value;
    return this;
  }
  /**
   * Client key for those methods such as 'google' which require oauth client key
   * @type {string}
   **/
  #clientKey: string = "";
  /**
   * Client key for those methods such as 'google' which require oauth client key
   * @returns {string}
   **/
  get clientKey() {
    return this.#clientKey;
  }
  /**
   * Client key for those methods such as 'google' which require oauth client key
   * @type {string}
   **/
  set clientKey(value: string) {
    this.#clientKey = String(value);
  }
  setClientKey(value: string) {
    this.clientKey = value;
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
    const d = data as Partial<PassportMethodDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.type !== undefined) {
      this.type = d.type;
    }
    if (d.region !== undefined) {
      this.region = d.region;
    }
    if (d.clientKey !== undefined) {
      this.clientKey = d.clientKey;
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
      region: this.#region,
      clientKey: this.#clientKey,
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
      region: "region",
      clientKey: "clientKey",
      workspaceId: "workspaceId",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of PassportMethodDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: PassportMethodDtoType) {
    return new PassportMethodDto(possibleDtoObject);
  }
  /**
   * Creates an instance of PassportMethodDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<PassportMethodDtoType>) {
    return new PassportMethodDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<PassportMethodDtoType>,
  ): InstanceType<typeof PassportMethodDto> {
    return new PassportMethodDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof PassportMethodDto> {
    return new PassportMethodDto(this.toJSON());
  }
}
export abstract class PassportMethodDtoFactory {
  abstract create(data: unknown): PassportMethodDto;
}
/**
 * The base type definition for passportMethodDto
 **/
export type PassportMethodDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   *
   * @type {"email" | "phone" | "google" | "facebook"}
   **/
  type: "email" | "phone" | "google" | "facebook";
  /**
   * The region which would be using this method of passports for authentication. In Fireback open-source, only 'global' is available.
   * @type {"global"}
   **/
  region: "global";
  /**
   * Client key for those methods such as 'google' which require oauth client key
   * @type {string}
   **/
  clientKey: string;
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
export namespace PassportMethodDtoType {}
