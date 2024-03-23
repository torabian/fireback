import { Route } from "react-router-dom";
import { PostCategoryArchiveScreen } from "./PostCategoryArchiveScreen";
import { PostCategoryEntityManager } from "./PostCategoryEntityManager";
import { PostCategorySingleScreen } from "./PostCategorySingleScreen";
import { PostCategoryEntity } from "src/sdk/fireback/modules/cms/PostCategoryEntity";
export function usePostCategoryRoutes() {
  return (
    <>
      <Route
        element={<PostCategoryEntityManager />}
        path={PostCategoryEntity.Navigation.Rcreate}
      />
      <Route
        element={<PostCategorySingleScreen />}
        path={PostCategoryEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<PostCategoryEntityManager />}
        path={PostCategoryEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<PostCategoryArchiveScreen />}
        path={PostCategoryEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
