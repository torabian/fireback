import { useUiState } from "../../hooks/uiStateContext";

import { useContext, useRef } from "react";

import { Panel } from "react-resizable-panels";
import { ResizeHandle } from "../../components/layouts/ResizeHandle";
import Sidebar from "../../components/layouts/Sidebar";
import { AppConfigProvider } from "../../hooks/appConfigTools";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { detectDeviceType } from "../../hooks/deviceInformation";

export const SidebarPanel = () => {
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
      if (detectDeviceType().isMobileView) {
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
