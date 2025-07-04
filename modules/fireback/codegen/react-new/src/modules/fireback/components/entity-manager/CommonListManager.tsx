import { QueryArchiveColumn } from "../../definitions/common";
import { Filter } from "../../definitions/definitions";
import { dxFilterToSqlAlike } from "../../hooks/datatabletools";
import { useDatatableFiltering } from "../../hooks/useDatatableFiltering";
import { useT } from "../../hooks/useT";
import { useGetTableViewSizingByUniqueId } from "../../sdk/modules/abac/useGetTableViewSizingByUniqueId";
import { usePatchTableViewSizing } from "../../sdk/modules/abac/usePatchTableViewSizing";
import {
  DataTypeProvider,
  Sorting,
  TableColumnWidthInfo,
} from "@devexpress/dx-react-grid";
import { FC, useEffect, useMemo, useRef, useState } from "react";
import { useQueryClient } from "react-query";
import { CommonDataTable } from "../common-data-table/CommonDataTable";
import Link from "../link/Link";
import { filtersToJsonQuery } from "./EnttityManagerHelper";
import { CardComponentType, FlatListMode } from "./FlatListMode";
import { MapListMode } from "./MapListMode";
import { CopyCell } from "./CopyCell";
import { OpenInNewRouter } from "./OpenInNewRouter";
import { PaginateTable } from "../common-data-table/PaginateTable";

const media = matchMedia("(max-width: 600px)");

function useViewMode() {
  const matchRef = useRef(media);

  const [view, setView] = useState<"datatable" | "card" | "map">(
    media.matches ? "card" : "datatable"
  );

  useEffect(() => {
    const query = matchRef.current;
    function listener() {
      if (query.matches) {
        setView("card");
      } else {
        setView("datatable");
      }
    }
    query.addEventListener("change", listener);

    return () => query.removeEventListener("change", listener);
  }, []);

  return { view };
}

function castSortToString(sorting?: Array<Sorting>): string {
  if (!sorting) {
    return "";
  }

  return sorting
    .map((item) => {
      let name = item.columnName;
      if (name === "createdFormatted" || name === "updatedFormatted") {
        name = name.replaceAll("Formatted", "");
      }
      return `${name} ${item.direction}`;
    })
    .join(",");
}

export const CommonListManager = ({
  children,
  columns,
  deleteHook,
  uniqueIdHrefHandler,
  withFilters,
  queryHook,
  onRecordsDeleted,
  selectable,
  id,
  RowDetail,
  withPreloads,
  queryFilters,
  deep,
  inlineInsertHook,
  bulkEditHook,
  urlMask,
  CardComponent,
}: {
  queryHook: any;
  RowDetail?: any;
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
  const t = useT();
  const { view } = useViewMode();
  const queryClient = useQueryClient();

  const { query } = useGetTableViewSizingByUniqueId({
    query: { uniqueId: queryHook.UKEY },
  });

  const [columnSizes, setColumnSizes] = useState<any>(
    columns.map((t) => ({ columnName: t.name, width: t.width }))
  );

  useEffect(() => {
    if ((query as any).data?.data?.sizes) {
      setColumnSizes(JSON.parse((query as any).data?.data?.sizes));
    } else {
      const table = localStorage.getItem(`table_${queryHook.UKEY}`);
      if (table) {
        setColumnSizes(JSON.parse(table));
      }
    }
  }, [(query as any).data?.data?.sizes]);

  const { submit: submitTableSizing } = usePatchTableViewSizing({
    queryClient,
  });

  const delHook =
    deleteHook &&
    deleteHook({
      queryClient,
    });

  const udf = useDatatableFiltering({
    urlMask: "",
    submitDelete: delHook?.submit,
    onRecordsDeleted: onRecordsDeleted
      ? () => onRecordsDeleted({ queryClient })
      : undefined,
  });

  const [defaultColumnWidths] = useState(
    columns.map((t) => ({ columnName: t.name, width: t.width }))
  );

  const onColumnWidthsChange = (nextColumnWidths: TableColumnWidthInfo[]) => {
    setColumnSizes(nextColumnWidths);
    const sizes = JSON.stringify(nextColumnWidths);
    submitTableSizing({
      uniqueId: queryHook.UKEY,
      sizes,
    });
    localStorage.setItem(`table_${queryHook.UKEY}`, sizes);
  };

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

  const f = [...(queryFilters || [])];

  const jsonQuery = useMemo(() => filtersToJsonQuery(f as any), [f]);

  const q = queryHook({
    query: {
      deep: deep === undefined ? true : deep,
      ...udf.debouncedFilters,
      withPreloads,
    },
    queryClient: queryClient,
  });

  q.jsonQuery = jsonQuery;

  const rows: any = q.query.data?.data?.items || [];

  return (
    <>
      {view === "map" && (
        <MapListMode
          columns={columns}
          deleteHook={deleteHook}
          uniqueIdHrefHandler={uniqueIdHrefHandler}
          q={q}
          udf={udf}
        />
      )}
      {view === "card" && (
        <FlatListMode
          columns={columns}
          CardComponent={CardComponent}
          jsonQuery={jsonQuery}
          deleteHook={deleteHook}
          uniqueIdHrefHandler={uniqueIdHrefHandler}
          q={q}
          udf={udf}
        />
      )}

      {view === "datatable" && (
        <PaginateTable
          udf={udf}
          selectable={selectable}
          bulkEditHook={bulkEditHook}
          RowDetail={RowDetail}
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
        </PaginateTable>
      )}
    </>
  );
};
