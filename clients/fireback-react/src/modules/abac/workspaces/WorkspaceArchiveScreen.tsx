import { useT } from "@/hooks/useT";

import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { WorkspaceList } from "./WorkspaceList";
import { WorkspaceEntity } from "@/sdk/fireback/modules/workspaces/WorkspaceEntity";

export const WorkspaceArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.workspaces}
        newEntityHandler={({ locale, router }) => {
          router.push(WorkspaceEntity.Navigation.create(locale));
        }}
      >
        <WorkspaceList />
      </CommonArchiveManager>
    </>
  );
};
