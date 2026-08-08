import { CommonListManager } from "../../../../fireback-ui/components/entity-manager/CommonListManager";
import { useWorkspaceInviteBrowseActionQuery } from "../../../sdk/abac/WorkspaceInviteBrowseAction";
import { useWorkspaceInviteAwareDeleteAction } from "../../../sdk/abac/WorkspaceInviteAwareDeleteAction";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings as uiStrings } from "../../../../fireback-ui/components/strings/translations";
import { strings } from "./strings/translations";
import { WorkspaceInviteNavigation } from "../../../sdk/navigation/AbacNavigation";
import { columns } from "./WorkspaceInviteColumns";

export const WorkspaceInviteList = () => {
  const s = useS(strings);
  const uiS = useS(uiStrings);

  return (
    <>
      <CommonListManager
        columns={columns(s, uiS)}
        queryHook={useWorkspaceInviteBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          WorkspaceInviteNavigation.single(uniqueId)
        }
        deleteHook={useWorkspaceInviteAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
