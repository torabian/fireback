import { CommonListManager } from "../../../../fireback-ui/components/entity-manager/CommonListManager";
import { useT } from "../../../../fireback-ui/hooks/useT";
import { useRoleAwareDeleteAction } from "../../../sdk/abac/RoleAwareDeleteAction";
import { useRoleBrowseActionQuery } from "../../../sdk/abac/RoleBrowseAction";
import { RoleNavigation } from "../../../sdk/navigation/AbacNavigation";
import { columns } from "./RoleColumns";

export const RoleList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useRoleBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          RoleNavigation.single(uniqueId)
        }
        deleteHook={useRoleAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
