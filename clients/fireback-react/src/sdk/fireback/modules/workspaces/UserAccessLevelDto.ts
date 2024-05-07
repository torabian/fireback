import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    UserRoleWorkspacePermissionDto,
} from "./UserRoleWorkspacePermissionDto"
// In this section we have sub entities related to this object
// Class body
export type UserAccessLevelDtoKeys =
  keyof typeof UserAccessLevelDto.Fields;
export class UserAccessLevelDto extends BaseDto {
  public capabilities?: string[] | null;
  public userRoleWorkspacePermissions?: UserRoleWorkspacePermissionDto[] | null;
    userRoleWorkspacePermissionsListId?: string[] | null;
  public workspaces?: string[] | null;
  public SQL?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      capabilities: 'capabilities',
        userRoleWorkspacePermissionsListId: 'userRoleWorkspacePermissionsListId',
      userRoleWorkspacePermissions$: 'userRoleWorkspacePermissions',
        userRoleWorkspacePermissions: UserRoleWorkspacePermissionDto.Fields,
      workspaces: 'workspaces',
      SQL: 'SQL',
}
  public static definition = {
  "name": "userAccessLevel",
  "fields": [
    {
      "name": "capabilities",
      "type": "arrayP",
      "primitive": "string",
      "computedType": "string[]",
      "gormMap": {}
    },
    {
      "name": "userRoleWorkspacePermissions",
      "type": "many2many",
      "target": "UserRoleWorkspacePermissionDto",
      "computedType": "UserRoleWorkspacePermissionDto[]",
      "gormMap": {}
    },
    {
      "name": "workspaces",
      "type": "arrayP",
      "primitive": "string",
      "computedType": "string[]",
      "gormMap": {}
    },
    {
      "name": "SQL",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
}
