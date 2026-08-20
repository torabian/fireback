import "react-data-grid/lib/styles.css";

import { useQueryClient } from "@tanstack/react-query";
import { debounce } from "lodash";
import { useEffect, useMemo, useRef } from "react";
import {
  type CalculatedColumn,
  DataGrid,
  type DataGridHandle,
  SelectColumn,
} from "react-data-grid";
import { useLocation } from "react-router-dom";
import { useDatatableFiltering } from "../../hooks/useDatatableFiltering";
import { useLocale } from "../../hooks/useLocale";
import { type QueryArchiveColumn } from "../../types/QueryArchiveColumn";
import { castColumns } from "../common-data-table/PaginateUtils";
import { useReindexedContent } from "../common-data-table/useReindex";
import { useTableSizingManager } from "./useTableSizingManager";

interface ListState {
  udf: ReturnType<typeof useDatatableFiltering>;
}

export const CommonListManager2 = ({
  children,
  columns,
  deleteHook,
  uniqueIdHrefHandler,
  queryHook,
  onRecordsDeleted,
  id,
}: {
  queryHook: ({ state }: { state: ListState }) => any;
  deleteHook?: any;
  columns: QueryArchiveColumn[] | any;
  id?: string;
  uniqueIdHrefHandler?: (id: string) => string;
  onRecordsDeleted?: ({ queryClient }: { queryClient: any }) => void;
  children?: any;
}) => {
  const queryClient = useQueryClient();
  const { pathname } = useLocation();
  const { dir } = useLocale();

  const { columnSizes, onColumnWidthsChange } = useTableSizingManager({
    columns,
    tableId: id,
  });

  const delHook =
    deleteHook &&
    deleteHook({
      queryClient,
    });

  const onRecordsDeleted$ = (items: string[]) => {
    if (onRecordsDeleted) {
      onRecordsDeleted({ queryClient });
    }
    deleteViaUniqueIds(items);
    // deleteViaUniqueIds only trims the locally accumulated indexedData.
    // Refetch the underlying list query too, so a since-scrolled-past page
    // doesn't bring the deleted row back on the next cursor advance, same
    // pattern FlatListMode's onRefresh already uses.
    q.query.refetch();
  };

  const udf = useDatatableFiltering({
    urlMask: "",
    submitDelete: delHook?.mutateAsync,
    onRecordsDeleted: onRecordsDeleted$,
  });

  const source = queryHook({ state: { udf } });

  const { indexedData, reindex, deleteViaUniqueIds } = useReindexedContent(udf);

  const q = source.query ? source : { query: source };

  // Accumulate pages as the cursor advances: reindex() appends the new
  // page's rows while udf.queryHash stays the same (cursor is stripped out
  // of it), and resets to just the new rows whenever the actual filters
  // change. See useReindex.tsx.
  useEffect(() => {
    if (!q.query.data) return;
    reindex(q.query.data?.data?.items || [], udf.queryHash);
  }, [q.query.data]);

  const rows: any = indexedData;

  const { setCursor, selection, setSelection } = udf;

  const cols = useMemo(() => {
    return [
      SelectColumn,
      ...castColumns(
        columns,
        (field, value) => {
          udf.setFilter({ [field]: value });
        },
        udf,
        columnSizes,
        uniqueIdHrefHandler,
        pathname,
      ),
    ];
  }, [columns, columnSizes]);

  const ref = useRef<DataGridHandle>();

  async function handleScroll(event: React.UIEvent<HTMLDivElement>) {
    if (q.query.isLoading || !isAtBottom(event)) return;

    // GResponse.next.cursor is "" once the server has no more rows to give.
    const nextCursor = q.query.data?.data.cursor;
    if (nextCursor) {
      setCursor(nextCursor);
    }
  }

  const onColumnResize = debounce(
    (column: CalculatedColumn<any, unknown>, width: number) => {
      const newSizes = cols.map((col: any) => {
        return {
          columnName: col.key,
          width: col.name === column.name ? width : col.width,
        };
      });

      onColumnWidthsChange(newSizes);
    },
    300,
  );

  // Note: BooleanTypeProvider (dx-react-grid) and `children` are kept
  // defined for parity with the previous PaginateTable2-based structure,
  // but were never actually rendered there either (PaginateTable2 never
  // rendered its `children` prop) - rendering DataTypeProvider standalone
  // (without a dx-react-grid Grid/PluginHost ancestor) throws at runtime.

  void children;

  return (
    <DataGrid
      columns={cols}
      onScroll={handleScroll}
      onColumnResize={onColumnResize}
      direction={dir as any}
      onSelectedRowsChange={(value) => {
        setSelection(Array.from(value));
      }}
      selectedRows={new Set(selection)}
      ref={ref}
      rows={rows}
      rowKeyGetter={(item) => item.uniqueId}
      style={{ height: "calc(100% - 2px)", margin: "1px -14px" }}
    />
  );
};

function isAtBottom({ currentTarget }: React.UIEvent<HTMLDivElement>): boolean {
  return (
    currentTarget.scrollTop + 300 >=
    currentTarget.scrollHeight - currentTarget.clientHeight
  );
}
