import { MCollection, MOne } from "@fireback/js-remote-ctx/common/operators";
import { PassportDto as PassportEntity } from "./PassportDto";
import { UserDto as UserEntity } from "./UserDto";
import { UserWorkspaceDto as UserWorkspaceEntity } from "./UserWorkspaceDto";
import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
import { withPrefix } from "@fireback/js-remote-ctx/common/withPrefix";
/**
 * The base class definition for userSessionDto
 **/
export class UserSessionDto {
  /**
   *
   * @type {PassportEntity}
   **/
  #passport?: MOne<PassportEntity> | null = undefined;
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
  set passport(
    value:
      | MOne<PassportEntity>
      | InstanceType<typeof PassportEntity>
      | null
      | undefined,
  ) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof MOne) {
      this.#passport = value;
    } else if (value instanceof PassportEntity) {
      this.#passport = MOne.of(value);
    } else {
      this.#passport = MOne.of(new PassportEntity(value));
    }
  }
  setPassport(
    value:
      | MOne<PassportEntity>
      | InstanceType<typeof PassportEntity>
      | null
      | undefined,
  ) {
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
  #userWorkspaces: MCollection<UserWorkspaceEntity> = MCollection.of([]);
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
  set userWorkspaces(
    value:
      | MCollection<UserWorkspaceEntity>
      | InstanceType<typeof UserWorkspaceEntity>[],
  ) {
    // When the passed value is already an array, we check if we need to
    // cast the inner items into class instance.
    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof UserWorkspaceEntity) {
        this.#userWorkspaces = MCollection.of(value);
      } else {
        this.#userWorkspaces = MCollection.of(
          value.map((item) => new UserWorkspaceEntity(item)),
        );
      }
      return;
    }
    // If the instance is already an MCollection, we assume it's all good.
    if (value instanceof MCollection) {
      this.#userWorkspaces = value;
      return;
    }
    // If the value is not array, and is not a MCollection, we need to be consider,
    // it might be eligible to be casted into MCollection.
    const { ok, value: mcastValue } = MCollection.cast<unknown>(value);
    if (ok) {
      this.#userWorkspaces = mcastValue as any;
      return;
    }
    console.warn(
      "Cannot assing value to userWorkspaces, because it needs MCollection instance or an Array.",
    );
  }
  setUserWorkspaces(
    value:
      | MCollection<UserWorkspaceEntity>
      | InstanceType<typeof UserWorkspaceEntity>[],
  ) {
    this.userWorkspaces = value;
    return this;
  }
  /**
   *
   * @type {UserEntity}
   **/
  #user?: MOne<UserEntity> | null = undefined;
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
  set user(
    value:
      | MOne<UserEntity>
      | InstanceType<typeof UserEntity>
      | null
      | undefined,
  ) {
    // For objects, the sub type needs to always be instance of the sub class.
    if (value instanceof MOne) {
      this.#user = value;
    } else if (value instanceof UserEntity) {
      this.#user = MOne.of(value);
    } else {
      this.#user = MOne.of(new UserEntity(value));
    }
  }
  setUser(
    value:
      | MOne<UserEntity>
      | InstanceType<typeof UserEntity>
      | null
      | undefined,
  ) {
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
