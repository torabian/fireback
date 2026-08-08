import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for regionalContentOptionalDto
 **/
export class RegionalContentOptionalDto {
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
  #content?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get content() {
    return this.#content;
  }
  /**
   *
   * @type {string}
   **/
  set content(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#content = correctType ? value : String(value);
  }
  setContent(value: string | null | undefined) {
    this.content = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #region?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get region() {
    return this.#region;
  }
  /**
   *
   * @type {string}
   **/
  set region(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#region = correctType ? value : String(value);
  }
  setRegion(value: string | null | undefined) {
    this.region = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #title?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get title() {
    return this.#title;
  }
  /**
   *
   * @type {string}
   **/
  set title(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#title = correctType ? value : String(value);
  }
  setTitle(value: string | null | undefined) {
    this.title = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #languageId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get languageId() {
    return this.#languageId;
  }
  /**
   *
   * @type {string}
   **/
  set languageId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#languageId = correctType ? value : String(value);
  }
  setLanguageId(value: string | null | undefined) {
    this.languageId = value;
    return this;
  }
  /**
   *
   * @type {any}
   **/
  #keyGroup?: any | null = undefined;
  /**
   *
   * @returns {any}
   **/
  get keyGroup() {
    return this.#keyGroup;
  }
  /**
   *
   * @type {any}
   **/
  set keyGroup(value: any | null | undefined) {
    this.#keyGroup = value;
  }
  setKeyGroup(value: any | null | undefined) {
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
    const d = data as Partial<RegionalContentOptionalDto>;
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
   * Creates an instance of RegionalContentOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: RegionalContentOptionalDtoType) {
    return new RegionalContentOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of RegionalContentOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<RegionalContentOptionalDtoType>) {
    return new RegionalContentOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<RegionalContentOptionalDtoType>,
  ): InstanceType<typeof RegionalContentOptionalDto> {
    return new RegionalContentOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof RegionalContentOptionalDto> {
    return new RegionalContentOptionalDto(this.toJSON());
  }
}
export abstract class RegionalContentOptionalDtoFactory {
  abstract create(data: unknown): RegionalContentOptionalDto;
}
/**
 * The base type definition for regionalContentOptionalDto
 **/
export type RegionalContentOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   *
   * @type {string}
   **/
  content?: string;
  /**
   *
   * @type {string}
   **/
  region?: string;
  /**
   *
   * @type {string}
   **/
  title?: string;
  /**
   *
   * @type {string}
   **/
  languageId?: string;
  /**
   *
   * @type {any}
   **/
  keyGroup?: any;
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
export namespace RegionalContentOptionalDtoType {}
