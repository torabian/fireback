/**
 * Tools for authentication based on fireback ABAC plugin
 */

import React, { useContext, useEffect, useRef, useState } from "react";
import { useResizeThreshold } from "./useResizeThreshold";
import { uuidv4 } from "./api";
import useResponsiveThresholds from "../components/layouts/useResponsiveThreshold";
import { detectDeviceType } from "./deviceInformation";

interface ActiveRoute {
  id: string;
  focused?: boolean;
  href?: string;
  initialEntries?: any;
}

export interface IUIStateProvider {
  sidebarVisible: boolean;
  threshold: string;
  routers: Array<ActiveRoute>;
  toggleSidebar: () => void;
  setSidebarRef: (ref: any) => void;
  persistSidebarSize: (value: number) => void;
  setFocusedRouter: (id: string) => void;
  updateSidebarSize: (size: number) => void;
  addRouter: (initialRoute?: string) => void;
  sidebarItemSelected: () => void;
  closeCurrentRouter: (id: string) => void;
  collapseLeftPanel: () => void;
  hide: () => void;
  show: () => void;
}

export const UIStateContext = React.createContext<IUIStateProvider>({
  sidebarVisible: false,
  threshold: "desktop",
  routers: [{ id: "url-router" }],
  toggleSidebar() {},
  setSidebarRef(ref) {},
  persistSidebarSize(value) {},
  setFocusedRouter(ref) {},
  closeCurrentRouter() {},
  sidebarItemSelected() {},
  collapseLeftPanel() {},
  addRouter() {},
  updateSidebarSize() {},
  hide() {},
  show() {},
});

export function useUiState() {
  return useContext(UIStateContext);
}

/**
 * This needs to be provided outside of the routers, because it would
 * manage how many routers are in the application
 * @param param0
 * @returns
 */
export function UIStateProvider({ children }: { children: React.ReactNode }) {
  const panelRef = useRef(null); // This is the panel on the sidebar
  const userPreferedWidth = useRef(null); // This is the panel on the sidebar
  const [sidebarVisible, setSidebarVisibility] = useState(false);

  const [routers, setRouters] = useState<Array<ActiveRoute>>([
    { id: "url-router" },
  ]);

  const persistSidebarSize = (newValue: number) => {
    userPreferedWidth.current = newValue;
    localStorage.setItem("sidebarState", newValue.toString());
  };

  useEffect(() => {
    const savedValue = localStorage.getItem("sidebarState");
    const m = savedValue !== null ? parseFloat(savedValue) : null;
    if (m) {
      userPreferedWidth.current = m;
    }
  }, []);

  const autoClose = useRef(false);

  const resize = (valuePrecentage: number) => {
    panelRef.current?.resize(valuePrecentage);
  };

  useResizeThreshold(768, (exceeded) => {
    resize(exceeded ? 0 : 20);
  });

  const addRouter = (initialRoute?: string) => {
    setRouters((routers) => [...routers, { id: uuidv4(), href: initialRoute }]);
  };

  const setFocusedRouter = (id: string) => {
    setRouters((routers) => {
      return routers.map((route) => {
        if (route.id === id) {
          return {
            ...route,
            focused: true,
          };
        }

        return {
          ...route,
          focused: false,
        };
      });
    });
  };

  const collapseLeftPanel = () => {
    if (panelRef.current && autoClose.current) {
      hide();
      autoClose.current = false;
    }
    updateSidebarSize(panelRef.current?.getSize());
  };

  const toggleSidebar = () => {
    const width = panelRef.current?.getSize();

    // Good sidebar size is at least 180px.
    const suggestedSize = (180 / window.innerWidth) * 100;
    let goodSize = suggestedSize;
    if (
      userPreferedWidth.current &&
      userPreferedWidth.current > suggestedSize
    ) {
      goodSize = userPreferedWidth.current;
    }

    // On phone, use maximum space
    if (detectDeviceType().isMobileView) {
      goodSize = 80;
    }

    if (width && width > 0) {
      resize(0);
      localStorage.setItem("sidebarState", "-1".toString());
      setSidebarVisibility(false);
    } else {
      localStorage.setItem("sidebarState", goodSize.toString());
      resize(goodSize);
      setSidebarVisibility(true);
    }
  };

  const setSidebarRef = (ref) => {
    panelRef.current = ref;
  };

  const hide = () => {
    resize(0);
    setSidebarVisibility(false);
  };

  const closeCurrentRouter = (id: string) => {
    setRouters((routers) => {
      return routers.filter((route) => route.id !== id);
    });
  };

  const updateSidebarSize = (size: number) => {
    resize(size);
  };

  const show = () => {
    if (panelRef.current) {
      resize(20);
      setSidebarVisibility(true);
    }
  };

  const handleThresholds = (name: string) => {
    if (name === "closed") {
      autoClose.current = true;
    } else {
      autoClose.current = false;
    }
  };

  const threshold = useResponsiveThresholds(
    ".sidebar-panel",
    [
      { name: "closed", value: 50 },
      { name: "tablet", value: 100 },
      { name: "desktop", value: 150 },
    ],
    handleThresholds,
    handleThresholds
  );

  const sidebarItemSelected = () => {
    if (window.innerWidth < 500) {
      hide();
    }
  };

  return (
    <UIStateContext.Provider
      value={{
        hide,
        sidebarItemSelected,
        addRouter,
        show,
        updateSidebarSize,
        setFocusedRouter,
        setSidebarRef,
        persistSidebarSize,
        closeCurrentRouter,
        threshold,
        collapseLeftPanel,
        routers,
        sidebarVisible,
        toggleSidebar,
      }}
    >
      {children}
    </UIStateContext.Provider>
  );
}
