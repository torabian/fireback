import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
import { useWorkspaceInviteBrowseActionQuery } from "@fireback/selfservice/sdk/abac/WorkspaceInviteBrowseAction";
import { useWorkspaceInviteAwareDeleteAction } from "@fireback/selfservice/sdk/abac/WorkspaceInviteAwareDeleteAction";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { strings } from "./strings/translations";
import { WorkspaceInviteNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { columns } from "./WorkspaceInviteColumns";
import { createUdfBrowseQueryHook } from "@fireback/ui-core/hooks/useUdfBrowseQuery";

export const WorkspaceInviteList = () => {
  const s = useS(strings);
  const uiS = useS(uiStrings);

  return (
    <>
      <CommonListManager
        columns={columns(s, uiS)}
        queryHook={createUdfBrowseQueryHook(
          useWorkspaceInviteBrowseActionQuery,
        )}
        uniqueIdHrefHandler={(uniqueId: string) =>
          WorkspaceInviteNavigation.single(uniqueId)
        }
        deleteHook={useWorkspaceInviteAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
