import { ErrorBoundary } from "react-error-boundary";

import { AuthProvider } from "../../hooks/authContext";

import { QueryClient } from "react-query";
import { Fallback } from "../../components/fallback/Fallback";

import { ToastContainer } from "react-toastify";
import { AppConfig } from "../../hooks/appConfigTools";
import { usePureLocale } from "../../hooks/usePureLocale";
import { WithFireback } from "./WithFireback";
import { ApplicationPanels } from "./ApplicationPanels";

export function AppTree({
  queryClient,
  ApplicationRoutes,
  WithSdk,
  mockServer,
  apiPrefix,
  config,
}: {
  queryClient: QueryClient;
  mockServer: any;
  ApplicationRoutes: any;
  WithSdk: any;
  apiPrefix?: string;
  config: AppConfig;
}) {
  const { locale } = usePureLocale();
  return (
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
            <ApplicationPanels
              queryClient={queryClient}
              ApplicationRoutes={ApplicationRoutes}
            />
          </WithSdk>
        </WithFireback>
      </ErrorBoundary>
      <ToastContainer />
    </AuthProvider>
  );
}
