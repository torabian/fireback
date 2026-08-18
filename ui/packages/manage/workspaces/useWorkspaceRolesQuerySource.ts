import { useQuery } from "@tanstack/react-query";
import { useFetchxContext } from "@fireback/js-remote-ctx/react/useFetchx";
import {
  QueryWorkspaceRolesAction,
  QueryWorkspaceRolesActionReq,
} from "@fireback/manage/sdk/abac/QueryWorkspaceRolesAction";
import { type UseRemoteQuery } from "@fireback/ui-core/types/remoteQuery";

// A role option, as returned by QueryWorkspaceRolesAction's GResponse envelope - kept
// local/minimal rather than importing RoleDto, since that class's constructor expects the
// full RoleEntity shape and this endpoint only ever returns a few fields of it.
export interface WorkspaceRoleOption {
  uniqueId: string;
  name: string;
  capabilitiesListId?: string[];
}

// Same {query, items, keyExtractor} adapter shape as useWorkspacesQuerySource, but for a
// *specific* workspace's roles - used by WorkspaceAddUserDrawer's (and any future
// workspace-role-changing UI's) role picker, so it only ever offers roles that actually
// belong to the workspace being acted on.
//
// This intentionally does NOT call the generic useRoleBrowseActionQuery with a
// "workspace-id" header override the way an earlier version of this hook did - that
// looked reasonable (fetchx does merge per-call headers over the global default ones,
// see js-remote-ctx/common/fetchx.ts) but doesn't actually work: RoleBrowseAction's
// security resolves the caller's capabilities against whatever workspace that header
// names, and root (the only caller of this drawer, per AddUserToWorkspaceAction's
// AllowOnRoot) has no real membership in most workspaces - it 401s with
// NotEnoughPermission the moment the header points anywhere but "root". Verified this
// empirically with a live server before writing this fix.
//
// QueryWorkspaceRolesAction (see modules/abac/QueryWorkspaceRolesActionImplementation.go)
// is the real fix: it keeps the caller's own Workspace-Id header pinned to "root" (proving
// they're actually root) while workspaceId - the workspace whose roles are being listed -
// travels as a separate body field instead. It's a POST action (root actions in this repo
// take their target ids in the body, not the URL/query string), so it doesn't come with a
// generated useQuery-shaped hook the way a GET browse action would - this wraps its static
// .Fetch() in a plain useQuery by hand instead.
export const useWorkspaceRolesQuerySource = (
  workspaceId: string,
  params?: UseRemoteQuery,
) => {
  const ctx = useFetchxContext();
  const query = useQuery({
    queryKey: ["QueryWorkspaceRolesAction", workspaceId],
    queryFn: async () => {
      const { response } = await QueryWorkspaceRolesAction.Fetch(
        { body: QueryWorkspaceRolesActionReq.with({ workspaceId }) },
        { ctx },
      );
      return response.result;
    },
    enabled: !!workspaceId,
  });
  const items = ((query.data as any)?.data?.items ??
    []) as WorkspaceRoleOption[];
  return {
    query: query as any,
    items,
    keyExtractor: (item: WorkspaceRoleOption) => item.uniqueId,
  };
};
