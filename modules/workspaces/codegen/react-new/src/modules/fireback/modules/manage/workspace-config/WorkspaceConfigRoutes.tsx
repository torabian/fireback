import { Route } from "react-router-dom";
import { WorkspaceConfigEntityManager } from "./WorkspaceConfigEntityManager";
import { WorkspaceConfigSingleScreen } from "./WorkspaceConfigSingleScreen";
export function useWorkspaceConfigRoutes() {
  return (
    <>
      <Route
        element={<WorkspaceConfigSingleScreen />}
        path={"workspace-config"}
      ></Route>
      <Route
        element={<WorkspaceConfigEntityManager />}
        path={"workspace-config/edit"}
      ></Route>
    </>
  );
}
