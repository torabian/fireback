import { Route } from "react-router-dom";
import { InternalStatsDashboard } from "./InternalStatsDashboard";

// internal-stats is a single dashboard screen, not an entity CRUD (no
// list/create/edit/archive the way capabilities/workspaces/... have) - one
// route is all it needs. Path matches the AppMenu Href main.go registers
// ("/manage/internal-stats" - see cmd/fireback/main.go).
export function useInternalStatsRoutes() {
  return (
    <>
      <Route element={<InternalStatsDashboard />} path={"internal-stats"} />
    </>
  );
}
