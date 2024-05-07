import { useT } from "@/fireback/hooks/useT";
import { CommonListManager } from "@/fireback/components/entity-manager/CommonListManager";
import { columns } from "./ProductSubmissionColumns";
import { ProductSubmissionEntity } from "src/sdk/fireback/modules/shop/ProductSubmissionEntity";
import { useGetProductSubmissions } from "src/sdk/fireback/modules/shop/useGetProductSubmissions";
import { useDeleteProductSubmission } from "@/sdk/fireback/modules/shop/useDeleteProductSubmission";
export const ProductSubmissionList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t) as any}
        queryHook={useGetProductSubmissions}
        withPreloads="Price,Price.Variations"
        uniqueIdHrefHandler={(uniqueId: string) =>
          ProductSubmissionEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteProductSubmission}
      ></CommonListManager>
    </>
  );
};
