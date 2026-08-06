import { usePageTitle } from "../../../components/page-title/PageTitle";

import { useT } from "../../../hooks/useT";
import { useS } from "../../../hooks/useS";

import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { useGetUsers } from "../../../sdk/modules/abac/useGetUsers";
import { UserEntity } from "../../../sdk/modules/abac/UserEntity";
import { columns } from "./UserColumns";
import { strings } from "./strings/translations";
import { usePostUserRemove } from "@/modules/fireback/sdk/modules/abac/usePostUserRemove";

export const UserList = () => {
  const t = useT();
  const s = useS(strings);
  usePageTitle(t.fbMenu.users);

  return (
    <>
      <CommonListManager
        columns={columns(t, s)}
        // CardComponent={UserCard}
        queryHook={useGetUsers}
        uniqueIdHrefHandler={(uniqueId: string) =>
          UserEntity.Navigation.single(uniqueId)
        }
        deleteHook={usePostUserRemove}
      ></CommonListManager>
    </>
  );
};
