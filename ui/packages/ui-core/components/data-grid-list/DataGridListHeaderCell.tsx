import { ArrowDownAZ, ArrowDownWideNarrow, ArrowDownZA, Languages } from "lucide-react";
import { useEffect, useState } from "react";
import { type RenderHeaderCellProps } from "react-data-grid";
import { useOverlay } from "../overlay/OverlayProvider";
import { strings as coreStrings } from "../strings/translations";
import { useS } from "../../hooks/useS";
import { useColumnFiltersContext } from "./hooks/useColumnFilters";
import { buildTStringFilterCondition } from "./jsonLogic";
import { TStringFilterDrawer, type TStringFilterDrawerResult } from "./TStringFilterDrawer";

// DataGridList's own header cell: sort button + filter input, wired to
// useColumnFilters instead of the old devexpress-Filter-shaped `Udf`
// (common-data-table/PaginateHeaderCell.tsx) - every keystroke here ends up in
// useColumnFilters' debounced JsonLogic tree (jsonLogic.ts) instead of a qs-
// encoded filters object.
//
// Reads the live ColumnFilters via useColumnFiltersContext(), not a prop - see that
// context's own doc comment (hooks/useColumnFilters.ts) for why: DataGridList's
// column definitions (and this cell's own renderHeaderCell closure) are deliberately
// NOT recreated on every filter keystroke, so a plain prop would go stale.
//
// filterType "tstring" (a complexes.TString column - see DatatableColumn.tsx's own
// doc comment) is the one exception: rather than a plain text input, it renders a
// button that opens TStringFilterDrawer to collect one search value per language,
// which then becomes an "or" of per-locale "contains" conditions (jsonLogic.ts's
// buildTStringFilterCondition) fed into setColumnFilter as a raw JsonLogic
// sub-expression - the same "{"/"[" escape hatch buildJsonLogic already supports for
// any column, just exercised programmatically here instead of hand-typed.
export function DataGridListHeaderCell<R>({
  column,
  filterType,
  sortable,
  filterable,
  filterKey,
}: RenderHeaderCellProps<R> & {
  filterType?: "string" | "date" | "tstring";
  filterable?: boolean;
  sortable?: boolean;
  // The field name to filter by, when it differs from column.key - see
  // DatatableColumn.tsx's own doc comment. Sorting still keys off
  // column.key - only the filter value/JsonLogic field name is affected.
  filterKey?: string;
}) {
  const columnFilters = useColumnFiltersContext();
  const sortField = column.key;
  const field = filterKey ?? column.key;
  const [internalValue, setInternalValue] = useState(columnFilters.values[field] ?? "");

  useEffect(() => {
    setInternalValue(columnFilters.values[field] ?? "");
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [columnFilters.values[field]]);

  const sort = columnFilters.sort?.columnName === sortField ? columnFilters.sort.direction : undefined;

  return (
    <>
      {sortable ? (
        <span className="data-table-sort-actions">
          <button
            className={`active-sort-col ${sort ? "active" : ""}`}
            onClick={() => columnFilters.toggleSort(sortField)}
          >
            {sort === "asc" && <ArrowDownAZ className="sort-icon" />}
            {sort === "desc" && <ArrowDownZA className="sort-icon" />}
            {!sort && <ArrowDownWideNarrow className="sort-icon" />}
          </button>
        </span>
      ) : null}
      {filterable && filterType === "tstring" ? (
        <TStringHeaderFilter field={field} columnTitle={(column.name as string) || ""} />
      ) : filterable ? (
        <input
          className="data-table-filter-input"
          type={filterType === "date" ? "date" : "text"}
          onKeyDown={(e) => {
            if (e.code === "Space") {
              e.stopPropagation();
            }
          }}
          value={internalValue}
          onChange={(e) => {
            setInternalValue(e.target.value);
            columnFilters.setColumnFilter(field, e.target.value);
          }}
          placeholder={(column.name as string) || ""}
        />
      ) : (
        <span>{column.name}</span>
      )}
    </>
  );
}

function TStringHeaderFilter({ field, columnTitle }: { field: string; columnTitle: string }) {
  const cs = useS(coreStrings);
  const columnFilters = useColumnFiltersContext();
  const { openDrawer } = useOverlay();

  // Per-locale values typed into the drawer, so reopening it shows what was there
  // before - held in useColumnFilters' own structuredValues (not local state here):
  // columnFilters.values[field] only ever holds the *serialized JsonLogic* string
  // built from these, not the map itself, and this component isn't guaranteed to
  // survive every DataGridList re-render with local state intact (react-data-grid
  // rebuilds header cells fairly often - on sort changes, on the underlying rows
  // changing after a fetch, ...). See setColumnFilterStructured's own doc comment.
  const perLocaleValues = (columnFilters.structuredValues[field] as Record<string, string>) ?? {};

  const activeCount = Object.values(perLocaleValues).filter((v) => v.trim() !== "").length;
  const firstActive = Object.entries(perLocaleValues).find(([, v]) => v.trim() !== "");

  const open = () => {
    openDrawer<TStringFilterDrawerResult>((props) => (
      <TStringFilterDrawer
        {...props}
        fieldLabel={columnTitle}
        initialValues={perLocaleValues}
      />
    )).promise.then(({ type, data }) => {
      if (type !== "resolved" || !data) {
        return;
      }
      const condition = buildTStringFilterCondition(field, data.values);
      columnFilters.setColumnFilterStructured(
        field,
        data.values,
        condition ? JSON.stringify(condition) : "",
      );
    });
  };

  return (
    <button
      type="button"
      className="data-table-filter-input data-table-filter-tstring"
      onClick={open}
      title={columnTitle}
    >
      <Languages className="sort-icon" />
      {activeCount > 0 ? (
        <span>
          {firstActive?.[1]}
          {activeCount > 1 ? ` (+${activeCount - 1})` : ""}
        </span>
      ) : (
        <span className="data-table-filter-tstring__placeholder">
          {cs.table.filter.filterPlaceholder}
        </span>
      )}
    </button>
  );
}
