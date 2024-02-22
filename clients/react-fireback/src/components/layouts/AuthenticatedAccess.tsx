import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import Link from "../link/Link";
import { useT } from "@/hooks/useT";

export function AuthenticatedAccess({
  children,
}: {
  children: React.ReactNode;
}) {
  const { isAuthenticated } = useContext(RemoteQueryContext);
  const t = useT();

  if (!isAuthenticated) {
    return (
      <div className="basic-error-box">
        <div>{t.authenticatedOnly}</div>
      </div>
    );
  }

  return <>{children}</>;
}
