import { type JsonLogicNode } from "./jsonLogic";

// One page of results, shaped like a single GResponse envelope's "data" object
// (ui/packages/js-remote-ctx/envelopes/google-json-style-guide) - the same shape
// every generated *BrowseAction hook already resolves to, so a real
// use*BrowseActionQuery-backed hook and DataGridListDemo's fake one both fit this
// without any adapter layer in between.
export interface DataGridListPage<T> {
  items: T[];
  next?: { cursor?: string };
  totalItems?: number;
  totalAvailableItems?: number;
}

// What DataGridList expects back from its `queryHook` prop - deliberately shaped
// like @tanstack/react-query's useInfiniteQuery() return value (data.pages[],
// fetchNextPage, hasNextPage, isFetchingNextPage, ...), since that's exactly what
// a real cursor-paged hook should be built on. `data.pages[n].data` is the
// GResponse-shaped page above, matching `query.data?.data?.items` - the same path
// every CommonListManager caller's queryHook result already reads from.
export interface DataGridListQueryResult<T> {
  data?: { pages: Array<{ data: DataGridListPage<T> }> };
  isLoading: boolean;
  isFetchingNextPage: boolean;
  isError?: boolean;
  fetchNextPage: () => void;
  hasNextPage?: boolean;
  refetch?: () => void;
  /**
   * Aborts the in-flight request, when the hook backing this query supports it
   * (e.g. wired up to an AbortController inside its queryFn) - powers the Cancel
   * button LoadingProgress shows during the first fetch. Hooks that don't
   * support cancellation simply omit this and DataGridList won't offer it.
   */
  cancel?: () => void;
}

export interface DataGridListQueryParams {
  cursor?: string;
  sort?: string;
  filter?: JsonLogicNode;
  itemsPerPage: number;
  queryClient?: unknown;
}

export type DataGridListQueryHook<T> = (
  params: DataGridListQueryParams,
) => DataGridListQueryResult<T>;
