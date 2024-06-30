import { Route } from "react-router-dom";
import { ProductArchiveScreen } from "./ProductArchiveScreen";
import { ProductEntityManager } from "./ProductEntityManager";
import { ProductSingleScreen } from "./ProductSingleScreen";
import { ProductEntity } from "src/sdk/fireback/modules/shop/ProductEntity";
export function useProductRoutes() {
  return (
    <>
      <Route
        element={<ProductEntityManager />}
        path={ ProductEntity.Navigation.Rcreate}
      />
      <Route
        element={<ProductSingleScreen />}
        path={ ProductEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<ProductEntityManager />}
        path={ ProductEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<ProductArchiveScreen />}
        path={  ProductEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}