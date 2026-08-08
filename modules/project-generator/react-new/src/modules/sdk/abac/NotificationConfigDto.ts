import { type PartialDeep } from "../sdk/common/fetchx";
/**
 * The base class definition for notificationConfigDto
 **/
export class NotificationConfigDto {
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
  #cascadeToSubWorkspaces!: boolean;
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
  set cascadeToSubWorkspaces(value: boolean) {
    this.#cascadeToSubWorkspaces = Boolean(value);
  }
  setCascadeToSubWorkspaces(value: boolean) {
    this.cascadeToSubWorkspaces = value;
    return this;
  }
  /**
   *
   * @type {boolean}
   **/
  #forcedCascadeEmailProvider!: boolean;
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
  set forcedCascadeEmailProvider(value: boolean) {
    this.#forcedCascadeEmailProvider = Boolean(value);
  }
  setForcedCascadeEmailProvider(value: boolean) {
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
  #inviteToWorkspaceContent: string = "";
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
  set inviteToWorkspaceContent(value: string) {
    this.#inviteToWorkspaceContent = String(value);
  }
  setInviteToWorkspaceContent(value: string) {
    this.inviteToWorkspaceContent = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteToWorkspaceContentExcerpt: string = "";
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
  set inviteToWorkspaceContentExcerpt(value: string) {
    this.#inviteToWorkspaceContentExcerpt = String(value);
  }
  setInviteToWorkspaceContentExcerpt(value: string) {
    this.inviteToWorkspaceContentExcerpt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteToWorkspaceContentDefault: string = "";
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
  set inviteToWorkspaceContentDefault(value: string) {
    this.#inviteToWorkspaceContentDefault = String(value);
  }
  setInviteToWorkspaceContentDefault(value: string) {
    this.inviteToWorkspaceContentDefault = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteToWorkspaceContentDefaultExcerpt: string = "";
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
  set inviteToWorkspaceContentDefaultExcerpt(value: string) {
    this.#inviteToWorkspaceContentDefaultExcerpt = String(value);
  }
  setInviteToWorkspaceContentDefaultExcerpt(value: string) {
    this.inviteToWorkspaceContentDefaultExcerpt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteToWorkspaceTitle: string = "";
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
  set inviteToWorkspaceTitle(value: string) {
    this.#inviteToWorkspaceTitle = String(value);
  }
  setInviteToWorkspaceTitle(value: string) {
    this.inviteToWorkspaceTitle = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #inviteToWorkspaceTitleDefault: string = "";
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
  set inviteToWorkspaceTitleDefault(value: string) {
    this.#inviteToWorkspaceTitleDefault = String(value);
  }
  setInviteToWorkspaceTitleDefault(value: string) {
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
  #forgetPasswordContent: string = "";
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
  set forgetPasswordContent(value: string) {
    this.#forgetPasswordContent = String(value);
  }
  setForgetPasswordContent(value: string) {
    this.forgetPasswordContent = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #forgetPasswordContentExcerpt: string = "";
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
  set forgetPasswordContentExcerpt(value: string) {
    this.#forgetPasswordContentExcerpt = String(value);
  }
  setForgetPasswordContentExcerpt(value: string) {
    this.forgetPasswordContentExcerpt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #forgetPasswordContentDefault: string = "";
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
  set forgetPasswordContentDefault(value: string) {
    this.#forgetPasswordContentDefault = String(value);
  }
  setForgetPasswordContentDefault(value: string) {
    this.forgetPasswordContentDefault = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #forgetPasswordContentDefaultExcerpt: string = "";
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
  set forgetPasswordContentDefaultExcerpt(value: string) {
    this.#forgetPasswordContentDefaultExcerpt = String(value);
  }
  setForgetPasswordContentDefaultExcerpt(value: string) {
    this.forgetPasswordContentDefaultExcerpt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #forgetPasswordTitle: string = "";
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
  set forgetPasswordTitle(value: string) {
    this.#forgetPasswordTitle = String(value);
  }
  setForgetPasswordTitle(value: string) {
    this.forgetPasswordTitle = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #forgetPasswordTitleDefault: string = "";
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
  set forgetPasswordTitleDefault(value: string) {
    this.#forgetPasswordTitleDefault = String(value);
  }
  setForgetPasswordTitleDefault(value: string) {
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
  #acceptLanguage: string = "";
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
  set acceptLanguage(value: string) {
    this.#acceptLanguage = String(value);
  }
  setAcceptLanguage(value: string) {
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
  #confirmEmailContent: string = "";
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
  set confirmEmailContent(value: string) {
    this.#confirmEmailContent = String(value);
  }
  setConfirmEmailContent(value: string) {
    this.confirmEmailContent = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #confirmEmailContentExcerpt: string = "";
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
  set confirmEmailContentExcerpt(value: string) {
    this.#confirmEmailContentExcerpt = String(value);
  }
  setConfirmEmailContentExcerpt(value: string) {
    this.confirmEmailContentExcerpt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #confirmEmailContentDefault: string = "";
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
  set confirmEmailContentDefault(value: string) {
    this.#confirmEmailContentDefault = String(value);
  }
  setConfirmEmailContentDefault(value: string) {
    this.confirmEmailContentDefault = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #confirmEmailContentDefaultExcerpt: string = "";
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
  set confirmEmailContentDefaultExcerpt(value: string) {
    this.#confirmEmailContentDefaultExcerpt = String(value);
  }
  setConfirmEmailContentDefaultExcerpt(value: string) {
    this.confirmEmailContentDefaultExcerpt = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #confirmEmailTitle: string = "";
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
  set confirmEmailTitle(value: string) {
    this.#confirmEmailTitle = String(value);
  }
  setConfirmEmailTitle(value: string) {
    this.confirmEmailTitle = value;
    return this;
  }
  /**
   *
   * @type {string}
   **/
  #confirmEmailTitleDefault: string = "";
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
  set confirmEmailTitleDefault(value: string) {
    this.#confirmEmailTitleDefault = String(value);
  }
  setConfirmEmailTitleDefault(value: string) {
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
    const d = data as Partial<NotificationConfigDto>;
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
   * Creates an instance of NotificationConfigDto, and possibleDtoObject
   * needs to satisfy the type requirement fully, otherwise typescript compile would
   * be complaining.
   **/
  static from(possibleDtoObject: NotificationConfigDtoType) {
    return new NotificationConfigDto(possibleDtoObject);
  }
  /**
   * Creates an instance of NotificationConfigDto, and partialDtoObject
   * needs to satisfy the type, but partially, and rest of the content would
   * be constructed according to data types and nullability.
   **/
  static with(partialDtoObject: PartialDeep<NotificationConfigDtoType>) {
    return new NotificationConfigDto(partialDtoObject);
  }
  copyWith(
    partial: PartialDeep<NotificationConfigDtoType>,
  ): InstanceType<typeof NotificationConfigDto> {
    return new NotificationConfigDto({ ...this.toJSON(), ...partial });
  }
  clone(): InstanceType<typeof NotificationConfigDto> {
    return new NotificationConfigDto(this.toJSON());
  }
}
export abstract class NotificationConfigDtoFactory {
  abstract create(data: unknown): NotificationConfigDto;
}
/**
 * The base type definition for notificationConfigDto
 **/
export type NotificationConfigDtoType = {
  /**
   *
   * @type {string}
   **/
  uniqueId?: string;
  /**
   *
   * @type {boolean}
   **/
  cascadeToSubWorkspaces: boolean;
  /**
   *
   * @type {boolean}
   **/
  forcedCascadeEmailProvider: boolean;
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
  inviteToWorkspaceContent: string;
  /**
   *
   * @type {string}
   **/
  inviteToWorkspaceContentExcerpt: string;
  /**
   *
   * @type {string}
   **/
  inviteToWorkspaceContentDefault: string;
  /**
   *
   * @type {string}
   **/
  inviteToWorkspaceContentDefaultExcerpt: string;
  /**
   *
   * @type {string}
   **/
  inviteToWorkspaceTitle: string;
  /**
   *
   * @type {string}
   **/
  inviteToWorkspaceTitleDefault: string;
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
  forgetPasswordContent: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordContentExcerpt: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordContentDefault: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordContentDefaultExcerpt: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordTitle: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordTitleDefault: string;
  /**
   *
   * @type {string}
   **/
  forgetPasswordSenderId?: string;
  /**
   *
   * @type {string}
   **/
  acceptLanguage: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailSenderId?: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailContent: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailContentExcerpt: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailContentDefault: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailContentDefaultExcerpt: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailTitle: string;
  /**
   *
   * @type {string}
   **/
  confirmEmailTitleDefault: string;
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
export namespace NotificationConfigDtoType {}
