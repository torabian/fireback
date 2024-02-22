import { MenuItem, MenuItemRendered, MenuRendered } from "@/definitions/common";
import { useLocale } from "@/hooks/useLocale";
import classNames from "classnames";
import ActiveLink from "../link/ActiveLink";
import { MenuItemContent } from "./MenuItemContent";
import { useContext } from "react";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useGetUserRoleWorkspaces } from "src/sdk/fireback/modules/workspaces/useGetUserRoleWorkspaces";
import { useQueryClient } from "react-query";
import { UserRoleWorkspaceEntity } from "src/sdk/fireback";

function renderMenu(
  menu: MenuItem,
  data: {
    urw?: UserRoleWorkspaceEntity;
    urws: UserRoleWorkspaceEntity[];
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

  if (hasChildren === false) {
    return null;
  }

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
  const { query: queryWorkspaces } = useGetUserRoleWorkspaces({
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
      <span className="d-flex align-items-center mb-3 mb-md-0 me-md-auto text-white text-decoration-none">
        <span className="category">{menu.label}</span>
      </span>
      <MenuUl items={menuRendered.children} />
    </div>
  );
}

function MenuUl({ items }: { items: MenuItemRendered[] }) {
  return (
    <ul className="nav nav-pills flex-column mb-auto">
      {items.map((item) => {
        return (
          <li
            className={classNames("nav-item")}
            key={`${item.href}_${item.label}`}
          >
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
      })}
    </ul>
  );
}
