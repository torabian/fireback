import EssentialApp from "@fireback/enterprise-shell/EssentialApp";
import { WithWasmServer } from "@fireback/wasm-server/WithWasmServer";
import { ApplicationRoutes } from "./ApplicationRoutes";

function App() {
  return (
    <WithWasmServer>
      <EssentialApp ApplicationRoutes={ApplicationRoutes} />
    </WithWasmServer>
  );
}

export default App;
