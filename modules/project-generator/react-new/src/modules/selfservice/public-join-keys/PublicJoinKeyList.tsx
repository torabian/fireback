import { CommonListManager } from "../../fireback-ui/components/entity-manager/CommonListManager";
import { usePublicJoinKeyBrowseActionQuery } from "../../sdk/abac/PublicJoinKeyBrowseAction";
import { usePublicJoinKeyAwareDeleteAction } from "../../sdk/abac/PublicJoinKeyAwareDeleteAction";
import { useS } from "../../fireback-ui/hooks/useS";
import { PublicJoinKeyDto } from "../../sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "../../sdk/navigation/AbacNavigation";
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
