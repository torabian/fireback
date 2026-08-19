import { Route } from "react-router-dom";
import { AnalyticsDashboard } from "./AnalyticsDashboard";

// analytics is a single dashboard screen, not an entity CRUD - one route is all it
// needs, same as internal-stats. Path matches the AppMenu Href main.go registers
// ("/manage/analytics" - see cmd/fireback/main.go).
export function useAnalyticsRoutes() {
  return (
    <>
      <Route element={<AnalyticsDashboard />} path={"analytics"} />
    </>
  );
}
