import { useT } from "@/modules/fireback/hooks/useT";
import { WorkspaceInviteList } from "./WorkspaceInviteList";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { WorkspaceInviteEntity } from "@/modules/fireback/sdk/modules/workspaces/WorkspaceInviteEntity";

export const WorkspaceInviteArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.workspaceInvites}
        newEntityHandler={({ locale, router }) => {
          router.push(WorkspaceInviteEntity.Navigation.create());
        }}
      >
        <WorkspaceInviteList />
      </CommonArchiveManager>
    </>
  );
};
