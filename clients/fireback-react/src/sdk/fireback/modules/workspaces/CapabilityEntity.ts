import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type CapabilityEntityKeys =
  keyof typeof CapabilityEntity.Fields;
export class CapabilityEntity extends BaseEntity {
  public name?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/capability/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/capability/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/capability/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/capabilities`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "capability/edit/:uniqueId",
      Rcreate: "capability/new",
      Rsingle: "capability/:uniqueId",
      Rquery: "capabilities",
  };
  public static definition = {
  "name": "capability",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "name",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliShort": "cap",
  "cliDescription": "Manage the capabilities inside the application, both builtin to core and custom defined ones"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
}
}
