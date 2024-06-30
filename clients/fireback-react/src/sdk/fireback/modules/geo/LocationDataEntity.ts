import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type LocationDataEntityKeys =
  keyof typeof LocationDataEntity.Fields;
export class LocationDataEntity extends BaseEntity {
  public children?: LocationDataEntity[] | null;
  public lat?: number | null;
  public lng?: number | null;
  public physicalAddress?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/location-data/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/location-data/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/location-data/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/location-data`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "location-data/edit/:uniqueId",
      Rcreate: "location-data/new",
      Rsingle: "location-data/:uniqueId",
      Rquery: "location-data",
  };
  public static definition = {
  "name": "locationData",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "lat",
      "type": "float64",
      "computedType": "number",
      "gormMap": {}
    },
    {
      "name": "lng",
      "type": "float64",
      "computedType": "number",
      "gormMap": {}
    },
    {
      "name": "physicalAddress",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliName": "location"
}
public static Fields = {
  ...BaseEntity.Fields,
      lat: 'lat',
      lng: 'lng',
      physicalAddress: 'physicalAddress',
}
}
