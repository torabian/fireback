import { WorkspaceTypeEntity } from "@/modules/fireback/sdk/modules/workspaces/WorkspaceTypeEntity";
import { WorkspaceTypeEntityManager } from "./WorkspaceTypeEntityManager";
import { Route } from "react-router-dom";
import { WorkspaceTypeArchiveScreen } from "./WorkspaceTypeArchiveScreen";
import { WorkspaceTypeSingleScreen } from "./WorkspaceTypeSingleScreen";

export function useWorkspaceTypeRoutes() {
  return (
    <>
      <Route
        element={<WorkspaceTypeEntityManager />}
        path={WorkspaceTypeEntity.Navigation.Rcreate}
      />
      <Route
        element={<WorkspaceTypeEntityManager />}
        path={WorkspaceTypeEntity.Navigation.Redit}
      />
      <Route
        element={<WorkspaceTypeSingleScreen />}
        path={WorkspaceTypeEntity.Navigation.Rsingle}
      />
      <Route
        element={<WorkspaceTypeArchiveScreen />}
        path={WorkspaceTypeEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
