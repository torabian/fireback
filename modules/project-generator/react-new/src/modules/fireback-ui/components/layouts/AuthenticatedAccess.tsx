import { RemoteQueryContext } from "../../../sdk/core/react-tools";
import { useContext } from "react";
import Link from "../link/Link";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";

export function AuthenticatedAccess({
  children,
}: {
  children: React.ReactNode;
}) {
  const { isAuthenticated } = useContext(RemoteQueryContext);
  const s = useS(strings);

  if (!isAuthenticated) {
    return (
      <div className="basic-error-box">
        <div>{s.authenticatedOnly}</div>
      </div>
    );
  }

  return <>{children}</>;
}
