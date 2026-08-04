// @ts-nocheck 
 // This no check has been added via fireback. 
import { type PartialDeep } from "../../sdk/common/fetchx";
/**
 * The base class definition for userRoleWorkspaceDto
 **/
export class UserRoleWorkspaceDto {
  /**
   *
   * @type {string}
   **/
  #roleId: string = "";
  /**
   *
   * @returns {string}
   **/
  get roleId() {
    return this.#roleId;
  }
  /**
   *
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
   * @type {string[]}
   **/
  #capabilities: string[] = [];
  /**
   *
   * @returns {string[]}
   **/
  get capabilities() {
    return this.#capabilities;
  }
  /**
   *
   * @type {string[]}
   **/
  set capabilities(value: string[]) {
    this.#capabilities = value;
  }
  setCapabilities(value: string[]) {
    this.capabilities = value;
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
    const d = data as Partial<UserRoleWorkspaceDto>;
    if (d.roleId !== undefined) {
      this.roleId = d.roleId;
    }
    if (d.capabilities !== undefined) {
      this.capabilities = d.capabilities;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      roleId: this.#roleId,
      capabilities: this.#capabilities,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      roleId: "roleId",
      capabilities$: "capabilities",
      get capabilities() {
        return "capabilities[:i]";
      },
    };
  }
  /**
   * Creates an instance of UserRoleWorkspaceDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: UserRoleWorkspaceDtoType) {
    return new UserRoleWorkspaceDto(possibleDtoObject);
  }
  /**
   * Creates an instance of UserRoleWorkspaceDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<UserRoleWorkspaceDtoType>) {
    return new UserRoleWorkspaceDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<UserRoleWorkspaceDtoType>,
  ): InstanceType<typeof UserRoleWorkspaceDto> {
    return new UserRoleWorkspaceDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof UserRoleWorkspaceDto> {
    return new UserRoleWorkspaceDto(this.toJSON());
  }
}
export abstract class UserRoleWorkspaceDtoFactory {
  abstract create(data: unknown): UserRoleWorkspaceDto;
}
/**
 * The base type definition for userRoleWorkspaceDto
 **/
export type UserRoleWorkspaceDtoType = {
  /**
   *
   * @type {string}
   **/
  roleId: string;
  /**
   *
   * @type {string[]}
   **/
  capabilities: string[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace UserRoleWorkspaceDtoType {}
