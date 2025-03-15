import {
  MenuItem,
  MenuItemRendered,
  MenuRendered,
} from "../../definitions/common";
import { useLocale } from "../../hooks/useLocale";
import classNames from "classnames";
import { useContext } from "react";
import { useQueryClient } from "react-query";
import ActiveLink from "../link/ActiveLink";
import { MenuItemContent } from "./MenuItemContent";
import { useGetUserWorkspaces } from "../../sdk/modules/workspaces/useGetUserWorkspaces";
import { RemoteQueryContext } from "../../hooks/RemoteQueryProvider";

function renderMenu(
  menu: MenuItem,
  data: {
    urw?: any;
    urws: any[];
    asPath: string;
  }
): MenuRendered | null {
  let hasChildren = false;
  const children = menu.children?.map((item): MenuItemRendered => {
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

    // // INACCURATE_MOCK_MODE is a feature I've added to show as much as content of the app,
    // // for demo purposes. It does not mean the app runs, we just show as much as things we can
    // // useful for github account. Could be fully replaced by a great mock system.
    // if (process.env.REACT_APP_INACCURATE_MOCK_MODE === "true") {
    //   hasChildren = true;
    //   isVisible = true;
    // }

    return {
      ...item,
      isActive: forceActive || false,
      isVisible,
    };
  });

  if (hasChildren === false && !menu.href) {
    return null;
  }

  return {
    name: menu.label,
    href: menu.href,
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
  const { selectedUrw } = useContext(RemoteQueryContext) as any;
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

  if (!menuRendered) {
    return null;
  }

  return (
    <div className="sidebar-menu-particle" onClick={onClick}>
      {menu.children?.length > 0 ? (
        <>
          <span className="d-flex align-items-center mb-3 mb-md-0 me-md-auto text-white text-decoration-none">
            <span className="category">{menu.label}</span>
          </span>
          <MenuUl items={menuRendered.children} />
        </>
      ) : (
        <>
          <LiItem item={menu as any} />
        </>
      )}
    </div>
  );
}

function MenuUl({ items }: { items: MenuItemRendered[] }) {
  return (
    <ul className="nav nav-pills flex-column mb-auto">
      {items.map((item) => {
        return <LiItem item={item} />;
      })}
    </ul>
  );
}

function LiItem({ item }: { item: MenuItemRendered }) {
  return (
    <li key={item.label} className={classNames("nav-item")}>
      {item.href && !item.onClick ? (
        <ActiveLink
          replace
          href={item.href}
          className="nav-link"
          aria-current="page"
          forceActive={item.isActive}
          scroll={null}
          inActiveClassName="text-white"
          activeClassName="active"
        >
          <MenuItemContent item={item} />
        </ActiveLink>
      ) : (
        <a
          className={classNames("nav-link", item.isActive && "active")}
          onClick={item.onClick}
        >
          <MenuItemContent item={item} />
        </a>
      )}
      {item.children && <MenuUl items={item.children as any} />}
    </li>
  );
}
