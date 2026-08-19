import "bootstrap/dist/css/bootstrap.css";
import { useEffect, useRef } from "react";
import "@fireback/styles/apple-family/styles.css";
import "@fireback/styles/styles.css";

// You do not have to use the mac-os family theme at all.
// this is the default theme which I use for mac desktop applications
// you could use it as a reference to build your own themes.
// themes are nothing special, rather than wrapping a set of css (scss) on a global name

import { WithFireback } from "@fireback/enterprise-shell/WithFireback";

import { useCheckAuthentication } from "@fireback/ui-core/components/layouts/ForcedAuthenticated";
import { BUILD_VARIABLES } from "@fireback/ui-core/hooks/build-variables";
import { useRtlClass } from "@fireback/ui-core/hooks/useRtlClass";
import { readUrlParam } from "@fireback/selfservice/auth.common";
import {
  useSelfServiceAuthenticateRoutes,
  useSelfServicePublicRoutes,
} from "@fireback/selfservice/SelfServiceRoutes";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {
  BrowserRouter,
  HashRouter,
  Navigate,
  Route,
  Routes,
} from "react-router-dom";

const useHashRouter = BUILD_VARIABLES.USE_HASH_ROUTER === "true";
const Router = useHashRouter ? HashRouter : BrowserRouter;

// useRtlClass() (unlike the manage app's EssentialRouter.tsx/EssentialApp.tsx,
// which already call it) was never wired up anywhere in this app at all - so
// switching language via LanguageSwitcher.tsx (Welcome.screen.tsx/
// ClassicPassport.screen.tsx) changed every piece of text but never actually
// flipped the page to rtl for fa/ar. It needs Router context (useLocale ->
// useRouter -> useLocation/useParams), which AppBody itself doesn't have
// until <Router> below is mounted, so this renders as a plain sibling of
// <Routes> inside it instead of being called directly in AppBody.
function RtlSync() {
  useRtlClass();
  return null;
}

function App() {
  const queryClient = useRef(new QueryClient());

  return (
    <QueryClientProvider client={queryClient.current}>
      <WithFireback config={{}} prefix={""} queryClient={queryClient.current}>
        <AppBody />
      </WithFireback>
    </QueryClientProvider>
  );
}

function AppBody() {
  const selfServicePublicRoutes = useSelfServicePublicRoutes();
  const selfServiceAuthenticateRoutes = useSelfServiceAuthenticateRoutes();
  const { session, checked } = useCheckAuthentication();

  // `checked` starts false until the async session check resolves. The old
  // `!session && checked ? public : authenticated` ternary treated "still
  // checking" the same as "authenticated" (its false branch) - so every
  // fresh, logged-out visitor briefly mounted the *authenticated* Router
  // first, whose catch-all (Navigate to /selfservice/passports, no query)
  // fires immediately since none of its own routes match a public path like
  // /welcome. That silently strips any ?redirect=/?workspace_type_id= query
  // string before the *correct* public router ever mounts a moment later
  // (once `checked` flips true) - by then the query is already gone, so
  // useTemporaryParamOptions (Welcome.presenter.tsx) never sees it and the
  // post-login redirect back to it never fires. Waiting for `checked` avoids
  // mounting either table - and touching window.location.hash at all - until
  // we actually know which one is correct.
  // useCompleteAuth's own onComplete (auth.common.tsx) is what normally consumes
  // ?redirect=... - but it only ever runs after an actual signin/signup mutation
  // completes. A visit that's *already* authenticated (an existing session, e.g.
  // ForcedAuthenticated.tsx's own "/selfservice?redirect=<original path>" link,
  // clicked by someone who never actually needed to log in) never goes through
  // that mutation at all, so redirect was silently ignored and the visitor just
  // landed on the default /passports screen instead. Checked once here, the same
  // two places (window.location.search, or embedded in the hash route's own
  // query) useTemporaryParamOptions already reads from.
  const redirectTarget =
    checked && session
      ? readUrlParam("redirect") || sessionStorage.getItem("redirect")
      : null;

  useEffect(() => {
    if (!redirectTarget) {
      return;
    }
    sessionStorage.removeItem("redirect");
    window.location.href = redirectTarget;
  }, [redirectTarget]);

  if (!checked || redirectTarget) {
    return null;
  }

  return (
    <>
      {!session ? (
        <Router>
          <RtlSync />
          <Routes>{selfServicePublicRoutes}</Routes>
        </Router>
      ) : (
        <Router>
          <RtlSync />
          <Routes>
            {selfServiceAuthenticateRoutes}
            <Route
              path="*"
              element={<Navigate to="/selfservice/passports" replace />}
            />
          </Routes>
        </Router>
      )}
    </>
  );
}

export default App;
