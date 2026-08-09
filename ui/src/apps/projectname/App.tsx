import EssentialApp from "../core/EssentialApp";
import { ApplicationRoutes } from "./ApplicationRoutes";
import { WithSdk } from "./WithSdk";
import { SessionGate } from "@/modules/fireback-ui/components/session-gate/SessionGate";
import { checkSessionViaWhoami } from "@/modules/fireback-ui/components/session-gate/checkSessionViaWhoami";

function App() {
  return (
    <SessionGate checkSession={checkSessionViaWhoami}>
      <EssentialApp ApplicationRoutes={ApplicationRoutes} WithSdk={WithSdk} />
    </SessionGate>
  );
}

export default App;
