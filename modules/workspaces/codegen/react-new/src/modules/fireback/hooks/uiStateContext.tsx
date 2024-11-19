/**
 * Tools for authentication based on fireback ABAC plugin
 */

import React, { useContext, useEffect, useRef, useState } from "react";
import { useResizeThreshold } from "./useResizeThreshold";

export interface IUIStateProvider {
  sidebarVisible: boolean;
  toggleSidebar: () => void;
  setSidebarRef: (ref: any) => void;
  updateSidebarSize: (size: number) => void;
  sidebarItemSelected: () => void;
  hide: () => void;
  show: () => void;
}

export const UIStateContext = React.createContext<IUIStateProvider>({
  sidebarVisible: false,
  toggleSidebar() {},
  setSidebarRef(ref) {},
  sidebarItemSelected() {},
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

  const sidebarState = useGetSidebarState(); // Get the value
  const setSidebarState = usePostSidebarState(); // Mutate the value

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

  const toggleSidebar = () => {
    const isVisible = panelRef.current?.getSize();
    if (isVisible) {
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

  const updateSidebarSize = (size: number) => {
    resize(size);
  };

  const show = () => {
    if (panelRef.current) {
      resize(20);
      setSidebarVisibility(true);
    }
  };

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
        show,
        updateSidebarSize,
        setSidebarRef,
        sidebarVisible,
        toggleSidebar,
      }}
    >
      {children}
    </UIStateContext.Provider>
  );
}
