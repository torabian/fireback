import { useT } from "../../hooks/useT";

import { CommonListManager } from "../../components/entity-manager/CommonListManager";
import { useDeleteWorkspaceType } from "../../sdk/modules/workspaces/useDeleteWorkspaceType";
import { useGetWorkspaceTypes } from "../../sdk/modules/workspaces/useGetWorkspaceTypes";
import { columns } from "./WorkspaceTypeColumns";

export const WorkspaceTypeList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetWorkspaceTypes}
        uniqueIdHrefHandler={(uniqueId: string) =>
          `/workspace-type/${uniqueId}`
        }
        deleteHook={useDeleteWorkspaceType}
      ></CommonListManager>
    </>
  );
};
