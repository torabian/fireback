import { Route } from "react-router-dom";
import { SendNotification } from "./SendNotification";

// A single action screen, not an entity CRUD - one route is all it needs, same as
// analytics/internal-stats. Path matches the AppMenu Href main.go registers
// ("/manage/notifications/send" - see cmd/fireback/main.go).
export function useNotificationsRoutes() {
  return (
    <>
      <Route element={<SendNotification />} path={"notifications/send"} />
    </>
  );
}
