// ~ auto:useMockImport

import {DriveMockServer} from './mock/api/drive';
import {EmailProviderMockServer} from './mock/api/emailprovider';
import {EmailSenderMockServer} from './mock/api/emailsender';
import {PublicJoinKeyMockServer} from './mock/api/public-join-key';
import {RoleMockServer} from './mock/api/roles';
import {SidebarMockServer} from './mock/api/sidebar';
import {UserMockServer} from './mock/api/users';
import {WorkspaceConfigMockServer} from './mock/api/workspace-config';
import {WorkspaceInviteMockServer} from './mock/api/workspace-invites';
import {WorkspaceTypeMockServer} from './mock/api/workspace-type';
import {WorkspaceMockServer} from './mock/api/workspaces';
import {AuthMockServer} from './mock/api/auth';

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
  // ~ auto:useMocknew
];
