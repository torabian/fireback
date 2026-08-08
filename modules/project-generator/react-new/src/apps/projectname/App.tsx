import EssentialApp from "../core/EssentialApp";
import { ApplicationRoutes } from "./ApplicationRoutes";
import { WithSdk } from "./WithSdk";

function App() {
  return (
    <EssentialApp ApplicationRoutes={ApplicationRoutes} WithSdk={WithSdk} />
  );
}

export default App;
