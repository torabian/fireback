import { useT } from "@/hooks/useT";
import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { OrderList } from "./OrderList";
import { OrderEntity } from "src/sdk/fireback/modules/shop/OrderEntity";
export const OrderArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.orders.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(OrderEntity.Navigation.create(locale));
      }}
    >
      <OrderList />
    </CommonArchiveManager>
  );
};