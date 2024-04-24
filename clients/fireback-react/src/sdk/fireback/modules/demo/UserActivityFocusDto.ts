import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type UserActivityFocusDtoKeys =
  keyof typeof UserActivityFocusDto.Fields;
export class UserActivityFocusDto extends BaseDto {
  public ids?: string[] | null;
public static Fields = {
  ...BaseEntity.Fields,
      ids: 'ids',
}
  public static definition = {
  "name": "userActivityFocus",
  "fields": [
    {
      "name": "ids",
      "type": "arrayP",
      "primitive": "string",
      "computedType": "string[]",
      "gormMap": {}
    }
  ]
}
}
