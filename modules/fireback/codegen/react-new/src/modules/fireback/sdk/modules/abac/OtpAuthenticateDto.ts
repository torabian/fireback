// @ts-nocheck 
 // This no check has been added via fireback. 
import { type PartialDeep } from "../../sdk/common/fetchx";
/**
 * The base class definition for otpAuthenticateDto
 **/
export class OtpAuthenticateDto {
  /**
   *
   * @type {string}
   **/
  #value: string = "";
  /**
   *
   * @returns {string}
   **/
  get value() {
    return this.#value;
  }
  /**
   *
   * @type {string}
   **/
  set value(value: string) {
    this.#value = String(value);
  }
  setValue(value: string) {
    this.value = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #otp: string = "";
  /**
   *
   * @returns {string}
   **/
  get otp() {
    return this.#otp;
  }
  /**
   *
   * @type {string}
   **/
  set otp(value: string) {
    this.#otp = String(value);
  }
  setOtp(value: string) {
    this.otp = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #type: string = "";
  /**
   *
   * @returns {string}
   **/
  get type() {
    return this.#type;
  }
  /**
   *
   * @type {string}
   **/
  set type(value: string) {
    this.#type = String(value);
  }
  setType(value: string) {
    this.type = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #password: string = "";
  /**
   *
   * @returns {string}
   **/
  get password() {
    return this.#password;
  }
  /**
   *
   * @type {string}
   **/
  set password(value: string) {
    this.#password = String(value);
  }
  setPassword(value: string) {
    this.password = value;
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
    const d = data as Partial<OtpAuthenticateDto>;
    if (d.value !== undefined) {
      this.value = d.value;
    }
    if (d.otp !== undefined) {
      this.otp = d.otp;
    }
    if (d.type !== undefined) {
      this.type = d.type;
    }
    if (d.password !== undefined) {
      this.password = d.password;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      value: this.#value,
      otp: this.#otp,
      type: this.#type,
      password: this.#password,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      value: "value",
      otp: "otp",
      type: "type",
      password: "password",
    };
  }
  /**
   * Creates an instance of OtpAuthenticateDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: OtpAuthenticateDtoType) {
    return new OtpAuthenticateDto(possibleDtoObject);
  }
  /**
   * Creates an instance of OtpAuthenticateDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<OtpAuthenticateDtoType>) {
    return new OtpAuthenticateDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<OtpAuthenticateDtoType>,
  ): InstanceType<typeof OtpAuthenticateDto> {
    return new OtpAuthenticateDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof OtpAuthenticateDto> {
    return new OtpAuthenticateDto(this.toJSON());
  }
}
export abstract class OtpAuthenticateDtoFactory {
  abstract create(data: unknown): OtpAuthenticateDto;
}
/**
 * The base type definition for otpAuthenticateDto
 **/
export type OtpAuthenticateDtoType = {
  /**
   *
   * @type {string}
   **/
  value: string;
  /**
   *
   * @type {string}
   **/
  otp: string;
  /**
   *
   * @type {string}
   **/
  type: string;
  /**
   *
   * @type {string}
   **/
  password: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace OtpAuthenticateDtoType {}
