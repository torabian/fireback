import { Route } from "react-router-dom";
import { PageCategoryArchiveScreen } from "./PageCategoryArchiveScreen";
import { PageCategoryEntityManager } from "./PageCategoryEntityManager";
import { PageCategorySingleScreen } from "./PageCategorySingleScreen";
import { PageCategoryEntity } from "src/sdk/fireback/modules/cms/PageCategoryEntity";
export function usePageCategoryRoutes() {
  return (
    <>
      <Route
        element={<PageCategoryEntityManager />}
        path={ PageCategoryEntity.Navigation.Rcreate}
      />
      <Route
        element={<PageCategorySingleScreen />}
        path={ PageCategoryEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<PageCategoryEntityManager />}
        path={ PageCategoryEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<PageCategoryArchiveScreen />}
        path={  PageCategoryEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}