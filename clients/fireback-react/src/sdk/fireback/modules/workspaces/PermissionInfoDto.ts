import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type PermissionInfoDtoKeys =
  keyof typeof PermissionInfoDto.Fields;
export class PermissionInfoDto extends BaseDto {
  public name?: string | null;
  public description?: string | null;
  public completeKey?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      description: 'description',
      completeKey: 'completeKey',
}
  public static definition = {
  "name": "permissionInfo",
  "fields": [
    {
      "name": "name",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "description",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "completeKey",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
}
