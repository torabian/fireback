/**
 * Fireback manage routes,
 * It's for administration a root level content.
 * Some components can be used for sub-level workspaces, but this is not planned yet
 *
 * All routes regarding manage are authenticated, they do not expose public components.
 */

import { Route } from "react-router-dom";
import { useWorkspaceConfigRoutes } from "./workspace-config/WorkspaceConfigRoutes";
import { usePassportMethodRoutes } from "./passport-method/PassportMethodRoutes";
import { useCapabilityRoutes } from "./capabilities/CapabilityRoutes";
import { useWorkspaceTypeRoutes } from "./workspace-types/WorkspaceTypeRoutes";
import { useWorkspaceRoutes } from "./workspaces/WorkspaceRoutes";
import { useDriveRoutes } from "./drive/DriveRoutes";
import { useEmailProviderRoutes } from "./mail-providers/EmailProviderRoutes";
import { useEmailSenderRoutes } from "./mail-senders/EmailSenderRoutes";
import { useUserRoutes } from "./users/UserRoutes";

export function useManageRoutes() {
  const capabilityRoutes = useCapabilityRoutes();
  const driveRoutes = useDriveRoutes();
  const mailProviderRoutes = useEmailProviderRoutes();
  const mailSenderRoutes = useEmailSenderRoutes();
  const passportMethodRoutes = usePassportMethodRoutes();
  const userRoutes = useUserRoutes();
  const workspaceConfigRoutes = useWorkspaceConfigRoutes();
  const workspaceTypeRoutes = useWorkspaceTypeRoutes();
  const workspaceRoutes = useWorkspaceRoutes();

  return (
    <Route path="manage">
      {capabilityRoutes}
      {driveRoutes}
      {mailProviderRoutes}
      {mailSenderRoutes}
      {passportMethodRoutes}
      {userRoutes}
      {workspaceConfigRoutes}
      {workspaceTypeRoutes}
      {workspaceRoutes}
    </Route>
  );
}
