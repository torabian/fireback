/**
 * Tools for authentication based on fireback ABAC plugin
 */

import React, { useContext, useState } from "react";

export interface IUIStateProvider {
  sidebarVisible: boolean;
  toggleSidebar: () => void;
}

export const UIStateContext = React.createContext<IUIStateProvider>({
  sidebarVisible: false,
  toggleSidebar() {},
});

export function useUiState() {
  return useContext(UIStateContext);
}

export function UIStateProvider({ children }: { children: React.ReactNode }) {
  const [sidebarVisible, setSidebarVisibility] = useState<boolean>(false);

  const toggleSidebar = () => {
    setSidebarVisibility((s) => !s);
  };

  return (
    <UIStateContext.Provider
      value={{
        sidebarVisible,
        toggleSidebar,
      }}
    >
      {children}
    </UIStateContext.Provider>
  );
}
