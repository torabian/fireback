import { useCommonArchiveExportTools } from "@/components/action-menu/ActionMenu";
import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { useT } from "@/hooks/useT";
import { RoleNavigationTools } from "src/sdk/fireback/modules/workspaces/role-navigation-tools";
import { RoleList } from "./RoleList";

export const RoleArchiveScreen = () => {
  const t = useT();

  useCommonArchiveExportTools();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={({ locale, router }) =>
          router.push(RoleNavigationTools.create(locale))
        }
        pageTitle={t.fbMenu.roles}
      >
        <RoleList />
      </CommonArchiveManager>
    </>
  );
};
