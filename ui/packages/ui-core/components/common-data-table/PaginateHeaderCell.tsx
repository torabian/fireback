import { get } from "lodash";
import { ArrowDownAZ, ArrowDownWideNarrow, ArrowDownZA } from "lucide-react";
import { useEffect, useState } from "react";
import { type RenderHeaderCellProps } from "react-data-grid";
import { type Udf } from "../../hooks/useDatatableFiltering";

export function FilterRenderer<R>({
  tabIndex,
  column,
  filterType,
  sortable,
  filterable,
  selectable,
  filterKey,
  udf,
}: RenderHeaderCellProps<R> & {
  filterType: any;
  filterable: boolean;
  selectable: boolean;
  sortable?: boolean;
  // The key udf.filters/setFilter actually reads/writes - see
  // DatatableColumn.tsx's own doc comment. Falls back to column.key so
  // callers that never set it keep working exactly as before.
  filterKey?: string;
  udf: Udf;
}) {
  // Same key the "sort" query string is built from (useDatatableFiltering's
  // toSortString does `${columnName} ${direction}` verbatim) - so this needs
  // to be the backend column name too, exactly like the filter field above.
  const field = filterKey ?? (column.key as string);

  // Single sort for now, assumes 1st one.
  const columnSort = udf.filters.sorting?.find(
    (col) => col.columnName === field,
  );

  const [internalValue, setInternalValue] = useState("");

  useEffect(() => {
    if (internalValue !== get(udf.filters, field)) {
      setInternalValue(get(udf.filters, field));
    }
  }, [udf.filters]);

  let sorting: "asc" | "desc" | undefined = undefined;
  if (columnSort?.columnName === field && columnSort?.direction == "asc") {
    sorting = "asc";
  }
  if (
    columnSort?.columnName === field &&
    columnSort?.direction == "desc"
  ) {
    sorting = "desc";
  }

  const onSortButtonClick = () => {
    if (columnSort) {
      if (columnSort?.direction === "desc") {
        udf.setSorting(
          udf.filters.sorting.filter((m) => m.columnName !== field),
        );
      }

      if (columnSort?.direction === "asc") {
        udf.setSorting(
          udf.filters.sorting.map((m) => {
            if (m.columnName === field) {
              return {
                ...m,
                direction: "desc",
              };
            }
            return m;
          }),
        );
      }
    } else {
      udf.setSorting([
        ...udf.filters.sorting,
        {
          columnName: field,
          direction: "asc",
        },
      ]);
    }
  };

  return (
    <>
      {sortable ? (
        <span className="data-table-sort-actions">
          <button
            className={`active-sort-col ${
              field == columnSort?.columnName ? "active" : ""
            }`}
            onClick={onSortButtonClick}
          >
            {sorting == "asc" ? <ArrowDownAZ className="sort-icon" /> : null}
            {sorting == "desc" ? <ArrowDownZA className="sort-icon" /> : null}
            {sorting === undefined ? (
              <ArrowDownWideNarrow className="sort-icon" />
            ) : null}
          </button>
        </span>
      ) : null}
      {filterable ? (
        <>
          {filterType === "date" ? (
            <input
              className="data-table-filter-input"
              tabIndex={tabIndex}
              value={internalValue}
              onChange={(e) => {
                setInternalValue(e.target.value);
                udf.setFilter({ [field]: e.target.value });
              }}
              placeholder={(column.name as any) || ""}
              type="date"
            />
          ) : (
            <input
              className="data-table-filter-input"
              tabIndex={tabIndex}
              onKeyDown={(e) => {
                if (e.code === "Space") {
                  e.stopPropagation();
                }
              }}
              value={internalValue}
              onChange={(e) => {
                setInternalValue(e.target.value);
                udf.setFilter({ [field]: e.target.value });
              }}
              placeholder={(column.name as any) || ""}
            />
          )}
        </>
      ) : (
        <span>{column.name}</span>
      )}
    </>
  );
}
