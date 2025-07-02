import {
  BrowserRouter,
  HashRouter,
  Navigate,
  Route,
  Routes,
} from "react-router-dom";

import { ReactNode } from "react";
import { useCheckAuthentication } from "../../components/layouts/ForcedAuthenticated";
import { useSelfServicePublicRoutes } from "../../modules/selfservice/SelfServiceRoutes";

const useHashRouter = process.env.REACT_APP_USE_HASH_ROUTER === "true";
const Router = useHashRouter ? HashRouter : BrowserRouter;

export const WithSelfServiceRoutes = ({
  children,
}: {
  children: ReactNode;
}) => {
  const { session, checked } = useCheckAuthentication();
  const selfServicePublicRoutes = useSelfServicePublicRoutes();

  if (!session && checked) {
    return (
      <Router future={{ v7_startTransition: true }}>
        <Routes>
          <Route path=":locale">{selfServicePublicRoutes}</Route>
          <Route
            path="*"
            element={<Navigate to="/en/selfservice/welcome" replace />}
          />
        </Routes>
      </Router>
    );
  }

  return <>{children}</>;
};
