import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { RoleEntity } from "../../sdk/modules/workspaces/RoleEntity";
import { useDeleteRole } from "../../sdk/modules/workspaces/useDeleteRole";
import { useGetRoles } from "../../sdk/modules/workspaces/useGetRoles";
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
