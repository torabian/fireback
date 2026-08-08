import { useT } from "../../../hooks/useT";
import { useWorkspaceBrowseActionQuery } from "../../../sdk/abac/WorkspaceBrowseAction";
import { useWorkspaceAwareDeleteAction } from "../../../sdk/abac/WorkspaceAwareDeleteAction";

import { CommonRowDetail } from "../../../components/detail-table/DetailTable";
import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { columns } from "./WorkspaceColumns";
import { WorkspaceNavigation } from "../../../sdk/navigation/AbacNavigation";

export const WorkspaceList = () => {
  const t = useT();
  const uniqueIdHrefHandler = (uniqueId: string) =>
    WorkspaceNavigation.single(uniqueId);

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useWorkspaceBrowseActionQuery}
        onRecordsDeleted={({ queryClient }) => {
          queryClient.invalidateQueries("*fireback.UserRoleWorkspace");
          queryClient.invalidateQueries("*fireback.WorkspaceEntity");
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
