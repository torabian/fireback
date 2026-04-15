import {
  BrowserRouter,
  HashRouter,
  Navigate,
  Route,
  Routes,
} from "react-router-dom";

import { type ReactNode, useContext, useEffect } from "react";
import { useCheckAuthentication } from "../../components/layouts/ForcedAuthenticated";
import { BUILD_VARIABLES } from "../../hooks/build-variables";
import { SelectWorkspaceScreen } from "../../modules/selfservice/SelectWorkspace.screen";
import { useSelfServicePublicRoutes } from "../../modules/selfservice/SelfServiceRoutes";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { useQueryUserRoleWorkspacesActionQuery } from "../../sdk/modules/abac/QueryUserRoleWorkspaces";

const useHashRouter = BUILD_VARIABLES.USE_HASH_ROUTER === "true";
const Router = useHashRouter ? HashRouter : BrowserRouter;

export const WithSelfServiceRoutes = ({
  children,
}: {
  children: ReactNode;
}) => {
  const { session, checked } = useCheckAuthentication();
  const selfServicePublicRoutes = useSelfServicePublicRoutes();
  const { selectedUrw, selectUrw } = useContext(RemoteQueryContext);

  const queryUrw = useQueryUserRoleWorkspacesActionQuery({
    cacheTime: 50,
    enabled: false
  });

  useEffect(() => {
    if ((session as any)?.userWorkspaces?.length === 1 && !selectedUrw) {
      queryUrw.refetch().then((resp) => {
        const items = resp?.data?.data?.items || [];
        if (items.length !== 1) {
          return;
        }

        selectUrw({
          roleId: items[0].roles?.[0]?.uniqueId,
          workspaceId: items[0].uniqueId,
        });
      });
    }
  }, [selectedUrw, session]);

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

  if (!selectedUrw && (session as any)?.userWorkspaces?.length > 1) {
    return (
      <Router future={{ v7_startTransition: true }}>
        <SelectWorkspaceScreen />
      </Router>
    );
  }

  return <>{children}</>;
};
