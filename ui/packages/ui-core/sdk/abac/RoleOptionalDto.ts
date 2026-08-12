import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for roleOptionalDto
 **/
export class RoleOptionalDto {
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
  #name?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get name() {
    return this.#name;
  }
  /**
   *
   * @type {string}
   **/
  set name(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#name = correctType ? value : String(value);
  }
  setName(value: string | null | undefined) {
    this.name = value;
    return this;
  }
  /**
   * The list of capability completeKeys granted to this role, stored directly as JSON (replaces the old many-to-many role_capabilities join table - Emi has no relation mechanism compatible with fireback's string-uniqueId FK convention, see other entities' xId fields).
   * @type {JSON}
   **/
  #capabilitiesListId!: JSON;
  /**
   * The list of capability completeKeys granted to this role, stored directly as JSON (replaces the old many-to-many role_capabilities join table - Emi has no relation mechanism compatible with fireback's string-uniqueId FK convention, see other entities' xId fields).
   * @returns {JSON}
   **/
  get capabilitiesListId() {
    return this.#capabilitiesListId;
  }
  /**
   * The list of capability completeKeys granted to this role, stored directly as JSON (replaces the old many-to-many role_capabilities join table - Emi has no relation mechanism compatible with fireback's string-uniqueId FK convention, see other entities' xId fields).
   * @type {JSON}
   **/
  set capabilitiesListId(value: JSON) {
    this.#capabilitiesListId = value;
  }
  setCapabilitiesListId(value: JSON) {
    this.capabilitiesListId = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #isDeletable?: boolean | null = undefined;
  /**
   *
   * @returns {boolean}
   **/
  get isDeletable() {
    return this.#isDeletable;
  }
  /**
   *
   * @type {boolean}
   **/
  set isDeletable(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#isDeletable = correctType ? value : Boolean(value);
  }
  setIsDeletable(value: boolean | null | undefined) {
    this.isDeletable = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #isUpdatable?: boolean | null = undefined;
  /**
   *
   * @returns {boolean}
   **/
  get isUpdatable() {
    return this.#isUpdatable;
  }
  /**
   *
   * @type {boolean}
   **/
  set isUpdatable(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#isUpdatable = correctType ? value : Boolean(value);
  }
  setIsUpdatable(value: boolean | null | undefined) {
    this.isUpdatable = value;
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
  /**
   *
   * @type {string}
   **/
  #userId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   *
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
    const d = data as Partial<RoleOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.name !== undefined) {
      this.name = d.name;
    }
    if (d.capabilitiesListId !== undefined) {
      this.capabilitiesListId = d.capabilitiesListId;
    }
    if (d.isDeletable !== undefined) {
      this.isDeletable = d.isDeletable;
    }
    if (d.isUpdatable !== undefined) {
      this.isUpdatable = d.isUpdatable;
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
      name: this.#name,
      capabilitiesListId: this.#capabilitiesListId,
      isDeletable: this.#isDeletable,
      isUpdatable: this.#isUpdatable,
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
      name: "name",
      capabilitiesListId: "capabilitiesListId",
      isDeletable: "isDeletable",
      isUpdatable: "isUpdatable",
      workspaceId: "workspaceId",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of RoleOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: RoleOptionalDtoType) {
    return new RoleOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of RoleOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<RoleOptionalDtoType>) {
    return new RoleOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<RoleOptionalDtoType>,
  ): InstanceType<typeof RoleOptionalDto> {
    return new RoleOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof RoleOptionalDto> {
    return new RoleOptionalDto(this.toJSON());
  }
}
export abstract class RoleOptionalDtoFactory {
  abstract create(data: unknown): RoleOptionalDto;
}
/**
 * The base type definition for roleOptionalDto
 **/
export type RoleOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   *
   * @type {string}
   **/
  name?: string;
  /**
   * The list of capability completeKeys granted to this role, stored directly as JSON (replaces the old many-to-many role_capabilities join table - Emi has no relation mechanism compatible with fireback's string-uniqueId FK convention, see other entities' xId fields).
   * @type {JSON}
   **/
  capabilitiesListId: JSON;
  /**
   *
   * @type {boolean}
   **/
  isDeletable?: boolean;
  /**
   *
   * @type {boolean}
   **/
  isUpdatable?: boolean;
  /**
   *
   * @type {string}
   **/
  workspaceId?: string;
  /**
   *
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
export namespace RoleOptionalDtoType {}
