import { Route } from "react-router-dom";
import { FirebackEssentialRouterManager } from "../../modules/fireback/apps/core/EssentialRouter";
import { DemoFormSelect } from "./demo/DemoFormSelect";
import { DemoScreen } from "./demo/DemoScreen";
import { DemoModal } from "./demo/DemoModal";
import { DemoFormDates } from "./demo/DemoFormDates";

// ~ auto:useRouteImport

export function ApplicationRoutes({ routerId }: { routerId?: string }) {
  // ~ auto:useRouteDefs

  return (
    <FirebackEssentialRouterManager routerId={routerId}>
      {/* ~ auto:useRouteJsx */}
      <Route path={"demo/form-select"} element={<DemoFormSelect />}></Route>
      <Route path={"demo/modals"} element={<DemoModal />}></Route>
      <Route path={"demo/form-date"} element={<DemoFormDates />}></Route>
      <Route path={"demo"} element={<DemoScreen />}></Route>
    </FirebackEssentialRouterManager>
  );
}
