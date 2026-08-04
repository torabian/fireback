import { useT } from "../../../hooks/useT";

import { WorkspaceTypeEntity } from "@/modules/fireback/sdk/modules/abac/WorkspaceTypeEntity";
import { usePostWorkspaceTypeRemove } from "@/modules/fireback/sdk/modules/abac/usePostWorkspaceTypeRemove";
import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { useGetWorkspaceTypes } from "../../../sdk/modules/abac/useGetWorkspaceTypes";
import { columns } from "./WorkspaceTypeColumns";

export const WorkspaceTypeList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetWorkspaceTypes}
        uniqueIdHrefHandler={(uniqueId: string) =>
          WorkspaceTypeEntity.Navigation.single(uniqueId)
        }
        deleteHook={usePostWorkspaceTypeRemove}
      ></CommonListManager>
    </>
  );
};
