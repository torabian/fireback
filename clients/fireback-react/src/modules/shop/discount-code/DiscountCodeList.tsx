import { useT } from "@/hooks/useT";
import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { columns } from "./DiscountCodeColumns";
import { DiscountCodeEntity } from "src/sdk/fireback/modules/shop/DiscountCodeEntity";
import { useGetDiscountCodes } from "src/sdk/fireback/modules/shop/useGetDiscountCodes";
import { useDeleteDiscountCode } from "@/sdk/fireback/modules/shop/useDeleteDiscountCode";
export const DiscountCodeList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetDiscountCodes}
        uniqueIdHrefHandler={(uniqueId: string) =>
          DiscountCodeEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteDiscountCode}
      ></CommonListManager>
    </>
  );
};