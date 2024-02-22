import { useT } from "@/hooks/useT";

import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { WorkspaceNavigationTools } from "src/sdk/fireback/modules/workspaces/workspace-navigation-tools";
import { WorkspaceList } from "./WorkspaceList";

export const WorkspaceArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.workspaces}
        newEntityHandler={({ locale, router }) => {
          router.push(WorkspaceNavigationTools.create(locale));
        }}
      >
        <WorkspaceList />
      </CommonArchiveManager>
    </>
  );
};
