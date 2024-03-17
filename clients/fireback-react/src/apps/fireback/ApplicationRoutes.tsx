import { useCategoryRoutes } from "@/modules/shop/category/CategoryRoutes";
import { useProductSubmissionRoutes } from "@/modules/shop/product-submission/ProductSubmissionRoutes";
import { useProductRoutes } from "@/modules/shop/product/ProductRoutes";
import { FirebackEssentialRouterManager } from "../core/EssentialRouter";

// ~ auto:useRouteImport

export function ApplicationRoutes() {
  // ~ auto:useRouteDefs

  const categoryRoutes = useCategoryRoutes();
  const productRoutes = useProductRoutes();
  const productSubmissionRoutes = useProductSubmissionRoutes();

  return (
    <FirebackEssentialRouterManager>
      {/* ~ auto:useRouteJsx */}
      {categoryRoutes}
      {productRoutes}
      {productSubmissionRoutes}
    </FirebackEssentialRouterManager>
  );
}
