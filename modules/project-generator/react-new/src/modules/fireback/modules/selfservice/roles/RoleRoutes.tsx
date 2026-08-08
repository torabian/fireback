import { Route } from "react-router-dom";
import { RoleEntityManager } from "./RoleEntityManager";
import { RoleNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
import { RoleSingleScreen } from "./RoleSingleScreen";
import { RoleArchiveScreen } from "./RoleArchiveScreen";

export function useRoleRoutes() {
  return (
    <>
      <Route
        element={<RoleEntityManager />}
        path={RoleNavigation.Rcreate}
      />
      <Route
        element={<RoleSingleScreen />}
        path={RoleNavigation.Rsingle}
      ></Route>
      <Route
        element={<RoleEntityManager />}
        path={RoleNavigation.Redit}
      ></Route>
      <Route
        element={<RoleArchiveScreen />}
        path={RoleNavigation.Rquery}
      ></Route>
    </>
  );
}
