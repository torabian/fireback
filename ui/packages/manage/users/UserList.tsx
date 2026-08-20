import { useUserAwareDeleteAction } from "@fireback/manage/sdk/abac/UserAwareDeleteAction";
import {
  UserBrowseActionQueryParams,
  useUserBrowseActionQuery,
} from "@fireback/manage/sdk/abac/UserBrowseAction";
import { usePageTitle } from "@fireback/ui-core/components/page-title/PageTitle";

import { strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { useS } from "@fireback/ui-core/hooks/useS";

import { CommonListManager2 } from "@fireback/ui-core/components/entity-manager/CommonListManager2";
import { UserNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { columns } from "./UserColumns";
import { strings } from "./strings/translations";

export const UserList = () => {
  const s = useS(strings);
  const uiS = useS(uiStrings);
  usePageTitle(s.menuTitle);

  return (
    <>
      <CommonListManager2
        columns={columns(s, uiS)}
        queryHook={({ state }) => {
          const qs = new UserBrowseActionQueryParams();
          qs.setItemsPerPage(state.udf.debouncedFilters.itemsPerPage);
          qs.setCursor(state.udf.debouncedFilters.cursor);
          return useUserBrowseActionQuery({ qs });
        }}
        uniqueIdHrefHandler={(uniqueId: string) =>
          UserNavigation.single(uniqueId)
        }
        deleteHook={useUserAwareDeleteAction}
      ></CommonListManager2>
    </>
  );
};
