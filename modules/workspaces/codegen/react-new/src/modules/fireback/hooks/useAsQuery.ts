import { QueryObserverResult, UseQueryResult } from "react-query";
import { IResponseList } from "../sdk/core/http-tools";
import { UseRemoteQuery } from "../sdk/core/react-tools";
import { jsonQueryFilter } from "./withJsonQuery";

export function createQuerySource<T>(items: T[]): (
  params: UseRemoteQuery & { items: T[] }
) => {
  query: UseQueryResult<IResponseList<T>, any>;
  items: T[];
} {
  return (otherParams) => useAsQuery<T>({ items, ...otherParams });
}

/**
 * @description This function would convert a static array and make it look like
 * a react-query query hook with similar signatures, pagination, etc.
 *
 * The details it can handle are similar to the query language we provide on Fireback,
 * but you could edit it to meet your needs instead.
 */
export function useAsQuery<T>(params: UseRemoteQuery & { items: T[] }): {
  query: UseQueryResult<IResponseList<T>, any>;
  items: T[];
} {
  let itemsPerPage = params.query?.itemsPerPage || 2;
  let startIndex = params.query.startIndex || 0;

  let items: T[] = params.items || [];
  if (params.query?.jsonQuery) {
    items = jsonQueryFilter(items, params.query.jsonQuery);
  }

  items = items.slice(startIndex, startIndex + itemsPerPage);

  const query: UseQueryResult<IResponseList<T>> = {
    data: {
      data: {
        items,
        totalItems: items.length,
        totalAvailableItems: items.length,
      } as any,
    },
    dataUpdatedAt: 0,
    error: null,
    errorUpdateCount: 0,
    errorUpdatedAt: 0,
    failureCount: 0,
    isError: false,
    isFetched: false,
    isFetchedAfterMount: false,
    isFetching: false,
    isIdle: false,
    isLoading: false,
    isLoadingError: false,
    isPlaceholderData: false,
    isPreviousData: false,
    isRefetchError: false,
    isRefetching: false,
    isStale: false,
    // isSuccess: false,
    remove() {
      console.log("Use as query has not implemented this.");
    },
    refetch() {
      console.log("Refetch is not working actually.");
      return Promise.resolve<QueryObserverResult<IResponseList<T>, unknown>>(
        undefined
      );
    },
    isSuccess: true,
    status: "success",
  };

  return {
    query,
    items,
  };
}
