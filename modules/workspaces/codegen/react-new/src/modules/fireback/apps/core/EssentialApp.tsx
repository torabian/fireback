import { ErrorBoundary } from "react-error-boundary";
import "react-toastify/dist/ReactToastify.css";

import {
  AppConfigContext,
  AppConfigProvider,
} from "@/modules/fireback/hooks/appConfigTools";
import { AuthProvider } from "@/modules/fireback/hooks/authContext";
import {
  UIStateProvider,
  useUiState,
} from "@/modules/fireback/hooks/uiStateContext";

import { Fallback } from "@/modules/fireback/components/fallback/Fallback";
import React, { useContext, useEffect, useRef } from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import { BrowserRouter, HashRouter, MemoryRouter } from "react-router-dom";

import { ActionMenuProvider } from "@/modules/fireback/components/action-menu/ActionMenu";
import { ResizeHandle } from "@/modules/fireback/components/layouts/ResizeHandle";
import Sidebar from "@/modules/fireback/components/layouts/Sidebar";
import {
  ModalManager,
  ModalProvider,
} from "@/modules/fireback/components/modal/Modal";
import { ReactiveSearchProvider } from "@/modules/fireback/components/reactive-search/ReactiveSearchContext";
import "@/styles/globals.scss";
import { Panel, PanelGroup } from "react-resizable-panels";
import { ToastContainer } from "react-toastify";
import { WithFireback } from "./WithFireback";

const useHashRouter = process.env.REACT_APP_USE_HASH_ROUTER === "true";
const Router = useHashRouter ? HashRouter : BrowserRouter;

export function EssentialApp({
  ApplicationRoutes,
  WithSdk,
  mockServer,
}: {
  mockServer: any;
  ApplicationRoutes: any;
  WithSdk: any;
}) {
  const [queryClient] = React.useState(() => new QueryClient());

  return (
    <QueryClientProvider client={queryClient}>
      <UIStateProvider>
        <AppTree
          ApplicationRoutes={ApplicationRoutes}
          WithSdk={WithSdk}
          mockServer={mockServer}
          queryClient={queryClient}
        />
      </UIStateProvider>
    </QueryClientProvider>
  );
}

function AppTree({
  queryClient,
  ApplicationRoutes,
  WithSdk,
  mockServer,
}: {
  queryClient: QueryClient;
  mockServer: any;
  ApplicationRoutes: any;
  WithSdk: any;
}) {
  const { config } = useContext(AppConfigContext);
  const { routers, setSidebarRef, setFocusedRouter } = useUiState();
  const panelRef = useRef(null); // Ref for the left panel

  useEffect(() => {
    setSidebarRef(panelRef.current);
  }, [panelRef.current]);

  return (
    <AuthProvider>
      <ErrorBoundary
        FallbackComponent={Fallback}
        onReset={(details) => {
          // Reset the state of your app so the error doesn't happen again
        }}
      >
        <PanelGroup direction="horizontal" style={{ height: "100vh" }}>
          <Router basename={useHashRouter ? undefined : process.env.PUBLIC_URL}>
            <WithFireback
              mockServer={mockServer}
              config={config}
              queryClient={queryClient}
            >
              <WithSdk
                mockServer={mockServer}
                config={config}
                queryClient={queryClient}
              >
                <Panel
                  style={{
                    position: "relative",
                    overflowY: "hidden",
                    height: "100vh",
                  }}
                  minSize={0}
                  ref={(ref) => (panelRef.current = ref)}
                >
                  <AppConfigProvider
                    initialConfig={{
                      remote: process.env.REACT_APP_REMOTE_SERVICE,
                    }}
                  >
                    <Sidebar miniSize={false} />
                  </AppConfigProvider>

                  <ResizeHandle />
                </Panel>
                <Panel
                  defaultSize={80 / routers.length}
                  minSize={10}
                  onClick={() => {
                    setFocusedRouter("url-router");
                  }}
                  style={{
                    position: "relative",
                    display: "flex",
                    width: "100%",
                  }}
                >
                  {routers.find((x) => x.id === "url-router")?.focused &&
                  routers.length ? (
                    <div className="focus-indicator"></div>
                  ) : null}

                  <AppConfigProvider
                    initialConfig={{
                      remote: process.env.REACT_APP_REMOTE_SERVICE,
                    }}
                  >
                    <ReactiveSearchProvider>
                      <ActionMenuProvider>
                        <ModalProvider>
                          <ApplicationRoutes routerId={"url-router"} />
                          <ModalManager />
                        </ModalProvider>
                        <ToastContainer />
                      </ActionMenuProvider>
                    </ReactiveSearchProvider>
                  </AppConfigProvider>
                  <ResizeHandle minimal />
                </Panel>
              </WithSdk>
            </WithFireback>
          </Router>
          {routers.length > 1
            ? routers
                .filter((x) => x.id !== "url-router")
                .map((router) => {
                  return (
                    <MemoryRouter>
                      <Panel
                        defaultSize={80 / routers.length}
                        minSize={10}
                        onClick={() => {
                          setFocusedRouter(router.id);
                        }}
                        style={{
                          position: "relative",
                          display: "flex",
                          width: "100%",
                        }}
                      >
                        {router.focused && routers.length ? (
                          <div className="focus-indicator"></div>
                        ) : null}
                        <AppConfigProvider
                          initialConfig={{
                            remote: process.env.REACT_APP_REMOTE_SERVICE,
                          }}
                        >
                          <WithFireback
                            mockServer={mockServer}
                            config={config}
                            queryClient={queryClient}
                          >
                            <WithSdk
                              mockServer={mockServer}
                              config={config}
                              queryClient={queryClient}
                            >
                              <ReactiveSearchProvider>
                                <ActionMenuProvider>
                                  <QueryClientProvider client={queryClient}>
                                    <ModalProvider>
                                      <ApplicationRoutes routerId={router.id} />
                                      <ModalManager />
                                    </ModalProvider>
                                  </QueryClientProvider>
                                  <ToastContainer />
                                </ActionMenuProvider>
                              </ReactiveSearchProvider>
                            </WithSdk>
                          </WithFireback>
                        </AppConfigProvider>
                        <ResizeHandle minimal />
                      </Panel>
                    </MemoryRouter>
                  );
                })
            : null}
        </PanelGroup>
      </ErrorBoundary>
      <ToastContainer />
    </AuthProvider>
  );
}

export default EssentialApp;
