import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    EmailProviderEntity,
} from "./EmailProviderEntity"
import {
    EmailSenderEntity,
} from "./EmailSenderEntity"
import {
    GsmProviderEntity,
} from "./GsmProviderEntity"
// In this section we have sub entities related to this object
// Class body
export type NotificationConfigEntityKeys =
  keyof typeof NotificationConfigEntity.Fields;
export class NotificationConfigEntity extends BaseEntity {
  public cascadeToSubWorkspaces?: boolean | null;
  public forcedCascadeEmailProvider?: boolean | null;
  public generalEmailProvider?: EmailProviderEntity | null;
      generalEmailProviderId?: string | null;
  public generalGsmProvider?: GsmProviderEntity | null;
      generalGsmProviderId?: string | null;
  public inviteToWorkspaceContent?: string | null;
  public inviteToWorkspaceContentExcerpt?: string | null;
  public inviteToWorkspaceContentDefault?: string | null;
  public inviteToWorkspaceContentDefaultExcerpt?: string | null;
  public inviteToWorkspaceTitle?: string | null;
  public inviteToWorkspaceTitleDefault?: string | null;
  public inviteToWorkspaceSender?: EmailSenderEntity | null;
      inviteToWorkspaceSenderId?: string | null;
  public accountCenterEmailSender?: EmailSenderEntity | null;
      accountCenterEmailSenderId?: string | null;
  public forgetPasswordContent?: string | null;
  public forgetPasswordContentExcerpt?: string | null;
  public forgetPasswordContentDefault?: string | null;
  public forgetPasswordContentDefaultExcerpt?: string | null;
  public forgetPasswordTitle?: string | null;
  public forgetPasswordTitleDefault?: string | null;
  public forgetPasswordSender?: EmailSenderEntity | null;
      forgetPasswordSenderId?: string | null;
  public acceptLanguage?: string | null;
  public confirmEmailSender?: EmailSenderEntity | null;
      confirmEmailSenderId?: string | null;
  public confirmEmailContent?: string | null;
  public confirmEmailContentExcerpt?: string | null;
  public confirmEmailContentDefault?: string | null;
  public confirmEmailContentDefaultExcerpt?: string | null;
  public confirmEmailTitle?: string | null;
  public confirmEmailTitleDefault?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/notification-config/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/notification-config/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/notification-config/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/notification-configs`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "notification-config/edit/:uniqueId",
      Rcreate: "notification-config/new",
      Rsingle: "notification-config/:uniqueId",
      Rquery: "notification-configs",
  };
  public static definition = {
  "name": "notificationConfig",
  "distinctBy": "workspace",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/workspaces/GsmProviderDefinitions.dyno.proto",
    "modules/workspaces/EmailProviderDefinitions.dyno.proto",
    "modules/workspaces/EmailSenderDefinitions.dyno.proto"
  ],
  "fields": [
    {
      "name": "cascadeToSubWorkspaces",
      "type": "bool",
      "computedType": "boolean",
      "gormMap": {}
    },
    {
      "name": "forcedCascadeEmailProvider",
      "type": "bool",
      "computedType": "boolean",
      "gormMap": {}
    },
    {
      "name": "generalEmailProvider",
      "type": "one",
      "target": "EmailProviderEntity",
      "computedType": "EmailProviderEntity",
      "gormMap": {}
    },
    {
      "name": "generalGsmProvider",
      "type": "one",
      "target": "GsmProviderEntity",
      "computedType": "GsmProviderEntity",
      "gormMap": {}
    },
    {
      "name": "inviteToWorkspaceContent",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "inviteToWorkspaceContentExcerpt",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "inviteToWorkspaceContentDefault",
      "type": "string",
      "computedType": "string",
      "gormMap": {},
      "sql": "false"
    },
    {
      "name": "inviteToWorkspaceContentDefaultExcerpt",
      "type": "string",
      "computedType": "string",
      "gormMap": {},
      "sql": "false"
    },
    {
      "name": "inviteToWorkspaceTitle",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "inviteToWorkspaceTitleDefault",
      "type": "string",
      "computedType": "string",
      "gormMap": {},
      "sql": "false"
    },
    {
      "name": "inviteToWorkspaceSender",
      "type": "one",
      "target": "EmailSenderEntity",
      "computedType": "EmailSenderEntity",
      "gormMap": {}
    },
    {
      "name": "accountCenterEmailSender",
      "type": "one",
      "target": "EmailSenderEntity",
      "computedType": "EmailSenderEntity",
      "gormMap": {}
    },
    {
      "name": "forgetPasswordContent",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "forgetPasswordContentExcerpt",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "forgetPasswordContentDefault",
      "type": "string",
      "computedType": "string",
      "gormMap": {},
      "sql": "false"
    },
    {
      "name": "forgetPasswordContentDefaultExcerpt",
      "type": "string",
      "computedType": "string",
      "gormMap": {},
      "sql": "false"
    },
    {
      "name": "forgetPasswordTitle",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "forgetPasswordTitleDefault",
      "type": "string",
      "computedType": "string",
      "gormMap": {},
      "sql": "false"
    },
    {
      "name": "forgetPasswordSender",
      "type": "one",
      "target": "EmailSenderEntity",
      "computedType": "EmailSenderEntity",
      "gormMap": {}
    },
    {
      "name": "acceptLanguage",
      "type": "text",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "confirmEmailSender",
      "type": "one",
      "target": "EmailSenderEntity",
      "computedType": "EmailSenderEntity",
      "gormMap": {}
    },
    {
      "name": "confirmEmailContent",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "confirmEmailContentExcerpt",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "confirmEmailContentDefault",
      "type": "string",
      "computedType": "string",
      "gormMap": {},
      "sql": "false"
    },
    {
      "name": "confirmEmailContentDefaultExcerpt",
      "type": "string",
      "computedType": "string",
      "gormMap": {},
      "sql": "false"
    },
    {
      "name": "confirmEmailTitle",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "confirmEmailTitleDefault",
      "type": "string",
      "computedType": "string",
      "gormMap": {},
      "sql": "false"
    }
  ],
  "cliShort": "config",
  "cliDescription": "Configuration for the notifications used in the app, such as default gsm number, email senders, and many more"
}
public static Fields = {
  ...BaseEntity.Fields,
      cascadeToSubWorkspaces: 'cascadeToSubWorkspaces',
      forcedCascadeEmailProvider: 'forcedCascadeEmailProvider',
          generalEmailProviderId: 'generalEmailProviderId',
      generalEmailProvider$: 'generalEmailProvider',
      generalEmailProvider: EmailProviderEntity.Fields,
          generalGsmProviderId: 'generalGsmProviderId',
      generalGsmProvider$: 'generalGsmProvider',
      generalGsmProvider: GsmProviderEntity.Fields,
      inviteToWorkspaceContent: 'inviteToWorkspaceContent',
      inviteToWorkspaceContentExcerpt: 'inviteToWorkspaceContentExcerpt',
      inviteToWorkspaceContentDefault: 'inviteToWorkspaceContentDefault',
      inviteToWorkspaceContentDefaultExcerpt: 'inviteToWorkspaceContentDefaultExcerpt',
      inviteToWorkspaceTitle: 'inviteToWorkspaceTitle',
      inviteToWorkspaceTitleDefault: 'inviteToWorkspaceTitleDefault',
          inviteToWorkspaceSenderId: 'inviteToWorkspaceSenderId',
      inviteToWorkspaceSender$: 'inviteToWorkspaceSender',
      inviteToWorkspaceSender: EmailSenderEntity.Fields,
          accountCenterEmailSenderId: 'accountCenterEmailSenderId',
      accountCenterEmailSender$: 'accountCenterEmailSender',
      accountCenterEmailSender: EmailSenderEntity.Fields,
      forgetPasswordContent: 'forgetPasswordContent',
      forgetPasswordContentExcerpt: 'forgetPasswordContentExcerpt',
      forgetPasswordContentDefault: 'forgetPasswordContentDefault',
      forgetPasswordContentDefaultExcerpt: 'forgetPasswordContentDefaultExcerpt',
      forgetPasswordTitle: 'forgetPasswordTitle',
      forgetPasswordTitleDefault: 'forgetPasswordTitleDefault',
          forgetPasswordSenderId: 'forgetPasswordSenderId',
      forgetPasswordSender$: 'forgetPasswordSender',
      forgetPasswordSender: EmailSenderEntity.Fields,
      acceptLanguage: 'acceptLanguage',
          confirmEmailSenderId: 'confirmEmailSenderId',
      confirmEmailSender$: 'confirmEmailSender',
      confirmEmailSender: EmailSenderEntity.Fields,
      confirmEmailContent: 'confirmEmailContent',
      confirmEmailContentExcerpt: 'confirmEmailContentExcerpt',
      confirmEmailContentDefault: 'confirmEmailContentDefault',
      confirmEmailContentDefaultExcerpt: 'confirmEmailContentDefaultExcerpt',
      confirmEmailTitle: 'confirmEmailTitle',
      confirmEmailTitleDefault: 'confirmEmailTitleDefault',
}
}
