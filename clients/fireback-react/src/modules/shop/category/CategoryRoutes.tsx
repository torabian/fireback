import { Route } from "react-router-dom";
import { CategoryArchiveScreen } from "./CategoryArchiveScreen";
import { CategoryEntityManager } from "./CategoryEntityManager";
import { CategorySingleScreen } from "./CategorySingleScreen";
import { CategoryEntity } from "src/sdk/fireback/modules/shop/CategoryEntity";
export function useCategoryRoutes() {
  return (
    <>
      <Route
        element={<CategoryEntityManager />}
        path={ CategoryEntity.Navigation.Rcreate}
      />
      <Route
        element={<CategorySingleScreen />}
        path={ CategoryEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<CategoryEntityManager />}
        path={ CategoryEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<CategoryArchiveScreen />}
        path={  CategoryEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}