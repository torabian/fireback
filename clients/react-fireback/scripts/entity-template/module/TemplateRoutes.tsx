import { Route } from "react-router-dom";
import { TemplateArchiveScreen } from "./TemplateArchiveScreen";
import { TemplateEntityManager } from "./TemplateEntityManager";
import { TemplateSingleScreen } from "./TemplateSingleScreen";
import { TemplateNavigationTools } from "src/sdk/xsdk/modules/xmodule/xnavigation";

export function useTemplateRoutes() {
  return (
    <>
      <Route
        element={<TemplateEntityManager />}
        path={TemplateNavigationTools.Rcreate}
      />
      <Route
        element={<TemplateSingleScreen />}
        path={TemplateNavigationTools.Rsingle}
      ></Route>
      <Route
        element={<TemplateEntityManager />}
        path={TemplateNavigationTools.Redit}
      ></Route>
      <Route
        element={<TemplateArchiveScreen />}
        path={TemplateNavigationTools.Rquery}
      ></Route>
    </>
  );
}
