import { useNavigate, useParams, Link, useLocation } from "react-router-dom";
import { localeFromPath } from "./hooks/localeFromPath";

export const RouterLink = (props: any) => {
  return (
    <Link {...props} to={props.href}>
      {props.children}
    </Link>
  );
};

export function useRouter() {
  const noPrefix = process.env.REACT_APP_NO_LOCALE_PREFIX === "true";
  const navigate = useNavigate();
  const params = useParams();
  const location$ = useLocation();
  const push = (
    path: string,
    actual?: string,
    params?: any,
    replace = false
  ) => {
    const locale = localeFromPath(window.location.pathname);
    let goToPath = path.replace("{locale}", locale);

    if (noPrefix) {
      if (goToPath.match(/\/[a-z]{2}\//)) {
        goToPath = goToPath.substring(3);
      }
    }

    navigate(goToPath, { replace });
  };

  const replace = (path: string, actual?: string, params?: any) => {
    push(path, actual, params, true);
  };

  return {
    asPath: location$.pathname,
    pathname: "",
    query: params,
    push,
    goBack: () => navigate(-1),
    goBackOrDefault: (goToIfNoBack: string) => navigate(-1),
    goForward: () => navigate(1),
    replace,
  };
}
