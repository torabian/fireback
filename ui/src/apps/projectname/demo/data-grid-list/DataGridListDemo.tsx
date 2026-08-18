import "./DataGridListDemo.css";
import { useRef, useState } from "react";
import { DataGridList } from "@fireback/ui-core/components/data-grid-list/DataGridList";
import { emiTypesColumns, compactColumns } from "./DataGridListDemoColumns";
import { useFakeCursorListQuery } from "./useFakeCursorListQuery";

// Three examples of DataGridList (ui-core/components/data-grid-list) - the
// react-data-grid-only successor to CommonListManager
// (ui-core/components/entity-manager/CommonListManager.tsx), which juggles 3
// render modes (datatable/card/map). DataGridList only ever renders the grid,
// paged by cursor instead of classic offset pagination, filtered by a debounced
// JsonLogic tree per column instead of qs-encoded filters, and scrolled
// infinitely via IntersectionObserver instead of a "load more" button.
//
// None of the data here is real - useFakeCursorListQuery generates a
// deterministic in-memory dataset and fakes network latency (and honors
// AbortSignal, so the Cancel button on the first-load spinner really cancels
// something) - see that file's own doc comment for why.
export function DataGridListDemo() {
  return (
    <div>
      <h1>DataGridList</h1>
      <p>
        The planned CommonListManager replacement - only ever the
        react-data-grid table (no card/map modes), built around cursor-based
        infinite scroll (<code>id(482)+sort(createdAt.desc)</code>, see
        ui-core/components/data-grid-list/hooks/useCursor.ts) and per-column
        JsonLogic filters (ui-core/components/data-grid-list/jsonLogic.ts -
        the same <code>{"{"}"and": [...]{"}"}</code> shape
        modules/fireback/jsonsql/JsonLogicToSql.go parses server-side).
      </p>

      <div className="mt-5 mb-5">
        <EmiTypesExample />
      </div>

      <div className="mt-5 mb-5">
        <BoxedExample />
      </div>

      <div className="mt-5 mb-5">
        <SelectionDeleteExample />
      </div>
    </div>
  );
}

// Example 1 - every emi field type (string/int/int64/float64/bool/complex/
// slice/enum/one/collection/object/any - see fakeRows.ts) as its own column, so
// this is the page to check when wiring up a new column type. Page-level
// scroll: DataGridList observes the browser viewport since no containerRef is
// given (see useInfiniteScroll.ts).
function EmiTypesExample() {
  return (
    <div>
      <h2>1. Every emi field type as a column</h2>
      <p>
        5,000 fake rows, scrolled on the page itself. Try filtering "title" or
        "status", and sorting a couple of columns - the request panel below
        the grid isn't shown here, but you can see the debounced JsonLogic
        query take shape in the Network tab (search param on the fake fetch is
        skipped since this is client-side, but a real queryHook would forward
        it exactly like this).
      </p>
      <DataGridList
        id="demo-emi-types"
        columns={emiTypesColumns}
        queryHook={(params) => useFakeCursorListQuery(5000, params)}
        uniqueIdHrefHandler={() => ""}
        selectable={false}
      />
    </div>
  );
}

// Example 2 - the same infinite-scroll/filter/sort machinery, but boxed inside
// a fixed-height, overflow:auto div. containerRef is what switches
// useInfiniteScroll's IntersectionObserver root from the viewport to this
// element - nothing else about DataGridList changes.
function BoxedExample() {
  const containerRef = useRef<HTMLDivElement | null>(null);

  return (
    <div>
      <h2>2. Contained in a scrollable box</h2>
      <p>
        Same dataset and columns (a leaner subset), but DataGridList only ever
        scrolls within this 420px box - the rest of the page stays put.
      </p>
      <div ref={containerRef} className="data-grid-list-demo__boxed">
        <DataGridList
          id="demo-boxed"
          columns={compactColumns}
          queryHook={(params) => useFakeCursorListQuery(5000, params)}
          uniqueIdHrefHandler={() => ""}
          selectable={false}
          containerRef={containerRef}
        />
      </div>
    </div>
  );
}

// Example 3 - selection + delete, ported from CommonListManager/
// useDatatableFiltering (see hooks/useRowSelectionDelete.ts): select a few
// rows, use the action menu (or Delete key) to remove them. deleteHook here is
// a fake mutation that just remembers "deleted" ids in React state and hands
// them to useFakeCursorListQuery to filter out - a real deleteHook is a
// generated use*AwareDeleteAction mutation hook instead, same as
// CommonListManager already expects.
function SelectionDeleteExample() {
  const [deletedIds, setDeletedIds] = useState<Set<string>>(new Set());

  const fakeDeleteHook = () => ({
    mutateAsync: async (body: string) => {
      const { uniqueIds } = JSON.parse(body) as { uniqueIds: string[] };
      setDeletedIds((prev) => new Set([...prev, ...uniqueIds]));
      return { data: { item: {} } };
    },
  });

  return (
    <div>
      <h2>3. Selection + delete</h2>
      <p>
        Select a few rows (checkbox column) - a "Delete" action appears in the
        action menu, and the Delete key works too, same as
        CommonListManager/useDatatableFiltering today. {deletedIds.size} row
        {deletedIds.size === 1 ? "" : "s"} deleted so far.
      </p>
      <DataGridList
        id="demo-delete"
        columns={compactColumns}
        queryHook={(params) => useFakeCursorListQuery(500, params, deletedIds)}
        uniqueIdHrefHandler={() => ""}
        deleteHook={fakeDeleteHook}
        onRecordsDeleted={(ids) =>
          setDeletedIds((prev) => new Set([...prev, ...ids]))
        }
        height={420}
      />
    </div>
  );
}

export default DataGridListDemo;
