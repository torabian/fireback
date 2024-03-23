import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
// In this section we have sub entities related to this object
// Class body
export type DiscountScopeEntityKeys =
  keyof typeof DiscountScopeEntity.Fields;
export class DiscountScopeEntity extends BaseEntity {
  public children?: DiscountScopeEntity[] | null;
  public name?: string | null;
  public description?: string | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/discount-scope/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/discount-scope/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/discount-scope/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/discount-scopes`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "discount-scope/edit/:uniqueId",
      Rcreate: "discount-scope/new",
      Rsingle: "discount-scope/:uniqueId",
      Rquery: "discount-scopes",
  };
  public static definition = {
  "name": "discountScope",
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
  "cliDescription": "Determine if the discount applies to the entire basket (total order) or per item, etc"
}
public static Fields = {
  ...BaseEntity.Fields,
      name: 'name',
      description: 'description',
}
}