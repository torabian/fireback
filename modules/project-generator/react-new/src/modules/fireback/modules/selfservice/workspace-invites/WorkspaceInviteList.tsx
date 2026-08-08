import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useWorkspaceInviteBrowseActionQuery } from "@/modules/fireback/sdk/abac/WorkspaceInviteBrowseAction";
import { useWorkspaceInviteAwareDeleteAction } from "@/modules/fireback/sdk/abac/WorkspaceInviteAwareDeleteAction";
import { useT } from "@/modules/fireback/hooks/useT";
import { WorkspaceInviteNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
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
