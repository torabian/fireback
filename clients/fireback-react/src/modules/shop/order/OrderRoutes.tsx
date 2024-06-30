import { Route } from "react-router-dom";
import { OrderArchiveScreen } from "./OrderArchiveScreen";
import { OrderEntityManager } from "./OrderEntityManager";
import { OrderSingleScreen } from "./OrderSingleScreen";
import { OrderEntity } from "src/sdk/fireback/modules/shop/OrderEntity";
export function useOrderRoutes() {
  return (
    <>
      <Route
        element={<OrderEntityManager />}
        path={ OrderEntity.Navigation.Rcreate}
      />
      <Route
        element={<OrderSingleScreen />}
        path={ OrderEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<OrderEntityManager />}
        path={ OrderEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<OrderArchiveScreen />}
        path={  OrderEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}