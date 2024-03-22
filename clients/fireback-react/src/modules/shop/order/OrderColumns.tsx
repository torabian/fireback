import { enTranslations } from "@/translations/en";
import { OrderEntity } from "src/sdk/fireback/modules/shop/OrderEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: OrderEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: OrderEntity.Fields.totalPrice$,
    title: t.orders.totalPrice,
    width: 100,
    getCellValue: (entity: OrderEntity) =>
      entity.totalPrice?.amount + " " + entity.totalPrice?.currencyId,
  },
  {
    name: OrderEntity.Fields.shippingAddress,
    title: t.orders.shippingAddress,
    width: 100,
  },
  {
    name: OrderEntity.Fields.paymentStatus$,
    title: t.orders.paymentStatus,
    getCellValue: (entity: OrderEntity) =>
      entity.paymentStatus?.name || entity.paymentStatusId,
    width: 100,
  },
  {
    name: OrderEntity.Fields.orderStatus$,
    title: t.orders.orderStatus,
    getCellValue: (entity: OrderEntity) =>
      entity.orderStatus?.name || entity.orderStatusId,
    width: 100,
  },
  {
    name: OrderEntity.Fields.invoiceNumber,
    title: t.orders.invoiceNumber,
    width: 100,
  },
  {
    name: OrderEntity.Fields.discountCode,
    title: t.orders.discountCode,
    width: 100,
  },
  {
    name: OrderEntity.Fields.items$,
    title: t.orders.items,
    width: 100,
    getCellValue: (entity: OrderEntity) => entity.items.length,
  },
];
