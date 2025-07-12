import { Route } from "react-router-dom";
import { usePaymentConfigRoutes } from "./config/PaymentConfigRoutes";
import { useInvoiceRoutes } from "./invoice/InvoiceRoutes";

export function usePaymentRoutes() {
  const paymentConfigRoutes = usePaymentConfigRoutes();
  const invoiceRoutes = useInvoiceRoutes();

  return (
    <Route path="payment">
      {paymentConfigRoutes}
      {invoiceRoutes}
    </Route>
  );
}
