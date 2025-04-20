import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { columns } from "./CapabilityColumns";
import { CapabilityEntity } from "@/modules/fireback/sdk/modules/fireback/CapabilityEntity";
import { useDeleteCapability } from "@/modules/fireback/sdk/modules/fireback/useDeleteCapability";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { useGetCapabilities } from "@/modules/fireback/sdk/modules/fireback/useGetCapabilities";
export const CapabilityList = () => {
  const s = useS(strings);
  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={useGetCapabilities}
        uniqueIdHrefHandler={(uniqueId: string) =>
          CapabilityEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteCapability}
      ></CommonListManager>
    </>
  );
};
