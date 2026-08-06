import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for appMenuDto
 **/
export class AppMenuDto {
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
   * Label that will be visible to user
   * @type {string}
   **/
  #label: string = "";
  /**
   * Label that will be visible to user
   * @returns {string}
   **/
  get label() {
    return this.#label;
  }
  /**
   * Label that will be visible to user
   * @type {string}
   **/
  set label(value: string) {
    this.#label = String(value);
  }
  setLabel(value: string) {
    this.label = value;
    return this;
  }
  /**
   * Location that will be navigated in case of click or selection on ui
   * @type {string}
   **/
  #href: string = "";
  /**
   * Location that will be navigated in case of click or selection on ui
   * @returns {string}
   **/
  get href() {
    return this.#href;
  }
  /**
   * Location that will be navigated in case of click or selection on ui
   * @type {string}
   **/
  set href(value: string) {
    this.#href = String(value);
  }
  setHref(value: string) {
    this.href = value;
    return this;
  }
  /**
   * Icon string address which matches the resources on the front-end apps.
   * @type {string}
   **/
  #icon: string = "";
  /**
   * Icon string address which matches the resources on the front-end apps.
   * @returns {string}
   **/
  get icon() {
    return this.#icon;
  }
  /**
   * Icon string address which matches the resources on the front-end apps.
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
   * Custom window location url matchers, for inner screens.
   * @type {string}
   **/
  #activeMatcher: string = "";
  /**
   * Custom window location url matchers, for inner screens.
   * @returns {string}
   **/
  get activeMatcher() {
    return this.#activeMatcher;
  }
  /**
   * Custom window location url matchers, for inner screens.
   * @type {string}
   **/
  set activeMatcher(value: string) {
    this.#activeMatcher = String(value);
  }
  setActiveMatcher(value: string) {
    this.activeMatcher = value;
    return this;
  }
  /**
   * The unique-id of the capability which is required for the menu to be visible.
   * @type {string}
   **/
  #capabilityId?: string | null = undefined;
  /**
   * The unique-id of the capability which is required for the menu to be visible.
   * @returns {string}
   **/
  get capabilityId() {
    return this.#capabilityId;
  }
  /**
   * The unique-id of the capability which is required for the menu to be visible.
   * @type {string}
   **/
  set capabilityId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#capabilityId = correctType ? value : String(value);
  }
  setCapabilityId(value: string | null | undefined) {
    this.capabilityId = value;
    return this;
  }
  /**
   * The unique-id of the parent menu item, for nested/tree menus.
   * @type {string}
   **/
  #parentId?: string | null = undefined;
  /**
   * The unique-id of the parent menu item, for nested/tree menus.
   * @returns {string}
   **/
  get parentId() {
    return this.#parentId;
  }
  /**
   * The unique-id of the parent menu item, for nested/tree menus.
   * @type {string}
   **/
  set parentId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#parentId = correctType ? value : String(value);
  }
  setParentId(value: string | null | undefined) {
    this.parentId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #workspaceId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get workspaceId() {
    return this.#workspaceId;
  }
  /**
   *
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
    const d = data as Partial<AppMenuDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.label !== undefined) {
      this.label = d.label;
    }
    if (d.href !== undefined) {
      this.href = d.href;
    }
    if (d.icon !== undefined) {
      this.icon = d.icon;
    }
    if (d.activeMatcher !== undefined) {
      this.activeMatcher = d.activeMatcher;
    }
    if (d.capabilityId !== undefined) {
      this.capabilityId = d.capabilityId;
    }
    if (d.parentId !== undefined) {
      this.parentId = d.parentId;
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
      label: this.#label,
      href: this.#href,
      icon: this.#icon,
      activeMatcher: this.#activeMatcher,
      capabilityId: this.#capabilityId,
      parentId: this.#parentId,
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
      label: "label",
      href: "href",
      icon: "icon",
      activeMatcher: "activeMatcher",
      capabilityId: "capabilityId",
      parentId: "parentId",
      workspaceId: "workspaceId",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of AppMenuDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: AppMenuDtoType) {
    return new AppMenuDto(possibleDtoObject);
  }
  /**
   * Creates an instance of AppMenuDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<AppMenuDtoType>) {
    return new AppMenuDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<AppMenuDtoType>,
  ): InstanceType<typeof AppMenuDto> {
    return new AppMenuDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof AppMenuDto> {
    return new AppMenuDto(this.toJSON());
  }
}
export abstract class AppMenuDtoFactory {
  abstract create(data: unknown): AppMenuDto;
}
/**
 * The base type definition for appMenuDto
 **/
export type AppMenuDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * Label that will be visible to user
   * @type {string}
   **/
  label: string;
  /**
   * Location that will be navigated in case of click or selection on ui
   * @type {string}
   **/
  href: string;
  /**
   * Icon string address which matches the resources on the front-end apps.
   * @type {string}
   **/
  icon: string;
  /**
   * Custom window location url matchers, for inner screens.
   * @type {string}
   **/
  activeMatcher: string;
  /**
   * The unique-id of the capability which is required for the menu to be visible.
   * @type {string}
   **/
  capabilityId?: string;
  /**
   * The unique-id of the parent menu item, for nested/tree menus.
   * @type {string}
   **/
  parentId?: string;
  /**
   *
   * @type {string}
   **/
  workspaceId?: string;
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
export namespace AppMenuDtoType {}
