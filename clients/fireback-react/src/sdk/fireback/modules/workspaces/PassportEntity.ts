import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    UserEntity,
} from "./UserEntity"
// In this section we have sub entities related to this object
// Class body
export type PassportEntityKeys =
  keyof typeof PassportEntity.Fields;
export class PassportEntity extends BaseEntity {
  public type?: string | null;
  public user?: UserEntity | null;
  public value?: string | null;
  public password?: string | null;
  public confirmed?: boolean | null;
  public accessToken?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/passport/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/passport/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/passport/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/passports`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "passport/edit/:uniqueId",
      Rcreate: "passport/new",
      Rsingle: "passport/:uniqueId",
      Rquery: "passports",
  };
  public static definition = {
  "name": "passport",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "type",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "user",
      "type": "one",
      "target": "UserEntity",
      "computedType": "UserEntity",
      "gormMap": {}
    },
    {
      "name": "value",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "password",
      "type": "string",
      "yaml": "-",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "confirmed",
      "type": "bool",
      "computedType": "boolean",
      "gormMap": {}
    },
    {
      "name": "accessToken",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      type: 'type',
      user$: 'user',
      user: UserEntity.Fields,
      value: 'value',
      password: 'password',
      confirmed: 'confirmed',
      accessToken: 'accessToken',
}
}
