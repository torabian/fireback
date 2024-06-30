import { useT } from "@/fireback/hooks/useT";
import { CommonArchiveManager } from "@/fireback/components/entity-manager/CommonArchiveManager";
import { PageCategoryList } from "./PageCategoryList";
import { PageCategoryEntity } from "src/sdk/fireback/modules/cms/PageCategoryEntity";
export const PageCategoryArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.pagecategories.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(PageCategoryEntity.Navigation.create(locale));
      }}
    >
      <PageCategoryList />
    </CommonArchiveManager>
  );
};
