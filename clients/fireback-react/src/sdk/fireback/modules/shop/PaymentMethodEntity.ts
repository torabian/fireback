import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type PaymentMethodEntityKeys =
  keyof typeof PaymentMethodEntity.Fields;
export class PaymentMethodEntity extends BaseEntity {
  public children?: PaymentMethodEntity[] | null;
  public name?: string | null;
  public description?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/payment-method/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/payment-method/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/payment-method/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/payment-methods`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "payment-method/edit/:uniqueId",
      Rcreate: "payment-method/new",
      Rsingle: "payment-method/:uniqueId",
      Rquery: "payment-methods",
  };
  public static definition = {
  "name": "paymentMethod",
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
  "cliDescription": "Method of payment"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      description: 'description',
}
}