import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type GeoCountryEntityKeys =
  keyof typeof GeoCountryEntity.Fields;
export class GeoCountryEntity extends BaseEntity {
  public children?: GeoCountryEntity[] | null;
  public status?: string | null;
  public flag?: string | null;
  public commonName?: string | null;
  public officialName?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-country/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-country/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-country/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-countries`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "geo-country/edit/:uniqueId",
      Rcreate: "geo-country/new",
      Rsingle: "geo-country/:uniqueId",
      Rquery: "geo-countries",
  };
  public static definition = {
  "name": "geoCountry",
  "http": {},
  "gormMap": {},
  "fields": [
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
      "name": "commonName",
      "type": "string",
      "translate": true,
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
  "cliName": "country"
}
public static Fields = {
  ...BaseEntity.Fields,
      status: 'status',
      flag: 'flag',
      commonName: 'commonName',
      officialName: 'officialName',
}
}