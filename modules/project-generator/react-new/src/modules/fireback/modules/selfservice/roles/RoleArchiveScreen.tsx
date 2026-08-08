import { useT } from "../../../hooks/useT";
import { RoleList } from "./RoleList";
import { useCommonArchiveExportTools } from "../../../components/action-menu/ActionMenu";
import { CommonArchiveManager } from "../../../components/entity-manager/CommonArchiveManager";
import { RoleNavigation } from "../../../sdk/navigation/AbacNavigation";

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
