import { Route } from "react-router-dom";
import { EmailProviderDto } from "../../../sdk/messaging/EmailProviderDto";
import { EmailProviderNavigation } from "../../../sdk/navigation/MessagingNavigation";
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
