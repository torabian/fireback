import { useT } from "@/modules/fireback/hooks/useT";

import { CommonRowDetail } from "@/modules/fireback/components/detail-table/DetailTable";
import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useGetCteWorkspaces } from "../../sdk/modules/workspaces/useGetCteWorkspaces";
import { columns } from "./WorkspaceColumns";
import { WorkspaceEntity } from "../../sdk/modules/workspaces/WorkspaceEntity";

export const WorkspaceList = () => {
  const t = useT();
  const uniqueIdHrefHandler = (uniqueId: string) =>
    WorkspaceEntity.Navigation.single(uniqueId);

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetCteWorkspaces}
        onRecordsDeleted={({ queryClient }) => {
          queryClient.invalidateQueries("*workspaces.UserRoleWorkspace");
          queryClient.invalidateQueries("*workspaces.WorkspaceEntity");
        }}
        RowDetail={(props: any) => (
          <CommonRowDetail
            {...props}
            columns={columns}
            uniqueIdHref
            Handler={uniqueIdHrefHandler}
          />
        )}
        uniqueIdHrefHandler={uniqueIdHrefHandler}
      ></CommonListManager>
    </>
  );
};
