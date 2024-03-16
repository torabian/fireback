import { useCategoryRoutes } from "@/modules/shop/product/category/CategoryRoutes";
import { FirebackEssentialRouterManager } from "../core/EssentialRouter";

// ~ auto:useRouteImport

export function ApplicationRoutes() {
  // ~ auto:useRouteDefs

  const categoryRoutes = useCategoryRoutes();

  return (
    <FirebackEssentialRouterManager>
      {/* ~ auto:useRouteJsx */}
      {categoryRoutes}
    </FirebackEssentialRouterManager>
  );
}
