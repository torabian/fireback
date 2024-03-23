import { Route } from "react-router-dom";
import { ShoppingCartArchiveScreen } from "./ShoppingCartArchiveScreen";
import { ShoppingCartEntityManager } from "./ShoppingCartEntityManager";
import { ShoppingCartSingleScreen } from "./ShoppingCartSingleScreen";
import { ShoppingCartEntity } from "src/sdk/fireback/modules/shop/ShoppingCartEntity";
export function useShoppingCartRoutes() {
  return (
    <>
      <Route
        element={<ShoppingCartEntityManager />}
        path={ ShoppingCartEntity.Navigation.Rcreate}
      />
      <Route
        element={<ShoppingCartSingleScreen />}
        path={ ShoppingCartEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<ShoppingCartEntityManager />}
        path={ ShoppingCartEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<ShoppingCartArchiveScreen />}
        path={  ShoppingCartEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}