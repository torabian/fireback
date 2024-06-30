import { Route } from "react-router-dom";
import { PostArchiveScreen } from "./PostArchiveScreen";
import { PostEntityManager } from "./PostEntityManager";
import { PostSingleScreen } from "./PostSingleScreen";
import { PostEntity } from "src/sdk/fireback/modules/cms/PostEntity";
export function usePostRoutes() {
  return (
    <>
      <Route
        element={<PostEntityManager />}
        path={ PostEntity.Navigation.Rcreate}
      />
      <Route
        element={<PostSingleScreen />}
        path={ PostEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<PostEntityManager />}
        path={ PostEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<PostArchiveScreen />}
        path={  PostEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}