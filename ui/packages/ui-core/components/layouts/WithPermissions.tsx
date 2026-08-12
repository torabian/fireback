import { userMeetsAccess } from "../../hooks/accessLevels";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";
import { useAuthentication } from "../../auth/AuthenticationContext";
import { useMemo } from "react";

export function useIsRoot() {
  const { selectedWorkspace } = useAuthentication();
  return selectedWorkspace?.workspaceId === "root";
}

export function WithPermissions({
  children,
  permissions,
  onlyRoot,
}: {
  children: React.ReactNode;
  permissions: string[] | undefined;
  onlyRoot?: boolean;
}) {
  const s = useS(strings);
  const { selectedWorkspace } = useAuthentication();

  const meets = useMemo(() => {
    if (selectedWorkspace?.workspaceId !== "root" && onlyRoot) {
      return false;
    }

    if (!permissions || permissions.length === 0) {
      return true;
    }

    return userMeetsAccess(selectedWorkspace as any, permissions[0]);
  }, [selectedWorkspace, permissions]);

  return (
    <>
      {meets ? (
        children
      ) : (
        <div className="basic-error-box">{s.lackOfPermission}</div>
      )}
    </>
  );
}
