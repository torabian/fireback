import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { WorkspaceInviteList } from "./WorkspaceInviteList";
import { CommonArchiveManager } from "../../fireback-ui/components/entity-manager/CommonArchiveManager";
import { WorkspaceInviteNavigation } from "../../sdk/navigation/AbacNavigation";

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
