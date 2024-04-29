import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    UserAccessLevelDto,
} from "./UserAccessLevelDto"
import {
    UserEntity,
} from "./UserEntity"
import {
    UserRoleWorkspacePermission,
} from "./UserRoleWorkspacePermission"
// In this section we have sub entities related to this object
// Class body
export type AuthResultDtoKeys =
  keyof typeof AuthResultDto.Fields;
export class AuthResultDto extends BaseDto {
  public workspaceId?: string | null;
  public userRoleWorkspacePermissions?: UserRoleWorkspacePermission[] | null;
    userRoleWorkspacePermissionsListId?: string[] | null;
  public internalSql?: string | null;
  public userId?: string | null;
  public userHas?: string[] | null;
  public workspaceHas?: string[] | null;
  public user?: UserEntity | null;
  public accessLevel?: UserAccessLevelDto | null;
      accessLevelId?: string | null;
public static Fields = {
  ...BaseEntity.Fields,
      workspaceId: 'workspaceId',
        userRoleWorkspacePermissionsListId: 'userRoleWorkspacePermissionsListId',
      userRoleWorkspacePermissions$: 'userRoleWorkspacePermissions',
        userRoleWorkspacePermissions: UserRoleWorkspacePermission.Fields,
      internalSql: 'internalSql',
      userId: 'userId',
      userHas: 'userHas',
      workspaceHas: 'workspaceHas',
      user$: 'user',
        user: UserEntity.Fields,
          accessLevelId: 'accessLevelId',
      accessLevel$: 'accessLevel',
        accessLevel: UserAccessLevelDto.Fields,
}
  public static definition = {
  "name": "authResult",
  "fields": [
    {
      "name": "workspaceId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "userRoleWorkspacePermissions",
      "type": "many2many",
      "target": "UserRoleWorkspacePermission",
      "computedType": "UserRoleWorkspacePermission[]",
      "gormMap": {}
    },
    {
      "name": "internalSql",
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
      "name": "userHas",
      "type": "arrayP",
      "primitive": "string",
      "computedType": "string[]",
      "gormMap": {}
    },
    {
      "name": "workspaceHas",
      "type": "arrayP",
      "primitive": "string",
      "computedType": "string[]",
      "gormMap": {}
    },
    {
      "name": "user",
      "type": "one",
      "target": "UserEntity",
      "computedType": "UserEntity",
      "gormMap": {}
    },
    {
      "name": "accessLevel",
      "type": "one",
      "target": "UserAccessLevelDto",
      "computedType": "UserAccessLevelDto",
      "gormMap": {}
    }
  ]
}
}
