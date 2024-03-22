import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type OrderStatusEntityKeys =
  keyof typeof OrderStatusEntity.Fields;
export class OrderStatusEntity extends BaseEntity {
  public children?: OrderStatusEntity[] | null;
  public name?: string | null;
  public description?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/order-status/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/order-status/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/order-status/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/order-statuses`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "order-status/edit/:uniqueId",
      Rcreate: "order-status/new",
      Rsingle: "order-status/:uniqueId",
      Rquery: "order-statuses",
  };
  public static definition = {
  "name": "orderStatus",
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
      "name": "description",
      "type": "string",
      "translate": true,
      "computedType": "string",
      "gormMap": {}
    }
  ],
  "cliDescription": "Status of an order"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      description: 'description',
}
}