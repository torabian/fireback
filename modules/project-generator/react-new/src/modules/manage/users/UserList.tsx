import { usePageTitle } from "../../fireback-ui/components/page-title/PageTitle";
import { useUserBrowseActionQuery } from "../../sdk/abac/UserBrowseAction";
import { useUserAwareDeleteAction } from "../../sdk/abac/UserAwareDeleteAction";

import { useS } from "../../fireback-ui/hooks/useS";
import { strings as uiStrings } from "../../fireback-ui/components/strings/translations";

import { CommonListManager } from "../../fireback-ui/components/entity-manager/CommonListManager";
import { UserNavigation } from "../../sdk/navigation/AbacNavigation";
import { columns } from "./UserColumns";
import { strings } from "./strings/translations";

export const UserList = () => {
  const s = useS(strings);
  const uiS = useS(uiStrings);
  usePageTitle(s.menuTitle);

  return (
    <>
      <CommonListManager
        columns={columns(s, uiS)}
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
