import { useContext, useEffect, useState } from "react";
import {
  RemoteQueryContext as FirebackContext,
  RemoteQueryProvider as FirebackQueryProvider,
} from "src/sdk/fireback/core/react-tools";
import Link from "@/components/link/Link";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useT } from "@/hooks/useT";
import { source } from "@/helpers/source";
import { AppConfigContext } from "@/hooks/appConfigTools";
import { osResources } from "@/components/mulittarget/multitarget-resource";

export const UserProfileCard = () => {
  const { session, signout } = useContext(RemoteQueryContext);
  const { config } = useContext(AppConfigContext);
  const t = useT();

  return (
    <div className="auth-profile-card with-fade-in auth-wrapper">
      <img src={source(osResources.user)} />
      <h2>
        {(session as any)?.user?.firstName} {(session as any)?.user?.lastName}
      </h2>

      <div>
        <Link
          className="go-to-the-app"
          href={(process.env.REACT_APP_DEFAULT_ROUTE || "").replace(
            "/{locale}",
            ""
          )}
        >
          {t.abac.backToApp}
        </Link>
      </div>
      <button className="btn btn-danger" onClick={signout}>
        {t.abac.signout}
      </button>
    </div>
  );
};

export const UserOsProfileCard = () => {
  const currentUser = "Offline";
  const t = useT();
  const { config } = useContext(AppConfigContext);
  const { options, session } = useContext(FirebackContext);
  const [pingResult, setPingResult] = useState<any>("Ping idle...");

  const pingTest = () => {
    // fetch(`http://10.0.2.2:59731/ping`, {
    fetch(`${config.remote}ping`, {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    })
      .then((response) => response.json())
      .then((response) => setPingResult(response))
      .catch((err) => setPingResult(err));
  };

  useEffect(() => {
    pingTest();
  }, []);

  return (
    <div className="auth-profile-card with-fade-in auth-wrapper">
      <img src={source(osResources.user)} />
      <h2>{t.signup.continueAs.replace("{currentUser}", currentUser)}</h2>

      <div
        className="disclaimer"
        dangerouslySetInnerHTML={{
          __html: t.signup.continueAsHint.replace("{currentUser}", currentUser),
        }}
      ></div>
    </div>
  );
};
