import { useT } from "@/hooks/useT";
import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { columns } from "./PostCategoryColumns";
import { PostCategoryEntity } from "src/sdk/fireback/modules/cms/PostCategoryEntity";
import { useDeletePostCategory } from "@/sdk/fireback/modules/cms/useDeletePostCategory";
import { useGetPostCategories } from "@/sdk/fireback/modules/cms/useGetPostCategories";
export const PostCategoryList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetPostCategories}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PostCategoryEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeletePostCategory}
      ></CommonListManager>
    </>
  );
};
