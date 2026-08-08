import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { usePublicJoinKeyBrowseActionQuery } from "@/modules/fireback/sdk/abac/PublicJoinKeyBrowseAction";
import { usePublicJoinKeyAwareDeleteAction } from "@/modules/fireback/sdk/abac/PublicJoinKeyAwareDeleteAction";
import { useS } from "@/modules/fireback/hooks/useS";
import { PublicJoinKeyDto } from "@/modules/fireback/sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
import { columns } from "./PublicJoinKeyColumns";
import { strings } from "./strings/translations";

export const PublicJoinKeyList = () => {
  const s = useS(strings);

  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={usePublicJoinKeyBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PublicJoinKeyNavigation.single(uniqueId)
        }
        deleteHook={usePublicJoinKeyAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
