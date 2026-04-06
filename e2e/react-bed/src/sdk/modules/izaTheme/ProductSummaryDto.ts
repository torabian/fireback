// @ts-nocheck 
 // This no check has been added via fireback. 
import { type PartialDeep } from "../../sdk/common/fetchx";
/**
 * The base class definition for productSummaryDto
 **/
export class ProductSummaryDto {
  /**
   * Product name, used in titles.
   * @type {string}
   **/
  #name: string = "";
  /**
   * Product name, used in titles.
   * @returns {string}
   **/
  get name() {
    return this.#name;
  }
  /**
   * Product name, used in titles.
   * @type {string}
   **/
  set name(value: string) {
    this.#name = String(value);
  }
  setName(value: string) {
    this.name = value;
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
    const d = data as Partial<ProductSummaryDto>;
    if (d.name !== undefined) {
      this.name = d.name;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      name: this.#name,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      name: "name",
    };
  }
  /**
   * Creates an instance of ProductSummaryDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ProductSummaryDtoType) {
    return new ProductSummaryDto(possibleDtoObject);
  }
  /**
   * Creates an instance of ProductSummaryDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ProductSummaryDtoType>) {
    return new ProductSummaryDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ProductSummaryDtoType>,
  ): InstanceType<typeof ProductSummaryDto> {
    return new ProductSummaryDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ProductSummaryDto> {
    return new ProductSummaryDto(this.toJSON());
  }
}
export abstract class ProductSummaryDtoFactory {
  abstract create(data: unknown): ProductSummaryDto;
}
/**
 * The base type definition for productSummaryDto
 **/
export type ProductSummaryDtoType = {
  /**
   * Product name, used in titles.
   * @type {string}
   **/
  name: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ProductSummaryDtoType {}
