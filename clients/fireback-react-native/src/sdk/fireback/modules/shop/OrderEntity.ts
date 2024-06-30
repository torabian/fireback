import {
    BaseDto,
    BaseEntity,
} from "../../core/definitions"
import {
    CurrencyEntity,
} from "../currency/CurrencyEntity"
import {
    DiscountCodeEntity,
} from "./DiscountCodeEntity"
import {
    OrderStatusEntity,
} from "./OrderStatusEntity"
import {
    PaymentStatusEntity,
} from "./PaymentStatusEntity"
import {
    ProductSubmissionEntity,
} from "./ProductSubmissionEntity"
// In this section we have sub entities related to this object
export class OrderTotalPrice extends BaseEntity {
  public amount?: number | null;
  public currency?: CurrencyEntity | null;
      currencyId?: string | null;
}
export class OrderItems extends BaseEntity {
  public quantity?: number | null;
  public price?: number | null;
  public product?: ProductSubmissionEntity | null;
      productId?: string | null;
  public productSnapshot?: any | null;
}
// Class body
export type OrderEntityKeys =
  keyof typeof OrderEntity.Fields;
export class OrderEntity extends BaseEntity {
  public children?: OrderEntity[] | null;
  public totalPrice?: OrderTotalPrice | null;
  public shippingAddress?: string | null;
  public paymentStatus?: PaymentStatusEntity | null;
      paymentStatusId?: string | null;
  public orderStatus?: OrderStatusEntity | null;
      orderStatusId?: string | null;
  public invoiceNumber?: string | null;
  public discountCode?: DiscountCodeEntity | null;
      discountCodeId?: string | null;
  public items?: OrderItems[] | null;
  public static Navigation = {
      edit(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/order/edit/${uniqueId}`;
      },
      create(locale?: string) {
          return `${locale ? '/' + locale : ''}/order/new`;
      },
      single(uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/order/${uniqueId}`;
      },
      query(params: any = {}, locale?: string) {
          return `${locale ? '/' + locale : ''}/orders`;
      },
      /**
      * Use R series while building router in CRA or nextjs, or react navigation for react Native
      * Might be useful in Angular as well.
      **/
      Redit: "order/edit/:uniqueId",
      Rcreate: "order/new",
      Rsingle: "order/:uniqueId",
      Rquery: "orders",
      rTotalPriceCreate: "order/:linkerId/total_price/new",
      rTotalPriceEdit: "order/:linkerId/total_price/edit/:uniqueId",
      editTotalPrice(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/order/${linkerId}/total_price/edit/${uniqueId}`;
      },
      createTotalPrice(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/order/${linkerId}/total_price/new`;
      },
      rItemsCreate: "order/:linkerId/items/new",
      rItemsEdit: "order/:linkerId/items/edit/:uniqueId",
      editItems(linkerId: string, uniqueId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/order/${linkerId}/items/edit/${uniqueId}`;
      },
      createItems(linkerId: string, locale?: string) {
          return `${locale ? '/' + locale : ''}/order/${linkerId}/items/new`;
      },
  };
  public static definition = {
  "permissions": [
    {
      "name": "Confirm Order",
      "key": "confirm",
      "description": "Allows a person in the workspace to confirm an order"
    }
  ],
  "name": "order",
  "http": {},
  "gormMap": {},
  "fields": [
    {
      "linkedTo": "OrderEntity",
      "name": "totalPrice",
      "type": "object",
      "computedType": "OrderTotalPrice",
      "gormMap": {},
      "fullName": "OrderTotalPrice",
      "fields": [
        {
          "name": "amount",
          "type": "float64",
          "validate": "required",
          "computedType": "number",
          "gormMap": {}
        },
        {
          "name": "currency",
          "type": "one",
          "target": "CurrencyEntity",
          "module": "currency",
          "computedType": "CurrencyEntity",
          "gormMap": {}
        }
      ]
    },
    {
      "description": "Final computed shipping address which will be print on the product",
      "name": "shippingAddress",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "paymentStatus",
      "type": "one",
      "target": "PaymentStatusEntity",
      "validate": "required",
      "computedType": "PaymentStatusEntity",
      "gormMap": {}
    },
    {
      "name": "orderStatus",
      "type": "one",
      "target": "OrderStatusEntity",
      "validate": "required",
      "computedType": "OrderStatusEntity",
      "gormMap": {}
    },
    {
      "name": "invoiceNumber",
      "type": "string",
      "computedType": "string",
      "gormMap": {}
    },
    {
      "name": "discountCode",
      "type": "one",
      "target": "DiscountCodeEntity",
      "computedType": "DiscountCodeEntity",
      "gormMap": {}
    },
    {
      "linkedTo": "OrderEntity",
      "name": "items",
      "type": "array",
      "computedType": "OrderItems[]",
      "gormMap": {},
      "fullName": "OrderItems",
      "fields": [
        {
          "name": "quantity",
          "type": "float64",
          "computedType": "number",
          "gormMap": {}
        },
        {
          "name": "price",
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
        },
        {
          "name": "productSnapshot",
          "type": "json",
          "computedType": "any",
          "gormMap": {}
        }
      ]
    }
  ],
  "cliDescription": "Placed orders by users, history, and their status"
}
public static Fields = {
  ...BaseEntity.Fields,
      totalPrice$: 'totalPrice',
      totalPrice: {
  ...BaseEntity.Fields,
      amount: 'amount',
          currencyId: 'currencyId',
      currency$: 'currency',
        currency: CurrencyEntity.Fields,
      },
      shippingAddress: 'shippingAddress',
          paymentStatusId: 'paymentStatusId',
      paymentStatus$: 'paymentStatus',
        paymentStatus: PaymentStatusEntity.Fields,
          orderStatusId: 'orderStatusId',
      orderStatus$: 'orderStatus',
        orderStatus: OrderStatusEntity.Fields,
      invoiceNumber: 'invoiceNumber',
          discountCodeId: 'discountCodeId',
      discountCode$: 'discountCode',
        discountCode: DiscountCodeEntity.Fields,
      items$: 'items',
      items: {
  ...BaseEntity.Fields,
      quantity: 'quantity',
      price: 'price',
          productId: 'productId',
      product$: 'product',
        product: ProductSubmissionEntity.Fields,
      productSnapshot: 'productSnapshot',
      },
}
}