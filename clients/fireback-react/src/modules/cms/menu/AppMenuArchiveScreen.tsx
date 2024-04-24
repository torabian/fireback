import { useT } from "@/hooks/useT";
import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { AppMenuList } from "./AppMenuList";
import { AppMenuEntity } from "src/sdk/fireback/modules/workspaces/AppMenuEntity";
export const AppMenuArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.appMenus.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(AppMenuEntity.Navigation.create(locale));
      }}
    >
      <AppMenuList />
    </CommonArchiveManager>
  );
};
