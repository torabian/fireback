import { type MenuItem } from "../../definitions/common";
import { source } from "../../hooks/source";
import { useUiState } from "../../hooks/uiStateContext";

import classNames from "classnames";
import React, { useContext } from "react";
import { BUILD_VARIABLES } from "../../hooks/build-variables";
import { detectDeviceType } from "../../hooks/deviceInformation";
import { useRemoteMenuResolver } from "../../hooks/useRemoteMenuResolver";
import { osResources } from "../../resources/resources";
import type { AppMenuOptionalDto } from "../../sdk/abac/AppMenuOptionalDto";
import { AppMenuEntity } from "../../sdk/modules/abac/AppMenuEntity";
import { ReactiveSearchContext } from "../reactive-search/ReactiveSearchContext";
import { CurrentUser } from "./CurrentUser";
import { MenuParticle } from "./MenuParticle";
import { useWorkspacesMenuPresenter } from "./useWorkspacesMenuPresenter";

export function dataMenuToMenu(
  data: AppMenuOptionalDto,
  permissionCheck: (permissionKey?: string | null) => boolean = () => true,
  locale: string,
): MenuItem | null {
  console.log(17, data.capabilityId);
  if (!permissionCheck(data.capabilityId)) {
    return null;
  }

  const children = (data.children || [])
    .map((v: AppMenuOptionalDto) => dataMenuToMenu(v, permissionCheck, locale))
    .filter(Boolean) as MenuItem[];

  let label = data.label || "";

  if (typeof label !== "string") {
    label = label[locale];
  }

  return {
    label,

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

function Sidebar({
  miniSize,
  onClose,
  sidebarItemSelectedExtra,
}: {
  miniSize: boolean;
  onClose?: () => void;
  sidebarItemSelectedExtra?: () => void;
}) {
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

  console.log(5, menus);

  return (
    <div
      data-wails-drag
      className={classNames(
        miniSize ? "sidebar-extra-small" : "",
        "sidebar",
        sidebarVisible ? "open" : "",
        "scrollable-element",
        detectDeviceType().isMobileView ? "has-bottom-tab" : undefined,
      )}
    >
      <button
        className="sidebar-close"
        onClick={() => (onClose ? onClose() : toggleSidebar())}
      >
        <img src={source(osResources.cancel)} />
      </button>

      {menus.map((menu) => (
        <MenuParticle
          onClick={() => {
            sidebarItemSelected();
            sidebarItemSelectedExtra?.();
          }}
          key={JSON.stringify(menu)}
          menu={menu}
        />
      ))}
      {BUILD_VARIABLES.GITHUB_DEMO === "true" && (
        <MenuParticle
          onClick={() => {
            sidebarItemSelected();
            sidebarItemSelectedExtra?.();
          }}
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
      <CurrentUser
        onClick={() => {
          sidebarItemSelected();
          sidebarItemSelectedExtra?.();
        }}
      />
    </div>
  );
}

export default React.memo(Sidebar);
