import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for okayResponseDto
 **/
export class OkayResponseDto {
  /**
   *
   * @type {number}
   **/
  #status: number = 0;
  /**
   *
   * @returns {number}
   **/
  get status() {
    return this.#status;
  }
  /**
   *
   * @type {number}
   **/
  set status(value: number) {
    const correctType = typeof value === "number";
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#status = parsedValue;
    }
  }
  setStatus(value: number) {
    this.status = value;
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
    const d = data as Partial<OkayResponseDto>;
    if (d.status !== undefined) {
      this.status = d.status;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      status: this.#status,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      status: "status",
    };
  }
  /**
   * Creates an instance of OkayResponseDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: OkayResponseDtoType) {
    return new OkayResponseDto(possibleDtoObject);
  }
  /**
   * Creates an instance of OkayResponseDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<OkayResponseDtoType>) {
    return new OkayResponseDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<OkayResponseDtoType>,
  ): InstanceType<typeof OkayResponseDto> {
    return new OkayResponseDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof OkayResponseDto> {
    return new OkayResponseDto(this.toJSON());
  }
}
export abstract class OkayResponseDtoFactory {
  abstract create(data: unknown): OkayResponseDto;
}
/**
 * The base type definition for okayResponseDto
 **/
export type OkayResponseDtoType = {
  /**
   *
   * @type {number}
   **/
  status: number;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace OkayResponseDtoType {}
