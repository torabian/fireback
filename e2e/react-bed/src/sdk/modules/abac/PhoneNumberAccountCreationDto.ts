// @ts-nocheck 
 // This no check has been added via fireback. 
import { type PartialDeep } from "../../sdk/common/fetchx";
/**
 * The base class definition for phoneNumberAccountCreationDto
 **/
export class PhoneNumberAccountCreationDto {
  /**
   *
   * @type {string}
   **/
  #phoneNumber: string = "";
  /**
   *
   * @returns {string}
   **/
  get phoneNumber() {
    return this.#phoneNumber;
  }
  /**
   *
   * @type {string}
   **/
  set phoneNumber(value: string) {
    this.#phoneNumber = String(value);
  }
  setPhoneNumber(value: string) {
    this.phoneNumber = value;
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
    const d = data as Partial<PhoneNumberAccountCreationDto>;
    if (d.phoneNumber !== undefined) {
      this.phoneNumber = d.phoneNumber;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      phoneNumber: this.#phoneNumber,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      phoneNumber: "phoneNumber",
    };
  }
  /**
   * Creates an instance of PhoneNumberAccountCreationDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: PhoneNumberAccountCreationDtoType) {
    return new PhoneNumberAccountCreationDto(possibleDtoObject);
  }
  /**
   * Creates an instance of PhoneNumberAccountCreationDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<PhoneNumberAccountCreationDtoType>,
  ) {
    return new PhoneNumberAccountCreationDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<PhoneNumberAccountCreationDtoType>,
  ): InstanceType<typeof PhoneNumberAccountCreationDto> {
    return new PhoneNumberAccountCreationDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof PhoneNumberAccountCreationDto> {
    return new PhoneNumberAccountCreationDto(this.toJSON());
  }
}
export abstract class PhoneNumberAccountCreationDtoFactory {
  abstract create(data: unknown): PhoneNumberAccountCreationDto;
}
/**
 * The base type definition for phoneNumberAccountCreationDto
 **/
export type PhoneNumberAccountCreationDtoType = {
  /**
   *
   * @type {string}
   **/
  phoneNumber: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace PhoneNumberAccountCreationDtoType {}
