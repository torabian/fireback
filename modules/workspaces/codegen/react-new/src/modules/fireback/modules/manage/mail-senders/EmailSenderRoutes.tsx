import { Route } from "react-router-dom";
import { EmailSenderEntityManager } from "./EmailSenderEntityManager";
import { EmailSenderEntity } from "@/modules/fireback/sdk/modules/workspaces/EmailSenderEntity";
import { EmailSenderSingleScreen } from "./EmailSenderSingleScreen";
import { EmailSenderArchiveScreen } from "./EmailSenderArchiveScreen";

export function useEmailSenderRoutes() {
  return (
    <>
      <Route
        element={<EmailSenderEntityManager />}
        path={EmailSenderEntity.Navigation.Rcreate}
      />
      <Route
        element={<EmailSenderSingleScreen />}
        path={EmailSenderEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<EmailSenderEntityManager />}
        path={EmailSenderEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<EmailSenderArchiveScreen />}
        path={EmailSenderEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
