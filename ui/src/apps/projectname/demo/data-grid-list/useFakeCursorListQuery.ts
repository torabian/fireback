import { useInfiniteQuery, useQueryClient } from "@tanstack/react-query";
import {
  type DataGridListQueryParams,
  type DataGridListQueryResult,
} from "@fireback/ui-core/components/data-grid-list/types";
import { decodeCursor, encodeCursor } from "@fireback/ui-core/components/data-grid-list/hooks/useCursor";
import { evaluateJsonLogic } from "./fakeJsonLogic";
import { generateFakeRow, type EmiTypesFakeRow } from "./fakeRows";

function delay(ms: number, signal: AbortSignal): Promise<void> {
  return new Promise((resolve, reject) => {
    const timer = setTimeout(resolve, ms);
    signal.addEventListener("abort", () => {
      clearTimeout(timer);
      reject(new DOMException("Aborted", "AbortError"));
    });
  });
}

function getSortValue(row: EmiTypesFakeRow, columnName: string): unknown {
  return (row as any)[columnName];
}

// Fake, client-side-only stand-in for a real cursor-paged use*BrowseActionQuery
// hook - exercises the exact contract DataGridList expects (types.ts's
// DataGridListQueryHook): a GResponse-shaped page per fetch, a cursor carrying
// both the resume id *and* the sort order that was active when it was minted
// (useCursor.ts), a debounced JsonLogic filter tree applied against the rows,
// and an abortable in-flight request wired to LoadingProgress's Cancel button.
//
// totalRows is the size of the whole fake dataset "on the server" - large
// enough (see DataGridListDemo.tsx) that scrolling through it actually needs
// several pages, so infinite scroll has something real to prove out.
export function useFakeCursorListQuery(
  totalRows: number,
  { cursor, sort, filter, itemsPerPage }: DataGridListQueryParams,
  /** Ids "deleted" client-side (see DataGridListDemo.tsx's Example 3) - filtered
   * out of the fake dataset, and part of the query key so a delete actually
   * triggers a refetch. */
  excludeIds?: Set<string>,
): DataGridListQueryResult<EmiTypesFakeRow> {
  const queryClient = useQueryClient();
  const excludeKey = excludeIds ? Array.from(excludeIds).sort().join(",") : "";
  const queryKey = [
    "fake-cursor-list",
    totalRows,
    sort,
    JSON.stringify(filter ?? null),
    itemsPerPage,
    excludeKey,
  ];

  const query = useInfiniteQuery({
    queryKey,
    initialPageParam: cursor,
    queryFn: async ({ pageParam, signal }) => {
      // Simulated latency - long enough that LoadingProgress/the "loading
      // more" bar are actually visible while testing, short enough not to be
      // annoying.
      await delay(500 + Math.random() * 400, signal);

      const resumeFrom = decodeCursor(pageParam as string | undefined);
      const startAfterId = resumeFrom ? parseInt(resumeFrom.id, 10) : 0;

      const [sortColumn, sortDirection] = sort ? sort.split(".") : [undefined, "asc"];

      // The "server": every id from 1..totalRows, deterministically generated,
      // optionally sorted and filtered - not remotely how you'd actually
      // implement this against a real database (a real backend sorts/filters
      // in SQL, not by materializing every row), but it's enough to prove the
      // cursor/filter/sort contract end to end against something that behaves
      // like real paginated data.
      let candidateIds = Array.from({ length: totalRows }, (_, i) => i + 1);
      let rows = candidateIds.map(generateFakeRow);

      if (excludeIds?.size) {
        rows = rows.filter((row) => !excludeIds.has(row.uniqueId));
      }

      if (filter) {
        rows = rows.filter((row) => evaluateJsonLogic(filter, row as any));
      }

      if (sortColumn) {
        rows = [...rows].sort((a, b) => {
          const av = getSortValue(a, sortColumn);
          const bv = getSortValue(b, sortColumn);
          const cmp = av === bv ? 0 : av < bv ? -1 : 1;
          return sortDirection === "desc" ? -cmp : cmp;
        });
      }

      const startIndex = resumeFrom
        ? rows.findIndex((row) => Number(row.uniqueId) === startAfterId) + 1
        : 0;

      const pageRows = rows.slice(startIndex, startIndex + itemsPerPage);
      const nextIndex = startIndex + pageRows.length;
      const hasMore = nextIndex < rows.length;

      return {
        data: {
          items: pageRows,
          totalItems: rows.length,
          next: hasMore
            ? {
                cursor: encodeCursor({
                  id: pageRows[pageRows.length - 1].uniqueId,
                  sort,
                }),
              }
            : undefined,
        },
      };
    },
    getNextPageParam: (lastPage) => lastPage.data.next?.cursor,
  });

  return {
    data: query.data,
    isLoading: query.isLoading,
    isFetchingNextPage: query.isFetchingNextPage,
    isError: query.isError,
    fetchNextPage: () => {
      query.fetchNextPage();
    },
    hasNextPage: query.hasNextPage,
    refetch: () => {
      query.refetch();
    },
    cancel: () => {
      queryClient.cancelQueries({ queryKey });
    },
  };
}
