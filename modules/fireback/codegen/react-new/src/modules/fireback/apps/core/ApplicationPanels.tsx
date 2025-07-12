import { QueryClient } from "react-query";
import { PanelGroup } from "react-resizable-panels";
import { BrowserRouter, HashRouter, MemoryRouter } from "react-router-dom";
import { TabbarMenu } from "../../components/tabbar-menu/TabbarMenu";
import { useUiState } from "../../hooks/uiStateContext";
import { ApplicationOutlet } from "./ApplicationOutlet";
import {
  PanelRouterWithSidebar,
  PanelRouterWrapper,
} from "./PanelRouterWrapper";
import classNames from "classnames";
import { detectDeviceType } from "../../hooks/deviceInformation";

const useHashRouter = process.env.REACT_APP_USE_HASH_ROUTER === "true";
const Router = useHashRouter ? HashRouter : BrowserRouter;

export function SidebarMultiRouterSetup({
  ApplicationRoutes,
  queryClient,
}: {
  ApplicationRoutes: any;
  queryClient: QueryClient;
}) {
  const { routers } = useUiState();
  const computedRouters = routers.map((item) => {
    return {
      ...item,
      initialEntries: item?.href ? [{ pathname: item?.href }] : undefined,
      Wrapper:
        item.id === "url-router" ? PanelRouterWithSidebar : PanelRouterWrapper,
      Router: item.id === "url-router" ? Router : MemoryRouter,
      showHandle: routers.filter((x) => x.id !== "url-router").length > 0,
    };
  });

  return (
    <PanelGroup
      direction="horizontal"
      className={classNames(
        "application-panels",
        detectDeviceType().isMobileView ? "has-bottom-tab" : undefined
      )}
    >
      {computedRouters.map((router, count) => {
        return (
          <router.Router
            key={router.id}
            future={{ v7_startTransition: true }}
            basename={useHashRouter ? undefined : process.env.PUBLIC_URL}
            initialEntries={router.initialEntries}
          >
            <router.Wrapper showHandle={router.showHandle} routerId={router.id}>
              <ApplicationOutlet
                routerId={router.id}
                ApplicationRoutes={ApplicationRoutes}
                queryClient={queryClient}
              />
            </router.Wrapper>
            {detectDeviceType().isMobileView ? <TabbarMenu /> : undefined}
          </router.Router>
        );
      })}
    </PanelGroup>
  );
}
