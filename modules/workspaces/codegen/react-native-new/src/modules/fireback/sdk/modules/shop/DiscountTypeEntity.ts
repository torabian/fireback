import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type DiscountTypeEntityKeys =
  keyof typeof DiscountTypeEntity.Fields;
export class DiscountTypeEntity extends BaseEntity {
  public children?: DiscountTypeEntity[] | null;
  public name?: string | null;
  public description?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/discount-type/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/discount-type/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/discount-type/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/discount-types`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "discount-type/edit/:uniqueId",
      Rcreate: "discount-type/new",
      Rsingle: "discount-type/:uniqueId",
      Rquery: "discount-types",
  };
  public static definition = {
  "name": "discountType",
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
  "cliDescription": "Types of the discounts"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      description: 'description',
}
}