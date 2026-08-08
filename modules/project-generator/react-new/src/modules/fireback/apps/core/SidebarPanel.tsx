import { useUiState } from "../../../fireback-ui/hooks/uiStateContext";

import { useRef } from "react";

import { Panel } from "react-resizable-panels";
import { ResizeHandle } from "../../../fireback-ui/components/layouts/ResizeHandle";
import Sidebar from "../../../fireback-ui/components/layouts/Sidebar";
import { AppConfigProvider } from "../../../fireback-ui/hooks/appConfigTools";
import { BUILD_VARIABLES } from "../../../fireback-ui/hooks/build-variables";
import { detectDeviceType } from "../../../fireback-ui/hooks/deviceInformation";

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
      defaultSize={getSize()}
      ref={onRef}
    >
      <AppConfigProvider
        initialConfig={{
          remote: BUILD_VARIABLES.REMOTE_SERVICE,
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
