// @ts-nocheck 
 // This no check has been added via fireback. 
import { type PartialDeep } from "../../sdk/common/fetchx";
/**
 * The base class definition for testMailDto
 **/
export class TestMailDto {
  /**
   *
   * @type {string}
   **/
  #senderId: string = "";
  /**
   *
   * @returns {string}
   **/
  get senderId() {
    return this.#senderId;
  }
  /**
   *
   * @type {string}
   **/
  set senderId(value: string) {
    this.#senderId = String(value);
  }
  setSenderId(value: string) {
    this.senderId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #toName: string = "";
  /**
   *
   * @returns {string}
   **/
  get toName() {
    return this.#toName;
  }
  /**
   *
   * @type {string}
   **/
  set toName(value: string) {
    this.#toName = String(value);
  }
  setToName(value: string) {
    this.toName = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #toEmail: string = "";
  /**
   *
   * @returns {string}
   **/
  get toEmail() {
    return this.#toEmail;
  }
  /**
   *
   * @type {string}
   **/
  set toEmail(value: string) {
    this.#toEmail = String(value);
  }
  setToEmail(value: string) {
    this.toEmail = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #subject: string = "";
  /**
   *
   * @returns {string}
   **/
  get subject() {
    return this.#subject;
  }
  /**
   *
   * @type {string}
   **/
  set subject(value: string) {
    this.#subject = String(value);
  }
  setSubject(value: string) {
    this.subject = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #content: string = "";
  /**
   *
   * @returns {string}
   **/
  get content() {
    return this.#content;
  }
  /**
   *
   * @type {string}
   **/
  set content(value: string) {
    this.#content = String(value);
  }
  setContent(value: string) {
    this.content = value;
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
    const d = data as Partial<TestMailDto>;
    if (d.senderId !== undefined) {
      this.senderId = d.senderId;
    }
    if (d.toName !== undefined) {
      this.toName = d.toName;
    }
    if (d.toEmail !== undefined) {
      this.toEmail = d.toEmail;
    }
    if (d.subject !== undefined) {
      this.subject = d.subject;
    }
    if (d.content !== undefined) {
      this.content = d.content;
    }
  }
  /**
   *	Special toJSON override, since the field are private,
   *	Json stringify won't see them unless we mention it explicitly.
   **/
  toJSON() {
    return {
      senderId: this.#senderId,
      toName: this.#toName,
      toEmail: this.#toEmail,
      subject: this.#subject,
      content: this.#content,
    };
  }
  toString() {
    return JSON.stringify(this);
  }
  static get Fields() {
    return {
      senderId: "senderId",
      toName: "toName",
      toEmail: "toEmail",
      subject: "subject",
      content: "content",
    };
  }
  /**
   * Creates an instance of TestMailDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: TestMailDtoType) {
    return new TestMailDto(possibleDtoObject);
  }
  /**
   * Creates an instance of TestMailDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<TestMailDtoType>) {
    return new TestMailDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<TestMailDtoType>,
  ): InstanceType<typeof TestMailDto> {
    return new TestMailDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof TestMailDto> {
    return new TestMailDto(this.toJSON());
  }
}
export abstract class TestMailDtoFactory {
  abstract create(data: unknown): TestMailDto;
}
/**
 * The base type definition for testMailDto
 **/
export type TestMailDtoType = {
  /**
   *
   * @type {string}
   **/
  senderId: string;
  /**
   *
   * @type {string}
   **/
  toName: string;
  /**
   *
   * @type {string}
   **/
  toEmail: string;
  /**
   *
   * @type {string}
   **/
  subject: string;
  /**
   *
   * @type {string}
   **/
  content: string;
};
// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace TestMailDtoType {}
