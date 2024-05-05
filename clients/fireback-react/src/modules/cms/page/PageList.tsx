import { useT } from "@/fireback/hooks/useT";
import { CommonListManager } from "@/fireback/components/entity-manager/CommonListManager";
import { columns } from "./PageColumns";
import { PageEntity } from "src/sdk/fireback/modules/cms/PageEntity";
import { useGetPages } from "src/sdk/fireback/modules/cms/useGetPages";
import { useDeletePage } from "@/sdk/fireback/modules/cms/useDeletePage";
export const PageList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t) as any}
        queryHook={useGetPages}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PageEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeletePage}
      ></CommonListManager>
    </>
  );
};
