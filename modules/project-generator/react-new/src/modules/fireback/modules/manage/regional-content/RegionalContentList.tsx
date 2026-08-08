import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useRegionalContentBrowseActionQuery } from "@/modules/fireback/sdk/abac/RegionalContentBrowseAction";
import { useRegionalContentAwareDeleteAction } from "@/modules/fireback/sdk/abac/RegionalContentAwareDeleteAction";
import { useS } from "@/modules/fireback/hooks/useS";
import { RegionalContentDto } from "@/modules/fireback/sdk/abac/RegionalContentDto";
import { RegionalContentNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
import { columns } from "./RegionalContentColumns";
import { strings } from "./strings/translations";
export const RegionalContentList = () => {
  const s = useS(strings);
  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={useRegionalContentBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          RegionalContentNavigation.single(uniqueId)
        }
        deleteHook={useRegionalContentAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
