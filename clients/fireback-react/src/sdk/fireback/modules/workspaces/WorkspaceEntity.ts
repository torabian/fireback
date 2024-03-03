import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type WorkspaceEntityKeys =
  keyof typeof WorkspaceEntity.Fields;
export class WorkspaceEntity extends BaseEntity {
  public description?: string | null;
  public name?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspaces`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "workspace/edit/:uniqueId",
      Rcreate: "workspace/new",
      Rsingle: "workspace/:uniqueId",
      Rquery: "workspaces",
  };
  public static definition = {
  "name": "workspace",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/workspaces/CapabilityDefinitions.dyno.proto"
  ],
  "fields": [
    {
      "name": "description",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "name",
      "type": "string",
      "validate": "required",
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliName": "ws",
  "cte": true
}
public static Fields = {
  ...BaseEntity.Fields,
      description: 'description',
      name: 'name',
}
}
