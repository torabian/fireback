import "react-data-grid/lib/styles.css";
import "./DataGridList.css";

import { useMemo } from "react";
import { DataGrid, SelectColumn, type Column } from "react-data-grid";
import { get } from "lodash";
import { type DatatableColumn } from "../../types/DatatableColumn";
import { LoadingProgress } from "../loading-progress/LoadingProgress";
import { strings as coreStrings } from "../strings/translations";
import { useS } from "../../hooks/useS";
import Link from "../link/Link";
import { CopyCell } from "../entity-manager/CopyCell";
import { OpenInNewRouter } from "../entity-manager/OpenInNewRouter";
import { DataGridListHeaderCell } from "./DataGridListHeaderCell";
import { ColumnFiltersProvider, useColumnFilters } from "./hooks/useColumnFilters";
import { useInfiniteScroll } from "./hooks/useInfiniteScroll";
import { useRowSelectionDelete } from "./hooks/useRowSelectionDelete";
import { sortingToString } from "./hooks/useCursor";
import { type DataGridListQueryHook } from "./types";

// DataGridList is meant to eventually replace CommonListManager
// (entity-manager/CommonListManager.tsx) - which juggles 3 render modes
// (datatable/card/map, see its `view` state and FlatListMode/MapListMode) - with
// something that is *only* the react-data-grid table, built around
// cursor-based infinite scroll and JsonLogic column filters instead of
// classic-pagination + qs-encoded filters. Every feature lives in its own hook
// file next to this one (hooks/) so any one of them can be swapped or dropped
// without touching the others:
//   - hooks/useCursor.ts          cursor encode/decode ("id(1)+sort(name.asc)")
//   - hooks/useColumnFilters.ts   per-column filter state -> debounced JsonLogic
//   - hooks/useInfiniteScroll.ts  IntersectionObserver-based "load more" trigger
//   - hooks/useRowSelectionDelete.ts   selection + delete, ported from
//                                       useDatatableFiltering's deleteItems/deleteAction
//
// Prop names deliberately mirror CommonListManagerProps (queryHook, columns,
// deleteHook, uniqueIdHrefHandler, selectable, onRecordsDeleted, id, children,
// ...) so swapping the import is close to the whole migration for a typical
// caller. Not carried over: RowDetail/CardComponent/MapListMode's props (there's
// only one render mode now) and inlineInsertHook/bulkEditHook (not implemented
// yet - accepted so passing them doesn't break the migration, but currently
// unused).
export interface DataGridListProps<T = any> {
  queryHook: DataGridListQueryHook<T>;
  columns: DatatableColumn[];
  deleteHook?: (args: { queryClient?: unknown }) => {
    mutateAsync: (body: string, opts: any) => Promise<any>;
  };
  uniqueIdHrefHandler?: (id: string) => string;
  withFilters?: boolean;
  onRecordsDeleted?: (items: string[]) => void;
  selectable?: boolean;
  id?: string;
  itemsPerPage?: number;
  /** Initial cursor to resume from, e.g. from a deep link. */
  initialCursor?: string;
  /**
   * Scroll container to observe for infinite scroll - leave unset when this
   * grid is the only scrollable thing on the page (observes the viewport
   * instead). See hooks/useInfiniteScroll.ts.
   */
  containerRef?: React.RefObject<HTMLElement | null>;
  /** Fixes the grid's own height so it scrolls internally even without containerRef. */
  height?: number | string;
  children?: React.ReactNode;
  // Accepted for CommonListManager prop-compatibility, currently no-ops:
  RowDetail?: any;
  CardComponent?: any;
  inlineInsertHook?: any;
  bulkEditHook?: any;
  urlMask?: string;
  withPreloads?: string;
  deep?: boolean;
  queryFilters?: any[];
}

function resolvePath(base: string, relative: string) {
  const stack = base.split("/").filter(Boolean);
  const parts = relative.split("/");
  parts.forEach((part) => {
    if (part === "..") stack.pop();
    else if (part !== "." && part !== "") stack.push(part);
  });
  return "/" + stack.join("/");
}

export function DataGridList<T extends { uniqueId?: string } = any>({
  queryHook,
  columns,
  deleteHook,
  uniqueIdHrefHandler,
  withFilters = true,
  onRecordsDeleted,
  selectable = true,
  itemsPerPage = 50,
  initialCursor,
  containerRef,
  height,
  children,
}: DataGridListProps<T>) {
  const cs = useS(coreStrings);
  const columnFilters = useColumnFilters();

  const query = queryHook({
    cursor: initialCursor,
    sort: sortingToString(columnFilters.sort),
    filter: columnFilters.hasActiveFilters ? columnFilters.jsonLogic : undefined,
    itemsPerPage,
  });

  const rows = useMemo<T[]>(
    () => query.data?.pages.flatMap((p) => p.data.items) ?? [],
    [query.data],
  );

  const totalItems = query.data?.pages[query.data.pages.length - 1]?.data.totalItems;

  const deleteMutation = deleteHook?.({});
  const rowSelection = useRowSelectionDelete({
    submitDelete: deleteMutation
      ? (body, opts) => deleteMutation.mutateAsync(body, opts)
      : undefined,
    onRecordsDeleted,
  });

  const { sentinelRef } = useInfiniteScroll({
    containerRef,
    enabled: !!query.hasNextPage && !query.isFetchingNextPage && !query.isLoading,
    onLoadMore: () => query.fetchNextPage(),
  });

  const pathname =
    typeof window !== "undefined" ? window.location.pathname : "";

  const cols = useMemo<Column<T>[]>(() => {
    const built = columns.map((col) => ({
      ...col,
      key: col.name,
      name: col.title,
      resizable: true,
      sortable: col.sortable,
      renderCell: ({ column, row }: { column: Column<T>; row: T }) => {
        if (column.key === "uniqueId") {
          const uniqueId = (row as any).uniqueId;
          let loc = uniqueIdHrefHandler ? uniqueIdHrefHandler(uniqueId) : "";
          if (loc.startsWith(".")) {
            loc = resolvePath(pathname, loc);
          }
          return (
            <div style={{ position: "relative" }}>
              <Link href={uniqueIdHrefHandler ? uniqueIdHrefHandler(uniqueId) : ""}>
                {uniqueId}
              </Link>
              <div className="cell-actions">
                <CopyCell value={uniqueId} />
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
      renderHeaderCell: (props: any) => (
        <DataGridListHeaderCell
          {...props}
          filterType={col.filterType}
          filterable={withFilters && col.filterable}
          sortable={col.sortable}
        />
      ),
    })) as Column<T>[];

    return selectable ? [SelectColumn, ...built] : built;
    // Deliberately NOT depending on columnFilters here - see ColumnFiltersContext's
    // doc comment (hooks/useColumnFilters.ts): column *definitions* only need to
    // change when the columns themselves do, and DataGridListHeaderCell reads the
    // live filter/sort state via useColumnFiltersContext() instead of a prop, so
    // rebuilding this (and remounting every header cell) on every keystroke would
    // only lose state for no benefit.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [columns, selectable, withFilters]);

  const isFirstLoad = query.isLoading && rows.length === 0;

  return (
    <ColumnFiltersProvider value={columnFilters}>
      <div
        className="data-grid-list"
        style={
          containerRef
            ? undefined
            : height
              ? { height, position: "relative" }
              : { position: "relative" }
        }
      >
        {children}

        {isFirstLoad ? (
          <LoadingProgress
            message={cs.components.loading}
            loadedCount={0}
            totalCount={totalItems}
            onCancel={query.cancel}
          />
        ) : (
          <>
            <DataGrid
              columns={cols}
              rows={rows}
              rowKeyGetter={(row: any) => row.uniqueId}
              selectedRows={new Set(rowSelection.selection)}
              onSelectedRowsChange={(value) =>
                rowSelection.setSelection(Array.from(value) as string[])
              }
              style={
                height
                  ? { height, blockSize: height }
                  : { height: "100%", minHeight: 400 }
              }
            />
            <div ref={sentinelRef} className="data-grid-list__sentinel" />
            {query.isFetchingNextPage && (
              <div className="data-grid-list__loading-more">
                <LoadingProgress
                  message={cs.components.loadingMore}
                  loadedCount={rows.length}
                  totalCount={totalItems}
                />
              </div>
            )}
          </>
        )}
      </div>
    </ColumnFiltersProvider>
  );
}

export default DataGridList;
