import { URLSearchParamsX } from "@fireback/js-remote-ctx/common/URLSearchParamsX";
import { buildJsonLogic, isEmptyJsonLogic } from "../components/data-grid-list/jsonLogic";
import { type Udf } from "./useDatatableFiltering";

// Meta keys useDatatableFiltering's `filters`/`debouncedFilters` mixes in
// alongside the actual per-column values (see CommonListManager2.tsx's
// castColumns onFilterChange, which does `udf.setFilter({ [field]: value })`
// into the very same object) - everything else left over is a real column
// filter, to be turned into a JsonLogic "contains" condition.
const FILTER_META_KEYS = new Set([
  "itemsPerPage",
  "startIndex",
  "cursor",
  "sort",
  "sorting",
]);

/**
 * Builds the qs a CommonListManager2 queryHook needs (filter/sort/itemsPerPage/
 * cursor) from a udf's debouncedFilters. Every emi Browse action shares this
 * exact qs shape (see e.g. UserBrowseAction.ts's `UserBrowseActionQueryParams
 * extends URLSearchParamsX`, with `filter`/`sort`/`itemsPerPage`/`cursor` as
 * its only fields) - and since each of those generated classes' typed
 * setFilter/setSort/... methods are themselves just thin wrappers over
 * URLSearchParamsX's own inherited `set()`, a plain `URLSearchParamsX` built
 * here serializes identically to one of those subclasses, without needing a
 * reference to the specific generated class at all.
 *
 * Deliberately a plain function, not a hook - so a queryHook that needs to
 * add its own qs params (scope, a fixed extra filter, ...) or pass extra
 * options through to its useXxxBrowseActionQuery call (headers, ctx,
 * overrideUrl, creatorFn, ...) can just call this for the common part and
 * take it from there, instead of useUdfBrowseQuery needing some bespoke
 * callback-based escape hatch for every possible customization:
 *
 *   queryHook={({ state }) => {
 *     const qs = buildUdfBrowseQs(state.udf);
 *     qs.set("workspaceId", currentWorkspaceId); // extra qs param
 *     return useUserBrowseActionQuery({ qs, headers: myHeaders }); // extra option
 *   }}
 */
export function buildUdfBrowseQs(
  udf: Pick<Udf, "debouncedFilters">,
): URLSearchParamsX {
  const qs = new URLSearchParamsX();
  const filters = udf.debouncedFilters;

  qs.set("itemsPerPage", filters.itemsPerPage);
  qs.set("sort", filters.sort);
  qs.set("cursor", filters.cursor);

  const conditions = Object.entries(filters)
    .filter(([field]) => !FILTER_META_KEYS.has(field))
    .map(([field, value]) => ({ field, value: String(value ?? "") }));
  const jsonLogic = buildJsonLogic(conditions);
  qs.set("filter", isEmptyJsonLogic(jsonLogic) ? null : JSON.stringify(jsonLogic));

  return qs;
}

/**
 * Convenience wrapper over buildUdfBrowseQs for the common case: no extra qs
 * params, no extra options - just hand udf straight to whichever
 * emi-generated `useXxxBrowseActionQuery` hook the caller passes in. Reach
 * for buildUdfBrowseQs directly instead (see its own doc comment) as soon as
 * you need to customize the qs or pass extra options through.
 *
 *   queryHook={({ state }) =>
 *     useUdfBrowseQuery(state.udf, useUserBrowseActionQuery)
 *   }
 */
export function useUdfBrowseQuery<TResult>(
  udf: Pick<Udf, "debouncedFilters">,
  useBrowseActionQuery: (options: { qs: any }) => TResult,
): TResult {
  return useBrowseActionQuery({ qs: buildUdfBrowseQs(udf) });
}

/**
 * One step more convenient than useUdfBrowseQuery for CommonListManager2's
 * queryHook prop specifically: pass it a useXxxBrowseActionQuery hook once,
 * get back a ready-made queryHook - no `({ state }) => ...` wrapper to write
 * per entity at all.
 *
 *   queryHook={createUdfBrowseQueryHook(useUserBrowseActionQuery)}
 *
 * Not named with a `use` prefix on purpose, same as the queryHook prop it
 * builds - it's a plain factory returning a callback CommonListManager2
 * invokes directly (`queryHook({ state })`), not something a component calls
 * as a hook of its own. Reach for buildUdfBrowseQs/useUdfBrowseQuery instead
 * (see their own doc comments) as soon as the qs or the query options need
 * anything beyond what udf.debouncedFilters already carries.
 */
export function createUdfBrowseQueryHook<TResult>(
  useBrowseActionQuery: (options: { qs: any }) => TResult,
): (args: { state: { udf: Pick<Udf, "debouncedFilters"> } }) => TResult {
  return ({ state }) => useUdfBrowseQuery(state.udf, useBrowseActionQuery);
}
