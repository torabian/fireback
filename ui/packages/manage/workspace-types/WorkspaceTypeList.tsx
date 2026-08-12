import { useS } from "@fireback/ui-core/hooks/useS";
import { strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { strings } from "./strings/translations";
import { useWorkspaceTypeBrowseActionQuery } from "@fireback/manage/sdk/abac/WorkspaceTypeBrowseAction";
import { useWorkspaceTypeAwareDeleteAction } from "@fireback/manage/sdk/abac/WorkspaceTypeAwareDeleteAction";

import { WorkspaceTypeNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
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
