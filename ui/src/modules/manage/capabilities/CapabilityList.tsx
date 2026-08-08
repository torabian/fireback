import { CommonListManager } from "../../fireback-ui/components/entity-manager/CommonListManager";
import { useS } from "../../fireback-ui/hooks/useS";
import { useCapabilityAwareDeleteAction } from "../../sdk/abac/CapabilityAwareDeleteAction";
import { useCapabilityBrowseActionQuery } from "../../sdk/abac/CapabilityBrowseAction";
import { CapabilityNavigation } from "../../sdk/navigation/AbacNavigation";
import { columns } from "./CapabilityColumns";
import { strings } from "./strings/translations";

export const CapabilityList = () => {
  const s = useS(strings);

  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={useCapabilityBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          CapabilityNavigation.single(uniqueId)
        }
        deleteHook={useCapabilityAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
