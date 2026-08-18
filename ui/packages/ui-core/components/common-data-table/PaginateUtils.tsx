import "react-data-grid/lib/styles.css";

import { get } from "lodash";
import { type ColumnOrColumnGroup } from "react-data-grid";
import { type DatatableColumn } from "../../types/DatatableColumn";
import { type Udf } from "../../hooks/useDatatableFiltering";
import { CopyCell } from "../entity-manager/CopyCell";
import { OpenInNewRouter } from "../entity-manager/OpenInNewRouter";
import Link from "../link/Link";
import { FilterRenderer } from "./PaginateHeaderCell";

export interface TableColumnWidthInfo {
  /** A column name. */
  columnName: string;
  /** A column width. */
  width: number | string;
}

/**
 * JSON.parse that returns undefined instead of throwing - used for values
 * that were never actually typechecked (a server response field typed as a
 * plain string, a raw localStorage read), so a malformed one is just treated
 * as "nothing saved" rather than crashing the render that reads it.
 */
export function parseJsonSafely(raw: string | null | undefined): unknown {
  if (!raw) {
    return undefined;
  }
  try {
    return JSON.parse(raw);
  } catch {
    return undefined;
  }
}

function isValidColumnWidth(width: unknown): width is number | string {
  if (typeof width === "number") {
    return Number.isFinite(width) && width > 0;
  }
  if (typeof width === "string") {
    return width.length > 0;
  }
  return false;
}

/**
 * Validates/cleans a possibly-untrusted TableColumnWidthInfo[] - parsed from
 * a server response or localStorage, so it's never actually typechecked at
 * runtime - before it's trusted anywhere near react-data-grid. One malformed
 * entry (a null width, a non-finite/negative width, the value not even being
 * an array - e.g. a 404's `{"error": ...}` body, or a stale localStorage
 * write from before a column was added/removed) used to reach
 * react-data-grid's own column-width computation and crash it with "Cannot
 * read properties of null (reading 'width')", taking the whole table down
 * (see CommonListManager.tsx). Returns undefined - never a partially-bad
 * array - for anything that doesn't check out, so callers fall back to their
 * own known-good default instead.
 */
export function normalizeColumnWidths(
  raw: unknown,
  columns: Array<{ name?: string }>,
): TableColumnWidthInfo[] | undefined {
  if (!Array.isArray(raw) || raw.length === 0) {
    return undefined;
  }

  const cleaned: TableColumnWidthInfo[] = [];
  for (const entry of raw) {
    if (
      !entry ||
      typeof entry !== "object" ||
      typeof (entry as any).columnName !== "string" ||
      !(entry as any).columnName ||
      !isValidColumnWidth((entry as any).width)
    ) {
      return undefined;
    }
    cleaned.push({
      columnName: (entry as any).columnName,
      width: (entry as any).width,
    });
  }

  // A saved sizing has to actually describe *this* table's columns - a
  // stale/foreign payload (a wholly different column set, e.g. two tables
  // once sharing the same tableKey - see CommonListManager.tsx's own
  // historical bug fix note above tableKey) is exactly the shape mismatch
  // that fed react-data-grid something it didn't expect.
  const knownNames = new Set(columns.map((c) => c.name).filter(Boolean));
  const coversKnownColumn = cleaned.some((c) => knownNames.has(c.columnName));
  if (!coversKnownColumn) {
    return undefined;
  }

  return cleaned;
}

function resolvePath(base, relative) {
  const stack = base.split("/").filter(Boolean);
  const parts = relative.split("/");

  parts.forEach((part) => {
    if (part === "..") {
      stack.pop();
    } else if (part !== "." && part !== "") {
      stack.push(part);
    }
  });

  return "/" + stack.join("/");
}

export const castColumns = (
  columns: DatatableColumn[],
  setFilter: (field: string, value: any) => void,
  udf: Udf,
  columnSizes: TableColumnWidthInfo[] = [],
  uniqueIdHrefHandler: (id: any) => string,
  currentRouterPath: string
): ColumnOrColumnGroup<any, unknown>[] => {
  return columns.map((col) => {
    const info = columnSizes.find((x) => x.columnName === col.name);
    // Defense in depth alongside CommonListManager.tsx's own normalization:
    // a bad width here (null, NaN, "") reaching react-data-grid's own column
    // layout math is what crashes it with "Cannot read properties of null
    // (reading 'width')" - fall back to the column's own declared default
    // rather than ever handing it something invalid.
    const hasValidWidth = info && isValidColumnWidth(info.width);

    return {
      ...col,
      key: col.name,
      renderCell: ({ column, row }) => {
        if (column.key === "uniqueId") {
          let loc = uniqueIdHrefHandler
            ? uniqueIdHrefHandler(row.uniqueId)
            : "";

          if (loc.startsWith(".")) {
            loc = resolvePath(currentRouterPath, loc);
          }

          return (
            <div style={{ position: "relative" }}>
              <Link
                href={uniqueIdHrefHandler && uniqueIdHrefHandler(row.uniqueId)}
              >
                {row.uniqueId}
              </Link>
              <div className="cell-actions">
                <CopyCell value={row.uniqueId} />
                <OpenInNewRouter value={loc} />
              </div>
            </div>
          );
        }

        if ((column as any).getCellValue) {
          return <>{(column as any).getCellValue(row)}</>;
        }
        return <span>{get(row, column.key as string)}</span>;
      },
      width: hasValidWidth ? info.width : col.width,
      name: col.title,
      resizable: true,
      sortable: col.sortable,
      renderHeaderCell: (p) => (
        <FilterRenderer<any>
          {...p}
          selectable={true}
          sortable={col.sortable}
          filterable={col.filterable}
          filterType={col.filterType}
          udf={udf}
        />
      ),
    };
  }) as ColumnOrColumnGroup<any, unknown>[];
};
