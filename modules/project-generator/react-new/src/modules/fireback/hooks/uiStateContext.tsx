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
  // react-resizable-panels only knows percentages of the group, so on its
  // own it keeps the sidebar's *share* constant as the window resizes —
  // which means its actual pixel width grows/shrinks with the window.
  // We want the opposite: the sidebar should hold a fixed pixel width and
  // let the router panel(s) absorb the change, same as a manual drag
  // would. This tracks that fixed target in pixels.
  const sidebarWidthPx = useRef<number | null>(null);
  const [sidebarVisible, setSidebarVisibility] = useState(false);

  const [routers, setRouters] = useState<Array<ActiveRoute>>([
    { id: "url-router" },
  ]);

  // The group's own rendered width — not window.innerWidth, which only
  // happens to match it in today's layout. Measuring the actual element
  // is what react-resizable-panels itself does internally to turn drag
  // deltas into percentages, so this stays correct even if the group
  // stops spanning the full viewport later (e.g. a fixed-width chrome
  // panel added outside it).
  const getGroupWidthPx = () => {
    const element = document.querySelector<HTMLElement>(".application-panels");
    return element?.getBoundingClientRect().width || window.innerWidth;
  };

  const syncSidebarWidthPx = () => {
    const size = panelRef.current?.getSize();
    if (size && size > 0) {
      sidebarWidthPx.current = (size / 100) * getGroupWidthPx();
    }
  };

  const persistSidebarSize = (newValue: number) => {
    userPreferedWidth.current = newValue;
    localStorage.setItem("sidebarState", newValue.toString());
    syncSidebarWidthPx();
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

  // Re-apply the sidebar's fixed pixel width whenever the panel GROUP
  // itself resizes, so it holds its size while the router panel(s)
  // grow/shrink instead of everything scaling together. Re-reads the
  // panel's live size each time rather than trusting React state, so it
  // never fights useResizeThreshold's collapse/expand at the 768px
  // breakpoint.
  //
  // Uses a ResizeObserver on the group element itself (same pattern as
  // useResponsiveThreshold.tsx) rather than a window "resize" listener:
  // it fires continuously as the element's own box changes — including
  // mid-drag while the OS window edge is being dragged — instead of only
  // once events settle, and it reports the group's actual pixel size
  // directly instead of us having to approximate it from the window.
  useEffect(() => {
    const element = document.querySelector<HTMLElement>(".application-panels");
    if (!element) {
      return;
    }

    const applyLockedWidth = (groupWidthPx: number) => {
      if (detectDeviceType().isMobileView) {
        return;
      }

      const currentSize = panelRef.current?.getSize();
      if (!currentSize || currentSize <= 0) {
        // Collapsed/hidden — leave it alone rather than forcing it open.
        return;
      }

      if (sidebarWidthPx.current == null) {
        // No fixed target established yet — this measurement just sets
        // the baseline; it takes effect from the next one onward.
        sidebarWidthPx.current = (currentSize / 100) * groupWidthPx;
        return;
      }

      const newPercentage = (sidebarWidthPx.current / groupWidthPx) * 100;
      resize(newPercentage);
    };

    const resizeObserver = new ResizeObserver((entries) => {
      for (const entry of entries) {
        applyLockedWidth(entry.contentRect.width);
      }
    });

    resizeObserver.observe(element);

    return () => {
      resizeObserver.unobserve(element);
      resizeObserver.disconnect();
    };
  }, []);

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
    const suggestedSize = (180 / getGroupWidthPx()) * 100;
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
      syncSidebarWidthPx();
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
      syncSidebarWidthPx();
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
    handleThresholds,
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
