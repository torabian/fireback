import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { useS } from "../../../hooks/useS";
import { useCapabilityAwareDeleteAction } from "../../../sdk/abac/CapabilityAwareDeleteAction";
import { useCapabilityBrowseActionQuery } from "../../../sdk/abac/CapabilityBrowseAction";
import { CapabilityDto } from "../../../sdk/abac/CapabilityDto";
import { CapabilityNavigation } from "../../../sdk/navigation/AbacNavigation";
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
