import { useT } from "@/hooks/useT";
import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { columns } from "./ShoppingCartColumns";
import { ShoppingCartEntity } from "src/sdk/fireback/modules/shop/ShoppingCartEntity";
import { useGetShoppingCarts } from "src/sdk/fireback/modules/shop/useGetShoppingCarts";
import { useDeleteShoppingCart } from "@/sdk/fireback/modules/shop/useDeleteShoppingCart";
export const ShoppingCartList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetShoppingCarts}
        uniqueIdHrefHandler={(uniqueId: string) =>
          ShoppingCartEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteShoppingCart}
      ></CommonListManager>
    </>
  );
};