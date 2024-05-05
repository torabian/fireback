import { useT } from "@/fireback/hooks/useT";

import { CommonListManager } from "@/fireback/components/entity-manager/CommonListManager";
import { useDeleteWorkspaceType } from "src/sdk/fireback/modules/workspaces/useDeleteWorkspaceType";
import { useGetWorkspaceTypes } from "src/sdk/fireback/modules/workspaces/useGetWorkspaceTypes";
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
