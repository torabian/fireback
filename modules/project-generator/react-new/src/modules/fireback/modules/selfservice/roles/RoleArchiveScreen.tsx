import { useT } from "@/modules/fireback/hooks/useT";
import { RoleList } from "./RoleList";
import { useCommonArchiveExportTools } from "@/modules/fireback/components/action-menu/ActionMenu";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { RoleNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";

export const RoleArchiveScreen = () => {
  const t = useT();

  useCommonArchiveExportTools();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={({ locale, router }) =>
          router.push(RoleNavigation.create())
        }
        pageTitle={t.fbMenu.roles}
      >
        <RoleList />
      </CommonArchiveManager>
    </>
  );
};
