import { Route } from "react-router-dom";
import { Signin } from "./auth/Signin";
import { Signup } from "./auth/SignupManager";
import { OtpPasswordPrimary } from "./auth/OtpPasswordPrimary";
import { JoinToWorkspace } from "./auth/JoinToWorkspace";
import { OtpPassword } from "./auth/OtpPassword";
// import { WorkspaceNotificationEntityManager } from "./workspaces/WorkspaceNotificationEntityManager";
import { WorkspaceInviteEntityManager } from "./workspace-invites/WorkspaceInviteEntityManager";
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
import { UserArchiveScreen } from "./users/UserArchiveScreen";
import { UserEntityManager } from "./users/UserEntityManager";
import { UserSingleScreen } from "./users/UserSingleScreen";
import { SignupTypeSelect } from "./auth/SignupTypeSelect";
import { WorkspaceTypeEntity } from "../sdk/modules/workspaces/WorkspaceTypeEntity";
import { WorkspaceEntity } from "../sdk/modules/workspaces/WorkspaceEntity";
import { EmailProviderEntity } from "../sdk/modules/workspaces/EmailProviderEntity";
import { EmailSenderEntity } from "../sdk/modules/workspaces/EmailSenderEntity";
import { RoleEntity } from "../sdk/modules/workspaces/RoleEntity";
import { UserEntity } from "../sdk/modules/workspaces/UserEntity";
import { PassportEntityManager } from "./passports/PassportEntityManager";

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
