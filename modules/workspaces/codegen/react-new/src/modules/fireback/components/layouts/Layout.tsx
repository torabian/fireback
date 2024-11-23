import { MenuItem } from "@/modules/fireback/definitions/common";
import { useT } from "@/modules/fireback/hooks/useT";
import { UploaderStatsCard } from "@/modules/fireback/modules/drive/UploaderStatsCard";
import { Outlet } from "react-router-dom";
import { ActionMenuManager } from "../action-menu/ActionMenu";
import { ForcedAuthenticated } from "./ForcedAuthenticated";
import Navbar from "./Navbar";

// We do not compile the pull to refresh for desktop and web
////// # if env.TARGET_TYPE == 'mobile' && !env.DISABLE_PULL_TO_REFRESH
import { useUiState } from "@/modules/fireback/hooks/uiStateContext";
import classNames from "classnames";
import { useContext, useEffect, useRef } from "react";
import { ReactiveSearchContext } from "../reactive-search/ReactiveSearchContext";
import { ReactiveSearchResult } from "../reactive-search/ReactiveSearchResult";
// @ts-ignore
function ContentSection({ children }: any) {
  return (
    <>
      <Outlet />
      {children}
    </>
  );
}
// /// # else
// // @ts-ignore
// function ContentSection({ children }: any) {
//   return (
//     <>
//       <Outlet />
//       {children}
//     </>
//   );
// }
// /// # endif

const Layout = ({
  children,
  navbarMenu,
  sidebarMenu,
  routerId,
}: {
  children?: React.ReactNode;
  sidebarMenu?: MenuItem | MenuItem[];
  navbarMenu?: MenuItem;
  routerId?: string;
}) => {
  const t = useT();
  const { result, phrase, reset } = useContext(ReactiveSearchContext);

  const {
    sidebarVisible,
    toggleSidebar: toggleSidebar$,
    hide,
    setSidebarRef,
    updateSidebarSize,
    routers,
    setFocusedRouter,
  } = useUiState();

  const boxRef = useRef(null);
  const panelRef = useRef(null); // Ref for the left panel
  const autoClose = useRef(false);

  useEffect(() => {
    setSidebarRef(panelRef.current);
  }, [panelRef.current]);

  // Cordova thingy
  // useEffect(() => {
  //   function listener() {
  //     alert("Keyboard viewed!");
  //   }

  //   window.addEventListener("keyboardDidShow", listener);

  //   return () => window.removeEventListener("keyboardDidShow", listener);
  // }, []);

  const onSearch = phrase.length > 0;

  return (
    <>
      <ForcedAuthenticated>
        <div style={{ display: "flex", width: "100%" }}>
          <div
            className={classNames(
              "sidebar-overlay",
              sidebarVisible ? "open" : ""
            )}
            onClick={(e) => {
              toggleSidebar$();
              e.stopPropagation();
            }}
          ></div>
          <div style={{ width: "100%", flex: 1 }}>
            <Navbar routerId={routerId} menu={navbarMenu} />
            <div className="content-section">
              {onSearch ? (
                <div className="content-container">
                  <ReactiveSearchResult
                    onComplete={() => reset()}
                    result={result}
                  />
                </div>
              ) : null}
              <div
                className="content-container"
                style={{ visibility: !onSearch ? undefined : "hidden" }}
              >
                <ContentSection>{children}</ContentSection>
              </div>
            </div>
          </div>

          <UploaderStatsCard />
        </div>
        <span className="general-action-menu mobile-view">
          <ActionMenuManager />
        </span>
      </ForcedAuthenticated>
    </>
  );
};

export default Layout;
