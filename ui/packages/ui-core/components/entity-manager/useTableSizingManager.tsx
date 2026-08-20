import { useEffect, useState } from "react";
import {
  normalizeColumnWidths,
  parseJsonSafely,
} from "../common-data-table/PaginateUtils";
import { useTableViewSizingGetActionQuery } from "../../sdk/interfacetools/TableViewSizingGetAction";
import type { QueryArchiveColumn } from "../../types/QueryArchiveColumn";
import { useTableViewSizingUpdateAction } from "../../sdk/interfacetools/TableViewSizingUpdateAction";

export interface ColumnSizingInfo {
  columnName: string;
  width: number;
}

export interface UseTableSizingManagerProps {
  columns: QueryArchiveColumn[];
  tableId?: string;
}

export const useTableSizingManager = ({
  columns,
  tableId,
}: UseTableSizingManagerProps) => {
  const tableKey = tableId ?? columns.map((t) => t.name).join(",");

  const query = useTableViewSizingGetActionQuery({
    params: { uniqueId: tableKey },
  });

  const tableSizingSizes = (query.data as any)?.data?.item?.sizes;

  const [columnSizes, setColumnSizes] = useState<any>(
    columns.map((t) => ({ columnName: t.name, width: t.width })),
  );

  const [defaultColumnWidths] = useState(
    columns.map((t) => ({ columnName: t.name, width: t.width })),
  );

  useEffect(() => {
    const fromServer = normalizeColumnWidths(
      parseJsonSafely(tableSizingSizes),
      columns,
    );
    if (fromServer) {
      setColumnSizes(fromServer);
      return;
    }

    const fromStorage = normalizeColumnWidths(
      parseJsonSafely(localStorage.getItem(`table_${tableKey}`)),
      columns,
    );

    if (fromStorage) {
      setColumnSizes(fromStorage);
    }
    // Neither had anything valid - leave columnSizes at its already-safe
    // initial default (each column's own declared width) rather than
    // clearing it out from under an in-progress render.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [tableSizingSizes]);

  // tableViewSizing is addressed by a caller-chosen uniqueId (tableKey, a
  // per-table, per-user key) - the update action upserts (creates on first save)
  // for that same uniqueId server-side.
  const { mutate: submitTableSizing } = useTableViewSizingUpdateAction({
    params: { uniqueId: tableKey },
  });

  const onColumnWidthsChange = (nextColumnWidths: ColumnSizingInfo[]) => {
    console.log(1, nextColumnWidths);
    setColumnSizes(nextColumnWidths);
    const sizes = JSON.stringify(nextColumnWidths);
    submitTableSizing({ sizes });
    localStorage.setItem(`table_${tableKey}`, sizes);
  };

  return { columnSizes, onColumnWidthsChange, defaultColumnWidths };
};
