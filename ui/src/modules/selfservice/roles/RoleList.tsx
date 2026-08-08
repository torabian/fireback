import { CommonListManager } from "../../fireback-ui/components/entity-manager/CommonListManager";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings as uiStrings } from "../../fireback-ui/components/strings/translations";
import { strings } from "./strings/translations";
import { useRoleAwareDeleteAction } from "../../sdk/abac/RoleAwareDeleteAction";
import { useRoleBrowseActionQuery } from "../../sdk/abac/RoleBrowseAction";
import { RoleNavigation } from "../../sdk/navigation/AbacNavigation";
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
