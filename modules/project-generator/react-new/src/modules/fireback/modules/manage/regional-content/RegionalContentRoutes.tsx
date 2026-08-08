import { Route } from "react-router-dom";
import { RegionalContentArchiveScreen } from "./RegionalContentArchiveScreen";
import { RegionalContentEntityManager } from "./RegionalContentEntityManager";
import { RegionalContentSingleScreen } from "./RegionalContentSingleScreen";
import { RegionalContentDto } from "@/modules/fireback/sdk/abac/RegionalContentDto";
import { RegionalContentNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
export function useRegionalContentRoutes() {
  return (
    <>
      <Route
        element={<RegionalContentEntityManager />}
        path={ RegionalContentNavigation.Rcreate}
      />
      <Route
        element={<RegionalContentSingleScreen />}
        path={ RegionalContentNavigation.Rsingle}
      ></Route>
      <Route
        element={<RegionalContentEntityManager />}
        path={ RegionalContentNavigation.Redit}
      ></Route>
      <Route
        element={<RegionalContentArchiveScreen />}
        path={  RegionalContentNavigation.Rquery}
      ></Route>
    </>
  );
}
