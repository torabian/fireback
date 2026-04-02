// @ts-nocheck 
 // This no check has been added via fireback. 
/**
 * The base class definition for classicAuthDto
 **/
export class ClassicAuthDto {
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
  /**
   *
   * @type {string}
   **/
  #firstName: string = "";
  /**
   *
   * @returns {string}
   **/
  get firstName() {
    return this.#firstName;
  }
  /**
   *
   * @type {string}
   **/
  set firstName(value: string) {
    this.#firstName = String(value);
  }
  setFirstName(value: string) {
    this.firstName = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #lastName: string = "";
  /**
   *
   * @returns {string}
   **/
  get lastName() {
    return this.#lastName;
  }
  /**
   *
   * @type {string}
   **/
  set lastName(value: string) {
    this.#lastName = String(value);
  }
  setLastName(value: string) {
    this.lastName = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteId: string = "";
  /**
   *
   * @returns {string}
   **/
  get inviteId() {
    return this.#inviteId;
  }
  /**
   *
   * @type {string}
   **/
  set inviteId(value: string) {
    this.#inviteId = String(value);
  }
  setInviteId(value: string) {
    this.inviteId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #publicJoinKeyId: string = "";
  /**
   *
   * @returns {string}
   **/
  get publicJoinKeyId() {
    return this.#publicJoinKeyId;
  }
  /**
   *
   * @type {string}
   **/
  set publicJoinKeyId(value: string) {
    this.#publicJoinKeyId = String(value);
  }
  setPublicJoinKeyId(value: string) {
    this.publicJoinKeyId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #workspaceTypeId: string = "";
  /**
   *
   * @returns {string}
   **/
  get workspaceTypeId() {
    return this.#workspaceTypeId;
  }
  /**
   *
   * @type {string}
   **/
  set workspaceTypeId(value: string) {
    this.#workspaceTypeId = String(value);
  }
  setWorkspaceTypeId(value: string) {
    this.workspaceTypeId = value;
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
    const d = data as Partial<ClassicAuthDto>;
    if (d.value !== undefined) {
      this.value = d.value;
    }
    if (d.password !== undefined) {
      this.password = d.password;
    }
    if (d.firstName !== undefined) {
      this.firstName = d.firstName;
    }
    if (d.lastName !== undefined) {
      this.lastName = d.lastName;
    }
    if (d.inviteId !== undefined) {
      this.inviteId = d.inviteId;
    }
    if (d.publicJoinKeyId !== undefined) {
      this.publicJoinKeyId = d.publicJoinKeyId;
    }
    if (d.workspaceTypeId !== undefined) {
      this.workspaceTypeId = d.workspaceTypeId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      value: this.#value,
      password: this.#password,
      firstName: this.#firstName,
      lastName: this.#lastName,
      inviteId: this.#inviteId,
      publicJoinKeyId: this.#publicJoinKeyId,
      workspaceTypeId: this.#workspaceTypeId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      value: "value",
      password: "password",
      firstName: "firstName",
      lastName: "lastName",
      inviteId: "inviteId",
      publicJoinKeyId: "publicJoinKeyId",
      workspaceTypeId: "workspaceTypeId",
    };
  }
  /**
   * Creates an instance of ClassicAuthDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: ClassicAuthDtoType) {
    return new ClassicAuthDto(possibleDtoObject);
  }
  /**
   * Creates an instance of ClassicAuthDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<ClassicAuthDtoType>) {
    return new ClassicAuthDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<ClassicAuthDtoType>,
  ): InstanceType<typeof ClassicAuthDto> {
    return new ClassicAuthDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof ClassicAuthDto> {
    return new ClassicAuthDto(this.toJSON());
  }
}
export abstract class ClassicAuthDtoFactory {
  abstract create(data: unknown): ClassicAuthDto;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for classicAuthDto
 **/
export type ClassicAuthDtoType = {
  /**
   *
   * @type {string}
   **/
  value: string;
  /**
   *
   * @type {string}
   **/
  password: string;
  /**
   *
   * @type {string}
   **/
  firstName: string;
  /**
   *
   * @type {string}
   **/
  lastName: string;
  /**
   *
   * @type {string}
   **/
  inviteId: string;
  /**
   *
   * @type {string}
   **/
  publicJoinKeyId: string;
  /**
   *
   * @type {string}
   **/
  workspaceTypeId: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace ClassicAuthDtoType {}
