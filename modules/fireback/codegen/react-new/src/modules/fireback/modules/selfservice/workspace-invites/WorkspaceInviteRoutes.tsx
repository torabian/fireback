import { WorkspaceInviteEntity } from "@/modules/fireback/sdk/modules/abac/WorkspaceInviteEntity";
import { WorkspaceInviteEntityManager } from "./WorkspaceInviteEntityManager";
import { Route } from "react-router-dom";
import { WorkspaceInviteSingleScreen } from "./WorkspaceInviteScreen";
import { WorkspaceInviteArchiveScreen } from "./WorkspaceInviteArchiveScreen";

export function useWorkspaceInviteRoutes() {
  return (
    <>
      <Route
        element={<WorkspaceInviteEntityManager />}
        path={WorkspaceInviteEntity.Navigation.Rcreate}
      />
      <Route
        element={<WorkspaceInviteEntityManager />}
        path={WorkspaceInviteEntity.Navigation.Redit}
      />
      <Route
        element={<WorkspaceInviteSingleScreen />}
        path={WorkspaceInviteEntity.Navigation.Rsingle}
      />
      <Route
        element={<WorkspaceInviteArchiveScreen />}
        path={WorkspaceInviteEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
