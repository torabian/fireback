/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { IError, QueryFilterRequest, RemoveReply } from "../../core/common";

export const protobufPackage = "";

export interface AppMenuCreateReply {
  data: AppMenuEntity | undefined;
  error: IError | undefined;
}

export interface AppMenuReply {
  data: AppMenuEntity | undefined;
  error: IError | undefined;
}

export interface AppMenuQueryReply {
  items: AppMenuEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface AppMenuEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: AppMenuEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(gorm:"foreignKey:LinkerId;references:UniqueId" json:"translations") */
  translations: AppMenuEntityPolyglot[];
  /** @tag(  yaml:"href"  ) */
  href?: string | undefined;
  /** @tag(  yaml:"icon"  ) */
  icon?: string | undefined;
  /** @tag(translate:"true"  yaml:"label"  ) */
  label?: string | undefined;
  /** @tag(  yaml:"activeMatcher"  ) */
  activeMatcher?: string | undefined;
  /** @tag(  yaml:"applyType"  ) */
  applyType?: string | undefined;
  /** One 2 one to external entity */
  capabilityId?: string | undefined;
  /** @tag(gorm:"foreignKey:CapabilityId;references:UniqueId" json:"capability" yaml:"capability" fbtype:"one") */
  capability: CapabilityEntity | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
  /** @tag(gorm:"-" sql:"-") */
  children: AppMenuEntity[];
}

/** Because it has translation field, we need a translation table for this */
export interface AppMenuEntityPolyglot {
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId") */
  linkerId: string;
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId") */
  languageId: string;
  /** @tag(yaml:"label" json:"label"); */
  label: string;
}

export interface BackupReqDto {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: BackupReqDto | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(yaml:"entities" fbtype:"arrayp") */
  entities: string[];
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface BackupTableMetaCreateReply {
  data: BackupTableMetaEntity | undefined;
  error: IError | undefined;
}

export interface BackupTableMetaReply {
  data: BackupTableMetaEntity | undefined;
  error: IError | undefined;
}

export interface BackupTableMetaQueryReply {
  items: BackupTableMetaEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface BackupTableMetaEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: BackupTableMetaEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"tableNameInDb"  ) */
  tableNameInDb?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface CapabilityCreateReply {
  data: CapabilityEntity | undefined;
  error: IError | undefined;
}

export interface CapabilityReply {
  data: CapabilityEntity | undefined;
  error: IError | undefined;
}

export interface CapabilityQueryReply {
  items: CapabilityEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface CapabilityEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: CapabilityEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"name"  ) */
  name?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface CapabilityChild {
  uniqueId: string;
  children: CapabilityChild[];
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
}

export interface CapabilitiesResult {
  capabilities: CapabilityEntity[];
  nested: CapabilityChild[];
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
}

export interface EmailConfirmationCreateReply {
  data: EmailConfirmationEntity | undefined;
  error: IError | undefined;
}

export interface EmailConfirmationReply {
  data: EmailConfirmationEntity | undefined;
  error: IError | undefined;
}

export interface EmailConfirmationQueryReply {
  items: EmailConfirmationEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface EmailConfirmationEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: EmailConfirmationEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** One 2 one to external entity */
  user: UserEntity | undefined;
  /** @tag(  yaml:"status"  ) */
  status?: string | undefined;
  /** @tag(  yaml:"email"  ) */
  email?: string | undefined;
  /** @tag(  yaml:"key"  ) */
  key?: string | undefined;
  /** @tag(  yaml:"expiresAt"  ) */
  expiresAt?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface EmailProviderCreateReply {
  data: EmailProviderEntity | undefined;
  error: IError | undefined;
}

export interface EmailProviderReply {
  data: EmailProviderEntity | undefined;
  error: IError | undefined;
}

export interface EmailProviderQueryReply {
  items: EmailProviderEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface EmailProviderEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: EmailProviderEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"type"  ) */
  type?: string | undefined;
  /** @tag(  yaml:"apiKey"  ) */
  apiKey?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface EmailSenderCreateReply {
  data: EmailSenderEntity | undefined;
  error: IError | undefined;
}

export interface EmailSenderReply {
  data: EmailSenderEntity | undefined;
  error: IError | undefined;
}

export interface EmailSenderQueryReply {
  items: EmailSenderEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface EmailSenderEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: EmailSenderEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"fromName"  ) */
  fromName?: string | undefined;
  /** @tag(  yaml:"fromEmailAddress"  ) */
  fromEmailAddress?: string | undefined;
  /** @tag(  yaml:"replyTo"  ) */
  replyTo?: string | undefined;
  /** @tag(  yaml:"nickName"  ) */
  nickName?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface TestMailDto {
  /** EmailSenderEntity senderId = 1; */
  senderId: string;
  toName: string;
  toEmail: string;
  subject: string;
  content: string;
}

export interface ForgetPasswordCreateReply {
  data: ForgetPasswordEntity | undefined;
  error: IError | undefined;
}

export interface ForgetPasswordReply {
  data: ForgetPasswordEntity | undefined;
  error: IError | undefined;
}

export interface ForgetPasswordQueryReply {
  items: ForgetPasswordEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface ForgetPasswordEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: ForgetPasswordEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** One 2 one to external entity */
  user: UserEntity | undefined;
  /** One 2 one to external entity */
  passportId?: string | undefined;
  /** @tag(gorm:"foreignKey:PassportId;references:UniqueId" json:"passport" yaml:"passport" fbtype:"one") */
  passport: PassportEntity | undefined;
  /** @tag(  yaml:"status"  ) */
  status?: string | undefined;
  validUntil: number;
  validUntilFormatted: string;
  blockedUntil: number;
  blockedUntilFormatted: string;
  /** @tag(  yaml:"secondsToUnblock"  ) */
  secondsToUnblock?: number | undefined;
  /** @tag(  yaml:"otp"  ) */
  otp?: string | undefined;
  /** @tag(  yaml:"recoveryAbsoluteUrl"  ) */
  recoveryAbsoluteUrl?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface GsmProviderCreateReply {
  data: GsmProviderEntity | undefined;
  error: IError | undefined;
}

export interface GsmProviderReply {
  data: GsmProviderEntity | undefined;
  error: IError | undefined;
}

export interface GsmProviderQueryReply {
  items: GsmProviderEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface GsmProviderEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: GsmProviderEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"apiKey"  ) */
  apiKey?: string | undefined;
  /** @tag(  yaml:"mainSenderNumber"  ) */
  mainSenderNumber?: string | undefined;
  /** @tag(  yaml:"type"  ) */
  type?: string | undefined;
  /** @tag(  yaml:"invokeUrl"  ) */
  invokeUrl?: string | undefined;
  /** @tag(  yaml:"invokeBody"  ) */
  invokeBody?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface NotificationConfigCreateReply {
  data: NotificationConfigEntity | undefined;
  error: IError | undefined;
}

export interface NotificationConfigReply {
  data: NotificationConfigEntity | undefined;
  error: IError | undefined;
}

export interface NotificationConfigQueryReply {
  items: NotificationConfigEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface NotificationConfigEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: NotificationConfigEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"cascadeToSubWorkspaces") */
  cascadeToSubWorkspaces?: boolean | undefined;
  /** @tag(  yaml:"forcedCascadeEmailProvider") */
  forcedCascadeEmailProvider?: boolean | undefined;
  /** One 2 one to external entity */
  generalEmailProviderId?: string | undefined;
  /** @tag(gorm:"foreignKey:GeneralEmailProviderId;references:UniqueId" json:"generalEmailProvider" yaml:"generalEmailProvider" fbtype:"one") */
  generalEmailProvider: EmailProviderEntity | undefined;
  /** One 2 one to external entity */
  generalGsmProviderId?: string | undefined;
  /** @tag(gorm:"foreignKey:GeneralGsmProviderId;references:UniqueId" json:"generalGsmProvider" yaml:"generalGsmProvider" fbtype:"one") */
  generalGsmProvider: GsmProviderEntity | undefined;
  /** @tag(  yaml:"inviteToWorkspaceContent"  gorm:"text") */
  inviteToWorkspaceContent?: string | undefined;
  /** @tag(  yaml:"inviteToWorkspaceContentExcerpt"  gorm:"text") */
  inviteToWorkspaceContentExcerpt?: string | undefined;
  /** @tag(  yaml:"inviteToWorkspaceContentDefault"  gorm:"text") */
  inviteToWorkspaceContentDefault?: string | undefined;
  /** @tag(  yaml:"inviteToWorkspaceContentDefaultExcerpt"  gorm:"text") */
  inviteToWorkspaceContentDefaultExcerpt?: string | undefined;
  /** @tag(  yaml:"inviteToWorkspaceTitle"  ) */
  inviteToWorkspaceTitle?: string | undefined;
  /** @tag(  yaml:"inviteToWorkspaceTitleDefault"  ) */
  inviteToWorkspaceTitleDefault?: string | undefined;
  /** One 2 one to external entity */
  inviteToWorkspaceSenderId?: string | undefined;
  /** @tag(gorm:"foreignKey:InviteToWorkspaceSenderId;references:UniqueId" json:"inviteToWorkspaceSender" yaml:"inviteToWorkspaceSender" fbtype:"one") */
  inviteToWorkspaceSender: EmailSenderEntity | undefined;
  /** @tag(  yaml:"forgetPasswordContent"  gorm:"text") */
  forgetPasswordContent?: string | undefined;
  /** @tag(  yaml:"forgetPasswordContentExcerpt"  gorm:"text") */
  forgetPasswordContentExcerpt?: string | undefined;
  /** @tag(  yaml:"forgetPasswordContentDefault"  gorm:"text") */
  forgetPasswordContentDefault?: string | undefined;
  /** @tag(  yaml:"forgetPasswordContentDefaultExcerpt"  gorm:"text") */
  forgetPasswordContentDefaultExcerpt?: string | undefined;
  /** @tag(  yaml:"forgetPasswordTitle"  gorm:"text") */
  forgetPasswordTitle?: string | undefined;
  /** @tag(  yaml:"forgetPasswordTitleDefault"  gorm:"text") */
  forgetPasswordTitleDefault?: string | undefined;
  /** One 2 one to external entity */
  forgetPasswordSenderId?: string | undefined;
  /** @tag(gorm:"foreignKey:ForgetPasswordSenderId;references:UniqueId" json:"forgetPasswordSender" yaml:"forgetPasswordSender" fbtype:"one") */
  forgetPasswordSender: EmailSenderEntity | undefined;
  /** @tag(  yaml:"acceptLanguage" gorm:"text") */
  acceptLanguage?: string | undefined;
  /** @tag( yaml:"acceptLanguageExcerpt" gorm:"text") */
  acceptLanguageExcerpt?: string | undefined;
  /** One 2 one to external entity */
  confirmEmailSenderId?: string | undefined;
  /** @tag(gorm:"foreignKey:ConfirmEmailSenderId;references:UniqueId" json:"confirmEmailSender" yaml:"confirmEmailSender" fbtype:"one") */
  confirmEmailSender: EmailSenderEntity | undefined;
  /** @tag(  yaml:"confirmEmailContent"  gorm:"text") */
  confirmEmailContent?: string | undefined;
  /** @tag(  yaml:"confirmEmailContentExcerpt"  gorm:"text") */
  confirmEmailContentExcerpt?: string | undefined;
  /** @tag(  yaml:"confirmEmailContentDefault"  gorm:"text") */
  confirmEmailContentDefault?: string | undefined;
  /** @tag(  yaml:"confirmEmailContentDefaultExcerpt"  gorm:"text") */
  confirmEmailContentDefaultExcerpt?: string | undefined;
  /** @tag(  yaml:"confirmEmailTitle"  ) */
  confirmEmailTitle?: string | undefined;
  /** @tag(  yaml:"confirmEmailTitleDefault"  ) */
  confirmEmailTitleDefault?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface CreateByMailResponse {
  error: IError | undefined;
  data: UserSessionDto | undefined;
}

export interface EmailAccountSignupDto {
  /** @tag(validate:"required") */
  email: string;
  /** @tag(validate:"required") */
  password: string;
  /** @tag(validate:"required") */
  firstName: string;
  /** @tag(validate:"required") */
  lastName: string;
  inviteId: string;
  publicJoinKeyId: string;
  workspaceTypeId: string;
}

export interface EmailAccountSigninDto {
  /** @tag(validate:"required") */
  email: string;
  /** @tag(validate:"required") */
  password: string;
}

export interface PhoneNumberAccountCreationDto {
  /** @tag(validate:"required") */
  phoneNumber: string;
}

export interface PhoneNumberUniversalAuthenticateDto {
  /** @tag(validate:"required") */
  phoneNumber: string;
}

export interface UserSessionDto {
  passport: PassportEntity | undefined;
  token: string;
  exchangeKey: string;
  /** @tag(json:"userRoleWorkspaces") */
  userRoleWorkspaces: UserRoleWorkspaceEntity[];
  user: UserEntity | undefined;
}

export interface OtpAuthenticateDto {
  /** @tag(validate:"required") */
  value: string;
  otp: string;
  /** @tag(validate:"required") */
  type: string;
  password: string;
}

export interface EmailOtpResponse {
  request: ForgetPasswordEntity | undefined;
  userSession: UserSessionDto | undefined;
}

export interface ResetEmailDto {
  password: string;
}

export interface PassportCreateReply {
  data: PassportEntity | undefined;
  error: IError | undefined;
}

export interface PassportReply {
  data: PassportEntity | undefined;
  error: IError | undefined;
}

export interface PassportQueryReply {
  items: PassportEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface PassportEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: PassportEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag( validate:"required" yaml:"type"  ) */
  type?: string | undefined;
  /** @tag( validate:"required" yaml:"value"  ) */
  value?: string | undefined;
  /** @tag(  yaml:"password"  ) */
  password?: string | undefined;
  /** @tag(  yaml:"confirmed") */
  confirmed?: boolean | undefined;
  /** @tag(  yaml:"accessToken"  ) */
  accessToken?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface PassportMethodCreateReply {
  data: PassportMethodEntity | undefined;
  error: IError | undefined;
}

export interface PassportMethodReply {
  data: PassportMethodEntity | undefined;
  error: IError | undefined;
}

export interface PassportMethodQueryReply {
  items: PassportMethodEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface PassportMethodEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: PassportMethodEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(gorm:"foreignKey:LinkerId;references:UniqueId" json:"translations") */
  translations: PassportMethodEntityPolyglot[];
  /** @tag(translate:"true" validate:"required" yaml:"name"  ) */
  name?: string | undefined;
  /** @tag( validate:"required" yaml:"type"  ) */
  type?: string | undefined;
  /** @tag( validate:"required" yaml:"region"  ) */
  region?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

/** Because it has translation field, we need a translation table for this */
export interface PassportMethodEntityPolyglot {
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId") */
  linkerId: string;
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId") */
  languageId: string;
  /** @tag(yaml:"name" json:"name"); */
  name: string;
}

export interface PendingWorkspaceInviteCreateReply {
  data: PendingWorkspaceInviteEntity | undefined;
  error: IError | undefined;
}

export interface PendingWorkspaceInviteReply {
  data: PendingWorkspaceInviteEntity | undefined;
  error: IError | undefined;
}

export interface PendingWorkspaceInviteQueryReply {
  items: PendingWorkspaceInviteEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface PendingWorkspaceInviteEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: PendingWorkspaceInviteEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"value"  ) */
  value?: string | undefined;
  /** @tag(  yaml:"type"  ) */
  type?: string | undefined;
  /** @tag(  yaml:"coverLetter"  ) */
  coverLetter?: string | undefined;
  /** @tag(  yaml:"workspaceName"  ) */
  workspaceName?: string | undefined;
  /** One 2 one to external entity */
  roleId?: string | undefined;
  /** @tag(gorm:"foreignKey:RoleId;references:UniqueId" json:"role" yaml:"role" fbtype:"one") */
  role: RoleEntity | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface PhoneConfirmationCreateReply {
  data: PhoneConfirmationEntity | undefined;
  error: IError | undefined;
}

export interface PhoneConfirmationReply {
  data: PhoneConfirmationEntity | undefined;
  error: IError | undefined;
}

export interface PhoneConfirmationQueryReply {
  items: PhoneConfirmationEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface PhoneConfirmationEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: PhoneConfirmationEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** One 2 one to external entity */
  user: UserEntity | undefined;
  /** @tag(  yaml:"status"  ) */
  status?: string | undefined;
  /** @tag(  yaml:"phoneNumber"  ) */
  phoneNumber?: string | undefined;
  /** @tag(  yaml:"key"  ) */
  key?: string | undefined;
  /** @tag(  yaml:"expiresAt"  ) */
  expiresAt?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface PreferenceCreateReply {
  data: PreferenceEntity | undefined;
  error: IError | undefined;
}

export interface PreferenceReply {
  data: PreferenceEntity | undefined;
  error: IError | undefined;
}

export interface PreferenceQueryReply {
  items: PreferenceEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface PreferenceEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: PreferenceEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"timezone"  ) */
  timezone?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface PublicJoinKeyCreateReply {
  data: PublicJoinKeyEntity | undefined;
  error: IError | undefined;
}

export interface PublicJoinKeyReply {
  data: PublicJoinKeyEntity | undefined;
  error: IError | undefined;
}

export interface PublicJoinKeyQueryReply {
  items: PublicJoinKeyEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface PublicJoinKeyEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: PublicJoinKeyEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** One 2 one to external entity */
  roleId?: string | undefined;
  /** @tag(gorm:"foreignKey:RoleId;references:UniqueId" json:"role" yaml:"role" fbtype:"one") */
  role: RoleEntity | undefined;
  /** One 2 one to external entity */
  workspace: WorkspaceEntity | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface RoleCreateReply {
  data: RoleEntity | undefined;
  error: IError | undefined;
}

export interface RoleReply {
  data: RoleEntity | undefined;
  error: IError | undefined;
}

export interface RoleQueryReply {
  items: RoleEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface RoleEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: RoleEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag( validate:"required,omitempty,min=1,max=200" yaml:"name"  ) */
  name?: string | undefined;
  /** Many 2 many entities */
  capabilitiesListId: string[];
  /** @tag(gorm:"many2many:role_capabilities;foreignKey:UniqueId;references:UniqueId" yaml:"capabilities" fbtype:"many2many") */
  capabilities: CapabilityEntity[];
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface TableViewSizingCreateReply {
  data: TableViewSizingEntity | undefined;
  error: IError | undefined;
}

export interface TableViewSizingReply {
  data: TableViewSizingEntity | undefined;
  error: IError | undefined;
}

export interface TableViewSizingQueryReply {
  items: TableViewSizingEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface TableViewSizingEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: TableViewSizingEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag( validate:"required" yaml:"tableName"  ) */
  tableName?: string | undefined;
  /** @tag(  yaml:"sizes"  ) */
  sizes?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface TokenCreateReply {
  data: TokenEntity | undefined;
  error: IError | undefined;
}

export interface TokenReply {
  data: TokenEntity | undefined;
  error: IError | undefined;
}

export interface TokenQueryReply {
  items: TokenEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface TokenEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: TokenEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** One 2 one to external entity */
  user: UserEntity | undefined;
  /** @tag(  yaml:"validUntil"  ) */
  validUntil?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface UserCreateReply {
  data: UserEntity | undefined;
  error: IError | undefined;
}

export interface UserReply {
  data: UserEntity | undefined;
  error: IError | undefined;
}

export interface UserQueryReply {
  items: UserEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface UserEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: UserEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"firstName"  ) */
  firstName?: string | undefined;
  /** @tag(  yaml:"lastName"  ) */
  lastName?: string | undefined;
  /** @tag(  yaml:"photo"  ) */
  photo?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface UserProfileCreateReply {
  data: UserProfileEntity | undefined;
  error: IError | undefined;
}

export interface UserProfileReply {
  data: UserProfileEntity | undefined;
  error: IError | undefined;
}

export interface UserProfileQueryReply {
  items: UserProfileEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface UserProfileEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: UserProfileEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"firstName"  ) */
  firstName?: string | undefined;
  /** @tag(  yaml:"lastName"  ) */
  lastName?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface UserRoleWorkspaceCreateReply {
  data: UserRoleWorkspaceEntity | undefined;
  error: IError | undefined;
}

export interface UserRoleWorkspaceReply {
  data: UserRoleWorkspaceEntity | undefined;
  error: IError | undefined;
}

export interface UserRoleWorkspaceQueryReply {
  items: UserRoleWorkspaceEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface UserRoleWorkspaceEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: UserRoleWorkspaceEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** One 2 one to external entity */
  user: UserEntity | undefined;
  /** One 2 one to external entity */
  roleId?: string | undefined;
  /** @tag(gorm:"foreignKey:RoleId;references:UniqueId" json:"role" yaml:"role" fbtype:"one") */
  role: RoleEntity | undefined;
  /** One 2 one to external entity */
  workspace: WorkspaceEntity | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface AcceptInviteDto {
  inviteUniqueId: string;
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
}

export interface AssignRoleDto {
  /** @tag(validate:"required") */
  roleId: string;
  /** @tag(validate:"required") */
  userId: string;
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
}

export interface WorkspaceDto {
  relations: UserRoleWorkspaceEntity[];
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
}

export interface ExchangeKeyInformationDto {
  key: string;
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
}

export interface UserAccessLevel {
  /** @tag(json:"capabilities") */
  capabilities: string[];
  /** @tag(json:"workspaces") */
  workspaces: string[];
  /** @tag(json:"sql") */
  SQL: string;
}

export interface AuthResult {
  /** @tag(json:"workspaceId") */
  workspaceId: string;
  /** @tag(json:"internalSql") */
  internalSql: string;
  /** @tag(json:"userId") */
  userId: string;
  /** @tag(json:"user") */
  user: UserEntity | undefined;
  /** @tag(json:"accessLevel") */
  accessLevel?: UserAccessLevel | undefined;
  /** @tag(json:"userHas") */
  userHas: string[];
}

export interface AuthContext {
  skipWorkspaceId: boolean;
  workspaceId: string;
  token: string;
  capabilities: string[];
}

export interface ReactiveSearchResultDto {
  /** @tag(yaml:"parent") */
  parent?: ReactiveSearchResultDto | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(  yaml:"phrase"  ) */
  phrase?: string | undefined;
  /** @tag(  yaml:"icon"  ) */
  icon?: string | undefined;
  /** @tag(  yaml:"description"  ) */
  description?: string | undefined;
  /** @tag(  yaml:"group"  ) */
  group?: string | undefined;
  /** @tag(  yaml:"uiLocation"  ) */
  uiLocation?: string | undefined;
  /** @tag(  yaml:"actionFn"  ) */
  actionFn?: string | undefined;
}

export interface WorkspaceConfigCreateReply {
  data: WorkspaceConfigEntity | undefined;
  error: IError | undefined;
}

export interface WorkspaceConfigReply {
  data: WorkspaceConfigEntity | undefined;
  error: IError | undefined;
}

export interface WorkspaceConfigQueryReply {
  items: WorkspaceConfigEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface WorkspaceConfigEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: WorkspaceConfigEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"disablePublicWorkspaceCreation"  ) */
  disablePublicWorkspaceCreation?: number | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface WorkspaceConfigDto {
  /** @tag(gorm:"foreignKey:WorkspaceId;references:UniqueId") */
  workspace: WorkspaceEntity | undefined;
  /** @tag(gorm:"size:100;") */
  workspaceId?: string | undefined;
  zoomClientId?: string | undefined;
  zoomClientSecret?: string | undefined;
  allowPublicToJoinTheWorkspace?: boolean | undefined;
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
}

export interface WorkspaceCreateReply {
  data: WorkspaceEntity | undefined;
  error: IError | undefined;
}

export interface WorkspaceReply {
  data: WorkspaceEntity | undefined;
  error: IError | undefined;
}

export interface WorkspaceQueryReply {
  items: WorkspaceEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface WorkspaceEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: WorkspaceEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"description"  ) */
  description?: string | undefined;
  /** @tag(  yaml:"name"  ) */
  name?: string | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
  /** @tag(gorm:"-" sql:"-") */
  children: WorkspaceEntity[];
}

export interface WorkspaceInviteCreateReply {
  data: WorkspaceInviteEntity | undefined;
  error: IError | undefined;
}

export interface WorkspaceInviteReply {
  data: WorkspaceInviteEntity | undefined;
  error: IError | undefined;
}

export interface WorkspaceInviteQueryReply {
  items: WorkspaceInviteEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface WorkspaceInviteEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: WorkspaceInviteEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(  yaml:"passportMode"  ) */
  passportMode?: string | undefined;
  /** @tag(  yaml:"coverLetter"  ) */
  coverLetter?: string | undefined;
  /** @tag(  yaml:"targetUserLocale"  ) */
  targetUserLocale?: string | undefined;
  /** @tag(  yaml:"email"  ) */
  email?: string | undefined;
  /** One 2 one to external entity */
  workspace: WorkspaceEntity | undefined;
  /** @tag(  yaml:"firstName"  ) */
  firstName?: string | undefined;
  /** @tag(  yaml:"lastName"  ) */
  lastName?: string | undefined;
  /** @tag(  yaml:"inviteeUserId"  ) */
  inviteeUserId?: string | undefined;
  /** @tag(  yaml:"phoneNumber"  ) */
  phoneNumber?: string | undefined;
  /** One 2 one to external entity */
  roleId?: string | undefined;
  /** @tag(gorm:"foreignKey:RoleId;references:UniqueId" json:"role" yaml:"role" fbtype:"one") */
  role: RoleEntity | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

export interface WorkspaceTypeCreateReply {
  data: WorkspaceTypeEntity | undefined;
  error: IError | undefined;
}

export interface WorkspaceTypeReply {
  data: WorkspaceTypeEntity | undefined;
  error: IError | undefined;
}

export interface WorkspaceTypeQueryReply {
  items: WorkspaceTypeEntity[];
  totalItems: number;
  itemsPerPage: number;
  startIndex: number;
  error: IError | undefined;
}

export interface WorkspaceTypeEntity {
  /** @tag(yaml:"visibility") */
  visibility?: string | undefined;
  /** @tag(yaml:"workspaceId") */
  workspaceId?: string | undefined;
  /** @tag(yaml:"linkerId") */
  linkerId?: string | undefined;
  /** @tag(yaml:"parentId") */
  parentId?: string | undefined;
  /** @tag(yaml:"parent") */
  parent?: WorkspaceTypeEntity | undefined;
  /** @tag(gorm:"primarykey;uniqueId;unique;not null;size:100;" yaml:"uniqueId") */
  uniqueId: string;
  /** @tag(yaml:"userId") */
  userId?: string | undefined;
  /** @tag(gorm:"foreignKey:LinkerId;references:UniqueId" json:"translations") */
  translations: WorkspaceTypeEntityPolyglot[];
  /** @tag(translate:"true" validate:"required,omitempty,min=1,max=250" yaml:"title"  ) */
  title?: string | undefined;
  /** @tag(translate:"true"  yaml:"description"  ) */
  description?: string | undefined;
  /** @tag( validate:"required,omitempty,min=2,max=50" yaml:"slug"  gorm:"unique;not null;size:100;") */
  slug?: string | undefined;
  /** One 2 one to external entity */
  roleId?: string | undefined;
  /** @tag(gorm:"foreignKey:RoleId;references:UniqueId" json:"role" yaml:"role" fbtype:"one") */
  role: RoleEntity | undefined;
  /** @tag(gorm:"type:int;name:rank") */
  rank: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  updated: number;
  /** @tag(gorm:"autoUpdateTime:nano") */
  created: number;
  /** @tag(sql:"-") */
  createdFormatted: string;
  /** @tag(sql:"-") */
  updatedFormatted: string;
}

/** Because it has translation field, we need a translation table for this */
export interface WorkspaceTypeEntityPolyglot {
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"linkerId" yaml:"linkerId") */
  linkerId: string;
  /** @tag(gorm:"uniqueId;not null;size:100;" json:"languageId" yaml:"languageId") */
  languageId: string;
  /** @tag(yaml:"title" json:"title"); */
  title: string;
  /** @tag(yaml:"description" json:"description"); */
  description: string;
}

function createBaseAppMenuCreateReply(): AppMenuCreateReply {
  return { data: undefined, error: undefined };
}

export const AppMenuCreateReply = {
  encode(
    message: AppMenuCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      AppMenuEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppMenuCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppMenuCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = AppMenuEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppMenuCreateReply {
    return {
      data: isSet(object.data)
        ? AppMenuEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: AppMenuCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? AppMenuEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<AppMenuCreateReply>, I>>(
    base?: I
  ): AppMenuCreateReply {
    return AppMenuCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AppMenuCreateReply>, I>>(
    object: I
  ): AppMenuCreateReply {
    const message = createBaseAppMenuCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? AppMenuEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseAppMenuReply(): AppMenuReply {
  return { data: undefined, error: undefined };
}

export const AppMenuReply = {
  encode(
    message: AppMenuReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      AppMenuEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppMenuReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppMenuReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = AppMenuEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppMenuReply {
    return {
      data: isSet(object.data)
        ? AppMenuEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: AppMenuReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? AppMenuEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<AppMenuReply>, I>>(
    base?: I
  ): AppMenuReply {
    return AppMenuReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AppMenuReply>, I>>(
    object: I
  ): AppMenuReply {
    const message = createBaseAppMenuReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? AppMenuEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseAppMenuQueryReply(): AppMenuQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const AppMenuQueryReply = {
  encode(
    message: AppMenuQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      AppMenuEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppMenuQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppMenuQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(AppMenuEntity.decode(reader, reader.uint32()));
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppMenuQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => AppMenuEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: AppMenuQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? AppMenuEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<AppMenuQueryReply>, I>>(
    base?: I
  ): AppMenuQueryReply {
    return AppMenuQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AppMenuQueryReply>, I>>(
    object: I
  ): AppMenuQueryReply {
    const message = createBaseAppMenuQueryReply();
    message.items =
      object.items?.map((e) => AppMenuEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseAppMenuEntity(): AppMenuEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    translations: [],
    href: undefined,
    icon: undefined,
    label: undefined,
    activeMatcher: undefined,
    applyType: undefined,
    capabilityId: undefined,
    capability: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
    children: [],
  };
}

export const AppMenuEntity = {
  encode(
    message: AppMenuEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      AppMenuEntity.encode(message.parent, writer.uint32(42).fork()).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    for (const v of message.translations) {
      AppMenuEntityPolyglot.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.href !== undefined) {
      writer.uint32(82).string(message.href);
    }
    if (message.icon !== undefined) {
      writer.uint32(90).string(message.icon);
    }
    if (message.label !== undefined) {
      writer.uint32(98).string(message.label);
    }
    if (message.activeMatcher !== undefined) {
      writer.uint32(106).string(message.activeMatcher);
    }
    if (message.applyType !== undefined) {
      writer.uint32(114).string(message.applyType);
    }
    if (message.capabilityId !== undefined) {
      writer.uint32(130).string(message.capabilityId);
    }
    if (message.capability !== undefined) {
      CapabilityEntity.encode(
        message.capability,
        writer.uint32(138).fork()
      ).ldelim();
    }
    if (message.rank !== 0) {
      writer.uint32(144).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(152).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(160).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(170).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(178).string(message.updatedFormatted);
    }
    for (const v of message.children) {
      AppMenuEntity.encode(v!, writer.uint32(186).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppMenuEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppMenuEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = AppMenuEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 8:
          message.translations.push(
            AppMenuEntityPolyglot.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.href = reader.string();
          break;
        case 11:
          message.icon = reader.string();
          break;
        case 12:
          message.label = reader.string();
          break;
        case 13:
          message.activeMatcher = reader.string();
          break;
        case 14:
          message.applyType = reader.string();
          break;
        case 16:
          message.capabilityId = reader.string();
          break;
        case 17:
          message.capability = CapabilityEntity.decode(reader, reader.uint32());
          break;
        case 18:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 19:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 20:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 21:
          message.createdFormatted = reader.string();
          break;
        case 22:
          message.updatedFormatted = reader.string();
          break;
        case 23:
          message.children.push(AppMenuEntity.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppMenuEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? AppMenuEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      translations: Array.isArray(object?.translations)
        ? object.translations.map((e: any) => AppMenuEntityPolyglot.fromJSON(e))
        : [],
      href: isSet(object.href) ? String(object.href) : undefined,
      icon: isSet(object.icon) ? String(object.icon) : undefined,
      label: isSet(object.label) ? String(object.label) : undefined,
      activeMatcher: isSet(object.activeMatcher)
        ? String(object.activeMatcher)
        : undefined,
      applyType: isSet(object.applyType) ? String(object.applyType) : undefined,
      capabilityId: isSet(object.capabilityId)
        ? String(object.capabilityId)
        : undefined,
      capability: isSet(object.capability)
        ? CapabilityEntity.fromJSON(object.capability)
        : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
      children: Array.isArray(object?.children)
        ? object.children.map((e: any) => AppMenuEntity.fromJSON(e))
        : [],
    };
  },

  toJSON(message: AppMenuEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? AppMenuEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    if (message.translations) {
      obj.translations = message.translations.map((e) =>
        e ? AppMenuEntityPolyglot.toJSON(e) : undefined
      );
    } else {
      obj.translations = [];
    }
    message.href !== undefined && (obj.href = message.href);
    message.icon !== undefined && (obj.icon = message.icon);
    message.label !== undefined && (obj.label = message.label);
    message.activeMatcher !== undefined &&
      (obj.activeMatcher = message.activeMatcher);
    message.applyType !== undefined && (obj.applyType = message.applyType);
    message.capabilityId !== undefined &&
      (obj.capabilityId = message.capabilityId);
    message.capability !== undefined &&
      (obj.capability = message.capability
        ? CapabilityEntity.toJSON(message.capability)
        : undefined);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    if (message.children) {
      obj.children = message.children.map((e) =>
        e ? AppMenuEntity.toJSON(e) : undefined
      );
    } else {
      obj.children = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AppMenuEntity>, I>>(
    base?: I
  ): AppMenuEntity {
    return AppMenuEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AppMenuEntity>, I>>(
    object: I
  ): AppMenuEntity {
    const message = createBaseAppMenuEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? AppMenuEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.translations =
      object.translations?.map((e) => AppMenuEntityPolyglot.fromPartial(e)) ||
      [];
    message.href = object.href ?? undefined;
    message.icon = object.icon ?? undefined;
    message.label = object.label ?? undefined;
    message.activeMatcher = object.activeMatcher ?? undefined;
    message.applyType = object.applyType ?? undefined;
    message.capabilityId = object.capabilityId ?? undefined;
    message.capability =
      object.capability !== undefined && object.capability !== null
        ? CapabilityEntity.fromPartial(object.capability)
        : undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    message.children =
      object.children?.map((e) => AppMenuEntity.fromPartial(e)) || [];
    return message;
  },
};

function createBaseAppMenuEntityPolyglot(): AppMenuEntityPolyglot {
  return { linkerId: "", languageId: "", label: "" };
}

export const AppMenuEntityPolyglot = {
  encode(
    message: AppMenuEntityPolyglot,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.linkerId !== "") {
      writer.uint32(10).string(message.linkerId);
    }
    if (message.languageId !== "") {
      writer.uint32(18).string(message.languageId);
    }
    if (message.label !== "") {
      writer.uint32(26).string(message.label);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): AppMenuEntityPolyglot {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppMenuEntityPolyglot();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.linkerId = reader.string();
          break;
        case 2:
          message.languageId = reader.string();
          break;
        case 3:
          message.label = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppMenuEntityPolyglot {
    return {
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : "",
      languageId: isSet(object.languageId) ? String(object.languageId) : "",
      label: isSet(object.label) ? String(object.label) : "",
    };
  },

  toJSON(message: AppMenuEntityPolyglot): unknown {
    const obj: any = {};
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.languageId !== undefined && (obj.languageId = message.languageId);
    message.label !== undefined && (obj.label = message.label);
    return obj;
  },

  create<I extends Exact<DeepPartial<AppMenuEntityPolyglot>, I>>(
    base?: I
  ): AppMenuEntityPolyglot {
    return AppMenuEntityPolyglot.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AppMenuEntityPolyglot>, I>>(
    object: I
  ): AppMenuEntityPolyglot {
    const message = createBaseAppMenuEntityPolyglot();
    message.linkerId = object.linkerId ?? "";
    message.languageId = object.languageId ?? "";
    message.label = object.label ?? "";
    return message;
  },
};

function createBaseBackupReqDto(): BackupReqDto {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    entities: [],
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const BackupReqDto = {
  encode(
    message: BackupReqDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      BackupReqDto.encode(message.parent, writer.uint32(42).fork()).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    for (const v of message.entities) {
      writer.uint32(74).string(v!);
    }
    if (message.rank !== 0) {
      writer.uint32(80).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(88).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(96).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(106).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(114).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BackupReqDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBackupReqDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = BackupReqDto.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.entities.push(reader.string());
          break;
        case 10:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 11:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 12:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.createdFormatted = reader.string();
          break;
        case 14:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BackupReqDto {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? BackupReqDto.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      entities: Array.isArray(object?.entities)
        ? object.entities.map((e: any) => String(e))
        : [],
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: BackupReqDto): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? BackupReqDto.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    if (message.entities) {
      obj.entities = message.entities.map((e) => e);
    } else {
      obj.entities = [];
    }
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<BackupReqDto>, I>>(
    base?: I
  ): BackupReqDto {
    return BackupReqDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<BackupReqDto>, I>>(
    object: I
  ): BackupReqDto {
    const message = createBaseBackupReqDto();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? BackupReqDto.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.entities = object.entities?.map((e) => e) || [];
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseBackupTableMetaCreateReply(): BackupTableMetaCreateReply {
  return { data: undefined, error: undefined };
}

export const BackupTableMetaCreateReply = {
  encode(
    message: BackupTableMetaCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      BackupTableMetaEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): BackupTableMetaCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBackupTableMetaCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = BackupTableMetaEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BackupTableMetaCreateReply {
    return {
      data: isSet(object.data)
        ? BackupTableMetaEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: BackupTableMetaCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? BackupTableMetaEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<BackupTableMetaCreateReply>, I>>(
    base?: I
  ): BackupTableMetaCreateReply {
    return BackupTableMetaCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<BackupTableMetaCreateReply>, I>>(
    object: I
  ): BackupTableMetaCreateReply {
    const message = createBaseBackupTableMetaCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? BackupTableMetaEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseBackupTableMetaReply(): BackupTableMetaReply {
  return { data: undefined, error: undefined };
}

export const BackupTableMetaReply = {
  encode(
    message: BackupTableMetaReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      BackupTableMetaEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): BackupTableMetaReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBackupTableMetaReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = BackupTableMetaEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BackupTableMetaReply {
    return {
      data: isSet(object.data)
        ? BackupTableMetaEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: BackupTableMetaReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? BackupTableMetaEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<BackupTableMetaReply>, I>>(
    base?: I
  ): BackupTableMetaReply {
    return BackupTableMetaReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<BackupTableMetaReply>, I>>(
    object: I
  ): BackupTableMetaReply {
    const message = createBaseBackupTableMetaReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? BackupTableMetaEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseBackupTableMetaQueryReply(): BackupTableMetaQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const BackupTableMetaQueryReply = {
  encode(
    message: BackupTableMetaQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      BackupTableMetaEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): BackupTableMetaQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBackupTableMetaQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            BackupTableMetaEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BackupTableMetaQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => BackupTableMetaEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: BackupTableMetaQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? BackupTableMetaEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<BackupTableMetaQueryReply>, I>>(
    base?: I
  ): BackupTableMetaQueryReply {
    return BackupTableMetaQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<BackupTableMetaQueryReply>, I>>(
    object: I
  ): BackupTableMetaQueryReply {
    const message = createBaseBackupTableMetaQueryReply();
    message.items =
      object.items?.map((e) => BackupTableMetaEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseBackupTableMetaEntity(): BackupTableMetaEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    tableNameInDb: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const BackupTableMetaEntity = {
  encode(
    message: BackupTableMetaEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      BackupTableMetaEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.tableNameInDb !== undefined) {
      writer.uint32(74).string(message.tableNameInDb);
    }
    if (message.rank !== 0) {
      writer.uint32(80).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(88).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(96).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(106).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(114).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): BackupTableMetaEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBackupTableMetaEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = BackupTableMetaEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.tableNameInDb = reader.string();
          break;
        case 10:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 11:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 12:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.createdFormatted = reader.string();
          break;
        case 14:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BackupTableMetaEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? BackupTableMetaEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      tableNameInDb: isSet(object.tableNameInDb)
        ? String(object.tableNameInDb)
        : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: BackupTableMetaEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? BackupTableMetaEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.tableNameInDb !== undefined &&
      (obj.tableNameInDb = message.tableNameInDb);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<BackupTableMetaEntity>, I>>(
    base?: I
  ): BackupTableMetaEntity {
    return BackupTableMetaEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<BackupTableMetaEntity>, I>>(
    object: I
  ): BackupTableMetaEntity {
    const message = createBaseBackupTableMetaEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? BackupTableMetaEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.tableNameInDb = object.tableNameInDb ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseCapabilityCreateReply(): CapabilityCreateReply {
  return { data: undefined, error: undefined };
}

export const CapabilityCreateReply = {
  encode(
    message: CapabilityCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      CapabilityEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): CapabilityCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCapabilityCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = CapabilityEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CapabilityCreateReply {
    return {
      data: isSet(object.data)
        ? CapabilityEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: CapabilityCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? CapabilityEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CapabilityCreateReply>, I>>(
    base?: I
  ): CapabilityCreateReply {
    return CapabilityCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CapabilityCreateReply>, I>>(
    object: I
  ): CapabilityCreateReply {
    const message = createBaseCapabilityCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? CapabilityEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseCapabilityReply(): CapabilityReply {
  return { data: undefined, error: undefined };
}

export const CapabilityReply = {
  encode(
    message: CapabilityReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      CapabilityEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CapabilityReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCapabilityReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = CapabilityEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CapabilityReply {
    return {
      data: isSet(object.data)
        ? CapabilityEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: CapabilityReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? CapabilityEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CapabilityReply>, I>>(
    base?: I
  ): CapabilityReply {
    return CapabilityReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CapabilityReply>, I>>(
    object: I
  ): CapabilityReply {
    const message = createBaseCapabilityReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? CapabilityEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseCapabilityQueryReply(): CapabilityQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const CapabilityQueryReply = {
  encode(
    message: CapabilityQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      CapabilityEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): CapabilityQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCapabilityQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(CapabilityEntity.decode(reader, reader.uint32()));
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CapabilityQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => CapabilityEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: CapabilityQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? CapabilityEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CapabilityQueryReply>, I>>(
    base?: I
  ): CapabilityQueryReply {
    return CapabilityQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CapabilityQueryReply>, I>>(
    object: I
  ): CapabilityQueryReply {
    const message = createBaseCapabilityQueryReply();
    message.items =
      object.items?.map((e) => CapabilityEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseCapabilityEntity(): CapabilityEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    name: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const CapabilityEntity = {
  encode(
    message: CapabilityEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      CapabilityEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.name !== undefined) {
      writer.uint32(74).string(message.name);
    }
    if (message.rank !== 0) {
      writer.uint32(80).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(88).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(96).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(106).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(114).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CapabilityEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCapabilityEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = CapabilityEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.name = reader.string();
          break;
        case 10:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 11:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 12:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.createdFormatted = reader.string();
          break;
        case 14:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CapabilityEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? CapabilityEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      name: isSet(object.name) ? String(object.name) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: CapabilityEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? CapabilityEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.name !== undefined && (obj.name = message.name);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<CapabilityEntity>, I>>(
    base?: I
  ): CapabilityEntity {
    return CapabilityEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CapabilityEntity>, I>>(
    object: I
  ): CapabilityEntity {
    const message = createBaseCapabilityEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? CapabilityEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.name = object.name ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseCapabilityChild(): CapabilityChild {
  return {
    uniqueId: "",
    children: [],
    visibility: undefined,
    updated: 0,
    created: 0,
  };
}

export const CapabilityChild = {
  encode(
    message: CapabilityChild,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.uniqueId !== "") {
      writer.uint32(10).string(message.uniqueId);
    }
    for (const v of message.children) {
      CapabilityChild.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.visibility !== undefined) {
      writer.uint32(26).string(message.visibility);
    }
    if (message.updated !== 0) {
      writer.uint32(32).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(40).int64(message.created);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CapabilityChild {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCapabilityChild();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uniqueId = reader.string();
          break;
        case 2:
          message.children.push(
            CapabilityChild.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.visibility = reader.string();
          break;
        case 4:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.created = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CapabilityChild {
    return {
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      children: Array.isArray(object?.children)
        ? object.children.map((e: any) => CapabilityChild.fromJSON(e))
        : [],
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
    };
  },

  toJSON(message: CapabilityChild): unknown {
    const obj: any = {};
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    if (message.children) {
      obj.children = message.children.map((e) =>
        e ? CapabilityChild.toJSON(e) : undefined
      );
    } else {
      obj.children = [];
    }
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    return obj;
  },

  create<I extends Exact<DeepPartial<CapabilityChild>, I>>(
    base?: I
  ): CapabilityChild {
    return CapabilityChild.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CapabilityChild>, I>>(
    object: I
  ): CapabilityChild {
    const message = createBaseCapabilityChild();
    message.uniqueId = object.uniqueId ?? "";
    message.children =
      object.children?.map((e) => CapabilityChild.fromPartial(e)) || [];
    message.visibility = object.visibility ?? undefined;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    return message;
  },
};

function createBaseCapabilitiesResult(): CapabilitiesResult {
  return {
    capabilities: [],
    nested: [],
    visibility: undefined,
    updated: 0,
    created: 0,
  };
}

export const CapabilitiesResult = {
  encode(
    message: CapabilitiesResult,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.capabilities) {
      CapabilityEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.nested) {
      CapabilityChild.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.visibility !== undefined) {
      writer.uint32(26).string(message.visibility);
    }
    if (message.updated !== 0) {
      writer.uint32(32).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(40).int64(message.created);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CapabilitiesResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCapabilitiesResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.capabilities.push(
            CapabilityEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.nested.push(CapabilityChild.decode(reader, reader.uint32()));
          break;
        case 3:
          message.visibility = reader.string();
          break;
        case 4:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.created = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CapabilitiesResult {
    return {
      capabilities: Array.isArray(object?.capabilities)
        ? object.capabilities.map((e: any) => CapabilityEntity.fromJSON(e))
        : [],
      nested: Array.isArray(object?.nested)
        ? object.nested.map((e: any) => CapabilityChild.fromJSON(e))
        : [],
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
    };
  },

  toJSON(message: CapabilitiesResult): unknown {
    const obj: any = {};
    if (message.capabilities) {
      obj.capabilities = message.capabilities.map((e) =>
        e ? CapabilityEntity.toJSON(e) : undefined
      );
    } else {
      obj.capabilities = [];
    }
    if (message.nested) {
      obj.nested = message.nested.map((e) =>
        e ? CapabilityChild.toJSON(e) : undefined
      );
    } else {
      obj.nested = [];
    }
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    return obj;
  },

  create<I extends Exact<DeepPartial<CapabilitiesResult>, I>>(
    base?: I
  ): CapabilitiesResult {
    return CapabilitiesResult.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CapabilitiesResult>, I>>(
    object: I
  ): CapabilitiesResult {
    const message = createBaseCapabilitiesResult();
    message.capabilities =
      object.capabilities?.map((e) => CapabilityEntity.fromPartial(e)) || [];
    message.nested =
      object.nested?.map((e) => CapabilityChild.fromPartial(e)) || [];
    message.visibility = object.visibility ?? undefined;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    return message;
  },
};

function createBaseEmailConfirmationCreateReply(): EmailConfirmationCreateReply {
  return { data: undefined, error: undefined };
}

export const EmailConfirmationCreateReply = {
  encode(
    message: EmailConfirmationCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      EmailConfirmationEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): EmailConfirmationCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailConfirmationCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = EmailConfirmationEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailConfirmationCreateReply {
    return {
      data: isSet(object.data)
        ? EmailConfirmationEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: EmailConfirmationCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? EmailConfirmationEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailConfirmationCreateReply>, I>>(
    base?: I
  ): EmailConfirmationCreateReply {
    return EmailConfirmationCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailConfirmationCreateReply>, I>>(
    object: I
  ): EmailConfirmationCreateReply {
    const message = createBaseEmailConfirmationCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? EmailConfirmationEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseEmailConfirmationReply(): EmailConfirmationReply {
  return { data: undefined, error: undefined };
}

export const EmailConfirmationReply = {
  encode(
    message: EmailConfirmationReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      EmailConfirmationEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): EmailConfirmationReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailConfirmationReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = EmailConfirmationEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailConfirmationReply {
    return {
      data: isSet(object.data)
        ? EmailConfirmationEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: EmailConfirmationReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? EmailConfirmationEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailConfirmationReply>, I>>(
    base?: I
  ): EmailConfirmationReply {
    return EmailConfirmationReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailConfirmationReply>, I>>(
    object: I
  ): EmailConfirmationReply {
    const message = createBaseEmailConfirmationReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? EmailConfirmationEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseEmailConfirmationQueryReply(): EmailConfirmationQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const EmailConfirmationQueryReply = {
  encode(
    message: EmailConfirmationQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      EmailConfirmationEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): EmailConfirmationQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailConfirmationQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            EmailConfirmationEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailConfirmationQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => EmailConfirmationEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: EmailConfirmationQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? EmailConfirmationEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailConfirmationQueryReply>, I>>(
    base?: I
  ): EmailConfirmationQueryReply {
    return EmailConfirmationQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailConfirmationQueryReply>, I>>(
    object: I
  ): EmailConfirmationQueryReply {
    const message = createBaseEmailConfirmationQueryReply();
    message.items =
      object.items?.map((e) => EmailConfirmationEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseEmailConfirmationEntity(): EmailConfirmationEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    user: undefined,
    status: undefined,
    email: undefined,
    key: undefined,
    expiresAt: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const EmailConfirmationEntity = {
  encode(
    message: EmailConfirmationEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      EmailConfirmationEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.user !== undefined) {
      UserEntity.encode(message.user, writer.uint32(82).fork()).ldelim();
    }
    if (message.status !== undefined) {
      writer.uint32(90).string(message.status);
    }
    if (message.email !== undefined) {
      writer.uint32(98).string(message.email);
    }
    if (message.key !== undefined) {
      writer.uint32(106).string(message.key);
    }
    if (message.expiresAt !== undefined) {
      writer.uint32(114).string(message.expiresAt);
    }
    if (message.rank !== 0) {
      writer.uint32(120).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(128).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(136).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(146).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(154).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): EmailConfirmationEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailConfirmationEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = EmailConfirmationEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 10:
          message.user = UserEntity.decode(reader, reader.uint32());
          break;
        case 11:
          message.status = reader.string();
          break;
        case 12:
          message.email = reader.string();
          break;
        case 13:
          message.key = reader.string();
          break;
        case 14:
          message.expiresAt = reader.string();
          break;
        case 15:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 18:
          message.createdFormatted = reader.string();
          break;
        case 19:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailConfirmationEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? EmailConfirmationEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      user: isSet(object.user) ? UserEntity.fromJSON(object.user) : undefined,
      status: isSet(object.status) ? String(object.status) : undefined,
      email: isSet(object.email) ? String(object.email) : undefined,
      key: isSet(object.key) ? String(object.key) : undefined,
      expiresAt: isSet(object.expiresAt) ? String(object.expiresAt) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: EmailConfirmationEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? EmailConfirmationEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.user !== undefined &&
      (obj.user = message.user ? UserEntity.toJSON(message.user) : undefined);
    message.status !== undefined && (obj.status = message.status);
    message.email !== undefined && (obj.email = message.email);
    message.key !== undefined && (obj.key = message.key);
    message.expiresAt !== undefined && (obj.expiresAt = message.expiresAt);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailConfirmationEntity>, I>>(
    base?: I
  ): EmailConfirmationEntity {
    return EmailConfirmationEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailConfirmationEntity>, I>>(
    object: I
  ): EmailConfirmationEntity {
    const message = createBaseEmailConfirmationEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? EmailConfirmationEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.user =
      object.user !== undefined && object.user !== null
        ? UserEntity.fromPartial(object.user)
        : undefined;
    message.status = object.status ?? undefined;
    message.email = object.email ?? undefined;
    message.key = object.key ?? undefined;
    message.expiresAt = object.expiresAt ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseEmailProviderCreateReply(): EmailProviderCreateReply {
  return { data: undefined, error: undefined };
}

export const EmailProviderCreateReply = {
  encode(
    message: EmailProviderCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      EmailProviderEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): EmailProviderCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailProviderCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = EmailProviderEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailProviderCreateReply {
    return {
      data: isSet(object.data)
        ? EmailProviderEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: EmailProviderCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? EmailProviderEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailProviderCreateReply>, I>>(
    base?: I
  ): EmailProviderCreateReply {
    return EmailProviderCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailProviderCreateReply>, I>>(
    object: I
  ): EmailProviderCreateReply {
    const message = createBaseEmailProviderCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? EmailProviderEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseEmailProviderReply(): EmailProviderReply {
  return { data: undefined, error: undefined };
}

export const EmailProviderReply = {
  encode(
    message: EmailProviderReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      EmailProviderEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EmailProviderReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailProviderReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = EmailProviderEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailProviderReply {
    return {
      data: isSet(object.data)
        ? EmailProviderEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: EmailProviderReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? EmailProviderEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailProviderReply>, I>>(
    base?: I
  ): EmailProviderReply {
    return EmailProviderReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailProviderReply>, I>>(
    object: I
  ): EmailProviderReply {
    const message = createBaseEmailProviderReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? EmailProviderEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseEmailProviderQueryReply(): EmailProviderQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const EmailProviderQueryReply = {
  encode(
    message: EmailProviderQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      EmailProviderEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): EmailProviderQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailProviderQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            EmailProviderEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailProviderQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => EmailProviderEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: EmailProviderQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? EmailProviderEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailProviderQueryReply>, I>>(
    base?: I
  ): EmailProviderQueryReply {
    return EmailProviderQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailProviderQueryReply>, I>>(
    object: I
  ): EmailProviderQueryReply {
    const message = createBaseEmailProviderQueryReply();
    message.items =
      object.items?.map((e) => EmailProviderEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseEmailProviderEntity(): EmailProviderEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    type: undefined,
    apiKey: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const EmailProviderEntity = {
  encode(
    message: EmailProviderEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      EmailProviderEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.type !== undefined) {
      writer.uint32(74).string(message.type);
    }
    if (message.apiKey !== undefined) {
      writer.uint32(82).string(message.apiKey);
    }
    if (message.rank !== 0) {
      writer.uint32(88).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(96).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(104).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(114).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(122).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EmailProviderEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailProviderEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = EmailProviderEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.type = reader.string();
          break;
        case 10:
          message.apiKey = reader.string();
          break;
        case 11:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 12:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.createdFormatted = reader.string();
          break;
        case 15:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailProviderEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? EmailProviderEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      type: isSet(object.type) ? String(object.type) : undefined,
      apiKey: isSet(object.apiKey) ? String(object.apiKey) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: EmailProviderEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? EmailProviderEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.type !== undefined && (obj.type = message.type);
    message.apiKey !== undefined && (obj.apiKey = message.apiKey);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailProviderEntity>, I>>(
    base?: I
  ): EmailProviderEntity {
    return EmailProviderEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailProviderEntity>, I>>(
    object: I
  ): EmailProviderEntity {
    const message = createBaseEmailProviderEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? EmailProviderEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.type = object.type ?? undefined;
    message.apiKey = object.apiKey ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseEmailSenderCreateReply(): EmailSenderCreateReply {
  return { data: undefined, error: undefined };
}

export const EmailSenderCreateReply = {
  encode(
    message: EmailSenderCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      EmailSenderEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): EmailSenderCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailSenderCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = EmailSenderEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailSenderCreateReply {
    return {
      data: isSet(object.data)
        ? EmailSenderEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: EmailSenderCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? EmailSenderEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailSenderCreateReply>, I>>(
    base?: I
  ): EmailSenderCreateReply {
    return EmailSenderCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailSenderCreateReply>, I>>(
    object: I
  ): EmailSenderCreateReply {
    const message = createBaseEmailSenderCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? EmailSenderEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseEmailSenderReply(): EmailSenderReply {
  return { data: undefined, error: undefined };
}

export const EmailSenderReply = {
  encode(
    message: EmailSenderReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      EmailSenderEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EmailSenderReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailSenderReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = EmailSenderEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailSenderReply {
    return {
      data: isSet(object.data)
        ? EmailSenderEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: EmailSenderReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? EmailSenderEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailSenderReply>, I>>(
    base?: I
  ): EmailSenderReply {
    return EmailSenderReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailSenderReply>, I>>(
    object: I
  ): EmailSenderReply {
    const message = createBaseEmailSenderReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? EmailSenderEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseEmailSenderQueryReply(): EmailSenderQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const EmailSenderQueryReply = {
  encode(
    message: EmailSenderQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      EmailSenderEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): EmailSenderQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailSenderQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(EmailSenderEntity.decode(reader, reader.uint32()));
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailSenderQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => EmailSenderEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: EmailSenderQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? EmailSenderEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailSenderQueryReply>, I>>(
    base?: I
  ): EmailSenderQueryReply {
    return EmailSenderQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailSenderQueryReply>, I>>(
    object: I
  ): EmailSenderQueryReply {
    const message = createBaseEmailSenderQueryReply();
    message.items =
      object.items?.map((e) => EmailSenderEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseEmailSenderEntity(): EmailSenderEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    fromName: undefined,
    fromEmailAddress: undefined,
    replyTo: undefined,
    nickName: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const EmailSenderEntity = {
  encode(
    message: EmailSenderEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      EmailSenderEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.fromName !== undefined) {
      writer.uint32(74).string(message.fromName);
    }
    if (message.fromEmailAddress !== undefined) {
      writer.uint32(82).string(message.fromEmailAddress);
    }
    if (message.replyTo !== undefined) {
      writer.uint32(90).string(message.replyTo);
    }
    if (message.nickName !== undefined) {
      writer.uint32(98).string(message.nickName);
    }
    if (message.rank !== 0) {
      writer.uint32(104).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(112).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(120).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(130).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(138).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EmailSenderEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailSenderEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = EmailSenderEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.fromName = reader.string();
          break;
        case 10:
          message.fromEmailAddress = reader.string();
          break;
        case 11:
          message.replyTo = reader.string();
          break;
        case 12:
          message.nickName = reader.string();
          break;
        case 13:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.createdFormatted = reader.string();
          break;
        case 17:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailSenderEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? EmailSenderEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      fromName: isSet(object.fromName) ? String(object.fromName) : undefined,
      fromEmailAddress: isSet(object.fromEmailAddress)
        ? String(object.fromEmailAddress)
        : undefined,
      replyTo: isSet(object.replyTo) ? String(object.replyTo) : undefined,
      nickName: isSet(object.nickName) ? String(object.nickName) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: EmailSenderEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? EmailSenderEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.fromName !== undefined && (obj.fromName = message.fromName);
    message.fromEmailAddress !== undefined &&
      (obj.fromEmailAddress = message.fromEmailAddress);
    message.replyTo !== undefined && (obj.replyTo = message.replyTo);
    message.nickName !== undefined && (obj.nickName = message.nickName);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailSenderEntity>, I>>(
    base?: I
  ): EmailSenderEntity {
    return EmailSenderEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailSenderEntity>, I>>(
    object: I
  ): EmailSenderEntity {
    const message = createBaseEmailSenderEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? EmailSenderEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.fromName = object.fromName ?? undefined;
    message.fromEmailAddress = object.fromEmailAddress ?? undefined;
    message.replyTo = object.replyTo ?? undefined;
    message.nickName = object.nickName ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseTestMailDto(): TestMailDto {
  return { senderId: "", toName: "", toEmail: "", subject: "", content: "" };
}

export const TestMailDto = {
  encode(
    message: TestMailDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.senderId !== "") {
      writer.uint32(10).string(message.senderId);
    }
    if (message.toName !== "") {
      writer.uint32(18).string(message.toName);
    }
    if (message.toEmail !== "") {
      writer.uint32(26).string(message.toEmail);
    }
    if (message.subject !== "") {
      writer.uint32(34).string(message.subject);
    }
    if (message.content !== "") {
      writer.uint32(434).string(message.content);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TestMailDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTestMailDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.senderId = reader.string();
          break;
        case 2:
          message.toName = reader.string();
          break;
        case 3:
          message.toEmail = reader.string();
          break;
        case 4:
          message.subject = reader.string();
          break;
        case 54:
          message.content = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TestMailDto {
    return {
      senderId: isSet(object.senderId) ? String(object.senderId) : "",
      toName: isSet(object.toName) ? String(object.toName) : "",
      toEmail: isSet(object.toEmail) ? String(object.toEmail) : "",
      subject: isSet(object.subject) ? String(object.subject) : "",
      content: isSet(object.content) ? String(object.content) : "",
    };
  },

  toJSON(message: TestMailDto): unknown {
    const obj: any = {};
    message.senderId !== undefined && (obj.senderId = message.senderId);
    message.toName !== undefined && (obj.toName = message.toName);
    message.toEmail !== undefined && (obj.toEmail = message.toEmail);
    message.subject !== undefined && (obj.subject = message.subject);
    message.content !== undefined && (obj.content = message.content);
    return obj;
  },

  create<I extends Exact<DeepPartial<TestMailDto>, I>>(base?: I): TestMailDto {
    return TestMailDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TestMailDto>, I>>(
    object: I
  ): TestMailDto {
    const message = createBaseTestMailDto();
    message.senderId = object.senderId ?? "";
    message.toName = object.toName ?? "";
    message.toEmail = object.toEmail ?? "";
    message.subject = object.subject ?? "";
    message.content = object.content ?? "";
    return message;
  },
};

function createBaseForgetPasswordCreateReply(): ForgetPasswordCreateReply {
  return { data: undefined, error: undefined };
}

export const ForgetPasswordCreateReply = {
  encode(
    message: ForgetPasswordCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      ForgetPasswordEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ForgetPasswordCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseForgetPasswordCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = ForgetPasswordEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ForgetPasswordCreateReply {
    return {
      data: isSet(object.data)
        ? ForgetPasswordEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: ForgetPasswordCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? ForgetPasswordEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ForgetPasswordCreateReply>, I>>(
    base?: I
  ): ForgetPasswordCreateReply {
    return ForgetPasswordCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ForgetPasswordCreateReply>, I>>(
    object: I
  ): ForgetPasswordCreateReply {
    const message = createBaseForgetPasswordCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? ForgetPasswordEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseForgetPasswordReply(): ForgetPasswordReply {
  return { data: undefined, error: undefined };
}

export const ForgetPasswordReply = {
  encode(
    message: ForgetPasswordReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      ForgetPasswordEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ForgetPasswordReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseForgetPasswordReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = ForgetPasswordEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ForgetPasswordReply {
    return {
      data: isSet(object.data)
        ? ForgetPasswordEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: ForgetPasswordReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? ForgetPasswordEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ForgetPasswordReply>, I>>(
    base?: I
  ): ForgetPasswordReply {
    return ForgetPasswordReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ForgetPasswordReply>, I>>(
    object: I
  ): ForgetPasswordReply {
    const message = createBaseForgetPasswordReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? ForgetPasswordEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseForgetPasswordQueryReply(): ForgetPasswordQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const ForgetPasswordQueryReply = {
  encode(
    message: ForgetPasswordQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      ForgetPasswordEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ForgetPasswordQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseForgetPasswordQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            ForgetPasswordEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ForgetPasswordQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => ForgetPasswordEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: ForgetPasswordQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? ForgetPasswordEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<ForgetPasswordQueryReply>, I>>(
    base?: I
  ): ForgetPasswordQueryReply {
    return ForgetPasswordQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ForgetPasswordQueryReply>, I>>(
    object: I
  ): ForgetPasswordQueryReply {
    const message = createBaseForgetPasswordQueryReply();
    message.items =
      object.items?.map((e) => ForgetPasswordEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseForgetPasswordEntity(): ForgetPasswordEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    user: undefined,
    passportId: undefined,
    passport: undefined,
    status: undefined,
    validUntil: 0,
    validUntilFormatted: "",
    blockedUntil: 0,
    blockedUntilFormatted: "",
    secondsToUnblock: undefined,
    otp: undefined,
    recoveryAbsoluteUrl: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const ForgetPasswordEntity = {
  encode(
    message: ForgetPasswordEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      ForgetPasswordEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.user !== undefined) {
      UserEntity.encode(message.user, writer.uint32(82).fork()).ldelim();
    }
    if (message.passportId !== undefined) {
      writer.uint32(98).string(message.passportId);
    }
    if (message.passport !== undefined) {
      PassportEntity.encode(
        message.passport,
        writer.uint32(106).fork()
      ).ldelim();
    }
    if (message.status !== undefined) {
      writer.uint32(114).string(message.status);
    }
    if (message.validUntil !== 0) {
      writer.uint32(128).int64(message.validUntil);
    }
    if (message.validUntilFormatted !== "") {
      writer.uint32(138).string(message.validUntilFormatted);
    }
    if (message.blockedUntil !== 0) {
      writer.uint32(152).int64(message.blockedUntil);
    }
    if (message.blockedUntilFormatted !== "") {
      writer.uint32(162).string(message.blockedUntilFormatted);
    }
    if (message.secondsToUnblock !== undefined) {
      writer.uint32(168).int64(message.secondsToUnblock);
    }
    if (message.otp !== undefined) {
      writer.uint32(178).string(message.otp);
    }
    if (message.recoveryAbsoluteUrl !== undefined) {
      writer.uint32(186).string(message.recoveryAbsoluteUrl);
    }
    if (message.rank !== 0) {
      writer.uint32(192).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(200).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(208).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(218).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(226).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ForgetPasswordEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseForgetPasswordEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = ForgetPasswordEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 10:
          message.user = UserEntity.decode(reader, reader.uint32());
          break;
        case 12:
          message.passportId = reader.string();
          break;
        case 13:
          message.passport = PassportEntity.decode(reader, reader.uint32());
          break;
        case 14:
          message.status = reader.string();
          break;
        case 16:
          message.validUntil = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.validUntilFormatted = reader.string();
          break;
        case 19:
          message.blockedUntil = longToNumber(reader.int64() as Long);
          break;
        case 20:
          message.blockedUntilFormatted = reader.string();
          break;
        case 21:
          message.secondsToUnblock = longToNumber(reader.int64() as Long);
          break;
        case 22:
          message.otp = reader.string();
          break;
        case 23:
          message.recoveryAbsoluteUrl = reader.string();
          break;
        case 24:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 25:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 26:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 27:
          message.createdFormatted = reader.string();
          break;
        case 28:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ForgetPasswordEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? ForgetPasswordEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      user: isSet(object.user) ? UserEntity.fromJSON(object.user) : undefined,
      passportId: isSet(object.passportId)
        ? String(object.passportId)
        : undefined,
      passport: isSet(object.passport)
        ? PassportEntity.fromJSON(object.passport)
        : undefined,
      status: isSet(object.status) ? String(object.status) : undefined,
      validUntil: isSet(object.validUntil) ? Number(object.validUntil) : 0,
      validUntilFormatted: isSet(object.validUntilFormatted)
        ? String(object.validUntilFormatted)
        : "",
      blockedUntil: isSet(object.blockedUntil)
        ? Number(object.blockedUntil)
        : 0,
      blockedUntilFormatted: isSet(object.blockedUntilFormatted)
        ? String(object.blockedUntilFormatted)
        : "",
      secondsToUnblock: isSet(object.secondsToUnblock)
        ? Number(object.secondsToUnblock)
        : undefined,
      otp: isSet(object.otp) ? String(object.otp) : undefined,
      recoveryAbsoluteUrl: isSet(object.recoveryAbsoluteUrl)
        ? String(object.recoveryAbsoluteUrl)
        : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: ForgetPasswordEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? ForgetPasswordEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.user !== undefined &&
      (obj.user = message.user ? UserEntity.toJSON(message.user) : undefined);
    message.passportId !== undefined && (obj.passportId = message.passportId);
    message.passport !== undefined &&
      (obj.passport = message.passport
        ? PassportEntity.toJSON(message.passport)
        : undefined);
    message.status !== undefined && (obj.status = message.status);
    message.validUntil !== undefined &&
      (obj.validUntil = Math.round(message.validUntil));
    message.validUntilFormatted !== undefined &&
      (obj.validUntilFormatted = message.validUntilFormatted);
    message.blockedUntil !== undefined &&
      (obj.blockedUntil = Math.round(message.blockedUntil));
    message.blockedUntilFormatted !== undefined &&
      (obj.blockedUntilFormatted = message.blockedUntilFormatted);
    message.secondsToUnblock !== undefined &&
      (obj.secondsToUnblock = Math.round(message.secondsToUnblock));
    message.otp !== undefined && (obj.otp = message.otp);
    message.recoveryAbsoluteUrl !== undefined &&
      (obj.recoveryAbsoluteUrl = message.recoveryAbsoluteUrl);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<ForgetPasswordEntity>, I>>(
    base?: I
  ): ForgetPasswordEntity {
    return ForgetPasswordEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ForgetPasswordEntity>, I>>(
    object: I
  ): ForgetPasswordEntity {
    const message = createBaseForgetPasswordEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? ForgetPasswordEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.user =
      object.user !== undefined && object.user !== null
        ? UserEntity.fromPartial(object.user)
        : undefined;
    message.passportId = object.passportId ?? undefined;
    message.passport =
      object.passport !== undefined && object.passport !== null
        ? PassportEntity.fromPartial(object.passport)
        : undefined;
    message.status = object.status ?? undefined;
    message.validUntil = object.validUntil ?? 0;
    message.validUntilFormatted = object.validUntilFormatted ?? "";
    message.blockedUntil = object.blockedUntil ?? 0;
    message.blockedUntilFormatted = object.blockedUntilFormatted ?? "";
    message.secondsToUnblock = object.secondsToUnblock ?? undefined;
    message.otp = object.otp ?? undefined;
    message.recoveryAbsoluteUrl = object.recoveryAbsoluteUrl ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseGsmProviderCreateReply(): GsmProviderCreateReply {
  return { data: undefined, error: undefined };
}

export const GsmProviderCreateReply = {
  encode(
    message: GsmProviderCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      GsmProviderEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): GsmProviderCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGsmProviderCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = GsmProviderEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GsmProviderCreateReply {
    return {
      data: isSet(object.data)
        ? GsmProviderEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: GsmProviderCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? GsmProviderEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GsmProviderCreateReply>, I>>(
    base?: I
  ): GsmProviderCreateReply {
    return GsmProviderCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GsmProviderCreateReply>, I>>(
    object: I
  ): GsmProviderCreateReply {
    const message = createBaseGsmProviderCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? GsmProviderEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseGsmProviderReply(): GsmProviderReply {
  return { data: undefined, error: undefined };
}

export const GsmProviderReply = {
  encode(
    message: GsmProviderReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      GsmProviderEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GsmProviderReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGsmProviderReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = GsmProviderEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GsmProviderReply {
    return {
      data: isSet(object.data)
        ? GsmProviderEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: GsmProviderReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? GsmProviderEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GsmProviderReply>, I>>(
    base?: I
  ): GsmProviderReply {
    return GsmProviderReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GsmProviderReply>, I>>(
    object: I
  ): GsmProviderReply {
    const message = createBaseGsmProviderReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? GsmProviderEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseGsmProviderQueryReply(): GsmProviderQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const GsmProviderQueryReply = {
  encode(
    message: GsmProviderQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      GsmProviderEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): GsmProviderQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGsmProviderQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(GsmProviderEntity.decode(reader, reader.uint32()));
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GsmProviderQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => GsmProviderEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: GsmProviderQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? GsmProviderEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<GsmProviderQueryReply>, I>>(
    base?: I
  ): GsmProviderQueryReply {
    return GsmProviderQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GsmProviderQueryReply>, I>>(
    object: I
  ): GsmProviderQueryReply {
    const message = createBaseGsmProviderQueryReply();
    message.items =
      object.items?.map((e) => GsmProviderEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseGsmProviderEntity(): GsmProviderEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    apiKey: undefined,
    mainSenderNumber: undefined,
    type: undefined,
    invokeUrl: undefined,
    invokeBody: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const GsmProviderEntity = {
  encode(
    message: GsmProviderEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      GsmProviderEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.apiKey !== undefined) {
      writer.uint32(74).string(message.apiKey);
    }
    if (message.mainSenderNumber !== undefined) {
      writer.uint32(82).string(message.mainSenderNumber);
    }
    if (message.type !== undefined) {
      writer.uint32(90).string(message.type);
    }
    if (message.invokeUrl !== undefined) {
      writer.uint32(98).string(message.invokeUrl);
    }
    if (message.invokeBody !== undefined) {
      writer.uint32(106).string(message.invokeBody);
    }
    if (message.rank !== 0) {
      writer.uint32(112).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(120).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(128).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(138).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(146).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GsmProviderEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGsmProviderEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = GsmProviderEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.apiKey = reader.string();
          break;
        case 10:
          message.mainSenderNumber = reader.string();
          break;
        case 11:
          message.type = reader.string();
          break;
        case 12:
          message.invokeUrl = reader.string();
          break;
        case 13:
          message.invokeBody = reader.string();
          break;
        case 14:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.createdFormatted = reader.string();
          break;
        case 18:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GsmProviderEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? GsmProviderEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      apiKey: isSet(object.apiKey) ? String(object.apiKey) : undefined,
      mainSenderNumber: isSet(object.mainSenderNumber)
        ? String(object.mainSenderNumber)
        : undefined,
      type: isSet(object.type) ? String(object.type) : undefined,
      invokeUrl: isSet(object.invokeUrl) ? String(object.invokeUrl) : undefined,
      invokeBody: isSet(object.invokeBody)
        ? String(object.invokeBody)
        : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: GsmProviderEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? GsmProviderEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.apiKey !== undefined && (obj.apiKey = message.apiKey);
    message.mainSenderNumber !== undefined &&
      (obj.mainSenderNumber = message.mainSenderNumber);
    message.type !== undefined && (obj.type = message.type);
    message.invokeUrl !== undefined && (obj.invokeUrl = message.invokeUrl);
    message.invokeBody !== undefined && (obj.invokeBody = message.invokeBody);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<GsmProviderEntity>, I>>(
    base?: I
  ): GsmProviderEntity {
    return GsmProviderEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GsmProviderEntity>, I>>(
    object: I
  ): GsmProviderEntity {
    const message = createBaseGsmProviderEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? GsmProviderEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.apiKey = object.apiKey ?? undefined;
    message.mainSenderNumber = object.mainSenderNumber ?? undefined;
    message.type = object.type ?? undefined;
    message.invokeUrl = object.invokeUrl ?? undefined;
    message.invokeBody = object.invokeBody ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseNotificationConfigCreateReply(): NotificationConfigCreateReply {
  return { data: undefined, error: undefined };
}

export const NotificationConfigCreateReply = {
  encode(
    message: NotificationConfigCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      NotificationConfigEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): NotificationConfigCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNotificationConfigCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = NotificationConfigEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NotificationConfigCreateReply {
    return {
      data: isSet(object.data)
        ? NotificationConfigEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: NotificationConfigCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? NotificationConfigEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<NotificationConfigCreateReply>, I>>(
    base?: I
  ): NotificationConfigCreateReply {
    return NotificationConfigCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<NotificationConfigCreateReply>, I>>(
    object: I
  ): NotificationConfigCreateReply {
    const message = createBaseNotificationConfigCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? NotificationConfigEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseNotificationConfigReply(): NotificationConfigReply {
  return { data: undefined, error: undefined };
}

export const NotificationConfigReply = {
  encode(
    message: NotificationConfigReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      NotificationConfigEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): NotificationConfigReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNotificationConfigReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = NotificationConfigEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NotificationConfigReply {
    return {
      data: isSet(object.data)
        ? NotificationConfigEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: NotificationConfigReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? NotificationConfigEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<NotificationConfigReply>, I>>(
    base?: I
  ): NotificationConfigReply {
    return NotificationConfigReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<NotificationConfigReply>, I>>(
    object: I
  ): NotificationConfigReply {
    const message = createBaseNotificationConfigReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? NotificationConfigEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseNotificationConfigQueryReply(): NotificationConfigQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const NotificationConfigQueryReply = {
  encode(
    message: NotificationConfigQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      NotificationConfigEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): NotificationConfigQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNotificationConfigQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            NotificationConfigEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NotificationConfigQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => NotificationConfigEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: NotificationConfigQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? NotificationConfigEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<NotificationConfigQueryReply>, I>>(
    base?: I
  ): NotificationConfigQueryReply {
    return NotificationConfigQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<NotificationConfigQueryReply>, I>>(
    object: I
  ): NotificationConfigQueryReply {
    const message = createBaseNotificationConfigQueryReply();
    message.items =
      object.items?.map((e) => NotificationConfigEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseNotificationConfigEntity(): NotificationConfigEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    cascadeToSubWorkspaces: undefined,
    forcedCascadeEmailProvider: undefined,
    generalEmailProviderId: undefined,
    generalEmailProvider: undefined,
    generalGsmProviderId: undefined,
    generalGsmProvider: undefined,
    inviteToWorkspaceContent: undefined,
    inviteToWorkspaceContentExcerpt: undefined,
    inviteToWorkspaceContentDefault: undefined,
    inviteToWorkspaceContentDefaultExcerpt: undefined,
    inviteToWorkspaceTitle: undefined,
    inviteToWorkspaceTitleDefault: undefined,
    inviteToWorkspaceSenderId: undefined,
    inviteToWorkspaceSender: undefined,
    forgetPasswordContent: undefined,
    forgetPasswordContentExcerpt: undefined,
    forgetPasswordContentDefault: undefined,
    forgetPasswordContentDefaultExcerpt: undefined,
    forgetPasswordTitle: undefined,
    forgetPasswordTitleDefault: undefined,
    forgetPasswordSenderId: undefined,
    forgetPasswordSender: undefined,
    acceptLanguage: undefined,
    acceptLanguageExcerpt: undefined,
    confirmEmailSenderId: undefined,
    confirmEmailSender: undefined,
    confirmEmailContent: undefined,
    confirmEmailContentExcerpt: undefined,
    confirmEmailContentDefault: undefined,
    confirmEmailContentDefaultExcerpt: undefined,
    confirmEmailTitle: undefined,
    confirmEmailTitleDefault: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const NotificationConfigEntity = {
  encode(
    message: NotificationConfigEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      NotificationConfigEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.cascadeToSubWorkspaces !== undefined) {
      writer.uint32(72).bool(message.cascadeToSubWorkspaces);
    }
    if (message.forcedCascadeEmailProvider !== undefined) {
      writer.uint32(80).bool(message.forcedCascadeEmailProvider);
    }
    if (message.generalEmailProviderId !== undefined) {
      writer.uint32(98).string(message.generalEmailProviderId);
    }
    if (message.generalEmailProvider !== undefined) {
      EmailProviderEntity.encode(
        message.generalEmailProvider,
        writer.uint32(106).fork()
      ).ldelim();
    }
    if (message.generalGsmProviderId !== undefined) {
      writer.uint32(122).string(message.generalGsmProviderId);
    }
    if (message.generalGsmProvider !== undefined) {
      GsmProviderEntity.encode(
        message.generalGsmProvider,
        writer.uint32(130).fork()
      ).ldelim();
    }
    if (message.inviteToWorkspaceContent !== undefined) {
      writer.uint32(138).string(message.inviteToWorkspaceContent);
    }
    if (message.inviteToWorkspaceContentExcerpt !== undefined) {
      writer.uint32(146).string(message.inviteToWorkspaceContentExcerpt);
    }
    if (message.inviteToWorkspaceContentDefault !== undefined) {
      writer.uint32(154).string(message.inviteToWorkspaceContentDefault);
    }
    if (message.inviteToWorkspaceContentDefaultExcerpt !== undefined) {
      writer.uint32(162).string(message.inviteToWorkspaceContentDefaultExcerpt);
    }
    if (message.inviteToWorkspaceTitle !== undefined) {
      writer.uint32(170).string(message.inviteToWorkspaceTitle);
    }
    if (message.inviteToWorkspaceTitleDefault !== undefined) {
      writer.uint32(178).string(message.inviteToWorkspaceTitleDefault);
    }
    if (message.inviteToWorkspaceSenderId !== undefined) {
      writer.uint32(194).string(message.inviteToWorkspaceSenderId);
    }
    if (message.inviteToWorkspaceSender !== undefined) {
      EmailSenderEntity.encode(
        message.inviteToWorkspaceSender,
        writer.uint32(202).fork()
      ).ldelim();
    }
    if (message.forgetPasswordContent !== undefined) {
      writer.uint32(210).string(message.forgetPasswordContent);
    }
    if (message.forgetPasswordContentExcerpt !== undefined) {
      writer.uint32(218).string(message.forgetPasswordContentExcerpt);
    }
    if (message.forgetPasswordContentDefault !== undefined) {
      writer.uint32(226).string(message.forgetPasswordContentDefault);
    }
    if (message.forgetPasswordContentDefaultExcerpt !== undefined) {
      writer.uint32(234).string(message.forgetPasswordContentDefaultExcerpt);
    }
    if (message.forgetPasswordTitle !== undefined) {
      writer.uint32(242).string(message.forgetPasswordTitle);
    }
    if (message.forgetPasswordTitleDefault !== undefined) {
      writer.uint32(250).string(message.forgetPasswordTitleDefault);
    }
    if (message.forgetPasswordSenderId !== undefined) {
      writer.uint32(266).string(message.forgetPasswordSenderId);
    }
    if (message.forgetPasswordSender !== undefined) {
      EmailSenderEntity.encode(
        message.forgetPasswordSender,
        writer.uint32(274).fork()
      ).ldelim();
    }
    if (message.acceptLanguage !== undefined) {
      writer.uint32(282).string(message.acceptLanguage);
    }
    if (message.acceptLanguageExcerpt !== undefined) {
      writer.uint32(290).string(message.acceptLanguageExcerpt);
    }
    if (message.confirmEmailSenderId !== undefined) {
      writer.uint32(306).string(message.confirmEmailSenderId);
    }
    if (message.confirmEmailSender !== undefined) {
      EmailSenderEntity.encode(
        message.confirmEmailSender,
        writer.uint32(314).fork()
      ).ldelim();
    }
    if (message.confirmEmailContent !== undefined) {
      writer.uint32(322).string(message.confirmEmailContent);
    }
    if (message.confirmEmailContentExcerpt !== undefined) {
      writer.uint32(330).string(message.confirmEmailContentExcerpt);
    }
    if (message.confirmEmailContentDefault !== undefined) {
      writer.uint32(338).string(message.confirmEmailContentDefault);
    }
    if (message.confirmEmailContentDefaultExcerpt !== undefined) {
      writer.uint32(346).string(message.confirmEmailContentDefaultExcerpt);
    }
    if (message.confirmEmailTitle !== undefined) {
      writer.uint32(354).string(message.confirmEmailTitle);
    }
    if (message.confirmEmailTitleDefault !== undefined) {
      writer.uint32(362).string(message.confirmEmailTitleDefault);
    }
    if (message.rank !== 0) {
      writer.uint32(368).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(376).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(384).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(394).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(402).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): NotificationConfigEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNotificationConfigEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = NotificationConfigEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.cascadeToSubWorkspaces = reader.bool();
          break;
        case 10:
          message.forcedCascadeEmailProvider = reader.bool();
          break;
        case 12:
          message.generalEmailProviderId = reader.string();
          break;
        case 13:
          message.generalEmailProvider = EmailProviderEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 15:
          message.generalGsmProviderId = reader.string();
          break;
        case 16:
          message.generalGsmProvider = GsmProviderEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 17:
          message.inviteToWorkspaceContent = reader.string();
          break;
        case 18:
          message.inviteToWorkspaceContentExcerpt = reader.string();
          break;
        case 19:
          message.inviteToWorkspaceContentDefault = reader.string();
          break;
        case 20:
          message.inviteToWorkspaceContentDefaultExcerpt = reader.string();
          break;
        case 21:
          message.inviteToWorkspaceTitle = reader.string();
          break;
        case 22:
          message.inviteToWorkspaceTitleDefault = reader.string();
          break;
        case 24:
          message.inviteToWorkspaceSenderId = reader.string();
          break;
        case 25:
          message.inviteToWorkspaceSender = EmailSenderEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 26:
          message.forgetPasswordContent = reader.string();
          break;
        case 27:
          message.forgetPasswordContentExcerpt = reader.string();
          break;
        case 28:
          message.forgetPasswordContentDefault = reader.string();
          break;
        case 29:
          message.forgetPasswordContentDefaultExcerpt = reader.string();
          break;
        case 30:
          message.forgetPasswordTitle = reader.string();
          break;
        case 31:
          message.forgetPasswordTitleDefault = reader.string();
          break;
        case 33:
          message.forgetPasswordSenderId = reader.string();
          break;
        case 34:
          message.forgetPasswordSender = EmailSenderEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 35:
          message.acceptLanguage = reader.string();
          break;
        case 36:
          message.acceptLanguageExcerpt = reader.string();
          break;
        case 38:
          message.confirmEmailSenderId = reader.string();
          break;
        case 39:
          message.confirmEmailSender = EmailSenderEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 40:
          message.confirmEmailContent = reader.string();
          break;
        case 41:
          message.confirmEmailContentExcerpt = reader.string();
          break;
        case 42:
          message.confirmEmailContentDefault = reader.string();
          break;
        case 43:
          message.confirmEmailContentDefaultExcerpt = reader.string();
          break;
        case 44:
          message.confirmEmailTitle = reader.string();
          break;
        case 45:
          message.confirmEmailTitleDefault = reader.string();
          break;
        case 46:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 47:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 48:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 49:
          message.createdFormatted = reader.string();
          break;
        case 50:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NotificationConfigEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? NotificationConfigEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      cascadeToSubWorkspaces: isSet(object.cascadeToSubWorkspaces)
        ? Boolean(object.cascadeToSubWorkspaces)
        : undefined,
      forcedCascadeEmailProvider: isSet(object.forcedCascadeEmailProvider)
        ? Boolean(object.forcedCascadeEmailProvider)
        : undefined,
      generalEmailProviderId: isSet(object.generalEmailProviderId)
        ? String(object.generalEmailProviderId)
        : undefined,
      generalEmailProvider: isSet(object.generalEmailProvider)
        ? EmailProviderEntity.fromJSON(object.generalEmailProvider)
        : undefined,
      generalGsmProviderId: isSet(object.generalGsmProviderId)
        ? String(object.generalGsmProviderId)
        : undefined,
      generalGsmProvider: isSet(object.generalGsmProvider)
        ? GsmProviderEntity.fromJSON(object.generalGsmProvider)
        : undefined,
      inviteToWorkspaceContent: isSet(object.inviteToWorkspaceContent)
        ? String(object.inviteToWorkspaceContent)
        : undefined,
      inviteToWorkspaceContentExcerpt: isSet(
        object.inviteToWorkspaceContentExcerpt
      )
        ? String(object.inviteToWorkspaceContentExcerpt)
        : undefined,
      inviteToWorkspaceContentDefault: isSet(
        object.inviteToWorkspaceContentDefault
      )
        ? String(object.inviteToWorkspaceContentDefault)
        : undefined,
      inviteToWorkspaceContentDefaultExcerpt: isSet(
        object.inviteToWorkspaceContentDefaultExcerpt
      )
        ? String(object.inviteToWorkspaceContentDefaultExcerpt)
        : undefined,
      inviteToWorkspaceTitle: isSet(object.inviteToWorkspaceTitle)
        ? String(object.inviteToWorkspaceTitle)
        : undefined,
      inviteToWorkspaceTitleDefault: isSet(object.inviteToWorkspaceTitleDefault)
        ? String(object.inviteToWorkspaceTitleDefault)
        : undefined,
      inviteToWorkspaceSenderId: isSet(object.inviteToWorkspaceSenderId)
        ? String(object.inviteToWorkspaceSenderId)
        : undefined,
      inviteToWorkspaceSender: isSet(object.inviteToWorkspaceSender)
        ? EmailSenderEntity.fromJSON(object.inviteToWorkspaceSender)
        : undefined,
      forgetPasswordContent: isSet(object.forgetPasswordContent)
        ? String(object.forgetPasswordContent)
        : undefined,
      forgetPasswordContentExcerpt: isSet(object.forgetPasswordContentExcerpt)
        ? String(object.forgetPasswordContentExcerpt)
        : undefined,
      forgetPasswordContentDefault: isSet(object.forgetPasswordContentDefault)
        ? String(object.forgetPasswordContentDefault)
        : undefined,
      forgetPasswordContentDefaultExcerpt: isSet(
        object.forgetPasswordContentDefaultExcerpt
      )
        ? String(object.forgetPasswordContentDefaultExcerpt)
        : undefined,
      forgetPasswordTitle: isSet(object.forgetPasswordTitle)
        ? String(object.forgetPasswordTitle)
        : undefined,
      forgetPasswordTitleDefault: isSet(object.forgetPasswordTitleDefault)
        ? String(object.forgetPasswordTitleDefault)
        : undefined,
      forgetPasswordSenderId: isSet(object.forgetPasswordSenderId)
        ? String(object.forgetPasswordSenderId)
        : undefined,
      forgetPasswordSender: isSet(object.forgetPasswordSender)
        ? EmailSenderEntity.fromJSON(object.forgetPasswordSender)
        : undefined,
      acceptLanguage: isSet(object.acceptLanguage)
        ? String(object.acceptLanguage)
        : undefined,
      acceptLanguageExcerpt: isSet(object.acceptLanguageExcerpt)
        ? String(object.acceptLanguageExcerpt)
        : undefined,
      confirmEmailSenderId: isSet(object.confirmEmailSenderId)
        ? String(object.confirmEmailSenderId)
        : undefined,
      confirmEmailSender: isSet(object.confirmEmailSender)
        ? EmailSenderEntity.fromJSON(object.confirmEmailSender)
        : undefined,
      confirmEmailContent: isSet(object.confirmEmailContent)
        ? String(object.confirmEmailContent)
        : undefined,
      confirmEmailContentExcerpt: isSet(object.confirmEmailContentExcerpt)
        ? String(object.confirmEmailContentExcerpt)
        : undefined,
      confirmEmailContentDefault: isSet(object.confirmEmailContentDefault)
        ? String(object.confirmEmailContentDefault)
        : undefined,
      confirmEmailContentDefaultExcerpt: isSet(
        object.confirmEmailContentDefaultExcerpt
      )
        ? String(object.confirmEmailContentDefaultExcerpt)
        : undefined,
      confirmEmailTitle: isSet(object.confirmEmailTitle)
        ? String(object.confirmEmailTitle)
        : undefined,
      confirmEmailTitleDefault: isSet(object.confirmEmailTitleDefault)
        ? String(object.confirmEmailTitleDefault)
        : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: NotificationConfigEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? NotificationConfigEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.cascadeToSubWorkspaces !== undefined &&
      (obj.cascadeToSubWorkspaces = message.cascadeToSubWorkspaces);
    message.forcedCascadeEmailProvider !== undefined &&
      (obj.forcedCascadeEmailProvider = message.forcedCascadeEmailProvider);
    message.generalEmailProviderId !== undefined &&
      (obj.generalEmailProviderId = message.generalEmailProviderId);
    message.generalEmailProvider !== undefined &&
      (obj.generalEmailProvider = message.generalEmailProvider
        ? EmailProviderEntity.toJSON(message.generalEmailProvider)
        : undefined);
    message.generalGsmProviderId !== undefined &&
      (obj.generalGsmProviderId = message.generalGsmProviderId);
    message.generalGsmProvider !== undefined &&
      (obj.generalGsmProvider = message.generalGsmProvider
        ? GsmProviderEntity.toJSON(message.generalGsmProvider)
        : undefined);
    message.inviteToWorkspaceContent !== undefined &&
      (obj.inviteToWorkspaceContent = message.inviteToWorkspaceContent);
    message.inviteToWorkspaceContentExcerpt !== undefined &&
      (obj.inviteToWorkspaceContentExcerpt =
        message.inviteToWorkspaceContentExcerpt);
    message.inviteToWorkspaceContentDefault !== undefined &&
      (obj.inviteToWorkspaceContentDefault =
        message.inviteToWorkspaceContentDefault);
    message.inviteToWorkspaceContentDefaultExcerpt !== undefined &&
      (obj.inviteToWorkspaceContentDefaultExcerpt =
        message.inviteToWorkspaceContentDefaultExcerpt);
    message.inviteToWorkspaceTitle !== undefined &&
      (obj.inviteToWorkspaceTitle = message.inviteToWorkspaceTitle);
    message.inviteToWorkspaceTitleDefault !== undefined &&
      (obj.inviteToWorkspaceTitleDefault =
        message.inviteToWorkspaceTitleDefault);
    message.inviteToWorkspaceSenderId !== undefined &&
      (obj.inviteToWorkspaceSenderId = message.inviteToWorkspaceSenderId);
    message.inviteToWorkspaceSender !== undefined &&
      (obj.inviteToWorkspaceSender = message.inviteToWorkspaceSender
        ? EmailSenderEntity.toJSON(message.inviteToWorkspaceSender)
        : undefined);
    message.forgetPasswordContent !== undefined &&
      (obj.forgetPasswordContent = message.forgetPasswordContent);
    message.forgetPasswordContentExcerpt !== undefined &&
      (obj.forgetPasswordContentExcerpt = message.forgetPasswordContentExcerpt);
    message.forgetPasswordContentDefault !== undefined &&
      (obj.forgetPasswordContentDefault = message.forgetPasswordContentDefault);
    message.forgetPasswordContentDefaultExcerpt !== undefined &&
      (obj.forgetPasswordContentDefaultExcerpt =
        message.forgetPasswordContentDefaultExcerpt);
    message.forgetPasswordTitle !== undefined &&
      (obj.forgetPasswordTitle = message.forgetPasswordTitle);
    message.forgetPasswordTitleDefault !== undefined &&
      (obj.forgetPasswordTitleDefault = message.forgetPasswordTitleDefault);
    message.forgetPasswordSenderId !== undefined &&
      (obj.forgetPasswordSenderId = message.forgetPasswordSenderId);
    message.forgetPasswordSender !== undefined &&
      (obj.forgetPasswordSender = message.forgetPasswordSender
        ? EmailSenderEntity.toJSON(message.forgetPasswordSender)
        : undefined);
    message.acceptLanguage !== undefined &&
      (obj.acceptLanguage = message.acceptLanguage);
    message.acceptLanguageExcerpt !== undefined &&
      (obj.acceptLanguageExcerpt = message.acceptLanguageExcerpt);
    message.confirmEmailSenderId !== undefined &&
      (obj.confirmEmailSenderId = message.confirmEmailSenderId);
    message.confirmEmailSender !== undefined &&
      (obj.confirmEmailSender = message.confirmEmailSender
        ? EmailSenderEntity.toJSON(message.confirmEmailSender)
        : undefined);
    message.confirmEmailContent !== undefined &&
      (obj.confirmEmailContent = message.confirmEmailContent);
    message.confirmEmailContentExcerpt !== undefined &&
      (obj.confirmEmailContentExcerpt = message.confirmEmailContentExcerpt);
    message.confirmEmailContentDefault !== undefined &&
      (obj.confirmEmailContentDefault = message.confirmEmailContentDefault);
    message.confirmEmailContentDefaultExcerpt !== undefined &&
      (obj.confirmEmailContentDefaultExcerpt =
        message.confirmEmailContentDefaultExcerpt);
    message.confirmEmailTitle !== undefined &&
      (obj.confirmEmailTitle = message.confirmEmailTitle);
    message.confirmEmailTitleDefault !== undefined &&
      (obj.confirmEmailTitleDefault = message.confirmEmailTitleDefault);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<NotificationConfigEntity>, I>>(
    base?: I
  ): NotificationConfigEntity {
    return NotificationConfigEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<NotificationConfigEntity>, I>>(
    object: I
  ): NotificationConfigEntity {
    const message = createBaseNotificationConfigEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? NotificationConfigEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.cascadeToSubWorkspaces = object.cascadeToSubWorkspaces ?? undefined;
    message.forcedCascadeEmailProvider =
      object.forcedCascadeEmailProvider ?? undefined;
    message.generalEmailProviderId = object.generalEmailProviderId ?? undefined;
    message.generalEmailProvider =
      object.generalEmailProvider !== undefined &&
      object.generalEmailProvider !== null
        ? EmailProviderEntity.fromPartial(object.generalEmailProvider)
        : undefined;
    message.generalGsmProviderId = object.generalGsmProviderId ?? undefined;
    message.generalGsmProvider =
      object.generalGsmProvider !== undefined &&
      object.generalGsmProvider !== null
        ? GsmProviderEntity.fromPartial(object.generalGsmProvider)
        : undefined;
    message.inviteToWorkspaceContent =
      object.inviteToWorkspaceContent ?? undefined;
    message.inviteToWorkspaceContentExcerpt =
      object.inviteToWorkspaceContentExcerpt ?? undefined;
    message.inviteToWorkspaceContentDefault =
      object.inviteToWorkspaceContentDefault ?? undefined;
    message.inviteToWorkspaceContentDefaultExcerpt =
      object.inviteToWorkspaceContentDefaultExcerpt ?? undefined;
    message.inviteToWorkspaceTitle = object.inviteToWorkspaceTitle ?? undefined;
    message.inviteToWorkspaceTitleDefault =
      object.inviteToWorkspaceTitleDefault ?? undefined;
    message.inviteToWorkspaceSenderId =
      object.inviteToWorkspaceSenderId ?? undefined;
    message.inviteToWorkspaceSender =
      object.inviteToWorkspaceSender !== undefined &&
      object.inviteToWorkspaceSender !== null
        ? EmailSenderEntity.fromPartial(object.inviteToWorkspaceSender)
        : undefined;
    message.forgetPasswordContent = object.forgetPasswordContent ?? undefined;
    message.forgetPasswordContentExcerpt =
      object.forgetPasswordContentExcerpt ?? undefined;
    message.forgetPasswordContentDefault =
      object.forgetPasswordContentDefault ?? undefined;
    message.forgetPasswordContentDefaultExcerpt =
      object.forgetPasswordContentDefaultExcerpt ?? undefined;
    message.forgetPasswordTitle = object.forgetPasswordTitle ?? undefined;
    message.forgetPasswordTitleDefault =
      object.forgetPasswordTitleDefault ?? undefined;
    message.forgetPasswordSenderId = object.forgetPasswordSenderId ?? undefined;
    message.forgetPasswordSender =
      object.forgetPasswordSender !== undefined &&
      object.forgetPasswordSender !== null
        ? EmailSenderEntity.fromPartial(object.forgetPasswordSender)
        : undefined;
    message.acceptLanguage = object.acceptLanguage ?? undefined;
    message.acceptLanguageExcerpt = object.acceptLanguageExcerpt ?? undefined;
    message.confirmEmailSenderId = object.confirmEmailSenderId ?? undefined;
    message.confirmEmailSender =
      object.confirmEmailSender !== undefined &&
      object.confirmEmailSender !== null
        ? EmailSenderEntity.fromPartial(object.confirmEmailSender)
        : undefined;
    message.confirmEmailContent = object.confirmEmailContent ?? undefined;
    message.confirmEmailContentExcerpt =
      object.confirmEmailContentExcerpt ?? undefined;
    message.confirmEmailContentDefault =
      object.confirmEmailContentDefault ?? undefined;
    message.confirmEmailContentDefaultExcerpt =
      object.confirmEmailContentDefaultExcerpt ?? undefined;
    message.confirmEmailTitle = object.confirmEmailTitle ?? undefined;
    message.confirmEmailTitleDefault =
      object.confirmEmailTitleDefault ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseCreateByMailResponse(): CreateByMailResponse {
  return { error: undefined, data: undefined };
}

export const CreateByMailResponse = {
  encode(
    message: CreateByMailResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(10).fork()).ldelim();
    }
    if (message.data !== undefined) {
      UserSessionDto.encode(message.data, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): CreateByMailResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateByMailResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.error = IError.decode(reader, reader.uint32());
          break;
        case 2:
          message.data = UserSessionDto.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateByMailResponse {
    return {
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
      data: isSet(object.data)
        ? UserSessionDto.fromJSON(object.data)
        : undefined,
    };
  },

  toJSON(message: CreateByMailResponse): unknown {
    const obj: any = {};
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    message.data !== undefined &&
      (obj.data = message.data
        ? UserSessionDto.toJSON(message.data)
        : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateByMailResponse>, I>>(
    base?: I
  ): CreateByMailResponse {
    return CreateByMailResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateByMailResponse>, I>>(
    object: I
  ): CreateByMailResponse {
    const message = createBaseCreateByMailResponse();
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    message.data =
      object.data !== undefined && object.data !== null
        ? UserSessionDto.fromPartial(object.data)
        : undefined;
    return message;
  },
};

function createBaseEmailAccountSignupDto(): EmailAccountSignupDto {
  return {
    email: "",
    password: "",
    firstName: "",
    lastName: "",
    inviteId: "",
    publicJoinKeyId: "",
    workspaceTypeId: "",
  };
}

export const EmailAccountSignupDto = {
  encode(
    message: EmailAccountSignupDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.email !== "") {
      writer.uint32(10).string(message.email);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    if (message.firstName !== "") {
      writer.uint32(26).string(message.firstName);
    }
    if (message.lastName !== "") {
      writer.uint32(34).string(message.lastName);
    }
    if (message.inviteId !== "") {
      writer.uint32(42).string(message.inviteId);
    }
    if (message.publicJoinKeyId !== "") {
      writer.uint32(50).string(message.publicJoinKeyId);
    }
    if (message.workspaceTypeId !== "") {
      writer.uint32(58).string(message.workspaceTypeId);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): EmailAccountSignupDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailAccountSignupDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.email = reader.string();
          break;
        case 2:
          message.password = reader.string();
          break;
        case 3:
          message.firstName = reader.string();
          break;
        case 4:
          message.lastName = reader.string();
          break;
        case 5:
          message.inviteId = reader.string();
          break;
        case 6:
          message.publicJoinKeyId = reader.string();
          break;
        case 7:
          message.workspaceTypeId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailAccountSignupDto {
    return {
      email: isSet(object.email) ? String(object.email) : "",
      password: isSet(object.password) ? String(object.password) : "",
      firstName: isSet(object.firstName) ? String(object.firstName) : "",
      lastName: isSet(object.lastName) ? String(object.lastName) : "",
      inviteId: isSet(object.inviteId) ? String(object.inviteId) : "",
      publicJoinKeyId: isSet(object.publicJoinKeyId)
        ? String(object.publicJoinKeyId)
        : "",
      workspaceTypeId: isSet(object.workspaceTypeId)
        ? String(object.workspaceTypeId)
        : "",
    };
  },

  toJSON(message: EmailAccountSignupDto): unknown {
    const obj: any = {};
    message.email !== undefined && (obj.email = message.email);
    message.password !== undefined && (obj.password = message.password);
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    message.inviteId !== undefined && (obj.inviteId = message.inviteId);
    message.publicJoinKeyId !== undefined &&
      (obj.publicJoinKeyId = message.publicJoinKeyId);
    message.workspaceTypeId !== undefined &&
      (obj.workspaceTypeId = message.workspaceTypeId);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailAccountSignupDto>, I>>(
    base?: I
  ): EmailAccountSignupDto {
    return EmailAccountSignupDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailAccountSignupDto>, I>>(
    object: I
  ): EmailAccountSignupDto {
    const message = createBaseEmailAccountSignupDto();
    message.email = object.email ?? "";
    message.password = object.password ?? "";
    message.firstName = object.firstName ?? "";
    message.lastName = object.lastName ?? "";
    message.inviteId = object.inviteId ?? "";
    message.publicJoinKeyId = object.publicJoinKeyId ?? "";
    message.workspaceTypeId = object.workspaceTypeId ?? "";
    return message;
  },
};

function createBaseEmailAccountSigninDto(): EmailAccountSigninDto {
  return { email: "", password: "" };
}

export const EmailAccountSigninDto = {
  encode(
    message: EmailAccountSigninDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.email !== "") {
      writer.uint32(10).string(message.email);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): EmailAccountSigninDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailAccountSigninDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.email = reader.string();
          break;
        case 2:
          message.password = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailAccountSigninDto {
    return {
      email: isSet(object.email) ? String(object.email) : "",
      password: isSet(object.password) ? String(object.password) : "",
    };
  },

  toJSON(message: EmailAccountSigninDto): unknown {
    const obj: any = {};
    message.email !== undefined && (obj.email = message.email);
    message.password !== undefined && (obj.password = message.password);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailAccountSigninDto>, I>>(
    base?: I
  ): EmailAccountSigninDto {
    return EmailAccountSigninDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailAccountSigninDto>, I>>(
    object: I
  ): EmailAccountSigninDto {
    const message = createBaseEmailAccountSigninDto();
    message.email = object.email ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBasePhoneNumberAccountCreationDto(): PhoneNumberAccountCreationDto {
  return { phoneNumber: "" };
}

export const PhoneNumberAccountCreationDto = {
  encode(
    message: PhoneNumberAccountCreationDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.phoneNumber !== "") {
      writer.uint32(10).string(message.phoneNumber);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PhoneNumberAccountCreationDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePhoneNumberAccountCreationDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.phoneNumber = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PhoneNumberAccountCreationDto {
    return {
      phoneNumber: isSet(object.phoneNumber) ? String(object.phoneNumber) : "",
    };
  },

  toJSON(message: PhoneNumberAccountCreationDto): unknown {
    const obj: any = {};
    message.phoneNumber !== undefined &&
      (obj.phoneNumber = message.phoneNumber);
    return obj;
  },

  create<I extends Exact<DeepPartial<PhoneNumberAccountCreationDto>, I>>(
    base?: I
  ): PhoneNumberAccountCreationDto {
    return PhoneNumberAccountCreationDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PhoneNumberAccountCreationDto>, I>>(
    object: I
  ): PhoneNumberAccountCreationDto {
    const message = createBasePhoneNumberAccountCreationDto();
    message.phoneNumber = object.phoneNumber ?? "";
    return message;
  },
};

function createBasePhoneNumberUniversalAuthenticateDto(): PhoneNumberUniversalAuthenticateDto {
  return { phoneNumber: "" };
}

export const PhoneNumberUniversalAuthenticateDto = {
  encode(
    message: PhoneNumberUniversalAuthenticateDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.phoneNumber !== "") {
      writer.uint32(10).string(message.phoneNumber);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PhoneNumberUniversalAuthenticateDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePhoneNumberUniversalAuthenticateDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.phoneNumber = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PhoneNumberUniversalAuthenticateDto {
    return {
      phoneNumber: isSet(object.phoneNumber) ? String(object.phoneNumber) : "",
    };
  },

  toJSON(message: PhoneNumberUniversalAuthenticateDto): unknown {
    const obj: any = {};
    message.phoneNumber !== undefined &&
      (obj.phoneNumber = message.phoneNumber);
    return obj;
  },

  create<I extends Exact<DeepPartial<PhoneNumberUniversalAuthenticateDto>, I>>(
    base?: I
  ): PhoneNumberUniversalAuthenticateDto {
    return PhoneNumberUniversalAuthenticateDto.fromPartial(base ?? {});
  },

  fromPartial<
    I extends Exact<DeepPartial<PhoneNumberUniversalAuthenticateDto>, I>
  >(object: I): PhoneNumberUniversalAuthenticateDto {
    const message = createBasePhoneNumberUniversalAuthenticateDto();
    message.phoneNumber = object.phoneNumber ?? "";
    return message;
  },
};

function createBaseUserSessionDto(): UserSessionDto {
  return {
    passport: undefined,
    token: "",
    exchangeKey: "",
    userRoleWorkspaces: [],
    user: undefined,
  };
}

export const UserSessionDto = {
  encode(
    message: UserSessionDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.passport !== undefined) {
      PassportEntity.encode(
        message.passport,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.token !== "") {
      writer.uint32(18).string(message.token);
    }
    if (message.exchangeKey !== "") {
      writer.uint32(26).string(message.exchangeKey);
    }
    for (const v of message.userRoleWorkspaces) {
      UserRoleWorkspaceEntity.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.user !== undefined) {
      UserEntity.encode(message.user, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserSessionDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserSessionDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.passport = PassportEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.token = reader.string();
          break;
        case 3:
          message.exchangeKey = reader.string();
          break;
        case 4:
          message.userRoleWorkspaces.push(
            UserRoleWorkspaceEntity.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.user = UserEntity.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserSessionDto {
    return {
      passport: isSet(object.passport)
        ? PassportEntity.fromJSON(object.passport)
        : undefined,
      token: isSet(object.token) ? String(object.token) : "",
      exchangeKey: isSet(object.exchangeKey) ? String(object.exchangeKey) : "",
      userRoleWorkspaces: Array.isArray(object?.userRoleWorkspaces)
        ? object.userRoleWorkspaces.map((e: any) =>
            UserRoleWorkspaceEntity.fromJSON(e)
          )
        : [],
      user: isSet(object.user) ? UserEntity.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: UserSessionDto): unknown {
    const obj: any = {};
    message.passport !== undefined &&
      (obj.passport = message.passport
        ? PassportEntity.toJSON(message.passport)
        : undefined);
    message.token !== undefined && (obj.token = message.token);
    message.exchangeKey !== undefined &&
      (obj.exchangeKey = message.exchangeKey);
    if (message.userRoleWorkspaces) {
      obj.userRoleWorkspaces = message.userRoleWorkspaces.map((e) =>
        e ? UserRoleWorkspaceEntity.toJSON(e) : undefined
      );
    } else {
      obj.userRoleWorkspaces = [];
    }
    message.user !== undefined &&
      (obj.user = message.user ? UserEntity.toJSON(message.user) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserSessionDto>, I>>(
    base?: I
  ): UserSessionDto {
    return UserSessionDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserSessionDto>, I>>(
    object: I
  ): UserSessionDto {
    const message = createBaseUserSessionDto();
    message.passport =
      object.passport !== undefined && object.passport !== null
        ? PassportEntity.fromPartial(object.passport)
        : undefined;
    message.token = object.token ?? "";
    message.exchangeKey = object.exchangeKey ?? "";
    message.userRoleWorkspaces =
      object.userRoleWorkspaces?.map((e) =>
        UserRoleWorkspaceEntity.fromPartial(e)
      ) || [];
    message.user =
      object.user !== undefined && object.user !== null
        ? UserEntity.fromPartial(object.user)
        : undefined;
    return message;
  },
};

function createBaseOtpAuthenticateDto(): OtpAuthenticateDto {
  return { value: "", otp: "", type: "", password: "" };
}

export const OtpAuthenticateDto = {
  encode(
    message: OtpAuthenticateDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.value !== "") {
      writer.uint32(10).string(message.value);
    }
    if (message.otp !== "") {
      writer.uint32(18).string(message.otp);
    }
    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }
    if (message.password !== "") {
      writer.uint32(34).string(message.password);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OtpAuthenticateDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOtpAuthenticateDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.value = reader.string();
          break;
        case 2:
          message.otp = reader.string();
          break;
        case 3:
          message.type = reader.string();
          break;
        case 4:
          message.password = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OtpAuthenticateDto {
    return {
      value: isSet(object.value) ? String(object.value) : "",
      otp: isSet(object.otp) ? String(object.otp) : "",
      type: isSet(object.type) ? String(object.type) : "",
      password: isSet(object.password) ? String(object.password) : "",
    };
  },

  toJSON(message: OtpAuthenticateDto): unknown {
    const obj: any = {};
    message.value !== undefined && (obj.value = message.value);
    message.otp !== undefined && (obj.otp = message.otp);
    message.type !== undefined && (obj.type = message.type);
    message.password !== undefined && (obj.password = message.password);
    return obj;
  },

  create<I extends Exact<DeepPartial<OtpAuthenticateDto>, I>>(
    base?: I
  ): OtpAuthenticateDto {
    return OtpAuthenticateDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<OtpAuthenticateDto>, I>>(
    object: I
  ): OtpAuthenticateDto {
    const message = createBaseOtpAuthenticateDto();
    message.value = object.value ?? "";
    message.otp = object.otp ?? "";
    message.type = object.type ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseEmailOtpResponse(): EmailOtpResponse {
  return { request: undefined, userSession: undefined };
}

export const EmailOtpResponse = {
  encode(
    message: EmailOtpResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.request !== undefined) {
      ForgetPasswordEntity.encode(
        message.request,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.userSession !== undefined) {
      UserSessionDto.encode(
        message.userSession,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EmailOtpResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEmailOtpResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.request = ForgetPasswordEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.userSession = UserSessionDto.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EmailOtpResponse {
    return {
      request: isSet(object.request)
        ? ForgetPasswordEntity.fromJSON(object.request)
        : undefined,
      userSession: isSet(object.userSession)
        ? UserSessionDto.fromJSON(object.userSession)
        : undefined,
    };
  },

  toJSON(message: EmailOtpResponse): unknown {
    const obj: any = {};
    message.request !== undefined &&
      (obj.request = message.request
        ? ForgetPasswordEntity.toJSON(message.request)
        : undefined);
    message.userSession !== undefined &&
      (obj.userSession = message.userSession
        ? UserSessionDto.toJSON(message.userSession)
        : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<EmailOtpResponse>, I>>(
    base?: I
  ): EmailOtpResponse {
    return EmailOtpResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EmailOtpResponse>, I>>(
    object: I
  ): EmailOtpResponse {
    const message = createBaseEmailOtpResponse();
    message.request =
      object.request !== undefined && object.request !== null
        ? ForgetPasswordEntity.fromPartial(object.request)
        : undefined;
    message.userSession =
      object.userSession !== undefined && object.userSession !== null
        ? UserSessionDto.fromPartial(object.userSession)
        : undefined;
    return message;
  },
};

function createBaseResetEmailDto(): ResetEmailDto {
  return { password: "" };
}

export const ResetEmailDto = {
  encode(
    message: ResetEmailDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.password !== "") {
      writer.uint32(10).string(message.password);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ResetEmailDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResetEmailDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.password = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ResetEmailDto {
    return { password: isSet(object.password) ? String(object.password) : "" };
  },

  toJSON(message: ResetEmailDto): unknown {
    const obj: any = {};
    message.password !== undefined && (obj.password = message.password);
    return obj;
  },

  create<I extends Exact<DeepPartial<ResetEmailDto>, I>>(
    base?: I
  ): ResetEmailDto {
    return ResetEmailDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ResetEmailDto>, I>>(
    object: I
  ): ResetEmailDto {
    const message = createBaseResetEmailDto();
    message.password = object.password ?? "";
    return message;
  },
};

function createBasePassportCreateReply(): PassportCreateReply {
  return { data: undefined, error: undefined };
}

export const PassportCreateReply = {
  encode(
    message: PassportCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PassportEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PassportCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePassportCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PassportEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PassportCreateReply {
    return {
      data: isSet(object.data)
        ? PassportEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PassportCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PassportEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PassportCreateReply>, I>>(
    base?: I
  ): PassportCreateReply {
    return PassportCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PassportCreateReply>, I>>(
    object: I
  ): PassportCreateReply {
    const message = createBasePassportCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PassportEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePassportReply(): PassportReply {
  return { data: undefined, error: undefined };
}

export const PassportReply = {
  encode(
    message: PassportReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PassportEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PassportReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePassportReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PassportEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PassportReply {
    return {
      data: isSet(object.data)
        ? PassportEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PassportReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PassportEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PassportReply>, I>>(
    base?: I
  ): PassportReply {
    return PassportReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PassportReply>, I>>(
    object: I
  ): PassportReply {
    const message = createBasePassportReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PassportEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePassportQueryReply(): PassportQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const PassportQueryReply = {
  encode(
    message: PassportQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      PassportEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PassportQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePassportQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(PassportEntity.decode(reader, reader.uint32()));
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PassportQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => PassportEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PassportQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? PassportEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PassportQueryReply>, I>>(
    base?: I
  ): PassportQueryReply {
    return PassportQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PassportQueryReply>, I>>(
    object: I
  ): PassportQueryReply {
    const message = createBasePassportQueryReply();
    message.items =
      object.items?.map((e) => PassportEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePassportEntity(): PassportEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    type: undefined,
    value: undefined,
    password: undefined,
    confirmed: undefined,
    accessToken: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const PassportEntity = {
  encode(
    message: PassportEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      PassportEntity.encode(message.parent, writer.uint32(42).fork()).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.type !== undefined) {
      writer.uint32(74).string(message.type);
    }
    if (message.value !== undefined) {
      writer.uint32(82).string(message.value);
    }
    if (message.password !== undefined) {
      writer.uint32(90).string(message.password);
    }
    if (message.confirmed !== undefined) {
      writer.uint32(96).bool(message.confirmed);
    }
    if (message.accessToken !== undefined) {
      writer.uint32(106).string(message.accessToken);
    }
    if (message.rank !== 0) {
      writer.uint32(112).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(120).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(128).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(138).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(146).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PassportEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePassportEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = PassportEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.type = reader.string();
          break;
        case 10:
          message.value = reader.string();
          break;
        case 11:
          message.password = reader.string();
          break;
        case 12:
          message.confirmed = reader.bool();
          break;
        case 13:
          message.accessToken = reader.string();
          break;
        case 14:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.createdFormatted = reader.string();
          break;
        case 18:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PassportEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? PassportEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      type: isSet(object.type) ? String(object.type) : undefined,
      value: isSet(object.value) ? String(object.value) : undefined,
      password: isSet(object.password) ? String(object.password) : undefined,
      confirmed: isSet(object.confirmed)
        ? Boolean(object.confirmed)
        : undefined,
      accessToken: isSet(object.accessToken)
        ? String(object.accessToken)
        : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: PassportEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? PassportEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.type !== undefined && (obj.type = message.type);
    message.value !== undefined && (obj.value = message.value);
    message.password !== undefined && (obj.password = message.password);
    message.confirmed !== undefined && (obj.confirmed = message.confirmed);
    message.accessToken !== undefined &&
      (obj.accessToken = message.accessToken);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<PassportEntity>, I>>(
    base?: I
  ): PassportEntity {
    return PassportEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PassportEntity>, I>>(
    object: I
  ): PassportEntity {
    const message = createBasePassportEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? PassportEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.type = object.type ?? undefined;
    message.value = object.value ?? undefined;
    message.password = object.password ?? undefined;
    message.confirmed = object.confirmed ?? undefined;
    message.accessToken = object.accessToken ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBasePassportMethodCreateReply(): PassportMethodCreateReply {
  return { data: undefined, error: undefined };
}

export const PassportMethodCreateReply = {
  encode(
    message: PassportMethodCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PassportMethodEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PassportMethodCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePassportMethodCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PassportMethodEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PassportMethodCreateReply {
    return {
      data: isSet(object.data)
        ? PassportMethodEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PassportMethodCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PassportMethodEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PassportMethodCreateReply>, I>>(
    base?: I
  ): PassportMethodCreateReply {
    return PassportMethodCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PassportMethodCreateReply>, I>>(
    object: I
  ): PassportMethodCreateReply {
    const message = createBasePassportMethodCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PassportMethodEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePassportMethodReply(): PassportMethodReply {
  return { data: undefined, error: undefined };
}

export const PassportMethodReply = {
  encode(
    message: PassportMethodReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PassportMethodEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PassportMethodReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePassportMethodReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PassportMethodEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PassportMethodReply {
    return {
      data: isSet(object.data)
        ? PassportMethodEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PassportMethodReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PassportMethodEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PassportMethodReply>, I>>(
    base?: I
  ): PassportMethodReply {
    return PassportMethodReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PassportMethodReply>, I>>(
    object: I
  ): PassportMethodReply {
    const message = createBasePassportMethodReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PassportMethodEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePassportMethodQueryReply(): PassportMethodQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const PassportMethodQueryReply = {
  encode(
    message: PassportMethodQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      PassportMethodEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PassportMethodQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePassportMethodQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            PassportMethodEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PassportMethodQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => PassportMethodEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PassportMethodQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? PassportMethodEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PassportMethodQueryReply>, I>>(
    base?: I
  ): PassportMethodQueryReply {
    return PassportMethodQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PassportMethodQueryReply>, I>>(
    object: I
  ): PassportMethodQueryReply {
    const message = createBasePassportMethodQueryReply();
    message.items =
      object.items?.map((e) => PassportMethodEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePassportMethodEntity(): PassportMethodEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    translations: [],
    name: undefined,
    type: undefined,
    region: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const PassportMethodEntity = {
  encode(
    message: PassportMethodEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      PassportMethodEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    for (const v of message.translations) {
      PassportMethodEntityPolyglot.encode(
        v!,
        writer.uint32(66).fork()
      ).ldelim();
    }
    if (message.name !== undefined) {
      writer.uint32(82).string(message.name);
    }
    if (message.type !== undefined) {
      writer.uint32(90).string(message.type);
    }
    if (message.region !== undefined) {
      writer.uint32(98).string(message.region);
    }
    if (message.rank !== 0) {
      writer.uint32(104).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(112).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(120).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(130).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(138).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PassportMethodEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePassportMethodEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = PassportMethodEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 8:
          message.translations.push(
            PassportMethodEntityPolyglot.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.name = reader.string();
          break;
        case 11:
          message.type = reader.string();
          break;
        case 12:
          message.region = reader.string();
          break;
        case 13:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.createdFormatted = reader.string();
          break;
        case 17:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PassportMethodEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? PassportMethodEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      translations: Array.isArray(object?.translations)
        ? object.translations.map((e: any) =>
            PassportMethodEntityPolyglot.fromJSON(e)
          )
        : [],
      name: isSet(object.name) ? String(object.name) : undefined,
      type: isSet(object.type) ? String(object.type) : undefined,
      region: isSet(object.region) ? String(object.region) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: PassportMethodEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? PassportMethodEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    if (message.translations) {
      obj.translations = message.translations.map((e) =>
        e ? PassportMethodEntityPolyglot.toJSON(e) : undefined
      );
    } else {
      obj.translations = [];
    }
    message.name !== undefined && (obj.name = message.name);
    message.type !== undefined && (obj.type = message.type);
    message.region !== undefined && (obj.region = message.region);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<PassportMethodEntity>, I>>(
    base?: I
  ): PassportMethodEntity {
    return PassportMethodEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PassportMethodEntity>, I>>(
    object: I
  ): PassportMethodEntity {
    const message = createBasePassportMethodEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? PassportMethodEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.translations =
      object.translations?.map((e) =>
        PassportMethodEntityPolyglot.fromPartial(e)
      ) || [];
    message.name = object.name ?? undefined;
    message.type = object.type ?? undefined;
    message.region = object.region ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBasePassportMethodEntityPolyglot(): PassportMethodEntityPolyglot {
  return { linkerId: "", languageId: "", name: "" };
}

export const PassportMethodEntityPolyglot = {
  encode(
    message: PassportMethodEntityPolyglot,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.linkerId !== "") {
      writer.uint32(10).string(message.linkerId);
    }
    if (message.languageId !== "") {
      writer.uint32(18).string(message.languageId);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PassportMethodEntityPolyglot {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePassportMethodEntityPolyglot();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.linkerId = reader.string();
          break;
        case 2:
          message.languageId = reader.string();
          break;
        case 3:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PassportMethodEntityPolyglot {
    return {
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : "",
      languageId: isSet(object.languageId) ? String(object.languageId) : "",
      name: isSet(object.name) ? String(object.name) : "",
    };
  },

  toJSON(message: PassportMethodEntityPolyglot): unknown {
    const obj: any = {};
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.languageId !== undefined && (obj.languageId = message.languageId);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  create<I extends Exact<DeepPartial<PassportMethodEntityPolyglot>, I>>(
    base?: I
  ): PassportMethodEntityPolyglot {
    return PassportMethodEntityPolyglot.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PassportMethodEntityPolyglot>, I>>(
    object: I
  ): PassportMethodEntityPolyglot {
    const message = createBasePassportMethodEntityPolyglot();
    message.linkerId = object.linkerId ?? "";
    message.languageId = object.languageId ?? "";
    message.name = object.name ?? "";
    return message;
  },
};

function createBasePendingWorkspaceInviteCreateReply(): PendingWorkspaceInviteCreateReply {
  return { data: undefined, error: undefined };
}

export const PendingWorkspaceInviteCreateReply = {
  encode(
    message: PendingWorkspaceInviteCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PendingWorkspaceInviteEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PendingWorkspaceInviteCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePendingWorkspaceInviteCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PendingWorkspaceInviteEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PendingWorkspaceInviteCreateReply {
    return {
      data: isSet(object.data)
        ? PendingWorkspaceInviteEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PendingWorkspaceInviteCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PendingWorkspaceInviteEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PendingWorkspaceInviteCreateReply>, I>>(
    base?: I
  ): PendingWorkspaceInviteCreateReply {
    return PendingWorkspaceInviteCreateReply.fromPartial(base ?? {});
  },

  fromPartial<
    I extends Exact<DeepPartial<PendingWorkspaceInviteCreateReply>, I>
  >(object: I): PendingWorkspaceInviteCreateReply {
    const message = createBasePendingWorkspaceInviteCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PendingWorkspaceInviteEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePendingWorkspaceInviteReply(): PendingWorkspaceInviteReply {
  return { data: undefined, error: undefined };
}

export const PendingWorkspaceInviteReply = {
  encode(
    message: PendingWorkspaceInviteReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PendingWorkspaceInviteEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PendingWorkspaceInviteReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePendingWorkspaceInviteReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PendingWorkspaceInviteEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PendingWorkspaceInviteReply {
    return {
      data: isSet(object.data)
        ? PendingWorkspaceInviteEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PendingWorkspaceInviteReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PendingWorkspaceInviteEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PendingWorkspaceInviteReply>, I>>(
    base?: I
  ): PendingWorkspaceInviteReply {
    return PendingWorkspaceInviteReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PendingWorkspaceInviteReply>, I>>(
    object: I
  ): PendingWorkspaceInviteReply {
    const message = createBasePendingWorkspaceInviteReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PendingWorkspaceInviteEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePendingWorkspaceInviteQueryReply(): PendingWorkspaceInviteQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const PendingWorkspaceInviteQueryReply = {
  encode(
    message: PendingWorkspaceInviteQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      PendingWorkspaceInviteEntity.encode(
        v!,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PendingWorkspaceInviteQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePendingWorkspaceInviteQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            PendingWorkspaceInviteEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PendingWorkspaceInviteQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => PendingWorkspaceInviteEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PendingWorkspaceInviteQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? PendingWorkspaceInviteEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PendingWorkspaceInviteQueryReply>, I>>(
    base?: I
  ): PendingWorkspaceInviteQueryReply {
    return PendingWorkspaceInviteQueryReply.fromPartial(base ?? {});
  },

  fromPartial<
    I extends Exact<DeepPartial<PendingWorkspaceInviteQueryReply>, I>
  >(object: I): PendingWorkspaceInviteQueryReply {
    const message = createBasePendingWorkspaceInviteQueryReply();
    message.items =
      object.items?.map((e) => PendingWorkspaceInviteEntity.fromPartial(e)) ||
      [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePendingWorkspaceInviteEntity(): PendingWorkspaceInviteEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    value: undefined,
    type: undefined,
    coverLetter: undefined,
    workspaceName: undefined,
    roleId: undefined,
    role: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const PendingWorkspaceInviteEntity = {
  encode(
    message: PendingWorkspaceInviteEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      PendingWorkspaceInviteEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.value !== undefined) {
      writer.uint32(74).string(message.value);
    }
    if (message.type !== undefined) {
      writer.uint32(82).string(message.type);
    }
    if (message.coverLetter !== undefined) {
      writer.uint32(90).string(message.coverLetter);
    }
    if (message.workspaceName !== undefined) {
      writer.uint32(98).string(message.workspaceName);
    }
    if (message.roleId !== undefined) {
      writer.uint32(114).string(message.roleId);
    }
    if (message.role !== undefined) {
      RoleEntity.encode(message.role, writer.uint32(122).fork()).ldelim();
    }
    if (message.rank !== 0) {
      writer.uint32(128).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(136).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(144).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(154).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(162).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PendingWorkspaceInviteEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePendingWorkspaceInviteEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = PendingWorkspaceInviteEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.value = reader.string();
          break;
        case 10:
          message.type = reader.string();
          break;
        case 11:
          message.coverLetter = reader.string();
          break;
        case 12:
          message.workspaceName = reader.string();
          break;
        case 14:
          message.roleId = reader.string();
          break;
        case 15:
          message.role = RoleEntity.decode(reader, reader.uint32());
          break;
        case 16:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 18:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 19:
          message.createdFormatted = reader.string();
          break;
        case 20:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PendingWorkspaceInviteEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? PendingWorkspaceInviteEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      value: isSet(object.value) ? String(object.value) : undefined,
      type: isSet(object.type) ? String(object.type) : undefined,
      coverLetter: isSet(object.coverLetter)
        ? String(object.coverLetter)
        : undefined,
      workspaceName: isSet(object.workspaceName)
        ? String(object.workspaceName)
        : undefined,
      roleId: isSet(object.roleId) ? String(object.roleId) : undefined,
      role: isSet(object.role) ? RoleEntity.fromJSON(object.role) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: PendingWorkspaceInviteEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? PendingWorkspaceInviteEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.value !== undefined && (obj.value = message.value);
    message.type !== undefined && (obj.type = message.type);
    message.coverLetter !== undefined &&
      (obj.coverLetter = message.coverLetter);
    message.workspaceName !== undefined &&
      (obj.workspaceName = message.workspaceName);
    message.roleId !== undefined && (obj.roleId = message.roleId);
    message.role !== undefined &&
      (obj.role = message.role ? RoleEntity.toJSON(message.role) : undefined);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<PendingWorkspaceInviteEntity>, I>>(
    base?: I
  ): PendingWorkspaceInviteEntity {
    return PendingWorkspaceInviteEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PendingWorkspaceInviteEntity>, I>>(
    object: I
  ): PendingWorkspaceInviteEntity {
    const message = createBasePendingWorkspaceInviteEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? PendingWorkspaceInviteEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.value = object.value ?? undefined;
    message.type = object.type ?? undefined;
    message.coverLetter = object.coverLetter ?? undefined;
    message.workspaceName = object.workspaceName ?? undefined;
    message.roleId = object.roleId ?? undefined;
    message.role =
      object.role !== undefined && object.role !== null
        ? RoleEntity.fromPartial(object.role)
        : undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBasePhoneConfirmationCreateReply(): PhoneConfirmationCreateReply {
  return { data: undefined, error: undefined };
}

export const PhoneConfirmationCreateReply = {
  encode(
    message: PhoneConfirmationCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PhoneConfirmationEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PhoneConfirmationCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePhoneConfirmationCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PhoneConfirmationEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PhoneConfirmationCreateReply {
    return {
      data: isSet(object.data)
        ? PhoneConfirmationEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PhoneConfirmationCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PhoneConfirmationEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PhoneConfirmationCreateReply>, I>>(
    base?: I
  ): PhoneConfirmationCreateReply {
    return PhoneConfirmationCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PhoneConfirmationCreateReply>, I>>(
    object: I
  ): PhoneConfirmationCreateReply {
    const message = createBasePhoneConfirmationCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PhoneConfirmationEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePhoneConfirmationReply(): PhoneConfirmationReply {
  return { data: undefined, error: undefined };
}

export const PhoneConfirmationReply = {
  encode(
    message: PhoneConfirmationReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PhoneConfirmationEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PhoneConfirmationReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePhoneConfirmationReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PhoneConfirmationEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PhoneConfirmationReply {
    return {
      data: isSet(object.data)
        ? PhoneConfirmationEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PhoneConfirmationReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PhoneConfirmationEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PhoneConfirmationReply>, I>>(
    base?: I
  ): PhoneConfirmationReply {
    return PhoneConfirmationReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PhoneConfirmationReply>, I>>(
    object: I
  ): PhoneConfirmationReply {
    const message = createBasePhoneConfirmationReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PhoneConfirmationEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePhoneConfirmationQueryReply(): PhoneConfirmationQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const PhoneConfirmationQueryReply = {
  encode(
    message: PhoneConfirmationQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      PhoneConfirmationEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PhoneConfirmationQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePhoneConfirmationQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            PhoneConfirmationEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PhoneConfirmationQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => PhoneConfirmationEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PhoneConfirmationQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? PhoneConfirmationEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PhoneConfirmationQueryReply>, I>>(
    base?: I
  ): PhoneConfirmationQueryReply {
    return PhoneConfirmationQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PhoneConfirmationQueryReply>, I>>(
    object: I
  ): PhoneConfirmationQueryReply {
    const message = createBasePhoneConfirmationQueryReply();
    message.items =
      object.items?.map((e) => PhoneConfirmationEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePhoneConfirmationEntity(): PhoneConfirmationEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    user: undefined,
    status: undefined,
    phoneNumber: undefined,
    key: undefined,
    expiresAt: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const PhoneConfirmationEntity = {
  encode(
    message: PhoneConfirmationEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      PhoneConfirmationEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.user !== undefined) {
      UserEntity.encode(message.user, writer.uint32(82).fork()).ldelim();
    }
    if (message.status !== undefined) {
      writer.uint32(90).string(message.status);
    }
    if (message.phoneNumber !== undefined) {
      writer.uint32(98).string(message.phoneNumber);
    }
    if (message.key !== undefined) {
      writer.uint32(106).string(message.key);
    }
    if (message.expiresAt !== undefined) {
      writer.uint32(114).string(message.expiresAt);
    }
    if (message.rank !== 0) {
      writer.uint32(120).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(128).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(136).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(146).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(154).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PhoneConfirmationEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePhoneConfirmationEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = PhoneConfirmationEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 10:
          message.user = UserEntity.decode(reader, reader.uint32());
          break;
        case 11:
          message.status = reader.string();
          break;
        case 12:
          message.phoneNumber = reader.string();
          break;
        case 13:
          message.key = reader.string();
          break;
        case 14:
          message.expiresAt = reader.string();
          break;
        case 15:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 18:
          message.createdFormatted = reader.string();
          break;
        case 19:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PhoneConfirmationEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? PhoneConfirmationEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      user: isSet(object.user) ? UserEntity.fromJSON(object.user) : undefined,
      status: isSet(object.status) ? String(object.status) : undefined,
      phoneNumber: isSet(object.phoneNumber)
        ? String(object.phoneNumber)
        : undefined,
      key: isSet(object.key) ? String(object.key) : undefined,
      expiresAt: isSet(object.expiresAt) ? String(object.expiresAt) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: PhoneConfirmationEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? PhoneConfirmationEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.user !== undefined &&
      (obj.user = message.user ? UserEntity.toJSON(message.user) : undefined);
    message.status !== undefined && (obj.status = message.status);
    message.phoneNumber !== undefined &&
      (obj.phoneNumber = message.phoneNumber);
    message.key !== undefined && (obj.key = message.key);
    message.expiresAt !== undefined && (obj.expiresAt = message.expiresAt);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<PhoneConfirmationEntity>, I>>(
    base?: I
  ): PhoneConfirmationEntity {
    return PhoneConfirmationEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PhoneConfirmationEntity>, I>>(
    object: I
  ): PhoneConfirmationEntity {
    const message = createBasePhoneConfirmationEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? PhoneConfirmationEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.user =
      object.user !== undefined && object.user !== null
        ? UserEntity.fromPartial(object.user)
        : undefined;
    message.status = object.status ?? undefined;
    message.phoneNumber = object.phoneNumber ?? undefined;
    message.key = object.key ?? undefined;
    message.expiresAt = object.expiresAt ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBasePreferenceCreateReply(): PreferenceCreateReply {
  return { data: undefined, error: undefined };
}

export const PreferenceCreateReply = {
  encode(
    message: PreferenceCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PreferenceEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PreferenceCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePreferenceCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PreferenceEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PreferenceCreateReply {
    return {
      data: isSet(object.data)
        ? PreferenceEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PreferenceCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PreferenceEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PreferenceCreateReply>, I>>(
    base?: I
  ): PreferenceCreateReply {
    return PreferenceCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PreferenceCreateReply>, I>>(
    object: I
  ): PreferenceCreateReply {
    const message = createBasePreferenceCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PreferenceEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePreferenceReply(): PreferenceReply {
  return { data: undefined, error: undefined };
}

export const PreferenceReply = {
  encode(
    message: PreferenceReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PreferenceEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PreferenceReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePreferenceReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PreferenceEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PreferenceReply {
    return {
      data: isSet(object.data)
        ? PreferenceEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PreferenceReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PreferenceEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PreferenceReply>, I>>(
    base?: I
  ): PreferenceReply {
    return PreferenceReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PreferenceReply>, I>>(
    object: I
  ): PreferenceReply {
    const message = createBasePreferenceReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PreferenceEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePreferenceQueryReply(): PreferenceQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const PreferenceQueryReply = {
  encode(
    message: PreferenceQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      PreferenceEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PreferenceQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePreferenceQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(PreferenceEntity.decode(reader, reader.uint32()));
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PreferenceQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => PreferenceEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PreferenceQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? PreferenceEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PreferenceQueryReply>, I>>(
    base?: I
  ): PreferenceQueryReply {
    return PreferenceQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PreferenceQueryReply>, I>>(
    object: I
  ): PreferenceQueryReply {
    const message = createBasePreferenceQueryReply();
    message.items =
      object.items?.map((e) => PreferenceEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePreferenceEntity(): PreferenceEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    timezone: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const PreferenceEntity = {
  encode(
    message: PreferenceEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      PreferenceEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.timezone !== undefined) {
      writer.uint32(74).string(message.timezone);
    }
    if (message.rank !== 0) {
      writer.uint32(80).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(88).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(96).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(106).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(114).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PreferenceEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePreferenceEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = PreferenceEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.timezone = reader.string();
          break;
        case 10:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 11:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 12:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.createdFormatted = reader.string();
          break;
        case 14:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PreferenceEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? PreferenceEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      timezone: isSet(object.timezone) ? String(object.timezone) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: PreferenceEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? PreferenceEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.timezone !== undefined && (obj.timezone = message.timezone);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<PreferenceEntity>, I>>(
    base?: I
  ): PreferenceEntity {
    return PreferenceEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PreferenceEntity>, I>>(
    object: I
  ): PreferenceEntity {
    const message = createBasePreferenceEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? PreferenceEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.timezone = object.timezone ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBasePublicJoinKeyCreateReply(): PublicJoinKeyCreateReply {
  return { data: undefined, error: undefined };
}

export const PublicJoinKeyCreateReply = {
  encode(
    message: PublicJoinKeyCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PublicJoinKeyEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PublicJoinKeyCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePublicJoinKeyCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PublicJoinKeyEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PublicJoinKeyCreateReply {
    return {
      data: isSet(object.data)
        ? PublicJoinKeyEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PublicJoinKeyCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PublicJoinKeyEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PublicJoinKeyCreateReply>, I>>(
    base?: I
  ): PublicJoinKeyCreateReply {
    return PublicJoinKeyCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PublicJoinKeyCreateReply>, I>>(
    object: I
  ): PublicJoinKeyCreateReply {
    const message = createBasePublicJoinKeyCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PublicJoinKeyEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePublicJoinKeyReply(): PublicJoinKeyReply {
  return { data: undefined, error: undefined };
}

export const PublicJoinKeyReply = {
  encode(
    message: PublicJoinKeyReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      PublicJoinKeyEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PublicJoinKeyReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePublicJoinKeyReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = PublicJoinKeyEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PublicJoinKeyReply {
    return {
      data: isSet(object.data)
        ? PublicJoinKeyEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PublicJoinKeyReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? PublicJoinKeyEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PublicJoinKeyReply>, I>>(
    base?: I
  ): PublicJoinKeyReply {
    return PublicJoinKeyReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PublicJoinKeyReply>, I>>(
    object: I
  ): PublicJoinKeyReply {
    const message = createBasePublicJoinKeyReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? PublicJoinKeyEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePublicJoinKeyQueryReply(): PublicJoinKeyQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const PublicJoinKeyQueryReply = {
  encode(
    message: PublicJoinKeyQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      PublicJoinKeyEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): PublicJoinKeyQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePublicJoinKeyQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            PublicJoinKeyEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PublicJoinKeyQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => PublicJoinKeyEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: PublicJoinKeyQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? PublicJoinKeyEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<PublicJoinKeyQueryReply>, I>>(
    base?: I
  ): PublicJoinKeyQueryReply {
    return PublicJoinKeyQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PublicJoinKeyQueryReply>, I>>(
    object: I
  ): PublicJoinKeyQueryReply {
    const message = createBasePublicJoinKeyQueryReply();
    message.items =
      object.items?.map((e) => PublicJoinKeyEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBasePublicJoinKeyEntity(): PublicJoinKeyEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    roleId: undefined,
    role: undefined,
    workspace: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const PublicJoinKeyEntity = {
  encode(
    message: PublicJoinKeyEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      PublicJoinKeyEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.roleId !== undefined) {
      writer.uint32(82).string(message.roleId);
    }
    if (message.role !== undefined) {
      RoleEntity.encode(message.role, writer.uint32(90).fork()).ldelim();
    }
    if (message.workspace !== undefined) {
      WorkspaceEntity.encode(
        message.workspace,
        writer.uint32(106).fork()
      ).ldelim();
    }
    if (message.rank !== 0) {
      writer.uint32(112).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(120).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(128).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(138).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(146).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PublicJoinKeyEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePublicJoinKeyEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = PublicJoinKeyEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 10:
          message.roleId = reader.string();
          break;
        case 11:
          message.role = RoleEntity.decode(reader, reader.uint32());
          break;
        case 13:
          message.workspace = WorkspaceEntity.decode(reader, reader.uint32());
          break;
        case 14:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.createdFormatted = reader.string();
          break;
        case 18:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PublicJoinKeyEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? PublicJoinKeyEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      roleId: isSet(object.roleId) ? String(object.roleId) : undefined,
      role: isSet(object.role) ? RoleEntity.fromJSON(object.role) : undefined,
      workspace: isSet(object.workspace)
        ? WorkspaceEntity.fromJSON(object.workspace)
        : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: PublicJoinKeyEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? PublicJoinKeyEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.roleId !== undefined && (obj.roleId = message.roleId);
    message.role !== undefined &&
      (obj.role = message.role ? RoleEntity.toJSON(message.role) : undefined);
    message.workspace !== undefined &&
      (obj.workspace = message.workspace
        ? WorkspaceEntity.toJSON(message.workspace)
        : undefined);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<PublicJoinKeyEntity>, I>>(
    base?: I
  ): PublicJoinKeyEntity {
    return PublicJoinKeyEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PublicJoinKeyEntity>, I>>(
    object: I
  ): PublicJoinKeyEntity {
    const message = createBasePublicJoinKeyEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? PublicJoinKeyEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.roleId = object.roleId ?? undefined;
    message.role =
      object.role !== undefined && object.role !== null
        ? RoleEntity.fromPartial(object.role)
        : undefined;
    message.workspace =
      object.workspace !== undefined && object.workspace !== null
        ? WorkspaceEntity.fromPartial(object.workspace)
        : undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseRoleCreateReply(): RoleCreateReply {
  return { data: undefined, error: undefined };
}

export const RoleCreateReply = {
  encode(
    message: RoleCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      RoleEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RoleCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRoleCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = RoleEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RoleCreateReply {
    return {
      data: isSet(object.data) ? RoleEntity.fromJSON(object.data) : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: RoleCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data ? RoleEntity.toJSON(message.data) : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<RoleCreateReply>, I>>(
    base?: I
  ): RoleCreateReply {
    return RoleCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RoleCreateReply>, I>>(
    object: I
  ): RoleCreateReply {
    const message = createBaseRoleCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? RoleEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseRoleReply(): RoleReply {
  return { data: undefined, error: undefined };
}

export const RoleReply = {
  encode(
    message: RoleReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      RoleEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RoleReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRoleReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = RoleEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RoleReply {
    return {
      data: isSet(object.data) ? RoleEntity.fromJSON(object.data) : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: RoleReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data ? RoleEntity.toJSON(message.data) : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<RoleReply>, I>>(base?: I): RoleReply {
    return RoleReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RoleReply>, I>>(
    object: I
  ): RoleReply {
    const message = createBaseRoleReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? RoleEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseRoleQueryReply(): RoleQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const RoleQueryReply = {
  encode(
    message: RoleQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      RoleEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RoleQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRoleQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(RoleEntity.decode(reader, reader.uint32()));
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RoleQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => RoleEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: RoleQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? RoleEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<RoleQueryReply>, I>>(
    base?: I
  ): RoleQueryReply {
    return RoleQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RoleQueryReply>, I>>(
    object: I
  ): RoleQueryReply {
    const message = createBaseRoleQueryReply();
    message.items = object.items?.map((e) => RoleEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseRoleEntity(): RoleEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    name: undefined,
    capabilitiesListId: [],
    capabilities: [],
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const RoleEntity = {
  encode(
    message: RoleEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      RoleEntity.encode(message.parent, writer.uint32(42).fork()).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.name !== undefined) {
      writer.uint32(74).string(message.name);
    }
    for (const v of message.capabilitiesListId) {
      writer.uint32(90).string(v!);
    }
    for (const v of message.capabilities) {
      CapabilityEntity.encode(v!, writer.uint32(98).fork()).ldelim();
    }
    if (message.rank !== 0) {
      writer.uint32(104).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(112).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(120).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(130).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(138).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RoleEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRoleEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = RoleEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.name = reader.string();
          break;
        case 11:
          message.capabilitiesListId.push(reader.string());
          break;
        case 12:
          message.capabilities.push(
            CapabilityEntity.decode(reader, reader.uint32())
          );
          break;
        case 13:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 16:
          message.createdFormatted = reader.string();
          break;
        case 17:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RoleEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? RoleEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      name: isSet(object.name) ? String(object.name) : undefined,
      capabilitiesListId: Array.isArray(object?.capabilitiesListId)
        ? object.capabilitiesListId.map((e: any) => String(e))
        : [],
      capabilities: Array.isArray(object?.capabilities)
        ? object.capabilities.map((e: any) => CapabilityEntity.fromJSON(e))
        : [],
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: RoleEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? RoleEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.name !== undefined && (obj.name = message.name);
    if (message.capabilitiesListId) {
      obj.capabilitiesListId = message.capabilitiesListId.map((e) => e);
    } else {
      obj.capabilitiesListId = [];
    }
    if (message.capabilities) {
      obj.capabilities = message.capabilities.map((e) =>
        e ? CapabilityEntity.toJSON(e) : undefined
      );
    } else {
      obj.capabilities = [];
    }
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<RoleEntity>, I>>(base?: I): RoleEntity {
    return RoleEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RoleEntity>, I>>(
    object: I
  ): RoleEntity {
    const message = createBaseRoleEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? RoleEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.name = object.name ?? undefined;
    message.capabilitiesListId = object.capabilitiesListId?.map((e) => e) || [];
    message.capabilities =
      object.capabilities?.map((e) => CapabilityEntity.fromPartial(e)) || [];
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseTableViewSizingCreateReply(): TableViewSizingCreateReply {
  return { data: undefined, error: undefined };
}

export const TableViewSizingCreateReply = {
  encode(
    message: TableViewSizingCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      TableViewSizingEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): TableViewSizingCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTableViewSizingCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = TableViewSizingEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TableViewSizingCreateReply {
    return {
      data: isSet(object.data)
        ? TableViewSizingEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: TableViewSizingCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? TableViewSizingEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<TableViewSizingCreateReply>, I>>(
    base?: I
  ): TableViewSizingCreateReply {
    return TableViewSizingCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TableViewSizingCreateReply>, I>>(
    object: I
  ): TableViewSizingCreateReply {
    const message = createBaseTableViewSizingCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? TableViewSizingEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseTableViewSizingReply(): TableViewSizingReply {
  return { data: undefined, error: undefined };
}

export const TableViewSizingReply = {
  encode(
    message: TableViewSizingReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      TableViewSizingEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): TableViewSizingReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTableViewSizingReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = TableViewSizingEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TableViewSizingReply {
    return {
      data: isSet(object.data)
        ? TableViewSizingEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: TableViewSizingReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? TableViewSizingEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<TableViewSizingReply>, I>>(
    base?: I
  ): TableViewSizingReply {
    return TableViewSizingReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TableViewSizingReply>, I>>(
    object: I
  ): TableViewSizingReply {
    const message = createBaseTableViewSizingReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? TableViewSizingEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseTableViewSizingQueryReply(): TableViewSizingQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const TableViewSizingQueryReply = {
  encode(
    message: TableViewSizingQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      TableViewSizingEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): TableViewSizingQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTableViewSizingQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            TableViewSizingEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TableViewSizingQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => TableViewSizingEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: TableViewSizingQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? TableViewSizingEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<TableViewSizingQueryReply>, I>>(
    base?: I
  ): TableViewSizingQueryReply {
    return TableViewSizingQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TableViewSizingQueryReply>, I>>(
    object: I
  ): TableViewSizingQueryReply {
    const message = createBaseTableViewSizingQueryReply();
    message.items =
      object.items?.map((e) => TableViewSizingEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseTableViewSizingEntity(): TableViewSizingEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    tableName: undefined,
    sizes: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const TableViewSizingEntity = {
  encode(
    message: TableViewSizingEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      TableViewSizingEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.tableName !== undefined) {
      writer.uint32(74).string(message.tableName);
    }
    if (message.sizes !== undefined) {
      writer.uint32(82).string(message.sizes);
    }
    if (message.rank !== 0) {
      writer.uint32(88).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(96).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(104).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(114).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(122).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): TableViewSizingEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTableViewSizingEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = TableViewSizingEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.tableName = reader.string();
          break;
        case 10:
          message.sizes = reader.string();
          break;
        case 11:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 12:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.createdFormatted = reader.string();
          break;
        case 15:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TableViewSizingEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? TableViewSizingEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      tableName: isSet(object.tableName) ? String(object.tableName) : undefined,
      sizes: isSet(object.sizes) ? String(object.sizes) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: TableViewSizingEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? TableViewSizingEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.tableName !== undefined && (obj.tableName = message.tableName);
    message.sizes !== undefined && (obj.sizes = message.sizes);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<TableViewSizingEntity>, I>>(
    base?: I
  ): TableViewSizingEntity {
    return TableViewSizingEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TableViewSizingEntity>, I>>(
    object: I
  ): TableViewSizingEntity {
    const message = createBaseTableViewSizingEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? TableViewSizingEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.tableName = object.tableName ?? undefined;
    message.sizes = object.sizes ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseTokenCreateReply(): TokenCreateReply {
  return { data: undefined, error: undefined };
}

export const TokenCreateReply = {
  encode(
    message: TokenCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      TokenEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TokenCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTokenCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = TokenEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TokenCreateReply {
    return {
      data: isSet(object.data) ? TokenEntity.fromJSON(object.data) : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: TokenCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data ? TokenEntity.toJSON(message.data) : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<TokenCreateReply>, I>>(
    base?: I
  ): TokenCreateReply {
    return TokenCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TokenCreateReply>, I>>(
    object: I
  ): TokenCreateReply {
    const message = createBaseTokenCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? TokenEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseTokenReply(): TokenReply {
  return { data: undefined, error: undefined };
}

export const TokenReply = {
  encode(
    message: TokenReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      TokenEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TokenReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTokenReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = TokenEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TokenReply {
    return {
      data: isSet(object.data) ? TokenEntity.fromJSON(object.data) : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: TokenReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data ? TokenEntity.toJSON(message.data) : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<TokenReply>, I>>(base?: I): TokenReply {
    return TokenReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TokenReply>, I>>(
    object: I
  ): TokenReply {
    const message = createBaseTokenReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? TokenEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseTokenQueryReply(): TokenQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const TokenQueryReply = {
  encode(
    message: TokenQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      TokenEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TokenQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTokenQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(TokenEntity.decode(reader, reader.uint32()));
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TokenQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => TokenEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: TokenQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? TokenEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<TokenQueryReply>, I>>(
    base?: I
  ): TokenQueryReply {
    return TokenQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TokenQueryReply>, I>>(
    object: I
  ): TokenQueryReply {
    const message = createBaseTokenQueryReply();
    message.items = object.items?.map((e) => TokenEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseTokenEntity(): TokenEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    user: undefined,
    validUntil: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const TokenEntity = {
  encode(
    message: TokenEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      TokenEntity.encode(message.parent, writer.uint32(42).fork()).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.user !== undefined) {
      UserEntity.encode(message.user, writer.uint32(82).fork()).ldelim();
    }
    if (message.validUntil !== undefined) {
      writer.uint32(90).string(message.validUntil);
    }
    if (message.rank !== 0) {
      writer.uint32(96).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(104).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(112).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(122).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(130).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TokenEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTokenEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = TokenEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 10:
          message.user = UserEntity.decode(reader, reader.uint32());
          break;
        case 11:
          message.validUntil = reader.string();
          break;
        case 12:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.createdFormatted = reader.string();
          break;
        case 16:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TokenEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? TokenEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      user: isSet(object.user) ? UserEntity.fromJSON(object.user) : undefined,
      validUntil: isSet(object.validUntil)
        ? String(object.validUntil)
        : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: TokenEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? TokenEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.user !== undefined &&
      (obj.user = message.user ? UserEntity.toJSON(message.user) : undefined);
    message.validUntil !== undefined && (obj.validUntil = message.validUntil);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<TokenEntity>, I>>(base?: I): TokenEntity {
    return TokenEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<TokenEntity>, I>>(
    object: I
  ): TokenEntity {
    const message = createBaseTokenEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? TokenEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.user =
      object.user !== undefined && object.user !== null
        ? UserEntity.fromPartial(object.user)
        : undefined;
    message.validUntil = object.validUntil ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseUserCreateReply(): UserCreateReply {
  return { data: undefined, error: undefined };
}

export const UserCreateReply = {
  encode(
    message: UserCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      UserEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = UserEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserCreateReply {
    return {
      data: isSet(object.data) ? UserEntity.fromJSON(object.data) : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: UserCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data ? UserEntity.toJSON(message.data) : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserCreateReply>, I>>(
    base?: I
  ): UserCreateReply {
    return UserCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserCreateReply>, I>>(
    object: I
  ): UserCreateReply {
    const message = createBaseUserCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? UserEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseUserReply(): UserReply {
  return { data: undefined, error: undefined };
}

export const UserReply = {
  encode(
    message: UserReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      UserEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = UserEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserReply {
    return {
      data: isSet(object.data) ? UserEntity.fromJSON(object.data) : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: UserReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data ? UserEntity.toJSON(message.data) : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserReply>, I>>(base?: I): UserReply {
    return UserReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserReply>, I>>(
    object: I
  ): UserReply {
    const message = createBaseUserReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? UserEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseUserQueryReply(): UserQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const UserQueryReply = {
  encode(
    message: UserQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      UserEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(UserEntity.decode(reader, reader.uint32()));
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => UserEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: UserQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? UserEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserQueryReply>, I>>(
    base?: I
  ): UserQueryReply {
    return UserQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserQueryReply>, I>>(
    object: I
  ): UserQueryReply {
    const message = createBaseUserQueryReply();
    message.items = object.items?.map((e) => UserEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseUserEntity(): UserEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    firstName: undefined,
    lastName: undefined,
    photo: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const UserEntity = {
  encode(
    message: UserEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      UserEntity.encode(message.parent, writer.uint32(42).fork()).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.firstName !== undefined) {
      writer.uint32(74).string(message.firstName);
    }
    if (message.lastName !== undefined) {
      writer.uint32(82).string(message.lastName);
    }
    if (message.photo !== undefined) {
      writer.uint32(90).string(message.photo);
    }
    if (message.rank !== 0) {
      writer.uint32(96).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(104).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(112).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(122).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(130).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = UserEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.firstName = reader.string();
          break;
        case 10:
          message.lastName = reader.string();
          break;
        case 11:
          message.photo = reader.string();
          break;
        case 12:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 15:
          message.createdFormatted = reader.string();
          break;
        case 16:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? UserEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      firstName: isSet(object.firstName) ? String(object.firstName) : undefined,
      lastName: isSet(object.lastName) ? String(object.lastName) : undefined,
      photo: isSet(object.photo) ? String(object.photo) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: UserEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? UserEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    message.photo !== undefined && (obj.photo = message.photo);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserEntity>, I>>(base?: I): UserEntity {
    return UserEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserEntity>, I>>(
    object: I
  ): UserEntity {
    const message = createBaseUserEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? UserEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.firstName = object.firstName ?? undefined;
    message.lastName = object.lastName ?? undefined;
    message.photo = object.photo ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseUserProfileCreateReply(): UserProfileCreateReply {
  return { data: undefined, error: undefined };
}

export const UserProfileCreateReply = {
  encode(
    message: UserProfileCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      UserProfileEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): UserProfileCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserProfileCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = UserProfileEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserProfileCreateReply {
    return {
      data: isSet(object.data)
        ? UserProfileEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: UserProfileCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? UserProfileEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserProfileCreateReply>, I>>(
    base?: I
  ): UserProfileCreateReply {
    return UserProfileCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserProfileCreateReply>, I>>(
    object: I
  ): UserProfileCreateReply {
    const message = createBaseUserProfileCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? UserProfileEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseUserProfileReply(): UserProfileReply {
  return { data: undefined, error: undefined };
}

export const UserProfileReply = {
  encode(
    message: UserProfileReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      UserProfileEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserProfileReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserProfileReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = UserProfileEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserProfileReply {
    return {
      data: isSet(object.data)
        ? UserProfileEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: UserProfileReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? UserProfileEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserProfileReply>, I>>(
    base?: I
  ): UserProfileReply {
    return UserProfileReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserProfileReply>, I>>(
    object: I
  ): UserProfileReply {
    const message = createBaseUserProfileReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? UserProfileEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseUserProfileQueryReply(): UserProfileQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const UserProfileQueryReply = {
  encode(
    message: UserProfileQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      UserProfileEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): UserProfileQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserProfileQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(UserProfileEntity.decode(reader, reader.uint32()));
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserProfileQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => UserProfileEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: UserProfileQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? UserProfileEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserProfileQueryReply>, I>>(
    base?: I
  ): UserProfileQueryReply {
    return UserProfileQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserProfileQueryReply>, I>>(
    object: I
  ): UserProfileQueryReply {
    const message = createBaseUserProfileQueryReply();
    message.items =
      object.items?.map((e) => UserProfileEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseUserProfileEntity(): UserProfileEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    firstName: undefined,
    lastName: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const UserProfileEntity = {
  encode(
    message: UserProfileEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      UserProfileEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.firstName !== undefined) {
      writer.uint32(74).string(message.firstName);
    }
    if (message.lastName !== undefined) {
      writer.uint32(82).string(message.lastName);
    }
    if (message.rank !== 0) {
      writer.uint32(88).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(96).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(104).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(114).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(122).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserProfileEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserProfileEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = UserProfileEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.firstName = reader.string();
          break;
        case 10:
          message.lastName = reader.string();
          break;
        case 11:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 12:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.createdFormatted = reader.string();
          break;
        case 15:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserProfileEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? UserProfileEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      firstName: isSet(object.firstName) ? String(object.firstName) : undefined,
      lastName: isSet(object.lastName) ? String(object.lastName) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: UserProfileEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? UserProfileEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserProfileEntity>, I>>(
    base?: I
  ): UserProfileEntity {
    return UserProfileEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserProfileEntity>, I>>(
    object: I
  ): UserProfileEntity {
    const message = createBaseUserProfileEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? UserProfileEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.firstName = object.firstName ?? undefined;
    message.lastName = object.lastName ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseUserRoleWorkspaceCreateReply(): UserRoleWorkspaceCreateReply {
  return { data: undefined, error: undefined };
}

export const UserRoleWorkspaceCreateReply = {
  encode(
    message: UserRoleWorkspaceCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      UserRoleWorkspaceEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): UserRoleWorkspaceCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserRoleWorkspaceCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = UserRoleWorkspaceEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserRoleWorkspaceCreateReply {
    return {
      data: isSet(object.data)
        ? UserRoleWorkspaceEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: UserRoleWorkspaceCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? UserRoleWorkspaceEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserRoleWorkspaceCreateReply>, I>>(
    base?: I
  ): UserRoleWorkspaceCreateReply {
    return UserRoleWorkspaceCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserRoleWorkspaceCreateReply>, I>>(
    object: I
  ): UserRoleWorkspaceCreateReply {
    const message = createBaseUserRoleWorkspaceCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? UserRoleWorkspaceEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseUserRoleWorkspaceReply(): UserRoleWorkspaceReply {
  return { data: undefined, error: undefined };
}

export const UserRoleWorkspaceReply = {
  encode(
    message: UserRoleWorkspaceReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      UserRoleWorkspaceEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): UserRoleWorkspaceReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserRoleWorkspaceReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = UserRoleWorkspaceEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserRoleWorkspaceReply {
    return {
      data: isSet(object.data)
        ? UserRoleWorkspaceEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: UserRoleWorkspaceReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? UserRoleWorkspaceEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserRoleWorkspaceReply>, I>>(
    base?: I
  ): UserRoleWorkspaceReply {
    return UserRoleWorkspaceReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserRoleWorkspaceReply>, I>>(
    object: I
  ): UserRoleWorkspaceReply {
    const message = createBaseUserRoleWorkspaceReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? UserRoleWorkspaceEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseUserRoleWorkspaceQueryReply(): UserRoleWorkspaceQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const UserRoleWorkspaceQueryReply = {
  encode(
    message: UserRoleWorkspaceQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      UserRoleWorkspaceEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): UserRoleWorkspaceQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserRoleWorkspaceQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            UserRoleWorkspaceEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserRoleWorkspaceQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => UserRoleWorkspaceEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: UserRoleWorkspaceQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? UserRoleWorkspaceEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserRoleWorkspaceQueryReply>, I>>(
    base?: I
  ): UserRoleWorkspaceQueryReply {
    return UserRoleWorkspaceQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserRoleWorkspaceQueryReply>, I>>(
    object: I
  ): UserRoleWorkspaceQueryReply {
    const message = createBaseUserRoleWorkspaceQueryReply();
    message.items =
      object.items?.map((e) => UserRoleWorkspaceEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseUserRoleWorkspaceEntity(): UserRoleWorkspaceEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    user: undefined,
    roleId: undefined,
    role: undefined,
    workspace: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const UserRoleWorkspaceEntity = {
  encode(
    message: UserRoleWorkspaceEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      UserRoleWorkspaceEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.user !== undefined) {
      UserEntity.encode(message.user, writer.uint32(82).fork()).ldelim();
    }
    if (message.roleId !== undefined) {
      writer.uint32(98).string(message.roleId);
    }
    if (message.role !== undefined) {
      RoleEntity.encode(message.role, writer.uint32(106).fork()).ldelim();
    }
    if (message.workspace !== undefined) {
      WorkspaceEntity.encode(
        message.workspace,
        writer.uint32(122).fork()
      ).ldelim();
    }
    if (message.rank !== 0) {
      writer.uint32(128).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(136).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(144).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(154).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(162).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): UserRoleWorkspaceEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserRoleWorkspaceEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = UserRoleWorkspaceEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 10:
          message.user = UserEntity.decode(reader, reader.uint32());
          break;
        case 12:
          message.roleId = reader.string();
          break;
        case 13:
          message.role = RoleEntity.decode(reader, reader.uint32());
          break;
        case 15:
          message.workspace = WorkspaceEntity.decode(reader, reader.uint32());
          break;
        case 16:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 18:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 19:
          message.createdFormatted = reader.string();
          break;
        case 20:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserRoleWorkspaceEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? UserRoleWorkspaceEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      user: isSet(object.user) ? UserEntity.fromJSON(object.user) : undefined,
      roleId: isSet(object.roleId) ? String(object.roleId) : undefined,
      role: isSet(object.role) ? RoleEntity.fromJSON(object.role) : undefined,
      workspace: isSet(object.workspace)
        ? WorkspaceEntity.fromJSON(object.workspace)
        : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: UserRoleWorkspaceEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? UserRoleWorkspaceEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.user !== undefined &&
      (obj.user = message.user ? UserEntity.toJSON(message.user) : undefined);
    message.roleId !== undefined && (obj.roleId = message.roleId);
    message.role !== undefined &&
      (obj.role = message.role ? RoleEntity.toJSON(message.role) : undefined);
    message.workspace !== undefined &&
      (obj.workspace = message.workspace
        ? WorkspaceEntity.toJSON(message.workspace)
        : undefined);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserRoleWorkspaceEntity>, I>>(
    base?: I
  ): UserRoleWorkspaceEntity {
    return UserRoleWorkspaceEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserRoleWorkspaceEntity>, I>>(
    object: I
  ): UserRoleWorkspaceEntity {
    const message = createBaseUserRoleWorkspaceEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? UserRoleWorkspaceEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.user =
      object.user !== undefined && object.user !== null
        ? UserEntity.fromPartial(object.user)
        : undefined;
    message.roleId = object.roleId ?? undefined;
    message.role =
      object.role !== undefined && object.role !== null
        ? RoleEntity.fromPartial(object.role)
        : undefined;
    message.workspace =
      object.workspace !== undefined && object.workspace !== null
        ? WorkspaceEntity.fromPartial(object.workspace)
        : undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseAcceptInviteDto(): AcceptInviteDto {
  return { inviteUniqueId: "", visibility: undefined, updated: 0, created: 0 };
}

export const AcceptInviteDto = {
  encode(
    message: AcceptInviteDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.inviteUniqueId !== "") {
      writer.uint32(10).string(message.inviteUniqueId);
    }
    if (message.visibility !== undefined) {
      writer.uint32(18).string(message.visibility);
    }
    if (message.updated !== 0) {
      writer.uint32(128).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(136).int64(message.created);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AcceptInviteDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAcceptInviteDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.inviteUniqueId = reader.string();
          break;
        case 2:
          message.visibility = reader.string();
          break;
        case 16:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.created = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AcceptInviteDto {
    return {
      inviteUniqueId: isSet(object.inviteUniqueId)
        ? String(object.inviteUniqueId)
        : "",
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
    };
  },

  toJSON(message: AcceptInviteDto): unknown {
    const obj: any = {};
    message.inviteUniqueId !== undefined &&
      (obj.inviteUniqueId = message.inviteUniqueId);
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    return obj;
  },

  create<I extends Exact<DeepPartial<AcceptInviteDto>, I>>(
    base?: I
  ): AcceptInviteDto {
    return AcceptInviteDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AcceptInviteDto>, I>>(
    object: I
  ): AcceptInviteDto {
    const message = createBaseAcceptInviteDto();
    message.inviteUniqueId = object.inviteUniqueId ?? "";
    message.visibility = object.visibility ?? undefined;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    return message;
  },
};

function createBaseAssignRoleDto(): AssignRoleDto {
  return {
    roleId: "",
    userId: "",
    visibility: undefined,
    updated: 0,
    created: 0,
  };
}

export const AssignRoleDto = {
  encode(
    message: AssignRoleDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.roleId !== "") {
      writer.uint32(10).string(message.roleId);
    }
    if (message.userId !== "") {
      writer.uint32(18).string(message.userId);
    }
    if (message.visibility !== undefined) {
      writer.uint32(26).string(message.visibility);
    }
    if (message.updated !== 0) {
      writer.uint32(128).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(136).int64(message.created);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AssignRoleDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAssignRoleDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.roleId = reader.string();
          break;
        case 2:
          message.userId = reader.string();
          break;
        case 3:
          message.visibility = reader.string();
          break;
        case 16:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.created = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AssignRoleDto {
    return {
      roleId: isSet(object.roleId) ? String(object.roleId) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
    };
  },

  toJSON(message: AssignRoleDto): unknown {
    const obj: any = {};
    message.roleId !== undefined && (obj.roleId = message.roleId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    return obj;
  },

  create<I extends Exact<DeepPartial<AssignRoleDto>, I>>(
    base?: I
  ): AssignRoleDto {
    return AssignRoleDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AssignRoleDto>, I>>(
    object: I
  ): AssignRoleDto {
    const message = createBaseAssignRoleDto();
    message.roleId = object.roleId ?? "";
    message.userId = object.userId ?? "";
    message.visibility = object.visibility ?? undefined;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    return message;
  },
};

function createBaseWorkspaceDto(): WorkspaceDto {
  return { relations: [], visibility: undefined, updated: 0, created: 0 };
}

export const WorkspaceDto = {
  encode(
    message: WorkspaceDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.relations) {
      UserRoleWorkspaceEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.visibility !== undefined) {
      writer.uint32(18).string(message.visibility);
    }
    if (message.updated !== 0) {
      writer.uint32(128).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(136).int64(message.created);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WorkspaceDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.relations.push(
            UserRoleWorkspaceEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.visibility = reader.string();
          break;
        case 16:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.created = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceDto {
    return {
      relations: Array.isArray(object?.relations)
        ? object.relations.map((e: any) => UserRoleWorkspaceEntity.fromJSON(e))
        : [],
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
    };
  },

  toJSON(message: WorkspaceDto): unknown {
    const obj: any = {};
    if (message.relations) {
      obj.relations = message.relations.map((e) =>
        e ? UserRoleWorkspaceEntity.toJSON(e) : undefined
      );
    } else {
      obj.relations = [];
    }
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceDto>, I>>(
    base?: I
  ): WorkspaceDto {
    return WorkspaceDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceDto>, I>>(
    object: I
  ): WorkspaceDto {
    const message = createBaseWorkspaceDto();
    message.relations =
      object.relations?.map((e) => UserRoleWorkspaceEntity.fromPartial(e)) ||
      [];
    message.visibility = object.visibility ?? undefined;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    return message;
  },
};

function createBaseExchangeKeyInformationDto(): ExchangeKeyInformationDto {
  return { key: "", visibility: undefined };
}

export const ExchangeKeyInformationDto = {
  encode(
    message: ExchangeKeyInformationDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.visibility !== undefined) {
      writer.uint32(18).string(message.visibility);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ExchangeKeyInformationDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseExchangeKeyInformationDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.visibility = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ExchangeKeyInformationDto {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
    };
  },

  toJSON(message: ExchangeKeyInformationDto): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.visibility !== undefined && (obj.visibility = message.visibility);
    return obj;
  },

  create<I extends Exact<DeepPartial<ExchangeKeyInformationDto>, I>>(
    base?: I
  ): ExchangeKeyInformationDto {
    return ExchangeKeyInformationDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ExchangeKeyInformationDto>, I>>(
    object: I
  ): ExchangeKeyInformationDto {
    const message = createBaseExchangeKeyInformationDto();
    message.key = object.key ?? "";
    message.visibility = object.visibility ?? undefined;
    return message;
  },
};

function createBaseUserAccessLevel(): UserAccessLevel {
  return { capabilities: [], workspaces: [], SQL: "" };
}

export const UserAccessLevel = {
  encode(
    message: UserAccessLevel,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.capabilities) {
      writer.uint32(10).string(v!);
    }
    for (const v of message.workspaces) {
      writer.uint32(18).string(v!);
    }
    if (message.SQL !== "") {
      writer.uint32(26).string(message.SQL);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserAccessLevel {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserAccessLevel();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.capabilities.push(reader.string());
          break;
        case 2:
          message.workspaces.push(reader.string());
          break;
        case 3:
          message.SQL = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserAccessLevel {
    return {
      capabilities: Array.isArray(object?.capabilities)
        ? object.capabilities.map((e: any) => String(e))
        : [],
      workspaces: Array.isArray(object?.workspaces)
        ? object.workspaces.map((e: any) => String(e))
        : [],
      SQL: isSet(object.SQL) ? String(object.SQL) : "",
    };
  },

  toJSON(message: UserAccessLevel): unknown {
    const obj: any = {};
    if (message.capabilities) {
      obj.capabilities = message.capabilities.map((e) => e);
    } else {
      obj.capabilities = [];
    }
    if (message.workspaces) {
      obj.workspaces = message.workspaces.map((e) => e);
    } else {
      obj.workspaces = [];
    }
    message.SQL !== undefined && (obj.SQL = message.SQL);
    return obj;
  },

  create<I extends Exact<DeepPartial<UserAccessLevel>, I>>(
    base?: I
  ): UserAccessLevel {
    return UserAccessLevel.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UserAccessLevel>, I>>(
    object: I
  ): UserAccessLevel {
    const message = createBaseUserAccessLevel();
    message.capabilities = object.capabilities?.map((e) => e) || [];
    message.workspaces = object.workspaces?.map((e) => e) || [];
    message.SQL = object.SQL ?? "";
    return message;
  },
};

function createBaseAuthResult(): AuthResult {
  return {
    workspaceId: "",
    internalSql: "",
    userId: "",
    user: undefined,
    accessLevel: undefined,
    userHas: [],
  };
}

export const AuthResult = {
  encode(
    message: AuthResult,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.workspaceId !== "") {
      writer.uint32(10).string(message.workspaceId);
    }
    if (message.internalSql !== "") {
      writer.uint32(18).string(message.internalSql);
    }
    if (message.userId !== "") {
      writer.uint32(26).string(message.userId);
    }
    if (message.user !== undefined) {
      UserEntity.encode(message.user, writer.uint32(34).fork()).ldelim();
    }
    if (message.accessLevel !== undefined) {
      UserAccessLevel.encode(
        message.accessLevel,
        writer.uint32(42).fork()
      ).ldelim();
    }
    for (const v of message.userHas) {
      writer.uint32(50).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AuthResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAuthResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.workspaceId = reader.string();
          break;
        case 2:
          message.internalSql = reader.string();
          break;
        case 3:
          message.userId = reader.string();
          break;
        case 4:
          message.user = UserEntity.decode(reader, reader.uint32());
          break;
        case 5:
          message.accessLevel = UserAccessLevel.decode(reader, reader.uint32());
          break;
        case 6:
          message.userHas.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AuthResult {
    return {
      workspaceId: isSet(object.workspaceId) ? String(object.workspaceId) : "",
      internalSql: isSet(object.internalSql) ? String(object.internalSql) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
      user: isSet(object.user) ? UserEntity.fromJSON(object.user) : undefined,
      accessLevel: isSet(object.accessLevel)
        ? UserAccessLevel.fromJSON(object.accessLevel)
        : undefined,
      userHas: Array.isArray(object?.userHas)
        ? object.userHas.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: AuthResult): unknown {
    const obj: any = {};
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.internalSql !== undefined &&
      (obj.internalSql = message.internalSql);
    message.userId !== undefined && (obj.userId = message.userId);
    message.user !== undefined &&
      (obj.user = message.user ? UserEntity.toJSON(message.user) : undefined);
    message.accessLevel !== undefined &&
      (obj.accessLevel = message.accessLevel
        ? UserAccessLevel.toJSON(message.accessLevel)
        : undefined);
    if (message.userHas) {
      obj.userHas = message.userHas.map((e) => e);
    } else {
      obj.userHas = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AuthResult>, I>>(base?: I): AuthResult {
    return AuthResult.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AuthResult>, I>>(
    object: I
  ): AuthResult {
    const message = createBaseAuthResult();
    message.workspaceId = object.workspaceId ?? "";
    message.internalSql = object.internalSql ?? "";
    message.userId = object.userId ?? "";
    message.user =
      object.user !== undefined && object.user !== null
        ? UserEntity.fromPartial(object.user)
        : undefined;
    message.accessLevel =
      object.accessLevel !== undefined && object.accessLevel !== null
        ? UserAccessLevel.fromPartial(object.accessLevel)
        : undefined;
    message.userHas = object.userHas?.map((e) => e) || [];
    return message;
  },
};

function createBaseAuthContext(): AuthContext {
  return {
    skipWorkspaceId: false,
    workspaceId: "",
    token: "",
    capabilities: [],
  };
}

export const AuthContext = {
  encode(
    message: AuthContext,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.skipWorkspaceId === true) {
      writer.uint32(32).bool(message.skipWorkspaceId);
    }
    if (message.workspaceId !== "") {
      writer.uint32(10).string(message.workspaceId);
    }
    if (message.token !== "") {
      writer.uint32(18).string(message.token);
    }
    for (const v of message.capabilities) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AuthContext {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAuthContext();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 4:
          message.skipWorkspaceId = reader.bool();
          break;
        case 1:
          message.workspaceId = reader.string();
          break;
        case 2:
          message.token = reader.string();
          break;
        case 3:
          message.capabilities.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AuthContext {
    return {
      skipWorkspaceId: isSet(object.skipWorkspaceId)
        ? Boolean(object.skipWorkspaceId)
        : false,
      workspaceId: isSet(object.workspaceId) ? String(object.workspaceId) : "",
      token: isSet(object.token) ? String(object.token) : "",
      capabilities: Array.isArray(object?.capabilities)
        ? object.capabilities.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: AuthContext): unknown {
    const obj: any = {};
    message.skipWorkspaceId !== undefined &&
      (obj.skipWorkspaceId = message.skipWorkspaceId);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.token !== undefined && (obj.token = message.token);
    if (message.capabilities) {
      obj.capabilities = message.capabilities.map((e) => e);
    } else {
      obj.capabilities = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AuthContext>, I>>(base?: I): AuthContext {
    return AuthContext.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AuthContext>, I>>(
    object: I
  ): AuthContext {
    const message = createBaseAuthContext();
    message.skipWorkspaceId = object.skipWorkspaceId ?? false;
    message.workspaceId = object.workspaceId ?? "";
    message.token = object.token ?? "";
    message.capabilities = object.capabilities?.map((e) => e) || [];
    return message;
  },
};

function createBaseReactiveSearchResultDto(): ReactiveSearchResultDto {
  return {
    parent: undefined,
    uniqueId: "",
    phrase: undefined,
    icon: undefined,
    description: undefined,
    group: undefined,
    uiLocation: undefined,
    actionFn: undefined,
  };
}

export const ReactiveSearchResultDto = {
  encode(
    message: ReactiveSearchResultDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.parent !== undefined) {
      ReactiveSearchResultDto.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.phrase !== undefined) {
      writer.uint32(74).string(message.phrase);
    }
    if (message.icon !== undefined) {
      writer.uint32(82).string(message.icon);
    }
    if (message.description !== undefined) {
      writer.uint32(90).string(message.description);
    }
    if (message.group !== undefined) {
      writer.uint32(98).string(message.group);
    }
    if (message.uiLocation !== undefined) {
      writer.uint32(106).string(message.uiLocation);
    }
    if (message.actionFn !== undefined) {
      writer.uint32(114).string(message.actionFn);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ReactiveSearchResultDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReactiveSearchResultDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 5:
          message.parent = ReactiveSearchResultDto.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 9:
          message.phrase = reader.string();
          break;
        case 10:
          message.icon = reader.string();
          break;
        case 11:
          message.description = reader.string();
          break;
        case 12:
          message.group = reader.string();
          break;
        case 13:
          message.uiLocation = reader.string();
          break;
        case 14:
          message.actionFn = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ReactiveSearchResultDto {
    return {
      parent: isSet(object.parent)
        ? ReactiveSearchResultDto.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      phrase: isSet(object.phrase) ? String(object.phrase) : undefined,
      icon: isSet(object.icon) ? String(object.icon) : undefined,
      description: isSet(object.description)
        ? String(object.description)
        : undefined,
      group: isSet(object.group) ? String(object.group) : undefined,
      uiLocation: isSet(object.uiLocation)
        ? String(object.uiLocation)
        : undefined,
      actionFn: isSet(object.actionFn) ? String(object.actionFn) : undefined,
    };
  },

  toJSON(message: ReactiveSearchResultDto): unknown {
    const obj: any = {};
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? ReactiveSearchResultDto.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.phrase !== undefined && (obj.phrase = message.phrase);
    message.icon !== undefined && (obj.icon = message.icon);
    message.description !== undefined &&
      (obj.description = message.description);
    message.group !== undefined && (obj.group = message.group);
    message.uiLocation !== undefined && (obj.uiLocation = message.uiLocation);
    message.actionFn !== undefined && (obj.actionFn = message.actionFn);
    return obj;
  },

  create<I extends Exact<DeepPartial<ReactiveSearchResultDto>, I>>(
    base?: I
  ): ReactiveSearchResultDto {
    return ReactiveSearchResultDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReactiveSearchResultDto>, I>>(
    object: I
  ): ReactiveSearchResultDto {
    const message = createBaseReactiveSearchResultDto();
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? ReactiveSearchResultDto.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.phrase = object.phrase ?? undefined;
    message.icon = object.icon ?? undefined;
    message.description = object.description ?? undefined;
    message.group = object.group ?? undefined;
    message.uiLocation = object.uiLocation ?? undefined;
    message.actionFn = object.actionFn ?? undefined;
    return message;
  },
};

function createBaseWorkspaceConfigCreateReply(): WorkspaceConfigCreateReply {
  return { data: undefined, error: undefined };
}

export const WorkspaceConfigCreateReply = {
  encode(
    message: WorkspaceConfigCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      WorkspaceConfigEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WorkspaceConfigCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceConfigCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = WorkspaceConfigEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceConfigCreateReply {
    return {
      data: isSet(object.data)
        ? WorkspaceConfigEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WorkspaceConfigCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? WorkspaceConfigEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceConfigCreateReply>, I>>(
    base?: I
  ): WorkspaceConfigCreateReply {
    return WorkspaceConfigCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceConfigCreateReply>, I>>(
    object: I
  ): WorkspaceConfigCreateReply {
    const message = createBaseWorkspaceConfigCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? WorkspaceConfigEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWorkspaceConfigReply(): WorkspaceConfigReply {
  return { data: undefined, error: undefined };
}

export const WorkspaceConfigReply = {
  encode(
    message: WorkspaceConfigReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      WorkspaceConfigEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WorkspaceConfigReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceConfigReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = WorkspaceConfigEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceConfigReply {
    return {
      data: isSet(object.data)
        ? WorkspaceConfigEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WorkspaceConfigReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? WorkspaceConfigEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceConfigReply>, I>>(
    base?: I
  ): WorkspaceConfigReply {
    return WorkspaceConfigReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceConfigReply>, I>>(
    object: I
  ): WorkspaceConfigReply {
    const message = createBaseWorkspaceConfigReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? WorkspaceConfigEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWorkspaceConfigQueryReply(): WorkspaceConfigQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const WorkspaceConfigQueryReply = {
  encode(
    message: WorkspaceConfigQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      WorkspaceConfigEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WorkspaceConfigQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceConfigQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            WorkspaceConfigEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceConfigQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => WorkspaceConfigEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WorkspaceConfigQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? WorkspaceConfigEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceConfigQueryReply>, I>>(
    base?: I
  ): WorkspaceConfigQueryReply {
    return WorkspaceConfigQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceConfigQueryReply>, I>>(
    object: I
  ): WorkspaceConfigQueryReply {
    const message = createBaseWorkspaceConfigQueryReply();
    message.items =
      object.items?.map((e) => WorkspaceConfigEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWorkspaceConfigEntity(): WorkspaceConfigEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    disablePublicWorkspaceCreation: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const WorkspaceConfigEntity = {
  encode(
    message: WorkspaceConfigEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      WorkspaceConfigEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.disablePublicWorkspaceCreation !== undefined) {
      writer.uint32(72).int64(message.disablePublicWorkspaceCreation);
    }
    if (message.rank !== 0) {
      writer.uint32(80).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(88).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(96).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(106).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(114).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WorkspaceConfigEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceConfigEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = WorkspaceConfigEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.disablePublicWorkspaceCreation = longToNumber(
            reader.int64() as Long
          );
          break;
        case 10:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 11:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 12:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.createdFormatted = reader.string();
          break;
        case 14:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceConfigEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? WorkspaceConfigEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      disablePublicWorkspaceCreation: isSet(
        object.disablePublicWorkspaceCreation
      )
        ? Number(object.disablePublicWorkspaceCreation)
        : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: WorkspaceConfigEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? WorkspaceConfigEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.disablePublicWorkspaceCreation !== undefined &&
      (obj.disablePublicWorkspaceCreation = Math.round(
        message.disablePublicWorkspaceCreation
      ));
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceConfigEntity>, I>>(
    base?: I
  ): WorkspaceConfigEntity {
    return WorkspaceConfigEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceConfigEntity>, I>>(
    object: I
  ): WorkspaceConfigEntity {
    const message = createBaseWorkspaceConfigEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? WorkspaceConfigEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.disablePublicWorkspaceCreation =
      object.disablePublicWorkspaceCreation ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseWorkspaceConfigDto(): WorkspaceConfigDto {
  return {
    workspace: undefined,
    workspaceId: undefined,
    zoomClientId: undefined,
    zoomClientSecret: undefined,
    allowPublicToJoinTheWorkspace: undefined,
    visibility: undefined,
    updated: 0,
    created: 0,
  };
}

export const WorkspaceConfigDto = {
  encode(
    message: WorkspaceConfigDto,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.workspace !== undefined) {
      WorkspaceEntity.encode(
        message.workspace,
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(26).string(message.workspaceId);
    }
    if (message.zoomClientId !== undefined) {
      writer.uint32(34).string(message.zoomClientId);
    }
    if (message.zoomClientSecret !== undefined) {
      writer.uint32(42).string(message.zoomClientSecret);
    }
    if (message.allowPublicToJoinTheWorkspace !== undefined) {
      writer.uint32(48).bool(message.allowPublicToJoinTheWorkspace);
    }
    if (message.visibility !== undefined) {
      writer.uint32(58).string(message.visibility);
    }
    if (message.updated !== 0) {
      writer.uint32(128).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(136).int64(message.created);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WorkspaceConfigDto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceConfigDto();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.workspace = WorkspaceEntity.decode(reader, reader.uint32());
          break;
        case 3:
          message.workspaceId = reader.string();
          break;
        case 4:
          message.zoomClientId = reader.string();
          break;
        case 5:
          message.zoomClientSecret = reader.string();
          break;
        case 6:
          message.allowPublicToJoinTheWorkspace = reader.bool();
          break;
        case 7:
          message.visibility = reader.string();
          break;
        case 16:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.created = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceConfigDto {
    return {
      workspace: isSet(object.workspace)
        ? WorkspaceEntity.fromJSON(object.workspace)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      zoomClientId: isSet(object.zoomClientId)
        ? String(object.zoomClientId)
        : undefined,
      zoomClientSecret: isSet(object.zoomClientSecret)
        ? String(object.zoomClientSecret)
        : undefined,
      allowPublicToJoinTheWorkspace: isSet(object.allowPublicToJoinTheWorkspace)
        ? Boolean(object.allowPublicToJoinTheWorkspace)
        : undefined,
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
    };
  },

  toJSON(message: WorkspaceConfigDto): unknown {
    const obj: any = {};
    message.workspace !== undefined &&
      (obj.workspace = message.workspace
        ? WorkspaceEntity.toJSON(message.workspace)
        : undefined);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.zoomClientId !== undefined &&
      (obj.zoomClientId = message.zoomClientId);
    message.zoomClientSecret !== undefined &&
      (obj.zoomClientSecret = message.zoomClientSecret);
    message.allowPublicToJoinTheWorkspace !== undefined &&
      (obj.allowPublicToJoinTheWorkspace =
        message.allowPublicToJoinTheWorkspace);
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceConfigDto>, I>>(
    base?: I
  ): WorkspaceConfigDto {
    return WorkspaceConfigDto.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceConfigDto>, I>>(
    object: I
  ): WorkspaceConfigDto {
    const message = createBaseWorkspaceConfigDto();
    message.workspace =
      object.workspace !== undefined && object.workspace !== null
        ? WorkspaceEntity.fromPartial(object.workspace)
        : undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.zoomClientId = object.zoomClientId ?? undefined;
    message.zoomClientSecret = object.zoomClientSecret ?? undefined;
    message.allowPublicToJoinTheWorkspace =
      object.allowPublicToJoinTheWorkspace ?? undefined;
    message.visibility = object.visibility ?? undefined;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    return message;
  },
};

function createBaseWorkspaceCreateReply(): WorkspaceCreateReply {
  return { data: undefined, error: undefined };
}

export const WorkspaceCreateReply = {
  encode(
    message: WorkspaceCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      WorkspaceEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WorkspaceCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = WorkspaceEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceCreateReply {
    return {
      data: isSet(object.data)
        ? WorkspaceEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WorkspaceCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? WorkspaceEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceCreateReply>, I>>(
    base?: I
  ): WorkspaceCreateReply {
    return WorkspaceCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceCreateReply>, I>>(
    object: I
  ): WorkspaceCreateReply {
    const message = createBaseWorkspaceCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? WorkspaceEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWorkspaceReply(): WorkspaceReply {
  return { data: undefined, error: undefined };
}

export const WorkspaceReply = {
  encode(
    message: WorkspaceReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      WorkspaceEntity.encode(message.data, writer.uint32(10).fork()).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WorkspaceReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = WorkspaceEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceReply {
    return {
      data: isSet(object.data)
        ? WorkspaceEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WorkspaceReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? WorkspaceEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceReply>, I>>(
    base?: I
  ): WorkspaceReply {
    return WorkspaceReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceReply>, I>>(
    object: I
  ): WorkspaceReply {
    const message = createBaseWorkspaceReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? WorkspaceEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWorkspaceQueryReply(): WorkspaceQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const WorkspaceQueryReply = {
  encode(
    message: WorkspaceQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      WorkspaceEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WorkspaceQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(WorkspaceEntity.decode(reader, reader.uint32()));
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => WorkspaceEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WorkspaceQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? WorkspaceEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceQueryReply>, I>>(
    base?: I
  ): WorkspaceQueryReply {
    return WorkspaceQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceQueryReply>, I>>(
    object: I
  ): WorkspaceQueryReply {
    const message = createBaseWorkspaceQueryReply();
    message.items =
      object.items?.map((e) => WorkspaceEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWorkspaceEntity(): WorkspaceEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    description: undefined,
    name: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
    children: [],
  };
}

export const WorkspaceEntity = {
  encode(
    message: WorkspaceEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      WorkspaceEntity.encode(message.parent, writer.uint32(42).fork()).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.description !== undefined) {
      writer.uint32(74).string(message.description);
    }
    if (message.name !== undefined) {
      writer.uint32(82).string(message.name);
    }
    if (message.rank !== 0) {
      writer.uint32(88).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(96).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(104).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(114).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(122).string(message.updatedFormatted);
    }
    for (const v of message.children) {
      WorkspaceEntity.encode(v!, writer.uint32(130).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WorkspaceEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = WorkspaceEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.description = reader.string();
          break;
        case 10:
          message.name = reader.string();
          break;
        case 11:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 12:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 13:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 14:
          message.createdFormatted = reader.string();
          break;
        case 15:
          message.updatedFormatted = reader.string();
          break;
        case 16:
          message.children.push(
            WorkspaceEntity.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? WorkspaceEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      description: isSet(object.description)
        ? String(object.description)
        : undefined,
      name: isSet(object.name) ? String(object.name) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
      children: Array.isArray(object?.children)
        ? object.children.map((e: any) => WorkspaceEntity.fromJSON(e))
        : [],
    };
  },

  toJSON(message: WorkspaceEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? WorkspaceEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.description !== undefined &&
      (obj.description = message.description);
    message.name !== undefined && (obj.name = message.name);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    if (message.children) {
      obj.children = message.children.map((e) =>
        e ? WorkspaceEntity.toJSON(e) : undefined
      );
    } else {
      obj.children = [];
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceEntity>, I>>(
    base?: I
  ): WorkspaceEntity {
    return WorkspaceEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceEntity>, I>>(
    object: I
  ): WorkspaceEntity {
    const message = createBaseWorkspaceEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? WorkspaceEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.description = object.description ?? undefined;
    message.name = object.name ?? undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    message.children =
      object.children?.map((e) => WorkspaceEntity.fromPartial(e)) || [];
    return message;
  },
};

function createBaseWorkspaceInviteCreateReply(): WorkspaceInviteCreateReply {
  return { data: undefined, error: undefined };
}

export const WorkspaceInviteCreateReply = {
  encode(
    message: WorkspaceInviteCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      WorkspaceInviteEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WorkspaceInviteCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceInviteCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = WorkspaceInviteEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceInviteCreateReply {
    return {
      data: isSet(object.data)
        ? WorkspaceInviteEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WorkspaceInviteCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? WorkspaceInviteEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceInviteCreateReply>, I>>(
    base?: I
  ): WorkspaceInviteCreateReply {
    return WorkspaceInviteCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceInviteCreateReply>, I>>(
    object: I
  ): WorkspaceInviteCreateReply {
    const message = createBaseWorkspaceInviteCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? WorkspaceInviteEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWorkspaceInviteReply(): WorkspaceInviteReply {
  return { data: undefined, error: undefined };
}

export const WorkspaceInviteReply = {
  encode(
    message: WorkspaceInviteReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      WorkspaceInviteEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WorkspaceInviteReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceInviteReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = WorkspaceInviteEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceInviteReply {
    return {
      data: isSet(object.data)
        ? WorkspaceInviteEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WorkspaceInviteReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? WorkspaceInviteEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceInviteReply>, I>>(
    base?: I
  ): WorkspaceInviteReply {
    return WorkspaceInviteReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceInviteReply>, I>>(
    object: I
  ): WorkspaceInviteReply {
    const message = createBaseWorkspaceInviteReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? WorkspaceInviteEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWorkspaceInviteQueryReply(): WorkspaceInviteQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const WorkspaceInviteQueryReply = {
  encode(
    message: WorkspaceInviteQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      WorkspaceInviteEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WorkspaceInviteQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceInviteQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            WorkspaceInviteEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceInviteQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => WorkspaceInviteEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WorkspaceInviteQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? WorkspaceInviteEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceInviteQueryReply>, I>>(
    base?: I
  ): WorkspaceInviteQueryReply {
    return WorkspaceInviteQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceInviteQueryReply>, I>>(
    object: I
  ): WorkspaceInviteQueryReply {
    const message = createBaseWorkspaceInviteQueryReply();
    message.items =
      object.items?.map((e) => WorkspaceInviteEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWorkspaceInviteEntity(): WorkspaceInviteEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    passportMode: undefined,
    coverLetter: undefined,
    targetUserLocale: undefined,
    email: undefined,
    workspace: undefined,
    firstName: undefined,
    lastName: undefined,
    inviteeUserId: undefined,
    phoneNumber: undefined,
    roleId: undefined,
    role: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const WorkspaceInviteEntity = {
  encode(
    message: WorkspaceInviteEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      WorkspaceInviteEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    if (message.passportMode !== undefined) {
      writer.uint32(74).string(message.passportMode);
    }
    if (message.coverLetter !== undefined) {
      writer.uint32(82).string(message.coverLetter);
    }
    if (message.targetUserLocale !== undefined) {
      writer.uint32(90).string(message.targetUserLocale);
    }
    if (message.email !== undefined) {
      writer.uint32(98).string(message.email);
    }
    if (message.workspace !== undefined) {
      WorkspaceEntity.encode(
        message.workspace,
        writer.uint32(114).fork()
      ).ldelim();
    }
    if (message.firstName !== undefined) {
      writer.uint32(122).string(message.firstName);
    }
    if (message.lastName !== undefined) {
      writer.uint32(130).string(message.lastName);
    }
    if (message.inviteeUserId !== undefined) {
      writer.uint32(138).string(message.inviteeUserId);
    }
    if (message.phoneNumber !== undefined) {
      writer.uint32(146).string(message.phoneNumber);
    }
    if (message.roleId !== undefined) {
      writer.uint32(170).string(message.roleId);
    }
    if (message.role !== undefined) {
      RoleEntity.encode(message.role, writer.uint32(178).fork()).ldelim();
    }
    if (message.rank !== 0) {
      writer.uint32(184).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(192).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(200).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(210).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(218).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WorkspaceInviteEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceInviteEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = WorkspaceInviteEntity.decode(
            reader,
            reader.uint32()
          );
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 9:
          message.passportMode = reader.string();
          break;
        case 10:
          message.coverLetter = reader.string();
          break;
        case 11:
          message.targetUserLocale = reader.string();
          break;
        case 12:
          message.email = reader.string();
          break;
        case 14:
          message.workspace = WorkspaceEntity.decode(reader, reader.uint32());
          break;
        case 15:
          message.firstName = reader.string();
          break;
        case 16:
          message.lastName = reader.string();
          break;
        case 17:
          message.inviteeUserId = reader.string();
          break;
        case 18:
          message.phoneNumber = reader.string();
          break;
        case 21:
          message.roleId = reader.string();
          break;
        case 22:
          message.role = RoleEntity.decode(reader, reader.uint32());
          break;
        case 23:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 24:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 25:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 26:
          message.createdFormatted = reader.string();
          break;
        case 27:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceInviteEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? WorkspaceInviteEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      passportMode: isSet(object.passportMode)
        ? String(object.passportMode)
        : undefined,
      coverLetter: isSet(object.coverLetter)
        ? String(object.coverLetter)
        : undefined,
      targetUserLocale: isSet(object.targetUserLocale)
        ? String(object.targetUserLocale)
        : undefined,
      email: isSet(object.email) ? String(object.email) : undefined,
      workspace: isSet(object.workspace)
        ? WorkspaceEntity.fromJSON(object.workspace)
        : undefined,
      firstName: isSet(object.firstName) ? String(object.firstName) : undefined,
      lastName: isSet(object.lastName) ? String(object.lastName) : undefined,
      inviteeUserId: isSet(object.inviteeUserId)
        ? String(object.inviteeUserId)
        : undefined,
      phoneNumber: isSet(object.phoneNumber)
        ? String(object.phoneNumber)
        : undefined,
      roleId: isSet(object.roleId) ? String(object.roleId) : undefined,
      role: isSet(object.role) ? RoleEntity.fromJSON(object.role) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: WorkspaceInviteEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? WorkspaceInviteEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.passportMode !== undefined &&
      (obj.passportMode = message.passportMode);
    message.coverLetter !== undefined &&
      (obj.coverLetter = message.coverLetter);
    message.targetUserLocale !== undefined &&
      (obj.targetUserLocale = message.targetUserLocale);
    message.email !== undefined && (obj.email = message.email);
    message.workspace !== undefined &&
      (obj.workspace = message.workspace
        ? WorkspaceEntity.toJSON(message.workspace)
        : undefined);
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    message.inviteeUserId !== undefined &&
      (obj.inviteeUserId = message.inviteeUserId);
    message.phoneNumber !== undefined &&
      (obj.phoneNumber = message.phoneNumber);
    message.roleId !== undefined && (obj.roleId = message.roleId);
    message.role !== undefined &&
      (obj.role = message.role ? RoleEntity.toJSON(message.role) : undefined);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceInviteEntity>, I>>(
    base?: I
  ): WorkspaceInviteEntity {
    return WorkspaceInviteEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceInviteEntity>, I>>(
    object: I
  ): WorkspaceInviteEntity {
    const message = createBaseWorkspaceInviteEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? WorkspaceInviteEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.passportMode = object.passportMode ?? undefined;
    message.coverLetter = object.coverLetter ?? undefined;
    message.targetUserLocale = object.targetUserLocale ?? undefined;
    message.email = object.email ?? undefined;
    message.workspace =
      object.workspace !== undefined && object.workspace !== null
        ? WorkspaceEntity.fromPartial(object.workspace)
        : undefined;
    message.firstName = object.firstName ?? undefined;
    message.lastName = object.lastName ?? undefined;
    message.inviteeUserId = object.inviteeUserId ?? undefined;
    message.phoneNumber = object.phoneNumber ?? undefined;
    message.roleId = object.roleId ?? undefined;
    message.role =
      object.role !== undefined && object.role !== null
        ? RoleEntity.fromPartial(object.role)
        : undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseWorkspaceTypeCreateReply(): WorkspaceTypeCreateReply {
  return { data: undefined, error: undefined };
}

export const WorkspaceTypeCreateReply = {
  encode(
    message: WorkspaceTypeCreateReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      WorkspaceTypeEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WorkspaceTypeCreateReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceTypeCreateReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = WorkspaceTypeEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceTypeCreateReply {
    return {
      data: isSet(object.data)
        ? WorkspaceTypeEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WorkspaceTypeCreateReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? WorkspaceTypeEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceTypeCreateReply>, I>>(
    base?: I
  ): WorkspaceTypeCreateReply {
    return WorkspaceTypeCreateReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceTypeCreateReply>, I>>(
    object: I
  ): WorkspaceTypeCreateReply {
    const message = createBaseWorkspaceTypeCreateReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? WorkspaceTypeEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWorkspaceTypeReply(): WorkspaceTypeReply {
  return { data: undefined, error: undefined };
}

export const WorkspaceTypeReply = {
  encode(
    message: WorkspaceTypeReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data !== undefined) {
      WorkspaceTypeEntity.encode(
        message.data,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WorkspaceTypeReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceTypeReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = WorkspaceTypeEntity.decode(reader, reader.uint32());
          break;
        case 2:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceTypeReply {
    return {
      data: isSet(object.data)
        ? WorkspaceTypeEntity.fromJSON(object.data)
        : undefined,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WorkspaceTypeReply): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = message.data
        ? WorkspaceTypeEntity.toJSON(message.data)
        : undefined);
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceTypeReply>, I>>(
    base?: I
  ): WorkspaceTypeReply {
    return WorkspaceTypeReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceTypeReply>, I>>(
    object: I
  ): WorkspaceTypeReply {
    const message = createBaseWorkspaceTypeReply();
    message.data =
      object.data !== undefined && object.data !== null
        ? WorkspaceTypeEntity.fromPartial(object.data)
        : undefined;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWorkspaceTypeQueryReply(): WorkspaceTypeQueryReply {
  return {
    items: [],
    totalItems: 0,
    itemsPerPage: 0,
    startIndex: 0,
    error: undefined,
  };
}

export const WorkspaceTypeQueryReply = {
  encode(
    message: WorkspaceTypeQueryReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      WorkspaceTypeEntity.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.totalItems !== 0) {
      writer.uint32(16).int64(message.totalItems);
    }
    if (message.itemsPerPage !== 0) {
      writer.uint32(24).int64(message.itemsPerPage);
    }
    if (message.startIndex !== 0) {
      writer.uint32(32).int64(message.startIndex);
    }
    if (message.error !== undefined) {
      IError.encode(message.error, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WorkspaceTypeQueryReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceTypeQueryReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(
            WorkspaceTypeEntity.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.totalItems = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.itemsPerPage = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.startIndex = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.error = IError.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceTypeQueryReply {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => WorkspaceTypeEntity.fromJSON(e))
        : [],
      totalItems: isSet(object.totalItems) ? Number(object.totalItems) : 0,
      itemsPerPage: isSet(object.itemsPerPage)
        ? Number(object.itemsPerPage)
        : 0,
      startIndex: isSet(object.startIndex) ? Number(object.startIndex) : 0,
      error: isSet(object.error) ? IError.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WorkspaceTypeQueryReply): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? WorkspaceTypeEntity.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    message.totalItems !== undefined &&
      (obj.totalItems = Math.round(message.totalItems));
    message.itemsPerPage !== undefined &&
      (obj.itemsPerPage = Math.round(message.itemsPerPage));
    message.startIndex !== undefined &&
      (obj.startIndex = Math.round(message.startIndex));
    message.error !== undefined &&
      (obj.error = message.error ? IError.toJSON(message.error) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceTypeQueryReply>, I>>(
    base?: I
  ): WorkspaceTypeQueryReply {
    return WorkspaceTypeQueryReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceTypeQueryReply>, I>>(
    object: I
  ): WorkspaceTypeQueryReply {
    const message = createBaseWorkspaceTypeQueryReply();
    message.items =
      object.items?.map((e) => WorkspaceTypeEntity.fromPartial(e)) || [];
    message.totalItems = object.totalItems ?? 0;
    message.itemsPerPage = object.itemsPerPage ?? 0;
    message.startIndex = object.startIndex ?? 0;
    message.error =
      object.error !== undefined && object.error !== null
        ? IError.fromPartial(object.error)
        : undefined;
    return message;
  },
};

function createBaseWorkspaceTypeEntity(): WorkspaceTypeEntity {
  return {
    visibility: undefined,
    workspaceId: undefined,
    linkerId: undefined,
    parentId: undefined,
    parent: undefined,
    uniqueId: "",
    userId: undefined,
    translations: [],
    title: undefined,
    description: undefined,
    slug: undefined,
    roleId: undefined,
    role: undefined,
    rank: 0,
    updated: 0,
    created: 0,
    createdFormatted: "",
    updatedFormatted: "",
  };
}

export const WorkspaceTypeEntity = {
  encode(
    message: WorkspaceTypeEntity,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.visibility !== undefined) {
      writer.uint32(10).string(message.visibility);
    }
    if (message.workspaceId !== undefined) {
      writer.uint32(18).string(message.workspaceId);
    }
    if (message.linkerId !== undefined) {
      writer.uint32(26).string(message.linkerId);
    }
    if (message.parentId !== undefined) {
      writer.uint32(34).string(message.parentId);
    }
    if (message.parent !== undefined) {
      WorkspaceTypeEntity.encode(
        message.parent,
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.uniqueId !== "") {
      writer.uint32(50).string(message.uniqueId);
    }
    if (message.userId !== undefined) {
      writer.uint32(58).string(message.userId);
    }
    for (const v of message.translations) {
      WorkspaceTypeEntityPolyglot.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.title !== undefined) {
      writer.uint32(82).string(message.title);
    }
    if (message.description !== undefined) {
      writer.uint32(90).string(message.description);
    }
    if (message.slug !== undefined) {
      writer.uint32(98).string(message.slug);
    }
    if (message.roleId !== undefined) {
      writer.uint32(114).string(message.roleId);
    }
    if (message.role !== undefined) {
      RoleEntity.encode(message.role, writer.uint32(122).fork()).ldelim();
    }
    if (message.rank !== 0) {
      writer.uint32(128).int64(message.rank);
    }
    if (message.updated !== 0) {
      writer.uint32(136).int64(message.updated);
    }
    if (message.created !== 0) {
      writer.uint32(144).int64(message.created);
    }
    if (message.createdFormatted !== "") {
      writer.uint32(154).string(message.createdFormatted);
    }
    if (message.updatedFormatted !== "") {
      writer.uint32(162).string(message.updatedFormatted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WorkspaceTypeEntity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceTypeEntity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.visibility = reader.string();
          break;
        case 2:
          message.workspaceId = reader.string();
          break;
        case 3:
          message.linkerId = reader.string();
          break;
        case 4:
          message.parentId = reader.string();
          break;
        case 5:
          message.parent = WorkspaceTypeEntity.decode(reader, reader.uint32());
          break;
        case 6:
          message.uniqueId = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        case 8:
          message.translations.push(
            WorkspaceTypeEntityPolyglot.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.title = reader.string();
          break;
        case 11:
          message.description = reader.string();
          break;
        case 12:
          message.slug = reader.string();
          break;
        case 14:
          message.roleId = reader.string();
          break;
        case 15:
          message.role = RoleEntity.decode(reader, reader.uint32());
          break;
        case 16:
          message.rank = longToNumber(reader.int64() as Long);
          break;
        case 17:
          message.updated = longToNumber(reader.int64() as Long);
          break;
        case 18:
          message.created = longToNumber(reader.int64() as Long);
          break;
        case 19:
          message.createdFormatted = reader.string();
          break;
        case 20:
          message.updatedFormatted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceTypeEntity {
    return {
      visibility: isSet(object.visibility)
        ? String(object.visibility)
        : undefined,
      workspaceId: isSet(object.workspaceId)
        ? String(object.workspaceId)
        : undefined,
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : undefined,
      parentId: isSet(object.parentId) ? String(object.parentId) : undefined,
      parent: isSet(object.parent)
        ? WorkspaceTypeEntity.fromJSON(object.parent)
        : undefined,
      uniqueId: isSet(object.uniqueId) ? String(object.uniqueId) : "",
      userId: isSet(object.userId) ? String(object.userId) : undefined,
      translations: Array.isArray(object?.translations)
        ? object.translations.map((e: any) =>
            WorkspaceTypeEntityPolyglot.fromJSON(e)
          )
        : [],
      title: isSet(object.title) ? String(object.title) : undefined,
      description: isSet(object.description)
        ? String(object.description)
        : undefined,
      slug: isSet(object.slug) ? String(object.slug) : undefined,
      roleId: isSet(object.roleId) ? String(object.roleId) : undefined,
      role: isSet(object.role) ? RoleEntity.fromJSON(object.role) : undefined,
      rank: isSet(object.rank) ? Number(object.rank) : 0,
      updated: isSet(object.updated) ? Number(object.updated) : 0,
      created: isSet(object.created) ? Number(object.created) : 0,
      createdFormatted: isSet(object.createdFormatted)
        ? String(object.createdFormatted)
        : "",
      updatedFormatted: isSet(object.updatedFormatted)
        ? String(object.updatedFormatted)
        : "",
    };
  },

  toJSON(message: WorkspaceTypeEntity): unknown {
    const obj: any = {};
    message.visibility !== undefined && (obj.visibility = message.visibility);
    message.workspaceId !== undefined &&
      (obj.workspaceId = message.workspaceId);
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.parentId !== undefined && (obj.parentId = message.parentId);
    message.parent !== undefined &&
      (obj.parent = message.parent
        ? WorkspaceTypeEntity.toJSON(message.parent)
        : undefined);
    message.uniqueId !== undefined && (obj.uniqueId = message.uniqueId);
    message.userId !== undefined && (obj.userId = message.userId);
    if (message.translations) {
      obj.translations = message.translations.map((e) =>
        e ? WorkspaceTypeEntityPolyglot.toJSON(e) : undefined
      );
    } else {
      obj.translations = [];
    }
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.slug !== undefined && (obj.slug = message.slug);
    message.roleId !== undefined && (obj.roleId = message.roleId);
    message.role !== undefined &&
      (obj.role = message.role ? RoleEntity.toJSON(message.role) : undefined);
    message.rank !== undefined && (obj.rank = Math.round(message.rank));
    message.updated !== undefined &&
      (obj.updated = Math.round(message.updated));
    message.created !== undefined &&
      (obj.created = Math.round(message.created));
    message.createdFormatted !== undefined &&
      (obj.createdFormatted = message.createdFormatted);
    message.updatedFormatted !== undefined &&
      (obj.updatedFormatted = message.updatedFormatted);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceTypeEntity>, I>>(
    base?: I
  ): WorkspaceTypeEntity {
    return WorkspaceTypeEntity.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceTypeEntity>, I>>(
    object: I
  ): WorkspaceTypeEntity {
    const message = createBaseWorkspaceTypeEntity();
    message.visibility = object.visibility ?? undefined;
    message.workspaceId = object.workspaceId ?? undefined;
    message.linkerId = object.linkerId ?? undefined;
    message.parentId = object.parentId ?? undefined;
    message.parent =
      object.parent !== undefined && object.parent !== null
        ? WorkspaceTypeEntity.fromPartial(object.parent)
        : undefined;
    message.uniqueId = object.uniqueId ?? "";
    message.userId = object.userId ?? undefined;
    message.translations =
      object.translations?.map((e) =>
        WorkspaceTypeEntityPolyglot.fromPartial(e)
      ) || [];
    message.title = object.title ?? undefined;
    message.description = object.description ?? undefined;
    message.slug = object.slug ?? undefined;
    message.roleId = object.roleId ?? undefined;
    message.role =
      object.role !== undefined && object.role !== null
        ? RoleEntity.fromPartial(object.role)
        : undefined;
    message.rank = object.rank ?? 0;
    message.updated = object.updated ?? 0;
    message.created = object.created ?? 0;
    message.createdFormatted = object.createdFormatted ?? "";
    message.updatedFormatted = object.updatedFormatted ?? "";
    return message;
  },
};

function createBaseWorkspaceTypeEntityPolyglot(): WorkspaceTypeEntityPolyglot {
  return { linkerId: "", languageId: "", title: "", description: "" };
}

export const WorkspaceTypeEntityPolyglot = {
  encode(
    message: WorkspaceTypeEntityPolyglot,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.linkerId !== "") {
      writer.uint32(10).string(message.linkerId);
    }
    if (message.languageId !== "") {
      writer.uint32(18).string(message.languageId);
    }
    if (message.title !== "") {
      writer.uint32(26).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WorkspaceTypeEntityPolyglot {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWorkspaceTypeEntityPolyglot();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.linkerId = reader.string();
          break;
        case 2:
          message.languageId = reader.string();
          break;
        case 3:
          message.title = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorkspaceTypeEntityPolyglot {
    return {
      linkerId: isSet(object.linkerId) ? String(object.linkerId) : "",
      languageId: isSet(object.languageId) ? String(object.languageId) : "",
      title: isSet(object.title) ? String(object.title) : "",
      description: isSet(object.description) ? String(object.description) : "",
    };
  },

  toJSON(message: WorkspaceTypeEntityPolyglot): unknown {
    const obj: any = {};
    message.linkerId !== undefined && (obj.linkerId = message.linkerId);
    message.languageId !== undefined && (obj.languageId = message.languageId);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    return obj;
  },

  create<I extends Exact<DeepPartial<WorkspaceTypeEntityPolyglot>, I>>(
    base?: I
  ): WorkspaceTypeEntityPolyglot {
    return WorkspaceTypeEntityPolyglot.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WorkspaceTypeEntityPolyglot>, I>>(
    object: I
  ): WorkspaceTypeEntityPolyglot {
    const message = createBaseWorkspaceTypeEntityPolyglot();
    message.linkerId = object.linkerId ?? "";
    message.languageId = object.languageId ?? "";
    message.title = object.title ?? "";
    message.description = object.description ?? "";
    return message;
  },
};

export interface AppMenus {
  AppMenuActionCreate(request: AppMenuEntity): Promise<AppMenuCreateReply>;
  AppMenuActionUpdate(request: AppMenuEntity): Promise<AppMenuCreateReply>;
  AppMenuActionQuery(request: QueryFilterRequest): Promise<AppMenuQueryReply>;
  AppMenuActionGetOne(request: QueryFilterRequest): Promise<AppMenuReply>;
  AppMenuActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class AppMenusClientImpl implements AppMenus {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "AppMenus";
    this.rpc = rpc;
    this.AppMenuActionCreate = this.AppMenuActionCreate.bind(this);
    this.AppMenuActionUpdate = this.AppMenuActionUpdate.bind(this);
    this.AppMenuActionQuery = this.AppMenuActionQuery.bind(this);
    this.AppMenuActionGetOne = this.AppMenuActionGetOne.bind(this);
    this.AppMenuActionRemove = this.AppMenuActionRemove.bind(this);
  }
  AppMenuActionCreate(request: AppMenuEntity): Promise<AppMenuCreateReply> {
    const data = AppMenuEntity.encode(request).finish();
    const promise = this.rpc.request(this.service, "AppMenuActionCreate", data);
    return promise.then((data) =>
      AppMenuCreateReply.decode(new _m0.Reader(data))
    );
  }

  AppMenuActionUpdate(request: AppMenuEntity): Promise<AppMenuCreateReply> {
    const data = AppMenuEntity.encode(request).finish();
    const promise = this.rpc.request(this.service, "AppMenuActionUpdate", data);
    return promise.then((data) =>
      AppMenuCreateReply.decode(new _m0.Reader(data))
    );
  }

  AppMenuActionQuery(request: QueryFilterRequest): Promise<AppMenuQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "AppMenuActionQuery", data);
    return promise.then((data) =>
      AppMenuQueryReply.decode(new _m0.Reader(data))
    );
  }

  AppMenuActionGetOne(request: QueryFilterRequest): Promise<AppMenuReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "AppMenuActionGetOne", data);
    return promise.then((data) => AppMenuReply.decode(new _m0.Reader(data)));
  }

  AppMenuActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "AppMenuActionRemove", data);
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface BackupTableMetas {
  BackupTableMetaActionCreate(
    request: BackupTableMetaEntity
  ): Promise<BackupTableMetaCreateReply>;
  BackupTableMetaActionUpdate(
    request: BackupTableMetaEntity
  ): Promise<BackupTableMetaCreateReply>;
  BackupTableMetaActionQuery(
    request: QueryFilterRequest
  ): Promise<BackupTableMetaQueryReply>;
  BackupTableMetaActionGetOne(
    request: QueryFilterRequest
  ): Promise<BackupTableMetaReply>;
  BackupTableMetaActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply>;
}

export class BackupTableMetasClientImpl implements BackupTableMetas {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "BackupTableMetas";
    this.rpc = rpc;
    this.BackupTableMetaActionCreate =
      this.BackupTableMetaActionCreate.bind(this);
    this.BackupTableMetaActionUpdate =
      this.BackupTableMetaActionUpdate.bind(this);
    this.BackupTableMetaActionQuery =
      this.BackupTableMetaActionQuery.bind(this);
    this.BackupTableMetaActionGetOne =
      this.BackupTableMetaActionGetOne.bind(this);
    this.BackupTableMetaActionRemove =
      this.BackupTableMetaActionRemove.bind(this);
  }
  BackupTableMetaActionCreate(
    request: BackupTableMetaEntity
  ): Promise<BackupTableMetaCreateReply> {
    const data = BackupTableMetaEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "BackupTableMetaActionCreate",
      data
    );
    return promise.then((data) =>
      BackupTableMetaCreateReply.decode(new _m0.Reader(data))
    );
  }

  BackupTableMetaActionUpdate(
    request: BackupTableMetaEntity
  ): Promise<BackupTableMetaCreateReply> {
    const data = BackupTableMetaEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "BackupTableMetaActionUpdate",
      data
    );
    return promise.then((data) =>
      BackupTableMetaCreateReply.decode(new _m0.Reader(data))
    );
  }

  BackupTableMetaActionQuery(
    request: QueryFilterRequest
  ): Promise<BackupTableMetaQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "BackupTableMetaActionQuery",
      data
    );
    return promise.then((data) =>
      BackupTableMetaQueryReply.decode(new _m0.Reader(data))
    );
  }

  BackupTableMetaActionGetOne(
    request: QueryFilterRequest
  ): Promise<BackupTableMetaReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "BackupTableMetaActionGetOne",
      data
    );
    return promise.then((data) =>
      BackupTableMetaReply.decode(new _m0.Reader(data))
    );
  }

  BackupTableMetaActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "BackupTableMetaActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface Capabilitys {
  CapabilityActionCreate(
    request: CapabilityEntity
  ): Promise<CapabilityCreateReply>;
  CapabilityActionUpdate(
    request: CapabilityEntity
  ): Promise<CapabilityCreateReply>;
  CapabilityActionQuery(
    request: QueryFilterRequest
  ): Promise<CapabilityQueryReply>;
  CapabilityActionGetOne(request: QueryFilterRequest): Promise<CapabilityReply>;
  CapabilityActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class CapabilitysClientImpl implements Capabilitys {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "Capabilitys";
    this.rpc = rpc;
    this.CapabilityActionCreate = this.CapabilityActionCreate.bind(this);
    this.CapabilityActionUpdate = this.CapabilityActionUpdate.bind(this);
    this.CapabilityActionQuery = this.CapabilityActionQuery.bind(this);
    this.CapabilityActionGetOne = this.CapabilityActionGetOne.bind(this);
    this.CapabilityActionRemove = this.CapabilityActionRemove.bind(this);
  }
  CapabilityActionCreate(
    request: CapabilityEntity
  ): Promise<CapabilityCreateReply> {
    const data = CapabilityEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CapabilityActionCreate",
      data
    );
    return promise.then((data) =>
      CapabilityCreateReply.decode(new _m0.Reader(data))
    );
  }

  CapabilityActionUpdate(
    request: CapabilityEntity
  ): Promise<CapabilityCreateReply> {
    const data = CapabilityEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CapabilityActionUpdate",
      data
    );
    return promise.then((data) =>
      CapabilityCreateReply.decode(new _m0.Reader(data))
    );
  }

  CapabilityActionQuery(
    request: QueryFilterRequest
  ): Promise<CapabilityQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CapabilityActionQuery",
      data
    );
    return promise.then((data) =>
      CapabilityQueryReply.decode(new _m0.Reader(data))
    );
  }

  CapabilityActionGetOne(
    request: QueryFilterRequest
  ): Promise<CapabilityReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CapabilityActionGetOne",
      data
    );
    return promise.then((data) => CapabilityReply.decode(new _m0.Reader(data)));
  }

  CapabilityActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "CapabilityActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface EmailConfirmations {
  EmailConfirmationActionCreate(
    request: EmailConfirmationEntity
  ): Promise<EmailConfirmationCreateReply>;
  EmailConfirmationActionUpdate(
    request: EmailConfirmationEntity
  ): Promise<EmailConfirmationCreateReply>;
  EmailConfirmationActionQuery(
    request: QueryFilterRequest
  ): Promise<EmailConfirmationQueryReply>;
  EmailConfirmationActionGetOne(
    request: QueryFilterRequest
  ): Promise<EmailConfirmationReply>;
  EmailConfirmationActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply>;
}

export class EmailConfirmationsClientImpl implements EmailConfirmations {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "EmailConfirmations";
    this.rpc = rpc;
    this.EmailConfirmationActionCreate =
      this.EmailConfirmationActionCreate.bind(this);
    this.EmailConfirmationActionUpdate =
      this.EmailConfirmationActionUpdate.bind(this);
    this.EmailConfirmationActionQuery =
      this.EmailConfirmationActionQuery.bind(this);
    this.EmailConfirmationActionGetOne =
      this.EmailConfirmationActionGetOne.bind(this);
    this.EmailConfirmationActionRemove =
      this.EmailConfirmationActionRemove.bind(this);
  }
  EmailConfirmationActionCreate(
    request: EmailConfirmationEntity
  ): Promise<EmailConfirmationCreateReply> {
    const data = EmailConfirmationEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailConfirmationActionCreate",
      data
    );
    return promise.then((data) =>
      EmailConfirmationCreateReply.decode(new _m0.Reader(data))
    );
  }

  EmailConfirmationActionUpdate(
    request: EmailConfirmationEntity
  ): Promise<EmailConfirmationCreateReply> {
    const data = EmailConfirmationEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailConfirmationActionUpdate",
      data
    );
    return promise.then((data) =>
      EmailConfirmationCreateReply.decode(new _m0.Reader(data))
    );
  }

  EmailConfirmationActionQuery(
    request: QueryFilterRequest
  ): Promise<EmailConfirmationQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailConfirmationActionQuery",
      data
    );
    return promise.then((data) =>
      EmailConfirmationQueryReply.decode(new _m0.Reader(data))
    );
  }

  EmailConfirmationActionGetOne(
    request: QueryFilterRequest
  ): Promise<EmailConfirmationReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailConfirmationActionGetOne",
      data
    );
    return promise.then((data) =>
      EmailConfirmationReply.decode(new _m0.Reader(data))
    );
  }

  EmailConfirmationActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailConfirmationActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface EmailProviders {
  EmailProviderActionCreate(
    request: EmailProviderEntity
  ): Promise<EmailProviderCreateReply>;
  EmailProviderActionUpdate(
    request: EmailProviderEntity
  ): Promise<EmailProviderCreateReply>;
  EmailProviderActionQuery(
    request: QueryFilterRequest
  ): Promise<EmailProviderQueryReply>;
  EmailProviderActionGetOne(
    request: QueryFilterRequest
  ): Promise<EmailProviderReply>;
  EmailProviderActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class EmailProvidersClientImpl implements EmailProviders {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "EmailProviders";
    this.rpc = rpc;
    this.EmailProviderActionCreate = this.EmailProviderActionCreate.bind(this);
    this.EmailProviderActionUpdate = this.EmailProviderActionUpdate.bind(this);
    this.EmailProviderActionQuery = this.EmailProviderActionQuery.bind(this);
    this.EmailProviderActionGetOne = this.EmailProviderActionGetOne.bind(this);
    this.EmailProviderActionRemove = this.EmailProviderActionRemove.bind(this);
  }
  EmailProviderActionCreate(
    request: EmailProviderEntity
  ): Promise<EmailProviderCreateReply> {
    const data = EmailProviderEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailProviderActionCreate",
      data
    );
    return promise.then((data) =>
      EmailProviderCreateReply.decode(new _m0.Reader(data))
    );
  }

  EmailProviderActionUpdate(
    request: EmailProviderEntity
  ): Promise<EmailProviderCreateReply> {
    const data = EmailProviderEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailProviderActionUpdate",
      data
    );
    return promise.then((data) =>
      EmailProviderCreateReply.decode(new _m0.Reader(data))
    );
  }

  EmailProviderActionQuery(
    request: QueryFilterRequest
  ): Promise<EmailProviderQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailProviderActionQuery",
      data
    );
    return promise.then((data) =>
      EmailProviderQueryReply.decode(new _m0.Reader(data))
    );
  }

  EmailProviderActionGetOne(
    request: QueryFilterRequest
  ): Promise<EmailProviderReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailProviderActionGetOne",
      data
    );
    return promise.then((data) =>
      EmailProviderReply.decode(new _m0.Reader(data))
    );
  }

  EmailProviderActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailProviderActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface EmailSenders {
  EmailSenderActionCreate(
    request: EmailSenderEntity
  ): Promise<EmailSenderCreateReply>;
  EmailSenderActionUpdate(
    request: EmailSenderEntity
  ): Promise<EmailSenderCreateReply>;
  EmailSenderActionQuery(
    request: QueryFilterRequest
  ): Promise<EmailSenderQueryReply>;
  EmailSenderActionGetOne(
    request: QueryFilterRequest
  ): Promise<EmailSenderReply>;
  EmailSenderActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class EmailSendersClientImpl implements EmailSenders {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "EmailSenders";
    this.rpc = rpc;
    this.EmailSenderActionCreate = this.EmailSenderActionCreate.bind(this);
    this.EmailSenderActionUpdate = this.EmailSenderActionUpdate.bind(this);
    this.EmailSenderActionQuery = this.EmailSenderActionQuery.bind(this);
    this.EmailSenderActionGetOne = this.EmailSenderActionGetOne.bind(this);
    this.EmailSenderActionRemove = this.EmailSenderActionRemove.bind(this);
  }
  EmailSenderActionCreate(
    request: EmailSenderEntity
  ): Promise<EmailSenderCreateReply> {
    const data = EmailSenderEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailSenderActionCreate",
      data
    );
    return promise.then((data) =>
      EmailSenderCreateReply.decode(new _m0.Reader(data))
    );
  }

  EmailSenderActionUpdate(
    request: EmailSenderEntity
  ): Promise<EmailSenderCreateReply> {
    const data = EmailSenderEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailSenderActionUpdate",
      data
    );
    return promise.then((data) =>
      EmailSenderCreateReply.decode(new _m0.Reader(data))
    );
  }

  EmailSenderActionQuery(
    request: QueryFilterRequest
  ): Promise<EmailSenderQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailSenderActionQuery",
      data
    );
    return promise.then((data) =>
      EmailSenderQueryReply.decode(new _m0.Reader(data))
    );
  }

  EmailSenderActionGetOne(
    request: QueryFilterRequest
  ): Promise<EmailSenderReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailSenderActionGetOne",
      data
    );
    return promise.then((data) =>
      EmailSenderReply.decode(new _m0.Reader(data))
    );
  }

  EmailSenderActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "EmailSenderActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface ForgetPasswords {
  ForgetPasswordActionCreate(
    request: ForgetPasswordEntity
  ): Promise<ForgetPasswordCreateReply>;
  ForgetPasswordActionUpdate(
    request: ForgetPasswordEntity
  ): Promise<ForgetPasswordCreateReply>;
  ForgetPasswordActionQuery(
    request: QueryFilterRequest
  ): Promise<ForgetPasswordQueryReply>;
  ForgetPasswordActionGetOne(
    request: QueryFilterRequest
  ): Promise<ForgetPasswordReply>;
  ForgetPasswordActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class ForgetPasswordsClientImpl implements ForgetPasswords {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "ForgetPasswords";
    this.rpc = rpc;
    this.ForgetPasswordActionCreate =
      this.ForgetPasswordActionCreate.bind(this);
    this.ForgetPasswordActionUpdate =
      this.ForgetPasswordActionUpdate.bind(this);
    this.ForgetPasswordActionQuery = this.ForgetPasswordActionQuery.bind(this);
    this.ForgetPasswordActionGetOne =
      this.ForgetPasswordActionGetOne.bind(this);
    this.ForgetPasswordActionRemove =
      this.ForgetPasswordActionRemove.bind(this);
  }
  ForgetPasswordActionCreate(
    request: ForgetPasswordEntity
  ): Promise<ForgetPasswordCreateReply> {
    const data = ForgetPasswordEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ForgetPasswordActionCreate",
      data
    );
    return promise.then((data) =>
      ForgetPasswordCreateReply.decode(new _m0.Reader(data))
    );
  }

  ForgetPasswordActionUpdate(
    request: ForgetPasswordEntity
  ): Promise<ForgetPasswordCreateReply> {
    const data = ForgetPasswordEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ForgetPasswordActionUpdate",
      data
    );
    return promise.then((data) =>
      ForgetPasswordCreateReply.decode(new _m0.Reader(data))
    );
  }

  ForgetPasswordActionQuery(
    request: QueryFilterRequest
  ): Promise<ForgetPasswordQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ForgetPasswordActionQuery",
      data
    );
    return promise.then((data) =>
      ForgetPasswordQueryReply.decode(new _m0.Reader(data))
    );
  }

  ForgetPasswordActionGetOne(
    request: QueryFilterRequest
  ): Promise<ForgetPasswordReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ForgetPasswordActionGetOne",
      data
    );
    return promise.then((data) =>
      ForgetPasswordReply.decode(new _m0.Reader(data))
    );
  }

  ForgetPasswordActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "ForgetPasswordActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface GsmProviders {
  GsmProviderActionCreate(
    request: GsmProviderEntity
  ): Promise<GsmProviderCreateReply>;
  GsmProviderActionUpdate(
    request: GsmProviderEntity
  ): Promise<GsmProviderCreateReply>;
  GsmProviderActionQuery(
    request: QueryFilterRequest
  ): Promise<GsmProviderQueryReply>;
  GsmProviderActionGetOne(
    request: QueryFilterRequest
  ): Promise<GsmProviderReply>;
  GsmProviderActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class GsmProvidersClientImpl implements GsmProviders {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "GsmProviders";
    this.rpc = rpc;
    this.GsmProviderActionCreate = this.GsmProviderActionCreate.bind(this);
    this.GsmProviderActionUpdate = this.GsmProviderActionUpdate.bind(this);
    this.GsmProviderActionQuery = this.GsmProviderActionQuery.bind(this);
    this.GsmProviderActionGetOne = this.GsmProviderActionGetOne.bind(this);
    this.GsmProviderActionRemove = this.GsmProviderActionRemove.bind(this);
  }
  GsmProviderActionCreate(
    request: GsmProviderEntity
  ): Promise<GsmProviderCreateReply> {
    const data = GsmProviderEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "GsmProviderActionCreate",
      data
    );
    return promise.then((data) =>
      GsmProviderCreateReply.decode(new _m0.Reader(data))
    );
  }

  GsmProviderActionUpdate(
    request: GsmProviderEntity
  ): Promise<GsmProviderCreateReply> {
    const data = GsmProviderEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "GsmProviderActionUpdate",
      data
    );
    return promise.then((data) =>
      GsmProviderCreateReply.decode(new _m0.Reader(data))
    );
  }

  GsmProviderActionQuery(
    request: QueryFilterRequest
  ): Promise<GsmProviderQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "GsmProviderActionQuery",
      data
    );
    return promise.then((data) =>
      GsmProviderQueryReply.decode(new _m0.Reader(data))
    );
  }

  GsmProviderActionGetOne(
    request: QueryFilterRequest
  ): Promise<GsmProviderReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "GsmProviderActionGetOne",
      data
    );
    return promise.then((data) =>
      GsmProviderReply.decode(new _m0.Reader(data))
    );
  }

  GsmProviderActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "GsmProviderActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface NotificationConfigs {
  NotificationConfigActionCreate(
    request: NotificationConfigEntity
  ): Promise<NotificationConfigCreateReply>;
  NotificationConfigActionUpdate(
    request: NotificationConfigEntity
  ): Promise<NotificationConfigCreateReply>;
  NotificationConfigActionQuery(
    request: QueryFilterRequest
  ): Promise<NotificationConfigQueryReply>;
  NotificationConfigActionGetOne(
    request: QueryFilterRequest
  ): Promise<NotificationConfigReply>;
  NotificationConfigActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply>;
}

export class NotificationConfigsClientImpl implements NotificationConfigs {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "NotificationConfigs";
    this.rpc = rpc;
    this.NotificationConfigActionCreate =
      this.NotificationConfigActionCreate.bind(this);
    this.NotificationConfigActionUpdate =
      this.NotificationConfigActionUpdate.bind(this);
    this.NotificationConfigActionQuery =
      this.NotificationConfigActionQuery.bind(this);
    this.NotificationConfigActionGetOne =
      this.NotificationConfigActionGetOne.bind(this);
    this.NotificationConfigActionRemove =
      this.NotificationConfigActionRemove.bind(this);
  }
  NotificationConfigActionCreate(
    request: NotificationConfigEntity
  ): Promise<NotificationConfigCreateReply> {
    const data = NotificationConfigEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "NotificationConfigActionCreate",
      data
    );
    return promise.then((data) =>
      NotificationConfigCreateReply.decode(new _m0.Reader(data))
    );
  }

  NotificationConfigActionUpdate(
    request: NotificationConfigEntity
  ): Promise<NotificationConfigCreateReply> {
    const data = NotificationConfigEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "NotificationConfigActionUpdate",
      data
    );
    return promise.then((data) =>
      NotificationConfigCreateReply.decode(new _m0.Reader(data))
    );
  }

  NotificationConfigActionQuery(
    request: QueryFilterRequest
  ): Promise<NotificationConfigQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "NotificationConfigActionQuery",
      data
    );
    return promise.then((data) =>
      NotificationConfigQueryReply.decode(new _m0.Reader(data))
    );
  }

  NotificationConfigActionGetOne(
    request: QueryFilterRequest
  ): Promise<NotificationConfigReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "NotificationConfigActionGetOne",
      data
    );
    return promise.then((data) =>
      NotificationConfigReply.decode(new _m0.Reader(data))
    );
  }

  NotificationConfigActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "NotificationConfigActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface Passports {
  PassportActionCreate(request: PassportEntity): Promise<PassportCreateReply>;
  PassportActionUpdate(request: PassportEntity): Promise<PassportCreateReply>;
  PassportActionQuery(request: QueryFilterRequest): Promise<PassportQueryReply>;
  PassportActionGetOne(request: QueryFilterRequest): Promise<PassportReply>;
  PassportActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class PassportsClientImpl implements Passports {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "Passports";
    this.rpc = rpc;
    this.PassportActionCreate = this.PassportActionCreate.bind(this);
    this.PassportActionUpdate = this.PassportActionUpdate.bind(this);
    this.PassportActionQuery = this.PassportActionQuery.bind(this);
    this.PassportActionGetOne = this.PassportActionGetOne.bind(this);
    this.PassportActionRemove = this.PassportActionRemove.bind(this);
  }
  PassportActionCreate(request: PassportEntity): Promise<PassportCreateReply> {
    const data = PassportEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PassportActionCreate",
      data
    );
    return promise.then((data) =>
      PassportCreateReply.decode(new _m0.Reader(data))
    );
  }

  PassportActionUpdate(request: PassportEntity): Promise<PassportCreateReply> {
    const data = PassportEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PassportActionUpdate",
      data
    );
    return promise.then((data) =>
      PassportCreateReply.decode(new _m0.Reader(data))
    );
  }

  PassportActionQuery(
    request: QueryFilterRequest
  ): Promise<PassportQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "PassportActionQuery", data);
    return promise.then((data) =>
      PassportQueryReply.decode(new _m0.Reader(data))
    );
  }

  PassportActionGetOne(request: QueryFilterRequest): Promise<PassportReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PassportActionGetOne",
      data
    );
    return promise.then((data) => PassportReply.decode(new _m0.Reader(data)));
  }

  PassportActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PassportActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface PassportMethods {
  PassportMethodActionCreate(
    request: PassportMethodEntity
  ): Promise<PassportMethodCreateReply>;
  PassportMethodActionUpdate(
    request: PassportMethodEntity
  ): Promise<PassportMethodCreateReply>;
  PassportMethodActionQuery(
    request: QueryFilterRequest
  ): Promise<PassportMethodQueryReply>;
  PassportMethodActionGetOne(
    request: QueryFilterRequest
  ): Promise<PassportMethodReply>;
  PassportMethodActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class PassportMethodsClientImpl implements PassportMethods {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "PassportMethods";
    this.rpc = rpc;
    this.PassportMethodActionCreate =
      this.PassportMethodActionCreate.bind(this);
    this.PassportMethodActionUpdate =
      this.PassportMethodActionUpdate.bind(this);
    this.PassportMethodActionQuery = this.PassportMethodActionQuery.bind(this);
    this.PassportMethodActionGetOne =
      this.PassportMethodActionGetOne.bind(this);
    this.PassportMethodActionRemove =
      this.PassportMethodActionRemove.bind(this);
  }
  PassportMethodActionCreate(
    request: PassportMethodEntity
  ): Promise<PassportMethodCreateReply> {
    const data = PassportMethodEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PassportMethodActionCreate",
      data
    );
    return promise.then((data) =>
      PassportMethodCreateReply.decode(new _m0.Reader(data))
    );
  }

  PassportMethodActionUpdate(
    request: PassportMethodEntity
  ): Promise<PassportMethodCreateReply> {
    const data = PassportMethodEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PassportMethodActionUpdate",
      data
    );
    return promise.then((data) =>
      PassportMethodCreateReply.decode(new _m0.Reader(data))
    );
  }

  PassportMethodActionQuery(
    request: QueryFilterRequest
  ): Promise<PassportMethodQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PassportMethodActionQuery",
      data
    );
    return promise.then((data) =>
      PassportMethodQueryReply.decode(new _m0.Reader(data))
    );
  }

  PassportMethodActionGetOne(
    request: QueryFilterRequest
  ): Promise<PassportMethodReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PassportMethodActionGetOne",
      data
    );
    return promise.then((data) =>
      PassportMethodReply.decode(new _m0.Reader(data))
    );
  }

  PassportMethodActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PassportMethodActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface PendingWorkspaceInvites {
  PendingWorkspaceInviteActionCreate(
    request: PendingWorkspaceInviteEntity
  ): Promise<PendingWorkspaceInviteCreateReply>;
  PendingWorkspaceInviteActionUpdate(
    request: PendingWorkspaceInviteEntity
  ): Promise<PendingWorkspaceInviteCreateReply>;
  PendingWorkspaceInviteActionQuery(
    request: QueryFilterRequest
  ): Promise<PendingWorkspaceInviteQueryReply>;
  PendingWorkspaceInviteActionGetOne(
    request: QueryFilterRequest
  ): Promise<PendingWorkspaceInviteReply>;
  PendingWorkspaceInviteActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply>;
}

export class PendingWorkspaceInvitesClientImpl
  implements PendingWorkspaceInvites
{
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "PendingWorkspaceInvites";
    this.rpc = rpc;
    this.PendingWorkspaceInviteActionCreate =
      this.PendingWorkspaceInviteActionCreate.bind(this);
    this.PendingWorkspaceInviteActionUpdate =
      this.PendingWorkspaceInviteActionUpdate.bind(this);
    this.PendingWorkspaceInviteActionQuery =
      this.PendingWorkspaceInviteActionQuery.bind(this);
    this.PendingWorkspaceInviteActionGetOne =
      this.PendingWorkspaceInviteActionGetOne.bind(this);
    this.PendingWorkspaceInviteActionRemove =
      this.PendingWorkspaceInviteActionRemove.bind(this);
  }
  PendingWorkspaceInviteActionCreate(
    request: PendingWorkspaceInviteEntity
  ): Promise<PendingWorkspaceInviteCreateReply> {
    const data = PendingWorkspaceInviteEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PendingWorkspaceInviteActionCreate",
      data
    );
    return promise.then((data) =>
      PendingWorkspaceInviteCreateReply.decode(new _m0.Reader(data))
    );
  }

  PendingWorkspaceInviteActionUpdate(
    request: PendingWorkspaceInviteEntity
  ): Promise<PendingWorkspaceInviteCreateReply> {
    const data = PendingWorkspaceInviteEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PendingWorkspaceInviteActionUpdate",
      data
    );
    return promise.then((data) =>
      PendingWorkspaceInviteCreateReply.decode(new _m0.Reader(data))
    );
  }

  PendingWorkspaceInviteActionQuery(
    request: QueryFilterRequest
  ): Promise<PendingWorkspaceInviteQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PendingWorkspaceInviteActionQuery",
      data
    );
    return promise.then((data) =>
      PendingWorkspaceInviteQueryReply.decode(new _m0.Reader(data))
    );
  }

  PendingWorkspaceInviteActionGetOne(
    request: QueryFilterRequest
  ): Promise<PendingWorkspaceInviteReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PendingWorkspaceInviteActionGetOne",
      data
    );
    return promise.then((data) =>
      PendingWorkspaceInviteReply.decode(new _m0.Reader(data))
    );
  }

  PendingWorkspaceInviteActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PendingWorkspaceInviteActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface PhoneConfirmations {
  PhoneConfirmationActionCreate(
    request: PhoneConfirmationEntity
  ): Promise<PhoneConfirmationCreateReply>;
  PhoneConfirmationActionUpdate(
    request: PhoneConfirmationEntity
  ): Promise<PhoneConfirmationCreateReply>;
  PhoneConfirmationActionQuery(
    request: QueryFilterRequest
  ): Promise<PhoneConfirmationQueryReply>;
  PhoneConfirmationActionGetOne(
    request: QueryFilterRequest
  ): Promise<PhoneConfirmationReply>;
  PhoneConfirmationActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply>;
}

export class PhoneConfirmationsClientImpl implements PhoneConfirmations {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "PhoneConfirmations";
    this.rpc = rpc;
    this.PhoneConfirmationActionCreate =
      this.PhoneConfirmationActionCreate.bind(this);
    this.PhoneConfirmationActionUpdate =
      this.PhoneConfirmationActionUpdate.bind(this);
    this.PhoneConfirmationActionQuery =
      this.PhoneConfirmationActionQuery.bind(this);
    this.PhoneConfirmationActionGetOne =
      this.PhoneConfirmationActionGetOne.bind(this);
    this.PhoneConfirmationActionRemove =
      this.PhoneConfirmationActionRemove.bind(this);
  }
  PhoneConfirmationActionCreate(
    request: PhoneConfirmationEntity
  ): Promise<PhoneConfirmationCreateReply> {
    const data = PhoneConfirmationEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PhoneConfirmationActionCreate",
      data
    );
    return promise.then((data) =>
      PhoneConfirmationCreateReply.decode(new _m0.Reader(data))
    );
  }

  PhoneConfirmationActionUpdate(
    request: PhoneConfirmationEntity
  ): Promise<PhoneConfirmationCreateReply> {
    const data = PhoneConfirmationEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PhoneConfirmationActionUpdate",
      data
    );
    return promise.then((data) =>
      PhoneConfirmationCreateReply.decode(new _m0.Reader(data))
    );
  }

  PhoneConfirmationActionQuery(
    request: QueryFilterRequest
  ): Promise<PhoneConfirmationQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PhoneConfirmationActionQuery",
      data
    );
    return promise.then((data) =>
      PhoneConfirmationQueryReply.decode(new _m0.Reader(data))
    );
  }

  PhoneConfirmationActionGetOne(
    request: QueryFilterRequest
  ): Promise<PhoneConfirmationReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PhoneConfirmationActionGetOne",
      data
    );
    return promise.then((data) =>
      PhoneConfirmationReply.decode(new _m0.Reader(data))
    );
  }

  PhoneConfirmationActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PhoneConfirmationActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface Preferences {
  PreferenceActionCreate(
    request: PreferenceEntity
  ): Promise<PreferenceCreateReply>;
  PreferenceActionUpdate(
    request: PreferenceEntity
  ): Promise<PreferenceCreateReply>;
  PreferenceActionQuery(
    request: QueryFilterRequest
  ): Promise<PreferenceQueryReply>;
  PreferenceActionGetOne(request: QueryFilterRequest): Promise<PreferenceReply>;
  PreferenceActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class PreferencesClientImpl implements Preferences {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "Preferences";
    this.rpc = rpc;
    this.PreferenceActionCreate = this.PreferenceActionCreate.bind(this);
    this.PreferenceActionUpdate = this.PreferenceActionUpdate.bind(this);
    this.PreferenceActionQuery = this.PreferenceActionQuery.bind(this);
    this.PreferenceActionGetOne = this.PreferenceActionGetOne.bind(this);
    this.PreferenceActionRemove = this.PreferenceActionRemove.bind(this);
  }
  PreferenceActionCreate(
    request: PreferenceEntity
  ): Promise<PreferenceCreateReply> {
    const data = PreferenceEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PreferenceActionCreate",
      data
    );
    return promise.then((data) =>
      PreferenceCreateReply.decode(new _m0.Reader(data))
    );
  }

  PreferenceActionUpdate(
    request: PreferenceEntity
  ): Promise<PreferenceCreateReply> {
    const data = PreferenceEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PreferenceActionUpdate",
      data
    );
    return promise.then((data) =>
      PreferenceCreateReply.decode(new _m0.Reader(data))
    );
  }

  PreferenceActionQuery(
    request: QueryFilterRequest
  ): Promise<PreferenceQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PreferenceActionQuery",
      data
    );
    return promise.then((data) =>
      PreferenceQueryReply.decode(new _m0.Reader(data))
    );
  }

  PreferenceActionGetOne(
    request: QueryFilterRequest
  ): Promise<PreferenceReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PreferenceActionGetOne",
      data
    );
    return promise.then((data) => PreferenceReply.decode(new _m0.Reader(data)));
  }

  PreferenceActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PreferenceActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface PublicJoinKeys {
  PublicJoinKeyActionCreate(
    request: PublicJoinKeyEntity
  ): Promise<PublicJoinKeyCreateReply>;
  PublicJoinKeyActionUpdate(
    request: PublicJoinKeyEntity
  ): Promise<PublicJoinKeyCreateReply>;
  PublicJoinKeyActionQuery(
    request: QueryFilterRequest
  ): Promise<PublicJoinKeyQueryReply>;
  PublicJoinKeyActionGetOne(
    request: QueryFilterRequest
  ): Promise<PublicJoinKeyReply>;
  PublicJoinKeyActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class PublicJoinKeysClientImpl implements PublicJoinKeys {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "PublicJoinKeys";
    this.rpc = rpc;
    this.PublicJoinKeyActionCreate = this.PublicJoinKeyActionCreate.bind(this);
    this.PublicJoinKeyActionUpdate = this.PublicJoinKeyActionUpdate.bind(this);
    this.PublicJoinKeyActionQuery = this.PublicJoinKeyActionQuery.bind(this);
    this.PublicJoinKeyActionGetOne = this.PublicJoinKeyActionGetOne.bind(this);
    this.PublicJoinKeyActionRemove = this.PublicJoinKeyActionRemove.bind(this);
  }
  PublicJoinKeyActionCreate(
    request: PublicJoinKeyEntity
  ): Promise<PublicJoinKeyCreateReply> {
    const data = PublicJoinKeyEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PublicJoinKeyActionCreate",
      data
    );
    return promise.then((data) =>
      PublicJoinKeyCreateReply.decode(new _m0.Reader(data))
    );
  }

  PublicJoinKeyActionUpdate(
    request: PublicJoinKeyEntity
  ): Promise<PublicJoinKeyCreateReply> {
    const data = PublicJoinKeyEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PublicJoinKeyActionUpdate",
      data
    );
    return promise.then((data) =>
      PublicJoinKeyCreateReply.decode(new _m0.Reader(data))
    );
  }

  PublicJoinKeyActionQuery(
    request: QueryFilterRequest
  ): Promise<PublicJoinKeyQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PublicJoinKeyActionQuery",
      data
    );
    return promise.then((data) =>
      PublicJoinKeyQueryReply.decode(new _m0.Reader(data))
    );
  }

  PublicJoinKeyActionGetOne(
    request: QueryFilterRequest
  ): Promise<PublicJoinKeyReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PublicJoinKeyActionGetOne",
      data
    );
    return promise.then((data) =>
      PublicJoinKeyReply.decode(new _m0.Reader(data))
    );
  }

  PublicJoinKeyActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "PublicJoinKeyActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface Roles {
  RoleActionCreate(request: RoleEntity): Promise<RoleCreateReply>;
  RoleActionUpdate(request: RoleEntity): Promise<RoleCreateReply>;
  RoleActionQuery(request: QueryFilterRequest): Promise<RoleQueryReply>;
  RoleActionGetOne(request: QueryFilterRequest): Promise<RoleReply>;
  RoleActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class RolesClientImpl implements Roles {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "Roles";
    this.rpc = rpc;
    this.RoleActionCreate = this.RoleActionCreate.bind(this);
    this.RoleActionUpdate = this.RoleActionUpdate.bind(this);
    this.RoleActionQuery = this.RoleActionQuery.bind(this);
    this.RoleActionGetOne = this.RoleActionGetOne.bind(this);
    this.RoleActionRemove = this.RoleActionRemove.bind(this);
  }
  RoleActionCreate(request: RoleEntity): Promise<RoleCreateReply> {
    const data = RoleEntity.encode(request).finish();
    const promise = this.rpc.request(this.service, "RoleActionCreate", data);
    return promise.then((data) => RoleCreateReply.decode(new _m0.Reader(data)));
  }

  RoleActionUpdate(request: RoleEntity): Promise<RoleCreateReply> {
    const data = RoleEntity.encode(request).finish();
    const promise = this.rpc.request(this.service, "RoleActionUpdate", data);
    return promise.then((data) => RoleCreateReply.decode(new _m0.Reader(data)));
  }

  RoleActionQuery(request: QueryFilterRequest): Promise<RoleQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "RoleActionQuery", data);
    return promise.then((data) => RoleQueryReply.decode(new _m0.Reader(data)));
  }

  RoleActionGetOne(request: QueryFilterRequest): Promise<RoleReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "RoleActionGetOne", data);
    return promise.then((data) => RoleReply.decode(new _m0.Reader(data)));
  }

  RoleActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "RoleActionRemove", data);
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface TableViewSizings {
  TableViewSizingActionCreate(
    request: TableViewSizingEntity
  ): Promise<TableViewSizingCreateReply>;
  TableViewSizingActionUpdate(
    request: TableViewSizingEntity
  ): Promise<TableViewSizingCreateReply>;
  TableViewSizingActionQuery(
    request: QueryFilterRequest
  ): Promise<TableViewSizingQueryReply>;
  TableViewSizingActionGetOne(
    request: QueryFilterRequest
  ): Promise<TableViewSizingReply>;
  TableViewSizingActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply>;
}

export class TableViewSizingsClientImpl implements TableViewSizings {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "TableViewSizings";
    this.rpc = rpc;
    this.TableViewSizingActionCreate =
      this.TableViewSizingActionCreate.bind(this);
    this.TableViewSizingActionUpdate =
      this.TableViewSizingActionUpdate.bind(this);
    this.TableViewSizingActionQuery =
      this.TableViewSizingActionQuery.bind(this);
    this.TableViewSizingActionGetOne =
      this.TableViewSizingActionGetOne.bind(this);
    this.TableViewSizingActionRemove =
      this.TableViewSizingActionRemove.bind(this);
  }
  TableViewSizingActionCreate(
    request: TableViewSizingEntity
  ): Promise<TableViewSizingCreateReply> {
    const data = TableViewSizingEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "TableViewSizingActionCreate",
      data
    );
    return promise.then((data) =>
      TableViewSizingCreateReply.decode(new _m0.Reader(data))
    );
  }

  TableViewSizingActionUpdate(
    request: TableViewSizingEntity
  ): Promise<TableViewSizingCreateReply> {
    const data = TableViewSizingEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "TableViewSizingActionUpdate",
      data
    );
    return promise.then((data) =>
      TableViewSizingCreateReply.decode(new _m0.Reader(data))
    );
  }

  TableViewSizingActionQuery(
    request: QueryFilterRequest
  ): Promise<TableViewSizingQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "TableViewSizingActionQuery",
      data
    );
    return promise.then((data) =>
      TableViewSizingQueryReply.decode(new _m0.Reader(data))
    );
  }

  TableViewSizingActionGetOne(
    request: QueryFilterRequest
  ): Promise<TableViewSizingReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "TableViewSizingActionGetOne",
      data
    );
    return promise.then((data) =>
      TableViewSizingReply.decode(new _m0.Reader(data))
    );
  }

  TableViewSizingActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "TableViewSizingActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface Tokens {
  TokenActionCreate(request: TokenEntity): Promise<TokenCreateReply>;
  TokenActionUpdate(request: TokenEntity): Promise<TokenCreateReply>;
  TokenActionQuery(request: QueryFilterRequest): Promise<TokenQueryReply>;
  TokenActionGetOne(request: QueryFilterRequest): Promise<TokenReply>;
  TokenActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class TokensClientImpl implements Tokens {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "Tokens";
    this.rpc = rpc;
    this.TokenActionCreate = this.TokenActionCreate.bind(this);
    this.TokenActionUpdate = this.TokenActionUpdate.bind(this);
    this.TokenActionQuery = this.TokenActionQuery.bind(this);
    this.TokenActionGetOne = this.TokenActionGetOne.bind(this);
    this.TokenActionRemove = this.TokenActionRemove.bind(this);
  }
  TokenActionCreate(request: TokenEntity): Promise<TokenCreateReply> {
    const data = TokenEntity.encode(request).finish();
    const promise = this.rpc.request(this.service, "TokenActionCreate", data);
    return promise.then((data) =>
      TokenCreateReply.decode(new _m0.Reader(data))
    );
  }

  TokenActionUpdate(request: TokenEntity): Promise<TokenCreateReply> {
    const data = TokenEntity.encode(request).finish();
    const promise = this.rpc.request(this.service, "TokenActionUpdate", data);
    return promise.then((data) =>
      TokenCreateReply.decode(new _m0.Reader(data))
    );
  }

  TokenActionQuery(request: QueryFilterRequest): Promise<TokenQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "TokenActionQuery", data);
    return promise.then((data) => TokenQueryReply.decode(new _m0.Reader(data)));
  }

  TokenActionGetOne(request: QueryFilterRequest): Promise<TokenReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "TokenActionGetOne", data);
    return promise.then((data) => TokenReply.decode(new _m0.Reader(data)));
  }

  TokenActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "TokenActionRemove", data);
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface Users {
  UserActionCreate(request: UserEntity): Promise<UserCreateReply>;
  UserActionUpdate(request: UserEntity): Promise<UserCreateReply>;
  UserActionQuery(request: QueryFilterRequest): Promise<UserQueryReply>;
  UserActionGetOne(request: QueryFilterRequest): Promise<UserReply>;
  UserActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class UsersClientImpl implements Users {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "Users";
    this.rpc = rpc;
    this.UserActionCreate = this.UserActionCreate.bind(this);
    this.UserActionUpdate = this.UserActionUpdate.bind(this);
    this.UserActionQuery = this.UserActionQuery.bind(this);
    this.UserActionGetOne = this.UserActionGetOne.bind(this);
    this.UserActionRemove = this.UserActionRemove.bind(this);
  }
  UserActionCreate(request: UserEntity): Promise<UserCreateReply> {
    const data = UserEntity.encode(request).finish();
    const promise = this.rpc.request(this.service, "UserActionCreate", data);
    return promise.then((data) => UserCreateReply.decode(new _m0.Reader(data)));
  }

  UserActionUpdate(request: UserEntity): Promise<UserCreateReply> {
    const data = UserEntity.encode(request).finish();
    const promise = this.rpc.request(this.service, "UserActionUpdate", data);
    return promise.then((data) => UserCreateReply.decode(new _m0.Reader(data)));
  }

  UserActionQuery(request: QueryFilterRequest): Promise<UserQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UserActionQuery", data);
    return promise.then((data) => UserQueryReply.decode(new _m0.Reader(data)));
  }

  UserActionGetOne(request: QueryFilterRequest): Promise<UserReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UserActionGetOne", data);
    return promise.then((data) => UserReply.decode(new _m0.Reader(data)));
  }

  UserActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UserActionRemove", data);
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface UserProfiles {
  UserProfileActionCreate(
    request: UserProfileEntity
  ): Promise<UserProfileCreateReply>;
  UserProfileActionUpdate(
    request: UserProfileEntity
  ): Promise<UserProfileCreateReply>;
  UserProfileActionQuery(
    request: QueryFilterRequest
  ): Promise<UserProfileQueryReply>;
  UserProfileActionGetOne(
    request: QueryFilterRequest
  ): Promise<UserProfileReply>;
  UserProfileActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class UserProfilesClientImpl implements UserProfiles {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "UserProfiles";
    this.rpc = rpc;
    this.UserProfileActionCreate = this.UserProfileActionCreate.bind(this);
    this.UserProfileActionUpdate = this.UserProfileActionUpdate.bind(this);
    this.UserProfileActionQuery = this.UserProfileActionQuery.bind(this);
    this.UserProfileActionGetOne = this.UserProfileActionGetOne.bind(this);
    this.UserProfileActionRemove = this.UserProfileActionRemove.bind(this);
  }
  UserProfileActionCreate(
    request: UserProfileEntity
  ): Promise<UserProfileCreateReply> {
    const data = UserProfileEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "UserProfileActionCreate",
      data
    );
    return promise.then((data) =>
      UserProfileCreateReply.decode(new _m0.Reader(data))
    );
  }

  UserProfileActionUpdate(
    request: UserProfileEntity
  ): Promise<UserProfileCreateReply> {
    const data = UserProfileEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "UserProfileActionUpdate",
      data
    );
    return promise.then((data) =>
      UserProfileCreateReply.decode(new _m0.Reader(data))
    );
  }

  UserProfileActionQuery(
    request: QueryFilterRequest
  ): Promise<UserProfileQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "UserProfileActionQuery",
      data
    );
    return promise.then((data) =>
      UserProfileQueryReply.decode(new _m0.Reader(data))
    );
  }

  UserProfileActionGetOne(
    request: QueryFilterRequest
  ): Promise<UserProfileReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "UserProfileActionGetOne",
      data
    );
    return promise.then((data) =>
      UserProfileReply.decode(new _m0.Reader(data))
    );
  }

  UserProfileActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "UserProfileActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface UserRoleWorkspaces {
  UserRoleWorkspaceActionCreate(
    request: UserRoleWorkspaceEntity
  ): Promise<UserRoleWorkspaceCreateReply>;
  UserRoleWorkspaceActionUpdate(
    request: UserRoleWorkspaceEntity
  ): Promise<UserRoleWorkspaceCreateReply>;
  UserRoleWorkspaceActionQuery(
    request: QueryFilterRequest
  ): Promise<UserRoleWorkspaceQueryReply>;
  UserRoleWorkspaceActionGetOne(
    request: QueryFilterRequest
  ): Promise<UserRoleWorkspaceReply>;
  UserRoleWorkspaceActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply>;
}

export class UserRoleWorkspacesClientImpl implements UserRoleWorkspaces {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "UserRoleWorkspaces";
    this.rpc = rpc;
    this.UserRoleWorkspaceActionCreate =
      this.UserRoleWorkspaceActionCreate.bind(this);
    this.UserRoleWorkspaceActionUpdate =
      this.UserRoleWorkspaceActionUpdate.bind(this);
    this.UserRoleWorkspaceActionQuery =
      this.UserRoleWorkspaceActionQuery.bind(this);
    this.UserRoleWorkspaceActionGetOne =
      this.UserRoleWorkspaceActionGetOne.bind(this);
    this.UserRoleWorkspaceActionRemove =
      this.UserRoleWorkspaceActionRemove.bind(this);
  }
  UserRoleWorkspaceActionCreate(
    request: UserRoleWorkspaceEntity
  ): Promise<UserRoleWorkspaceCreateReply> {
    const data = UserRoleWorkspaceEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "UserRoleWorkspaceActionCreate",
      data
    );
    return promise.then((data) =>
      UserRoleWorkspaceCreateReply.decode(new _m0.Reader(data))
    );
  }

  UserRoleWorkspaceActionUpdate(
    request: UserRoleWorkspaceEntity
  ): Promise<UserRoleWorkspaceCreateReply> {
    const data = UserRoleWorkspaceEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "UserRoleWorkspaceActionUpdate",
      data
    );
    return promise.then((data) =>
      UserRoleWorkspaceCreateReply.decode(new _m0.Reader(data))
    );
  }

  UserRoleWorkspaceActionQuery(
    request: QueryFilterRequest
  ): Promise<UserRoleWorkspaceQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "UserRoleWorkspaceActionQuery",
      data
    );
    return promise.then((data) =>
      UserRoleWorkspaceQueryReply.decode(new _m0.Reader(data))
    );
  }

  UserRoleWorkspaceActionGetOne(
    request: QueryFilterRequest
  ): Promise<UserRoleWorkspaceReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "UserRoleWorkspaceActionGetOne",
      data
    );
    return promise.then((data) =>
      UserRoleWorkspaceReply.decode(new _m0.Reader(data))
    );
  }

  UserRoleWorkspaceActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "UserRoleWorkspaceActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface WorkspaceConfigs {
  WorkspaceConfigActionCreate(
    request: WorkspaceConfigEntity
  ): Promise<WorkspaceConfigCreateReply>;
  WorkspaceConfigActionUpdate(
    request: WorkspaceConfigEntity
  ): Promise<WorkspaceConfigCreateReply>;
  WorkspaceConfigActionQuery(
    request: QueryFilterRequest
  ): Promise<WorkspaceConfigQueryReply>;
  WorkspaceConfigActionGetOne(
    request: QueryFilterRequest
  ): Promise<WorkspaceConfigReply>;
  WorkspaceConfigActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply>;
}

export class WorkspaceConfigsClientImpl implements WorkspaceConfigs {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "WorkspaceConfigs";
    this.rpc = rpc;
    this.WorkspaceConfigActionCreate =
      this.WorkspaceConfigActionCreate.bind(this);
    this.WorkspaceConfigActionUpdate =
      this.WorkspaceConfigActionUpdate.bind(this);
    this.WorkspaceConfigActionQuery =
      this.WorkspaceConfigActionQuery.bind(this);
    this.WorkspaceConfigActionGetOne =
      this.WorkspaceConfigActionGetOne.bind(this);
    this.WorkspaceConfigActionRemove =
      this.WorkspaceConfigActionRemove.bind(this);
  }
  WorkspaceConfigActionCreate(
    request: WorkspaceConfigEntity
  ): Promise<WorkspaceConfigCreateReply> {
    const data = WorkspaceConfigEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceConfigActionCreate",
      data
    );
    return promise.then((data) =>
      WorkspaceConfigCreateReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceConfigActionUpdate(
    request: WorkspaceConfigEntity
  ): Promise<WorkspaceConfigCreateReply> {
    const data = WorkspaceConfigEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceConfigActionUpdate",
      data
    );
    return promise.then((data) =>
      WorkspaceConfigCreateReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceConfigActionQuery(
    request: QueryFilterRequest
  ): Promise<WorkspaceConfigQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceConfigActionQuery",
      data
    );
    return promise.then((data) =>
      WorkspaceConfigQueryReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceConfigActionGetOne(
    request: QueryFilterRequest
  ): Promise<WorkspaceConfigReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceConfigActionGetOne",
      data
    );
    return promise.then((data) =>
      WorkspaceConfigReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceConfigActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceConfigActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface Workspaces {
  WorkspaceActionCreate(
    request: WorkspaceEntity
  ): Promise<WorkspaceCreateReply>;
  WorkspaceActionUpdate(
    request: WorkspaceEntity
  ): Promise<WorkspaceCreateReply>;
  WorkspaceActionQuery(
    request: QueryFilterRequest
  ): Promise<WorkspaceQueryReply>;
  WorkspaceActionGetOne(request: QueryFilterRequest): Promise<WorkspaceReply>;
  WorkspaceActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class WorkspacesClientImpl implements Workspaces {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "Workspaces";
    this.rpc = rpc;
    this.WorkspaceActionCreate = this.WorkspaceActionCreate.bind(this);
    this.WorkspaceActionUpdate = this.WorkspaceActionUpdate.bind(this);
    this.WorkspaceActionQuery = this.WorkspaceActionQuery.bind(this);
    this.WorkspaceActionGetOne = this.WorkspaceActionGetOne.bind(this);
    this.WorkspaceActionRemove = this.WorkspaceActionRemove.bind(this);
  }
  WorkspaceActionCreate(
    request: WorkspaceEntity
  ): Promise<WorkspaceCreateReply> {
    const data = WorkspaceEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceActionCreate",
      data
    );
    return promise.then((data) =>
      WorkspaceCreateReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceActionUpdate(
    request: WorkspaceEntity
  ): Promise<WorkspaceCreateReply> {
    const data = WorkspaceEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceActionUpdate",
      data
    );
    return promise.then((data) =>
      WorkspaceCreateReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceActionQuery(
    request: QueryFilterRequest
  ): Promise<WorkspaceQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceActionQuery",
      data
    );
    return promise.then((data) =>
      WorkspaceQueryReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceActionGetOne(request: QueryFilterRequest): Promise<WorkspaceReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceActionGetOne",
      data
    );
    return promise.then((data) => WorkspaceReply.decode(new _m0.Reader(data)));
  }

  WorkspaceActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface WorkspaceInvites {
  WorkspaceInviteActionCreate(
    request: WorkspaceInviteEntity
  ): Promise<WorkspaceInviteCreateReply>;
  WorkspaceInviteActionUpdate(
    request: WorkspaceInviteEntity
  ): Promise<WorkspaceInviteCreateReply>;
  WorkspaceInviteActionQuery(
    request: QueryFilterRequest
  ): Promise<WorkspaceInviteQueryReply>;
  WorkspaceInviteActionGetOne(
    request: QueryFilterRequest
  ): Promise<WorkspaceInviteReply>;
  WorkspaceInviteActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply>;
}

export class WorkspaceInvitesClientImpl implements WorkspaceInvites {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "WorkspaceInvites";
    this.rpc = rpc;
    this.WorkspaceInviteActionCreate =
      this.WorkspaceInviteActionCreate.bind(this);
    this.WorkspaceInviteActionUpdate =
      this.WorkspaceInviteActionUpdate.bind(this);
    this.WorkspaceInviteActionQuery =
      this.WorkspaceInviteActionQuery.bind(this);
    this.WorkspaceInviteActionGetOne =
      this.WorkspaceInviteActionGetOne.bind(this);
    this.WorkspaceInviteActionRemove =
      this.WorkspaceInviteActionRemove.bind(this);
  }
  WorkspaceInviteActionCreate(
    request: WorkspaceInviteEntity
  ): Promise<WorkspaceInviteCreateReply> {
    const data = WorkspaceInviteEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceInviteActionCreate",
      data
    );
    return promise.then((data) =>
      WorkspaceInviteCreateReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceInviteActionUpdate(
    request: WorkspaceInviteEntity
  ): Promise<WorkspaceInviteCreateReply> {
    const data = WorkspaceInviteEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceInviteActionUpdate",
      data
    );
    return promise.then((data) =>
      WorkspaceInviteCreateReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceInviteActionQuery(
    request: QueryFilterRequest
  ): Promise<WorkspaceInviteQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceInviteActionQuery",
      data
    );
    return promise.then((data) =>
      WorkspaceInviteQueryReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceInviteActionGetOne(
    request: QueryFilterRequest
  ): Promise<WorkspaceInviteReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceInviteActionGetOne",
      data
    );
    return promise.then((data) =>
      WorkspaceInviteReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceInviteActionRemove(
    request: QueryFilterRequest
  ): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceInviteActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

export interface WorkspaceTypes {
  WorkspaceTypeActionCreate(
    request: WorkspaceTypeEntity
  ): Promise<WorkspaceTypeCreateReply>;
  WorkspaceTypeActionUpdate(
    request: WorkspaceTypeEntity
  ): Promise<WorkspaceTypeCreateReply>;
  WorkspaceTypeActionQuery(
    request: QueryFilterRequest
  ): Promise<WorkspaceTypeQueryReply>;
  WorkspaceTypeActionGetOne(
    request: QueryFilterRequest
  ): Promise<WorkspaceTypeReply>;
  WorkspaceTypeActionRemove(request: QueryFilterRequest): Promise<RemoveReply>;
}

export class WorkspaceTypesClientImpl implements WorkspaceTypes {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "WorkspaceTypes";
    this.rpc = rpc;
    this.WorkspaceTypeActionCreate = this.WorkspaceTypeActionCreate.bind(this);
    this.WorkspaceTypeActionUpdate = this.WorkspaceTypeActionUpdate.bind(this);
    this.WorkspaceTypeActionQuery = this.WorkspaceTypeActionQuery.bind(this);
    this.WorkspaceTypeActionGetOne = this.WorkspaceTypeActionGetOne.bind(this);
    this.WorkspaceTypeActionRemove = this.WorkspaceTypeActionRemove.bind(this);
  }
  WorkspaceTypeActionCreate(
    request: WorkspaceTypeEntity
  ): Promise<WorkspaceTypeCreateReply> {
    const data = WorkspaceTypeEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceTypeActionCreate",
      data
    );
    return promise.then((data) =>
      WorkspaceTypeCreateReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceTypeActionUpdate(
    request: WorkspaceTypeEntity
  ): Promise<WorkspaceTypeCreateReply> {
    const data = WorkspaceTypeEntity.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceTypeActionUpdate",
      data
    );
    return promise.then((data) =>
      WorkspaceTypeCreateReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceTypeActionQuery(
    request: QueryFilterRequest
  ): Promise<WorkspaceTypeQueryReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceTypeActionQuery",
      data
    );
    return promise.then((data) =>
      WorkspaceTypeQueryReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceTypeActionGetOne(
    request: QueryFilterRequest
  ): Promise<WorkspaceTypeReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceTypeActionGetOne",
      data
    );
    return promise.then((data) =>
      WorkspaceTypeReply.decode(new _m0.Reader(data))
    );
  }

  WorkspaceTypeActionRemove(request: QueryFilterRequest): Promise<RemoveReply> {
    const data = QueryFilterRequest.encode(request).finish();
    const promise = this.rpc.request(
      this.service,
      "WorkspaceTypeActionRemove",
      data
    );
    return promise.then((data) => RemoveReply.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & {
      [K in Exclude<keyof I, KeysOfUnion<P>>]: never;
    };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error(
      "Value is larger than Number.MAX_SAFE_INTEGER"
    );
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
