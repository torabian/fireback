import { useT } from "@/fireback/hooks/useT";

import { CommonArchiveManager } from "@/fireback/components/entity-manager/CommonArchiveManager";
import { WorkspaceInviteList } from "./WorkspaceInviteList";

export const WorkspaceInviteArchiveScreen = () => {
  const t = useT();

  return (
    <>
      <CommonArchiveManager
        pageTitle={t.fbMenu.workspaceInvites}
        newEntityHandler={({ locale, router }) => {
          router.push(`/${locale}/workspace/invite/new`);
        }}
      >
        <WorkspaceInviteList />
      </CommonArchiveManager>
    </>
  );
};
