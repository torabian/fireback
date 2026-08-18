import { createContext, useContext, useMemo, useState } from "react";
import { useDebouncedEffect } from "../../../hooks/useDebouncedEffect";
import { buildJsonLogic, isEmptyJsonLogic, type JsonLogicNode } from "../jsonLogic";
import { type SingleSort } from "./useCursor";

// Per-column filter state + debounce, feeding straight into DataGridList's
// JsonLogic query builder (jsonLogic.ts) - the DataGridList analogue of
// useDatatableFiltering's `filters`/`debouncedFilters` pair, but scoped to just
// column values (selection/delete lives in its own useRowSelectionDelete hook,
// paging/sorting the caller owns via useCursor) and producing a JsonLogic tree
// instead of a raw qs-encoded object.
//
// Every keystroke updates `values` immediately (so inputs stay responsive), but
// `debouncedJsonLogic` - what callers actually fetch with - only updates
// `delayMs` after the *last* keystroke across *any* column, same one-timer
// coalescing useDatatableFiltering's withDebounce already does.
export function useColumnFilters({ delayMs = 500 }: { delayMs?: number } = {}) {
  const [values, setValues] = useState<Record<string, string>>({});
  const [debouncedValues, setDebouncedValues] = useState<Record<string, string>>({});
  const [sort, setSort] = useState<SingleSort | undefined>(undefined);

  // Structured, per-field "what did the user actually pick" state - for a column
  // whose filter isn't just typed text (e.g. TStringHeaderFilter's per-locale values,
  // built into `values[field]` as a serialized JsonLogic string), this is where the
  // *pre-serialization* shape lives, so reopening its editor (a drawer, a picker,
  // whatever) can show what was there before. Deliberately lives here - in
  // useColumnFilters' own state, one level up from any individual header cell - rather
  // than as header-cell-local state: DataGridList's rows/columns re-render a fair bit
  // (every fetch, every filter, every sort), and a header cell surviving all of that
  // with its identity/local state intact isn't guaranteed the way this hook's own
  // state (tied to DataGridList's own component instance) is.
  const [structuredValues, setStructuredValues] = useState<Record<string, unknown>>({});

  const { withDebounce } = useDebouncedEffect();

  const setColumnFilter = (field: string, value: string) => {
    const next = { ...values, [field]: value };
    setValues(next);
    withDebounce(() => setDebouncedValues(next), delayMs);
  };

  // Same as setColumnFilter, but also remembers the pre-serialization value under
  // structuredValues[field] - see that field's own doc comment above.
  const setColumnFilterStructured = (field: string, structured: unknown, value: string) => {
    setStructuredValues((prev) => ({ ...prev, [field]: structured }));
    setColumnFilter(field, value);
  };

  const clearColumnFilter = (field: string) => setColumnFilter(field, "");

  const jsonLogic = useMemo<JsonLogicNode>(
    () =>
      buildJsonLogic(
        Object.entries(debouncedValues).map(([field, value]) => ({ field, value })),
      ),
    [debouncedValues],
  );

  const hasActiveFilters = !isEmptyJsonLogic(jsonLogic);

  const toggleSort = (columnName: string) => {
    setSort((current) => {
      if (!current || current.columnName !== columnName) {
        return { columnName, direction: "asc" };
      }
      if (current.direction === "asc") {
        return { columnName, direction: "desc" };
      }
      return undefined;
    });
  };

  return {
    values,
    setColumnFilter,
    structuredValues,
    setColumnFilterStructured,
    clearColumnFilter,
    debouncedValues,
    jsonLogic,
    hasActiveFilters,
    sort,
    setSort,
    toggleSort,
  };
}

export type ColumnFilters = ReturnType<typeof useColumnFilters>;

// DataGridListHeaderCell reads the live ColumnFilters through this context instead of
// as a plain prop threaded through the memoized `cols`/`renderHeaderCell` closure in
// DataGridList.tsx. That closure is deliberately only rebuilt when the *column
// definitions* themselves change (columns/selectable/withFilters), not on every
// keystroke - rebuilding it per keystroke would remount every header cell (react-data-
// grid treats a new columns array as new columns), which for TStringHeaderFilter
// (data-grid-list/DataGridListHeaderCell.tsx) meant its own locally-held per-locale
// values were wiped out the instant its own filter update landed. Context reads happen
// at render time, not closure-creation time, so it sidesteps the whole issue: the
// column definitions can stay stable while the filter *values* they read are always
// current.
const ColumnFiltersContext = createContext<ColumnFilters | null>(null);

export const ColumnFiltersProvider = ColumnFiltersContext.Provider;

export function useColumnFiltersContext(): ColumnFilters {
  const ctx = useContext(ColumnFiltersContext);
  if (!ctx) {
    throw new Error("useColumnFiltersContext must be used inside a ColumnFiltersProvider");
  }
  return ctx;
}
