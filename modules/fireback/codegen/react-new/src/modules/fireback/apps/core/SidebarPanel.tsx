import { useUiState } from "../../hooks/uiStateContext";

import { useRef } from "react";

import { Panel } from "react-resizable-panels";
import { ResizeHandle } from "../../components/layouts/ResizeHandle";
import Sidebar from "../../components/layouts/Sidebar";
import { AppConfigProvider } from "../../hooks/appConfigTools";
import { detectDeviceType } from "../../hooks/deviceInformation";

const getSize = () => {
  if (detectDeviceType().isMobileView) {
    return 0;
  }

  const savedValue = localStorage.getItem("sidebarState");
  const m = savedValue !== null ? parseFloat(savedValue) : null;

  if (m <= 0) {
    return 0;
  }

  return m * 1.3;
};

export const SidebarPanel = () => {
  const { setSidebarRef, persistSidebarSize } = useUiState();
  const panelRef = useRef(null);

  const onRef = (ref) => {
    panelRef.current = ref;
    setSidebarRef(panelRef.current);
  };

  return (
    <Panel
      style={{
        position: "relative",
        overflowY: "hidden",
        height: "100vh",
      }}
      minSize={0}
      defaultSize={getSize()}
      ref={onRef}
    >
      <AppConfigProvider
        initialConfig={{
          remote: process.env.REACT_APP_REMOTE_SERVICE,
        }}
      >
        <Sidebar miniSize={false} />
      </AppConfigProvider>

      {!detectDeviceType().isMobileView && (
        <ResizeHandle
          onDragComplete={() => {
            persistSidebarSize(panelRef.current?.getSize());
          }}
        />
      )}
    </Panel>
  );
};
