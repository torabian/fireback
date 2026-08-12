import { Route } from "react-router-dom";
import { PublicJoinKeyEntityManager } from "./PublicJoinKeyEntityManager";
import { PublicJoinKeyDto } from "@fireback/selfservice/sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { PublicJoinKeySingleScreen } from "./PublicJoinKeySingleScreen";
import { PublicJoinKeyArchiveScreen } from "./PublicJoinKeyArchiveScreen";

export function usePublicJoinKeyRoutes() {
  return (
    <>
      <Route
        element={<PublicJoinKeyEntityManager />}
        path={PublicJoinKeyNavigation.Rcreate}
      />
      <Route
        element={<PublicJoinKeySingleScreen />}
        path={PublicJoinKeyNavigation.Rsingle}
      ></Route>
      <Route
        element={<PublicJoinKeyEntityManager />}
        path={PublicJoinKeyNavigation.Redit}
      ></Route>
      <Route
        element={<PublicJoinKeyArchiveScreen />}
        path={PublicJoinKeyNavigation.Rquery}
      ></Route>
    </>
  );
}
