import { DataTypeProvider, type Filter } from "@devexpress/dx-react-grid";
import { useQueryClient } from "@tanstack/react-query";
import { useDatatableFiltering } from "../../hooks/useDatatableFiltering";
import { type QueryArchiveColumn } from "../../types/QueryArchiveColumn";
import { PaginateTable2 } from "../common-data-table/PaginateTable2";
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

  return (
    <>
      <PaginateTable2
        udf={udf}
        selectable={selectable}
        bulkEditHook={bulkEditHook}
        reindex={reindex}
        indexedData={indexedData}
        uniqueIdHrefHandler={uniqueIdHrefHandler}
        onColumnWidthsChange={onColumnWidthsChange}
        columns={columns}
        columnSizes={columnSizes}
        inlineInsertHook={inlineInsertHook}
        rows={rows}
        defaultColumnWidths={defaultColumnWidths as any}
        query={q.query}
        booleanColumns={["uniqueId"]}
        withFilters={withFilters}
      >
        <BooleanTypeProvider for={["uniqueId"]} />

        {children}
      </PaginateTable2>
    </>
  );
};

function isAtBottom({ currentTarget }: React.UIEvent<HTMLDivElement>): boolean {
  return (
    currentTarget.scrollTop + 300 >=
    currentTarget.scrollHeight - currentTarget.clientHeight
  );
}
