import { userMeetsAccess } from "../../hooks/accessLevels";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";
import { RemoteQueryContext } from "../../../fireback/sdk/core/react-tools";
import { useContext, useMemo } from "react";

export function useIsRoot() {
  const { selectedUrw } = useContext(RemoteQueryContext);
  return selectedUrw?.workspaceId === "root";
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
  const { selectedUrw } = useContext(RemoteQueryContext);

  const meets = useMemo(() => {
    if (selectedUrw?.workspaceId !== "root" && onlyRoot) {
      return false;
    }

    if (!permissions || permissions.length === 0) {
      return true;
    }

    return userMeetsAccess(selectedUrw as any, permissions[0]);
  }, [selectedUrw, permissions]);

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
