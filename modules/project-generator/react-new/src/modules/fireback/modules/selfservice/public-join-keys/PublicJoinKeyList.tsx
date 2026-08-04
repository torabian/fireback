import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { PublicJoinKeyEntity } from "@/modules/fireback/sdk/modules/abac/PublicJoinKeyEntity";
import { useGetPublicJoinKeys } from "@/modules/fireback/sdk/modules/abac/useGetPublicJoinKeys";
import { usePostPublicJoinKeyRemove } from "@/modules/fireback/sdk/modules/abac/usePostPublicJoinKeyRemove";
import { columns } from "./PublicJoinKeyColumns";
import { strings } from "./strings/translations";

export const PublicJoinKeyList = () => {
  const s = useS(strings);

  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={useGetPublicJoinKeys}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PublicJoinKeyEntity.Navigation.single(uniqueId)
        }
        deleteHook={usePostPublicJoinKeyRemove}
      ></CommonListManager>
    </>
  );
};
