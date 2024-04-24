import { Route } from "react-router-dom";
import { AppMenuArchiveScreen } from "./AppMenuArchiveScreen";
import { AppMenuEntityManager } from "./AppMenuEntityManager";
import { AppMenuSingleScreen } from "./AppMenuSingleScreen";
import { AppMenuEntity } from "src/sdk/fireback/modules/workspaces/AppMenuEntity";
export function useAppMenuRoutes() {
  return (
    <>
      <Route
        element={<AppMenuEntityManager />}
        path={ AppMenuEntity.Navigation.Rcreate}
      />
      <Route
        element={<AppMenuSingleScreen />}
        path={ AppMenuEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<AppMenuEntityManager />}
        path={ AppMenuEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<AppMenuArchiveScreen />}
        path={  AppMenuEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}