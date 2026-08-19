import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for notificationOptionalDto
 **/
export class NotificationOptionalDto {
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
   * UniqueId of the user this notification was sent to.
   * @type {string}
   **/
  #userId?: string | null = undefined;
  /**
   * UniqueId of the user this notification was sent to.
   * @returns {string}
   **/
  get userId() {
    return this.#userId;
  }
  /**
   * UniqueId of the user this notification was sent to.
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
   * UniqueId of the (root) user who sent this notification.
   * @type {string}
   **/
  #senderId?: string | null = undefined;
  /**
   * UniqueId of the (root) user who sent this notification.
   * @returns {string}
   **/
  get senderId() {
    return this.#senderId;
  }
  /**
   * UniqueId of the (root) user who sent this notification.
   * @type {string}
   **/
  set senderId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#senderId = correctType ? value : String(value);
  }
  setSenderId(value: string | null | undefined) {
    this.senderId = value;
    return this;
  }
  /**
   * Short notification title.
   * @type {string}
   **/
  #title?: string | null = undefined;
  /**
   * Short notification title.
   * @returns {string}
   **/
  get title() {
    return this.#title;
  }
  /**
   * Short notification title.
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
   * Notification message body.
   * @type {string}
   **/
  #body?: string | null = undefined;
  /**
   * Notification message body.
   * @returns {string}
   **/
  get body() {
    return this.#body;
  }
  /**
   * Notification message body.
   * @type {string}
   **/
  set body(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#body = correctType ? value : String(value);
  }
  setBody(value: string | null | undefined) {
    this.body = value;
    return this;
  }
  /**
   * Whether the recipient has read this notification yet.
   * @type {boolean}
   **/
  #isRead?: boolean | null = undefined;
  /**
   * Whether the recipient has read this notification yet.
   * @returns {boolean}
   **/
  get isRead() {
    return this.#isRead;
  }
  /**
   * Whether the recipient has read this notification yet.
   * @type {boolean}
   **/
  set isRead(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#isRead = correctType ? value : Boolean(value);
  }
  setIsRead(value: boolean | null | undefined) {
    this.isRead = value;
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
    const d = data as Partial<NotificationOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.userId !== undefined) {
      this.userId = d.userId;
    }
    if (d.senderId !== undefined) {
      this.senderId = d.senderId;
    }
    if (d.title !== undefined) {
      this.title = d.title;
    }
    if (d.body !== undefined) {
      this.body = d.body;
    }
    if (d.isRead !== undefined) {
      this.isRead = d.isRead;
    }
    if (d.workspaceId !== undefined) {
      this.workspaceId = d.workspaceId;
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
      userId: this.#userId,
      senderId: this.#senderId,
      title: this.#title,
      body: this.#body,
      isRead: this.#isRead,
      workspaceId: this.#workspaceId,
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
      userId: "userId",
      senderId: "senderId",
      title: "title",
      body: "body",
      isRead: "isRead",
      workspaceId: "workspaceId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of NotificationOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: NotificationOptionalDtoType) {
    return new NotificationOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of NotificationOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<NotificationOptionalDtoType>) {
    return new NotificationOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<NotificationOptionalDtoType>,
  ): InstanceType<typeof NotificationOptionalDto> {
    return new NotificationOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof NotificationOptionalDto> {
    return new NotificationOptionalDto(this.toJSON());
  }
}
export abstract class NotificationOptionalDtoFactory {
  abstract create(data: unknown): NotificationOptionalDto;
}
/**
 * The base type definition for notificationOptionalDto
 **/
export type NotificationOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * UniqueId of the user this notification was sent to.
   * @type {string}
   **/
  userId?: string;
  /**
   * UniqueId of the (root) user who sent this notification.
   * @type {string}
   **/
  senderId?: string;
  /**
   * Short notification title.
   * @type {string}
   **/
  title?: string;
  /**
   * Notification message body.
   * @type {string}
   **/
  body?: string;
  /**
   * Whether the recipient has read this notification yet.
   * @type {boolean}
   **/
  isRead?: boolean;
  /**
   * The unique-id of the workspace which content belongs to.
   * @type {string}
   **/
  workspaceId?: string;
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
export namespace NotificationOptionalDtoType {}
