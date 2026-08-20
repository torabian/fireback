import { useUserAwareDeleteAction } from "@fireback/manage/sdk/abac/UserAwareDeleteAction";
import {
  UserBrowseActionQueryParams,
  useUserBrowseActionQuery,
} from "@fireback/manage/sdk/abac/UserBrowseAction";
import { usePageTitle } from "@fireback/ui-core/components/page-title/PageTitle";

import { strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { useS } from "@fireback/ui-core/hooks/useS";

import { CommonListManager2 } from "@fireback/ui-core/components/entity-manager/CommonListManager2";
import {
  buildJsonLogic,
  isEmptyJsonLogic,
} from "@fireback/ui-core/components/data-grid-list/jsonLogic";
import { UserNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { columns } from "./UserColumns";
import { strings } from "./strings/translations";

// Meta keys useDatatableFiltering's `filters` mixes in alongside the actual
// per-column values (see CommonListManager2.tsx's castColumns onFilterChange,
// which does `udf.setFilter({ [field]: value })` into the very same object) -
// everything else left over is a real column filter.
const FILTER_META_KEYS = new Set([
  "itemsPerPage",
  "startIndex",
  "cursor",
  "sort",
  "sorting",
]);

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

          qs.setSort(state.udf.debouncedFilters.sort);
          // Same JsonLogic shape DataGridList's useColumnFilters builds (see
          // jsonLogic.ts's buildJsonLogic doc comment) - reusing it here keeps
          // the "filter" query param identical between CommonListManager2 and
          // DataGridList, since both eventually reach ApplyQueryFilter/
          // jsonlogic2sql on the backend.
          const conditions = Object.entries(state.udf.debouncedFilters)
            .filter(([field]) => !FILTER_META_KEYS.has(field))
            .map(([field, value]) => ({ field, value: String(value ?? "") }));
          const jsonLogic = buildJsonLogic(conditions);
          qs.setFilter(
            isEmptyJsonLogic(jsonLogic) ? null : JSON.stringify(jsonLogic),
          );

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
