import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    UserEntity,
} from "./UserEntity"
import {
    WorkspaceEntity,
} from "./WorkspaceEntity"
// In this section we have sub entities related to this object
// Class body
export type UserWorkspaceEntityKeys =
  keyof typeof UserWorkspaceEntity.Fields;
export class UserWorkspaceEntity extends BaseEntity {
  public children?: UserWorkspaceEntity[] | null;
  public user?: UserEntity | null;
  public workspace?: WorkspaceEntity | null;
  public userPermissions?: string[] | null;
  public rolePermission?: unknown[] | null;
  public workspacePermissions?: string[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/user-workspace/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/user-workspace/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/user-workspace/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/user-workspaces`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "user-workspace/edit/:uniqueId",
      Rcreate: "user-workspace/new",
      Rsingle: "user-workspace/:uniqueId",
      Rquery: "user-workspaces",
  };
  public static definition = {
  "name": "userWorkspace",
  "security": {
    "ActionRequires": null,
    "resolveStrategy": "user"
  },
  "http": {},
  "gormMap": {
    "workspaceId": "index:userworkspace_idx,unique",
    "userId": "index:userworkspace_idx,unique"
  },
  "fields": [
    {
      "name": "user",
      "type": "one",
      "target": "UserEntity",
      "computedType": "UserEntity",
      "gormMap": {}
    },
    {
      "name": "workspace",
      "type": "one",
      "target": "WorkspaceEntity",
      "computedType": "WorkspaceEntity",
      "gormMap": {}
    },
    {
      "name": "userPermissions",
      "type": "arrayP",
      "primitive": "string",
      "computedType": "string[]",
      "gormMap": {},
      "sql": "-"
    },
    {
      "name": "rolePermission",
      "type": "arrayP",
      "primitive": "UserRoleWorkspaceDto",
      "computedType": "unknown[]",
      "gormMap": {},
      "sql": "-"
    },
    {
      "name": "workspacePermissions",
      "type": "arrayP",
      "primitive": "string",
      "computedType": "string[]",
      "gormMap": {},
      "sql": "-"
    }
  ],
  "cliShort": "user",
  "cliDescription": "Manage the workspaces that user belongs to (either its himselves or adding by invitation)",
  "postFormatter": "UserWorkspacePostFormatter"
}
public static Fields = {
  ...BaseEntity.Fields,
      user$: 'user',
        user: UserEntity.Fields,
      workspace$: 'workspace',
        workspace: WorkspaceEntity.Fields,
      userPermissions: 'userPermissions',
      rolePermission: 'rolePermission',
      workspacePermissions: 'workspacePermissions',
}
}