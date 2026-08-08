import { usePageTitle } from "../../../../fireback-ui/components/page-title/PageTitle";
import { useUserBrowseActionQuery } from "../../../sdk/abac/UserBrowseAction";
import { useUserAwareDeleteAction } from "../../../sdk/abac/UserAwareDeleteAction";

import { useT } from "../../../../fireback-ui/hooks/useT";
import { useS } from "../../../../fireback-ui/hooks/useS";

import { CommonListManager } from "../../../../fireback-ui/components/entity-manager/CommonListManager";
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
