import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { RoleList } from "./RoleList";
import { useCommonArchiveExportTools } from "@fireback/ui-core/components/action-menu/ActionMenu";
import { CommonArchiveManager } from "@fireback/ui-core/components/entity-manager/CommonArchiveManager";
import { RoleNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";

export const RoleArchiveScreen = () => {
  const s = useS(strings);

  useCommonArchiveExportTools();

  return (
    <>
      <CommonArchiveManager
        newEntityHandler={({ locale, router }) =>
          router.push(RoleNavigation.create())
        }
        pageTitle={s.menuTitle}
      >
        <RoleList />
      </CommonArchiveManager>
    </>
  );
};
