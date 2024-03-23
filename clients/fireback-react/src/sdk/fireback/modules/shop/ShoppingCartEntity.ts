import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    ProductSubmissionEntity,
} from "./ProductSubmissionEntity"
// In this section we have sub entities related to this object
export class ShoppingCartItems extends BaseEntity {
  public quantity?: number | null;
  public product?: ProductSubmissionEntity | null;
      productId?: string | null;
}
// Class body
export type ShoppingCartEntityKeys =
  keyof typeof ShoppingCartEntity.Fields;
export class ShoppingCartEntity extends BaseEntity {
  public children?: ShoppingCartEntity[] | null;
  public items?: ShoppingCartItems[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/shopping-cart/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/shopping-cart/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/shopping-cart/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/shopping-carts`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "shopping-cart/edit/:uniqueId",
      Rcreate: "shopping-cart/new",
      Rsingle: "shopping-cart/:uniqueId",
      Rquery: "shopping-carts",
      rItemsCreate: "shopping-cart/:linkerId/items/new",
      rItemsEdit: "shopping-cart/:linkerId/items/edit/:uniqueId",
      editItems(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/shopping-cart/${linkerId}/items/edit/${uniqueId}`;
      },
      createItems(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/shopping-cart/${linkerId}/items/new`;
      },
  };
  public static definition = {
  "name": "shoppingCart",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "linkedTo": "ShoppingCartEntity",
      "name": "items",
      "type": "array",
      "computedType": "ShoppingCartItems[]",
      "gormMap": {},
      "fullName": "ShoppingCartItems",
      "fields": [
        {
          "name": "quantity",
          "type": "float64",
          "computedType": "number",
          "gormMap": {}
        },
        {
          "name": "product",
          "type": "one",
          "target": "ProductSubmissionEntity",
          "computedType": "ProductSubmissionEntity",
          "gormMap": {}
        }
      ]
    }
  ],
  "cliDescription": "Manage the active shopping carts (not ordered yet of the store)"
}
public static Fields = {
  ...BaseEntity.Fields,
      items$: 'items',
      items: {
  ...BaseEntity.Fields,
      quantity: 'quantity',
          productId: 'productId',
      product$: 'product',
        product: ProductSubmissionEntity.Fields,
      },
}
}