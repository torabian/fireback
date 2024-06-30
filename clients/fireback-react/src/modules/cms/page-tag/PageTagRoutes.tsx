import { Route } from "react-router-dom";
import { PageTagArchiveScreen } from "./PageTagArchiveScreen";
import { PageTagEntityManager } from "./PageTagEntityManager";
import { PageTagSingleScreen } from "./PageTagSingleScreen";
import { PageTagEntity } from "src/sdk/fireback/modules/cms/PageTagEntity";
export function usePageTagRoutes() {
  return (
    <>
      <Route
        element={<PageTagEntityManager />}
        path={PageTagEntity.Navigation.Rcreate}
      />
      <Route
        element={<PageTagSingleScreen />}
        path={PageTagEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<PageTagEntityManager />}
        path={PageTagEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<PageTagArchiveScreen />}
        path={PageTagEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
