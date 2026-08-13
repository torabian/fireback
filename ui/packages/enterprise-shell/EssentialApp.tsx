import "react-toastify/dist/ReactToastify.css";
import "@fireback/styles/styles.css";
// You do not have to use the mac-os family theme at all.
// this is the default theme which I use for mac desktop applications
// you could use it as a reference to build your own themes.
// themes are nothing special, rather than wrapping a set of css (scss) on a global name
import "@fireback/styles/apple-family/styles.css";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import React, { useContext, useEffect } from "react";
import { ToastContainer } from "react-toastify";

import { Fallback } from "@fireback/ui-core/components/fallback/Fallback";
import { AppConfigContext } from "@fireback/ui-core/hooks/appConfigTools";
import { UIStateProvider } from "@fireback/ui-core/hooks/uiStateContext";
import { usePureLocale } from "@fireback/ui-core/hooks/usePureLocale";
import { ErrorBoundary } from "react-error-boundary";
import { SidebarMultiRouterSetup } from "./ApplicationPanels";
import { WithFireback } from "./WithFireback";
import { WithSelfServiceRoutes } from "./WithSelfServiceRoutes";
import { useRtlClass } from "@fireback/ui-core/hooks/useRtlClass";

export function EssentialApp({
  ApplicationRoutes,
  apiPrefix,
}: {
  ApplicationRoutes: any;
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

  useEffect(() => {
    document
      .querySelector("html")
      ?.setAttribute("dir", ["fa", "ar"].includes(locale) ? "rtl" : "ltr");
  }, []);

  return (
    <QueryClientProvider client={queryClient}>
      {/* Bug fix: hooks/toast.ts's Toast() (and every error/success message routed
          through it - httpErrorHanlder, CommonEntityManager's save toast, ...) called
          react-toastify's toast() this whole time, but nothing ever rendered: toast()
          only queues a notification, react-toastify still needs a mounted
          <ToastContainer/> somewhere in the tree to actually portal it onto the page -
          there wasn't one anywhere in either app (manage or self-service, both go
          through this shared EssentialApp). Mounted once, here, covers both. */}
      <ToastContainer />
      <UIStateProvider>
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
            <WithSelfServiceRoutes>
              <SidebarMultiRouterSetup
                queryClient={queryClient}
                ApplicationRoutes={ApplicationRoutes}
              />
            </WithSelfServiceRoutes>
          </WithFireback>
        </ErrorBoundary>
      </UIStateProvider>
    </QueryClientProvider>
  );
}

export default EssentialApp;
