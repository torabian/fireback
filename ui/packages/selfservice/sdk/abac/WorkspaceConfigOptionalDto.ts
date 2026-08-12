import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for workspaceConfigOptionalDto
 **/
export class WorkspaceConfigOptionalDto {
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
   * Enables the recaptcha2 for authentication flow.
   * @type {boolean}
   **/
  #enableRecaptcha2?: boolean | null = undefined;
  /**
   * Enables the recaptcha2 for authentication flow.
   * @returns {boolean}
   **/
  get enableRecaptcha2() {
    return this.#enableRecaptcha2;
  }
  /**
   * Enables the recaptcha2 for authentication flow.
   * @type {boolean}
   **/
  set enableRecaptcha2(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#enableRecaptcha2 = correctType ? value : Boolean(value);
  }
  setEnableRecaptcha2(value: boolean | null | undefined) {
    this.enableRecaptcha2 = value;
    return this;
  }
  /**
   * Enables the otp option. It's not forcing it, so user can choose if they want otp or password.
   * @type {boolean}
   **/
  #enableOtp?: boolean | null = undefined;
  /**
   * Enables the otp option. It's not forcing it, so user can choose if they want otp or password.
   * @returns {boolean}
   **/
  get enableOtp() {
    return this.#enableOtp;
  }
  /**
   * Enables the otp option. It's not forcing it, so user can choose if they want otp or password.
   * @type {boolean}
   **/
  set enableOtp(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#enableOtp = correctType ? value : Boolean(value);
  }
  setEnableOtp(value: boolean | null | undefined) {
    this.enableOtp = value;
    return this;
  }
  /**
   * Forces the user to have otp verification before can create an account. They can define their password still.
   * @type {boolean}
   **/
  #requireOtpOnSignup?: boolean | null = undefined;
  /**
   * Forces the user to have otp verification before can create an account. They can define their password still.
   * @returns {boolean}
   **/
  get requireOtpOnSignup() {
    return this.#requireOtpOnSignup;
  }
  /**
   * Forces the user to have otp verification before can create an account. They can define their password still.
   * @type {boolean}
   **/
  set requireOtpOnSignup(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#requireOtpOnSignup = correctType ? value : Boolean(value);
  }
  setRequireOtpOnSignup(value: boolean | null | undefined) {
    this.requireOtpOnSignup = value;
    return this;
  }
  /**
   * Forces the user to use otp when signing in. Even if they have password set, they won't use it and only will be able to signin using that otp.
   * @type {boolean}
   **/
  #requireOtpOnSignin?: boolean | null = undefined;
  /**
   * Forces the user to use otp when signing in. Even if they have password set, they won't use it and only will be able to signin using that otp.
   * @returns {boolean}
   **/
  get requireOtpOnSignin() {
    return this.#requireOtpOnSignin;
  }
  /**
   * Forces the user to use otp when signing in. Even if they have password set, they won't use it and only will be able to signin using that otp.
   * @type {boolean}
   **/
  set requireOtpOnSignin(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#requireOtpOnSignin = correctType ? value : Boolean(value);
  }
  setRequireOtpOnSignin(value: boolean | null | undefined) {
    this.requireOtpOnSignin = value;
    return this;
  }
  /**
   * Secret which would be used to decrypt if the recaptcha is correct. Should not be available publicly.
   * @type {string}
   **/
  #recaptcha2ServerKey?: string | null = undefined;
  /**
   * Secret which would be used to decrypt if the recaptcha is correct. Should not be available publicly.
   * @returns {string}
   **/
  get recaptcha2ServerKey() {
    return this.#recaptcha2ServerKey;
  }
  /**
   * Secret which would be used to decrypt if the recaptcha is correct. Should not be available publicly.
   * @type {string}
   **/
  set recaptcha2ServerKey(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#recaptcha2ServerKey = correctType ? value : String(value);
  }
  setRecaptcha2ServerKey(value: string | null | undefined) {
    this.recaptcha2ServerKey = value;
    return this;
  }
  /**
   * Secret which would be used for recaptcha2 on the client side. Can be publicly visible, and upon authenticating users it would be sent to front-end.
   * @type {string}
   **/
  #recaptcha2ClientKey?: string | null = undefined;
  /**
   * Secret which would be used for recaptcha2 on the client side. Can be publicly visible, and upon authenticating users it would be sent to front-end.
   * @returns {string}
   **/
  get recaptcha2ClientKey() {
    return this.#recaptcha2ClientKey;
  }
  /**
   * Secret which would be used for recaptcha2 on the client side. Can be publicly visible, and upon authenticating users it would be sent to front-end.
   * @type {string}
   **/
  set recaptcha2ClientKey(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#recaptcha2ClientKey = correctType ? value : String(value);
  }
  setRecaptcha2ClientKey(value: string | null | undefined) {
    this.recaptcha2ClientKey = value;
    return this;
  }
  /**
   * Enables user to make 2FA using apps such as google authenticator or microsoft authenticator.
   * @type {boolean}
   **/
  #enableTotp?: boolean | null = undefined;
  /**
   * Enables user to make 2FA using apps such as google authenticator or microsoft authenticator.
   * @returns {boolean}
   **/
  get enableTotp() {
    return this.#enableTotp;
  }
  /**
   * Enables user to make 2FA using apps such as google authenticator or microsoft authenticator.
   * @type {boolean}
   **/
  set enableTotp(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#enableTotp = correctType ? value : Boolean(value);
  }
  setEnableTotp(value: boolean | null | undefined) {
    this.enableTotp = value;
    return this;
  }
  /**
   * Forces the user to setup a 2FA in order to access their account. Users which did not setup this won't be affected.
   * @type {boolean}
   **/
  #forceTotp?: boolean | null = undefined;
  /**
   * Forces the user to setup a 2FA in order to access their account. Users which did not setup this won't be affected.
   * @returns {boolean}
   **/
  get forceTotp() {
    return this.#forceTotp;
  }
  /**
   * Forces the user to setup a 2FA in order to access their account. Users which did not setup this won't be affected.
   * @type {boolean}
   **/
  set forceTotp(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#forceTotp = correctType ? value : Boolean(value);
  }
  setForceTotp(value: boolean | null | undefined) {
    this.forceTotp = value;
    return this;
  }
  /**
   * Forces users who want to create account using phone number to also set a password on their account
   * @type {boolean}
   **/
  #forcePasswordOnPhone?: boolean | null = undefined;
  /**
   * Forces users who want to create account using phone number to also set a password on their account
   * @returns {boolean}
   **/
  get forcePasswordOnPhone() {
    return this.#forcePasswordOnPhone;
  }
  /**
   * Forces users who want to create account using phone number to also set a password on their account
   * @type {boolean}
   **/
  set forcePasswordOnPhone(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#forcePasswordOnPhone = correctType ? value : Boolean(value);
  }
  setForcePasswordOnPhone(value: boolean | null | undefined) {
    this.forcePasswordOnPhone = value;
    return this;
  }
  /**
   * Forces the creation of account using phone number to ask for user first name and last name
   * @type {boolean}
   **/
  #forcePersonNameOnPhone?: boolean | null = undefined;
  /**
   * Forces the creation of account using phone number to ask for user first name and last name
   * @returns {boolean}
   **/
  get forcePersonNameOnPhone() {
    return this.#forcePersonNameOnPhone;
  }
  /**
   * Forces the creation of account using phone number to ask for user first name and last name
   * @type {boolean}
   **/
  set forcePersonNameOnPhone(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#forcePersonNameOnPhone = correctType ? value : Boolean(value);
  }
  setForcePersonNameOnPhone(value: boolean | null | undefined) {
    this.forcePersonNameOnPhone = value;
    return this;
  }
  /**
   * The unique-id of the email provider service, which will be used to send the messages using it's service.
   * @type {string}
   **/
  #generalEmailProviderId?: string | null = undefined;
  /**
   * The unique-id of the email provider service, which will be used to send the messages using it's service.
   * @returns {string}
   **/
  get generalEmailProviderId() {
    return this.#generalEmailProviderId;
  }
  /**
   * The unique-id of the email provider service, which will be used to send the messages using it's service.
   * @type {string}
   **/
  set generalEmailProviderId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#generalEmailProviderId = correctType ? value : String(value);
  }
  setGeneralEmailProviderId(value: string | null | undefined) {
    this.generalEmailProviderId = value;
    return this;
  }
  /**
   * The unique-id of the general service which would be used to send text messages (sms).
   * @type {string}
   **/
  #generalGsmProviderId?: string | null = undefined;
  /**
   * The unique-id of the general service which would be used to send text messages (sms).
   * @returns {string}
   **/
  get generalGsmProviderId() {
    return this.#generalGsmProviderId;
  }
  /**
   * The unique-id of the general service which would be used to send text messages (sms).
   * @type {string}
   **/
  set generalGsmProviderId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#generalGsmProviderId = correctType ? value : String(value);
  }
  setGeneralGsmProviderId(value: string | null | undefined) {
    this.generalGsmProviderId = value;
    return this;
  }
  /**
   * The unique-id of the template used as default when a user is inviting a third-party into their own workspace.
   * @type {string}
   **/
  #inviteToWorkspaceContentId?: string | null = undefined;
  /**
   * The unique-id of the template used as default when a user is inviting a third-party into their own workspace.
   * @returns {string}
   **/
  get inviteToWorkspaceContentId() {
    return this.#inviteToWorkspaceContentId;
  }
  /**
   * The unique-id of the template used as default when a user is inviting a third-party into their own workspace.
   * @type {string}
   **/
  set inviteToWorkspaceContentId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#inviteToWorkspaceContentId = correctType ? value : String(value);
  }
  setInviteToWorkspaceContentId(value: string | null | undefined) {
    this.inviteToWorkspaceContentId = value;
    return this;
  }
  /**
   * The unique-id of the template used to fill the message for email one-time-password requests.
   * @type {string}
   **/
  #emailOtpContentId?: string | null = undefined;
  /**
   * The unique-id of the template used to fill the message for email one-time-password requests.
   * @returns {string}
   **/
  get emailOtpContentId() {
    return this.#emailOtpContentId;
  }
  /**
   * The unique-id of the template used to fill the message for email one-time-password requests.
   * @type {string}
   **/
  set emailOtpContentId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#emailOtpContentId = correctType ? value : String(value);
  }
  setEmailOtpContentId(value: string | null | undefined) {
    this.emailOtpContentId = value;
    return this;
  }
  /**
   * The unique-id of the template used for OTP text messages, including the one time password code.
   * @type {string}
   **/
  #smsOtpContentId?: string | null = undefined;
  /**
   * The unique-id of the template used for OTP text messages, including the one time password code.
   * @returns {string}
   **/
  get smsOtpContentId() {
    return this.#smsOtpContentId;
  }
  /**
   * The unique-id of the template used for OTP text messages, including the one time password code.
   * @type {string}
   **/
  set smsOtpContentId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#smsOtpContentId = correctType ? value : String(value);
  }
  setSmsOtpContentId(value: string | null | undefined) {
    this.smsOtpContentId = value;
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
    const d = data as Partial<WorkspaceConfigOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.enableRecaptcha2 !== undefined) {
      this.enableRecaptcha2 = d.enableRecaptcha2;
    }
    if (d.enableOtp !== undefined) {
      this.enableOtp = d.enableOtp;
    }
    if (d.requireOtpOnSignup !== undefined) {
      this.requireOtpOnSignup = d.requireOtpOnSignup;
    }
    if (d.requireOtpOnSignin !== undefined) {
      this.requireOtpOnSignin = d.requireOtpOnSignin;
    }
    if (d.recaptcha2ServerKey !== undefined) {
      this.recaptcha2ServerKey = d.recaptcha2ServerKey;
    }
    if (d.recaptcha2ClientKey !== undefined) {
      this.recaptcha2ClientKey = d.recaptcha2ClientKey;
    }
    if (d.enableTotp !== undefined) {
      this.enableTotp = d.enableTotp;
    }
    if (d.forceTotp !== undefined) {
      this.forceTotp = d.forceTotp;
    }
    if (d.forcePasswordOnPhone !== undefined) {
      this.forcePasswordOnPhone = d.forcePasswordOnPhone;
    }
    if (d.forcePersonNameOnPhone !== undefined) {
      this.forcePersonNameOnPhone = d.forcePersonNameOnPhone;
    }
    if (d.generalEmailProviderId !== undefined) {
      this.generalEmailProviderId = d.generalEmailProviderId;
    }
    if (d.generalGsmProviderId !== undefined) {
      this.generalGsmProviderId = d.generalGsmProviderId;
    }
    if (d.inviteToWorkspaceContentId !== undefined) {
      this.inviteToWorkspaceContentId = d.inviteToWorkspaceContentId;
    }
    if (d.emailOtpContentId !== undefined) {
      this.emailOtpContentId = d.emailOtpContentId;
    }
    if (d.smsOtpContentId !== undefined) {
      this.smsOtpContentId = d.smsOtpContentId;
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
      enableRecaptcha2: this.#enableRecaptcha2,
      enableOtp: this.#enableOtp,
      requireOtpOnSignup: this.#requireOtpOnSignup,
      requireOtpOnSignin: this.#requireOtpOnSignin,
      recaptcha2ServerKey: this.#recaptcha2ServerKey,
      recaptcha2ClientKey: this.#recaptcha2ClientKey,
      enableTotp: this.#enableTotp,
      forceTotp: this.#forceTotp,
      forcePasswordOnPhone: this.#forcePasswordOnPhone,
      forcePersonNameOnPhone: this.#forcePersonNameOnPhone,
      generalEmailProviderId: this.#generalEmailProviderId,
      generalGsmProviderId: this.#generalGsmProviderId,
      inviteToWorkspaceContentId: this.#inviteToWorkspaceContentId,
      emailOtpContentId: this.#emailOtpContentId,
      smsOtpContentId: this.#smsOtpContentId,
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
      enableRecaptcha2: "enableRecaptcha2",
      enableOtp: "enableOtp",
      requireOtpOnSignup: "requireOtpOnSignup",
      requireOtpOnSignin: "requireOtpOnSignin",
      recaptcha2ServerKey: "recaptcha2ServerKey",
      recaptcha2ClientKey: "recaptcha2ClientKey",
      enableTotp: "enableTotp",
      forceTotp: "forceTotp",
      forcePasswordOnPhone: "forcePasswordOnPhone",
      forcePersonNameOnPhone: "forcePersonNameOnPhone",
      generalEmailProviderId: "generalEmailProviderId",
      generalGsmProviderId: "generalGsmProviderId",
      inviteToWorkspaceContentId: "inviteToWorkspaceContentId",
      emailOtpContentId: "emailOtpContentId",
      smsOtpContentId: "smsOtpContentId",
      workspaceId: "workspaceId",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of WorkspaceConfigOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: WorkspaceConfigOptionalDtoType) {
    return new WorkspaceConfigOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of WorkspaceConfigOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<WorkspaceConfigOptionalDtoType>) {
    return new WorkspaceConfigOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<WorkspaceConfigOptionalDtoType>,
  ): InstanceType<typeof WorkspaceConfigOptionalDto> {
    return new WorkspaceConfigOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof WorkspaceConfigOptionalDto> {
    return new WorkspaceConfigOptionalDto(this.toJSON());
  }
}
export abstract class WorkspaceConfigOptionalDtoFactory {
  abstract create(data: unknown): WorkspaceConfigOptionalDto;
}
/**
 * The base type definition for workspaceConfigOptionalDto
 **/
export type WorkspaceConfigOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   * Enables the recaptcha2 for authentication flow.
   * @type {boolean}
   **/
  enableRecaptcha2?: boolean;
  /**
   * Enables the otp option. It's not forcing it, so user can choose if they want otp or password.
   * @type {boolean}
   **/
  enableOtp?: boolean;
  /**
   * Forces the user to have otp verification before can create an account. They can define their password still.
   * @type {boolean}
   **/
  requireOtpOnSignup?: boolean;
  /**
   * Forces the user to use otp when signing in. Even if they have password set, they won't use it and only will be able to signin using that otp.
   * @type {boolean}
   **/
  requireOtpOnSignin?: boolean;
  /**
   * Secret which would be used to decrypt if the recaptcha is correct. Should not be available publicly.
   * @type {string}
   **/
  recaptcha2ServerKey?: string;
  /**
   * Secret which would be used for recaptcha2 on the client side. Can be publicly visible, and upon authenticating users it would be sent to front-end.
   * @type {string}
   **/
  recaptcha2ClientKey?: string;
  /**
   * Enables user to make 2FA using apps such as google authenticator or microsoft authenticator.
   * @type {boolean}
   **/
  enableTotp?: boolean;
  /**
   * Forces the user to setup a 2FA in order to access their account. Users which did not setup this won't be affected.
   * @type {boolean}
   **/
  forceTotp?: boolean;
  /**
   * Forces users who want to create account using phone number to also set a password on their account
   * @type {boolean}
   **/
  forcePasswordOnPhone?: boolean;
  /**
   * Forces the creation of account using phone number to ask for user first name and last name
   * @type {boolean}
   **/
  forcePersonNameOnPhone?: boolean;
  /**
   * The unique-id of the email provider service, which will be used to send the messages using it's service.
   * @type {string}
   **/
  generalEmailProviderId?: string;
  /**
   * The unique-id of the general service which would be used to send text messages (sms).
   * @type {string}
   **/
  generalGsmProviderId?: string;
  /**
   * The unique-id of the template used as default when a user is inviting a third-party into their own workspace.
   * @type {string}
   **/
  inviteToWorkspaceContentId?: string;
  /**
   * The unique-id of the template used to fill the message for email one-time-password requests.
   * @type {string}
   **/
  emailOtpContentId?: string;
  /**
   * The unique-id of the template used for OTP text messages, including the one time password code.
   * @type {string}
   **/
  smsOtpContentId?: string;
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
export namespace WorkspaceConfigOptionalDtoType {}
