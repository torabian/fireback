import { Route } from "react-router-dom";

import { DemoFormSelect } from "./demo/DemoFormSelect";
import { DemoScreen } from "./demo/DemoScreen";
import { DemoModal } from "./demo/DemoModal";
import { DemoFormDates } from "./demo/DemoFormDates";
import { FirebackEssentialRouterManager } from "@fireback/enterprise-shell/EssentialRouter";
import { DemoJsf } from "./demo/DemoJsf";

export function ApplicationRoutes({ routerId }: { routerId?: string }) {
  return (
    <FirebackEssentialRouterManager routerId={routerId}>
      {/* ~ auto:useRouteJsx */}
      <Route path={"demo/form-select"} element={<DemoFormSelect />}></Route>
      <Route path={"demo/jsf"} element={<DemoJsf />}></Route>
      <Route path={"demo/modals"} element={<DemoModal />}></Route>
      <Route path={"demo/form-date"} element={<DemoFormDates />}></Route>
      <Route path={"demo"} element={<DemoScreen />}></Route>
    </FirebackEssentialRouterManager>
  );
}
