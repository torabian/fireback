import { Route } from "react-router-dom";
import { RegionalContentArchiveScreen } from "./RegionalContentArchiveScreen";
import { RegionalContentEntityManager } from "./RegionalContentEntityManager";
import { RegionalContentSingleScreen } from "./RegionalContentSingleScreen";
import { RegionalContentEntity } from "@/modules/fireback/sdk/modules/abac/RegionalContentEntity";
export function useRegionalContentRoutes() {
  return (
    <>
      <Route
        element={<RegionalContentEntityManager />}
        path={ RegionalContentEntity.Navigation.Rcreate}
      />
      <Route
        element={<RegionalContentSingleScreen />}
        path={ RegionalContentEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<RegionalContentEntityManager />}
        path={ RegionalContentEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<RegionalContentArchiveScreen />}
        path={  RegionalContentEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
