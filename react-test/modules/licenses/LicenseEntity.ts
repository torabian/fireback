import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    CapabilityEntity,
} from "../workspaces/CapabilityEntity"
// In this section we have sub entities related to this object
export class LicensePermissions extends BaseEntity {
  public capability?: CapabilityEntity | null;
      capabilityId?: string | null;
}
// Class body
export type LicenseEntityKeys =
  keyof typeof LicenseEntity.Fields;
export class LicenseEntity extends BaseEntity {
  public children?: LicenseEntity[] | null;
  public name?: string | null;
  public signedLicense?: string | null;
  public validityStartDate?: Date | null;
  public validityEndDate?: Date | null;
  public permissions?: LicensePermissions[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/license/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/license/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/license/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/licenses`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "license/edit/:uniqueId",
      Rcreate: "license/new",
      Rsingle: "license/:uniqueId",
      Rquery: "licenses",
      rPermissionsCreate: "license/:linkerId/permissions/new",
      rPermissionsEdit: "license/:linkerId/permissions/edit/:uniqueId",
      editPermissions(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/license/${linkerId}/permissions/edit/${uniqueId}`;
      },
      createPermissions(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/license/${linkerId}/permissions/new`;
      },
  };
  public static definition = {
  "name": "license",
  "queryScope": "specific",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/workspaces/CapabilityDefinitions.dyno.proto"
  ],
  "fields": [
    {
      "name": "name",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "signedLicense",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "validityStartDate",
      "type": "date",
      "computedType": "Date",
      "gormMap": {}
    },
    {
      "name": "validityEndDate",
      "type": "date",
      "computedType": "Date",
      "gormMap": {}
    },
    {
      "linkedTo": "LicenseEntity",
      "name": "permissions",
      "type": "array",
      "computedType": "LicensePermissions[]",
      "gormMap": {},
      "fullName": "LicensePermissions",
      "fields": [
        {
          "name": "capability",
          "type": "one",
          "target": "CapabilityEntity",
          "module": "workspaces",
          "computedType": "CapabilityEntity",
          "gormMap": {}
        }
      ]
    }
  ],
  "cliDescription": "Manage the licenses in the app (either to issue, or to activate current product)"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      signedLicense: 'signedLicense',
      validityStartDate: 'validityStartDate',
      validityEndDate: 'validityEndDate',
      permissions$: 'permissions',
      permissions: {
  ...BaseEntity.Fields,
          capabilityId: 'capabilityId',
      capability$: 'capability',
      capability: CapabilityEntity.Fields,
      },
}
}