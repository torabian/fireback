import { MenuItem } from "@/definitions/common";
import { useT } from "@/hooks/useT";
import { UploaderStatsCard } from "@/modules/drive/UploaderStatsCard";
import { Outlet } from "react-router-dom";
import { ForcedAuthenticated } from "./ForcedAuthenticated";
import Navbar from "./Navbar";
import Sidebar from "./Sidebar";
import { ActionMenuManager } from "../action-menu/ActionMenu";

// We do not compile the pull to refresh for desktop and web
////// # if env.TARGET_TYPE == 'mobile' && !env.DISABLE_PULL_TO_REFRESH
import ReactPullToRefresh from "react-pull-to-refresh";
import { useContext, useEffect } from "react";
import { ReactiveSearchResult } from "../reactive-search/ReactiveSearchResult";
import { ReactiveSearchContext } from "../reactive-search/ReactiveSearchContext";
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
}: {
  children?: React.ReactNode;
  sidebarMenu?: MenuItem | MenuItem[];
  navbarMenu?: MenuItem;
}) => {
  const t = useT();
  const { result, phrase, reset } = useContext(ReactiveSearchContext);
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
        <div style={{ display: "flex" }}>
          {sidebarMenu && <Sidebar menu={sidebarMenu} />}
          <Navbar menu={navbarMenu} />
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
