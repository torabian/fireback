import { dataMenuToMenu } from "@/fireback/components/layouts/Sidebar";
import { MenuItem } from "@/fireback/definitions/common";
import { useGetCteAppMenus } from "@/sdk/fireback/modules/workspaces/useGetCteAppMenus";
import { useContext, useEffect } from "react";
import { useQueryClient } from "react-query";
import { useLocale } from "./useLocale";
import { RemoteQueryContext } from "@/sdk/fireback/core/react-tools";
import { userMeetsAccess2 } from "./accessLevels";

/**
 *
 * @param menuGroup Use it later for getting different menu items for navbar, other places, etc
 */
export function useRemoteMenuResolver(menuGroup: string): MenuItem[] {
  const queryClient = useQueryClient();
  const { selectedUrw } = useContext(RemoteQueryContext) as any;

  const { query } = useGetCteAppMenus({
    queryClient,
    queryOptions: { refetchOnWindowFocus: false },
    query: {
      itemsPerPage: 9999,
    },
  });
  const { locale } = useLocale();
  useEffect(() => {
    query.refetch();
  }, [locale]);

  let result: MenuItem[] = [];

  const visibilityCheck = (permissionKey: string): boolean => {
    if (!permissionKey) {
      return true;
    }
    return userMeetsAccess2(selectedUrw, permissionKey);
  };

  if (query.data?.data?.items && query.data?.data?.items.length) {
    result = query.data?.data?.items
      .map((item) => dataMenuToMenu(item, visibilityCheck))
      .filter(Boolean);
  }

  return result;
}
