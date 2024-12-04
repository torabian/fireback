import { useT } from "../../hooks/useT";

import { useGetWorkspaceInvites } from "../../sdk/modules/workspaces/useGetWorkspaceInvites";

import { CommonListManager } from "../../components/entity-manager/CommonListManager";
import { useDeleteWorkspaceInvite } from "../../sdk/modules/workspaces/useDeleteWorkspaceInvite";
import { columns } from "./WorkspaceInviteColumns";

export const WorkspaceInviteList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetWorkspaceInvites}
        uniqueIdHrefHandler={(uniqueId: string) =>
          `/workspace/invite/${uniqueId}`
        }
        deleteHook={useDeleteWorkspaceInvite}
      ></CommonListManager>
    </>
  );
};
