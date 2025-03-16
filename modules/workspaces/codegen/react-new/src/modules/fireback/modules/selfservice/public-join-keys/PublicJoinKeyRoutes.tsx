import { Route } from "react-router-dom";
import { PublicJoinKeyEntityManager } from "./PublicJoinKeyEntityManager";
import { PublicJoinKeyEntity } from "@/modules/fireback/sdk/modules/workspaces/PublicJoinKeyEntity";
import { PublicJoinKeySingleScreen } from "./PublicJoinKeySingleScreen";
import { PublicJoinKeyArchiveScreen } from "./PublicJoinKeyArchiveScreen";

export function usePublicJoinKeyRoutes() {
  return (
    <>
      <Route
        element={<PublicJoinKeyEntityManager />}
        path={PublicJoinKeyEntity.Navigation.Rcreate}
      />
      <Route
        element={<PublicJoinKeySingleScreen />}
        path={PublicJoinKeyEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<PublicJoinKeyEntityManager />}
        path={PublicJoinKeyEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<PublicJoinKeyArchiveScreen />}
        path={PublicJoinKeyEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
