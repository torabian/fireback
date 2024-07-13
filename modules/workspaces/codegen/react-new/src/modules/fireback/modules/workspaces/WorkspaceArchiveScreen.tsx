import { useT } from "@/modules/fireback/hooks/useT";

import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { WorkspaceList } from "./WorkspaceList";
import { WorkspaceEntity } from "../../sdk/modules/workspaces/WorkspaceEntity";

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
