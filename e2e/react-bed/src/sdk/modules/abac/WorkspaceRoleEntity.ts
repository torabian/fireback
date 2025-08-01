/*
*	Generated by fireback 1.2.3
*	Written by Ali Torabi.
*	Checkout the repository for licenses and contribution: https://github.com/torabian/fireback
*/
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
  public children?: WorkspaceRoleEntity[] | null;
  public userWorkspace?: UserWorkspaceEntity | null;
      userWorkspaceId?: string | null;
  public role?: RoleEntity | null;
      roleId?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : '..'}/workspace-role/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : '..'}/workspace-role/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : '..'}/workspace-role/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : '..'}/workspace-roles`;
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
  "rpc": {
    "query": {}
  },
  "permRewrite": {
    "replace": "root.modules",
    "with": "root.manage"
  },
  "name": "workspaceRole",
  "features": {},
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
  "description": "Manage roles assigned to an specific workspace or created by the workspace itself"
}
public static Fields = {
  ...BaseEntity.Fields,
          userWorkspaceId: `userWorkspaceId`,
      userWorkspace$: `userWorkspace`,
        userWorkspace: UserWorkspaceEntity.Fields,
          roleId: `roleId`,
      role$: `role`,
        role: RoleEntity.Fields,
}
}
