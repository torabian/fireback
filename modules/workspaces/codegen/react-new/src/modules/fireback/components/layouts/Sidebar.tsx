import { MenuItem } from "../../definitions/common";
import { source } from "../../hooks/source";
import { useUiState } from "../../hooks/uiStateContext";
import { onPermission, onPermissionInRoot } from "../../hooks/accessLevels";
import { useT } from "../../hooks/useT";

import classNames from "classnames";
import React, { useContext } from "react";
import { useQueryClient } from "react-query";
import { ReactiveSearchContext } from "../reactive-search/ReactiveSearchContext";
import { CurrentUser } from "./CurrentUser";
import { MenuParticle } from "./MenuParticle";
import { WorkspacesMenuParticle } from "./WorkspacesMenuParticle";
import { AppMenuEntity } from "../../sdk/modules/workspaces/AppMenuEntity";
import { useGetWidgetAreas } from "../../sdk/modules/widget/useGetWidgetAreas";
import { osResources } from "../../resources/resources";
import { useRemoteMenuResolver } from "../../hooks/useRemoteMenuResolver";

export function dataMenuToMenu(
  data: AppMenuEntity,
  permissionCheck: (permissionKey?: string | null) => boolean = () => true
): MenuItem | null {
  if (!permissionCheck(data.capabilityId)) {
    return null;
  }

  const children = (data.children || [])
    .map((v: AppMenuEntity) => dataMenuToMenu(v, permissionCheck))
    .filter(Boolean) as MenuItem[];

  return {
    label: data.label || "",

    children,
    displayFn: castMenuDefinitionToDisplayFn(data),
    icon: data.icon,
    href: data.href,
    activeMatcher: data.activeMatcher
      ? new RegExp(data.activeMatcher)
      : undefined,
  };
}

function castMenuDefinitionToDisplayFn(data: AppMenuEntity) {
  if (data.applyType === "permission" && data.capabilityId) {
    return onPermission(data.capabilityId);
  }

  if (data.applyType === "permissionInRoot" && data.capabilityId) {
    return onPermissionInRoot(data.capabilityId);
  }

  return () => true;
}

export const defaultNavbar: MenuItem = {
  label: "Navbar",
  children: [],
};

function Sidebar({ miniSize }: { miniSize: boolean }) {
  const {
    sidebarVisible,
    toggleSidebar: toggleSidebar$,
    sidebarItemSelected,
  } = useUiState();
  const menu = useRemoteMenuResolver("sidebar");

  const { reset } = useContext(ReactiveSearchContext);

  const toggleSidebar = () => {
    reset();
    toggleSidebar$();
  };

  const queryClient = useQueryClient();
  const { query } = useGetWidgetAreas({ queryClient, query: {} });
  const t = useT();

  // let dashboardMenu: MenuItem = {
  //   label: t.dashboards,
  //   children: [
  //     ...(query.data?.data?.items || []).map((item) => {
  //       return {
  //         children: [],
  //         label: item.name || "",
  //         href: "/dashboard/" + item.uniqueId,
  //         icon: osResources.dashboard,
  //       };
  //     }),
  //     {
  //       children: [],
  //       label: t.widgetPicker.widgets || "",
  //       href: "/widgets",
  //       icon: osResources.dashboard,
  //     },
  //   ],
  // };

  if (!menu) {
    return null;
  }

  let menus: MenuItem[] = [];
  if (Array.isArray(menu)) {
    menus = [...menu];
  } else if ((menu as any).children?.length) {
    menus.push(menu);
  }

  return (
    <div
      data-wails-drag
      className={classNames(
        miniSize ? "sidebar-extra-small" : "",
        "sidebar",
        sidebarVisible ? "open" : "",
        "scrollable-element"
      )}
      style={{ display: "flex", height: "100vh" }}
    >
      <button className="sidebar-close" onClick={toggleSidebar}>
        <img src={source(osResources.cancel)} />
      </button>

      {/* {process.env.REACT_APP_FEATURE_DASHBOARD === "true" && (
            <MenuParticle onClick={toggleSidebar} menu={dashboardMenu} />
          )} */}

      {menus.map((menu) => (
        <MenuParticle
          onClick={sidebarItemSelected}
          key={menu.label}
          menu={menu}
        />
      ))}
      <WorkspacesMenuParticle onClick={sidebarItemSelected} />

      <CurrentUser onClick={sidebarItemSelected} />
    </div>
  );
}

export default React.memo(Sidebar);
