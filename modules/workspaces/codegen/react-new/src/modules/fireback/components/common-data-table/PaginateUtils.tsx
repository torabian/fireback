import "react-data-grid/lib/styles.css";

import { get } from "lodash";
import { ColumnOrColumnGroup } from "react-data-grid";
import { DatatableColumn } from "../../definitions/definitions";
import { Udf } from "../../hooks/useDatatableFiltering";
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

export const castColumns = (
  columns: DatatableColumn[],
  setFilter: (field: string, value: any) => void,
  udf: Udf,
  columnSizes: TableColumnWidthInfo[] = [],
  uniqueIdHrefHandler: (id: any) => void
): ColumnOrColumnGroup<any, unknown>[] => {
  return columns.map((col) => {
    const info = columnSizes.find((x) => x.columnName === col.name);

    return {
      ...col,
      key: col.name,
      renderCell: ({ column, row }) => {
        if (column.key === "uniqueId") {
          return (
            <div style={{ position: "relative" }}>
              <Link
                href={uniqueIdHrefHandler && uniqueIdHrefHandler(row.uniqueId)}
              >
                {row.uniqueId}
              </Link>
              <CopyCell value={row.uniqueId} />
              <OpenInNewRouter value={row.uniqueId} />
            </div>
          );
        }

        if ((column as any).getCellValue) {
          return <>{(column as any).getCellValue(row)}</>;
        }
        return <span>{get(row, column.key as string)}</span>;
      },
      width: info ? info.width : col.width,
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
