import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { useGetWorkspaceInvites } from "@/modules/fireback/sdk/modules/abac/useGetWorkspaceInvites";
import { usePostWorkspaceInviteRemove } from "@/modules/fireback/sdk/modules/abac/usePostWorkspaceInviteRemove";
import { WorkspaceInviteEntity } from "@/modules/fireback/sdk/modules/abac/WorkspaceInviteEntity";
import { columns } from "./WorkspaceInviteColumns";

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
        deleteHook={usePostWorkspaceInviteRemove}
      ></CommonListManager>
    </>
  );
};
