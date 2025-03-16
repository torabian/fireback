import { Route } from "react-router-dom";
import { UserEntityManager } from "./UserEntityManager";
import { UserEntity } from "@/modules/fireback/sdk/modules/workspaces/UserEntity";
import { UserSingleScreen } from "./UserSingleScreen";
import { UserArchiveScreen } from "./UserArchiveScreen";

export function useUserRoutes() {
  return (
    <>
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
}
