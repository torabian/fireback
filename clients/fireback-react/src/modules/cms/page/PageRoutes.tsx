import { Route } from "react-router-dom";
import { PageArchiveScreen } from "./PageArchiveScreen";
import { PageEntityManager } from "./PageEntityManager";
import { PageSingleScreen } from "./PageSingleScreen";
import { PageEntity } from "src/sdk/fireback/modules/cms/PageEntity";
export function usePageRoutes() {
  return (
    <>
      <Route
        element={<PageEntityManager />}
        path={PageEntity.Navigation.Rcreate}
      />
      <Route
        element={<PageSingleScreen />}
        path={PageEntity.Navigation.Rsingle}
      ></Route>
      <Route
        element={<PageEntityManager />}
        path={PageEntity.Navigation.Redit}
      ></Route>
      <Route
        element={<PageArchiveScreen />}
        path={PageEntity.Navigation.Rquery}
      ></Route>
    </>
  );
}
