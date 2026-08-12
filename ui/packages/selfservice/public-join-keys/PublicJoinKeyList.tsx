import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
import { usePublicJoinKeyBrowseActionQuery } from "@fireback/selfservice/sdk/abac/PublicJoinKeyBrowseAction";
import { usePublicJoinKeyAwareDeleteAction } from "@fireback/selfservice/sdk/abac/PublicJoinKeyAwareDeleteAction";
import { useS } from "@fireback/ui-core/hooks/useS";
import { PublicJoinKeyDto } from "@fireback/selfservice/sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
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
