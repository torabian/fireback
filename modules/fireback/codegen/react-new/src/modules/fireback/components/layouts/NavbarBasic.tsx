import { MenuItem } from "../../definitions/common";
import classNames from "classnames";
import React from "react";
import ActiveLink from "../link/ActiveLink";

function Navbar({ menu }: { menu?: MenuItem }) {
  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-light">
      <div className="container-fluid">
        <a className="navbar-brand" href="#">
          {process.env.productName}
        </a>
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
        <div className="collapse navbar-collapse" id="navbarSupportedContent">
          <ul className="navbar-nav me-auto mb-2 mb-lg-0">
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
          <form className="d-flex">
            <input
              className="form-control me-2"
              type="search"
              placeholder="Search"
              aria-label="Search"
            />
            <button className="btn btn-outline-success" type="submit">
              Search
            </button>
          </form>
        </div>
      </div>
    </nav>
  );
}

export default React.memo(Navbar);
