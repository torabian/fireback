import { useT } from "@/hooks/useT";
import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { columns } from "./PageCategoryColumns";
import { PageCategoryEntity } from "src/sdk/fireback/modules/cms/PageCategoryEntity";
import { useDeletePageCategory } from "@/sdk/fireback/modules/cms/useDeletePageCategory";
import { useGetPageCategories } from "@/sdk/fireback/modules/cms/useGetPageCategories";
export const PageCategoryList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetPageCategories}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PageCategoryEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeletePageCategory}
      ></CommonListManager>
    </>
  );
};
