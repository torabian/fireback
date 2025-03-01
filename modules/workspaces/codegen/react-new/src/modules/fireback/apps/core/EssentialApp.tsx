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

import React, { useContext, useEffect, useRef, useState } from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import {
  BrowserRouter,
  HashRouter,
  MemoryRouter,
  Navigate,
  Route,
  Routes,
} from "react-router-dom";
import { Fallback } from "../../components/fallback/Fallback";

import { Panel, PanelGroup } from "react-resizable-panels";
import { ToastContainer } from "react-toastify";
import { ActionMenuProvider } from "../../components/action-menu/ActionMenu";
import {
  ForcedAuthenticated,
  useCheckAuthentication,
} from "../../components/layouts/ForcedAuthenticated";
import { ResizeHandle } from "../../components/layouts/ResizeHandle";
import Sidebar from "../../components/layouts/Sidebar";
import { ModalManager, ModalProvider } from "../../components/modal/Modal";
import { ReactiveSearchProvider } from "../../components/reactive-search/ReactiveSearchContext";
import {
  AppConfig,
  AppConfigContext,
  AppConfigProvider,
} from "../../hooks/appConfigTools";
import { useAbacModulePublicRoutes } from "../../modules/AbacModuleRoutes";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { WithFireback } from "./WithFireback";
import { useLocale } from "../../hooks/useLocale";
import { usePureLocale } from "../../hooks/usePureLocale";

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
  const { session } = useContext(RemoteQueryContext);
  const abacModulePublicRoutes = useAbacModulePublicRoutes();
  const { config } = useContext(AppConfigContext);

  return (
    <QueryClientProvider client={queryClient}>
      <UIStateProvider>
        <>
          <AppTree
            config={config}
            ApplicationRoutes={ApplicationRoutes}
            WithSdk={WithSdk}
            mockServer={mockServer}
            apiPrefix={apiPrefix}
            queryClient={queryClient}
          />
        </>
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

function ApplicationPanels({
  ApplicationRoutes,
  queryClient,
}: {
  ApplicationRoutes: any;
  queryClient: QueryClient;
}) {
  const { routers, setFocusedRouter } = useUiState();
  const { session, checked } = useCheckAuthentication();

  const abacModulePublicRoutes = useAbacModulePublicRoutes();

  if (!session && checked) {
    return (
      <Router>
        <Routes>
          <Route path=":locale">{abacModulePublicRoutes}</Route>
          <Route path="*" element={<Navigate to="/en/signin2" replace />} />
        </Routes>
      </Router>
    );
  }

  return (
    <PanelGroup direction="horizontal" style={{ height: "100vh" }}>
      <Router basename={useHashRouter ? undefined : process.env.PUBLIC_URL}>
        <ForcedAuthenticated>
          <SidebarPanel />
          <GeneralPanel
            ApplicationRoutes={ApplicationRoutes}
            showHandle={routers.filter((x) => x.id !== "url-router").length > 0}
          />
        </ForcedAuthenticated>
      </Router>
      {routers.length > 1
        ? routers
            .filter((x) => x.id !== "url-router")
            .map((router, count) => {
              return (
                <MemoryRouter>
                  <Panel
                    order={count + 2}
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
                    </AppConfigProvider>
                    <ResizeHandle minimal />
                  </Panel>
                </MemoryRouter>
              );
            })
        : null}
    </PanelGroup>
  );
}
function detectDeviceType() {
  const userAgent =
    navigator.userAgent || navigator.vendor || (window as any).opera;
  const isTouch = "ontouchstart" in window || navigator.maxTouchPoints > 0;
  const screenWidth = window.innerWidth || document.documentElement.clientWidth;

  const mobileRegex =
    /android|iphone|ipad|ipod|blackberry|windows phone|opera mini|iemobile/i;
  const desktopRegex = /windows|macintosh|linux|x11/i;

  if (mobileRegex.test(userAgent) || (isTouch && screenWidth < 1024)) {
    return "mobile";
  } else if (
    desktopRegex.test(userAgent) ||
    (!isTouch && screenWidth >= 1024)
  ) {
    return "desktop";
  }

  // Fallback: Check screen width
  return screenWidth < 1024 ? "mobile" : "desktop";
}

const SidebarPanel = () => {
  const { routers, setSidebarRef, setFocusedRouter } = useUiState();
  const panelRef = useRef(null); // Ref for the left panel
  const { session } = useContext(RemoteQueryContext);

  useEffect(() => {
    setSidebarRef(panelRef.current);
    if (detectDeviceType() === "mobile") {
      setTimeout(() => {
        panelRef.current?.resize(0);
      }, 0);
    }
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

const GeneralPanel = ({
  ApplicationRoutes,
  showHandle,
}: {
  ApplicationRoutes: any;
  showHandle: boolean;
}) => {
  const { routers, setSidebarRef, setFocusedRouter } = useUiState();
  const { session } = useContext(RemoteQueryContext);

  return (
    <Panel
      order={2}
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
      {showHandle ? <ResizeHandle minimal /> : null}
    </Panel>
  );
};
export default EssentialApp;
