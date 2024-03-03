import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type ExchangeKeyInformationDtoKeys =
  keyof typeof ExchangeKeyInformationDto.Fields;
export class ExchangeKeyInformationDto extends BaseDto {
  public key?: string | null;
  public visibility?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      key: 'key',
      visibility: 'visibility',
}
  public static definition = {
  "name": "exchangeKeyInformation",
  "fields": [
    {
      "name": "key",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "visibility",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
}
