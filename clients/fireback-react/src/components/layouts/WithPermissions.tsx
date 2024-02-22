import { userMeetsAccess } from "@/hooks/accessLevels";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
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
  const t = useT();
  const { selectedUrw } = useContext(RemoteQueryContext);

  const meets = useMemo(() => {
    if (process.env.REACT_APP_INACCURATE_MOCK_MODE === "true") {
      return true;
    }
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
        <div className="basic-error-box">{t.lackOfPermission}</div>
      )}
    </>
  );
}
