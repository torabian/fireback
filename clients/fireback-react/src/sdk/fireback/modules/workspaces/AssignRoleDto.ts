import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type AssignRoleDtoKeys =
  keyof typeof AssignRoleDto.Fields;
export class AssignRoleDto extends BaseDto {
  public roleId?: string | null;
  public userId?: string | null;
  public visibility?: string | null;
  public updated?: number | null;
  public created?: number | null;
public static Fields = {
  ...BaseEntity.Fields,
      roleId: 'roleId',
      userId: 'userId',
      visibility: 'visibility',
      updated: 'updated',
      created: 'created',
}
  public static definition = {
  "name": "assignRole",
  "fields": [
    {
      "name": "roleId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "userId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "visibility",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "updated",
      "type": "int64",
      "computedType": "number",
      "gormMap": {}
    },
    {
      "name": "created",
      "type": "int64",
      "computedType": "number",
      "gormMap": {}
    }
  ]
}
}
