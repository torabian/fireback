import { useS } from "../../fireback-ui/hooks/useS";
import { strings as uiStrings } from "../../fireback-ui/components/strings/translations";
import { strings } from "./strings/translations";
import { useWorkspaceTypeBrowseActionQuery } from "../../sdk/abac/WorkspaceTypeBrowseAction";
import { useWorkspaceTypeAwareDeleteAction } from "../../sdk/abac/WorkspaceTypeAwareDeleteAction";

import { WorkspaceTypeNavigation } from "../../sdk/navigation/AbacNavigation";
import { CommonListManager } from "../../fireback-ui/components/entity-manager/CommonListManager";
import { columns } from "./WorkspaceTypeColumns";

export const WorkspaceTypeList = () => {
  const s = useS(strings);
  const uiS = useS(uiStrings);

  return (
    <>
      <CommonListManager
        columns={columns(s, uiS)}
        queryHook={useWorkspaceTypeBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          WorkspaceTypeNavigation.single(uniqueId)
        }
        deleteHook={useWorkspaceTypeAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
