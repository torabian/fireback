// @ts-nocheck 
 // This no check has been added via fireback. 
import { type PartialDeep } from "../../sdk/common/fetchx";
/**
 * The base class definition for exchangeKeyInformationDto
 **/
export class ExchangeKeyInformationDto {
  /**
   *
   * @type {string}
   **/
  #key: string = "";
  /**
   *
   * @returns {string}
   **/
  get key() {
    return this.#key;
  }
  /**
   *
   * @type {string}
   **/
  set key(value: string) {
    this.#key = String(value);
  }
  setKey(value: string) {
    this.key = value;
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
    const d = data as Partial<ExchangeKeyInformationDto>;
    if (d.key !== undefined) {
      this.key = d.key;
    }
    if (d.visibility !== undefined) {
      this.visibility = d.visibility;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      key: this.#key,
      visibility: this.#visibility,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      key: "key",
      visibility: "visibility",
    };
  }
  /**
   * Creates an instance of ExchangeKeyInformationDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ExchangeKeyInformationDtoType) {
    return new ExchangeKeyInformationDto(possibleDtoObject);
  }
  /**
   * Creates an instance of ExchangeKeyInformationDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ExchangeKeyInformationDtoType>) {
    return new ExchangeKeyInformationDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ExchangeKeyInformationDtoType>,
  ): InstanceType<typeof ExchangeKeyInformationDto> {
    return new ExchangeKeyInformationDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ExchangeKeyInformationDto> {
    return new ExchangeKeyInformationDto(this.toJSON());
  }
}
export abstract class ExchangeKeyInformationDtoFactory {
  abstract create(data: unknown): ExchangeKeyInformationDto;
}
/**
 * The base type definition for exchangeKeyInformationDto
 **/
export type ExchangeKeyInformationDtoType = {
  /**
   *
   * @type {string}
   **/
  key: string;
  /**
   *
   * @type {string}
   **/
  visibility: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ExchangeKeyInformationDtoType {}
