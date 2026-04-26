import { dataMenuToMenu } from "../components/layouts/Sidebar";
import { type MenuItem } from "../definitions/common";
import { useGetCteAppMenus } from "../sdk/modules/abac/useGetCteAppMenus";
import { useContext, useEffect } from "react";
import { useQueryClient } from "react-query";
import { useLocale } from "./useLocale";
import { RemoteQueryContext } from "../sdk/core/react-tools";
import { userMeetsAccess2 } from "./accessLevels";
import { useQueryUserRoleWorkspacesActionQuery } from "../sdk/modules/abac/QueryUserRoleWorkspaces";

/**
 *
 * @param menuGroup Use it later for getting different menu items for navbar, other places, etc
 */
export function useRemoteMenuResolver(menuGroup: string): MenuItem[] {
  const queryClient = useQueryClient();
  const { selectedUrw, session } = useContext(RemoteQueryContext);
  const queryUrw = useQueryUserRoleWorkspacesActionQuery({
    enabled: !!session?.token,
  });

  const enabled = !queryUrw.isError && queryUrw.isSuccess && !!session?.token;

  const { query } = useGetCteAppMenus({
    queryClient,
    queryOptions: {
      refetchOnWindowFocus: false,
      enabled,
    },
    query: {
      itemsPerPage: 9999,
    },
  });

  const { locale } = useLocale();

  useEffect(() => {
    query.refetch();
  }, [locale]);

  let result: MenuItem[] = [];

  const visibilityCheck = (permissionKey?: string | null): boolean => {
    if (!permissionKey) {
      return true;
    }

    return userMeetsAccess2(
      selectedUrw,
      queryUrw.data?.data?.items || [],
      permissionKey,
    );
  };

  if (query.data?.data?.items && query.data?.data?.items.length) {
    result = query.data?.data?.items
      .map((item) => dataMenuToMenu(item, visibilityCheck))
      .filter(Boolean) as MenuItem[];
  }

  return result;
}
