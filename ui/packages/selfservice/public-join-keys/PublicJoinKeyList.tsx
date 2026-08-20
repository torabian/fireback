import { usePublicJoinKeyAwareDeleteAction } from "@fireback/selfservice/sdk/abac/PublicJoinKeyAwareDeleteAction";
import { usePublicJoinKeyBrowseActionQuery } from "@fireback/selfservice/sdk/abac/PublicJoinKeyBrowseAction";
import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { createUdfBrowseQueryHook } from "@fireback/ui-core/hooks/useUdfBrowseQuery";
import { PublicJoinKeyNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { columns } from "./PublicJoinKeyColumns";
import { strings } from "./strings/translations";

export const PublicJoinKeyList = () => {
  const s = useS(strings);

  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={createUdfBrowseQueryHook(usePublicJoinKeyBrowseActionQuery)}
        uniqueIdHrefHandler={(uniqueId: string) =>
          PublicJoinKeyNavigation.single(uniqueId)
        }
        deleteHook={usePublicJoinKeyAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
