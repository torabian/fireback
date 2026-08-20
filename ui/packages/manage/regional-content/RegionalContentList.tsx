import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
import { useRegionalContentBrowseActionQuery } from "@fireback/manage/sdk/abac/RegionalContentBrowseAction";
import { useRegionalContentAwareDeleteAction } from "@fireback/manage/sdk/abac/RegionalContentAwareDeleteAction";
import { useS } from "@fireback/ui-core/hooks/useS";
import { RegionalContentDto } from "@fireback/manage/sdk/abac/RegionalContentDto";
import { RegionalContentNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { columns } from "./RegionalContentColumns";
import { strings } from "./strings/translations";
import { createUdfBrowseQueryHook } from "@fireback/ui-core/hooks/useUdfBrowseQuery";
export const RegionalContentList = () => {
  const s = useS(strings);
  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={createUdfBrowseQueryHook(
          useRegionalContentBrowseActionQuery,
        )}
        uniqueIdHrefHandler={(uniqueId: string) =>
          RegionalContentNavigation.single(uniqueId)
        }
        deleteHook={useRegionalContentAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
