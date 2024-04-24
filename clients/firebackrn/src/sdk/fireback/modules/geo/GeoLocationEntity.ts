import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    GeoLocationTypeEntity,
} from "./GeoLocationTypeEntity"
// In this section we have sub entities related to this object
// Class body
export type GeoLocationEntityKeys =
  keyof typeof GeoLocationEntity.Fields;
export class GeoLocationEntity extends BaseEntity {
  public children?: GeoLocationEntity[] | null;
  public name?: string | null;
  public code?: string | null;
  public type?: GeoLocationTypeEntity | null;
      typeId?: string | null;
  public status?: string | null;
  public flag?: string | null;
  public officialName?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-location/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-location/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-location/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-locations`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "geo-location/edit/:uniqueId",
      Rcreate: "geo-location/new",
      Rsingle: "geo-location/:uniqueId",
      Rquery: "geo-locations",
  };
  public static definition = {
  "name": "geoLocation",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/geo/GeoLocationTypeDefinitions.dyno.proto"
  ],
  "fields": [
    {
      "name": "name",
      "type": "string",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "code",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "type",
      "type": "one",
      "target": "GeoLocationTypeEntity",
      "computedType": "GeoLocationTypeEntity",
      "gormMap": {}
    },
    {
      "name": "status",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "flag",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "officialName",
      "type": "string",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliName": "location",
  "cte": true
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      code: 'code',
          typeId: 'typeId',
      type$: 'type',
        type: GeoLocationTypeEntity.Fields,
      status: 'status',
      flag: 'flag',
      officialName: 'officialName',
}
}