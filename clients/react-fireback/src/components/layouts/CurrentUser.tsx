import { useRouter } from "@/Router";
import { useT } from "@/hooks/useT";
import { useContext } from "react";
import { useQueryClient } from "react-query";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import Link from "../link/Link";
import { source } from "@/helpers/source";
import { osResources } from "../mulittarget/multitarget-resource";

export function CurrentUser({ onClick }: { onClick: () => void }) {
  const { isAuthenticated, signout } = useContext(RemoteQueryContext);
  const router = useRouter();
  const t = useT();
  const queryClient = useQueryClient();
  const signout$ = () => {
    signout();
    queryClient.setQueriesData("*workspaces.UserRoleWorkspace", []);
    if (process.env.REACT_APP_NAVIGATE_ON_SIGNOUT) {
      router.push(
        process.env.REACT_APP_NAVIGATE_ON_SIGNOUT,
        process.env.REACT_APP_NAVIGATE_ON_SIGNOUT
      );
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
          <a onClick={signout$} className="nav-link text-white">
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
