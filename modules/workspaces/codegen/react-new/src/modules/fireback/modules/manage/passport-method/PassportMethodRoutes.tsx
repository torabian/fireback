import { Route } from "react-router-dom";
import { PassportMethodArchiveScreen } from "./PassportMethodArchiveScreen";
import { PassportMethodEntityManager } from "./PassportMethodEntityManager";
import { PassportMethodSingleScreen } from "./PassportMethodSingleScreen";
import { PassportMethodEntity } from "@/modules/fireback/sdk/modules/workspaces/PassportMethodEntity";
export function usePassportMethodRoutes() {
  return (
    <>
      <Route
        element={<PassportMethodEntityManager />}
        path={ PassportMethodEntity.Navigation.Rcreate}
      />
      <Route
        element={<PassportMethodSingleScreen />}
        path={ PassportMethodEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<PassportMethodEntityManager />}
        path={ PassportMethodEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<PassportMethodArchiveScreen />}
        path={  PassportMethodEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
