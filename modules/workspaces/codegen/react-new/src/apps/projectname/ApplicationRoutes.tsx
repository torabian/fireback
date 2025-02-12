import { Route } from "react-router-dom";
import { FirebackEssentialRouterManager } from "../../modules/fireback/apps/core/EssentialRouter";
import { DemoScreen } from "./demo/DemoScreen";

// ~ auto:useRouteImport

export function ApplicationRoutes({ routerId }: { routerId?: string }) {
  // ~ auto:useRouteDefs

  return (
    <FirebackEssentialRouterManager routerId={routerId}>
      {/* ~ auto:useRouteJsx */}
      <Route path={"demo"} element={<DemoScreen />}></Route>
    </FirebackEssentialRouterManager>
  );
}
