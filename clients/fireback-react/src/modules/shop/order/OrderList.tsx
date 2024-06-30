import { useT } from "@/fireback/hooks/useT";
import { CommonListManager } from "@/fireback/components/entity-manager/CommonListManager";
import { columns } from "./OrderColumns";
import { OrderEntity } from "src/sdk/fireback/modules/shop/OrderEntity";
import { useGetOrders } from "src/sdk/fireback/modules/shop/useGetOrders";
import { useDeleteOrder } from "@/sdk/fireback/modules/shop/useDeleteOrder";
export const OrderList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetOrders}
        uniqueIdHrefHandler={(uniqueId: string) =>
          OrderEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteOrder}
      ></CommonListManager>
    </>
  );
};
