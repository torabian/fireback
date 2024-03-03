import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    WorkspaceEntity,
} from "./WorkspaceEntity"
// In this section we have sub entities related to this object
// Class body
export type WorkspaceConfigEntityKeys =
  keyof typeof WorkspaceConfigEntity.Fields;
export class WorkspaceConfigEntity extends BaseEntity {
  public disablePublicWorkspaceCreation?: number | null;
  public workspace?: WorkspaceEntity | null;
  public zoomClientId?: string | null;
  public zoomClientSecret?: string | null;
  public allowPublicToJoinTheWorkspace?: boolean | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-config/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-config/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-config/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-configs`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "workspace-config/edit/:uniqueId",
      Rcreate: "workspace-config/new",
      Rsingle: "workspace-config/:uniqueId",
      Rquery: "workspace-configs",
  };
  public static definition = {
  "name": "workspaceConfig",
  "distinctBy": "workspace",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "disablePublicWorkspaceCreation",
      "type": "int64",
      "default": 1,
      "computedType": "number",
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
      "name": "zoomClientId",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "zoomClientSecret",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "allowPublicToJoinTheWorkspace",
      "type": "bool",
      "computedType": "boolean",
      "gormMap": {}
    }
  ],
  "cliName": "config"
}
public static Fields = {
  ...BaseEntity.Fields,
      disablePublicWorkspaceCreation: 'disablePublicWorkspaceCreation',
      workspace$: 'workspace',
      workspace: WorkspaceEntity.Fields,
      zoomClientId: 'zoomClientId',
      zoomClientSecret: 'zoomClientSecret',
      allowPublicToJoinTheWorkspace: 'allowPublicToJoinTheWorkspace',
}
}
