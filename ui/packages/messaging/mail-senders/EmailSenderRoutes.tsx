import { Route } from "react-router-dom";
import { EmailSenderEntityManager } from "./EmailSenderEntityManager";
import { EmailSenderDto } from "@fireback/messaging/sdk/messaging/EmailSenderDto";
import { EmailSenderNavigation } from "@fireback/ui-core/sdk/navigation/MessagingNavigation";
import { EmailSenderSingleScreen } from "./EmailSenderSingleScreen";
import { EmailSenderArchiveScreen } from "./EmailSenderArchiveScreen";

export function useEmailSenderRoutes() {
  return (
    <>
      <Route
        element={<EmailSenderEntityManager />}
        path={EmailSenderNavigation.Rcreate}
      />
      <Route
        element={<EmailSenderSingleScreen />}
        path={EmailSenderNavigation.Rsingle}
      ></Route>
      <Route
        element={<EmailSenderEntityManager />}
        path={EmailSenderNavigation.Redit}
      ></Route>
      <Route
        element={<EmailSenderArchiveScreen />}
        path={EmailSenderNavigation.Rquery}
      ></Route>
    </>
  );
}
