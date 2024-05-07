import { ErrorBoundary } from "react-error-boundary";
import "react-toastify/dist/ReactToastify.css";

import { AuthProvider } from "@/fireback/hooks/authContext";
import { UIStateProvider } from "@/fireback/hooks/uiStateContext";
import {
  AppConfigContext,
  AppConfigProvider,
} from "@/fireback/hooks/appConfigTools";

import { Fallback } from "@/fireback/components/fallback/Fallback";
import React, { useContext, useRef } from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import { BrowserRouter, HashRouter } from "react-router-dom";
import { ApplicationRoutes } from "./ApplicationRoutes";
import { WithFireback } from "../../fireback/apps/core/WithFireback";

import "@/styles/globals.scss";
import { ToastContainer } from "react-toastify";
import { FirebackMockServer } from "./mockServer";
import { ModalManager, ModalProvider } from "@/fireback/components/modal/Modal";
import { ReactiveSearchProvider } from "@/fireback/components/reactive-search/ReactiveSearchContext";
import { ActionMenuProvider } from "@/fireback/components/action-menu/ActionMenu";

const useHashRouter = process.env.REACT_APP_USE_HASH_ROUTER === "true";
const Router = useHashRouter ? HashRouter : BrowserRouter;

function App() {
  const [queryClient] = React.useState(() => new QueryClient());
  const { config } = useContext(AppConfigContext);
  const mockServer = useRef<any>(FirebackMockServer);
  return (
    <AuthProvider>
      <ErrorBoundary
        FallbackComponent={Fallback}
        onReset={(details) => {
          // Reset the state of your app so the error doesn't happen again
        }}
      >
        <Router basename={useHashRouter ? undefined : process.env.PUBLIC_URL}>
          <AppConfigProvider
            initialConfig={{ remote: process.env.REACT_APP_REMOTE_SERVICE }}
          >
            <WithFireback
              mockServer={mockServer}
              config={config}
              queryClient={queryClient}
            >
              <UIStateProvider>
                <ReactiveSearchProvider>
                  <ActionMenuProvider>
                    <QueryClientProvider client={queryClient}>
                      <ModalProvider>
                        <ApplicationRoutes />
                        <ModalManager />
                      </ModalProvider>
                    </QueryClientProvider>
                    <ToastContainer />
                  </ActionMenuProvider>
                </ReactiveSearchProvider>
              </UIStateProvider>
            </WithFireback>
          </AppConfigProvider>
        </Router>
      </ErrorBoundary>
      <ToastContainer />
    </AuthProvider>
  );
}

export default App;
