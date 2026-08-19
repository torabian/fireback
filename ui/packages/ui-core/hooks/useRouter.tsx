import { useNavigate, useParams, Link, useLocation } from "react-router-dom";

export const RouterLink = (props: any) => {
  return (
    <Link {...props} to={props.href}>
      {props.children}
    </Link>
  );
};

export function useRouter() {
  const navigate = useNavigate();
  const params = useParams();
  const location$ = useLocation();
  // Bug fix: this used to also replace a literal "{locale}" placeholder in
  // `path` with a locale parsed back out of window.location.pathname, and
  // (dead code - the flag driving it was hardcoded false) strip a leading
  // "/xx/" locale segment when set. Routes no longer carry a locale segment
  // at all (see EssentialRouter.tsx/App.tsx's own history) and no caller
  // builds a "{locale}"-templated path anymore, so this is just a plain
  // passthrough now.
  const push = (
    path: string,
    actual?: string,
    params?: any,
    replace = false
  ) => {
    navigate(path, { replace, state: params });
  };

  const replace = (path: string, actual?: string, params?: any) => {
    push(path, actual, params, true);
  };

  // Bug fix: this used to unconditionally call navigate(-1), discarding
  // goToIfNoBack entirely - so landing on a page directly (a fresh page load, e.g.
  // visiting a link or a browser refresh, rather than clicking to it from elsewhere
  // in the app) and then cancelling/submitting had nowhere in the app's own history to
  // go back to. navigate(-1) in that case falls through to whatever the browser's
  // real history had before this SPA ever loaded (often nothing render-able), which
  // hangs the page on a pending navigation instead of landing anywhere in the app.
  // React Router sets the current location's `key` to the literal string "default"
  // only for that first, no-prior-app-entry case - every entry actually pushed by the
  // app (via `push`/`navigate`) gets a real generated key - so it's what
  // distinguishes "safe to go back" from "must use the fallback" here.
  const goBackOrDefault = (goToIfNoBack: string) => {
    if (location$.key && location$.key !== "default") {
      navigate(-1);
    } else {
      push(goToIfNoBack);
    }
  };

  return {
    asPath: location$.pathname,
    state: location$.state,
    pathname: "",
    query: params,
    push,
    goBack: () => navigate(-1),
    goBackOrDefault,
    goForward: () => navigate(1),
    replace,
  };
}
