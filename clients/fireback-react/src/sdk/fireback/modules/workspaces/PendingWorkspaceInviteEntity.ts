import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    RoleEntity,
} from "./RoleEntity"
// In this section we have sub entities related to this object
// Class body
export type PendingWorkspaceInviteEntityKeys =
  keyof typeof PendingWorkspaceInviteEntity.Fields;
export class PendingWorkspaceInviteEntity extends BaseEntity {
  public value?: string | null;
  public type?: string | null;
  public coverLetter?: string | null;
  public workspaceName?: string | null;
  public role?: RoleEntity | null;
      roleId?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/pending-workspace-invite/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/pending-workspace-invite/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/pending-workspace-invite/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/pending-workspace-invites`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "pending-workspace-invite/edit/:uniqueId",
      Rcreate: "pending-workspace-invite/new",
      Rsingle: "pending-workspace-invite/:uniqueId",
      Rquery: "pending-workspace-invites",
  };
  public static definition = {
  "name": "pendingWorkspaceInvite",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/workspaces/RoleDefinitions.dyno.proto"
  ],
  "fields": [
    {
      "name": "value",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "type",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "coverLetter",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "workspaceName",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "role",
      "type": "one",
      "target": "RoleEntity",
      "computedType": "RoleEntity",
      "gormMap": {}
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      value: 'value',
      type: 'type',
      coverLetter: 'coverLetter',
      workspaceName: 'workspaceName',
          roleId: 'roleId',
      role$: 'role',
      role: RoleEntity.Fields,
}
}
