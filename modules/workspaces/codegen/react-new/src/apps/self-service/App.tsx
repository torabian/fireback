import { useRef } from "react";
import "bootstrap/dist/css/bootstrap.css";
import "../../modules/fireback/styles/styles.scss";
import "../../modules/fireback/styles/apple-family/styles.scss";

// You do not have to use the mac-os family theme at all.
// this is the default theme which I use for mac desktop applications
// you could use it as a reference to build your own themes.
// themes are nothing special, rather than wrapping a set of css (scss) on a global name

import { WithFireback } from "@/modules/fireback/apps/core/WithFireback";
import {
  useAbacAuthenticatedRoutes,
  useAbacModulePublicRoutes,
  useSelfServiceAuthenticatedRoutes,
} from "@/modules/fireback/modules/AbacModuleRoutes";
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
  const abacModulePublicRoutes = useAbacModulePublicRoutes();
  const selfServiceAuthenticatedRoutes = useSelfServiceAuthenticatedRoutes();
  const { session, checked } = useCheckAuthentication();

  return (
    <>
      {!session && checked ? (
        <Router>
          <Routes>
            <Route path=":locale">{abacModulePublicRoutes}</Route>
            <Route path="*" element={<Navigate to="/en/signin2" replace />} />
          </Routes>
        </Router>
      ) : (
        <Router>
          <Routes>
            <Route path=":locale">{selfServiceAuthenticatedRoutes}</Route>
            <Route
              path="*"
              element={<Navigate to="/en/auth/passports" replace />}
            />
          </Routes>
        </Router>
      )}
    </>
  );
}

export default App;
