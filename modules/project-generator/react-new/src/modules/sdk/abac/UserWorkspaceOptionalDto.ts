import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for userWorkspaceOptionalDto
 **/
export class UserWorkspaceOptionalDto {
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
   * The unique-id of the user this record belongs to.
   * @type {string}
   **/
  #userId?: string | null = undefined;
  /**
   * The unique-id of the user this record belongs to.
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * The unique-id of the user this record belongs to.
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
   * @type {any}
   **/
  #userPermissions?: any | null = undefined;
  /**
   *
   * @returns {any}
   **/
  get userPermissions() {
    return this.#userPermissions;
  }
  /**
   *
   * @type {any}
   **/
  set userPermissions(value: any | null | undefined) {
    this.#userPermissions = value;
  }
  setUserPermissions(value: any | null | undefined) {
    this.userPermissions = value;
    return this;
  }
  /**
   *
   * @type {any}
   **/
  #rolePermission?: any | null = undefined;
  /**
   *
   * @returns {any}
   **/
  get rolePermission() {
    return this.#rolePermission;
  }
  /**
   *
   * @type {any}
   **/
  set rolePermission(value: any | null | undefined) {
    this.#rolePermission = value;
  }
  setRolePermission(value: any | null | undefined) {
    this.rolePermission = value;
    return this;
  }
  /**
   *
   * @type {any}
   **/
  #workspacePermissions?: any | null = undefined;
  /**
   *
   * @returns {any}
   **/
  get workspacePermissions() {
    return this.#workspacePermissions;
  }
  /**
   *
   * @type {any}
   **/
  set workspacePermissions(value: any | null | undefined) {
    this.#workspacePermissions = value;
  }
  setWorkspacePermissions(value: any | null | undefined) {
    this.workspacePermissions = value;
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
    const d = data as Partial<UserWorkspaceOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
    }
    if (d.userPermissions !== undefined) {
      this.userPermissions = d.userPermissions;
    }
    if (d.rolePermission !== undefined) {
      this.rolePermission = d.rolePermission;
    }
    if (d.workspacePermissions !== undefined) {
      this.workspacePermissions = d.workspacePermissions;
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
      workspaceId: this.#workspaceId,
      userPermissions: this.#userPermissions,
      rolePermission: this.#rolePermission,
      workspacePermissions: this.#workspacePermissions,
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
      workspaceId: "workspaceId",
      userPermissions: "userPermissions",
      rolePermission: "rolePermission",
      workspacePermissions: "workspacePermissions",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of UserWorkspaceOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: UserWorkspaceOptionalDtoType) {
    return new UserWorkspaceOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of UserWorkspaceOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<UserWorkspaceOptionalDtoType>) {
    return new UserWorkspaceOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<UserWorkspaceOptionalDtoType>,
  ): InstanceType<typeof UserWorkspaceOptionalDto> {
    return new UserWorkspaceOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof UserWorkspaceOptionalDto> {
    return new UserWorkspaceOptionalDto(this.toJSON());
  }
}
export abstract class UserWorkspaceOptionalDtoFactory {
  abstract create(data: unknown): UserWorkspaceOptionalDto;
}
/**
 * The base type definition for userWorkspaceOptionalDto
 **/
export type UserWorkspaceOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * The unique-id of the user this record belongs to.
   * @type {string}
   **/
  userId?: string;
  /**
   * The unique-id of the workspace which content belongs to.
   * @type {string}
   **/
  workspaceId?: string;
  /**
   *
   * @type {any}
   **/
  userPermissions?: any;
  /**
   *
   * @type {any}
   **/
  rolePermission?: any;
  /**
   *
   * @type {any}
   **/
  workspacePermissions?: any;
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
export namespace UserWorkspaceOptionalDtoType {}
