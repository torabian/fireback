import { Route } from "react-router-dom";
import { RoleEntityManager } from "./RoleEntityManager";
import { RoleEntity } from "@/modules/fireback/sdk/modules/workspaces/RoleEntity";
import { RoleSingleScreen } from "./RoleSingleScreen";
import { RoleArchiveScreen } from "./RoleArchiveScreen";

export function useRoleRoutes() {
  return (
    <>
      <Route
        element={<RoleEntityManager />}
        path={RoleEntity.Navigation.Rcreate}
      />
      <Route
        element={<RoleSingleScreen />}
        path={RoleEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<RoleEntityManager />}
        path={RoleEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<RoleArchiveScreen />}
        path={RoleEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
