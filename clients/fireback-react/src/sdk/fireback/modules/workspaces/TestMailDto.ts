import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type TestMailDtoKeys =
  keyof typeof TestMailDto.Fields;
export class TestMailDto extends BaseDto {
  public senderId?: string | null;
  public toName?: string | null;
  public toEmail?: string | null;
  public subject?: string | null;
  public content?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      senderId: 'senderId',
      toName: 'toName',
      toEmail: 'toEmail',
      subject: 'subject',
      content: 'content',
}
  public static definition = {
  "name": "testMail",
  "fields": [
    {
      "name": "senderId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "toName",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "toEmail",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "subject",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "content",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
}
