import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    RoleEntity,
} from "./RoleEntity"
import {
    UserWorkspaceEntity,
} from "./UserWorkspaceEntity"
// In this section we have sub entities related to this object
// Class body
export type WorkspaceRoleEntityKeys =
  keyof typeof WorkspaceRoleEntity.Fields;
export class WorkspaceRoleEntity extends BaseEntity {
  public userWorkspace?: UserWorkspaceEntity | null;
      userWorkspaceId?: string | null;
  public role?: RoleEntity | null;
      roleId?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-role/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-role/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-role/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-roles`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "workspace-role/edit/:uniqueId",
      Rcreate: "workspace-role/new",
      Rsingle: "workspace-role/:uniqueId",
      Rquery: "workspace-roles",
  };
  public static definition = {
  "name": "workspaceRole",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "userWorkspace",
      "type": "one",
      "target": "UserWorkspaceEntity",
      "idFieldGorm": "index:workspacerole_idx,unique",
      "computedType": "UserWorkspaceEntity",
      "gormMap": {}
    },
    {
      "name": "role",
      "type": "one",
      "target": "RoleEntity",
      "idFieldGorm": "index:workspacerole_idx,unique",
      "computedType": "RoleEntity",
      "gormMap": {}
    }
  ],
  "cliShort": "role",
  "cliDescription": "Manage roles assigned to an specific workspace or created by the workspace itself"
}
public static Fields = {
  ...BaseEntity.Fields,
          userWorkspaceId: 'userWorkspaceId',
      userWorkspace$: 'userWorkspace',
      userWorkspace: UserWorkspaceEntity.Fields,
          roleId: 'roleId',
      role$: 'role',
      role: RoleEntity.Fields,
}
}
