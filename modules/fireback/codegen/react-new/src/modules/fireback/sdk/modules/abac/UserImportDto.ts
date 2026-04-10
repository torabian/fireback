// @ts-nocheck 
 // This no check has been added via fireback. 
import { type PartialDeep } from "../../sdk/common/fetchx";
import { withPrefix } from "../../sdk/common/withPrefix";
/**
 * The base class definition for userImportDto
 **/
export class UserImportDto {
  /**
   *
   * @type {string}
   **/
  #avatar: string = "";
  /**
   *
   * @returns {string}
   **/
  get avatar() {
    return this.#avatar;
  }
  /**
   *
   * @type {string}
   **/
  set avatar(value: string) {
    this.#avatar = String(value);
  }
  setAvatar(value: string) {
    this.avatar = value;
    return this;
  }
  /**
   *
   * @type {UserImportDto.Passports}
   **/
  #passports: InstanceType<typeof UserImportDto.Passports>[] = [];
  /**
   *
   * @returns {UserImportDto.Passports}
   **/
  get passports() {
    return this.#passports;
  }
  /**
   *
   * @type {UserImportDto.Passports}
   **/
  set passports(value: InstanceType<typeof UserImportDto.Passports>[]) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof UserImportDto.Passports) {
      this.#passports = value;
    } else {
      this.#passports = value.map((item) => new UserImportDto.Passports(item));
    }
  }
  setPassports(value: InstanceType<typeof UserImportDto.Passports>[]) {
    this.passports = value;
    return this;
  }
  /**
   *
   * @type {UserImportDto.Address}
   **/
  #address!: InstanceType<typeof UserImportDto.Address>;
  /**
   *
   * @returns {UserImportDto.Address}
   **/
  get address() {
    return this.#address;
  }
  /**
   *
   * @type {UserImportDto.Address}
   **/
  set address(value: InstanceType<typeof UserImportDto.Address>) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof UserImportDto.Address) {
      this.#address = value;
    } else {
      this.#address = new UserImportDto.Address(value);
    }
  }
  setAddress(value: InstanceType<typeof UserImportDto.Address>) {
    this.address = value;
    return this;
  }
  /**
   * The base class definition for passports
   **/
  static Passports = class Passports {
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
      const d = data as Partial<Passports>;
      if (d.value !== undefined) {
        this.value = d.value;
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
        password: this.#password,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        value: "value",
        password: "password",
      };
    }
    /**
     * Creates an instance of UserImportDto.Passports, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: UserImportDtoType.PassportsType) {
      return new UserImportDto.Passports(possibleDtoObject);
    }
    /**
     * Creates an instance of UserImportDto.Passports, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<UserImportDtoType.PassportsType>,
    ) {
      return new UserImportDto.Passports(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<UserImportDtoType.PassportsType>,
    ): InstanceType<typeof UserImportDto.Passports> {
      return new UserImportDto.Passports({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof UserImportDto.Passports> {
      return new UserImportDto.Passports(this.toJSON());
    }
  };
  /**
   * The base class definition for address
   **/
  static Address = class Address {
    /**
     *
     * @type {string}
     **/
    #street: string = "";
    /**
     *
     * @returns {string}
     **/
    get street() {
      return this.#street;
    }
    /**
     *
     * @type {string}
     **/
    set street(value: string) {
      this.#street = String(value);
    }
    setStreet(value: string) {
      this.street = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #zipCode: string = "";
    /**
     *
     * @returns {string}
     **/
    get zipCode() {
      return this.#zipCode;
    }
    /**
     *
     * @type {string}
     **/
    set zipCode(value: string) {
      this.#zipCode = String(value);
    }
    setZipCode(value: string) {
      this.zipCode = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #city: string = "";
    /**
     *
     * @returns {string}
     **/
    get city() {
      return this.#city;
    }
    /**
     *
     * @type {string}
     **/
    set city(value: string) {
      this.#city = String(value);
    }
    setCity(value: string) {
      this.city = value;
      return this;
    }
    /**
     *
     * @type {string}
     **/
    #country: string = "";
    /**
     *
     * @returns {string}
     **/
    get country() {
      return this.#country;
    }
    /**
     *
     * @type {string}
     **/
    set country(value: string) {
      this.#country = String(value);
    }
    setCountry(value: string) {
      this.country = value;
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
      const d = data as Partial<Address>;
      if (d.street !== undefined) {
        this.street = d.street;
      }
      if (d.zipCode !== undefined) {
        this.zipCode = d.zipCode;
      }
      if (d.city !== undefined) {
        this.city = d.city;
      }
      if (d.country !== undefined) {
        this.country = d.country;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        street: this.#street,
        zipCode: this.#zipCode,
        city: this.#city,
        country: this.#country,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        street: "street",
        zipCode: "zipCode",
        city: "city",
        country: "country",
      };
    }
    /**
     * Creates an instance of UserImportDto.Address, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: UserImportDtoType.AddressType) {
      return new UserImportDto.Address(possibleDtoObject);
    }
    /**
     * Creates an instance of UserImportDto.Address, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(partialDtoObject: PartialDeep<UserImportDtoType.AddressType>) {
      return new UserImportDto.Address(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<UserImportDtoType.AddressType>,
    ): InstanceType<typeof UserImportDto.Address> {
      return new UserImportDto.Address({ ...this.toJSON(), ...partial });
    }
    clone(): InstanceType<typeof UserImportDto.Address> {
      return new UserImportDto.Address(this.toJSON());
    }
  };
  constructor(data: unknown = undefined) {
    if (data === null || data === undefined) {
      this.#lateInitFields();
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
    const d = data as Partial<UserImportDto>;
    if (d.avatar !== undefined) {
      this.avatar = d.avatar;
    }
    if (d.passports !== undefined) {
      this.passports = d.passports;
    }
    if (d.address !== undefined) {
      this.address = d.address;
    }
    this.#lateInitFields(data);
  }
  /**
   * These are the class instances, which need to be initialised, regardless of the constructor incoming data
   **/
  #lateInitFields(data = {}) {
    const d = data as Partial<UserImportDto>;
    if (!(d.address instanceof UserImportDto.Address)) {
      this.address = new UserImportDto.Address(d.address || {});
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      avatar: this.#avatar,
      passports: this.#passports,
      address: this.#address,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      avatar: "avatar",
      passports$: "passports",
      get passports() {
        return withPrefix("passports[:i]", UserImportDto.Passports.Fields);
      },
      address$: "address",
      get address() {
        return withPrefix("address", UserImportDto.Address.Fields);
      },
    };
  }
  /**
   * Creates an instance of UserImportDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: UserImportDtoType) {
    return new UserImportDto(possibleDtoObject);
  }
  /**
   * Creates an instance of UserImportDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<UserImportDtoType>) {
    return new UserImportDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<UserImportDtoType>,
  ): InstanceType<typeof UserImportDto> {
    return new UserImportDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof UserImportDto> {
    return new UserImportDto(this.toJSON());
  }
}
export abstract class UserImportDtoFactory {
  abstract create(data: unknown): UserImportDto;
}
/**
 * The base type definition for userImportDto
 **/
export type UserImportDtoType = {
  /**
   *
   * @type {string}
   **/
  avatar: string;
  /**
   *
   * @type {UserImportDtoType.PassportsType[]}
   **/
  passports: UserImportDtoType.PassportsType[];
  /**
   *
   * @type {UserImportDtoType.AddressType}
   **/
  address: UserImportDtoType.AddressType;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace UserImportDtoType {
  /**
   * The base type definition for passportsType
   **/
  export type PassportsType = {
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
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace PassportsType {}
  /**
   * The base type definition for addressType
   **/
  export type AddressType = {
    /**
     *
     * @type {string}
     **/
    street: string;
    /**
     *
     * @type {string}
     **/
    zipCode: string;
    /**
     *
     * @type {string}
     **/
    city: string;
    /**
     *
     * @type {string}
     **/
    country: string;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace AddressType {}
}
