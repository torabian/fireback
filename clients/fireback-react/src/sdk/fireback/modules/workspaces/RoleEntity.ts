import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    CapabilityEntity,
} from "./CapabilityEntity"
// In this section we have sub entities related to this object
// Class body
export type RoleEntityKeys =
  keyof typeof RoleEntity.Fields;
export class RoleEntity extends BaseEntity {
  public name?: string | null;
  public capabilities?: CapabilityEntity[] | null;
    capabilitiesListId?: string[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/role/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/role/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/role/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/roles`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "role/edit/:uniqueId",
      Rcreate: "role/new",
      Rsingle: "role/:uniqueId",
      Rquery: "roles",
  };
  public static definition = {
  "name": "role",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "name",
      "type": "string",
      "validate": "required,omitempty,min=1,max=200",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "capabilities",
      "type": "many2many",
      "target": "CapabilityEntity",
      "computedType": "CapabilityEntity[]",
      "gormMap": {}
    }
  ],
  "cliDescription": "Manage roles within the workspaces, or root configuration"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
        capabilitiesListId: 'capabilitiesListId',
      capabilities$: 'capabilities',
      capabilities: CapabilityEntity.Fields,
}
}
