import { WorkspaceInviteNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
import { WorkspaceInviteEntityManager } from "./WorkspaceInviteEntityManager";
import { Route } from "react-router-dom";
import { WorkspaceInviteSingleScreen } from "./WorkspaceInviteScreen";
import { WorkspaceInviteArchiveScreen } from "./WorkspaceInviteArchiveScreen";

export function useWorkspaceInviteRoutes() {
  return (
    <>
      <Route
        element={<WorkspaceInviteEntityManager />}
        path={WorkspaceInviteNavigation.Rcreate}
      />
      <Route
        element={<WorkspaceInviteEntityManager />}
        path={WorkspaceInviteNavigation.Redit}
      />
      <Route
        element={<WorkspaceInviteSingleScreen />}
        path={WorkspaceInviteNavigation.Rsingle}
      />
      <Route
        element={<WorkspaceInviteArchiveScreen />}
        path={WorkspaceInviteNavigation.Rquery}
      ></Route>
    </>
  );
}
