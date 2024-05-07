import { useT } from "@/fireback/hooks/useT";
import { CommonListManager } from "@/fireback/components/entity-manager/CommonListManager";
import { columns } from "./BrandColumns";
import { BrandEntity } from "src/sdk/fireback/modules/shop/BrandEntity";
import { useGetBrands } from "src/sdk/fireback/modules/shop/useGetBrands";
import { useDeleteBrand } from "@/sdk/fireback/modules/shop/useDeleteBrand";
export const BrandList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetBrands}
        uniqueIdHrefHandler={(uniqueId: string) =>
          BrandEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteBrand}
      ></CommonListManager>
    </>
  );
};
