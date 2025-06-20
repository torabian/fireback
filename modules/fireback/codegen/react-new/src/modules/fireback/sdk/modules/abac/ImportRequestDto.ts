/*
*	Generated by fireback 1.2.3
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
*/
    import {
        BaseDto,
        BaseEntity,
    } from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type ImportRequestDtoKeys =
  keyof typeof ImportRequestDto.Fields;
export class ImportRequestDto extends BaseDto {
  public file?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      file: `file`,
}
  public static definition = {
  "name": "importRequest",
  "fields": [
    {
      "name": "file",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
}
