import { useS } from "../../fireback-ui/hooks/useS";
import { strings as uiStrings } from "../../fireback-ui/components/strings/translations";
import { strings } from "./strings/translations";
import { useWorkspaceBrowseActionQuery } from "../../sdk/abac/WorkspaceBrowseAction";
import { useWorkspaceAwareDeleteAction } from "../../sdk/abac/WorkspaceAwareDeleteAction";

import { CommonRowDetail } from "../../fireback-ui/components/detail-table/DetailTable";
import { CommonListManager } from "../../fireback-ui/components/entity-manager/CommonListManager";
import { columns } from "./WorkspaceColumns";
import { WorkspaceNavigation } from "../../sdk/navigation/AbacNavigation";

export const WorkspaceList = () => {
  const s = useS(strings);
  const uiS = useS(uiStrings);
  const uniqueIdHrefHandler = (uniqueId: string) =>
    WorkspaceNavigation.single(uniqueId);

  return (
    <>
      <CommonListManager
        columns={columns(s, uiS)}
        queryHook={useWorkspaceBrowseActionQuery}
        deleteHook={useWorkspaceAwareDeleteAction}
        onRecordsDeleted={({ queryClient }) => {
          queryClient.invalidateQueries("*fireback.UserRoleWorkspace");
          queryClient.invalidateQueries("*fireback.WorkspaceEntity");
        }}
        RowDetail={(props: any) => (
          <CommonRowDetail
            {...props}
            columns={() => columns(s, uiS)}
            uniqueIdHref
            Handler={uniqueIdHrefHandler}
          />
        )}
        uniqueIdHrefHandler={uniqueIdHrefHandler}
      ></CommonListManager>
    </>
  );
};
