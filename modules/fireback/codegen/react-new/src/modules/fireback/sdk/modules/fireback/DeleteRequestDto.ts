// @ts-nocheck 
 // This no check has been added via fireback. 
import { type PartialDeep } from "../../sdk/common/fetchx";
/**
 * The base class definition for deleteRequestDto
 **/
export class DeleteRequestDto {
  /**
   * The query selector which would be used to delete the content.
   * @type {string}
   **/
  #query: string = "";
  /**
   * The query selector which would be used to delete the content.
   * @returns {string}
   **/
  get query() {
    return this.#query;
  }
  /**
   * The query selector which would be used to delete the content.
   * @type {string}
   **/
  set query(value: string) {
    this.#query = String(value);
  }
  setQuery(value: string) {
    this.query = value;
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
    const d = data as Partial<DeleteRequestDto>;
    if (d.query !== undefined) {
      this.query = d.query;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      query: this.#query,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      query: "query",
    };
  }
  /**
   * Creates an instance of DeleteRequestDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: DeleteRequestDtoType) {
    return new DeleteRequestDto(possibleDtoObject);
  }
  /**
   * Creates an instance of DeleteRequestDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<DeleteRequestDtoType>) {
    return new DeleteRequestDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<DeleteRequestDtoType>,
  ): InstanceType<typeof DeleteRequestDto> {
    return new DeleteRequestDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof DeleteRequestDto> {
    return new DeleteRequestDto(this.toJSON());
  }
}
export abstract class DeleteRequestDtoFactory {
  abstract create(data: unknown): DeleteRequestDto;
}
/**
 * The base type definition for deleteRequestDto
 **/
export type DeleteRequestDtoType = {
  /**
   * The query selector which would be used to delete the content.
   * @type {string}
   **/
  query: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace DeleteRequestDtoType {}
