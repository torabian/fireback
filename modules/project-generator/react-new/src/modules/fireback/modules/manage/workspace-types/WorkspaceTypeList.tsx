import { useT } from "../../../../fireback-ui/hooks/useT";
import { useWorkspaceTypeBrowseActionQuery } from "../../../sdk/abac/WorkspaceTypeBrowseAction";
import { useWorkspaceTypeAwareDeleteAction } from "../../../sdk/abac/WorkspaceTypeAwareDeleteAction";

import { WorkspaceTypeNavigation } from "../../../sdk/navigation/AbacNavigation";
import { CommonListManager } from "../../../../fireback-ui/components/entity-manager/CommonListManager";
import { columns } from "./WorkspaceTypeColumns";

export const WorkspaceTypeList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useWorkspaceTypeBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          WorkspaceTypeNavigation.single(uniqueId)
        }
        deleteHook={useWorkspaceTypeAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
