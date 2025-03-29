import { get } from "lodash";
import { ArrowDownAZ, ArrowDownWideNarrow, ArrowDownZA } from "lucide-react";
import { useEffect, useState } from "react";
import { RenderHeaderCellProps } from "react-data-grid";
import { Udf } from "../../hooks/useDatatableFiltering";

export function FilterRenderer<R>({
  tabIndex,
  column,
  filterType,
  sortable,
  filterable,
  selectable,
  udf,
}: RenderHeaderCellProps<R> & {
  filterType: any;
  filterable: boolean;
  selectable: boolean;
  sortable?: boolean;
  udf: Udf;
}) {
  // Single sort for now, assumes 1st one.
  const columnSort = udf.filters.sorting?.find(
    (col) => col.columnName === column.key
  );

  const [internalValue, setInternalValue] = useState("");

  useEffect(() => {
    if (internalValue !== get(udf.filters, column.key)) {
      setInternalValue(get(udf.filters, column.key));
    }
  }, [udf.filters]);

  let sorting: "asc" | "desc" | undefined = undefined;
  if (columnSort?.columnName === column.key && columnSort?.direction == "asc") {
    sorting = "asc";
  }
  if (
    columnSort?.columnName === column.key &&
    columnSort?.direction == "desc"
  ) {
    sorting = "desc";
  }

  const onSortButtonClick = () => {
    if (columnSort) {
      if (columnSort?.direction === "desc") {
        udf.setSorting(
          udf.filters.sorting.filter((m) => m.columnName !== column.key)
        );
      }

      if (columnSort?.direction === "asc") {
        udf.setSorting(
          udf.filters.sorting.map((m) => {
            if (m.columnName === column.key) {
              return {
                ...m,
                direction: "desc",
              };
            }
            return m;
          })
        );
      }
    } else {
      udf.setSorting([
        ...udf.filters.sorting,
        {
          columnName: column.key.toString(),
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
              column.key == columnSort?.columnName ? "active" : ""
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
                udf.setFilter({ [column.key]: e.target.value });
              }}
              placeholder={(column.name as any) || ""}
              type="date"
            />
          ) : (
            <input
              className="data-table-filter-input"
              tabIndex={tabIndex}
              value={internalValue}
              onChange={(e) => {
                setInternalValue(e.target.value);
                udf.setFilter({ [column.key]: e.target.value });
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
