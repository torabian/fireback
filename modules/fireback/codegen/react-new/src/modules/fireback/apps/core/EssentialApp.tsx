import { ErrorBoundary } from "react-error-boundary";

import "react-toastify/dist/ReactToastify.css";
import "../../../../modules/fireback/styles/styles.scss";
// You do not have to use the mac-os family theme at all.
// this is the default theme which I use for mac desktop applications
// you could use it as a reference to build your own themes.
// themes are nothing special, rather than wrapping a set of css (scss) on a global name
import "../../../..//modules/fireback/styles/apple-family/styles.scss";

import { AuthProvider } from "../../hooks/authContext";
import { UIStateProvider } from "../../hooks/uiStateContext";

import React, { useContext, useEffect } from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import { Fallback } from "../../components/fallback/Fallback";

import { ToastContainer } from "react-toastify";
import { AppConfig, AppConfigContext } from "../../hooks/appConfigTools";
import { usePureLocale } from "../../hooks/usePureLocale";
import { WithFireback } from "./WithFireback";
import { ApplicationPanels } from "./ApplicationPanels";
import { AppTree } from "./AppTree";

export function EssentialApp({
  ApplicationRoutes,
  WithSdk,
  mockServer,
  apiPrefix,
}: {
  mockServer: any;
  ApplicationRoutes: any;
  WithSdk: any;
  apiPrefix?: string;
}) {
  const [queryClient] = React.useState(() => new QueryClient());
  const { config } = useContext(AppConfigContext);

  useEffect(() => {
    if ("serviceWorker" in navigator && "PushManager" in window) {
      navigator.serviceWorker.register("sw.js").then((reg) => {});
    }
  }, []);

  return (
    <QueryClientProvider client={queryClient}>
      <UIStateProvider>
        <AppTree
          config={config}
          ApplicationRoutes={ApplicationRoutes}
          WithSdk={WithSdk}
          mockServer={mockServer}
          apiPrefix={apiPrefix}
          queryClient={queryClient}
        />
      </UIStateProvider>
    </QueryClientProvider>
  );
}

export default EssentialApp;
