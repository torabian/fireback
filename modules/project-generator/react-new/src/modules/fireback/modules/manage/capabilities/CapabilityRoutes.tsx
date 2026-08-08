import { Route } from "react-router-dom";
import { CapabilityArchiveScreen } from "./CapabilityArchiveScreen";
import { CapabilityEntityManager } from "./CapabilityEntityManager";
import { CapabilitySingleScreen } from "./CapabilitySingleScreen";
import { CapabilityDto } from "../../../sdk/abac/CapabilityDto";
import { CapabilityNavigation } from "../../../sdk/navigation/AbacNavigation";
export function useCapabilityRoutes() {
  return (
    <>
      <Route
        element={<CapabilityEntityManager />}
        path={CapabilityNavigation.Rcreate}
      />
      <Route
        element={<CapabilitySingleScreen />}
        path={CapabilityNavigation.Rsingle}
      ></Route>
      <Route
        element={<CapabilityEntityManager />}
        path={CapabilityNavigation.Redit}
      ></Route>
      <Route
        element={<CapabilityArchiveScreen />}
        path={CapabilityNavigation.Rquery}
      ></Route>
    </>
  );
}
