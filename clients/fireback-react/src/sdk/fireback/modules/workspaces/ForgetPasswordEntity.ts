import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    PassportEntity,
} from "./PassportEntity"
import {
    UserEntity,
} from "./UserEntity"
// In this section we have sub entities related to this object
// Class body
export type ForgetPasswordEntityKeys =
  keyof typeof ForgetPasswordEntity.Fields;
export class ForgetPasswordEntity extends BaseEntity {
  public user?: UserEntity | null;
  public passport?: PassportEntity | null;
      passportId?: string | null;
  public status?: string | null;
  public validUntil?: string | null;
  public blockedUntil?: string | null;
  public secondsToUnblock?: number | null;
  public otp?: string | null;
  public recoveryAbsoluteUrl?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/forget-password/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/forget-password/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/forget-password/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/forget-passwords`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "forget-password/edit/:uniqueId",
      Rcreate: "forget-password/new",
      Rsingle: "forget-password/:uniqueId",
      Rquery: "forget-passwords",
  };
  public static definition = {
  "name": "forgetPassword",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/workspaces/UserDefinitions.dyno.proto",
    "modules/workspaces/PassportDefinitions.dyno.proto"
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
      "name": "passport",
      "type": "one",
      "target": "PassportEntity",
      "computedType": "PassportEntity",
      "gormMap": {}
    },
    {
      "name": "status",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "validUntil",
      "type": "datenano",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "blockedUntil",
      "type": "datenano",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "secondsToUnblock",
      "type": "int64",
      "computedType": "number",
      "gormMap": {}
    },
    {
      "name": "otp",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "recoveryAbsoluteUrl",
      "type": "string",
      "computedType": "string",
      "gormMap": {},
      "sql": "false"
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      user$: 'user',
      user: UserEntity.Fields,
          passportId: 'passportId',
      passport$: 'passport',
      passport: PassportEntity.Fields,
      status: 'status',
      validUntil: 'validUntil',
      blockedUntil: 'blockedUntil',
      secondsToUnblock: 'secondsToUnblock',
      otp: 'otp',
      recoveryAbsoluteUrl: 'recoveryAbsoluteUrl',
}
}
