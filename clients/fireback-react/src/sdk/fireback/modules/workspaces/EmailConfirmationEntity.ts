import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    UserEntity,
} from "./UserEntity"
// In this section we have sub entities related to this object
// Class body
export type EmailConfirmationEntityKeys =
  keyof typeof EmailConfirmationEntity.Fields;
export class EmailConfirmationEntity extends BaseEntity {
  public user?: UserEntity | null;
  public status?: string | null;
  public email?: string | null;
  public key?: string | null;
  public expiresAt?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/email-confirmation/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/email-confirmation/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/email-confirmation/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/email-confirmations`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "email-confirmation/edit/:uniqueId",
      Rcreate: "email-confirmation/new",
      Rsingle: "email-confirmation/:uniqueId",
      Rquery: "email-confirmations",
  };
  public static definition = {
  "name": "emailConfirmation",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/workspaces/UserDefinitions.dyno.proto"
  ],
  "fields": [
    {
      "name": "user",
      "type": "one",
      "target": "UserEntity",
      "computedType": "UserEntity",
      "gormMap": {}
    },
    {
      "name": "status",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "email",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "key",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "expiresAt",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      user$: 'user',
      user: UserEntity.Fields,
      status: 'status',
      email: 'email',
      key: 'key',
      expiresAt: 'expiresAt',
}
}
