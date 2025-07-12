import { Route } from "react-router-dom";
import { InvoiceArchiveScreen } from "./InvoiceArchiveScreen";
import { InvoiceEntityManager } from "./InvoiceEntityManager";
import { InvoiceSingleScreen } from "./InvoiceSingleScreen";
import { InvoiceEntity } from "@/modules/fireback/sdk/modules/payment/InvoiceEntity";
export function useInvoiceRoutes() {
  return (
    <>
      <Route
        element={<InvoiceEntityManager />}
        path={InvoiceEntity.Navigation.Rcreate}
      />
      <Route
        element={<InvoiceSingleScreen />}
        path={InvoiceEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<InvoiceEntityManager />}
        path={InvoiceEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<InvoiceArchiveScreen />}
        path={InvoiceEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
