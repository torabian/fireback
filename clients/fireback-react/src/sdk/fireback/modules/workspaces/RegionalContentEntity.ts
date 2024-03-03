import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type RegionalContentEntityKeys =
  keyof typeof RegionalContentEntity.Fields;
export class RegionalContentEntity extends BaseEntity {
  public content?: string | null;
  public region?: string | null;
  public title?: string | null;
  public languageId?: string | null;
  public keyGroup?: "SMS_OTP" | "EMAIL_OTP" | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/regional-content/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/regional-content/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/regional-content/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/regional-contents`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "regional-content/edit/:uniqueId",
      Rcreate: "regional-content/new",
      Rsingle: "regional-content/:uniqueId",
      Rquery: "regional-contents",
  };
  public static definition = {
  "name": "regionalContent",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "content",
      "type": "html",
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
    },
    {
      "name": "title",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "languageId",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "keyGroup",
      "type": "enum",
      "validate": "required",
      "of": [
        {
          "k": "SMS_OTP"
        },
        {
          "k": "EMAIL_OTP"
        }
      ],
      "computedType": "\"SMS_OTP\" | \"EMAIL_OTP\"",
      "gormMap": {}
    }
  ],
  "cliShort": "rc",
  "cliDescription": "Email templates, sms templates or other textual content which can be accessed."
}
public static Fields = {
  ...BaseEntity.Fields,
      content: 'content',
      region: 'region',
      title: 'title',
      languageId: 'languageId',
      keyGroup: 'keyGroup',
}
}
