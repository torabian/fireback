import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type UserRoleWorkspacePermissionDtoKeys =
  keyof typeof UserRoleWorkspacePermissionDto.Fields;
export class UserRoleWorkspacePermissionDto extends BaseDto {
  public workspaceId?: string | null;
  public userId?: string | null;
  public roleId?: string | null;
  public capabilityId?: string | null;
  public type?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      workspaceId: 'workspaceId',
      userId: 'userId',
      roleId: 'roleId',
      capabilityId: 'capabilityId',
      type: 'type',
}
  public static definition = {
  "name": "userRoleWorkspacePermission",
  "fields": [
    {
      "name": "workspaceId",
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
      "name": "roleId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "capabilityId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "type",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
}
