import type { QueryClient, UseQueryOptions } from "@tanstack/react-query";

/**
 * Shape of the options `useAsQuery`/`FormSelect`'s `querySource` accept for a
 * single query - `query.startIndex`/`itemsPerPage` in particular. Moved from
 * the old `sdk/core/react-tools.tsx`, trimmed to what's actually used: the
 * `Query`/`RemoteRequestOption`/`ExecApi` types it previously referenced came
 * from a `./http-tools` module that doesn't exist in this checkout (a
 * type-only import, silently erased at build time and never actually
 * resolved), so they're `any` here rather than reconstructed.
 */
export interface UseRemoteQuery {
  query?: any;
  queryClient?: QueryClient;
  execFnOverride?: any;
  queryOptions?: UseQueryOptions<any>;
  unauthorized?: boolean;
  UseRemoteQuery?: (options: any) => any;
  optionFn?: (data: any) => any;
}
