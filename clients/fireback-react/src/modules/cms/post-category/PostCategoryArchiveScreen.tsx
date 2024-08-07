import { useT } from "@/fireback/hooks/useT";
import { CommonArchiveManager } from "@/fireback/components/entity-manager/CommonArchiveManager";
import { PostCategoryList } from "./PostCategoryList";
import { PostCategoryEntity } from "src/sdk/fireback/modules/cms/PostCategoryEntity";
export const PostCategoryArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.postcategories.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(PostCategoryEntity.Navigation.create(locale));
      }}
    >
      <PostCategoryList />
    </CommonArchiveManager>
  );
};
