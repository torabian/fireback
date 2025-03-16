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

import React, { useContext, useEffect, useRef } from "react";
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
import { usePureLocale } from "../../hooks/usePureLocale";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { WithFireback } from "./WithFireback";
import { useSelfServicePublicRoutes } from "../../modules/selfservice/SelfServiceRoutes";

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

  const selfServicePublicRoutes = useSelfServicePublicRoutes();

  if (!session && checked) {
    return (
      <Router>
        <Routes>
          <Route path=":locale">{selfServicePublicRoutes}</Route>
          <Route
            path="*"
            element={<Navigate to="/en/selfservice/welcome" replace />}
          />
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
  const { setSidebarRef, persistSidebarSize } = useUiState();
  const panelRef = useRef(null);
  const { session } = useContext(RemoteQueryContext);

  const onRef = (ref) => {
    if (!ref || panelRef.current) {
      return null;
    }

    panelRef.current = ref;
    setSidebarRef(panelRef.current);

    setTimeout(() => {
      if (detectDeviceType() === "mobile") {
        panelRef.current?.resize(0);
      } else {
        const savedValue = localStorage.getItem("sidebarState");
        const m = savedValue !== null ? parseFloat(savedValue) : null;
        const suggestedSize = (180 / window.innerWidth) * 100;
        const stockSize = m !== null && m > suggestedSize ? m : suggestedSize;

        panelRef.current?.resize(stockSize);
      }
    }, 0);
  };

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
      defaultSize={0}
      ref={onRef}
    >
      <AppConfigProvider
        initialConfig={{
          remote: process.env.REACT_APP_REMOTE_SERVICE,
        }}
      >
        <Sidebar miniSize={false} />
      </AppConfigProvider>

      <ResizeHandle
        onDragComplete={() => {
          persistSidebarSize(panelRef.current?.getSize());
        }}
      />
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
