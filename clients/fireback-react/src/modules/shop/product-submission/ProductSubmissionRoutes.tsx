import { Route } from "react-router-dom";
import { ProductSubmissionArchiveScreen } from "./ProductSubmissionArchiveScreen";
import { ProductSubmissionEntityManager } from "./ProductSubmissionEntityManager";
import { ProductSubmissionSingleScreen } from "./ProductSubmissionSingleScreen";
import { ProductSubmissionEntity } from "src/sdk/fireback/modules/shop/ProductSubmissionEntity";
export function useProductSubmissionRoutes() {
  return (
    <>
      <Route
        element={<ProductSubmissionEntityManager />}
        path={ ProductSubmissionEntity.Navigation.Rcreate}
      />
      <Route
        element={<ProductSubmissionSingleScreen />}
        path={ ProductSubmissionEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<ProductSubmissionEntityManager />}
        path={ ProductSubmissionEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<ProductSubmissionArchiveScreen />}
        path={  ProductSubmissionEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}