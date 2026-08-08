import { useT } from "@/modules/fireback/hooks/useT";
import { WorkspaceInviteList } from "./WorkspaceInviteList";
import { CommonArchiveManager } from "@/modules/fireback/components/entity-manager/CommonArchiveManager";
import { WorkspaceInviteNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";

export const WorkspaceInviteArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.workspaceInvites}
        newEntityHandler={({ locale, router }) => {
          router.push(WorkspaceInviteNavigation.create());
        }}
      >
        <WorkspaceInviteList />
      </CommonArchiveManager>
    </>
  );
};
