import { Route } from "react-router-dom";
import { TagArchiveScreen } from "./TagArchiveScreen";
import { TagEntityManager } from "./TagEntityManager";
import { TagSingleScreen } from "./TagSingleScreen";
import { TagEntity } from "src/sdk/fireback/modules/shop/TagEntity";
export function useTagRoutes() {
  return (
    <>
      <Route
        element={<TagEntityManager />}
        path={ TagEntity.Navigation.Rcreate}
      />
      <Route
        element={<TagSingleScreen />}
        path={ TagEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<TagEntityManager />}
        path={ TagEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<TagArchiveScreen />}
        path={  TagEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}