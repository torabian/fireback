import { ReactNode, useContext } from "react";
import { Panel } from "react-resizable-panels";
import { ResizeHandle } from "../../components/layouts/ResizeHandle";
import { useUiState } from "../../hooks/uiStateContext";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { SidebarPanel } from "./SidebarPanel";

export const PanelRouterWithSidebar = ({
  routerId,
  children,
  showHandle,
}: {
  children: ReactNode;
  showHandle: boolean;
  routerId: string;
}) => {
  return (
    <>
      <SidebarPanel />
      <PanelRouterWrapper showHandle={showHandle} routerId={routerId}>
        {children}
      </PanelRouterWrapper>
    </>
  );
};

export const PanelRouterWrapper = ({
  showHandle,
  routerId,
  children,
}: {
  showHandle: boolean;
  children: ReactNode;
  routerId: string;
}) => {
  const { routers, setFocusedRouter } = useUiState();
  const { session } = useContext(RemoteQueryContext);

  return (
    <Panel
      order={2}
      defaultSize={!session ? 100 : 80 / routers.length}
      minSize={10}
      onClick={() => {
        setFocusedRouter(routerId);
      }}
      style={{
        position: "relative",
        display: "flex",
        width: "100%",
      }}
    >
      {routers.find((x) => x.id === routerId)?.focused && routers.length ? (
        <div className="focus-indicator"></div>
      ) : null}

      {children}

      {showHandle ? <ResizeHandle minimal /> : null}
    </Panel>
  );
};
