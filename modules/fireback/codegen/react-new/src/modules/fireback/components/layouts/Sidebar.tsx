import { MenuItem } from "../../definitions/common";
import { source } from "../../hooks/source";
import { useUiState } from "../../hooks/uiStateContext";

import classNames from "classnames";
import React, { useContext } from "react";
import { useRemoteMenuResolver } from "../../hooks/useRemoteMenuResolver";
import { osResources } from "../../resources/resources";
import { AppMenuEntity } from "../../sdk/modules/abac/AppMenuEntity";
import { ReactiveSearchContext } from "../reactive-search/ReactiveSearchContext";
import { CurrentUser } from "./CurrentUser";
import { MenuParticle } from "./MenuParticle";
import { useWorkspacesMenuPresenter } from "./useWorkspacesMenuPresenter";

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

  if (!menu) {
    return null;
  }

  let menus: MenuItem[] = [];
  if (Array.isArray(menu)) {
    menus = [...menu];
  } else if ((menu as any).children?.length) {
    menus.push(menu);
  }

  const { menus: workspaceMenus } = useWorkspacesMenuPresenter();
  menus.push(workspaceMenus[0]);

  return (
    <div
      data-wails-drag
      className={classNames(
        miniSize ? "sidebar-extra-small" : "",
        "sidebar",
        sidebarVisible ? "open" : "",
        "scrollable-element"
      )}
      style={{ display: "flex", height: "calc(100vh - 60px)" }}
    >
      <button className="sidebar-close" onClick={toggleSidebar}>
        <img src={source(osResources.cancel)} />
      </button>

      {menus.map((menu) => (
        <MenuParticle
          onClick={sidebarItemSelected}
          key={menu.label}
          menu={menu}
        />
      ))}
      {process.env.REACT_APP_GITHUB_DEMO === "true" && (
        <MenuParticle
          onClick={sidebarItemSelected}
          menu={{
            label: "Demo",
            children: [
              {
                label: "Form select",
                icon: "/ios-theme/icons/settings.svg",
                children: [],
                href: "/demo/form-select",
              },
              {
                label: "Form Date/Time",
                icon: "/ios-theme/icons/settings.svg",
                children: [],
                href: "/demo/form-date",
              },
              {
                label: "Overlays & Modal",
                icon: "/ios-theme/icons/settings.svg",
                children: [],
                href: "/demo/modals",
              },
            ],
          }}
        />
      )}
      <CurrentUser onClick={sidebarItemSelected} />
    </div>
  );
}

export default React.memo(Sidebar);
