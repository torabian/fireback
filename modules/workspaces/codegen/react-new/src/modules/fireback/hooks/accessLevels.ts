import { DisplayDetectionProps } from "../definitions/common";
import { CapabilityEntity } from "../sdk/modules/workspaces/CapabilityEntity";
import { UserRoleWorkspaceDtoKeys } from "../sdk/modules/workspaces/UserRoleWorkspaceDto";
import { UserWorkspaceEntity } from "../sdk/modules/workspaces/UserWorkspaceEntity";
import { QueryUserRoleWorkspacesActionResDto } from "../sdk/modules/workspaces/WorkspacesActionsDto";

export function userMeetsAccess(urw: any, perm: string): boolean {
  let hasPermission = false;

  for (const item of (urw?.role?.capabilities || []) as CapabilityEntity[]) {
    if (
      item.uniqueId === perm ||
      item.uniqueId === "root.*" ||
      (item?.uniqueId?.endsWith(".*") &&
        perm.includes(item.uniqueId.replace("*", "")))
    ) {
      hasPermission = true;
      break;
    }
  }

  return hasPermission;
}

export function userMeetsAccess2(
  state: { roleId: string; workspaceId: string },
  urw: QueryUserRoleWorkspacesActionResDto[],
  perm: string
): boolean {
  let workspaceMeets = false;
  let roleMeets = false;

  if (!state) {
    return false;
  }

  const workspace = urw.find((item) => item.uniqueId === state.workspaceId);

  // If there is no workspace, then there is no chance that user meets any permission there
  if (!workspace) {
    return false;
  }

  for (const item of workspace.capabilities || []) {
    if (new RegExp(item).test(perm)) {
      workspaceMeets = true;
      break;
    }
  }

  const role = (workspace.roles || []).find(
    (role) => role.uniqueId === state.roleId
  );

  // If there is not role, means there is no chance.
  if (!role) {
    return false;
  }

  for (const item of role.capabilities || []) {
    if (new RegExp(item).test(perm)) {
      roleMeets = true;
      break;
    }
  }

  return workspaceMeets && roleMeets;
}

export function onPermissionInRoot(permission: string) {
  return function (props: DisplayDetectionProps) {
    if (props.selectedUrw?.workspaceId !== "root") {
      return false;
    }

    if (props.selectedUrw) {
      return userMeetsAccess(props.selectedUrw, permission);
    }

    return false;
  };
}

export function onPermission(permission: string) {
  return function (props: DisplayDetectionProps) {
    if (props.selectedUrw) {
      return userMeetsAccess(props.selectedUrw, permission);
    }

    return false;
  };
}
