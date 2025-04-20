import { useT } from "@/modules/fireback/hooks/useT";
import { columns } from "./WorkspaceInviteColumns";
import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useGetWorkspaceInvites } from "@/modules/fireback/sdk/modules/abac/useGetWorkspaceInvites";
import { WorkspaceInviteEntity } from "@/modules/fireback/sdk/modules/abac/WorkspaceInviteEntity";
import { useDeleteWorkspaceInvite } from "@/modules/fireback/sdk/modules/abac/useDeleteWorkspaceInvite";

export const WorkspaceInviteList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetWorkspaceInvites}
        uniqueIdHrefHandler={(uniqueId: string) =>
          WorkspaceInviteEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteWorkspaceInvite}
      ></CommonListManager>
    </>
  );
};
