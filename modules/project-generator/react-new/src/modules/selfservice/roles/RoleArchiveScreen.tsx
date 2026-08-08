import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { RoleList } from "./RoleList";
import { useCommonArchiveExportTools } from "../../fireback-ui/components/action-menu/ActionMenu";
import { CommonArchiveManager } from "../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { RoleNavigation } from "../../sdk/navigation/AbacNavigation";

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
