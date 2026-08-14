import { useAuthentication } from "@fireback/auth-client";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";

export function AuthenticatedAccess({
  children,
}: {
  children: React.ReactNode;
}) {
  const { isAuthenticated } = useAuthentication();
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
