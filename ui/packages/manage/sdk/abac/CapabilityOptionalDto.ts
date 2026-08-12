import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for capabilityOptionalDto
 **/
export class CapabilityOptionalDto {
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
   *
   * @type {string}
   **/
  #description?: string | null = undefined;
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
  set description(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#description = correctType ? value : String(value);
  }
  setDescription(value: string | null | undefined) {
    this.description = value;
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
    const d = data as Partial<CapabilityOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.name !== undefined) {
      this.name = d.name;
    }
    if (d.description !== undefined) {
      this.description = d.description;
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
      description: this.#description,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      name: "name",
      description: "description",
    };
  }
  /**
   * Creates an instance of CapabilityOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CapabilityOptionalDtoType) {
    return new CapabilityOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of CapabilityOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<CapabilityOptionalDtoType>) {
    return new CapabilityOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CapabilityOptionalDtoType>,
  ): InstanceType<typeof CapabilityOptionalDto> {
    return new CapabilityOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CapabilityOptionalDto> {
    return new CapabilityOptionalDto(this.toJSON());
  }
}
export abstract class CapabilityOptionalDtoFactory {
  abstract create(data: unknown): CapabilityOptionalDto;
}
/**
 * The base type definition for capabilityOptionalDto
 **/
export type CapabilityOptionalDtoType = {
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
   *
   * @type {string}
   **/
  description?: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CapabilityOptionalDtoType {}
