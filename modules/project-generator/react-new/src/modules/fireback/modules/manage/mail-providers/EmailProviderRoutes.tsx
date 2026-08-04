import { Route } from "react-router-dom";
import { EmailProviderEntity } from "@/modules/fireback/sdk/modules/abac/EmailProviderEntity";
import { EmailProviderEntityManager } from "./EmailProviderEntityManager";
import { EmailProviderSingleScreen } from "./EmailProviderSingleScreen";
import { EmailProviderArchiveScreen } from "./EmailProviderArchiveScreen";

export function useEmailProviderRoutes() {
  return (
    <>
      <Route
        element={<EmailProviderEntityManager />}
        path={EmailProviderEntity.Navigation.Rcreate}
      />
      <Route
        element={<EmailProviderSingleScreen />}
        path={EmailProviderEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<EmailProviderEntityManager />}
        path={EmailProviderEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<EmailProviderArchiveScreen />}
        path={EmailProviderEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
