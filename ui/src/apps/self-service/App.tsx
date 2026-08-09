import { useRef } from "react";
import "bootstrap/dist/css/bootstrap.css";
import "../../modules/styles/styles.css";
import "../../modules/styles/apple-family/styles.css";

// You do not have to use the mac-os family theme at all.
// this is the default theme which I use for mac desktop applications
// you could use it as a reference to build your own themes.
// themes are nothing special, rather than wrapping a set of css (scss) on a global name

import { WithFireback } from "../core/WithFireback";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {
  BrowserRouter,
  HashRouter,
  Navigate,
  Route,
  Routes,
} from "react-router-dom";
import { useCheckAuthentication } from "@/modules/fireback-ui/components/layouts/ForcedAuthenticated";
import {
  useSelfServiceAuthenticateRoutes,
  useSelfServicePublicRoutes,
} from "@/modules/selfservice/SelfServiceRoutes";
import { BUILD_VARIABLES } from "@/modules/fireback-ui/hooks/build-variables";
import { SessionGate } from "@/modules/fireback-ui/components/session-gate/SessionGate";
import { checkSessionViaWhoami } from "@/modules/fireback-ui/components/session-gate/checkSessionViaWhoami";

const useHashRouter = BUILD_VARIABLES.USE_HASH_ROUTER === "true";
const Router = useHashRouter ? HashRouter : BrowserRouter;

function App() {
  const queryClient = useRef(new QueryClient());

  return (
    <SessionGate checkSession={checkSessionViaWhoami}>
      <QueryClientProvider client={queryClient.current}>
        <WithFireback
          config={{}}
          prefix={""}
          queryClient={queryClient.current}
        >
          <AppBody />
        </WithFireback>
      </QueryClientProvider>
    </SessionGate>
  );
}

function AppBody() {
  const selfServicePublicRoutes = useSelfServicePublicRoutes();
  const selfServiceAuthenticateRoutes = useSelfServiceAuthenticateRoutes();
  const { session, checked } = useCheckAuthentication();

  return (
    <>
      {!session && checked ? (
        <Router>
          <Routes>
            <Route path=":locale">{selfServicePublicRoutes}</Route>
            <Route
              path="*"
              element={<Navigate to="/en/selfservice/welcome" replace />}
            />
          </Routes>
        </Router>
      ) : (
        <Router>
          <Routes>
            <Route path=":locale">{selfServiceAuthenticateRoutes}</Route>
            <Route
              path="*"
              element={<Navigate to="/en/selfservice/passports" replace />}
            />
          </Routes>
        </Router>
      )}
    </>
  );
}

export default App;
