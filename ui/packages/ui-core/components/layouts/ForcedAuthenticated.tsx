import classNames from "classnames";
import { useEffect, useState } from "react";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";
import { useAuthentication } from "@fireback/auth-client";
import Link from "../link/Link";

export function useCheckAuthentication() {
  const { session, checked } = useAuthentication();
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
  const s = useS(strings);
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
          <div>{s.forcedLayout.forcedLayoutGeneralMessage}</div>
          <Link
            className="btn btn-secondary"
            replace
            href={`/selfservice?redirect=${encodeURIComponent(
              window.location.pathname,
            )}`}
          >
            {s.signinInstead}
          </Link>
        </>
      ) : (
        <>
          <span className="anim-loader"></span>
          <div>{s.forcedLayout.checkingSession}</div>
        </>
      )}
    </div>
  );
}
