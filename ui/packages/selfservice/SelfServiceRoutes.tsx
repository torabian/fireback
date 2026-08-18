/**
 * Self service allows to connect any application with Fireback user flow,
 * invitations, permissions etc.
 * Components here will be used in a separate project, and other apps can redirect
 * or open these routes in a webview or iframe.
 * Note that manage also might use components from selfservice, the root login is using
 * the same flow and components.
 */

import { AnimatedRouteWrapper } from "@fireback/ui-core/components/swipe-transition/SwipeTransition";
import { usePureLocale } from "@fireback/ui-core/hooks/usePureLocale";
import { Navigate, Route } from "react-router-dom";
import { AuthMethod } from "./auth.common";
import { ChangePasswordScreen } from "./ChangePassword.screen";
import { ClassicPassportScreen } from "./ClassicPassport.screen";
import { ClassicPassportAccountCreation } from "./ClassicPassportAccountCreation.screen";
import { ClassicSigninPassword } from "./ClassicSigninPassword.screen";
import { OtpScreen } from "./Otp.screen";
import { ResetPasswordScreen } from "./ResetPassword.screen";
import { usePublicJoinKeyRoutes } from "./public-join-keys/PublicJoinKeyRoutes";
import { useRoleRoutes } from "./roles/RoleRoutes";
import { SelfServiceHome } from "./SelfServiceHome";
import { TotpEnter } from "./TotpEnter.screen";
import { TotpSetup } from "./TotpSetup.screen";
import { useUserInvitationRoutes } from "./user-invitations/UserInvitationRoutes";
import { UserPassportsScreen } from "./UserPassports.screen";
import { WelcomeScreen } from "./Welcome.screen";
import { useWorkspaceInviteRoutes } from "./workspace-invites/WorkspaceInviteRoutes";

/**
 * Public routes are those which do not require user to be authenticate,
 * or might be. Such as login form, etc.
 */
export function useSelfServicePublicRoutes() {
  const { locale } = usePureLocale();

  return (
    <>
      <Route path="selfservice">
        <Route path={"welcome"} element={<WelcomeScreen />}></Route>
        <Route
          path={"email"}
          element={<ClassicPassportScreen method={AuthMethod.Email} />}
        ></Route>
        <Route
          path={"phone"}
          element={<ClassicPassportScreen method={AuthMethod.Phone} />}
        ></Route>

        <Route path={"totp-setup"} element={<TotpSetup />}></Route>
        <Route path={"totp-enter"} element={<TotpEnter />}></Route>
        <Route
          path={"complete"}
          element={<ClassicPassportAccountCreation />}
        ></Route>
        <Route path={"password"} element={<ClassicSigninPassword />}></Route>

        <Route path={"otp"} element={<OtpScreen />}></Route>

        {/* Reached from an emailed/texted link (ClassicPassportRequestOtpAction /
            SendPassportResetEmailAction - see this route's own ?value= query param
            handling in ResetPassword.presenter.tsx), not from in-app navigation - and
            registered again, identically, under useSelfServiceAuthenticateRoutes below,
            since App.tsx picks whichever of these two route tables to mount based on
            whether *this browser* already has an active session at all, which has
            nothing to do with whose passport the link is actually for (an admin
            testing a link, or a shared/kiosk browser already signed in as someone
            else, must still land on this screen rather than being silently bounced to
            their own /selfservice/passports by the authenticated table's catch-all). */}
        <Route path={"reset-password"} element={<ResetPasswordScreen />}></Route>
      </Route>

      <Route
        path="*"
        element={<Navigate to={`/${locale}/selfservice/welcome`} replace />}
      />
    </>
  );
}

/**
 * Routes that require user to be authenticated and session to be active,
 * such change change password, etc.
 */
export function useSelfServiceAuthenticateRoutes() {
  const publicJoinKeys = usePublicJoinKeyRoutes();
  const roleRoutes = useRoleRoutes();
  const userInvitationRoutes = useUserInvitationRoutes();
  const workspaceInviteRoutes = useWorkspaceInviteRoutes();

  return (
    <Route path="selfservice">
      <Route path={"passports"} element={<UserPassportsScreen />}></Route>
      <Route
        path={"change-password/:uniqueId"}
        element={<ChangePasswordScreen />}
      ></Route>
      {/* See the matching route's own comment under useSelfServicePublicRoutes above -
          registered in both tables on purpose. */}
      <Route path={"reset-password"} element={<ResetPasswordScreen />}></Route>
      {publicJoinKeys}
      {roleRoutes}
      {userInvitationRoutes}
      {workspaceInviteRoutes}

      <Route
        path=""
        element={
          <AnimatedRouteWrapper>
            <SelfServiceHome />
          </AnimatedRouteWrapper>
        }
      />
    </Route>
  );
}
