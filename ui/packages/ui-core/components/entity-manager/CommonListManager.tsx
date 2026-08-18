import {
  DataTypeProvider,
  type Filter,
  type Sorting,
  type TableColumnWidthInfo,
} from "@devexpress/dx-react-grid";
import { useQueryClient } from "@tanstack/react-query";
import { useEffect, useRef, useState } from "react";
import { useTableViewSizingGetActionQuery } from "@fireback/ui-core/sdk/interfacetools/TableViewSizingGetAction";
import { useTableViewSizingUpdateAction } from "@fireback/ui-core/sdk/interfacetools/TableViewSizingUpdateAction";
import { useDatatableFiltering } from "../../hooks/useDatatableFiltering";
import { type QueryArchiveColumn } from "../../types/QueryArchiveColumn";
import { PaginateTable } from "../common-data-table/PaginateTable";
import {
  normalizeColumnWidths,
  parseJsonSafely,
} from "../common-data-table/PaginateUtils";
import { useReindexedContent } from "../common-data-table/useReindex";
import Link from "../link/Link";
import { type CardComponentType, FlatListMode } from "./FlatListMode";
import { MapListMode } from "./MapListMode";

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
  const { view } = useViewMode();
  const queryClient = useQueryClient();

  // Bug fix: none of the generated use*BrowseActionQuery hooks actually set a
  // static .UKEY (it's referenced here but never assigned anywhere in the sdk
  // codegen), so queryHook.UKEY was always undefined - every single list table in
  // the app (workspaces, roles, email providers, ...) ended up reading and writing
  // the *same* server-side TableViewSizing row (uniqueId "undefined") and the same
  // "table_undefined" localStorage key. Resizing columns on one table (with N
  // columns) then corrupted every other table's saved sizes array, which
  // react-data-grid's internal useColumnWidths doesn't expect - a length/shape
  // mismatch there is what threw "Cannot read properties of null (reading
  // 'width')" on completely unrelated screens (e.g. email-providers, after having
  // resized a column somewhere else). queryHook.name is stable and unique per
  // generated hook (arrow functions assigned to a named export get that name per
  // the JS spec) even though .UKEY never gets set, so it's a safe fallback -
  // falls back further to a column-name signature only if even that's somehow
  // empty (e.g. an inline/anonymous queryHook).
  const tableKey =
    id ??
    queryHook.UKEY ??
    queryHook.name ??
    columns.map((t) => t.name).join(",");

  const query = useTableViewSizingGetActionQuery({
    params: { uniqueId: tableKey },
  });

  const [columnSizes, setColumnSizes] = useState<any>(
    columns.map((t) => ({ columnName: t.name, width: t.width })),
  );

  const tableSizingSizes = (query.data as any)?.data?.item?.sizes;

  // A 404 (no sizing saved for this tableKey yet - the normal state for any
  // table nobody has resized) - or any other non-2xx - never throws here:
  // GResponse.inject only ever populates data.item from a real body.data.item,
  // so an error-shaped `{"error": ...}` response just leaves it null and
  // tableSizingSizes falls straight through to undefined below. What isn't
  // safe on its own is trusting the *shape* of whatever string does come back
  // (from the server, or from localStorage - itself just as untyped) - a
  // single malformed entry used to reach react-data-grid's own column-width
  // computation and crash it with "Cannot read properties of null (reading
  // 'width')", taking the whole table down. parseJsonSafely + normalizeColumnWidths
  // (see PaginateUtils.tsx) is the one gate everything has to clear before it's
  // ever trusted; anything that doesn't parse, isn't shaped right, or doesn't
  // even describe one of this table's own columns is treated exactly like "no
  // saved sizing" instead of being applied and breaking the table.
  //
  // columns/tableKey are intentionally not in the dependency array - callers
  // like CapabilityList.tsx pass a fresh `columns` array literal every render,
  // so depending on it here would re-run (and re-setColumnSizes with a new
  // array reference) every render.
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
    localStorage.setItem(`table_${tableKey}`, sizes);
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
