import { createEntityNavigation } from "./createEntityNavigation";

export const EmailConfirmationNavigation = createEntityNavigation(
  "email-confirmation",
  "email-confirmations"
);
export const NotificationConfigNavigation = createEntityNavigation(
  "notification-config",
  "notification-configs"
);
export const PassportNavigation = createEntityNavigation(
  "passport",
  "passports"
);
export const PassportMethodNavigation = createEntityNavigation(
  "passport-method",
  "passport-methods"
);
export const PublicJoinKeyNavigation = createEntityNavigation(
  "public-join-key",
  "public-join-keys"
);
export const RegionalContentNavigation = createEntityNavigation(
  "regional-content",
  "regional-contents"
);
export const RoleNavigation = createEntityNavigation("role", "roles");
export const UserNavigation = createEntityNavigation("user", "users");
export const UserWorkspaceNavigation = createEntityNavigation(
  "user-workspace",
  "user-workspaces"
);
export const WorkspaceNavigation = createEntityNavigation(
  "workspace",
  "workspaces"
);
export const WorkspaceConfigNavigation = createEntityNavigation(
  "workspace-config",
  "workspace-configs"
);
export const WorkspaceInviteNavigation = createEntityNavigation(
  "workspace-invite",
  "workspace-invites"
);
export const WorkspaceRoleNavigation = createEntityNavigation(
  "workspace-role",
  "workspace-roles"
);
export const WorkspaceTypeNavigation = createEntityNavigation(
  "workspace-type",
  "workspace-types"
);
export const CapabilityNavigation = createEntityNavigation(
  "capability",
  "capabilities"
);
