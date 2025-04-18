import { useRouter } from "../../hooks/useRouter";
import { useT } from "../../hooks/useT";
import { useContext } from "react";
import { useQueryClient } from "react-query";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import Link from "../link/Link";
import { source } from "../../hooks/source";
import { osResources } from "../../resources/resources";
import { ModalContext } from "../modal/Modal";

export function CurrentUser({ onClick }: { onClick: () => void }) {
  const { isAuthenticated, signout } = useContext(RemoteQueryContext);
  const router = useRouter();
  const t = useT();
  const queryClient = useQueryClient();
  const useModal = useContext(ModalContext);
  const signout$ = () => {
    onClick();
    signout();
    queryClient.setQueriesData("*workspaces.UserRoleWorkspace", []);
    if (process.env.REACT_APP_NAVIGATE_ON_SIGNOUT) {
      router.push(
        process.env.REACT_APP_NAVIGATE_ON_SIGNOUT,
        process.env.REACT_APP_NAVIGATE_ON_SIGNOUT
      );
    }
  };

  const onSignoutClick = () => {
    if (confirm("Are you sure to leave the app?")) {
      signout$();
    }
  };

  if (!isAuthenticated) {
    return (
      <Link className="user-signin-section" href="/signin" onClick={onClick}>
        <img src={process.env.PUBLIC_URL + "/common/user.svg"} />
        {t.currentUser.signin}
      </Link>
    );
  }

  return (
    <div className="sidebar-menu-particle mt-5">
      <ul className="nav nav-pills flex-column mb-auto">
        <li className="nav-item">
          <a onClick={onSignoutClick} className="nav-link text-white">
            <span>
              <img className="menu-icon" src={source(osResources.turnoff)} />
              <span className="nav-link-text">{t.currentUser.signout}</span>
            </span>
          </a>
        </li>
      </ul>
    </div>
  );
}
