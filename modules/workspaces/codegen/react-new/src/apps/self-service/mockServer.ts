// ~ auto:useMockImport

import { AuthMockServer } from "@/modules/fireback/mock/api/auth";
import { DriveMockServer } from "@/modules/fireback/mock/api/drive";
import { EmailProviderMockServer } from "@/modules/fireback/mock/api/emailprovider";
import { EmailSenderMockServer } from "@/modules/fireback/mock/api/emailsender";
import { PublicJoinKeyMockServer } from "@/modules/fireback/mock/api/public-join-key";
import { RoleMockServer } from "@/modules/fireback/mock/api/roles";
import { SidebarMockServer } from "@/modules/fireback/mock/api/sidebar";
import { UserMockServer } from "@/modules/fireback/mock/api/users";
import { WorkspaceConfigMockServer } from "@/modules/fireback/mock/api/workspace-config";
import { WorkspaceInviteMockServer } from "@/modules/fireback/mock/api/workspace-invites";
import { WorkspaceTypeMockServer } from "@/modules/fireback/mock/api/workspace-type";
import { WorkspaceMockServer } from "@/modules/fireback/mock/api/workspaces";
import { WorkspaceConfigMockProvider } from "@/modules/fireback/modules/root/workspace-config/WorkspaceConfigMockProvider";

export const FirebackMockServer = [
  new AuthMockServer(),
  new RoleMockServer(),
  new SidebarMockServer(),
  new UserMockServer(),
  new WorkspaceTypeMockServer(),
  new DriveMockServer(),
  new EmailProviderMockServer(),
  new EmailSenderMockServer(),
  new WorkspaceInviteMockServer(),
  new PublicJoinKeyMockServer(),
  new WorkspaceConfigMockServer(),
  new WorkspaceMockServer(),
  new WorkspaceConfigMockProvider(),
  // ~ auto:useMocknew
];
