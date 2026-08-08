import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { useWorkspaceInviteBrowseActionQuery } from "../../../sdk/abac/WorkspaceInviteBrowseAction";
import { useWorkspaceInviteAwareDeleteAction } from "../../../sdk/abac/WorkspaceInviteAwareDeleteAction";
import { useT } from "../../../hooks/useT";
import { WorkspaceInviteNavigation } from "../../../sdk/navigation/AbacNavigation";
import { columns } from "./WorkspaceInviteColumns";

export const WorkspaceInviteList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useWorkspaceInviteBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          WorkspaceInviteNavigation.single(uniqueId)
        }
        deleteHook={useWorkspaceInviteAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
