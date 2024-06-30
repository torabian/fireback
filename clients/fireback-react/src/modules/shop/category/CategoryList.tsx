import { useT } from "@/fireback/hooks/useT";
import { CommonListManager } from "@/fireback/components/entity-manager/CommonListManager";
import { columns } from "./CategoryColumns";
import { CategoryEntity } from "src/sdk/fireback/modules/shop/CategoryEntity";
import { useGetCategories } from "src/sdk/fireback/modules/shop/useGetCategories";
import { useDeleteCategory } from "@/sdk/fireback/modules/shop/useDeleteCategory";
export const CategoryList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetCategories}
        uniqueIdHrefHandler={(uniqueId: string) =>
          CategoryEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteCategory}
      ></CommonListManager>
    </>
  );
};
