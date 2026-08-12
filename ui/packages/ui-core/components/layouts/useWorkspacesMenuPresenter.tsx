import { useMemo } from "react";

import { useQueryUserRoleWorkspacesActionQuery } from "@fireback/ui-core/sdk/abac/QueryUserRoleWorkspacesAction";
import { useAuthentication } from "../../auth/AuthenticationContext";
import type { MArray } from "@fireback/js-remote-ctx/common/operators";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";
import type { MenuItem } from "../../types/MenuItem";

export enum MacTagsColor {
  Green = "#00bd00",
  Red = "#ff0313",
  Orange = "#fa7a00",
  Yellow = "#f4b700",
  Blue = "#0072ff",
  Purple = "#ad41d1",
  Grey = "#717176",
}

/**
 * It computes the menu items related to the workspaces, and active role generally
 * used for the sidebar and returns them as MenuItem
 * @param param0
 * @returns
 */
export function useWorkspacesMenuPresenter() {
  const s = useS(strings);
  const { selectedWorkspace, selectWorkspace, session } = useAuthentication();

  const queryUrw = useQueryUserRoleWorkspacesActionQuery({
    enabled: !!session?.token,
  });

  const items = queryUrw.data?.data?.items || [];
  const recomputeKey =
    (items || []).map((item) => item.uniqueId).join("-") +
    "_" +
    selectedWorkspace?.roleId +
    "_" +
    selectedWorkspace?.workspaceId;

  const menus: MenuItem[] = useMemo(() => {
    const workspacesAndRolesList: MenuItem[] = [];
    items.forEach((workspace) => {
      (workspace.roles as MArray<any>).get().forEach((role) => {
        workspacesAndRolesList.push({
          key: `${role.uniqueId}_${workspace.uniqueId}`,
          label: `${workspace.name} (${role.name})`,
          children: [],
          forceActive:
            selectedWorkspace?.roleId === role.uniqueId &&
            selectedWorkspace?.workspaceId === workspace.uniqueId,
          color:
            workspace.uniqueId === "root"
              ? MacTagsColor.Orange
              : MacTagsColor.Green,
          onClick: () => {
            selectWorkspace({
              roleId: role.uniqueId,
              workspaceId: workspace.uniqueId,
            });
          },
        });
      });
    });

    return [
      {
        label: s.workspacesSideTitle,
        children: workspacesAndRolesList.sort((a, b) =>
          a.key < b.key ? -1 : 1,
        ),
      },
    ];
  }, [recomputeKey]);

  return { menus };
}
