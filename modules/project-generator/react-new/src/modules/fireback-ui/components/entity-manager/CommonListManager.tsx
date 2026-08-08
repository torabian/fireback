import {
  DataTypeProvider,
  type Sorting,
  type TableColumnWidthInfo,
} from "@devexpress/dx-react-grid";
import { useEffect, useMemo, useRef, useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { type QueryArchiveColumn } from "../../../fireback/definitions/common";
import { type Filter } from "../../../fireback/definitions/definitions";
import { useDatatableFiltering } from "../../hooks/useDatatableFiltering";
import { useT } from "../../hooks/useT";
import { useTableViewSizingGetActionQuery } from "../../../fireback/sdk/interfacetools/TableViewSizingGetAction";
import { useTableViewSizingUpdateAction } from "../../../fireback/sdk/interfacetools/TableViewSizingUpdateAction";
import { PaginateTable } from "../common-data-table/PaginateTable";
import Link from "../link/Link";
import { filtersToJsonQuery } from "./EnttityManagerHelper";
import { type CardComponentType, FlatListMode } from "./FlatListMode";
import { MapListMode } from "./MapListMode";
import { useReindexedContent } from "../common-data-table/useReindex";

// `queryHook` is a *hook function* (e.g. `useGetCapabilities` from the old
// "react-query" v3 generator, or `useCapabilityBrowseActionQuery` from the
// newer "@tanstack/react-query" v5 action generator), not a resolved query
// result — the two generators don't even share the same UseQueryResult
// shape, so the return type is kept loose here and narrowed with `any`
// where it's consumed below. It's called from inside this component and,
// on the older generated hooks, also carries a static `UKEY` string used
// as a cache/storage key (not yet present on the newer action hooks).
type QueryHook = ((args: { query?: any; queryClient?: any }) => any) & {
  UKEY?: string;
};

const media = matchMedia("(max-width: 600px)");

function useViewMode() {
  const matchRef = useRef(media);

  const [view, setView] = useState<"datatable" | "card" | "map">(
    media.matches ? "card" : "datatable",
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
  queryHook: QueryHook;
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

  const query = useTableViewSizingGetActionQuery({
    params: { uniqueId: queryHook.UKEY },
  });

  const [columnSizes, setColumnSizes] = useState<any>(
    columns.map((t) => ({ columnName: t.name, width: t.width })),
  );

  const tableSizingSizes = (query.data as any)?.data?.item?.sizes;

  useEffect(() => {
    if (tableSizingSizes) {
      setColumnSizes(JSON.parse(tableSizingSizes));
    } else {
      const table = localStorage.getItem(`table_${queryHook.UKEY}`);
      if (table) {
        setColumnSizes(JSON.parse(table));
      }
    }
  }, [tableSizingSizes]);

  // tableViewSizing is addressed by a caller-chosen uniqueId (queryHook.UKEY, a
  // per-table, per-user key) - the update action upserts (creates on first save)
  // for that same uniqueId server-side.
  const { mutate: submitTableSizing } = useTableViewSizingUpdateAction({
    params: { uniqueId: queryHook.UKEY },
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
  };

  const udf = useDatatableFiltering({
    urlMask: "",
    submitDelete: delHook?.mutateAsync,
    onRecordsDeleted: onRecordsDeleted$,
  });

  const source = queryHook({
    query: {
      deep: deep === undefined ? true : deep,
      ...udf.debouncedFilters,
      withPreloads,
    },
    queryClient: queryClient,
  });

  const { indexedData, reindex, deleteViaUniqueIds } = useReindexedContent(udf);

  const [defaultColumnWidths] = useState(
    columns.map((t) => ({ columnName: t.name, width: t.width })),
  );

  const onColumnWidthsChange = (nextColumnWidths: TableColumnWidthInfo[]) => {
    setColumnSizes(nextColumnWidths);
    const sizes = JSON.stringify(nextColumnWidths);
    submitTableSizing({ sizes });
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

  const q = source.query ? source : { query: source };
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
        </PaginateTable>
      )}
    </>
  );
};
