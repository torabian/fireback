// @ts-nocheck 
 // This no check has been added via fireback. 
/**
 * The base class definition for assignRoleDto
 **/
export class AssignRoleDto {
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
   * @type {string}
   **/
  #userId: string = "";
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
  set userId(value: string) {
    this.#userId = String(value);
  }
  setUserId(value: string) {
    this.userId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #visibility: string = "";
  /**
   *
   * @returns {string}
   **/
  get visibility() {
    return this.#visibility;
  }
  /**
   *
   * @type {string}
   **/
  set visibility(value: string) {
    this.#visibility = String(value);
  }
  setVisibility(value: string) {
    this.visibility = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #updated: number = 0;
  /**
   *
   * @returns {number}
   **/
  get updated() {
    return this.#updated;
  }
  /**
   *
   * @type {number}
   **/
  set updated(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#updated = parsedValue;
    }
  }
  setUpdated(value: number) {
    this.updated = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #created: number = 0;
  /**
   *
   * @returns {number}
   **/
  get created() {
    return this.#created;
  }
  /**
   *
   * @type {number}
   **/
  set created(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#created = parsedValue;
    }
  }
  setCreated(value: number) {
    this.created = value;
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
    const d = data as Partial<AssignRoleDto>;
    if (d.roleId !== undefined) {
      this.roleId = d.roleId;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.visibility !== undefined) {
      this.visibility = d.visibility;
    }
    if (d.updated !== undefined) {
      this.updated = d.updated;
    }
    if (d.created !== undefined) {
      this.created = d.created;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      roleId: this.#roleId,
      userId: this.#userId,
      visibility: this.#visibility,
      updated: this.#updated,
      created: this.#created,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      roleId: "roleId",
      userId: "userId",
      visibility: "visibility",
      updated: "updated",
      created: "created",
    };
  }
  /**
   * Creates an instance of AssignRoleDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: AssignRoleDtoType) {
    return new AssignRoleDto(possibleDtoObject);
  }
  /**
   * Creates an instance of AssignRoleDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<AssignRoleDtoType>) {
    return new AssignRoleDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<AssignRoleDtoType>,
  ): InstanceType<typeof AssignRoleDto> {
    return new AssignRoleDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof AssignRoleDto> {
    return new AssignRoleDto(this.toJSON());
  }
}
export abstract class AssignRoleDtoFactory {
  abstract create(data: unknown): AssignRoleDto;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for assignRoleDto
 **/
export type AssignRoleDtoType = {
  /**
   *
   * @type {string}
   **/
  roleId: string;
  /**
   *
   * @type {string}
   **/
  userId: string;
  /**
   *
   * @type {string}
   **/
  visibility: string;
  /**
   *
   * @type {number}
   **/
  updated: number;
  /**
   *
   * @type {number}
   **/
  created: number;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace AssignRoleDtoType {}
