import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type EmailAccountSigninDtoKeys =
  keyof typeof EmailAccountSigninDto.Fields;
export class EmailAccountSigninDto extends BaseDto {
  public email?: string | null;
  public password?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      email: 'email',
      password: 'password',
}
  public static definition = {
  "name": "emailAccountSignin",
  "fields": [
    {
      "name": "email",
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
