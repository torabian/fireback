import { usePageTitle } from "@fireback/ui-core/components/page-title/PageTitle";
import { useUserBrowseActionQuery } from "@fireback/manage/sdk/abac/UserBrowseAction";
import { useUserAwareDeleteAction } from "@fireback/manage/sdk/abac/UserAwareDeleteAction";

import { useS } from "@fireback/ui-core/hooks/useS";
import { strings as uiStrings } from "@fireback/ui-core/components/strings/translations";

import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
import { UserNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
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
