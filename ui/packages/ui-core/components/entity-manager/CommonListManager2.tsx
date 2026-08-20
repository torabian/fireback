import "react-data-grid/lib/styles.css";

import { DataTypeProvider, type Filter } from "@devexpress/dx-react-grid";
import { useQueryClient } from "@tanstack/react-query";
import { debounce } from "lodash";
import { useMemo, useRef } from "react";
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
import Link from "../link/Link";
import { type CardComponentType } from "./FlatListMode";
import { useTableSizingManager } from "./useTableSizingManager";

interface ListState {
  udf: ReturnType<typeof useDatatableFiltering>;
}

export const CommonListManager2 = ({
  children,
  columns,
  deleteHook,
  uniqueIdHrefHandler,
  withFilters,
  queryHook,
  onRecordsDeleted,
  selectable,
  id,

  withPreloads,
  queryFilters,
  deep,
  inlineInsertHook,
  bulkEditHook,
  urlMask,
  CardComponent,
}: {
  queryHook: ({ state }: { state: ListState }) => any;
  bulkEditHook?: any;
  inlineInsertHook?: any;
  deleteHook?: any;
  columns: QueryArchiveColumn[] | any;
  id?: string;
  urlMask?: string;
  withPreloads?: string;
  uniqueIdHrefHandler?: (id: string) => string;
  deep?: boolean;
  selectable?: boolean;
  withFilters?: boolean;
  onRecordsDeleted?: ({ queryClient }: { queryClient: any }) => void;
  children?: any;
  queryFilters?: Array<Filter | undefined>;
  CardComponent?: CardComponentType<unknown>;
}) => {
  const queryClient = useQueryClient();
  const { pathname } = useLocation();
  const { dir } = useLocale();

  const { columnSizes, onColumnWidthsChange, defaultColumnWidths } =
    useTableSizingManager({
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
    // PaginateTable now renders straight off q.query.data (its own reindex/
    // indexedData wiring is temporarily disabled - see PaginateTable.tsx),
    // so deleteViaUniqueIds above no longer has anywhere to put the removed
    // row: it only updates indexedData, which nothing reads anymore. Refetch
    // the underlying list query itself so the deleted row actually leaves
    // the table, same pattern FlatListMode's onRefresh already uses.
    q.query.refetch();
  };

  const udf = useDatatableFiltering({
    urlMask: "",
    submitDelete: delHook?.mutateAsync,
    onRecordsDeleted: onRecordsDeleted$,
  });

  const source = queryHook({ state: { udf } });

  const { indexedData, reindex, deleteViaUniqueIds } = useReindexedContent(udf);

  let UniqueIdCellRenderer = ({ value }: any) => (
    <div style={{ position: "relative" }}>
      <Link href={uniqueIdHrefHandler && uniqueIdHrefHandler(value)}>
        {value}
      </Link>
      {/* <CopyCell value={value} />
      <OpenInNewRouter value={value} /> */}
    </div>
  );

  let BooleanTypeProvider = (props: any) => (
    <DataTypeProvider formatterComponent={UniqueIdCellRenderer} {...props} />
  );

  const q = source.query ? source : { query: source };
  const rows: any = q.query.data?.data?.items || [];

  const { setStartIndex, selection, setSelection } = udf;

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
    alert(2);
    setStartIndex(indexedData.length);
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
  void BooleanTypeProvider;
  void children;

  return (
    <>
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
    </>
  );
};

function isAtBottom({ currentTarget }: React.UIEvent<HTMLDivElement>): boolean {
  return (
    currentTarget.scrollTop + 300 >=
    currentTarget.scrollHeight - currentTarget.clientHeight
  );
}
