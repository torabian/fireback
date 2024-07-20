import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    GeoCountryEntity,
} from "./GeoCountryEntity"
// In this section we have sub entities related to this object
// Class body
export type GeoProvinceEntityKeys =
  keyof typeof GeoProvinceEntity.Fields;
export class GeoProvinceEntity extends BaseEntity {
  public children?: GeoProvinceEntity[] | null;
  public name?: string | null;
  public country?: GeoCountryEntity | null;
      countryId?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-province/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-province/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-province/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-provinces`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "geo-province/edit/:uniqueId",
      Rcreate: "geo-province/new",
      Rsingle: "geo-province/:uniqueId",
      Rquery: "geo-provinces",
  };
  public static definition = {
  "name": "geoProvince",
  "http": {},
  "gormMap": {},
  "importList": [
    "modules/geo/GeoCountryDefinitions.dyno.proto"
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
      "name": "country",
      "type": "one",
      "target": "GeoCountryEntity",
      "computedType": "GeoCountryEntity",
      "gormMap": {}
    }
  ],
  "cliName": "province"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
          countryId: 'countryId',
      country$: 'country',
        country: GeoCountryEntity.Fields,
}
}