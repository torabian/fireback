import { useRef } from "react";
import "bootstrap/dist/css/bootstrap.css";
import "../../modules/fireback/styles/styles.scss";
import "../../modules/fireback/styles/apple-family/styles.scss";

// You do not have to use the mac-os family theme at all.
// this is the default theme which I use for mac desktop applications
// you could use it as a reference to build your own themes.
// themes are nothing special, rather than wrapping a set of css (scss) on a global name

import { WithFireback } from "@/modules/fireback/apps/core/WithFireback";

import { QueryClient, QueryClientProvider } from "react-query";
import {
  BrowserRouter,
  HashRouter,
  Navigate,
  Route,
  Routes,
} from "react-router-dom";
import { FirebackMockServer } from "./mockServer";
import { useCheckAuthentication } from "@/modules/fireback/components/layouts/ForcedAuthenticated";
import {
  useSelfServiceAuthenticateRoutes,
  useSelfServicePublicRoutes,
} from "@/modules/fireback/modules/selfservice/SelfServiceRoutes";

const useHashRouter = process.env.REACT_APP_USE_HASH_ROUTER === "true";
const Router = useHashRouter ? HashRouter : BrowserRouter;

function App() {
  const mockServer = useRef<any>(FirebackMockServer);

  const queryClient = useRef(new QueryClient());

  return (
    <QueryClientProvider client={queryClient.current}>
      <WithFireback
        mockServer={mockServer}
        config={{}}
        prefix={""}
        queryClient={queryClient.current}
      >
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
              element={<Navigate to="/en/selftservice" replace />}
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
