import { Navigate, Route } from "react-router-dom";
// import { WorkspaceNotificationEntityManager } from "./workspaces/WorkspaceNotificationEntityManager";
import { EmailProviderEntity } from "../sdk/modules/workspaces/EmailProviderEntity";
import { EmailSenderEntity } from "../sdk/modules/workspaces/EmailSenderEntity";
import { RoleEntity } from "../sdk/modules/workspaces/RoleEntity";
import { UserEntity } from "../sdk/modules/workspaces/UserEntity";
import { WorkspaceEntity } from "../sdk/modules/workspaces/WorkspaceEntity";
import { WorkspaceTypeEntity } from "../sdk/modules/workspaces/WorkspaceTypeEntity";
import { AuthMethod } from "./auth2/auth.common";
import { ClassicPassportScreen } from "./auth2/ClassicPassport.screen";
import { ClassicPassportAccountCreation } from "./auth2/ClassicPassportAccountCreation.screen";
import { ClassicSigninPassword } from "./auth2/ClassicSigninPassword.screen";
import { OtpScreen } from "./auth2/Otp.screen";
import { WelcomeScreen } from "./auth2/Welcome.screen";
import { EmailProviderArchiveScreen } from "./mail-providers/MailProviderArchiveScreen";
import { EmailProviderEntityManager } from "./mail-providers/MailProviderEntityManager";
import { EmailProviderSingleScreen } from "./mail-providers/MailProviderSingleScreen";
import { EmailSenderArchiveScreen } from "./mail-senders/EmailSenderArchiveScreen";
import { EmailSenderEntityManager } from "./mail-senders/EmailSenderEntityManager";
import { EmailSenderSingleScreen } from "./mail-senders/EmailSenderSingleScreen";
import { PassportEntityManager } from "./passports/PassportEntityManager";
import { PublicJoinKeyArchiveScreen } from "./public-join-keys/PublicJoinKeyArchiveScreen";
import { PublicJoinKeyEntityManager } from "./public-join-keys/PublicJoinKeyEntityManager";
import { PublicJoinKeySingleScreen } from "./public-join-keys/PublicJoinKeySingleScreen";
import { RoleArchiveScreen } from "./roles/RoleArchiveScreen";
import { RoleEntityManager } from "./roles/RoleEntityManager";
import { RoleSingleScreen } from "./roles/RoleSingleScreen";
import { UserArchiveScreen } from "./users/UserArchiveScreen";
import { UserEntityManager } from "./users/UserEntityManager";
import { UserInvitationArchiveScreen } from "./users/UserInvitationArchiveScreen";
import { UserSingleScreen } from "./users/UserSingleScreen";
import { WorkspaceInviteArchiveScreen } from "./workspace-invites/WorkspaceInviteArchiveScreen";
import { WorkspaceInviteEntityManager } from "./workspace-invites/WorkspaceInviteEntityManager";
import { WorkspaceTypeArchiveScreen } from "./workspace-types/WorkspaceTypeArchiveScreen";
import { WorkspaceTypeEntityManager } from "./workspace-types/WorkspaceTypeEntityManager";
import { WorkspaceTypeSingleScreen } from "./workspace-types/WorkspaceTypeSingleScreen";
import { WorkspaceArchiveScreen } from "./workspaces/WorkspaceArchiveScreen";
import { WorkspaceEntityManager } from "./workspaces/WorkspaceEntityManager";
import { WorkspaceSingleScreen } from "./workspaces/WorkspaceSingleScreen";
import { TotpSetup } from "./auth2/TotpSetup.screen";
import { TotpEnter } from "./auth2/TotpEnter.screen";
import { useWorkspaceConfigRoutes } from "./root/workspace-config/WorkspaceConfigRoutes";
import { ChangePasswordScreen } from "./auth2/ChangePassword.screen";

export const useAbacModulePublicRoutes = () => {
  return (
    <>
      <Route path={"welcome"} element={<WelcomeScreen />}></Route>
      <Route
        path={"auth/email"}
        element={<ClassicPassportScreen method={AuthMethod.Email} />}
      ></Route>
      <Route
        path={"auth/phone"}
        element={<ClassicPassportScreen method={AuthMethod.Phone} />}
      ></Route>

      <Route path={"auth/totp-setup"} element={<TotpSetup />}></Route>
      <Route path={"auth/totp-enter"} element={<TotpEnter />}></Route>
      <Route
        path={"auth/complete"}
        element={<ClassicPassportAccountCreation />}
      ></Route>
      <Route path={"auth/password"} element={<ClassicSigninPassword />}></Route>
      <Route path={"auth/otp"} element={<OtpScreen />}></Route>

      <Route path="*" element={<Navigate to="/en/welcome" replace />} />
    </>
  );
};

export const useAbacAuthenticatedRoutes = () => {
  const configWorkspaces = useWorkspaceConfigRoutes();

  return (
    <>
      {configWorkspaces}
      <Route
        element={<WorkspaceInviteEntityManager />}
        path={"workspace/invite/new"}
      />

      <Route
        path={"auth/change-password"}
        element={<ChangePasswordScreen />}
      ></Route>

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

      <Route element={<PassportEntityManager />} path={"passport"}></Route>

      {/* <Route
        element={<WorkspaceNotificationEntityManager />}
        path={"workspace/config"}
      /> */}
      <Route
        element={<WorkspaceTypeEntityManager />}
        path={WorkspaceTypeEntity.Navigation.Rcreate}
      />
      <Route
        element={<WorkspaceTypeSingleScreen />}
        path={WorkspaceTypeEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<WorkspaceTypeEntityManager />}
        path={WorkspaceTypeEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<WorkspaceTypeArchiveScreen />}
        path={WorkspaceTypeEntity.Navigation.Rquery}
      ></Route>

      <Route
        element={<WorkspaceEntityManager />}
        path={WorkspaceEntity.Navigation.Rcreate}
      />
      <Route
        element={<WorkspaceSingleScreen />}
        path={WorkspaceEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<WorkspaceEntityManager />}
        path={WorkspaceEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<WorkspaceArchiveScreen />}
        path={WorkspaceEntity.Navigation.Rquery}
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
        path={EmailProviderEntity.Navigation.Rcreate}
      />
      <Route
        element={<EmailProviderSingleScreen />}
        path={EmailProviderEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<EmailProviderEntityManager />}
        path={EmailProviderEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<EmailProviderArchiveScreen />}
        path={EmailProviderEntity.Navigation.Rquery}
      ></Route>

      <Route
        element={<EmailSenderEntityManager />}
        path={EmailSenderEntity.Navigation.Rcreate}
      />
      <Route
        element={<EmailSenderSingleScreen />}
        path={EmailSenderEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<EmailSenderEntityManager />}
        path={EmailSenderEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<EmailSenderArchiveScreen />}
        path={EmailSenderEntity.Navigation.Rquery}
      ></Route>

      <Route
        element={<RoleEntityManager />}
        path={RoleEntity.Navigation.Rcreate}
      />
      <Route
        element={<RoleSingleScreen />}
        path={RoleEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<RoleEntityManager />}
        path={RoleEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<RoleArchiveScreen />}
        path={RoleEntity.Navigation.Rquery}
      ></Route>
      <Route
        element={<UserEntityManager />}
        path={UserEntity.Navigation.Rcreate}
      />
      <Route
        element={<UserSingleScreen />}
        path={UserEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<UserEntityManager />}
        path={UserEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<UserArchiveScreen />}
        path={UserEntity.Navigation.Rquery}
      ></Route>
    </>
  );
};
