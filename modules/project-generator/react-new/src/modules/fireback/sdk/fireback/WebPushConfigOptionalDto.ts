import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for webPushConfigOptionalDto
 **/
export class WebPushConfigOptionalDto {
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
   * The json content of the web push after getting it from browser
   * @type {JSON}
   **/
  #subscription!: JSON;
  /**
   * The json content of the web push after getting it from browser
   * @returns {JSON}
   **/
  get subscription() {
    return this.#subscription;
  }
  /**
   * The json content of the web push after getting it from browser
   * @type {JSON}
   **/
  set subscription(value: JSON) {
    this.#subscription = value;
  }
  setSubscription(value: JSON) {
    this.subscription = value;
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
    const d = data as Partial<WebPushConfigOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.subscription !== undefined) {
      this.subscription = d.subscription;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      subscription: this.#subscription,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      subscription: "subscription",
    };
  }
  /**
   * Creates an instance of WebPushConfigOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WebPushConfigOptionalDtoType) {
    return new WebPushConfigOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WebPushConfigOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WebPushConfigOptionalDtoType>) {
    return new WebPushConfigOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WebPushConfigOptionalDtoType>,
  ): InstanceType<typeof WebPushConfigOptionalDto> {
    return new WebPushConfigOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WebPushConfigOptionalDto> {
    return new WebPushConfigOptionalDto(this.toJSON());
  }
}
export abstract class WebPushConfigOptionalDtoFactory {
  abstract create(data: unknown): WebPushConfigOptionalDto;
}
/**
 * The base type definition for webPushConfigOptionalDto
 **/
export type WebPushConfigOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * The json content of the web push after getting it from browser
   * @type {JSON}
   **/
  subscription: JSON;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace WebPushConfigOptionalDtoType {}
