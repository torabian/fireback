import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { useRoleAwareDeleteAction } from "@/modules/fireback/sdk/abac/RoleAwareDeleteAction";
import { useRoleBrowseActionQuery } from "@/modules/fireback/sdk/abac/RoleBrowseAction";
import { RoleNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
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
