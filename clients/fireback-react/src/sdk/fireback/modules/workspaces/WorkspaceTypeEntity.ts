import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    CapabilityEntity,
} from "./CapabilityEntity"
// In this section we have sub entities related to this object
// Class body
export type WorkspaceTypeEntityKeys =
  keyof typeof WorkspaceTypeEntity.Fields;
export class WorkspaceTypeEntity extends BaseEntity {
  public children?: WorkspaceTypeEntity[] | null;
  public title?: string | null;
  public capabilities?: CapabilityEntity[] | null;
    capabilitiesListId?: string[] | null;
  public description?: string | null;
  public slug?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-type/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-type/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-type/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/workspace-types`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "workspace-type/edit/:uniqueId",
      Rcreate: "workspace-type/new",
      Rsingle: "workspace-type/:uniqueId",
      Rquery: "workspace-types",
  };
  public static definition = {
  "name": "workspaceType",
  "distinctBy": "workspace",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "title",
      "type": "string",
      "validate": "required,omitempty,min=1,max=250",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "capabilities",
      "type": "many2many",
      "target": "CapabilityEntity",
      "computedType": "CapabilityEntity[]",
      "gormMap": {}
    },
    {
      "name": "description",
      "type": "string",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "slug",
      "type": "string",
      "validate": "required,omitempty,min=2,max=50",
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliName": "type"
}
public static Fields = {
  ...BaseEntity.Fields,
      title: 'title',
        capabilitiesListId: 'capabilitiesListId',
      capabilities$: 'capabilities',
        capabilities: CapabilityEntity.Fields,
      description: 'description',
      slug: 'slug',
}
}