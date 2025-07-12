import { Route } from "react-router-dom";
import { PaymentConfigEntityManager } from "./PaymentConfigEntityManager";
import { PaymentConfigSingleScreen } from "./PaymentConfigSingleScreen";

export function usePaymentConfigRoutes() {
  return (
    <>
      <Route element={<PaymentConfigEntityManager />} path={"config/edit"} />
      <Route element={<PaymentConfigSingleScreen />} path={"config"}></Route>
    </>
  );
}
