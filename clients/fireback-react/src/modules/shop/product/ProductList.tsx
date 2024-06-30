import { useT } from "@/fireback/hooks/useT";
import { CommonListManager } from "@/fireback/components/entity-manager/CommonListManager";
import { columns } from "./ProductColumns";
import { ProductEntity } from "src/sdk/fireback/modules/shop/ProductEntity";
import { useGetProducts } from "src/sdk/fireback/modules/shop/useGetProducts";
import { useDeleteProduct } from "@/sdk/fireback/modules/shop/useDeleteProduct";
export const ProductList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t) as any}
        queryHook={useGetProducts}
        uniqueIdHrefHandler={(uniqueId: string) =>
          ProductEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteProduct}
      ></CommonListManager>
    </>
  );
};
