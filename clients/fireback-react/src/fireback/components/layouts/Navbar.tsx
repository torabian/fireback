import { MenuItem } from "@/fireback/definitions/common";
import { getOS } from "@/fireback/hooks/useHtmlClass";
import { useRouter } from "@/Router";
import classNames from "classnames";
import React from "react";
import { ActionMenuManager } from "../action-menu/ActionMenu";
import ActiveLink from "../link/ActiveLink";
import Link from "../link/Link";
import { PageTitleManager } from "../page-title/PageTitle";
import { useUiState } from "@/fireback/hooks/uiStateContext";
import { source } from "@/fireback/hooks/source";
import { osResources } from "@/resources/resources";
import { ReactiveSearch } from "../reactive-search/ReactiveSearch";

function Navbar({ menu }: { menu?: MenuItem }) {
  const router = useRouter();
  const { toggleSidebar } = useUiState();

  return (
    <nav
      className="navbar navbar-expand-lg navbar-light"
      style={{ "--wails-draggable": "drag" } as any}
    >
      <div className="container-fluid">
        <div className="page-navigator">
          <button className="navbar-menu-icon" onClick={toggleSidebar}>
            <img src={source(osResources.menu)} />
          </button>
        </div>
        <ActionMenuManager filter={({ id }) => id === "navigation"} />
        <div className="page-navigator">
          {/* <button onClick={router.goBack}>
            <img src={source(osResources.left)} />
          </button> */}

          {/* <button onClick={router.goForward}>
            <img src={source(osResources.right)} />
          </button> */}
        </div>

        <span className="navbar-brand">
          <PageTitleManager />
        </span>
        {getOS() === "web" && (
          <button
            className="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarSupportedContent"
            aria-controls="navbarSupportedContent"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span className="navbar-toggler-icon"></span>
          </button>
        )}
        <div
          className={getOS() === "web" ? "collapse navbar-collapse" : ""}
          id="navbarSupportedContent"
        >
          <ul className="navbar-nav ms-auto mb-2 mb-lg-0">
            {(menu?.children || []).map((item) => (
              <li
                className={classNames(
                  "nav-item",
                  item.children?.length && "dropdown"
                )}
                key={`${item.label}_${item.href}`}
              >
                {item.children.length ? (
                  <>
                    <ActiveLink
                      className="nav-link dropdown-toggle"
                      href={item.href}
                      id="navbarDropdown"
                      role="button"
                      data-bs-toggle="dropdown"
                      aria-expanded="false"
                    >
                      <span>{item.label}</span>
                    </ActiveLink>

                    {item?.children || [] ? (
                      <ul
                        className="dropdown-menu"
                        aria-labelledby="navbarDropdown"
                      >
                        {(item?.children || []).map((item) => {
                          return (
                            <li
                              className={classNames(
                                "nav-item",
                                item.children?.length && "dropdown"
                              )}
                              key={`${item.label}_${item.href}`}
                            >
                              <ActiveLink
                                className="dropdown-item"
                                href={item.href}
                              >
                                <span>{item.label}</span>
                              </ActiveLink>
                            </li>
                          );
                        })}
                      </ul>
                    ) : null}
                  </>
                ) : (
                  <ActiveLink
                    className="nav-link active"
                    aria-current="page"
                    href={item.href}
                  >
                    <span>{item.label}</span>
                  </ActiveLink>
                )}
              </li>
            ))}
          </ul>
          <span className="general-action-menu desktop-view">
            <ActionMenuManager filter={({ id }) => id !== "navigation"} />
          </span>
          <ReactiveSearch />
        </div>
      </div>
    </nav>
  );
}

export default React.memo(Navbar);
