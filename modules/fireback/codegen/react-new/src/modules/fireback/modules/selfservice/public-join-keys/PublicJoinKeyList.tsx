import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { useDeletePublicJoinKey } from "@/modules/fireback/sdk/modules/abac/useDeletePublicJoinKey";
import { useGetPublicJoinKeys } from "@/modules/fireback/sdk/modules/abac/useGetPublicJoinKeys";
import { columns } from "./PublicJoinKeyColumns";
import { PublicJoinKeyEntity } from "@/modules/fireback/sdk/modules/abac/PublicJoinKeyEntity";

export const PublicJoinKeyList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetPublicJoinKeys}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PublicJoinKeyEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeletePublicJoinKey}
      ></CommonListManager>
    </>
  );
};
