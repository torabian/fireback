/**
 * Fireback manage routes,
 * It's for administration a root level content.
 * Some components can be used for sub-level workspaces, but this is not planned yet
 *
 * All routes regarding manage are authenticated, they do not expose public components.
 */

import { Route } from "react-router-dom";
import { DashboardScreen } from "./DashboardScreen";
import { AnimatedRouteWrapper } from "@/modules/fireback/apps/core/SwipeTransition";

export function useMobileKitRoutes() {
  return (
    <Route path="">
      <Route
        element={
          <AnimatedRouteWrapper>
            <DashboardScreen />
          </AnimatedRouteWrapper>
        }
        path={"dashboard"}
      />
    </Route>
  );
}
