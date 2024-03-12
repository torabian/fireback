import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type WidgetEntityKeys =
  keyof typeof WidgetEntity.Fields;
export class WidgetEntity extends BaseEntity {
  public children?: WidgetEntity[] | null;
  public name?: string | null;
  public family?: string | null;
  public providerKey?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/widget/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/widget/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/widget/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/widgets`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "widget/edit/:uniqueId",
      Rcreate: "widget/new",
      Rsingle: "widget/:uniqueId",
      Rquery: "widgets",
  };
  public static definition = {
  "name": "widget",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "name",
      "type": "string",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "family",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "providerKey",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliDescription": "Widget is an item which can be placed on a widget area, such as weather widget"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      family: 'family',
      providerKey: 'providerKey',
}
}