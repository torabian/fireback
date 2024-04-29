import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    CategoryEntity,
} from "./CategoryEntity"
import {
    ProductSubmissionEntity,
} from "./ProductSubmissionEntity"
// In this section we have sub entities related to this object
// Class body
export type DiscountCodeEntityKeys =
  keyof typeof DiscountCodeEntity.Fields;
export class DiscountCodeEntity extends BaseEntity {
  public children?: DiscountCodeEntity[] | null;
  public series?: string | null;
  public limit?: number | null;
  public valid?: any | null;
    public validStart?: string[] | null;
    public validEnd?: string[] | null;
  public appliedProducts?: ProductSubmissionEntity[] | null;
    appliedProductsListId?: string[] | null;
  public excludedProducts?: ProductSubmissionEntity[] | null;
    excludedProductsListId?: string[] | null;
  public appliedCategories?: CategoryEntity[] | null;
    appliedCategoriesListId?: string[] | null;
  public excludedCategories?: CategoryEntity[] | null;
    excludedCategoriesListId?: string[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/discount-code/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/discount-code/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/discount-code/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/discount-codes`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "discount-code/edit/:uniqueId",
      Rcreate: "discount-code/new",
      Rsingle: "discount-code/:uniqueId",
      Rquery: "discount-codes",
  };
  public static definition = {
  "name": "discountCode",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "name": "series",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "limit",
      "type": "int64",
      "computedType": "number",
      "gormMap": {}
    },
    {
      "name": "valid",
      "type": "daterange",
      "computedType": "any",
      "gormMap": {}
    },
    {
      "name": "appliedProducts",
      "type": "many2many",
      "target": "ProductSubmissionEntity",
      "computedType": "ProductSubmissionEntity[]",
      "gormMap": {}
    },
    {
      "name": "excludedProducts",
      "type": "many2many",
      "target": "ProductSubmissionEntity",
      "computedType": "ProductSubmissionEntity[]",
      "gormMap": {}
    },
    {
      "name": "appliedCategories",
      "type": "many2many",
      "target": "CategoryEntity",
      "computedType": "CategoryEntity[]",
      "gormMap": {}
    },
    {
      "name": "excludedCategories",
      "type": "many2many",
      "target": "CategoryEntity",
      "computedType": "CategoryEntity[]",
      "gormMap": {}
    }
  ],
  "cliDescription": "List of all discount codes inside the application"
}
public static Fields = {
  ...BaseEntity.Fields,
      series: 'series',
      limit: 'limit',
      validStart: 'validStart',
      validEnd: 'validEnd',
      valid: 'valid',
        appliedProductsListId: 'appliedProductsListId',
      appliedProducts$: 'appliedProducts',
        appliedProducts: ProductSubmissionEntity.Fields,
        excludedProductsListId: 'excludedProductsListId',
      excludedProducts$: 'excludedProducts',
        excludedProducts: ProductSubmissionEntity.Fields,
        appliedCategoriesListId: 'appliedCategoriesListId',
      appliedCategories$: 'appliedCategories',
        appliedCategories: CategoryEntity.Fields,
        excludedCategoriesListId: 'excludedCategoriesListId',
      excludedCategories$: 'excludedCategories',
        excludedCategories: CategoryEntity.Fields,
}
}