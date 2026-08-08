import { useT } from "../../../../fireback-ui/hooks/useT";
import { WorkspaceInviteList } from "./WorkspaceInviteList";
import { CommonArchiveManager } from "../../../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { WorkspaceInviteNavigation } from "../../../sdk/navigation/AbacNavigation";

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
