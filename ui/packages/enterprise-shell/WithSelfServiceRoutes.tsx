import { BrowserRouter, HashRouter, Routes } from "react-router-dom";

import { useAuthentication } from "@fireback/auth-client";
import { SelectWorkspaceScreen } from "@fireback/selfservice/SelectWorkspace.screen";
import { useSelfServicePublicRoutes } from "@fireback/selfservice/SelfServiceRoutes";
import { useCheckAuthentication } from "@fireback/ui-core/components/layouts/ForcedAuthenticated";
import { SessionGate } from "@fireback/ui-core/components/session-gate/SessionGate";
import { checkSessionViaWhoami } from "@fireback/ui-core/components/session-gate/checkSessionViaWhoami";
import { BUILD_VARIABLES } from "@fireback/ui-core/hooks/build-variables";
import { useWasmFetchOverride } from "@fireback/wasm-server/WasmFetchContext";
import { type ReactNode } from "react";

const useHashRouter = BUILD_VARIABLES.USE_HASH_ROUTER === "true";
const Router = useHashRouter ? HashRouter : BrowserRouter;

export const WithSelfServiceRoutes = ({
  children,
}: {
  children: ReactNode;
}) => {
  const { session, checked } = useCheckAuthentication();
  const selfServicePublicRoutes = useSelfServicePublicRoutes();
  const { selectedWorkspace } = useAuthentication();
  // undefined outside of an app that mounts WithWasmServer (see
  // WasmFetchContext.ts) - checkSessionViaWhoami then falls back to a real
  // network fetch, unchanged from before.
  const wasmFetchOverride = useWasmFetchOverride();

  // AuthenticationSession.workspaces is already a plain array (see
  // mapRawSessionToAuthenticationSession) - no MCollection/localStorage
  // round-trip shape-juggling needed here anymore.
  const userWorkspaceCount = session?.workspaces?.length ?? 0;

  // Unauthenticated: self-service's own public routes (welcome/signup/signin)
  // render here directly, deliberately outside SessionGate below - there's no
  // session yet to verify, and gating an anonymous visitor's first paint
  // behind a whoami check would just be a pointless spinner before the exact
  // screen that lets them get a session in the first place.
  if (!session && checked) {
    return (
      <Router future={{ v7_startTransition: true }}>
        <Routes>{selfServicePublicRoutes}</Routes>
      </Router>
    );
  }

  // From here on a session exists (at least optimistically, from storage) -
  // SessionGate confirms it's actually still valid via whoami before anything
  // that assumes real authenticated access renders: the workspace picker, or
  // the main app itself.
  return (
    <SessionGate
      checkSession={() => checkSessionViaWhoami(wasmFetchOverride)}
    >
      {!selectedWorkspace && userWorkspaceCount > 1 ? (
        <Router future={{ v7_startTransition: true }}>
          <SelectWorkspaceScreen />
        </Router>
      ) : (
        children
      )}
    </SessionGate>
  );
};
