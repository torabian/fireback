import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { useCapabilityBrowseActionQuery } from "@/modules/fireback/sdk/fireback/CapabilityBrowseAction";
import { CapabilityEntity } from "@/modules/fireback/sdk/modules/fireback/CapabilityEntity";
import { usePostCapabilityRemove } from "@/modules/fireback/sdk/modules/fireback/usePostCapabilityRemove";
import { columns } from "./CapabilityColumns";
import { strings } from "./strings/translations";
import { useCapabilityAwareDeleteAction } from "@/modules/fireback/sdk/fireback/CapabilityAwareDeleteAction";
export const CapabilityList = () => {
  const s = useS(strings);

  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={useCapabilityBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          CapabilityEntity.Navigation.single(uniqueId)
        }
        deleteHook={useCapabilityAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
