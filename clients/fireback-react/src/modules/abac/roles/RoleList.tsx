import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { useT } from "@/hooks/useT";
import { useDeleteRole } from "@/sdk/fireback/modules/workspaces/useDeleteRole";
import { useGetRoles } from "@/sdk/fireback/modules/workspaces/useGetRoles";
import { RoleNavigationTools } from "src/sdk/fireback/modules/workspaces/role-navigation-tools";
import { columns } from "./RoleColumns";

export const RoleList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetRoles}
        uniqueIdHrefHandler={(uniqueId: string) =>
          RoleNavigationTools.single(uniqueId)
        }
        deleteHook={useDeleteRole}
      ></CommonListManager>
    </>
  );
};
