import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    EmailProviderEntity,
} from "./EmailProviderEntity"
import {
    GsmProviderEntity,
} from "./GsmProviderEntity"
import {
    UserSessionDto,
} from "./UserSessionDto"
import {
    WorkspaceEntity,
} from "./WorkspaceEntity"
export class SendEmailActionReqDto {
  public toAddress?: string | null;
  public body?: string | null;
public static Fields = {
      toAddress: 'toAddress',
      body: 'body',
}
}
export class SendEmailActionResDto {
  public queueId?: string | null;
public static Fields = {
      queueId: 'queueId',
}
}
export class SendEmailWithProviderActionReqDto {
  public emailProvider?: EmailProviderEntity | null;
      emailProviderId?: string | null;
  public toAddress?: string | null;
  public body?: string | null;
public static Fields = {
          emailProviderId: 'emailProviderId',
      emailProvider$: 'emailProvider',
      emailProvider: EmailProviderEntity.Fields,
      toAddress: 'toAddress',
      body: 'body',
}
}
export class SendEmailWithProviderActionResDto {
  public queueId?: string | null;
public static Fields = {
      queueId: 'queueId',
}
}
export class GsmSendSmsActionReqDto {
  public toNumber?: string | null;
  public body?: string | null;
public static Fields = {
      toNumber: 'toNumber',
      body: 'body',
}
}
export class GsmSendSmsActionResDto {
  public queueId?: string | null;
public static Fields = {
      queueId: 'queueId',
}
}
export class GsmSendSmsWithProviderActionReqDto {
  public gsmProvider?: GsmProviderEntity | null;
      gsmProviderId?: string | null;
  public toNumber?: string | null;
  public body?: string | null;
public static Fields = {
          gsmProviderId: 'gsmProviderId',
      gsmProvider$: 'gsmProvider',
      gsmProvider: GsmProviderEntity.Fields,
      toNumber: 'toNumber',
      body: 'body',
}
}
export class GsmSendSmsWithProviderActionResDto {
  public queueId?: string | null;
public static Fields = {
      queueId: 'queueId',
}
}
export class ClassicSigninActionReqDto {
  public value?: string | null;
  public password?: string | null;
public static Fields = {
      value: 'value',
      password: 'password',
}
}
export class ClassicSignupActionReqDto {
  public value?: string | null;
  public type?: "phonenumber" | "email" | null;
  public password?: string | null;
  public firstName?: string | null;
  public lastName?: string | null;
  public inviteId?: string | null;
  public publicJoinKeyId?: string | null;
  public workspaceTypeId?: string | null;
public static Fields = {
      value: 'value',
      type: 'type',
      password: 'password',
      firstName: 'firstName',
      lastName: 'lastName',
      inviteId: 'inviteId',
      publicJoinKeyId: 'publicJoinKeyId',
      workspaceTypeId: 'workspaceTypeId',
}
}
export class CreateWorkspaceActionReqDto {
  public name?: string | null;
  public workspace?: WorkspaceEntity | null;
  public workspaceId?: string | null;
public static Fields = {
      name: 'name',
      workspace$: 'workspace',
      workspace: WorkspaceEntity.Fields,
      workspaceId: 'workspaceId',
}
}
export class CheckClassicPassportActionReqDto {
  public value?: string | null;
public static Fields = {
      value: 'value',
}
}
export class CheckClassicPassportActionResDto {
  public exists?: boolean | null;
public static Fields = {
      exists: 'exists',
}
}
export class ClassicPassportOtpActionReqDto {
  public value?: string | null;
  public otp?: string | null;
public static Fields = {
      value: 'value',
      otp: 'otp',
}
}
export class ClassicPassportOtpActionResDto {
  public suspendUntil?: number | null;
  public session?: UserSessionDto | null;
      sessionId?: string | null;
  public validUntil?: number | null;
  public blockedUntil?: number | null;
  public secondsToUnblock?: number | null;
public static Fields = {
      suspendUntil: 'suspendUntil',
          sessionId: 'sessionId',
      session$: 'session',
      session: UserSessionDto.Fields,
      validUntil: 'validUntil',
      blockedUntil: 'blockedUntil',
      secondsToUnblock: 'secondsToUnblock',
}
}