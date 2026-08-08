import { ApplicationRoutes } from "./ApplicationRoutes";
import { WithSdk } from "./WithSdk";

import EssentialApp from "@/modules/fireback/apps/core/EssentialApp";

function App() {
  return (
    <EssentialApp ApplicationRoutes={ApplicationRoutes} WithSdk={WithSdk} />
  );
}

export default App;
