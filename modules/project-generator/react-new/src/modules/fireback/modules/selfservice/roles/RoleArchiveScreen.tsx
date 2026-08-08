import { useT } from "../../../../fireback-ui/hooks/useT";
import { RoleList } from "./RoleList";
import { useCommonArchiveExportTools } from "../../../../fireback-ui/components/action-menu/ActionMenu";
import { CommonArchiveManager } from "../../../../fireback-ui/components/entity-manager/CommonArchiveManager";
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
