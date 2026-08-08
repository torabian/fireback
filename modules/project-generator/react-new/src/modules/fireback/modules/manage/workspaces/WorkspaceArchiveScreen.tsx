import { useT } from "../../../hooks/useT";

import { CommonArchiveManager } from "../../../components/entity-manager/CommonArchiveManager";
import { WorkspaceList } from "./WorkspaceList";
import { WorkspaceNavigation } from "../../../sdk/navigation/AbacNavigation";

export const WorkspaceArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.workspaces}
        newEntityHandler={({ locale, router }) => {
          router.push(WorkspaceNavigation.create());
        }}
      >
        <WorkspaceList />
      </CommonArchiveManager>
    </>
  );
};
