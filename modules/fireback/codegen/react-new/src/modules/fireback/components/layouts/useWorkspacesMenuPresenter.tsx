import { useContext, useMemo } from "react";
import { MacTagsColor, type MenuItem } from "../../definitions/common";
import { useT } from "../../hooks/useT";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { useQueryUserRoleWorkspacesActionQuery } from "../../sdk/modules/abac/QueryUserRoleWorkspaces";

/**
 * It computes the menu items related to the workspaces, and active role generally
 * used for the sidebar and returns them as MenuItem
 * @param param0
 * @returns
 */
export function useWorkspacesMenuPresenter() {
  const t = useT();
  const { selectedUrw, selectUrw } = useContext(RemoteQueryContext);
  const queryUrw = useQueryUserRoleWorkspacesActionQuery({
    cacheTime: 50,
  });

  const items = queryUrw.data?.data?.items || [];
  const recomputeKey =
    items.map((item) => item.uniqueId).join("-") +
    "_" +
    selectedUrw?.roleId +
    "_" +
    selectedUrw?.workspaceId;

  const menus: MenuItem[] = useMemo(() => {
    const workspacesAndRolesList: MenuItem[] = [];
    items.forEach((workspace) => {
      workspace.roles.forEach((role) => {
        workspacesAndRolesList.push({
          key: `${role.uniqueId}_${workspace.uniqueId}`,
          label: `${workspace.name} (${role.name})`,
          children: [],
          forceActive:
            selectedUrw?.roleId === role.uniqueId &&
            selectedUrw?.workspaceId === workspace.uniqueId,
          color:
            workspace.uniqueId === "root"
              ? MacTagsColor.Orange
              : MacTagsColor.Green,
          onClick: () => {
            selectUrw({
              roleId: role.uniqueId,
              workspaceId: workspace.uniqueId,
            } as any);
          },
        });
      });
    });

    return [
      {
        label: t.wokspaces.sidetitle,
        children: workspacesAndRolesList.sort((a, b) =>
          a.key < b.key ? -1 : 1
        ),
      },
    ];
  }, [recomputeKey]);

  return { menus };
}
