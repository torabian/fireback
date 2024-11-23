/**
 * Tools for authentication based on fireback ABAC plugin
 */

import React, { useContext, useEffect, useRef, useState } from "react";
import { useResizeThreshold } from "./useResizeThreshold";
import { uuidv4 } from "./api";
import useResponsiveThresholds from "../components/layouts/useResponsiveThreshold";

interface ActiveRoute {
  id: string;
  focused?: boolean;
}

export interface IUIStateProvider {
  sidebarVisible: boolean;
  threshold: string;
  routers: Array<ActiveRoute>;
  toggleSidebar: () => void;
  setSidebarRef: (ref: any) => void;
  setFocusedRouter: (id: string) => void;
  updateSidebarSize: (size: number) => void;
  addRouter: () => void;
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

// Hook to set a numeric value in localStorage
function usePostSidebarState() {
  return (newValue) => {
    if (typeof newValue === "number") {
      localStorage.setItem("sidebarState", newValue.toString());
    } else {
      console.error("Sidebar state must be a number.");
    }
  };
}

// Hook to get the numeric value from localStorage
function useGetSidebarState() {
  const [state, setState] = useState(null);

  useEffect(() => {
    const fetchState = async () => {
      const savedValue = localStorage.getItem("sidebarState");
      setState(savedValue !== null ? parseFloat(savedValue) : null);
    };

    fetchState();
  }, []);

  return state;
}

export function UIStateProvider({ children }: { children: React.ReactNode }) {
  const panelRef = useRef(null); // This is the panel on the sidebar
  const [sidebarVisible, setSidebarVisibility] = useState(false);
  const [routers, setRouters] = useState<Array<ActiveRoute>>([
    { id: "url-router" },
  ]);

  const sidebarState = useGetSidebarState(); // Get the value
  const setSidebarState = usePostSidebarState(); // Mutate the value
  const autoClose = useRef(false);

  useEffect(() => {
    if (sidebarState || sidebarState === 0) {
      panelRef.current?.resize(sidebarState);
    }
  }, [sidebarState]);

  const resize = (valuePrecentage: number) => {
    setSidebarState(valuePrecentage);
    panelRef.current?.resize(valuePrecentage);
  };

  useResizeThreshold(768, (exceeded) => {
    resize(exceeded ? 0 : 20);
  });

  const addRouter = () => {
    setRouters((routers) => [...routers, { id: uuidv4() }]);
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

    if (width && width > 0) {
      resize(0);
      setSidebarVisibility(false);
    } else {
      resize(20);
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
