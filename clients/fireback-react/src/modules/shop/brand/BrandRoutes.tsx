import { Route } from "react-router-dom";
import { BrandArchiveScreen } from "./BrandArchiveScreen";
import { BrandEntityManager } from "./BrandEntityManager";
import { BrandSingleScreen } from "./BrandSingleScreen";
import { BrandEntity } from "src/sdk/fireback/modules/shop/BrandEntity";
export function useBrandRoutes() {
  return (
    <>
      <Route
        element={<BrandEntityManager />}
        path={ BrandEntity.Navigation.Rcreate}
      />
      <Route
        element={<BrandSingleScreen />}
        path={ BrandEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<BrandEntityManager />}
        path={ BrandEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<BrandArchiveScreen />}
        path={  BrandEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}