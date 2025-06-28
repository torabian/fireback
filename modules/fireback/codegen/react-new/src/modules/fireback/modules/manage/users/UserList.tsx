import { usePageTitle } from "../../../components/page-title/PageTitle";

import { useT } from "../../../hooks/useT";

import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { useDeleteUser } from "../../../sdk/modules/abac/useDeleteUser";
import { useGetUsers } from "../../../sdk/modules/abac/useGetUsers";
import { UserEntity } from "../../../sdk/modules/abac/UserEntity";
import { columns } from "./UserColumns";

export const UserList = () => {
  const t = useT();
  usePageTitle(t.fbMenu.users);

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        // CardComponent={UserCard}
        queryHook={useGetUsers}
        uniqueIdHrefHandler={(uniqueId: string) =>
          UserEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteUser}
      ></CommonListManager>
    </>
  );
};
