import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { useCapabilityAwareDeleteAction } from "@fireback/manage/sdk/abac/CapabilityAwareDeleteAction";
import { useCapabilityBrowseActionQuery } from "@fireback/manage/sdk/abac/CapabilityBrowseAction";
import { CapabilityNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { columns } from "./CapabilityColumns";
import { strings } from "./strings/translations";
import { createUdfBrowseQueryHook } from "@fireback/ui-core/hooks/useUdfBrowseQuery";

export const CapabilityList = () => {
  const s = useS(strings);

  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={createUdfBrowseQueryHook(useCapabilityBrowseActionQuery)}
        uniqueIdHrefHandler={(uniqueId: string) =>
          CapabilityNavigation.single(uniqueId)
        }
        deleteHook={useCapabilityAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
