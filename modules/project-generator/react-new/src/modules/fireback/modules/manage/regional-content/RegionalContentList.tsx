import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { useRegionalContentBrowseActionQuery } from "../../../sdk/abac/RegionalContentBrowseAction";
import { useRegionalContentAwareDeleteAction } from "../../../sdk/abac/RegionalContentAwareDeleteAction";
import { useS } from "../../../hooks/useS";
import { RegionalContentDto } from "../../../sdk/abac/RegionalContentDto";
import { RegionalContentNavigation } from "../../../sdk/navigation/AbacNavigation";
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
