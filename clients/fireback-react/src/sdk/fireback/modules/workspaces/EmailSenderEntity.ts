import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type EmailSenderEntityKeys =
  keyof typeof EmailSenderEntity.Fields;
export class EmailSenderEntity extends BaseEntity {
  public fromName?: string | null;
  public fromEmailAddress?: string | null;
  public replyTo?: string | null;
  public nickName?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/email-sender/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/email-sender/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/email-sender/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/email-senders`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "email-sender/edit/:uniqueId",
      Rcreate: "email-sender/new",
      Rsingle: "email-sender/:uniqueId",
      Rquery: "email-senders",
  };
  public static definition = {
  "name": "emailSender",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "fromName",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "fromEmailAddress",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "replyTo",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "nickName",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      fromName: 'fromName',
      fromEmailAddress: 'fromEmailAddress',
      replyTo: 'replyTo',
      nickName: 'nickName',
}
}
