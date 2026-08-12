import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { strings } from "./strings/translations";
import { useRoleAwareDeleteAction } from "@fireback/selfservice/sdk/abac/RoleAwareDeleteAction";
import { useRoleBrowseActionQuery } from "@fireback/selfservice/sdk/abac/RoleBrowseAction";
import { RoleNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { columns } from "./RoleColumns";

export const RoleList = () => {
  const s = useS(strings);
  const uiS = useS(uiStrings);

  return (
    <>
      <CommonListManager
        columns={columns(s, uiS)}
        queryHook={useRoleBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          RoleNavigation.single(uniqueId)
        }
        deleteHook={useRoleAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
