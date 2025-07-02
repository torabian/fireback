import "react-data-grid/lib/styles.css";

import { debounce } from "lodash";
import { useEffect, useMemo, useRef } from "react";
import { CalculatedColumn, DataGrid, DataGridHandle } from "react-data-grid";
import { UseQueryResult } from "react-query";
import { DatatableColumn } from "../../definitions/definitions";
import { Udf } from "../../hooks/useDatatableFiltering";
import { useT } from "../../hooks/useT";
import { useReindexedContent } from "./useReindex";
import { castColumns, TableColumnWidthInfo } from "./PaginateUtils";
import { useLocation } from "react-router-dom";

export function PaginateTable({
  columns,
  query,
  columnSizes,
  onColumnWidthsChange,
  udf,
  tableClass,
  uniqueIdHrefHandler,
}: {
  rows: any[];
  columns: DatatableColumn[];
  booleanColumns?: string[];
  permissions?: string[];
  withFilters?: boolean;
  withPagination?: boolean;
  selectable?: boolean;
  uniqueIdHrefHandler?: (id: string) => string;
  RowDetail?: any;
  tableClass?: string;
  query?: UseQueryResult<any, any>;
  defaultColumnWidths: TableColumnWidthInfo[];
  udf: Udf;
  children?: React.ReactNode;
  columnSizes?: TableColumnWidthInfo[];
  onColumnWidthsChange?:
    | ((nextColumnWidths: TableColumnWidthInfo[]) => void)
    | undefined;
  bulkEditHook?: any;
  inlineInsertHook?: any;
}) {
  const t = useT();
  const { pathname } = useLocation();

  const {
    filters,
    setSorting,
    setStartIndex,
    selection,
    setSelection,
    setPageSize,
    onFiltersChange,
  } = udf;

  console.log(100, pathname);

  const cols = useMemo(() => {
    return castColumns(
      columns,
      (field, value) => {
        udf.setFilter({ [field]: value });
      },
      udf,
      columnSizes,
      uniqueIdHrefHandler,
      pathname
    );
  }, [columns, columnSizes]);

  const { indexedData, reindex } = useReindexedContent(udf);
  const ref = useRef<DataGridHandle>();

  useEffect(() => {
    const rows: any = query.data?.data?.items || [];

    reindex(rows, udf.queryHash, () => {
      ref.current.element.scrollTo({ top: 0, left: 0 });
    });
  }, [query.data?.data?.items]);

  async function handleScroll(event: React.UIEvent<HTMLDivElement>) {
    if (query.isLoading || !isAtBottom(event)) return;
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
    300
  );

  return (
    <>
      <DataGrid
        className={tableClass}
        columns={cols}
        onScroll={handleScroll}
        onColumnResize={onColumnResize}
        ref={ref}
        rows={indexedData}
        rowKeyGetter={(item) => item.uniqueId}
        style={{ height: "calc(100% - 2px)", margin: "1px -14px" }}
      />
    </>
  );
}

function isAtBottom({ currentTarget }: React.UIEvent<HTMLDivElement>): boolean {
  return (
    currentTarget.scrollTop + 300 >=
    currentTarget.scrollHeight - currentTarget.clientHeight
  );
}
