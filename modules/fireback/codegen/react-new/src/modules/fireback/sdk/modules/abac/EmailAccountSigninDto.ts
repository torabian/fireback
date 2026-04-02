// @ts-nocheck 
 // This no check has been added via fireback. 
/**
 * The base class definition for emailAccountSigninDto
 **/
export class EmailAccountSigninDto {
  /**
   *
   * @type {string}
   **/
  #email: string = "";
  /**
   *
   * @returns {string}
   **/
  get email() {
    return this.#email;
  }
  /**
   *
   * @type {string}
   **/
  set email(value: string) {
    this.#email = String(value);
  }
  setEmail(value: string) {
    this.email = value;
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
    const d = data as Partial<EmailAccountSigninDto>;
    if (d.email !== undefined) {
      this.email = d.email;
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
      email: this.#email,
      password: this.#password,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      email: "email",
      password: "password",
    };
  }
  /**
   * Creates an instance of EmailAccountSigninDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: EmailAccountSigninDtoType) {
    return new EmailAccountSigninDto(possibleDtoObject);
  }
  /**
   * Creates an instance of EmailAccountSigninDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<EmailAccountSigninDtoType>) {
    return new EmailAccountSigninDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<EmailAccountSigninDtoType>,
  ): InstanceType<typeof EmailAccountSigninDto> {
    return new EmailAccountSigninDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof EmailAccountSigninDto> {
    return new EmailAccountSigninDto(this.toJSON());
  }
}
export abstract class EmailAccountSigninDtoFactory {
  abstract create(data: unknown): EmailAccountSigninDto;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for emailAccountSigninDto
 **/
export type EmailAccountSigninDtoType = {
  /**
   *
   * @type {string}
   **/
  email: string;
  /**
   *
   * @type {string}
   **/
  password: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace EmailAccountSigninDtoType {}
