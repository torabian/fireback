import "bootstrap/dist/css/bootstrap.css";
import { useRef } from "react";
import "@fireback/styles/apple-family/styles.css";
import "@fireback/styles/styles.css";

// You do not have to use the mac-os family theme at all.
// this is the default theme which I use for mac desktop applications
// you could use it as a reference to build your own themes.
// themes are nothing special, rather than wrapping a set of css (scss) on a global name

import { WithFireback } from "../core/WithFireback";

import { useCheckAuthentication } from "@fireback/ui-core/components/layouts/ForcedAuthenticated";
import { BUILD_VARIABLES } from "@fireback/ui-core/hooks/build-variables";
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
