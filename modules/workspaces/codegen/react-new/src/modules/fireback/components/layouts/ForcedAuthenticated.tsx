import classNames from "classnames";
import { useContext, useEffect, useState } from "react";
import { useT } from "../../hooks/useT";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import Link from "../link/Link";

export function useCheckAuthentication() {
  const { session, checked } = useContext(RemoteQueryContext);
  const [loadComplete, setLoadComplete] = useState(false);
  const needsAuthentication = checked && !session;
  const [isFading, setFading] = useState(false);

  useEffect(() => {
    if (checked && session) {
      setFading(true);
      setTimeout(() => {
        setLoadComplete(true);
        // make sure the amount is same as the fade.
      }, 500);
    }
  }, [checked, session]);

  return {
    session,
    checked,
    needsAuthentication,
    loadComplete,
    setLoadComplete,
    isFading,
  };
}

export function ForcedAuthenticated({
  byPass,
  children,
}: {
  byPass?: boolean;
  children: React.ReactNode;
}) {
  const t = useT();
  const { loadComplete, needsAuthentication, session, isFading } =
    useCheckAuthentication();

  if (loadComplete && session) {
    return <>{children}</>;
  }
  return (
    <div
      className={classNames("unauthorized-forced-area", {
        "fade-out": isFading,
      })}
    >
      {needsAuthentication ? (
        <>
          <div>{t.forcedLayout.forcedLayoutGeneralMessage}</div>
          <Link
            className="btn btn-secondary"
            replace
            href={`/signin2?redirect=${encodeURIComponent(
              window.location.pathname
            )}`}
          >
            {t.signinInstead}
          </Link>
        </>
      ) : (
        <>
          <span className="anim-loader"></span>
          <div>{t.forcedLayout.checkingSession}</div>
        </>
      )}
    </div>
  );
}
