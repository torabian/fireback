import { ErrorBoundary } from "react-error-boundary";


import "react-toastify/dist/ReactToastify.css";
import "../../../../modules/fireback/styles/styles.scss";
// You do not have to use the mac-os family theme at all.
// this is the default theme which I use for mac desktop applications
// you could use it as a reference to build your own themes.
// themes are nothing special, rather than wrapping a set of css (scss) on a global name
import "../../../..//modules/fireback/styles/apple-family/styles.scss";


import { AuthProvider } from "../../hooks/authContext";
import { UIStateProvider, useUiState } from "../../hooks/uiStateContext";

import { Fallback } from "../../components/fallback/Fallback";
import React, { useContext, useEffect, useRef } from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import { BrowserRouter, HashRouter, MemoryRouter } from "react-router-dom";

import { ActionMenuProvider } from "../../components/action-menu/ActionMenu";
import { ResizeHandle } from "../../components/layouts/ResizeHandle";
import Sidebar from "../../components/layouts/Sidebar";
import { ModalManager, ModalProvider } from "../../components/modal/Modal";
import { ReactiveSearchProvider } from "../../components/reactive-search/ReactiveSearchContext";
import { Panel, PanelGroup } from "react-resizable-panels";
import { ToastContainer } from "react-toastify";
import { WithFireback } from "./WithFireback";
import {
  AppConfigContext,
  AppConfigProvider,
} from "../../hooks/appConfigTools";
import { RemoteQueryContext } from "../../sdk/core/react-tools";

const useHashRouter = process.env.REACT_APP_USE_HASH_ROUTER === "true";
const Router = useHashRouter ? HashRouter : BrowserRouter;

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

  return (
    <QueryClientProvider client={queryClient}>
      <UIStateProvider>
        <AppTree
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

function AppTree({
  queryClient,
  ApplicationRoutes,
  WithSdk,
  mockServer,
  apiPrefix,
}: {
  queryClient: QueryClient;
  mockServer: any;
  ApplicationRoutes: any;
  WithSdk: any;
  apiPrefix?: string;
}) {
  const { config } = useContext(AppConfigContext);
  const { routers, setSidebarRef, setFocusedRouter } = useUiState();

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
              prefix={apiPrefix}
              queryClient={queryClient}
            >
              <WithSdk
                mockServer={mockServer}
                prefix={apiPrefix}
                config={config}
                queryClient={queryClient}
              >
                <SidebarPanel />
                <GeneralPanel ApplicationRoutes={ApplicationRoutes} />
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

const SidebarPanel = () => {
  const { routers, setSidebarRef, setFocusedRouter } = useUiState();
  const panelRef = useRef(null); // Ref for the left panel
  const { session } = useContext(RemoteQueryContext);

  useEffect(() => {
    setSidebarRef(panelRef.current);
  }, [panelRef.current]);

  if (!session) {
    return null;
  }

  return (
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
  );
};

const GeneralPanel = ({ ApplicationRoutes }: { ApplicationRoutes: any }) => {
  const { routers, setSidebarRef, setFocusedRouter } = useUiState();
  const { session } = useContext(RemoteQueryContext);

  return (
    <Panel
      defaultSize={!session ? 100 : 80 / routers.length}
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
      {routers.find((x) => x.id === "url-router")?.focused && routers.length ? (
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
  );
};
export default EssentialApp;
