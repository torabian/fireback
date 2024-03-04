import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
export class TimezoneGroupUtcItems extends BaseEntity {
  public name?: string | null;
}
// Class body
export type TimezoneGroupEntityKeys =
  keyof typeof TimezoneGroupEntity.Fields;
export class TimezoneGroupEntity extends BaseEntity {
  public children?: TimezoneGroupEntity[] | null;
  public value?: string | null;
  public abbr?: string | null;
  public offset?: number | null;
  public isdst?: boolean | null;
  public text?: string | null;
  public utcItems?: TimezoneGroupUtcItems[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/timezone-group/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/timezone-group/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/timezone-group/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/timezone-groups`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "timezone-group/edit/:uniqueId",
      Rcreate: "timezone-group/new",
      Rsingle: "timezone-group/:uniqueId",
      Rquery: "timezone-groups",
      rUtcItemsCreate: "timezone-group/:linkerId/utc_items/new",
      rUtcItemsEdit: "timezone-group/:linkerId/utc_items/edit/:uniqueId",
      editUtcItems(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/timezone-group/${linkerId}/utc_items/edit/${uniqueId}`;
      },
      createUtcItems(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/timezone-group/${linkerId}/utc_items/new`;
      },
  };
  public static definition = {
  "name": "timezoneGroup",
  "queryScope": "public",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "value",
      "type": "string",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "abbr",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "offset",
      "type": "int64",
      "computedType": "number",
      "gormMap": {}
    },
    {
      "name": "isdst",
      "type": "bool",
      "computedType": "boolean",
      "gormMap": {}
    },
    {
      "name": "text",
      "type": "string",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    },
    {
      "linkedTo": "TimezoneGroupEntity",
      "name": "utcItems",
      "type": "array",
      "computedType": "TimezoneGroupUtcItems[]",
      "gormMap": {},
      "fullName": "TimezoneGroupUtcItems",
      "fields": [
        {
          "name": "name",
          "type": "string",
          "validate": "required",
          "translate": true,
          "computedType": "string",
          "gormMap": {}
        }
      ]
    }
  ],
  "cliName": "tz"
}
public static Fields = {
  ...BaseEntity.Fields,
      value: 'value',
      abbr: 'abbr',
      offset: 'offset',
      isdst: 'isdst',
      text: 'text',
      utcItems$: 'utcItems',
      utcItems: {
  ...BaseEntity.Fields,
      name: 'name',
      },
}
}
