import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for userWorkspaceDto
 **/
export class UserWorkspaceDto {
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
   * @type {string[]}
   **/
  #userPermissions: string[] = [];
  /**
   *
   * @returns {string[]}
   **/
  get userPermissions() {
    return this.#userPermissions;
  }
  /**
   *
   * @type {string[]}
   **/
  set userPermissions(value: string[]) {
    this.#userPermissions = value;
  }
  setUserPermissions(value: string[]) {
    this.userPermissions = value;
    return this;
  }
  /**
   *
   * @type {unknown[]}
   **/
  #rolePermission: unknown[] = [];
  /**
   *
   * @returns {unknown[]}
   **/
  get rolePermission() {
    return this.#rolePermission;
  }
  /**
   *
   * @type {unknown[]}
   **/
  set rolePermission(value: unknown[]) {
    this.#rolePermission = value;
  }
  setRolePermission(value: unknown[]) {
    this.rolePermission = value;
    return this;
  }
  /**
   *
   * @type {string[]}
   **/
  #workspacePermissions: string[] = [];
  /**
   *
   * @returns {string[]}
   **/
  get workspacePermissions() {
    return this.#workspacePermissions;
  }
  /**
   *
   * @type {string[]}
   **/
  set workspacePermissions(value: string[]) {
    this.#workspacePermissions = value;
  }
  setWorkspacePermissions(value: string[]) {
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
    const d = data as Partial<UserWorkspaceDto>;
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
      userPermissions$: "userPermissions",
      get userPermissions() {
        return "userPermissions[:i]";
      },
      rolePermission$: "rolePermission",
      get rolePermission() {
        return "rolePermission[:i]";
      },
      workspacePermissions$: "workspacePermissions",
      get workspacePermissions() {
        return "workspacePermissions[:i]";
      },
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of UserWorkspaceDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: UserWorkspaceDtoType) {
    return new UserWorkspaceDto(possibleDtoObject);
  }
  /**
   * Creates an instance of UserWorkspaceDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<UserWorkspaceDtoType>) {
    return new UserWorkspaceDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<UserWorkspaceDtoType>,
  ): InstanceType<typeof UserWorkspaceDto> {
    return new UserWorkspaceDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof UserWorkspaceDto> {
    return new UserWorkspaceDto(this.toJSON());
  }
}
export abstract class UserWorkspaceDtoFactory {
  abstract create(data: unknown): UserWorkspaceDto;
}
/**
 * The base type definition for userWorkspaceDto
 **/
export type UserWorkspaceDtoType = {
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
   * @type {string[]}
   **/
  userPermissions: string[];
  /**
   *
   * @type {unknown[]}
   **/
  rolePermission: unknown[];
  /**
   *
   * @type {string[]}
   **/
  workspacePermissions: string[];
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
export namespace UserWorkspaceDtoType {}
