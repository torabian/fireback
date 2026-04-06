// @ts-nocheck 
 // This no check has been added via fireback. 
import { type PartialDeep } from "../../sdk/common/fetchx";
/**
 * The base class definition for importRequestDto
 **/
export class ImportRequestDto {
  /**
   *
   * @type {string}
   **/
  #file: string = "";
  /**
   *
   * @returns {string}
   **/
  get file() {
    return this.#file;
  }
  /**
   *
   * @type {string}
   **/
  set file(value: string) {
    this.#file = String(value);
  }
  setFile(value: string) {
    this.file = value;
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
    const d = data as Partial<ImportRequestDto>;
    if (d.file !== undefined) {
      this.file = d.file;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      file: this.#file,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      file: "file",
    };
  }
  /**
   * Creates an instance of ImportRequestDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ImportRequestDtoType) {
    return new ImportRequestDto(possibleDtoObject);
  }
  /**
   * Creates an instance of ImportRequestDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ImportRequestDtoType>) {
    return new ImportRequestDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ImportRequestDtoType>,
  ): InstanceType<typeof ImportRequestDto> {
    return new ImportRequestDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ImportRequestDto> {
    return new ImportRequestDto(this.toJSON());
  }
}
export abstract class ImportRequestDtoFactory {
  abstract create(data: unknown): ImportRequestDto;
}
/**
 * The base type definition for importRequestDto
 **/
export type ImportRequestDtoType = {
  /**
   *
   * @type {string}
   **/
  file: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ImportRequestDtoType {}
