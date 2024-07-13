import { useT } from "@/modules/fireback/hooks/useT";

import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { columns } from "./PublicJoinKeyColumns";
import { useGetPublicJoinKeys } from "../../sdk/modules/workspaces/useGetPublicJoinKeys";
import { useDeletePublicJoinKey } from "../../sdk/modules/workspaces/useDeletePublicJoinKey";

export const PublicJoinKeyList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetPublicJoinKeys}
        uniqueIdHrefHandler={(uniqueId: string) => `/publicjoinkey/${uniqueId}`}
        deleteHook={useDeletePublicJoinKey}
      ></CommonListManager>
    </>
  );
};
