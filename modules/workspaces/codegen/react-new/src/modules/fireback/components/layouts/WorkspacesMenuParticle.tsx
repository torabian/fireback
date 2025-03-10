import { MacTagsColor, MenuItem } from "../../definitions/common";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { useContext, useEffect } from "react";
import { useQueryClient } from "react-query";
import { MenuParticle } from "./MenuParticle";
import { useT } from "../../hooks/useT";
import { groupBy } from "lodash";
import { useGetUserWorkspaces } from "../../sdk/modules/workspaces/useGetUserWorkspaces";

export function WorkspacesMenuParticle({ onClick }: { onClick: () => void }) {
  const t = useT();
  const { selectedUrw, selectUrw } = useContext(RemoteQueryContext) as any;

  const queryClient = useQueryClient();
  const { query: queryWorkspaces } = useGetUserWorkspaces({
    queryClient,
    query: {},
    queryOptions: {
      cacheTime: 0,
      refetchOnWindowFocus: false,
    },
  });

  useEffect(() => {
    if (queryWorkspaces.data?.data?.items?.length && !selectedUrw) {
      selectUrw(queryWorkspaces.data?.data?.items[0]);
    }
  }, [queryWorkspaces.data?.data?.items]);

  const menus: MenuItem[] = [];

  const groupped = groupBy(
    queryWorkspaces.data?.data?.items || [],
    (t) => t.workspaceId
  );

  const otherMenu: any = {
    label: t.wokspaces.sidetitle,
    children: [],
  };

  for (const workspaceId of Object.keys(groupped)) {
    const urws = groupped[workspaceId];

    if (urws.length > 1) {
      menus.push({
        name:
          urws[0].workspace?.name || urws[0].workspaceId || t.unnamedWorkspace,
        children: urws.map((urw) => {
          return {
            key: urw.uniqueId,
            children: [],

            onClick() {
              if (urw.workspaceId) {
                selectUrw(urw);
              }
            },
            forceActive: selectedUrw?.uniqueId === urw.uniqueId,
            color:
              urw.workspaceId === "root"
                ? MacTagsColor.Orange
                : MacTagsColor.Green,
            label: "" || t.unnamedRole,
          };
        }),
      } as any);
    } else {
      otherMenu.children.push({
        key: urws[0].uniqueId,
        children: [],

        onClick() {
          if (urws[0].workspaceId) {
            selectUrw(urws[0]);
          }
        },
        forceActive: selectedUrw?.uniqueId === urws[0].uniqueId,
        color:
          urws[0].workspaceId === "root"
            ? MacTagsColor.Orange
            : MacTagsColor.Green,
        label:
          urws[0].workspace?.name || urws[0].workspaceId || t.unnamedWorkspace,
      });
    }
  }

  menus.push(otherMenu);
  if (menus.length === 1 && menus[0].children.length === 0) {
    return null;
  }

  return (
    <>
      {menus.map((menu: any, index) => {
        return (
          <MenuParticle
            onClick={onClick}
            key={`${menu.key}_${menu.label}_${index}`}
            menu={menu}
          />
        );
      })}
    </>
  );
}
