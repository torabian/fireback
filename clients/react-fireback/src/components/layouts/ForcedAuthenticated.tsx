import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import Link from "../link/Link";
import { useT } from "@/hooks/useT";

export function ForcedAuthenticated({
  byPass,
  children,
}: {
  byPass?: boolean;
  children: React.ReactNode;
}) {
  const { session } = useContext(RemoteQueryContext);
  const t = useT();
  if (process.env.REACT_APP_FORCE_AUTHENTICATION === "true" && !session) {
    return (
      <div className="unauthorized-forced-area">
        <div>{t.forcedLayout.forcedLayoutGeneralMessage}</div>

        <Link
          className="btn btn-secondary"
          replace
          href={`/signin?redirect=${encodeURIComponent(
            window.location.pathname
          )}`}
        >
          {t.signinInstead}
        </Link>
      </div>
    );
  }

  return <>{children}</>;
}
