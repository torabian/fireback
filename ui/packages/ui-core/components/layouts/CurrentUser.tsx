import { useQueryClient } from "@tanstack/react-query";
import { BUILD_VARIABLES } from "../../hooks/build-variables";
import { source } from "../../hooks/source";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { osResources } from "../../hooks/resources";
import { useAuthentication } from "@fireback/auth-client/AuthenticationContext";
import Link from "../link/Link";
import { strings } from "../strings/translations";

export function CurrentUser({ onClick }: { onClick: () => void }) {
  const { isAuthenticated, signout } = useAuthentication();
  const router = useRouter();
  const s = useS(strings);
  const queryClient = useQueryClient();
  const signout$ = () => {
    onClick();
    signout();
    queryClient.setQueriesData("*fireback.UserRoleWorkspace", []);
    if (BUILD_VARIABLES.NAVIGATE_ON_SIGNOUT) {
      router.push(
        BUILD_VARIABLES.NAVIGATE_ON_SIGNOUT,
        BUILD_VARIABLES.NAVIGATE_ON_SIGNOUT,
      );
    }
  };

  const onSignoutClick = () => {
    if (confirm(s.components.confirmLeaveApp)) {
      signout$();
    }
  };

  if (!isAuthenticated) {
    return (
      <Link className="user-signin-section" href="/signin" onClick={onClick}>
        <img src={BUILD_VARIABLES.PUBLIC_URL + "/common/user.svg"} />
        {s.currentUser.signin}
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
              <span className="nav-link-text">{s.currentUser.signout}</span>
            </span>
          </a>
        </li>
      </ul>
    </div>
  );
}
