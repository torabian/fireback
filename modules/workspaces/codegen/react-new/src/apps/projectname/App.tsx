import { useRef } from "react";
import { ApplicationRoutes } from "./ApplicationRoutes";
import { WithSdk } from "./WithSdk";

import { FirebackMockServer } from "./mockServer";
import EssentialApp from "@/modules/fireback/apps/core/EssentialApp";
import { Webrtc } from "./Webrtc";

function App() {
  const mockServer = useRef<any>(FirebackMockServer);
  return (
    <>
      <Webrtc />
    </>
  );
}

export default App;
