import { Route } from "react-router-dom";
import { GsmProviderArchiveScreen } from "./GsmProviderArchiveScreen";
import { GsmProviderEntityManager } from "./GsmProviderEntityManager";
import { GsmProviderSingleScreen } from "./GsmProviderSingleScreen";
import { GsmProviderDto } from "../../sdk/messaging/GsmProviderDto";
import { GsmProviderNavigation } from "../../sdk/navigation/MessagingNavigation";
export function useGsmProviderRoutes() {
  return (
    <>
      <Route
        element={<GsmProviderEntityManager />}
        path={GsmProviderNavigation.Rcreate}
      />
      <Route
        element={<GsmProviderSingleScreen />}
        path={GsmProviderNavigation.Rsingle}
      ></Route>
      <Route
        element={<GsmProviderEntityManager />}
        path={GsmProviderNavigation.Redit}
      ></Route>
      <Route
        element={<GsmProviderArchiveScreen />}
        path={GsmProviderNavigation.Rquery}
      ></Route>
    </>
  );
}
