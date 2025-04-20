import { Route } from "react-router-dom";

import { UserInvitationArchiveScreen } from "./UserInvitationArchiveScreen";

export function useUserInvitationRoutes() {
  return (
    <>
      <Route
        element={<UserInvitationArchiveScreen />}
        path={"user-invitations"}
      ></Route>
    </>
  );
}
