import { Route } from "react-router-dom";
import { UserEntityManager } from "./UserEntityManager";
import { UserNavigation } from "../../../sdk/navigation/AbacNavigation";
import { UserSingleScreen } from "./UserSingleScreen";
import { UserArchiveScreen } from "./UserArchiveScreen";

export function useUserRoutes() {
  return (
    <>
      <Route
        element={<UserEntityManager />}
        path={UserNavigation.Rcreate}
      />
      <Route
        element={<UserSingleScreen />}
        path={UserNavigation.Rsingle}
      ></Route>
      <Route
        element={<UserEntityManager />}
        path={UserNavigation.Redit}
      ></Route>
      <Route
        element={<UserArchiveScreen />}
        path={UserNavigation.Rquery}
      ></Route>
    </>
  );
}
