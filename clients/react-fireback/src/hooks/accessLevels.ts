import { DisplayDetectionProps } from "@/definitions/common";
import { CapabilityEntity, UserRoleWorkspaceEntity } from "src/sdk/fireback";

export function userMeetsAccess(
  urw: UserRoleWorkspaceEntity,
  perm: string
): boolean {
  let hasPermission = false;

  for (const item of (urw?.role?.capabilities || []) as CapabilityEntity[]) {
    // if (perm === "root/examsessionreview/query")
    //   console.log("Comparing", item, urw, perm);
    if (
      item.uniqueId === perm ||
      item.uniqueId === "root/*" ||
      (item.uniqueId.endsWith("/*") &&
        perm.includes(item.uniqueId.replace("*", "")))
    ) {
      hasPermission = true;
      break;
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
