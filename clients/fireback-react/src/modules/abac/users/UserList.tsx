import { usePageTitle } from "@/components/page-title/PageTitle";

import { useT } from "@/hooks/useT";

import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { useDeleteUser } from "src/sdk/fireback/modules/workspaces/useDeleteUser";
import { useGetUsers } from "src/sdk/fireback/modules/workspaces/useGetUsers";
import { columns } from "./UserColumns";
import { UserEntity } from "@/sdk/fireback/modules/workspaces/UserEntity";

export const UserList = () => {
  const t = useT();
  usePageTitle(t.fbMenu.users);

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
