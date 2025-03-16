import { WorkspaceEntity } from "@/modules/fireback/sdk/modules/workspaces/WorkspaceEntity";
import { WorkspaceEntityManager } from "./WorkspaceEntityManager";
import { Route } from "react-router-dom";
import { WorkspaceArchiveScreen } from "./WorkspaceArchiveScreen";
import { WorkspaceSingleScreen } from "./WorkspaceSingleScreen";

export function useWorkspaceRoutes() {
  return (
    <>
      <Route
        element={<WorkspaceEntityManager />}
        path={WorkspaceEntity.Navigation.Rcreate}
      />
      <Route
        element={<WorkspaceEntityManager />}
        path={WorkspaceEntity.Navigation.Redit}
      />
      <Route
        element={<WorkspaceSingleScreen />}
        path={WorkspaceEntity.Navigation.Rsingle}
      />
      <Route
        element={<WorkspaceArchiveScreen />}
        path={WorkspaceEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
