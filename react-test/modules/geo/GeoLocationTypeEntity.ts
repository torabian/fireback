import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type GeoLocationTypeEntityKeys =
  keyof typeof GeoLocationTypeEntity.Fields;
export class GeoLocationTypeEntity extends BaseEntity {
  public children?: GeoLocationTypeEntity[] | null;
  public name?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-location-type/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-location-type/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-location-type/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-location-types`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "geo-location-type/edit/:uniqueId",
      Rcreate: "geo-location-type/new",
      Rsingle: "geo-location-type/:uniqueId",
      Rquery: "geo-location-types",
  };
  public static definition = {
  "name": "geoLocationType",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "name",
      "type": "string",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliName": "type"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
}
}