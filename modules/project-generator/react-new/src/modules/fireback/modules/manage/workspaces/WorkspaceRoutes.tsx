import { WorkspaceNavigation } from "../../../sdk/navigation/AbacNavigation";
import { WorkspaceEntityManager } from "./WorkspaceEntityManager";
import { Route } from "react-router-dom";
import { WorkspaceArchiveScreen } from "./WorkspaceArchiveScreen";
import { WorkspaceSingleScreen } from "./WorkspaceSingleScreen";

export function useWorkspaceRoutes() {
  return (
    <>
      <Route
        element={<WorkspaceEntityManager />}
        path={WorkspaceNavigation.Rcreate}
      />
      <Route
        element={<WorkspaceEntityManager />}
        path={WorkspaceNavigation.Redit}
      />
      <Route
        element={<WorkspaceSingleScreen />}
        path={WorkspaceNavigation.Rsingle}
      />
      <Route
        element={<WorkspaceArchiveScreen />}
        path={WorkspaceNavigation.Rquery}
      ></Route>
    </>
  );
}
