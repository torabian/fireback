import { Route } from "react-router-dom";
import { CapabilityArchiveScreen } from "./CapabilityArchiveScreen";
import { CapabilityEntityManager } from "./CapabilityEntityManager";
import { CapabilitySingleScreen } from "./CapabilitySingleScreen";
import { CapabilityEntity } from "@/modules/fireback/sdk/modules/workspaces/CapabilityEntity";
export function useCapabilityRoutes() {
  return (
    <>
      <Route
        element={<CapabilityEntityManager />}
        path={ CapabilityEntity.Navigation.Rcreate}
      />
      <Route
        element={<CapabilitySingleScreen />}
        path={ CapabilityEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<CapabilityEntityManager />}
        path={ CapabilityEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<CapabilityArchiveScreen />}
        path={  CapabilityEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
