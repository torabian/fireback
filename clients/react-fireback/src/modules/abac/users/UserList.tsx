import { usePageTitle } from "@/components/page-title/PageTitle";

import { useT } from "@/hooks/useT";

import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { useDeleteUser } from "src/sdk/fireback/modules/workspaces/useDeleteUser";
import { useGetUsers } from "src/sdk/fireback/modules/workspaces/useGetUsers";
import { UserNavigationTools } from "src/sdk/fireback/modules/workspaces/user-navigation-tools";
import { columns } from "./UserColumns";

export const UserList = () => {
  const t = useT();
  usePageTitle(t.fbMenu.users);

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetUsers}
        uniqueIdHrefHandler={(uniqueId: string) =>
          UserNavigationTools.single(uniqueId)
        }
        deleteHook={useDeleteUser}
      ></CommonListManager>
    </>
  );
};
