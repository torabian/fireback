import "react-toastify/dist/ReactToastify.css";

import { useRef } from "react";
import { ApplicationRoutes } from "./ApplicationRoutes";
import { WithSdk } from "./WithSdk";

import EssentialApp from "@/modules/fireback/apps/core/EssentialApp";
import { FirebackMockServer } from "./mockServer";

function App() {
  const mockServer = useRef<any>(FirebackMockServer);
  return (
    <EssentialApp
      ApplicationRoutes={ApplicationRoutes}
      mockServer={mockServer}
      WithSdk={WithSdk}
    />
  );
}

export default App;
