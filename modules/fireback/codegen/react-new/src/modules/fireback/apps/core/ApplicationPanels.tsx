import { useContext } from "react";
import { useUiState } from "../../hooks/uiStateContext";
import { QueryClient, QueryClientProvider } from "react-query";
import {
  BrowserRouter,
  HashRouter,
  MemoryRouter,
  Navigate,
  Route,
  Routes,
} from "react-router-dom";

import { Panel, PanelGroup } from "react-resizable-panels";
import { ToastContainer } from "react-toastify";
import { ActionMenuProvider } from "../../components/action-menu/ActionMenu";
import {
  ForcedAuthenticated,
  useCheckAuthentication,
} from "../../components/layouts/ForcedAuthenticated";
import { ResizeHandle } from "../../components/layouts/ResizeHandle";
import { ModalManager, ModalProvider } from "../../components/modal/Modal";
import { ReactiveSearchProvider } from "../../components/reactive-search/ReactiveSearchContext";
import { TabbarMenu } from "../../components/tabbar-menu/TabbarMenu";
import { AppConfigProvider } from "../../hooks/appConfigTools";
import { useSelfServicePublicRoutes } from "../../modules/selfservice/SelfServiceRoutes";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { useGetUrwQuery } from "../../sdk/modules/abac/useGetUrwQuery";
import { GeneralPanel } from "./GeneralPanel";
import { SidebarPanel } from "./SidebarPanel";

const useHashRouter = process.env.REACT_APP_USE_HASH_ROUTER === "true";
const Router = useHashRouter ? HashRouter : BrowserRouter;

export function ApplicationPanels({
  ApplicationRoutes,
  queryClient,
}: {
  ApplicationRoutes: any;
  queryClient: QueryClient;
}) {
  const { routers, setFocusedRouter } = useUiState();
  const { session, checked } = useCheckAuthentication();
  const { selectedUrw, selectUrw } = useContext(RemoteQueryContext);
  const selfServicePublicRoutes = useSelfServicePublicRoutes();
  const { query: queryWorkspaces } = useGetUrwQuery({
    queryOptions: { cacheTime: 50 },
    query: {},
  });

  if (!session && checked) {
    return (
      <Router future={{ v7_startTransition: true }}>
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
    <>
      <PanelGroup
        direction="horizontal"
        style={{ height: "calc(100vh - 60px)" }}
      >
        <Router
          future={{ v7_startTransition: true }}
          basename={useHashRouter ? undefined : process.env.PUBLIC_URL}
        >
          <ForcedAuthenticated>
            <SidebarPanel />
            <GeneralPanel
              ApplicationRoutes={ApplicationRoutes}
              showHandle={
                routers.filter((x) => x.id !== "url-router").length > 0
              }
            />
            <TabbarMenu />
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
    </>
  );
}
