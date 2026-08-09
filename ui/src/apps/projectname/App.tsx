import EssentialApp from "../core/EssentialApp";
import { ApplicationRoutes } from "./ApplicationRoutes";
import { WithSdk } from "./WithSdk";
import {
  SessionGate,
  noopCheckSession,
} from "@/modules/fireback-ui/components/session-gate/SessionGate";

// TODO: swap noopCheckSession for a real session/whoami API call. Until then
// this resolves instantly, so boot behavior is unchanged.
function App() {
  return (
    <SessionGate checkSession={noopCheckSession}>
      <EssentialApp ApplicationRoutes={ApplicationRoutes} WithSdk={WithSdk} />
    </SessionGate>
  );
}

export default App;
