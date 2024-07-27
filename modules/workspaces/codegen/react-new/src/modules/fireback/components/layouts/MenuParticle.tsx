import {
  MenuItem,
  MenuItemRendered,
  MenuRendered,
} from "@/modules/fireback/definitions/common";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { useGetUserWorkspaces } from "../../sdk/modules/workspaces/useGetUserWorkspaces";
import classNames from "classnames";
import { useContext } from "react";
import { useQueryClient } from "react-query";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import ActiveLink from "../link/ActiveLink";
import { MenuItemContent } from "./MenuItemContent";

function renderMenu(
  menu: MenuItem,
  data: {
    urw?: any;
    urws: any[];
    asPath: string;
  }
): MenuRendered | null {
  let hasChildren = false;
  const children = menu.children.map((item): MenuItemRendered => {
    let forceActive = item.activeMatcher
      ? item.activeMatcher.test(data.asPath)
      : undefined;

    if (item.forceActive) {
      forceActive = true;
    }

    let isVisible = item.displayFn
      ? item.displayFn({
          location: "here",
          asPath: data.asPath,
          selectedUrw: data.urw,
          userRoleWorkspaces: data.urws,
        })
      : true;

    if (isVisible) {
      hasChildren = true;
    }

    // INACCURATE_MOCK_MODE is a feature I've added to show as much as content of the app,
    // for demo purposes. It does not mean the app runs, we just show as much as things we can
    // useful for github account. Could be fully replaced by a great mock system.
    if (process.env.REACT_APP_INACCURATE_MOCK_MODE === "true") {
      hasChildren = true;
      isVisible = true;
    }

    return {
      ...item,
      // isActive: true,
      // isVisible: true,
      isActive: forceActive || false,
      isVisible,
    };
  });

  // if (hasChildren === false) {
  //   return null;
  // }

  return {
    name: menu.label,
    children,
  };
}
export function MenuParticle({
  menu,
  onClick,
}: {
  menu: MenuItem;
  onClick: () => void;
}) {
  const { asPath } = useLocale();
  const queryClient = useQueryClient();
  const { selectedUrw } = useContext(RemoteQueryContext);
  const { query: queryWorkspaces } = useGetUserWorkspaces({
    queryClient,
    query: {},
    queryOptions: {
      refetchOnWindowFocus: false,
      cacheTime: 0,
    },
  });

  const menuRendered = renderMenu(menu, {
    asPath,
    urw: selectedUrw as any,
    urws: queryWorkspaces.data?.data?.items || [],
  });

  // Here we decide if there is a single item, or has children.
  // if it doesn't have children, means itself is a menu

  if (!menu.children || menu.children.length === 0) {
    return (
      <div className="sidebar-menu-particle">
        <ul className="nav nav-pills flex-column mb-auto">
          <MenuLi menu={menu as any} />
        </ul>
      </div>
    );
  }

  return (
    <div className="sidebar-menu-particle" onClick={onClick}>
      <span className="d-flex align-items-center mb-3 mb-md-0 me-md-auto text-white text-decoration-none">
        <span className="category">{menu.label}</span>
      </span>
      <MenuUl items={menuRendered.children} />
    </div>
  );
}

function MenuLi({ menu }: { menu: MenuItemRendered }) {
  return (
    <li className={classNames("nav-item")}>
      {menu.href && !menu.onClick ? (
        <ActiveLink
          replace
          href={menu.href}
          className="nav-link"
          aria-current="page"
          forceActive={menu.isActive}
          scroll={null}
          inActiveClassName="text-white"
          activeClassName="active"
        >
          <MenuItemContent item={menu} />
        </ActiveLink>
      ) : (
        <a
          className={classNames("nav-link", menu.isActive && "active")}
          onClick={menu.onClick}
        >
          <MenuItemContent item={menu} />
        </a>
      )}
      {menu.children && <MenuUl items={menu.children as any} />}
    </li>
  );
}

function MenuUl({ items }: { items: MenuItemRendered[] }) {
  return (
    <ul className="nav nav-pills flex-column mb-auto">
      {items.map((menu) => {
        return <MenuLi menu={menu} />;
      })}
    </ul>
  );
}
