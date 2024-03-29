import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { useT } from "@/hooks/useT";
import { RoleEntity } from "@/sdk/fireback/modules/workspaces/RoleEntity";
import { useDeleteRole } from "@/sdk/fireback/modules/workspaces/useDeleteRole";
import { useGetRoles } from "@/sdk/fireback/modules/workspaces/useGetRoles";
import { columns } from "./RoleColumns";

export const RoleList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetRoles}
        uniqueIdHrefHandler={(uniqueId: string) =>
          RoleEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteRole}
      ></CommonListManager>
    </>
  );
};
