import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for workspaceTypeDto
 **/
export class WorkspaceTypeDto {
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
  #title: string = "";
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
  set title(value: string) {
    this.#title = String(value);
  }
  setTitle(value: string) {
    this.title = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #description: string = "";
  /**
   *
   * @returns {string}
   **/
  get description() {
    return this.#description;
  }
  /**
   *
   * @type {string}
   **/
  set description(value: string) {
    this.#description = String(value);
  }
  setDescription(value: string) {
    this.description = value;
    return this;
  }
  /**
   * Unique, URL-safe identifier for this workspace type. Must start with "/", contain only lowercase a-z and dashes after that, and be at most 50 characters long. Format is enforced in ValidateTheWorkspaceTypeEntity (WorkspaceTypeActions.go); gorm:"unique" here is the DB-level backstop for the uniqueness half of that.
   * @type {string}
   **/
  #slug: string = "";
  /**
   * Unique, URL-safe identifier for this workspace type. Must start with "/", contain only lowercase a-z and dashes after that, and be at most 50 characters long. Format is enforced in ValidateTheWorkspaceTypeEntity (WorkspaceTypeActions.go); gorm:"unique" here is the DB-level backstop for the uniqueness half of that.
   * @returns {string}
   **/
  get slug() {
    return this.#slug;
  }
  /**
   * Unique, URL-safe identifier for this workspace type. Must start with "/", contain only lowercase a-z and dashes after that, and be at most 50 characters long. Format is enforced in ValidateTheWorkspaceTypeEntity (WorkspaceTypeActions.go); gorm:"unique" here is the DB-level backstop for the uniqueness half of that.
   * @type {string}
   **/
  set slug(value: string) {
    this.#slug = String(value);
  }
  setSlug(value: string) {
    this.slug = value;
    return this;
  }
  /**
   * The role which will be used to define the functionality of this workspace. Role needs to be created before hand, and only roles which belong to root workspace are possible to be selected.
   * @type {string}
   **/
  #roleId: string = "";
  /**
   * The role which will be used to define the functionality of this workspace. Role needs to be created before hand, and only roles which belong to root workspace are possible to be selected.
   * @returns {string}
   **/
  get roleId() {
    return this.#roleId;
  }
  /**
   * The role which will be used to define the functionality of this workspace. Role needs to be created before hand, and only roles which belong to root workspace are possible to be selected.
   * @type {string}
   **/
  set roleId(value: string) {
    this.#roleId = String(value);
  }
  setRoleId(value: string) {
    this.roleId = value;
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
    const d = data as Partial<WorkspaceTypeDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.title !== undefined) {
      this.title = d.title;
    }
    if (d.description !== undefined) {
      this.description = d.description;
    }
    if (d.slug !== undefined) {
      this.slug = d.slug;
    }
    if (d.roleId !== undefined) {
      this.roleId = d.roleId;
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
      title: this.#title,
      description: this.#description,
      slug: this.#slug,
      roleId: this.#roleId,
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
      title: "title",
      description: "description",
      slug: "slug",
      roleId: "roleId",
      workspaceId: "workspaceId",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of WorkspaceTypeDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WorkspaceTypeDtoType) {
    return new WorkspaceTypeDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WorkspaceTypeDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WorkspaceTypeDtoType>) {
    return new WorkspaceTypeDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WorkspaceTypeDtoType>,
  ): InstanceType<typeof WorkspaceTypeDto> {
    return new WorkspaceTypeDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WorkspaceTypeDto> {
    return new WorkspaceTypeDto(this.toJSON());
  }
}
export abstract class WorkspaceTypeDtoFactory {
  abstract create(data: unknown): WorkspaceTypeDto;
}
/**
 * The base type definition for workspaceTypeDto
 **/
export type WorkspaceTypeDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   *
   * @type {string}
   **/
  title: string;
  /**
   *
   * @type {string}
   **/
  description: string;
  /**
   * Unique, URL-safe identifier for this workspace type. Must start with "/", contain only lowercase a-z and dashes after that, and be at most 50 characters long. Format is enforced in ValidateTheWorkspaceTypeEntity (WorkspaceTypeActions.go); gorm:"unique" here is the DB-level backstop for the uniqueness half of that.
   * @type {string}
   **/
  slug: string;
  /**
   * The role which will be used to define the functionality of this workspace. Role needs to be created before hand, and only roles which belong to root workspace are possible to be selected.
   * @type {string}
   **/
  roleId: string;
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
export namespace WorkspaceTypeDtoType {}
