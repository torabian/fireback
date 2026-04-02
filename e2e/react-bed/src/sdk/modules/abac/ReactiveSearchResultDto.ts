// @ts-nocheck 
 // This no check has been added via fireback. 
/**
 * The base class definition for reactiveSearchResultDto
 **/
export class ReactiveSearchResultDto {
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
  #phrase: string = "";
  /**
   *
   * @returns {string}
   **/
  get phrase() {
    return this.#phrase;
  }
  /**
   *
   * @type {string}
   **/
  set phrase(value: string) {
    this.#phrase = String(value);
  }
  setPhrase(value: string) {
    this.phrase = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #icon: string = "";
  /**
   *
   * @returns {string}
   **/
  get icon() {
    return this.#icon;
  }
  /**
   *
   * @type {string}
   **/
  set icon(value: string) {
    this.#icon = String(value);
  }
  setIcon(value: string) {
    this.icon = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #description: string = "";
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
  set description(value: string) {
    this.#description = String(value);
  }
  setDescription(value: string) {
    this.description = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #group: string = "";
  /**
   *
   * @returns {string}
   **/
  get group() {
    return this.#group;
  }
  /**
   *
   * @type {string}
   **/
  set group(value: string) {
    this.#group = String(value);
  }
  setGroup(value: string) {
    this.group = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #uiLocation: string = "";
  /**
   *
   * @returns {string}
   **/
  get uiLocation() {
    return this.#uiLocation;
  }
  /**
   *
   * @type {string}
   **/
  set uiLocation(value: string) {
    this.#uiLocation = String(value);
  }
  setUiLocation(value: string) {
    this.uiLocation = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #actionFn: string = "";
  /**
   *
   * @returns {string}
   **/
  get actionFn() {
    return this.#actionFn;
  }
  /**
   *
   * @type {string}
   **/
  set actionFn(value: string) {
    this.#actionFn = String(value);
  }
  setActionFn(value: string) {
    this.actionFn = value;
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
    const d = data as Partial<ReactiveSearchResultDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.phrase !== undefined) {
      this.phrase = d.phrase;
    }
    if (d.icon !== undefined) {
      this.icon = d.icon;
    }
    if (d.description !== undefined) {
      this.description = d.description;
    }
    if (d.group !== undefined) {
      this.group = d.group;
    }
    if (d.uiLocation !== undefined) {
      this.uiLocation = d.uiLocation;
    }
    if (d.actionFn !== undefined) {
      this.actionFn = d.actionFn;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      phrase: this.#phrase,
      icon: this.#icon,
      description: this.#description,
      group: this.#group,
      uiLocation: this.#uiLocation,
      actionFn: this.#actionFn,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      phrase: "phrase",
      icon: "icon",
      description: "description",
      group: "group",
      uiLocation: "uiLocation",
      actionFn: "actionFn",
    };
  }
  /**
   * Creates an instance of ReactiveSearchResultDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ReactiveSearchResultDtoType) {
    return new ReactiveSearchResultDto(possibleDtoObject);
  }
  /**
   * Creates an instance of ReactiveSearchResultDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ReactiveSearchResultDtoType>) {
    return new ReactiveSearchResultDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ReactiveSearchResultDtoType>,
  ): InstanceType<typeof ReactiveSearchResultDto> {
    return new ReactiveSearchResultDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ReactiveSearchResultDto> {
    return new ReactiveSearchResultDto(this.toJSON());
  }
}
export abstract class ReactiveSearchResultDtoFactory {
  abstract create(data: unknown): ReactiveSearchResultDto;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for reactiveSearchResultDto
 **/
export type ReactiveSearchResultDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId: string;
  /**
   *
   * @type {string}
   **/
  phrase: string;
  /**
   *
   * @type {string}
   **/
  icon: string;
  /**
   *
   * @type {string}
   **/
  description: string;
  /**
   *
   * @type {string}
   **/
  group: string;
  /**
   *
   * @type {string}
   **/
  uiLocation: string;
  /**
   *
   * @type {string}
   **/
  actionFn: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ReactiveSearchResultDtoType {}
