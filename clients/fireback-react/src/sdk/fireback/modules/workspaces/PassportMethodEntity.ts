import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type PassportMethodEntityKeys =
  keyof typeof PassportMethodEntity.Fields;
export class PassportMethodEntity extends BaseEntity {
  public name?: string | null;
  public type?: string | null;
  public region?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/passport-method/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/passport-method/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/passport-method/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/passport-methods`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "passport-method/edit/:uniqueId",
      Rcreate: "passport-method/new",
      Rsingle: "passport-method/:uniqueId",
      Rquery: "passport-methods",
  };
  public static definition = {
  "name": "passportMethod",
  "queryScope": "public",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "name",
      "type": "string",
      "validate": "required",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "type",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "region",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliShort": "method",
  "cliDescription": "Login/Signup methods which are available in the app for different regions (Email, Phone Number, Google, etc)"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      type: 'type',
      region: 'region',
}
}
