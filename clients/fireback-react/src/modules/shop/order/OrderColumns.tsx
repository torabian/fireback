import { enTranslations } from "@/translations/en";
import { OrderEntity } from "src/sdk/fireback/modules/shop/OrderEntity";
export const columns = (t: typeof enTranslations) => [
  {
    name: OrderEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: OrderEntity.Fields.totalPrice,
    title: t.orders.totalPrice,
    width: 100,
  },    
  {
    name: OrderEntity.Fields.shippingAddress,
    title: t.orders.shippingAddress,
    width: 100,
  },    
  {
    name: OrderEntity.Fields.paymentStatus,
    title: t.orders.paymentStatus,
    width: 100,
  },    
  {
    name: OrderEntity.Fields.orderStatus,
    title: t.orders.orderStatus,
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
    name: OrderEntity.Fields.items,
    title: t.orders.items,
    width: 100,
  },    
];