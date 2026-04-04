// @ts-nocheck 
 // This no check has been added via fireback. 
import { type PartialDeep } from "../../sdk/common/fetchx";
/**
 * The base class definition for permissionInfoDto
 **/
export class PermissionInfoDto {
  /**
   *
   * @type {string}
   **/
  #name: string = "";
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
  set name(value: string) {
    this.#name = String(value);
  }
  setName(value: string) {
    this.name = value;
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
   *
   * @type {string}
   **/
  #completeKey: string = "";
  /**
   *
   * @returns {string}
   **/
  get completeKey() {
    return this.#completeKey;
  }
  /**
   *
   * @type {string}
   **/
  set completeKey(value: string) {
    this.#completeKey = String(value);
  }
  setCompleteKey(value: string) {
    this.completeKey = value;
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
    const d = data as Partial<PermissionInfoDto>;
    if (d.name !== undefined) {
      this.name = d.name;
    }
    if (d.description !== undefined) {
      this.description = d.description;
    }
    if (d.completeKey !== undefined) {
      this.completeKey = d.completeKey;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      name: this.#name,
      description: this.#description,
      completeKey: this.#completeKey,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      name: "name",
      description: "description",
      completeKey: "completeKey",
    };
  }
  /**
   * Creates an instance of PermissionInfoDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: PermissionInfoDtoType) {
    return new PermissionInfoDto(possibleDtoObject);
  }
  /**
   * Creates an instance of PermissionInfoDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<PermissionInfoDtoType>) {
    return new PermissionInfoDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<PermissionInfoDtoType>,
  ): InstanceType<typeof PermissionInfoDto> {
    return new PermissionInfoDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof PermissionInfoDto> {
    return new PermissionInfoDto(this.toJSON());
  }
}
export abstract class PermissionInfoDtoFactory {
  abstract create(data: unknown): PermissionInfoDto;
}
/**
 * The base type definition for permissionInfoDto
 **/
export type PermissionInfoDtoType = {
  /**
   *
   * @type {string}
   **/
  name: string;
  /**
   *
   * @type {string}
   **/
  description: string;
  /**
   *
   * @type {string}
   **/
  completeKey: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PermissionInfoDtoType {}
