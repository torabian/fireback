import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";

import { CommonArchiveManager } from "../../../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { WorkspaceList } from "./WorkspaceList";
import { WorkspaceNavigation } from "../../../sdk/navigation/AbacNavigation";

export const WorkspaceArchiveScreen = () => {
  const s = useS(strings);

  return (
    <>
      <CommonArchiveManager
        pageTitle={s.menuTitle}
        newEntityHandler={({ locale, router }) => {
          router.push(WorkspaceNavigation.create());
        }}
      >
        <WorkspaceList />
      </CommonArchiveManager>
    </>
  );
};
