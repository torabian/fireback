import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type PreferenceEntityKeys =
  keyof typeof PreferenceEntity.Fields;
export class PreferenceEntity extends BaseEntity {
  public timezone?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/preference/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/preference/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/preference/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/preferences`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "preference/edit/:uniqueId",
      Rcreate: "preference/new",
      Rsingle: "preference/:uniqueId",
      Rquery: "preferences",
  };
  public static definition = {
  "name": "preference",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "timezone",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ]
}
public static Fields = {
  ...BaseEntity.Fields,
      timezone: 'timezone',
}
}
