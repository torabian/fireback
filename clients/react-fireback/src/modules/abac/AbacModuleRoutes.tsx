import { Route } from "react-router-dom";
import { Signin } from "../authentication/Signin";
import { Signup } from "../authentication/SignupManager";
import { OtpPasswordPrimary } from "../authentication/OtpPasswordPrimary";
import { JoinToWorkspace } from "../authentication/JoinToWorkspace";
import { OtpPassword } from "../authentication/OtpPassword";
import { WorkspaceNotificationEntityManager } from "./workspaces/WorkspaceNotificationEntityManager";
import { WorkspaceInviteEntityManager } from "./workspace-invites/WorkspaceInviteEntityManager";
import { EmailProviderNavigationTools } from "src/sdk/fireback/modules/workspaces/email-provider-navigation-tools";
import { EmailSenderNavigationTools } from "src/sdk/fireback/modules/workspaces/email-sender-navigation-tools";
import { RoleNavigationTools } from "src/sdk/fireback/modules/workspaces/role-navigation-tools";
import { WorkspaceNavigationTools } from "src/sdk/fireback/modules/workspaces/workspace-navigation-tools";
import { WorkspaceTypeNavigationTools } from "src/sdk/fireback/modules/workspaces/workspace-type-navigation-tools";
import { EmailProviderArchiveScreen } from "./mail-providers/MailProviderArchiveScreen";
import { EmailProviderEntityManager } from "./mail-providers/MailProviderEntityManager";
import { EmailProviderSingleScreen } from "./mail-providers/MailProviderSingleScreen";
import { EmailSenderArchiveScreen } from "./mail-senders/EmailSenderArchiveScreen";
import { EmailSenderEntityManager } from "./mail-senders/EmailSenderEntityManager";
import { EmailSenderSingleScreen } from "./mail-senders/EmailSenderSingleScreen";
import { PublicJoinKeyArchiveScreen } from "./public-join-keys/PublicJoinKeyArchiveScreen";
import { PublicJoinKeyEntityManager } from "./public-join-keys/PublicJoinKeyEntityManager";
import { PublicJoinKeySingleScreen } from "./public-join-keys/PublicJoinKeySingleScreen";
import { RoleArchiveScreen } from "./roles/RoleArchiveScreen";
import { RoleEntityManager } from "./roles/RoleEntityManager";
import { RoleSingleScreen } from "./roles/RoleSingleScreen";
import { UserInvitationArchiveScreen } from "./users/UserInvitationArchiveScreen";
import { WorkspaceInviteArchiveScreen } from "./workspace-invites/WorkspaceInviteArchiveScreen";
import { WorkspaceTypeArchiveScreen } from "./workspace-types/WorkspaceTypeArchiveScreen";
import { WorkspaceTypeEntityManager } from "./workspace-types/WorkspaceTypeEntityManager";
import { WorkspaceTypeSingleScreen } from "./workspace-types/WorkspaceTypeSingleScreen";
import { WorkspaceArchiveScreen } from "./workspaces/WorkspaceArchiveScreen";
import { WorkspaceEntityManager } from "./workspaces/WorkspaceEntityManager";
import { WorkspaceSingleScreen } from "./workspaces/WorkspaceSingleScreen";
import { UserNavigationTools } from "src/sdk/fireback/modules/workspaces/user-navigation-tools";
import { UserArchiveScreen } from "./users/UserArchiveScreen";
import { UserEntityManager } from "./users/UserEntityManager";
import { UserSingleScreen } from "./users/UserSingleScreen";
import { SignupTypeSelect } from "../authentication/SignupTypeSelect";

export const useAbacModulePublicRoutes = () => {
  return (
    <>
      <Route path={"signin"} element={<Signin />}></Route>
      <Route path={"signup/team/:joinKey"} element={<Signup />}></Route>
      <Route path={"signup/:workspaceTypeId"} element={<Signup />}></Route>
      <Route element={<SignupTypeSelect />} path={"signup"} />

      <Route path={"auth"} element={<OtpPasswordPrimary />}></Route>
      <Route
        path={"join/:uniqueId"}
        element={<JoinToWorkspace onSuccess={() => {}} />}
      ></Route>
      <Route path={"otp"} element={<OtpPassword />}></Route>
    </>
  );
};

export const useAbacAuthenticatedRoutes = () => {
  return (
    <>
      <Route
        element={<WorkspaceInviteEntityManager />}
        path={"workspace/invite/new"}
      />

      <Route
        element={<WorkspaceInviteEntityManager />}
        path={"workspace/invite/:uniqueId"}
      />
      <Route
        element={<WorkspaceInviteArchiveScreen />}
        path={"workspace-invites"}
      ></Route>

      <Route
        element={<UserInvitationArchiveScreen />}
        path={"invitations"}
      ></Route>

      <Route
        element={<WorkspaceNotificationEntityManager />}
        path={"workspace/config"}
      />
      <Route
        element={<WorkspaceTypeEntityManager />}
        path={WorkspaceTypeNavigationTools.Rcreate}
      />
      <Route
        element={<WorkspaceTypeSingleScreen />}
        path={WorkspaceTypeNavigationTools.Rsingle}
      ></Route>
      <Route
        element={<WorkspaceTypeEntityManager />}
        path={WorkspaceTypeNavigationTools.Redit}
      ></Route>
      <Route
        element={<WorkspaceTypeArchiveScreen />}
        path={WorkspaceTypeNavigationTools.Rquery}
      ></Route>

      <Route
        element={<WorkspaceEntityManager />}
        path={WorkspaceNavigationTools.Rcreate}
      />
      <Route
        element={<WorkspaceSingleScreen />}
        path={WorkspaceNavigationTools.Rsingle}
      ></Route>
      <Route
        element={<WorkspaceEntityManager />}
        path={WorkspaceNavigationTools.Redit}
      ></Route>
      <Route
        element={<WorkspaceArchiveScreen />}
        path={WorkspaceNavigationTools.Rquery}
      ></Route>

      <Route
        element={<PublicJoinKeyEntityManager />}
        path={"publicjoinkey/edit/:uniqueId"}
      />
      <Route
        element={<PublicJoinKeySingleScreen />}
        path={"publicjoinkey/:uniqueId"}
      ></Route>
      <Route
        element={<PublicJoinKeyEntityManager />}
        path={"publicjoinkey/new"}
      ></Route>
      <Route
        element={<PublicJoinKeyArchiveScreen />}
        path={"publicjoinkeys"}
      ></Route>

      <Route
        element={<EmailProviderEntityManager />}
        path={EmailProviderNavigationTools.Rcreate}
      />
      <Route
        element={<EmailProviderSingleScreen />}
        path={EmailProviderNavigationTools.Rsingle}
      ></Route>
      <Route
        element={<EmailProviderEntityManager />}
        path={EmailProviderNavigationTools.Redit}
      ></Route>
      <Route
        element={<EmailProviderArchiveScreen />}
        path={EmailProviderNavigationTools.Rquery}
      ></Route>

      <Route
        element={<EmailSenderEntityManager />}
        path={EmailSenderNavigationTools.Rcreate}
      />
      <Route
        element={<EmailSenderSingleScreen />}
        path={EmailSenderNavigationTools.Rsingle}
      ></Route>
      <Route
        element={<EmailSenderEntityManager />}
        path={EmailSenderNavigationTools.Redit}
      ></Route>
      <Route
        element={<EmailSenderArchiveScreen />}
        path={EmailSenderNavigationTools.Rquery}
      ></Route>

      <Route
        element={<RoleEntityManager />}
        path={RoleNavigationTools.Rcreate}
      />
      <Route
        element={<RoleSingleScreen />}
        path={RoleNavigationTools.Rsingle}
      ></Route>
      <Route
        element={<RoleEntityManager />}
        path={RoleNavigationTools.Redit}
      ></Route>
      <Route
        element={<RoleArchiveScreen />}
        path={RoleNavigationTools.Rquery}
      ></Route>
      <Route
        element={<UserEntityManager />}
        path={UserNavigationTools.Rcreate}
      />
      <Route
        element={<UserSingleScreen />}
        path={UserNavigationTools.Rsingle}
      ></Route>
      <Route
        element={<UserEntityManager />}
        path={UserNavigationTools.Redit}
      ></Route>
      <Route
        element={<UserArchiveScreen />}
        path={UserNavigationTools.Rquery}
      ></Route>
    </>
  );
};
