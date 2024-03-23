import { Route } from "react-router-dom";
import { PostTagArchiveScreen } from "./PostTagArchiveScreen";
import { PostTagEntityManager } from "./PostTagEntityManager";
import { PostTagSingleScreen } from "./PostTagSingleScreen";
import { PostTagEntity } from "src/sdk/fireback/modules/cms/PostTagEntity";
export function usePostTagRoutes() {
  return (
    <>
      <Route
        element={<PostTagEntityManager />}
        path={PostTagEntity.Navigation.Rcreate}
      />
      <Route
        element={<PostTagSingleScreen />}
        path={PostTagEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<PostTagEntityManager />}
        path={PostTagEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<PostTagArchiveScreen />}
        path={PostTagEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
