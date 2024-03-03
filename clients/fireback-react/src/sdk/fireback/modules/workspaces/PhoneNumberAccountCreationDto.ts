import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type PhoneNumberAccountCreationDtoKeys =
  keyof typeof PhoneNumberAccountCreationDto.Fields;
export class PhoneNumberAccountCreationDto extends BaseDto {
  public phoneNumber?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      phoneNumber: 'phoneNumber',
}
  public static definition = {
  "name": "phoneNumberAccountCreation",
  "fields": [
    {
      "name": "phoneNumber",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
}
