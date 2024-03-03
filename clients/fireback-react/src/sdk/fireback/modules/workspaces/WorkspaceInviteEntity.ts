import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    RoleEntity,
} from "./RoleEntity"
import {
    WorkspaceEntity,
} from "./WorkspaceEntity"
// In this section we have sub entities related to this object
// Class body
export type WorkspaceInviteEntityKeys =
  keyof typeof WorkspaceInviteEntity.Fields;
export class WorkspaceInviteEntity extends BaseEntity {
  public coverLetter?: string | null;
  public targetUserLocale?: string | null;
  public value?: string | null;
  public workspace?: WorkspaceEntity | null;
  public firstName?: string | null;
  public lastName?: string | null;
  public used?: boolean | null;
  public role?: RoleEntity | null;
      roleId?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-invite/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-invite/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-invite/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-invites`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "workspace-invite/edit/:uniqueId",
      Rcreate: "workspace-invite/new",
      Rsingle: "workspace-invite/:uniqueId",
      Rquery: "workspace-invites",
  };
  public static definition = {
  "name": "workspaceInvite",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "coverLetter",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "targetUserLocale",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "value",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "workspace",
      "type": "one",
      "target": "WorkspaceEntity",
      "validate": "required",
      "computedType": "WorkspaceEntity",
      "gormMap": {}
    },
    {
      "name": "firstName",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "lastName",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "used",
      "type": "bool",
      "computedType": "boolean",
      "gormMap": {}
    },
    {
      "name": "role",
      "type": "one",
      "target": "RoleEntity",
      "validate": "required",
      "computedType": "RoleEntity",
      "gormMap": {}
    }
  ],
  "cliDescription": "Active invitations for non-users or already users to join an specific workspace"
}
public static Fields = {
  ...BaseEntity.Fields,
      coverLetter: 'coverLetter',
      targetUserLocale: 'targetUserLocale',
      value: 'value',
      workspace$: 'workspace',
      workspace: WorkspaceEntity.Fields,
      firstName: 'firstName',
      lastName: 'lastName',
      used: 'used',
          roleId: 'roleId',
      role$: 'role',
      role: RoleEntity.Fields,
}
}
