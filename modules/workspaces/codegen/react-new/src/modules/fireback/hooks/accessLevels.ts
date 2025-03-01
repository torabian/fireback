import { DisplayDetectionProps } from "../definitions/common";
import { CapabilityEntity } from "../sdk/modules/workspaces/CapabilityEntity";
import { UserWorkspaceEntity } from "../sdk/modules/workspaces/UserWorkspaceEntity";

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
  urw: UserWorkspaceEntity,
  perm: string
): boolean {
  let hasPermission = false;

  for (const item of urw?.workspacePermissions || []) {
    if (new RegExp(item).test(perm)) {
      return true;
    }
  }

  return hasPermission;
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
