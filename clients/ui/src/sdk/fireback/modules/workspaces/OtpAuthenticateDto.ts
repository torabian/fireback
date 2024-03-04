import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type OtpAuthenticateDtoKeys =
  keyof typeof OtpAuthenticateDto.Fields;
export class OtpAuthenticateDto extends BaseDto {
  public value?: string | null;
  public otp?: string | null;
  public type?: string | null;
  public password?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      value: 'value',
      otp: 'otp',
      type: 'type',
      password: 'password',
}
  public static definition = {
  "name": "otpAuthenticate",
  "fields": [
    {
      "name": "value",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "otp",
      "type": "string",
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
      "name": "password",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
}
