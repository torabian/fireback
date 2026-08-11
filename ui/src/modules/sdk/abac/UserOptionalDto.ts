import { type PartialDeep } from "../sdk/common/fetchx";
import { withPrefix } from "../sdk/common/withPrefix";
/**
 * The base class definition for userOptionalDto
 **/
export class UserOptionalDto {
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
  #firstName?: string | null = undefined;
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
  set firstName(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#firstName = correctType ? value : String(value);
  }
  setFirstName(value: string | null | undefined) {
    this.firstName = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #lastName?: string | null = undefined;
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
  set lastName(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#lastName = correctType ? value : String(value);
  }
  setLastName(value: string | null | undefined) {
    this.lastName = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #photo?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get photo() {
    return this.#photo;
  }
  /**
   *
   * @type {string}
   **/
  set photo(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#photo = correctType ? value : String(value);
  }
  setPhoto(value: string | null | undefined) {
    this.photo = value;
    return this;
  }
  /**
   *
   * @type {number}
   **/
  #gender?: number | null = undefined;
  /**
   *
   * @returns {number}
   **/
  get gender() {
    return this.#gender;
  }
  /**
   *
   * @type {number}
   **/
  set gender(value: number | null | undefined) {
    const correctType =
      typeof value === "number" || value === undefined || value === null;
    const parsedValue = correctType ? value : Number(value);
    if (!Number.isNaN(parsedValue)) {
      this.#gender = parsedValue;
    }
  }
  setGender(value: number | null | undefined) {
    this.gender = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #title?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get title() {
    return this.#title;
  }
  /**
   *
   * @type {string}
   **/
  set title(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#title = correctType ? value : String(value);
  }
  setTitle(value: string | null | undefined) {
    this.title = value;
    return this;
  }
  /**
   *
   * @type {XDate}
   **/
  #birthDate!: XDate;
  /**
   *
   * @returns {XDate}
   **/
  get birthDate() {
    return this.#birthDate;
  }
  /**
   *
   * @type {XDate}
   **/
  set birthDate(value: XDate) {
    this.#birthDate = value;
  }
  setBirthDate(value: XDate) {
    this.birthDate = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #avatar?: string | null = undefined;
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
  set avatar(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#avatar = correctType ? value : String(value);
  }
  setAvatar(value: string | null | undefined) {
    this.avatar = value;
    return this;
  }
  /**
   * User last connecting ip address
   * @type {string}
   **/
  #lastIpAddress?: string | null = undefined;
  /**
   * User last connecting ip address
   * @returns {string}
   **/
  get lastIpAddress() {
    return this.#lastIpAddress;
  }
  /**
   * User last connecting ip address
   * @type {string}
   **/
  set lastIpAddress(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#lastIpAddress = correctType ? value : String(value);
  }
  setLastIpAddress(value: string | null | undefined) {
    this.lastIpAddress = value;
    return this;
  }
  /**
   * User primary address location. Can be useful for simple projects that a user is associated with a single address.
   * @type {UserOptionalDto.PrimaryAddress}
   **/
  #primaryAddress?:
    | InstanceType<typeof UserOptionalDto.PrimaryAddress>
    | null
    | undefined
    | null = undefined;
  /**
   * User primary address location. Can be useful for simple projects that a user is associated with a single address.
   * @returns {UserOptionalDto.PrimaryAddress}
   **/
  get primaryAddress() {
    return this.#primaryAddress;
  }
  /**
   * User primary address location. Can be useful for simple projects that a user is associated with a single address.
   * @type {UserOptionalDto.PrimaryAddress}
   **/
  set primaryAddress(
    value:
      | InstanceType<typeof UserOptionalDto.PrimaryAddress>
      | null
      | undefined
      | null
      | undefined,
  ) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof UserOptionalDto.PrimaryAddress) {
      this.#primaryAddress = value;
    } else {
      this.#primaryAddress = new UserOptionalDto.PrimaryAddress(value);
    }
  }
  setPrimaryAddress(
    value:
      | InstanceType<typeof UserOptionalDto.PrimaryAddress>
      | null
      | undefined
      | null
      | undefined,
  ) {
    this.primaryAddress = value;
    return this;
  }
  /**
   * Contact phone number for this user (separate from any passport used to sign in).
   * @type {string}
   **/
  #phoneNumber?: string | null = undefined;
  /**
   * Contact phone number for this user (separate from any passport used to sign in).
   * @returns {string}
   **/
  get phoneNumber() {
    return this.#phoneNumber;
  }
  /**
   * Contact phone number for this user (separate from any passport used to sign in).
   * @type {string}
   **/
  set phoneNumber(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#phoneNumber = correctType ? value : String(value);
  }
  setPhoneNumber(value: string | null | undefined) {
    this.phoneNumber = value;
    return this;
  }
  /**
   * The user's job title/role, e.g. "Support Engineer".
   * @type {string}
   **/
  #jobTitle?: string | null = undefined;
  /**
   * The user's job title/role, e.g. "Support Engineer".
   * @returns {string}
   **/
  get jobTitle() {
    return this.#jobTitle;
  }
  /**
   * The user's job title/role, e.g. "Support Engineer".
   * @type {string}
   **/
  set jobTitle(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#jobTitle = correctType ? value : String(value);
  }
  setJobTitle(value: string | null | undefined) {
    this.jobTitle = value;
    return this;
  }
  /**
   * The company or organization the user is associated with.
   * @type {string}
   **/
  #company?: string | null = undefined;
  /**
   * The company or organization the user is associated with.
   * @returns {string}
   **/
  get company() {
    return this.#company;
  }
  /**
   * The company or organization the user is associated with.
   * @type {string}
   **/
  set company(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#company = correctType ? value : String(value);
  }
  setCompany(value: string | null | undefined) {
    this.company = value;
    return this;
  }
  /**
   * Free-form biography/notes about the user.
   * @type {string}
   **/
  #bio?: string | null = undefined;
  /**
   * Free-form biography/notes about the user.
   * @returns {string}
   **/
  get bio() {
    return this.#bio;
  }
  /**
   * Free-form biography/notes about the user.
   * @type {string}
   **/
  set bio(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#bio = correctType ? value : String(value);
  }
  setBio(value: string | null | undefined) {
    this.bio = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #userId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   *
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
  /**
   * The base class definition for primaryAddress
   **/
  static PrimaryAddress = class PrimaryAddress {
    /**
     * Street address, building number
     * @type {string}
     **/
    #addressLine1?: string | null = undefined;
    /**
     * Street address, building number
     * @returns {string}
     **/
    get addressLine1() {
      return this.#addressLine1;
    }
    /**
     * Street address, building number
     * @type {string}
     **/
    set addressLine1(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#addressLine1 = correctType ? value : String(value);
    }
    setAddressLine1(value: string | null | undefined) {
      this.addressLine1 = value;
      return this;
    }
    /**
     * Apartment, suite, floor (optional)
     * @type {string}
     **/
    #addressLine2?: string | null = undefined;
    /**
     * Apartment, suite, floor (optional)
     * @returns {string}
     **/
    get addressLine2() {
      return this.#addressLine2;
    }
    /**
     * Apartment, suite, floor (optional)
     * @type {string}
     **/
    set addressLine2(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#addressLine2 = correctType ? value : String(value);
    }
    setAddressLine2(value: string | null | undefined) {
      this.addressLine2 = value;
      return this;
    }
    /**
     * City or locality
     * @type {string}
     **/
    #city?: string | null = undefined;
    /**
     * City or locality
     * @returns {string}
     **/
    get city() {
      return this.#city;
    }
    /**
     * City or locality
     * @type {string}
     **/
    set city(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#city = correctType ? value : String(value);
    }
    setCity(value: string | null | undefined) {
      this.city = value;
      return this;
    }
    /**
     * State, region, or province
     * @type {string}
     **/
    #stateOrProvince?: string | null = undefined;
    /**
     * State, region, or province
     * @returns {string}
     **/
    get stateOrProvince() {
      return this.#stateOrProvince;
    }
    /**
     * State, region, or province
     * @type {string}
     **/
    set stateOrProvince(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#stateOrProvince = correctType ? value : String(value);
    }
    setStateOrProvince(value: string | null | undefined) {
      this.stateOrProvince = value;
      return this;
    }
    /**
     * ZIP or postal code
     * @type {string}
     **/
    #postalCode?: string | null = undefined;
    /**
     * ZIP or postal code
     * @returns {string}
     **/
    get postalCode() {
      return this.#postalCode;
    }
    /**
     * ZIP or postal code
     * @type {string}
     **/
    set postalCode(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#postalCode = correctType ? value : String(value);
    }
    setPostalCode(value: string | null | undefined) {
      this.postalCode = value;
      return this;
    }
    /**
     * ISO 3166-1 alpha-2 (e.g., "US", "DE")
     * @type {string}
     **/
    #countryCode?: string | null = undefined;
    /**
     * ISO 3166-1 alpha-2 (e.g., "US", "DE")
     * @returns {string}
     **/
    get countryCode() {
      return this.#countryCode;
    }
    /**
     * ISO 3166-1 alpha-2 (e.g., "US", "DE")
     * @type {string}
     **/
    set countryCode(value: string | null | undefined) {
      const correctType =
        typeof value === "string" || value === undefined || value === null;
      this.#countryCode = correctType ? value : String(value);
    }
    setCountryCode(value: string | null | undefined) {
      this.countryCode = value;
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
      const d = data as Partial<PrimaryAddress>;
      if (d.addressLine1 !== undefined) {
        this.addressLine1 = d.addressLine1;
      }
      if (d.addressLine2 !== undefined) {
        this.addressLine2 = d.addressLine2;
      }
      if (d.city !== undefined) {
        this.city = d.city;
      }
      if (d.stateOrProvince !== undefined) {
        this.stateOrProvince = d.stateOrProvince;
      }
      if (d.postalCode !== undefined) {
        this.postalCode = d.postalCode;
      }
      if (d.countryCode !== undefined) {
        this.countryCode = d.countryCode;
      }
    }
    /**
     *	Special toJSON override, since the field are private,
     *	Json stringify won't see them unless we mention it explicitly.
     **/
    toJSON() {
      return {
        addressLine1: this.#addressLine1,
        addressLine2: this.#addressLine2,
        city: this.#city,
        stateOrProvince: this.#stateOrProvince,
        postalCode: this.#postalCode,
        countryCode: this.#countryCode,
      };
    }
    toString() {
      return JSON.stringify(this);
    }
    static get Fields() {
      return {
        addressLine1: "addressLine1",
        addressLine2: "addressLine2",
        city: "city",
        stateOrProvince: "stateOrProvince",
        postalCode: "postalCode",
        countryCode: "countryCode",
      };
    }
    /**
     * Creates an instance of UserOptionalDto.PrimaryAddress, and possibleDtoObject
     * needs to satisfy the type requirement fully, otherwise typescript compile would
     * be complaining.
     **/
    static from(possibleDtoObject: UserOptionalDtoType.PrimaryAddressType) {
      return new UserOptionalDto.PrimaryAddress(possibleDtoObject);
    }
    /**
     * Creates an instance of UserOptionalDto.PrimaryAddress, and partialDtoObject
     * needs to satisfy the type, but partially, and rest of the content would
     * be constructed according to data types and nullability.
     **/
    static with(
      partialDtoObject: PartialDeep<UserOptionalDtoType.PrimaryAddressType>,
    ) {
      return new UserOptionalDto.PrimaryAddress(partialDtoObject);
    }
    copyWith(
      partial: PartialDeep<UserOptionalDtoType.PrimaryAddressType>,
    ): InstanceType<typeof UserOptionalDto.PrimaryAddress> {
      return new UserOptionalDto.PrimaryAddress({
        ...this.toJSON(),
        ...partial,
      });
    }
    clone(): InstanceType<typeof UserOptionalDto.PrimaryAddress> {
      return new UserOptionalDto.PrimaryAddress(this.toJSON());
    }
  };
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
    const d = data as Partial<UserOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.firstName !== undefined) {
      this.firstName = d.firstName;
    }
    if (d.lastName !== undefined) {
      this.lastName = d.lastName;
    }
    if (d.photo !== undefined) {
      this.photo = d.photo;
    }
    if (d.gender !== undefined) {
      this.gender = d.gender;
    }
    if (d.title !== undefined) {
      this.title = d.title;
    }
    if (d.birthDate !== undefined) {
      this.birthDate = d.birthDate;
    }
    if (d.avatar !== undefined) {
      this.avatar = d.avatar;
    }
    if (d.lastIpAddress !== undefined) {
      this.lastIpAddress = d.lastIpAddress;
    }
    if (d.primaryAddress !== undefined) {
      this.primaryAddress = d.primaryAddress;
    }
    if (d.phoneNumber !== undefined) {
      this.phoneNumber = d.phoneNumber;
    }
    if (d.jobTitle !== undefined) {
      this.jobTitle = d.jobTitle;
    }
    if (d.company !== undefined) {
      this.company = d.company;
    }
    if (d.bio !== undefined) {
      this.bio = d.bio;
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
      firstName: this.#firstName,
      lastName: this.#lastName,
      photo: this.#photo,
      gender: this.#gender,
      title: this.#title,
      birthDate: this.#birthDate,
      avatar: this.#avatar,
      lastIpAddress: this.#lastIpAddress,
      primaryAddress: this.#primaryAddress,
      phoneNumber: this.#phoneNumber,
      jobTitle: this.#jobTitle,
      company: this.#company,
      bio: this.#bio,
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
      firstName: "firstName",
      lastName: "lastName",
      photo: "photo",
      gender: "gender",
      title: "title",
      birthDate: "birthDate",
      avatar: "avatar",
      lastIpAddress: "lastIpAddress",
      primaryAddress$: "primaryAddress",
      get primaryAddress() {
        return withPrefix(
          "primaryAddress",
          UserOptionalDto.PrimaryAddress.Fields,
        );
      },
      phoneNumber: "phoneNumber",
      jobTitle: "jobTitle",
      company: "company",
      bio: "bio",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of UserOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: UserOptionalDtoType) {
    return new UserOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of UserOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<UserOptionalDtoType>) {
    return new UserOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<UserOptionalDtoType>,
  ): InstanceType<typeof UserOptionalDto> {
    return new UserOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof UserOptionalDto> {
    return new UserOptionalDto(this.toJSON());
  }
}
export abstract class UserOptionalDtoFactory {
  abstract create(data: unknown): UserOptionalDto;
}
/**
 * The base type definition for userOptionalDto
 **/
export type UserOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   *
   * @type {string}
   **/
  firstName?: string;
  /**
   *
   * @type {string}
   **/
  lastName?: string;
  /**
   *
   * @type {string}
   **/
  photo?: string;
  /**
   *
   * @type {number}
   **/
  gender?: number;
  /**
   *
   * @type {string}
   **/
  title?: string;
  /**
   *
   * @type {XDate}
   **/
  birthDate: XDate;
  /**
   *
   * @type {string}
   **/
  avatar?: string;
  /**
   * User last connecting ip address
   * @type {string}
   **/
  lastIpAddress?: string;
  /**
   * User primary address location. Can be useful for simple projects that a user is associated with a single address.
   * @type {UserOptionalDtoType.PrimaryAddressType}
   **/
  primaryAddress?: UserOptionalDtoType.PrimaryAddressType;
  /**
   * Contact phone number for this user (separate from any passport used to sign in).
   * @type {string}
   **/
  phoneNumber?: string;
  /**
   * The user's job title/role, e.g. "Support Engineer".
   * @type {string}
   **/
  jobTitle?: string;
  /**
   * The company or organization the user is associated with.
   * @type {string}
   **/
  company?: string;
  /**
   * Free-form biography/notes about the user.
   * @type {string}
   **/
  bio?: string;
  /**
   *
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
export namespace UserOptionalDtoType {
  /**
   * The base type definition for primaryAddressType
   **/
  export type PrimaryAddressType = {
    /**
     * Street address, building number
     * @type {string}
     **/
    addressLine1?: string;
    /**
     * Apartment, suite, floor (optional)
     * @type {string}
     **/
    addressLine2?: string;
    /**
     * City or locality
     * @type {string}
     **/
    city?: string;
    /**
     * State, region, or province
     * @type {string}
     **/
    stateOrProvince?: string;
    /**
     * ZIP or postal code
     * @type {string}
     **/
    postalCode?: string;
    /**
     * ISO 3166-1 alpha-2 (e.g., "US", "DE")
     * @type {string}
     **/
    countryCode?: string;
  };
  // eslint-disable-next-line @typescript-eslint/no-namespace
  export namespace PrimaryAddressType {}
}
