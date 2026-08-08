import "react-toastify/dist/ReactToastify.css";
import "../../modules/styles/styles.scss";
// You do not have to use the mac-os family theme at all.
// this is the default theme which I use for mac desktop applications
// you could use it as a reference to build your own themes.
// themes are nothing special, rather than wrapping a set of css (scss) on a global name
import "../../modules/styles/apple-family/styles.scss";

import {
  QueryClient,
  QueryClient as QueryClient2,
  QueryClientProvider,
  QueryClientProvider as QueryClientProvider2,
} from "@tanstack/react-query";
import React, { useContext, useEffect } from "react";

import { ErrorBoundary } from "react-error-boundary";
import { Fallback } from "@/modules/fireback-ui/components/fallback/Fallback";
import { AppConfigContext } from "@/modules/fireback-ui/hooks/appConfigTools";
import { AuthProvider } from "@/modules/fireback-ui/hooks/authContext";
import { usePureLocale } from "@/modules/fireback-ui/hooks/usePureLocale";
import { SidebarMultiRouterSetup } from "./ApplicationPanels";
import { WithFireback } from "./WithFireback";
import { WithSelfServiceRoutes } from "./WithSelfServiceRoutes";
import { UIStateProvider } from "@/modules/fireback-ui/hooks/uiStateContext";

export function EssentialApp({
  ApplicationRoutes,
  WithSdk,
  apiPrefix,
}: {
  ApplicationRoutes: any;
  WithSdk: any;
  apiPrefix?: string;
}) {
  const [queryClient] = React.useState(() => new QueryClient());
  const [queryClient2] = React.useState(() => new QueryClient2());
  const { config } = useContext(AppConfigContext);

  useEffect(() => {
    if ("serviceWorker" in navigator && "PushManager" in window) {
      navigator.serviceWorker.register("sw.js").then((reg) => {});
    }
  }, []);

  const { locale } = usePureLocale();

  return (
    <QueryClientProvider client={queryClient}>
      <QueryClientProvider2 client={queryClient2}>
        <UIStateProvider>
          <AuthProvider>
            <ErrorBoundary
              FallbackComponent={Fallback}
              onReset={(details) => {
                // Reset the state of your app so the error doesn't happen again
              }}
            >
              <WithFireback
                config={config}
                prefix={apiPrefix}
                queryClient={queryClient}
                locale={locale}
              >
                <WithSdk
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
      </QueryClientProvider2>
    </QueryClientProvider>
  );
}

export default EssentialApp;
