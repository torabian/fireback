import { useCommonArchiveExportTools } from "@/components/action-menu/ActionMenu";
import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { useT } from "@/hooks/useT";
import { RoleEntity } from "@/sdk/fireback/modules/workspaces/RoleEntity";
import { RoleList } from "./RoleList";

export const RoleArchiveScreen = () => {
  const t = useT();

  useCommonArchiveExportTools();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={({ locale, router }) =>
          router.push(RoleEntity.Navigation.create(locale))
        }
        pageTitle={t.fbMenu.roles}
      >
        <RoleList />
      </CommonArchiveManager>
    </>
  );
};
