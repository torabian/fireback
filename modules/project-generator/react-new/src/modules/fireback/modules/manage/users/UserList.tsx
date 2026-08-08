import { usePageTitle } from "../../../components/page-title/PageTitle";
import { useUserBrowseActionQuery } from "../../../sdk/abac/UserBrowseAction";
import { useUserAwareDeleteAction } from "../../../sdk/abac/UserAwareDeleteAction";

import { useT } from "../../../hooks/useT";
import { useS } from "../../../hooks/useS";

import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { UserNavigation } from "../../../sdk/navigation/AbacNavigation";
import { columns } from "./UserColumns";
import { strings } from "./strings/translations";

export const UserList = () => {
  const t = useT();
  const s = useS(strings);
  usePageTitle(t.fbMenu.users);

  return (
    <>
      <CommonListManager
        columns={columns(t, s)}
        // CardComponent={UserCard}
        queryHook={useUserBrowseActionQuery}
        uniqueIdHrefHandler={(uniqueId: string) =>
          UserNavigation.single(uniqueId)
        }
        deleteHook={useUserAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
