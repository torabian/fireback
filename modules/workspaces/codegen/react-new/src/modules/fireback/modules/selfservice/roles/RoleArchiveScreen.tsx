import { useT } from "@/modules/fireback/hooks/useT";
import { RoleList } from "./RoleList";
import { useCommonArchiveExportTools } from "@/modules/fireback/components/action-menu/ActionMenu";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { RoleEntity } from "@/modules/fireback/sdk/modules/workspaces/RoleEntity";

export const RoleArchiveScreen = () => {
  const t = useT();

  useCommonArchiveExportTools();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={({ locale, router }) =>
          router.push(RoleEntity.Navigation.create())
        }
        pageTitle={t.fbMenu.roles}
      >
        <RoleList />
      </CommonArchiveManager>
    </>
  );
};
