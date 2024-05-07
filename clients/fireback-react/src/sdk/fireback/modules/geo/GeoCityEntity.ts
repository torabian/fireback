import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    GeoCountryEntity,
} from "./GeoCountryEntity"
import {
    GeoProvinceEntity,
} from "./GeoProvinceEntity"
import {
    GeoStateEntity,
} from "./GeoStateEntity"
// In this section we have sub entities related to this object
// Class body
export type GeoCityEntityKeys =
  keyof typeof GeoCityEntity.Fields;
export class GeoCityEntity extends BaseEntity {
  public children?: GeoCityEntity[] | null;
  public name?: string | null;
  public province?: GeoProvinceEntity | null;
      provinceId?: string | null;
  public state?: GeoStateEntity | null;
      stateId?: string | null;
  public country?: GeoCountryEntity | null;
      countryId?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-city/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-city/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-city/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-cities`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "geo-city/edit/:uniqueId",
      Rcreate: "geo-city/new",
      Rsingle: "geo-city/:uniqueId",
      Rquery: "geo-cities",
  };
  public static definition = {
  "name": "geoCity",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/geo/GeoStateDefinitions.dyno.proto",
    "modules/geo/GeoProvinceDefinitions.dyno.proto",
    "modules/geo/GeoCountryDefinitions.dyno.proto"
  ],
  "fields": [
    {
      "name": "name",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "province",
      "type": "one",
      "target": "GeoProvinceEntity",
      "computedType": "GeoProvinceEntity",
      "gormMap": {}
    },
    {
      "name": "state",
      "type": "one",
      "target": "GeoStateEntity",
      "computedType": "GeoStateEntity",
      "gormMap": {}
    },
    {
      "name": "country",
      "type": "one",
      "target": "GeoCountryEntity",
      "computedType": "GeoCountryEntity",
      "gormMap": {}
    }
  ],
  "cliName": "city"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
          provinceId: 'provinceId',
      province$: 'province',
        province: GeoProvinceEntity.Fields,
          stateId: 'stateId',
      state$: 'state',
        state: GeoStateEntity.Fields,
          countryId: 'countryId',
      country$: 'country',
        country: GeoCountryEntity.Fields,
}
}