import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { WorkspaceInviteList } from "./WorkspaceInviteList";
import { CommonArchiveManager } from "@fireback/ui-core/components/entity-manager/CommonArchiveManager";
import { WorkspaceInviteNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";

export const WorkspaceInviteArchiveScreen = () => {
  const s = useS(strings);

  return (
    <>
      <CommonArchiveManager
        pageTitle={s.menuTitle}
        newEntityHandler={({ locale, router }) => {
          router.push(WorkspaceInviteNavigation.create());
        }}
      >
        <WorkspaceInviteList />
      </CommonArchiveManager>
    </>
  );
};
