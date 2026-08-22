import { Route } from "react-router-dom";
import { MessagingConfigEntityManager } from "./MessagingConfigEntityManager";
import { MessagingConfigSingleScreen } from "./MessagingConfigSingleScreen";
export function useMessagingConfigRoutes() {
  return (
    <>
      <Route
        element={<MessagingConfigSingleScreen />}
        path={"messaging-config"}
      ></Route>
      <Route
        element={<MessagingConfigEntityManager />}
        path={"messaging-config/edit"}
      ></Route>
    </>
  );
}
