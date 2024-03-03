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
export type PublicJoinKeyEntityKeys =
  keyof typeof PublicJoinKeyEntity.Fields;
export class PublicJoinKeyEntity extends BaseEntity {
  public role?: RoleEntity | null;
      roleId?: string | null;
  public workspace?: WorkspaceEntity | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/public-join-key/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/public-join-key/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/public-join-key/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/public-join-keys`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "public-join-key/edit/:uniqueId",
      Rcreate: "public-join-key/new",
      Rsingle: "public-join-key/:uniqueId",
      Rquery: "public-join-keys",
  };
  public static definition = {
  "name": "publicJoinKey",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/workspaces/RoleDefinitions.dyno.proto",
    "modules/workspaces/WorkspaceDefinitions.dyno.proto"
  ],
  "fields": [
    {
      "name": "role",
      "type": "one",
      "target": "RoleEntity",
      "computedType": "RoleEntity",
      "gormMap": {}
    },
    {
      "name": "workspace",
      "type": "one",
      "target": "WorkspaceEntity",
      "computedType": "WorkspaceEntity",
      "gormMap": {}
    }
  ],
  "cliDescription": "Joining to different workspaces using a public link directly"
}
public static Fields = {
  ...BaseEntity.Fields,
          roleId: 'roleId',
      role$: 'role',
      role: RoleEntity.Fields,
      workspace$: 'workspace',
      workspace: WorkspaceEntity.Fields,
}
}
