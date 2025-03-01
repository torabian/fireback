import { WorkspaceConfigEntity } from "@/modules/fireback/sdk/modules/workspaces/WorkspaceConfigEntity";
import { Route } from "react-router-dom";
import { WorkspaceConfigEntityManager } from "./WorkspaceConfigEntityManager";
import { WorkspaceConfigSingleScreen } from "./WorkspaceConfigSingleScreen";
export function useWorkspaceConfigRoutes() {
  return (
    <>
      <Route
        element={<WorkspaceConfigSingleScreen />}
        path={"root/workspace/config"}
      ></Route>
      <Route
        element={<WorkspaceConfigEntityManager />}
        path={"root/workspace/config/edit"}
      ></Route>
    </>
  );
}
