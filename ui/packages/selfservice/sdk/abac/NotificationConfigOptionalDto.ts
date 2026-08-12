import { type PartialDeep } from "@fireback/js-remote-ctx/common/fetchx";
/**
 * The base class definition for notificationConfigOptionalDto
 **/
export class NotificationConfigOptionalDto {
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
   * @type {boolean}
   **/
  #cascadeToSubWorkspaces?: boolean | null = undefined;
  /**
   *
   * @returns {boolean}
   **/
  get cascadeToSubWorkspaces() {
    return this.#cascadeToSubWorkspaces;
  }
  /**
   *
   * @type {boolean}
   **/
  set cascadeToSubWorkspaces(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#cascadeToSubWorkspaces = correctType ? value : Boolean(value);
  }
  setCascadeToSubWorkspaces(value: boolean | null | undefined) {
    this.cascadeToSubWorkspaces = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #forcedCascadeEmailProvider?: boolean | null = undefined;
  /**
   *
   * @returns {boolean}
   **/
  get forcedCascadeEmailProvider() {
    return this.#forcedCascadeEmailProvider;
  }
  /**
   *
   * @type {boolean}
   **/
  set forcedCascadeEmailProvider(value: boolean | null | undefined) {
    const correctType =
      value === true ||
      value === false ||
      value === undefined ||
      value === null;
    this.#forcedCascadeEmailProvider = correctType ? value : Boolean(value);
  }
  setForcedCascadeEmailProvider(value: boolean | null | undefined) {
    this.forcedCascadeEmailProvider = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #generalEmailProviderId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get generalEmailProviderId() {
    return this.#generalEmailProviderId;
  }
  /**
   *
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
   *
   * @type {string}
   **/
  #generalGsmProviderId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get generalGsmProviderId() {
    return this.#generalGsmProviderId;
  }
  /**
   *
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
   *
   * @type {string}
   **/
  #inviteToWorkspaceContent?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get inviteToWorkspaceContent() {
    return this.#inviteToWorkspaceContent;
  }
  /**
   *
   * @type {string}
   **/
  set inviteToWorkspaceContent(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#inviteToWorkspaceContent = correctType ? value : String(value);
  }
  setInviteToWorkspaceContent(value: string | null | undefined) {
    this.inviteToWorkspaceContent = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteToWorkspaceContentExcerpt?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get inviteToWorkspaceContentExcerpt() {
    return this.#inviteToWorkspaceContentExcerpt;
  }
  /**
   *
   * @type {string}
   **/
  set inviteToWorkspaceContentExcerpt(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#inviteToWorkspaceContentExcerpt = correctType ? value : String(value);
  }
  setInviteToWorkspaceContentExcerpt(value: string | null | undefined) {
    this.inviteToWorkspaceContentExcerpt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteToWorkspaceContentDefault?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get inviteToWorkspaceContentDefault() {
    return this.#inviteToWorkspaceContentDefault;
  }
  /**
   *
   * @type {string}
   **/
  set inviteToWorkspaceContentDefault(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#inviteToWorkspaceContentDefault = correctType ? value : String(value);
  }
  setInviteToWorkspaceContentDefault(value: string | null | undefined) {
    this.inviteToWorkspaceContentDefault = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteToWorkspaceContentDefaultExcerpt?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get inviteToWorkspaceContentDefaultExcerpt() {
    return this.#inviteToWorkspaceContentDefaultExcerpt;
  }
  /**
   *
   * @type {string}
   **/
  set inviteToWorkspaceContentDefaultExcerpt(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#inviteToWorkspaceContentDefaultExcerpt = correctType
      ? value
      : String(value);
  }
  setInviteToWorkspaceContentDefaultExcerpt(value: string | null | undefined) {
    this.inviteToWorkspaceContentDefaultExcerpt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteToWorkspaceTitle?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get inviteToWorkspaceTitle() {
    return this.#inviteToWorkspaceTitle;
  }
  /**
   *
   * @type {string}
   **/
  set inviteToWorkspaceTitle(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#inviteToWorkspaceTitle = correctType ? value : String(value);
  }
  setInviteToWorkspaceTitle(value: string | null | undefined) {
    this.inviteToWorkspaceTitle = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteToWorkspaceTitleDefault?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get inviteToWorkspaceTitleDefault() {
    return this.#inviteToWorkspaceTitleDefault;
  }
  /**
   *
   * @type {string}
   **/
  set inviteToWorkspaceTitleDefault(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#inviteToWorkspaceTitleDefault = correctType ? value : String(value);
  }
  setInviteToWorkspaceTitleDefault(value: string | null | undefined) {
    this.inviteToWorkspaceTitleDefault = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteToWorkspaceSenderId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get inviteToWorkspaceSenderId() {
    return this.#inviteToWorkspaceSenderId;
  }
  /**
   *
   * @type {string}
   **/
  set inviteToWorkspaceSenderId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#inviteToWorkspaceSenderId = correctType ? value : String(value);
  }
  setInviteToWorkspaceSenderId(value: string | null | undefined) {
    this.inviteToWorkspaceSenderId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #accountCenterEmailSenderId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get accountCenterEmailSenderId() {
    return this.#accountCenterEmailSenderId;
  }
  /**
   *
   * @type {string}
   **/
  set accountCenterEmailSenderId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#accountCenterEmailSenderId = correctType ? value : String(value);
  }
  setAccountCenterEmailSenderId(value: string | null | undefined) {
    this.accountCenterEmailSenderId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #forgetPasswordContent?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get forgetPasswordContent() {
    return this.#forgetPasswordContent;
  }
  /**
   *
   * @type {string}
   **/
  set forgetPasswordContent(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#forgetPasswordContent = correctType ? value : String(value);
  }
  setForgetPasswordContent(value: string | null | undefined) {
    this.forgetPasswordContent = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #forgetPasswordContentExcerpt?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get forgetPasswordContentExcerpt() {
    return this.#forgetPasswordContentExcerpt;
  }
  /**
   *
   * @type {string}
   **/
  set forgetPasswordContentExcerpt(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#forgetPasswordContentExcerpt = correctType ? value : String(value);
  }
  setForgetPasswordContentExcerpt(value: string | null | undefined) {
    this.forgetPasswordContentExcerpt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #forgetPasswordContentDefault?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get forgetPasswordContentDefault() {
    return this.#forgetPasswordContentDefault;
  }
  /**
   *
   * @type {string}
   **/
  set forgetPasswordContentDefault(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#forgetPasswordContentDefault = correctType ? value : String(value);
  }
  setForgetPasswordContentDefault(value: string | null | undefined) {
    this.forgetPasswordContentDefault = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #forgetPasswordContentDefaultExcerpt?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get forgetPasswordContentDefaultExcerpt() {
    return this.#forgetPasswordContentDefaultExcerpt;
  }
  /**
   *
   * @type {string}
   **/
  set forgetPasswordContentDefaultExcerpt(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#forgetPasswordContentDefaultExcerpt = correctType
      ? value
      : String(value);
  }
  setForgetPasswordContentDefaultExcerpt(value: string | null | undefined) {
    this.forgetPasswordContentDefaultExcerpt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #forgetPasswordTitle?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get forgetPasswordTitle() {
    return this.#forgetPasswordTitle;
  }
  /**
   *
   * @type {string}
   **/
  set forgetPasswordTitle(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#forgetPasswordTitle = correctType ? value : String(value);
  }
  setForgetPasswordTitle(value: string | null | undefined) {
    this.forgetPasswordTitle = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #forgetPasswordTitleDefault?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get forgetPasswordTitleDefault() {
    return this.#forgetPasswordTitleDefault;
  }
  /**
   *
   * @type {string}
   **/
  set forgetPasswordTitleDefault(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#forgetPasswordTitleDefault = correctType ? value : String(value);
  }
  setForgetPasswordTitleDefault(value: string | null | undefined) {
    this.forgetPasswordTitleDefault = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #forgetPasswordSenderId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get forgetPasswordSenderId() {
    return this.#forgetPasswordSenderId;
  }
  /**
   *
   * @type {string}
   **/
  set forgetPasswordSenderId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#forgetPasswordSenderId = correctType ? value : String(value);
  }
  setForgetPasswordSenderId(value: string | null | undefined) {
    this.forgetPasswordSenderId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #acceptLanguage?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get acceptLanguage() {
    return this.#acceptLanguage;
  }
  /**
   *
   * @type {string}
   **/
  set acceptLanguage(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#acceptLanguage = correctType ? value : String(value);
  }
  setAcceptLanguage(value: string | null | undefined) {
    this.acceptLanguage = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #confirmEmailSenderId?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get confirmEmailSenderId() {
    return this.#confirmEmailSenderId;
  }
  /**
   *
   * @type {string}
   **/
  set confirmEmailSenderId(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#confirmEmailSenderId = correctType ? value : String(value);
  }
  setConfirmEmailSenderId(value: string | null | undefined) {
    this.confirmEmailSenderId = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #confirmEmailContent?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get confirmEmailContent() {
    return this.#confirmEmailContent;
  }
  /**
   *
   * @type {string}
   **/
  set confirmEmailContent(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#confirmEmailContent = correctType ? value : String(value);
  }
  setConfirmEmailContent(value: string | null | undefined) {
    this.confirmEmailContent = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #confirmEmailContentExcerpt?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get confirmEmailContentExcerpt() {
    return this.#confirmEmailContentExcerpt;
  }
  /**
   *
   * @type {string}
   **/
  set confirmEmailContentExcerpt(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#confirmEmailContentExcerpt = correctType ? value : String(value);
  }
  setConfirmEmailContentExcerpt(value: string | null | undefined) {
    this.confirmEmailContentExcerpt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #confirmEmailContentDefault?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get confirmEmailContentDefault() {
    return this.#confirmEmailContentDefault;
  }
  /**
   *
   * @type {string}
   **/
  set confirmEmailContentDefault(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#confirmEmailContentDefault = correctType ? value : String(value);
  }
  setConfirmEmailContentDefault(value: string | null | undefined) {
    this.confirmEmailContentDefault = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #confirmEmailContentDefaultExcerpt?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get confirmEmailContentDefaultExcerpt() {
    return this.#confirmEmailContentDefaultExcerpt;
  }
  /**
   *
   * @type {string}
   **/
  set confirmEmailContentDefaultExcerpt(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#confirmEmailContentDefaultExcerpt = correctType
      ? value
      : String(value);
  }
  setConfirmEmailContentDefaultExcerpt(value: string | null | undefined) {
    this.confirmEmailContentDefaultExcerpt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #confirmEmailTitle?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get confirmEmailTitle() {
    return this.#confirmEmailTitle;
  }
  /**
   *
   * @type {string}
   **/
  set confirmEmailTitle(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#confirmEmailTitle = correctType ? value : String(value);
  }
  setConfirmEmailTitle(value: string | null | undefined) {
    this.confirmEmailTitle = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #confirmEmailTitleDefault?: string | null = undefined;
  /**
   *
   * @returns {string}
   **/
  get confirmEmailTitleDefault() {
    return this.#confirmEmailTitleDefault;
  }
  /**
   *
   * @type {string}
   **/
  set confirmEmailTitleDefault(value: string | null | undefined) {
    const correctType =
      typeof value === "string" || value === undefined || value === null;
    this.#confirmEmailTitleDefault = correctType ? value : String(value);
  }
  setConfirmEmailTitleDefault(value: string | null | undefined) {
    this.confirmEmailTitleDefault = value;
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
    const d = data as Partial<NotificationConfigOptionalDto>;
    if (d.uniqueId !== undefined) {
      this.uniqueId = d.uniqueId;
    }
    if (d.cascadeToSubWorkspaces !== undefined) {
      this.cascadeToSubWorkspaces = d.cascadeToSubWorkspaces;
    }
    if (d.forcedCascadeEmailProvider !== undefined) {
      this.forcedCascadeEmailProvider = d.forcedCascadeEmailProvider;
    }
    if (d.generalEmailProviderId !== undefined) {
      this.generalEmailProviderId = d.generalEmailProviderId;
    }
    if (d.generalGsmProviderId !== undefined) {
      this.generalGsmProviderId = d.generalGsmProviderId;
    }
    if (d.inviteToWorkspaceContent !== undefined) {
      this.inviteToWorkspaceContent = d.inviteToWorkspaceContent;
    }
    if (d.inviteToWorkspaceContentExcerpt !== undefined) {
      this.inviteToWorkspaceContentExcerpt = d.inviteToWorkspaceContentExcerpt;
    }
    if (d.inviteToWorkspaceContentDefault !== undefined) {
      this.inviteToWorkspaceContentDefault = d.inviteToWorkspaceContentDefault;
    }
    if (d.inviteToWorkspaceContentDefaultExcerpt !== undefined) {
      this.inviteToWorkspaceContentDefaultExcerpt =
        d.inviteToWorkspaceContentDefaultExcerpt;
    }
    if (d.inviteToWorkspaceTitle !== undefined) {
      this.inviteToWorkspaceTitle = d.inviteToWorkspaceTitle;
    }
    if (d.inviteToWorkspaceTitleDefault !== undefined) {
      this.inviteToWorkspaceTitleDefault = d.inviteToWorkspaceTitleDefault;
    }
    if (d.inviteToWorkspaceSenderId !== undefined) {
      this.inviteToWorkspaceSenderId = d.inviteToWorkspaceSenderId;
    }
    if (d.accountCenterEmailSenderId !== undefined) {
      this.accountCenterEmailSenderId = d.accountCenterEmailSenderId;
    }
    if (d.forgetPasswordContent !== undefined) {
      this.forgetPasswordContent = d.forgetPasswordContent;
    }
    if (d.forgetPasswordContentExcerpt !== undefined) {
      this.forgetPasswordContentExcerpt = d.forgetPasswordContentExcerpt;
    }
    if (d.forgetPasswordContentDefault !== undefined) {
      this.forgetPasswordContentDefault = d.forgetPasswordContentDefault;
    }
    if (d.forgetPasswordContentDefaultExcerpt !== undefined) {
      this.forgetPasswordContentDefaultExcerpt =
        d.forgetPasswordContentDefaultExcerpt;
    }
    if (d.forgetPasswordTitle !== undefined) {
      this.forgetPasswordTitle = d.forgetPasswordTitle;
    }
    if (d.forgetPasswordTitleDefault !== undefined) {
      this.forgetPasswordTitleDefault = d.forgetPasswordTitleDefault;
    }
    if (d.forgetPasswordSenderId !== undefined) {
      this.forgetPasswordSenderId = d.forgetPasswordSenderId;
    }
    if (d.acceptLanguage !== undefined) {
      this.acceptLanguage = d.acceptLanguage;
    }
    if (d.confirmEmailSenderId !== undefined) {
      this.confirmEmailSenderId = d.confirmEmailSenderId;
    }
    if (d.confirmEmailContent !== undefined) {
      this.confirmEmailContent = d.confirmEmailContent;
    }
    if (d.confirmEmailContentExcerpt !== undefined) {
      this.confirmEmailContentExcerpt = d.confirmEmailContentExcerpt;
    }
    if (d.confirmEmailContentDefault !== undefined) {
      this.confirmEmailContentDefault = d.confirmEmailContentDefault;
    }
    if (d.confirmEmailContentDefaultExcerpt !== undefined) {
      this.confirmEmailContentDefaultExcerpt =
        d.confirmEmailContentDefaultExcerpt;
    }
    if (d.confirmEmailTitle !== undefined) {
      this.confirmEmailTitle = d.confirmEmailTitle;
    }
    if (d.confirmEmailTitleDefault !== undefined) {
      this.confirmEmailTitleDefault = d.confirmEmailTitleDefault;
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
      cascadeToSubWorkspaces: this.#cascadeToSubWorkspaces,
      forcedCascadeEmailProvider: this.#forcedCascadeEmailProvider,
      generalEmailProviderId: this.#generalEmailProviderId,
      generalGsmProviderId: this.#generalGsmProviderId,
      inviteToWorkspaceContent: this.#inviteToWorkspaceContent,
      inviteToWorkspaceContentExcerpt: this.#inviteToWorkspaceContentExcerpt,
      inviteToWorkspaceContentDefault: this.#inviteToWorkspaceContentDefault,
      inviteToWorkspaceContentDefaultExcerpt:
        this.#inviteToWorkspaceContentDefaultExcerpt,
      inviteToWorkspaceTitle: this.#inviteToWorkspaceTitle,
      inviteToWorkspaceTitleDefault: this.#inviteToWorkspaceTitleDefault,
      inviteToWorkspaceSenderId: this.#inviteToWorkspaceSenderId,
      accountCenterEmailSenderId: this.#accountCenterEmailSenderId,
      forgetPasswordContent: this.#forgetPasswordContent,
      forgetPasswordContentExcerpt: this.#forgetPasswordContentExcerpt,
      forgetPasswordContentDefault: this.#forgetPasswordContentDefault,
      forgetPasswordContentDefaultExcerpt:
        this.#forgetPasswordContentDefaultExcerpt,
      forgetPasswordTitle: this.#forgetPasswordTitle,
      forgetPasswordTitleDefault: this.#forgetPasswordTitleDefault,
      forgetPasswordSenderId: this.#forgetPasswordSenderId,
      acceptLanguage: this.#acceptLanguage,
      confirmEmailSenderId: this.#confirmEmailSenderId,
      confirmEmailContent: this.#confirmEmailContent,
      confirmEmailContentExcerpt: this.#confirmEmailContentExcerpt,
      confirmEmailContentDefault: this.#confirmEmailContentDefault,
      confirmEmailContentDefaultExcerpt:
        this.#confirmEmailContentDefaultExcerpt,
      confirmEmailTitle: this.#confirmEmailTitle,
      confirmEmailTitleDefault: this.#confirmEmailTitleDefault,
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
      cascadeToSubWorkspaces: "cascadeToSubWorkspaces",
      forcedCascadeEmailProvider: "forcedCascadeEmailProvider",
      generalEmailProviderId: "generalEmailProviderId",
      generalGsmProviderId: "generalGsmProviderId",
      inviteToWorkspaceContent: "inviteToWorkspaceContent",
      inviteToWorkspaceContentExcerpt: "inviteToWorkspaceContentExcerpt",
      inviteToWorkspaceContentDefault: "inviteToWorkspaceContentDefault",
      inviteToWorkspaceContentDefaultExcerpt:
        "inviteToWorkspaceContentDefaultExcerpt",
      inviteToWorkspaceTitle: "inviteToWorkspaceTitle",
      inviteToWorkspaceTitleDefault: "inviteToWorkspaceTitleDefault",
      inviteToWorkspaceSenderId: "inviteToWorkspaceSenderId",
      accountCenterEmailSenderId: "accountCenterEmailSenderId",
      forgetPasswordContent: "forgetPasswordContent",
      forgetPasswordContentExcerpt: "forgetPasswordContentExcerpt",
      forgetPasswordContentDefault: "forgetPasswordContentDefault",
      forgetPasswordContentDefaultExcerpt:
        "forgetPasswordContentDefaultExcerpt",
      forgetPasswordTitle: "forgetPasswordTitle",
      forgetPasswordTitleDefault: "forgetPasswordTitleDefault",
      forgetPasswordSenderId: "forgetPasswordSenderId",
      acceptLanguage: "acceptLanguage",
      confirmEmailSenderId: "confirmEmailSenderId",
      confirmEmailContent: "confirmEmailContent",
      confirmEmailContentExcerpt: "confirmEmailContentExcerpt",
      confirmEmailContentDefault: "confirmEmailContentDefault",
      confirmEmailContentDefaultExcerpt: "confirmEmailContentDefaultExcerpt",
      confirmEmailTitle: "confirmEmailTitle",
      confirmEmailTitleDefault: "confirmEmailTitleDefault",
      workspaceId: "workspaceId",
      userId: "userId",
      createdAt: "createdAt",
      updatedAt: "updatedAt",
    };
  }
  /**
   * Creates an instance of NotificationConfigOptionalDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: NotificationConfigOptionalDtoType) {
    return new NotificationConfigOptionalDto(possibleDtoObject);
  }
  /**
   * Creates an instance of NotificationConfigOptionalDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(
    partialDtoObject: PartialDeep<NotificationConfigOptionalDtoType>,
  ) {
    return new NotificationConfigOptionalDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<NotificationConfigOptionalDtoType>,
  ): InstanceType<typeof NotificationConfigOptionalDto> {
    return new NotificationConfigOptionalDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof NotificationConfigOptionalDto> {
    return new NotificationConfigOptionalDto(this.toJSON());
  }
}
export abstract class NotificationConfigOptionalDtoFactory {
  abstract create(data: unknown): NotificationConfigOptionalDto;
}
/**
 * The base type definition for notificationConfigOptionalDto
 **/
export type NotificationConfigOptionalDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   *
   * @type {boolean}
   **/
  cascadeToSubWorkspaces?: boolean;
  /**
   *
   * @type {boolean}
   **/
  forcedCascadeEmailProvider?: boolean;
  /**
   *
   * @type {string}
   **/
  generalEmailProviderId?: string;
  /**
   *
   * @type {string}
   **/
  generalGsmProviderId?: string;
  /**
   *
   * @type {string}
   **/
  inviteToWorkspaceContent?: string;
  /**
   *
   * @type {string}
   **/
  inviteToWorkspaceContentExcerpt?: string;
  /**
   *
   * @type {string}
   **/
  inviteToWorkspaceContentDefault?: string;
  /**
   *
   * @type {string}
   **/
  inviteToWorkspaceContentDefaultExcerpt?: string;
  /**
   *
   * @type {string}
   **/
  inviteToWorkspaceTitle?: string;
  /**
   *
   * @type {string}
   **/
  inviteToWorkspaceTitleDefault?: string;
  /**
   *
   * @type {string}
   **/
  inviteToWorkspaceSenderId?: string;
  /**
   *
   * @type {string}
   **/
  accountCenterEmailSenderId?: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordContent?: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordContentExcerpt?: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordContentDefault?: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordContentDefaultExcerpt?: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordTitle?: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordTitleDefault?: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordSenderId?: string;
  /**
   *
   * @type {string}
   **/
  acceptLanguage?: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailSenderId?: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailContent?: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailContentExcerpt?: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailContentDefault?: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailContentDefaultExcerpt?: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailTitle?: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailTitleDefault?: string;
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
export namespace NotificationConfigOptionalDtoType {}
