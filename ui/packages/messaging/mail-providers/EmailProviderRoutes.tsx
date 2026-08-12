import { Route } from "react-router-dom";
import { EmailProviderDto } from "@fireback/messaging/sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "@fireback/ui-core/sdk/navigation/MessagingNavigation";
import { EmailProviderEntityManager } from "./EmailProviderEntityManager";
import { EmailProviderSingleScreen } from "./EmailProviderSingleScreen";
import { EmailProviderArchiveScreen } from "./EmailProviderArchiveScreen";

export function useEmailProviderRoutes() {
  return (
    <>
      <Route
        element={<EmailProviderEntityManager />}
        path={EmailProviderNavigation.Rcreate}
      />
      <Route
        element={<EmailProviderSingleScreen />}
        path={EmailProviderNavigation.Rsingle}
      ></Route>
      <Route
        element={<EmailProviderEntityManager />}
        path={EmailProviderNavigation.Redit}
      ></Route>
      <Route
        element={<EmailProviderArchiveScreen />}
        path={EmailProviderNavigation.Rquery}
      ></Route>
    </>
  );
}
