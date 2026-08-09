import { useUiState } from "../../modules/fireback-ui/hooks/uiStateContext";

import { useRef } from "react";

import { Panel } from "react-resizable-panels";
import { ResizeHandle } from "../../modules/fireback-ui/components/layouts/ResizeHandle";
import Sidebar from "../../modules/fireback-ui/components/layouts/Sidebar";
import { AppConfigProvider } from "../../modules/fireback-ui/hooks/appConfigTools";
import { BUILD_VARIABLES } from "../../modules/fireback-ui/hooks/build-variables";
import { detectDeviceType } from "../../modules/fireback-ui/hooks/deviceInformation";

// Percentage width the sidebar opens to when there's no saved preference yet
// (a first-time visitor) - matches the fallback useResizeThreshold(768, ...)
// already resizes to elsewhere in this app.
const DEFAULT_SIDEBAR_SIZE = 20;

const getSize = () => {
  if (detectDeviceType().isMobileView) {
    return 0;
  }

  const savedValue = localStorage.getItem("sidebarState");
  if (savedValue === null) {
    // No preference saved yet - open with a sensible default rather than
    // collapsed. Bug fix: `parseFloat(null-ish) <= 0` used to be true here,
    // which collapsed the sidebar for every first-time visitor - the exact
    // same outcome as toggleSidebar()'s own "explicitly collapsed" sentinel
    // (saved as the string "-1"), even though nothing had actually been
    // saved at all yet.
    return DEFAULT_SIDEBAR_SIZE;
  }

  const m = parseFloat(savedValue);

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
