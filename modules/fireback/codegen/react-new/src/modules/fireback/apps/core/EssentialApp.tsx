import "react-toastify/dist/ReactToastify.css";
import "../../../../modules/fireback/styles/styles.scss";
// You do not have to use the mac-os family theme at all.
// this is the default theme which I use for mac desktop applications
// you could use it as a reference to build your own themes.
// themes are nothing special, rather than wrapping a set of css (scss) on a global name
import "../../../..//modules/fireback/styles/apple-family/styles.scss";

import { UIStateProvider } from "../../hooks/uiStateContext";

import React, { useContext, useEffect } from "react";
import { QueryClient, QueryClientProvider } from "react-query";

import { ErrorBoundary } from "react-error-boundary";
import { Fallback } from "../../components/fallback/Fallback";
import { AppConfigContext } from "../../hooks/appConfigTools";
import { AuthProvider } from "../../hooks/authContext";
import { usePureLocale } from "../../hooks/usePureLocale";
import { SidebarMultiRouterSetup } from "./ApplicationPanels";
import { WithFireback } from "./WithFireback";
import { WithSelfServiceRoutes } from "./WithSelfServiceRoutes";

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

  const { locale } = usePureLocale();

  return (
    <QueryClientProvider client={queryClient}>
      <UIStateProvider>
        <AuthProvider>
          <ErrorBoundary
            FallbackComponent={Fallback}
            onReset={(details) => {
              // Reset the state of your app so the error doesn't happen again
            }}
          >
            <WithFireback
              mockServer={mockServer}
              config={config}
              prefix={apiPrefix}
              queryClient={queryClient}
              locale={locale}
            >
              <WithSdk
                mockServer={mockServer}
                prefix={apiPrefix}
                config={config}
                queryClient={queryClient}
              >
                <WithSelfServiceRoutes>
                  <SidebarMultiRouterSetup
                    queryClient={queryClient}
                    ApplicationRoutes={ApplicationRoutes}
                  />
                </WithSelfServiceRoutes>
              </WithSdk>
            </WithFireback>
          </ErrorBoundary>
        </AuthProvider>
      </UIStateProvider>
    </QueryClientProvider>
  );
}

export default EssentialApp;
