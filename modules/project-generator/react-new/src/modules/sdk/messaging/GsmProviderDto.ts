import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for gsmProviderDto
 **/
export class GsmProviderDto {
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
   *
   * @type {string}
   **/
  #apiKey: string = "";
  /**
   *
   * @returns {string}
   **/
  get apiKey() {
    return this.#apiKey;
  }
  /**
   *
   * @type {string}
   **/
  set apiKey(value: string) {
    this.#apiKey = String(value);
  }
  setApiKey(value: string) {
    this.apiKey = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #mainSenderNumber: string = "";
  /**
   *
   * @returns {string}
   **/
  get mainSenderNumber() {
    return this.#mainSenderNumber;
  }
  /**
   *
   * @type {string}
   **/
  set mainSenderNumber(value: string) {
    this.#mainSenderNumber = String(value);
  }
  setMainSenderNumber(value: string) {
    this.mainSenderNumber = value;
    return this;
  }
  /**
   *
   * @type {"url" | "terminal" | "mediana"}
   **/
  #type!: "url" | "terminal" | "mediana";
  /**
   *
   * @returns {"url" | "terminal" | "mediana"}
   **/
  get type() {
    return this.#type;
  }
  /**
   *
   * @type {"url" | "terminal" | "mediana"}
   **/
  set type(value: "url" | "terminal" | "mediana") {
    this.#type = value;
  }
  setType(value: "url" | "terminal" | "mediana") {
    this.type = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #invokeUrl: string = "";
  /**
   *
   * @returns {string}
   **/
  get invokeUrl() {
    return this.#invokeUrl;
  }
  /**
   *
   * @type {string}
   **/
  set invokeUrl(value: string) {
    this.#invokeUrl = String(value);
  }
  setInvokeUrl(value: string) {
    this.invokeUrl = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #invokeBody: string = "";
  /**
   *
   * @returns {string}
   **/
  get invokeBody() {
    return this.#invokeBody;
  }
  /**
   *
   * @type {string}
   **/
  set invokeBody(value: string) {
    this.#invokeBody = String(value);
  }
  setInvokeBody(value: string) {
    this.invokeBody = value;
    return this;
  }
  /**
   * The unique-id of the workspace which content belongs to.
   * @type {string}
   **/
  #workspaceId?: string | null = undefined;
  /**
   * The unique-id of the workspace which content belongs to.
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   * The unique-id of the workspace which content belongs to.
   * @type {string}
   **/
  set workspaceId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#workspaceId = correctType ? value : String(value);
  }
  setWorkspaceId(value: string | null | undefined) {
    this.workspaceId = value;
    return this;
  }
  /**
   * The unique-id of the user which created/owns the record.
   * @type {string}
   **/
  #userId?: string | null = undefined;
  /**
   * The unique-id of the user which created/owns the record.
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * The unique-id of the user which created/owns the record.
   * @type {string}
   **/
  set userId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#userId = correctType ? value : String(value);
  }
  setUserId(value: string | null | undefined) {
    this.userId = value;
    return this;
  }
  /**
   *
   * @type {PlainTime}
   **/
  #createdAt!: PlainTime;
  /**
   *
   * @returns {PlainTime}
   **/
  get createdAt() {
    return this.#createdAt;
  }
  /**
   *
   * @type {PlainTime}
   **/
  set createdAt(value: PlainTime) {
    this.#createdAt = value;
  }
  setCreatedAt(value: PlainTime) {
    this.createdAt = value;
    return this;
  }
  /**
   *
   * @type {PlainTime}
   **/
  #updatedAt!: PlainTime;
  /**
   *
   * @returns {PlainTime}
   **/
  get updatedAt() {
    return this.#updatedAt;
  }
  /**
   *
   * @type {PlainTime}
   **/
  set updatedAt(value: PlainTime) {
    this.#updatedAt = value;
  }
  setUpdatedAt(value: PlainTime) {
    this.updatedAt = value;
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
    const d = data as Partial<GsmProviderDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.apiKey !== undefined) {
      this.apiKey = d.apiKey;
    }
    if (d.mainSenderNumber !== undefined) {
      this.mainSenderNumber = d.mainSenderNumber;
    }
    if (d.type !== undefined) {
      this.type = d.type;
    }
    if (d.invokeUrl !== undefined) {
      this.invokeUrl = d.invokeUrl;
    }
    if (d.invokeBody !== undefined) {
      this.invokeBody = d.invokeBody;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.createdAt !== undefined) {
      this.createdAt = d.createdAt;
    }
    if (d.updatedAt !== undefined) {
      this.updatedAt = d.updatedAt;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      uniqueId: this.#uniqueId,
      apiKey: this.#apiKey,
      mainSenderNumber: this.#mainSenderNumber,
      type: this.#type,
      invokeUrl: this.#invokeUrl,
      invokeBody: this.#invokeBody,
      workspaceId: this.#workspaceId,
      userId: this.#userId,
      createdAt: this.#createdAt,
      updatedAt: this.#updatedAt,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      uniqueId: "uniqueId",
      apiKey: "apiKey",
      mainSenderNumber: "mainSenderNumber",
      type: "type",
      invokeUrl: "invokeUrl",
      invokeBody: "invokeBody",
      workspaceId: "workspaceId",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of GsmProviderDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: GsmProviderDtoType) {
    return new GsmProviderDto(possibleDtoObject);
  }
  /**
   * Creates an instance of GsmProviderDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<GsmProviderDtoType>) {
    return new GsmProviderDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<GsmProviderDtoType>,
  ): InstanceType<typeof GsmProviderDto> {
    return new GsmProviderDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof GsmProviderDto> {
    return new GsmProviderDto(this.toJSON());
  }
}
export abstract class GsmProviderDtoFactory {
  abstract create(data: unknown): GsmProviderDto;
}
/**
 * The base type definition for gsmProviderDto
 **/
export type GsmProviderDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   *
   * @type {string}
   **/
  apiKey: string;
  /**
   *
   * @type {string}
   **/
  mainSenderNumber: string;
  /**
   *
   * @type {"url" | "terminal" | "mediana"}
   **/
  type: "url" | "terminal" | "mediana";
  /**
   *
   * @type {string}
   **/
  invokeUrl: string;
  /**
   *
   * @type {string}
   **/
  invokeBody: string;
  /**
   * The unique-id of the workspace which content belongs to.
   * @type {string}
   **/
  workspaceId?: string;
  /**
   * The unique-id of the user which created/owns the record.
   * @type {string}
   **/
  userId?: string;
  /**
   *
   * @type {PlainTime}
   **/
  createdAt: PlainTime;
  /**
   *
   * @type {PlainTime}
   **/
  updatedAt: PlainTime;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace GsmProviderDtoType {}
