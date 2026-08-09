import {
  BrowserRouter,
  HashRouter,
  Navigate,
  Route,
  Routes,
} from "react-router-dom";

import { type ReactNode, useContext, useEffect } from "react";
import { useCheckAuthentication } from "../../modules/fireback-ui/components/layouts/ForcedAuthenticated";
import { BUILD_VARIABLES } from "../../modules/fireback-ui/hooks/build-variables";
import { SelectWorkspaceScreen } from "../../modules/selfservice/SelectWorkspace.screen";
import { useSelfServicePublicRoutes } from "../../modules/selfservice/SelfServiceRoutes";
import {
  RemoteQueryContext,
  collectionToArray,
} from "../../modules/sdk/core/react-tools";
import { useQueryUserRoleWorkspacesActionQuery } from "../../modules/sdk/abac/QueryUserRoleWorkspacesAction";
import { SessionGate } from "@/modules/fireback-ui/components/session-gate/SessionGate";
import { checkSessionViaWhoami } from "@/modules/fireback-ui/components/session-gate/checkSessionViaWhoami";

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
    enabled: false,
  });

  // userWorkspaces' actual shape varies at runtime - a real MCollection
  // right after signin/signup, or a plain array/tagged object once it's been
  // through the JSON round-trip react-tools.tsx's saveSession/getSession use
  // for localStorage (see collectionToArray's doc comment there). Assuming
  // just one shape either crashed (`.len is not a function`) or silently
  // matched nothing (`.length`/`.length` on an MCollection is always
  // undefined) - both disabled this auto-select effect and the
  // multi-workspace picker below for every session, leaving no workspace-id
  // ever selected.
  const userWorkspaceCount = collectionToArray(
    (session as any)?.userWorkspaces,
  ).length;

  useEffect(() => {
    if (userWorkspaceCount === 1 && !selectedUrw) {
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

  // Unauthenticated: self-service's own public routes (welcome/signup/signin)
  // render here directly, deliberately outside SessionGate below - there's no
  // session yet to verify, and gating an anonymous visitor's first paint
  // behind a whoami check would just be a pointless spinner before the exact
  // screen that lets them get a session in the first place.
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

  // From here on a session exists (at least optimistically, from storage) -
  // SessionGate confirms it's actually still valid via whoami before anything
  // that assumes real authenticated access renders: the workspace picker, or
  // the main app itself.
  return (
    <SessionGate checkSession={checkSessionViaWhoami}>
      {!selectedUrw && userWorkspaceCount > 1 ? (
        <Router future={{ v7_startTransition: true }}>
          <SelectWorkspaceScreen />
        </Router>
      ) : (
        children
      )}
    </SessionGate>
  );
};
