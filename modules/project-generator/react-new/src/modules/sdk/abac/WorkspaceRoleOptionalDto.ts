import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for workspaceRoleOptionalDto
 **/
export class WorkspaceRoleOptionalDto {
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
   * The unique-id of the userWorkspace this role assignment belongs to.
   * @type {string}
   **/
  #userWorkspaceId?: string | null = undefined;
  /**
   * The unique-id of the userWorkspace this role assignment belongs to.
   * @returns {string}
   **/
  get userWorkspaceId() {
    return this.#userWorkspaceId;
  }
  /**
   * The unique-id of the userWorkspace this role assignment belongs to.
   * @type {string}
   **/
  set userWorkspaceId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#userWorkspaceId = correctType ? value : String(value);
  }
  setUserWorkspaceId(value: string | null | undefined) {
    this.userWorkspaceId = value;
    return this;
  }
  /**
   * The unique-id of the assigned role.
   * @type {string}
   **/
  #roleId?: string | null = undefined;
  /**
   * The unique-id of the assigned role.
   * @returns {string}
   **/
  get roleId() {
    return this.#roleId;
  }
  /**
   * The unique-id of the assigned role.
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
    const d = data as Partial<WorkspaceRoleOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.userWorkspaceId !== undefined) {
      this.userWorkspaceId = d.userWorkspaceId;
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
      userWorkspaceId: this.#userWorkspaceId,
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
      userWorkspaceId: "userWorkspaceId",
      roleId: "roleId",
      workspaceId: "workspaceId",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of WorkspaceRoleOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WorkspaceRoleOptionalDtoType) {
    return new WorkspaceRoleOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WorkspaceRoleOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WorkspaceRoleOptionalDtoType>) {
    return new WorkspaceRoleOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WorkspaceRoleOptionalDtoType>,
  ): InstanceType<typeof WorkspaceRoleOptionalDto> {
    return new WorkspaceRoleOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WorkspaceRoleOptionalDto> {
    return new WorkspaceRoleOptionalDto(this.toJSON());
  }
}
export abstract class WorkspaceRoleOptionalDtoFactory {
  abstract create(data: unknown): WorkspaceRoleOptionalDto;
}
/**
 * The base type definition for workspaceRoleOptionalDto
 **/
export type WorkspaceRoleOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * The unique-id of the userWorkspace this role assignment belongs to.
   * @type {string}
   **/
  userWorkspaceId?: string;
  /**
   * The unique-id of the assigned role.
   * @type {string}
   **/
  roleId?: string;
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
export namespace WorkspaceRoleOptionalDtoType {}
