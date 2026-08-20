import { useUserAwareDeleteAction } from "@fireback/manage/sdk/abac/UserAwareDeleteAction";
import { useUserBrowseActionQuery } from "@fireback/manage/sdk/abac/UserBrowseAction";
import { usePageTitle } from "@fireback/ui-core/components/page-title/PageTitle";

import { strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { useS } from "@fireback/ui-core/hooks/useS";

import { CommonListManager } from "@fireback/ui-core/components/entity-manager/CommonListManager";
import { createUdfBrowseQueryHook } from "@fireback/ui-core/hooks/useUdfBrowseQuery";
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
        queryHook={createUdfBrowseQueryHook(useUserBrowseActionQuery)}
        uniqueIdHrefHandler={(uniqueId: string) =>
          UserNavigation.single(uniqueId)
        }
        deleteHook={useUserAwareDeleteAction}
      ></CommonListManager>
    </>
  );
};
