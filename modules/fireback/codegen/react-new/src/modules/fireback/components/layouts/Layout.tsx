import { Outlet } from "react-router-dom";
import { MenuItem } from "../../definitions/common";
import { useT } from "../../hooks/useT";
import { UploaderStatsCard } from "../../modules/manage/drive/UploaderStatsCard";
import { ActionMenuManager } from "../action-menu/ActionMenu";
import Navbar from "./Navbar";

// We do not compile the pull to refresh for desktop and web
////// # if env.TARGET_TYPE == 'mobile' && !env.DISABLE_PULL_TO_REFRESH
import classNames from "classnames";
import { useContext } from "react";
import { useUiState } from "../../hooks/uiStateContext";
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

  const { sidebarVisible, toggleSidebar: toggleSidebar$ } = useUiState();

  const onSearch = phrase.length > 0;

  return (
    <>
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
    </>
  );
};

export default Layout;
