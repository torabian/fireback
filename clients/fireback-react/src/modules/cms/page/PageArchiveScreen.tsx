import { useT } from "@/fireback/hooks/useT";
import { CommonArchiveManager } from "@/fireback/components/entity-manager/CommonArchiveManager";
import { PageList } from "./PageList";
import { PageEntity } from "src/sdk/fireback/modules/cms/PageEntity";
export const PageArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.pages.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(PageEntity.Navigation.create(locale));
      }}
    >
      <PageList />
    </CommonArchiveManager>
  );
};
