import { Route } from "react-router-dom";
import { PassportMethodArchiveScreen } from "./PassportMethodArchiveScreen";
import { PassportMethodEntityManager } from "./PassportMethodEntityManager";
import { PassportMethodSingleScreen } from "./PassportMethodSingleScreen";
import { PassportMethodNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
export function usePassportMethodRoutes() {
  return (
    <>
      <Route
        element={<PassportMethodEntityManager />}
        path={PassportMethodNavigation.Rcreate}
      />
      <Route
        element={<PassportMethodSingleScreen />}
        path={PassportMethodNavigation.Rsingle}
      ></Route>
      <Route
        element={<PassportMethodEntityManager />}
        path={PassportMethodNavigation.Redit}
      ></Route>
      <Route
        element={<PassportMethodArchiveScreen />}
        path={PassportMethodNavigation.Rquery}
      ></Route>
    </>
  );
}
