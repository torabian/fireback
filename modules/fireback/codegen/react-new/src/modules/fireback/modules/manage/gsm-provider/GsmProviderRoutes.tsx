import { Route } from "react-router-dom";
import { GsmProviderArchiveScreen } from "./GsmProviderArchiveScreen";
import { GsmProviderEntityManager } from "./GsmProviderEntityManager";
import { GsmProviderSingleScreen } from "./GsmProviderSingleScreen";
import { GsmProviderEntity } from "@/modules/fireback/sdk/modules/abac/GsmProviderEntity";
export function useGsmProviderRoutes() {
  return (
    <>
      <Route
        element={<GsmProviderEntityManager />}
        path={ GsmProviderEntity.Navigation.Rcreate}
      />
      <Route
        element={<GsmProviderSingleScreen />}
        path={ GsmProviderEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<GsmProviderEntityManager />}
        path={ GsmProviderEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<GsmProviderArchiveScreen />}
        path={  GsmProviderEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
