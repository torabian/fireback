import {
  useUserBrowseActionQuery,
  UserBrowseActionQueryParams,
} from "@fireback/ui-core/sdk/abac/UserBrowseAction";
import { type UserDto } from "@fireback/ui-core/sdk/abac/UserDto";
import { type UseRemoteQuery } from "../types/remoteQuery";

// Adapts useUserBrowseActionQuery (the generated per-action hook, returning a raw
// @tanstack/react-query result whose payload is the {data: {items, ...}} envelope)
// into the {query, items, keyExtractor} shape FormSelect/FormSelectMultiple's
// querySource prop expects - same pattern as useEmailProvidersQuerySource/
// useRolesQuerySource. Used by the "add existing user to workspace" drawer
// (ui/packages/manage/workspaces/WorkspaceAddUserDrawer.tsx) to search across every
// user in the system (UserBrowseAction is root-only and not workspace-scoped, unlike
// roles).
export const useUsersQuerySource = (params?: UseRemoteQuery) => {
  const query = useUserBrowseActionQuery({
    qs: new UserBrowseActionQueryParams({
      itemsPerPage: params?.query?.itemsPerPage ?? 200,
      searchPhrase: params?.query?.searchPhrase || undefined,
    }),
  });
  const items = ((query.data as any)?.data?.items ?? []) as UserDto[];
  return {
    query: query as any,
    items,
    keyExtractor: (item: UserDto) => item.uniqueId,
  };
};
