import { dataMenuToMenu } from "../components/layouts/Sidebar";
import { useCteAppMenusActionQuery } from "@fireback/ui-core/sdk/interfacetools/CteAppMenusAction";
import { AppMenuOptionalDto } from "@fireback/ui-core/sdk/interfacetools/AppMenuOptionalDto";
import { useQueryUserRoleWorkspacesActionQuery } from "@fireback/ui-core/sdk/abac/QueryUserRoleWorkspacesAction";
import { useAuthentication } from "@fireback/auth-client/AuthenticationContext";
import { GResponse } from "@fireback/js-remote-ctx/envelopes";
import { userMeetsAccess2 } from "./accessLevels";
import { useLocale } from "./useLocale";
import type { MenuItem } from "../types/MenuItem";

// AppMenuOptionalDto is Emi-generated from the "appMenu" entity - Emi has no
// self-referencing dto, so it has no "children" field at all, and its constructor
// only copies its own known fields. /cte-app-menus's response nests real children
// (with their own capabilityId, etc.) under each item though - see
// AppMenuActions.CteQuery / AppMenuTreeNode on the backend. AppMenuTreeItem +
// toAppMenuTree below are the client-side mirror of that hand-written backend type,
// recursively preserving "children" instead of losing it to the plain
// AppMenuOptionalDto cast (which is CteAppMenusAction's default creatorFn).
class AppMenuTreeItem extends AppMenuOptionalDto {
  children: AppMenuTreeItem[] = [];
}

function toAppMenuTree(item: unknown): AppMenuTreeItem {
  const dto = new AppMenuTreeItem(item);
  const rawChildren = (item as { children?: unknown[] })?.children ?? [];
  dto.children = rawChildren.map(toAppMenuTree);
  return dto;
}

/**
 *
 * @param menuGroup Use it later for getting different menu items for navbar, other places, etc
 */
export function useRemoteMenuResolver(menuGroup: string): MenuItem[] {
  const { locale } = useLocale();
  const { selectedWorkspace, session } = useAuthentication();
  const queryUrw = useQueryUserRoleWorkspacesActionQuery({
    enabled: !!session?.token,
  });

  const enabled = !queryUrw.isError && queryUrw.isSuccess && !!session?.token;

  const { data } = useCteAppMenusActionQuery({
    enabled,
    creatorFn: toAppMenuTree,
  });

  let result: MenuItem[] = [];

  const visibilityCheck = (permissionKey?: string | null): boolean => {
    if (!permissionKey) {
      return true;
    }

    return userMeetsAccess2(
      selectedWorkspace as any,
      queryUrw.data?.data?.items || [],
      permissionKey,
    );
  };

  if (data instanceof GResponse) {
    // `?.` on `data?.data?.items` only guards `data`/`data.data` being
    // null/undefined - `items` itself can still legitimately come back as a
    // bare `null` (e.g. cte-app-menus has no rows yet, a fresh install with
    // no appmenu seeders run), and calling `.map` directly on that crashed
    // the whole Sidebar (caught by its ErrorBoundary, but still broken).
    result = (data?.data?.items ?? [])
      .map((item) => {
        return dataMenuToMenu(item, visibilityCheck, locale);
      })
      .filter(Boolean) as MenuItem[];
  }

  return result;
}
