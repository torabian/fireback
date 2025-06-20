/*
*	Generated by fireback 1.2.3
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
*/
    import {
        BaseDto,
        BaseEntity,
    } from "../../core/definitions"
    import {
        UserEntity,
    } from "./UserEntity"
// In this section we have sub entities related to this object
// Class body
export type PhoneConfirmationEntityKeys =
  keyof typeof PhoneConfirmationEntity.Fields;
export class PhoneConfirmationEntity extends BaseEntity {
  public children?: PhoneConfirmationEntity[] | null;
  public user?: UserEntity | null;
  public status?: string | null;
  public phoneNumber?: string | null;
  public key?: string | null;
  public expiresAt?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : '..'}/phone-confirmation/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : '..'}/phone-confirmation/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : '..'}/phone-confirmation/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : '..'}/phone-confirmations`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "phone-confirmation/edit/:uniqueId",
      Rcreate: "phone-confirmation/new",
      Rsingle: "phone-confirmation/:uniqueId",
      Rquery: "phone-confirmations",
  };
  public static definition = {
  "rpc": {
    "query": {}
  },
  "permRewrite": {
    "replace": "root.modules",
    "with": "root.manage"
  },
  "name": "phoneConfirmation",
  "features": {},
  "gormMap": {},
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
      "name": "phoneNumber",
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
      user$: `user`,
        user: UserEntity.Fields,
      status: `status`,
      phoneNumber: `phoneNumber`,
      key: `key`,
      expiresAt: `expiresAt`,
}
}
