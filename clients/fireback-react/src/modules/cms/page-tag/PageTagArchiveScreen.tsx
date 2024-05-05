import { useT } from "@/fireback/hooks/useT";
import { CommonArchiveManager } from "@/fireback/components/entity-manager/CommonArchiveManager";
import { PageTagList } from "./PageTagList";
import { PageTagEntity } from "src/sdk/fireback/modules/cms/PageTagEntity";
export const PageTagArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.pagetags.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(PageTagEntity.Navigation.create(locale));
      }}
    >
      <PageTagList />
    </CommonArchiveManager>
  );
};
