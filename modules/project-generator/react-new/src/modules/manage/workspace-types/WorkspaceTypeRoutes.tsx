import { WorkspaceTypeNavigation } from "../../sdk/navigation/AbacNavigation";
import { WorkspaceTypeEntityManager } from "./WorkspaceTypeEntityManager";
import { Route } from "react-router-dom";
import { WorkspaceTypeArchiveScreen } from "./WorkspaceTypeArchiveScreen";
import { WorkspaceTypeSingleScreen } from "./WorkspaceTypeSingleScreen";

export function useWorkspaceTypeRoutes() {
  return (
    <>
      <Route
        element={<WorkspaceTypeEntityManager />}
        path={WorkspaceTypeNavigation.Rcreate}
      />
      <Route
        element={<WorkspaceTypeEntityManager />}
        path={WorkspaceTypeNavigation.Redit}
      />
      <Route
        element={<WorkspaceTypeSingleScreen />}
        path={WorkspaceTypeNavigation.Rsingle}
      />
      <Route
        element={<WorkspaceTypeArchiveScreen />}
        path={WorkspaceTypeNavigation.Rquery}
      ></Route>
    </>
  );
}
