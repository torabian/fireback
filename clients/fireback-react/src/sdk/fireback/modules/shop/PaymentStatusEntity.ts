import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type PaymentStatusEntityKeys =
  keyof typeof PaymentStatusEntity.Fields;
export class PaymentStatusEntity extends BaseEntity {
  public children?: PaymentStatusEntity[] | null;
  public name?: string | null;
  public description?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/payment-status/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/payment-status/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/payment-status/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/payment-statuses`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "payment-status/edit/:uniqueId",
      Rcreate: "payment-status/new",
      Rsingle: "payment-status/:uniqueId",
      Rquery: "payment-statuses",
  };
  public static definition = {
  "name": "paymentStatus",
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
  "cliDescription": "Status of an payment"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      description: 'description',
}
}