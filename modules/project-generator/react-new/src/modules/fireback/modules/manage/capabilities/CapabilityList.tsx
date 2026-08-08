import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { useCapabilityAwareDeleteAction } from "@/modules/fireback/sdk/abac/CapabilityAwareDeleteAction";
import { useCapabilityBrowseActionQuery } from "@/modules/fireback/sdk/abac/CapabilityBrowseAction";
import { CapabilityDto } from "@/modules/fireback/sdk/abac/CapabilityDto";
import { CapabilityNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
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
