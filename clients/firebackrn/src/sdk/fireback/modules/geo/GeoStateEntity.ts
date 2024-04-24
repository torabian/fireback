import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type GeoStateEntityKeys =
  keyof typeof GeoStateEntity.Fields;
export class GeoStateEntity extends BaseEntity {
  public children?: GeoStateEntity[] | null;
  public name?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-state/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-state/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-state/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/geo-states`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "geo-state/edit/:uniqueId",
      Rcreate: "geo-state/new",
      Rsingle: "geo-state/:uniqueId",
      Rquery: "geo-states",
  };
  public static definition = {
  "name": "geoState",
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
  "cliName": "state"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
}
}