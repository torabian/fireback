import { dataMenuToMenu } from "@/components/layouts/Sidebar";
import { MenuItem } from "@/definitions/common";
import { useGetCteAppMenus } from "@/sdk/fireback/modules/workspaces/useGetCteAppMenus";
import { useEffect } from "react";
import { useQueryClient } from "react-query";
import { useLocale } from "./useLocale";

/**
 *
 * @param menuGroup Use it later for getting different menu items for navbar, other places, etc
 */
export function useRemoteMenuResolver(menuGroup: string): MenuItem[] {
  const queryClient = useQueryClient();

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
  if (query.data?.data?.items && query.data?.data?.items.length) {
    result = query.data?.data?.items.map((item) => dataMenuToMenu(item));
  }

  return result;
}
