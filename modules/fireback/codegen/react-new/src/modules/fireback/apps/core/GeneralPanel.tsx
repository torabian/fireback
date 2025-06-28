import { useUiState } from "../../hooks/uiStateContext";
import { useContext } from "react";
import { Panel } from "react-resizable-panels";
import { ToastContainer } from "react-toastify";
import { ActionMenuProvider } from "../../components/action-menu/ActionMenu";
import { ResizeHandle } from "../../components/layouts/ResizeHandle";
import { ModalManager, ModalProvider } from "../../components/modal/Modal";
import { OverlayProvider } from "../../components/overlay/OverlayProvider";
import { ReactiveSearchProvider } from "../../components/reactive-search/ReactiveSearchContext";
import { AppConfigProvider } from "../../hooks/appConfigTools";
import { RemoteQueryContext } from "../../sdk/core/react-tools";

export const GeneralPanel = ({
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
              <OverlayProvider>
                <ApplicationRoutes routerId={"url-router"} />
                <ModalManager />
              </OverlayProvider>
            </ModalProvider>
            <ToastContainer />
          </ActionMenuProvider>
        </ReactiveSearchProvider>
      </AppConfigProvider>
      {showHandle ? <ResizeHandle minimal /> : null}
    </Panel>
  );
};
