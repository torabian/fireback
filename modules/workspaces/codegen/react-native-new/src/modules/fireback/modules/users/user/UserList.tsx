import { useT } from "@/hooks/useT";
import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { columns } from "./UserColumns";
import { UserEntity } from "src/sdk/fireback/modules/workspaces/UserEntity";
import { useGetUsers } from "src/sdk/fireback/modules/workspaces/useGetUsers";
import { useDeleteUser } from "@/sdk/fireback/modules/workspaces/useDeleteUser";
export const UserList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetUsers}
        uniqueIdHrefHandler={(uniqueId: string) =>
          UserEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteUser}
      ></CommonListManager>
    </>
  );
};