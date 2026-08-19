/**
 * Fireback manage routes,
 * It's for administration a root level content.
 * Some components can be used for sub-level workspaces, but this is not planned yet
 *
 * All routes regarding manage are authenticated, they do not expose public components.
 */

import { Route } from "react-router-dom";
import { useAnalyticsRoutes } from "./analytics/AnalyticsRoutes";
import { useCapabilityRoutes } from "./capabilities/CapabilityRoutes";
import { useEmailProviderRoutes } from "@fireback/messaging/mail-providers/EmailProviderRoutes";
import { useEmailSenderRoutes } from "@fireback/messaging/mail-senders/EmailSenderRoutes";
import { useInternalStatsRoutes } from "./internal-stats/InternalStatsRoutes";
import { useNotificationsRoutes } from "./notifications/NotificationsRoutes";
import { usePassportMethodRoutes } from "./passport-method/PassportMethodRoutes";
import { useRegionalContentRoutes } from "./regional-content/RegionalContentRoutes";
import { useUserRoutes } from "./users/UserRoutes";
import { useWorkspaceConfigRoutes } from "./workspace-config/WorkspaceConfigRoutes";
import { useWorkspaceTypeRoutes } from "./workspace-types/WorkspaceTypeRoutes";
import { useWorkspaceRoutes } from "./workspaces/WorkspaceRoutes";
import { useGsmProviderRoutes } from "@fireback/messaging/gsm-provider/GsmProviderRoutes";

export function useManageRoutes() {
  const analyticsRoutes = useAnalyticsRoutes();
  const capabilityRoutes = useCapabilityRoutes();
  const mailProviderRoutes = useEmailProviderRoutes();
  const gsmProviderRoutes = useGsmProviderRoutes();
  const mailSenderRoutes = useEmailSenderRoutes();
  const internalStatsRoutes = useInternalStatsRoutes();
  const notificationsRoutes = useNotificationsRoutes();
  const passportMethodRoutes = usePassportMethodRoutes();
  const userRoutes = useUserRoutes();
  const workspaceConfigRoutes = useWorkspaceConfigRoutes();
  const workspaceTypeRoutes = useWorkspaceTypeRoutes();
  const workspaceRoutes = useWorkspaceRoutes();
  const regionalContentRoutes = useRegionalContentRoutes();

  return (
    <Route path="manage">
      {analyticsRoutes}
      {capabilityRoutes}
      {mailProviderRoutes}
      {mailSenderRoutes}
      {internalStatsRoutes}
      {notificationsRoutes}
      {passportMethodRoutes}
      {userRoutes}
      {gsmProviderRoutes}
      {workspaceConfigRoutes}
      {workspaceTypeRoutes}
      {workspaceRoutes}
      {regionalContentRoutes}
    </Route>
  );
}
