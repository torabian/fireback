import { useT } from "../../../hooks/useT";
import { useWorkspaceTypeBrowseActionQuery } from "@/modules/fireback/sdk/abac/WorkspaceTypeBrowseAction";
import { useWorkspaceTypeAwareDeleteAction } from "@/modules/fireback/sdk/abac/WorkspaceTypeAwareDeleteAction";

import { WorkspaceTypeNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
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
