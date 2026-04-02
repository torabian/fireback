// @ts-nocheck 
 // This no check has been added via fireback. 
import { PassportEntity } from "./PassportEntity";
import { UserEntity } from "./UserEntity";
import { UserWorkspaceEntity } from "./UserWorkspaceEntity";
import { withPrefix } from "../../sdk/common/withPrefix";
/**
 * The base class definition for userSessionDto
 **/
export class UserSessionDto {
  /**
   *
   * @type {PassportEntity}
   **/
  #passport?: PassportEntity | null = undefined;
  /**
   *
   * @returns {PassportEntity}
   **/
  get passport() {
    return this.#passport;
  }
  /**
   *
   * @type {PassportEntity}
   **/
  set passport(value: PassportEntity | null | undefined) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof PassportEntity) {
      this.#passport = value;
    } else {
      this.#passport = new PassportEntity(value);
    }
  }
  setPassport(value: PassportEntity | null | undefined) {
    this.passport = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #token: string = "";
  /**
   *
   * @returns {string}
   **/
  get token() {
    return this.#token;
  }
  /**
   *
   * @type {string}
   **/
  set token(value: string) {
    this.#token = String(value);
  }
  setToken(value: string) {
    this.token = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #exchangeKey: string = "";
  /**
   *
   * @returns {string}
   **/
  get exchangeKey() {
    return this.#exchangeKey;
  }
  /**
   *
   * @type {string}
   **/
  set exchangeKey(value: string) {
    this.#exchangeKey = String(value);
  }
  setExchangeKey(value: string) {
    this.exchangeKey = value;
    return this;
  }
  /**
   *
   * @type {UserWorkspaceEntity[]}
   **/
  #userWorkspaces: UserWorkspaceEntity[] = [];
  /**
   *
   * @returns {UserWorkspaceEntity[]}
   **/
  get userWorkspaces() {
    return this.#userWorkspaces;
  }
  /**
   *
   * @type {UserWorkspaceEntity[]}
   **/
  set userWorkspaces(value: UserWorkspaceEntity[]) {
    // For arrays, you only can pass arrays to the object
    if (!Array.isArray(value)) {
      return;
    }
    if (value.length > 0 && value[0] instanceof UserWorkspaceEntity) {
      this.#userWorkspaces = value;
    } else {
      this.#userWorkspaces = value.map((item) => new UserWorkspaceEntity(item));
    }
  }
  setUserWorkspaces(value: UserWorkspaceEntity[]) {
    this.userWorkspaces = value;
    return this;
  }
  /**
   *
   * @type {UserEntity}
   **/
  #user?: UserEntity | null = undefined;
  /**
   *
   * @returns {UserEntity}
   **/
  get user() {
    return this.#user;
  }
  /**
   *
   * @type {UserEntity}
   **/
  set user(value: UserEntity | null | undefined) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof UserEntity) {
      this.#user = value;
    } else {
      this.#user = new UserEntity(value);
    }
  }
  setUser(value: UserEntity | null | undefined) {
    this.user = value;
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
    const d = data as Partial<UserSessionDto>;
    if (d.passport !== undefined) {
      this.passport = d.passport;
    }
    if (d.token !== undefined) {
      this.token = d.token;
    }
    if (d.exchangeKey !== undefined) {
      this.exchangeKey = d.exchangeKey;
    }
    if (d.userWorkspaces !== undefined) {
      this.userWorkspaces = d.userWorkspaces;
    }
    if (d.user !== undefined) {
      this.user = d.user;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      passport: this.#passport,
      token: this.#token,
      exchangeKey: this.#exchangeKey,
      userWorkspaces: this.#userWorkspaces,
      user: this.#user,
      userId: this.#userId,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      passport: "passport",
      token: "token",
      exchangeKey: "exchangeKey",
      userWorkspaces$: "userWorkspaces",
      get userWorkspaces() {
        return withPrefix("userWorkspaces[:i]", UserWorkspaceEntity.Fields);
      },
      user: "user",
      userId: "userId",
    };
  }
  /**
   * Creates an instance of UserSessionDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: UserSessionDtoType) {
    return new UserSessionDto(possibleDtoObject);
  }
  /**
   * Creates an instance of UserSessionDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<UserSessionDtoType>) {
    return new UserSessionDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<UserSessionDtoType>,
  ): InstanceType<typeof UserSessionDto> {
    return new UserSessionDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof UserSessionDto> {
    return new UserSessionDto(this.toJSON());
  }
}
export abstract class UserSessionDtoFactory {
  abstract create(data: unknown): UserSessionDto;
}
type PartialDeep<T> = {
  [P in keyof T]?: T[P] extends Array<infer U>
    ? Array<PartialDeep<U>>
    : T[P] extends object
      ? PartialDeep<T[P]>
      : T[P];
};
/**
 * The base type definition for userSessionDto
 **/
export type UserSessionDtoType = {
  /**
   *
   * @type {PassportEntity}
   **/
  passport?: PassportEntity;
  /**
   *
   * @type {string}
   **/
  token: string;
  /**
   *
   * @type {string}
   **/
  exchangeKey: string;
  /**
   *
   * @type {UserWorkspaceEntity[]}
   **/
  userWorkspaces: UserWorkspaceEntity[];
  /**
   *
   * @type {UserEntity}
   **/
  user?: UserEntity;
  /**
   *
   * @type {string}
   **/
  userId?: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace UserSessionDtoType {}
