import { MCollection } from "../sdk/common/operators";
import { type PartialDeep } from "../sdk/common/fetchx";
import { withPrefix } from "../sdk/common/withPrefix";
/**
 * The base class definition for capabilityInfoDto
 **/
export class CapabilityInfoDto {
  /**
   *
   * @type {string}
   **/
  #uniqueId: string = "";
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
  set uniqueId(value: string) {
    this.#uniqueId = String(value);
  }
  setUniqueId(value: string) {
    this.uniqueId = value;
    return this;
  }
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
   * @type {CapabilityInfoDto}
   **/
  #children: MCollection<CapabilityInfoDto> = MCollection.of([]);
  /**
   *
   * @returns {CapabilityInfoDto}
   **/
  get children() {
    return this.#children;
  }
  /**
   *
   * @type {CapabilityInfoDto}
   **/
  set children(
    value:
      | MCollection<CapabilityInfoDto>
      | InstanceType<typeof CapabilityInfoDto>[],
  ) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof CapabilityInfoDto) {
        this.#children = MCollection.of(value);
      } else {
        this.#children = MCollection.of(
          value.map((item) => new CapabilityInfoDto(item)),
        );
      }
      return;
    }
    // If the instance is already an MCollection, we assume it's all good.
    if (value instanceof MCollection) {
      this.#children = value;
      return;
    }
    // If the value is not array, and is not a MCollection, we need to be consider,
    // it might be eligible to be casted into MCollection.
    const { ok, value: mcastValue } = MCollection.cast<unknown>(value);
    if (ok) {
      this.#children = mcastValue as any;
      return;
    }
    console.warn(
      "Cannot assing value to children, because it needs MCollection instance or an Array.",
    );
  }
  setChildren(
    value:
      | MCollection<CapabilityInfoDto>
      | InstanceType<typeof CapabilityInfoDto>[],
  ) {
    this.children = value;
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
    const d = data as Partial<CapabilityInfoDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.name !== undefined) {
      this.name = d.name;
    }
    if (d.children !== undefined) {
      this.children = d.children;
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
      children: this.#children,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      name: "name",
      children$: "children",
      get children() {
        return withPrefix("children[:i]", CapabilityInfoDto.Fields);
      },
    };
  }
  /**
   * Creates an instance of CapabilityInfoDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: CapabilityInfoDtoType) {
    return new CapabilityInfoDto(possibleDtoObject);
  }
  /**
   * Creates an instance of CapabilityInfoDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<CapabilityInfoDtoType>) {
    return new CapabilityInfoDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<CapabilityInfoDtoType>,
  ): InstanceType<typeof CapabilityInfoDto> {
    return new CapabilityInfoDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof CapabilityInfoDto> {
    return new CapabilityInfoDto(this.toJSON());
  }
}
export abstract class CapabilityInfoDtoFactory {
  abstract create(data: unknown): CapabilityInfoDto;
}
/**
 * The base type definition for capabilityInfoDto
 **/
export type CapabilityInfoDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId: string;
  /**
   *
   * @type {string}
   **/
  name: string;
  /**
   *
   * @type {CapabilityInfoDtoType[]}
   **/
  children: CapabilityInfoDtoType[];
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace CapabilityInfoDtoType {}
