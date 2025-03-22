import "react-data-grid/lib/styles.css";

import { ColumnOrColumnGroup, DataGrid } from "react-data-grid";
import { UseQueryResult } from "react-query";
import { DatatableColumn } from "../../definitions/definitions";
import { useT } from "../../hooks/useT";
import { get } from "lodash";
import { useEffect, useMemo } from "react";
import { useReindexedContent } from "./useReindex";

interface TableColumnWidthInfo {
  /** A column name. */
  columnName: string;
  /** A column width. */
  width: number | string;
}

const castColumns = (
  columns: DatatableColumn[]
): ColumnOrColumnGroup<any, unknown>[] => {
  return columns.map((col) => {
    return {
      ...col,
      key: col.name,
      renderCell: ({ column, row }) => {
        return <span>{get(row, column.name as string)}</span>;
      },
      title: col.title,
      resizable: true,
    };
  }) as ColumnOrColumnGroup<any, unknown>[];
};

export function PaginateTable({
  rows,
  columns,
  withFilters,
  booleanColumns,
  query,
  columnSizes,
  defaultColumnWidths,
  onColumnWidthsChange,
  udf,
  RowDetail,
  permissions,
  withPagination,
  children,
  bulkEditHook,
  inlineInsertHook,
  tableClass,
}: {
  rows: any[];
  columns: DatatableColumn[];
  booleanColumns?: string[];
  permissions?: string[];
  withFilters?: boolean;
  withPagination?: boolean;
  RowDetail?: any;
  tableClass?: string;
  query?: UseQueryResult<any, any>;
  defaultColumnWidths: TableColumnWidthInfo[];
  udf: any;
  children?: React.ReactNode;
  columnSizes?: TableColumnWidthInfo[];
  onColumnWidthsChange?:
    | ((nextColumnWidths: TableColumnWidthInfo[]) => void)
    | undefined;
  bulkEditHook?: any;
  inlineInsertHook?: any;
}) {
  const t = useT();

  const {
    filters,
    setSorting,
    setStartIndex,
    selection,
    setSelection,
    setPageSize,
    onFiltersChange,
  } = udf;

  const cols = useMemo(() => castColumns(columns), [columns]);

  const { indexedData, reindex } = useReindexedContent(udf);

  useEffect(() => {
    const rows: any = query.data?.data?.items || [];

    reindex(rows, query.data?.jsonQuery);
  }, [query.data?.data?.items]);

  console.log("Reindexed:", indexedData.length);

  async function handleScroll(event: React.UIEvent<HTMLDivElement>) {
    if (query.isLoading || !isAtBottom(event)) return;
    console.log("Setting to:", indexedData.length);
    setStartIndex(indexedData.length);
  }

  return (
    <DataGrid
      columns={cols}
      onScroll={handleScroll}
      rows={indexedData}
      rowKeyGetter={(item) => item.uniqueId}
      style={{ height: "calc(100% - 2px)", margin: "1px -14px" }}
      //   style={{ height: "calc(100vh - 90px)", marginTop: "20px" }}
    />
  );
}

function isAtBottom({ currentTarget }: React.UIEvent<HTMLDivElement>): boolean {
  return (
    currentTarget.scrollTop + 100 >=
    currentTarget.scrollHeight - currentTarget.clientHeight
  );
}
