import { useT } from "@/modules/fireback/hooks/useT";
import { columns } from "./RoleColumns";
import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useGetRoles } from "@/modules/fireback/sdk/modules/abac/useGetRoles";
import { RoleEntity } from "@/modules/fireback/sdk/modules/abac/RoleEntity";
import { useDeleteRole } from "@/modules/fireback/sdk/modules/abac/useDeleteRole";

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
