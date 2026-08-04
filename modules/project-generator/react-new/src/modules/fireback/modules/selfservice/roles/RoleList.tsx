import { CommonListManager } from "@/modules/fireback/components/entity-manager/CommonListManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { RoleEntity } from "@/modules/fireback/sdk/modules/abac/RoleEntity";
import { useGetRoles } from "@/modules/fireback/sdk/modules/abac/useGetRoles";
import { usePostRoleRemove } from "@/modules/fireback/sdk/modules/abac/usePostRoleRemove";
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
        deleteHook={usePostRoleRemove}
      ></CommonListManager>
    </>
  );
};
