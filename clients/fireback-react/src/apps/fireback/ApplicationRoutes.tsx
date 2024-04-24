import { useAppMenuRoutes } from "@/modules/cms/menu/AppMenuRoutes";
import { usePageCategoryRoutes } from "@/modules/cms/page-category/PageCategoryRoutes";
import { usePageTagRoutes } from "@/modules/cms/page-tag/PageTagRoutes";
import { usePageRoutes } from "@/modules/cms/page/PageRoutes";
import { usePostCategoryRoutes } from "@/modules/cms/post-category/PostCategoryRoutes";
import { usePostTagRoutes } from "@/modules/cms/post-tag/PostTagRoutes";
import { usePostRoutes } from "@/modules/cms/post/PostRoutes";
import { useBrandRoutes } from "@/modules/shop/brand/BrandRoutes";
import { useCategoryRoutes } from "@/modules/shop/category/CategoryRoutes";
import { useDiscountCodeRoutes } from "@/modules/shop/discount-code/DiscountCodeRoutes";
import { useOrderRoutes } from "@/modules/shop/order/OrderRoutes";
import { useProductSubmissionRoutes } from "@/modules/shop/product-submission/ProductSubmissionRoutes";
import { useProductRoutes } from "@/modules/shop/product/ProductRoutes";
import { useShoppingCartRoutes } from "@/modules/shop/shopping-cart/ShoppingCartRoutes";
import { useTagRoutes } from "@/modules/shop/tag/TagRoutes";
import { FirebackEssentialRouterManager } from "../core/EssentialRouter";

// ~ auto:useRouteImport

export function ApplicationRoutes() {
  // ~ auto:useRouteDefs

  const categoryRoutes = useCategoryRoutes();
  const productRoutes = useProductRoutes();
  const productSubmissionRoutes = useProductSubmissionRoutes();
  const tagRoutes = useTagRoutes();
  const brandRoutes = useBrandRoutes();
  const pageRoutes = usePageRoutes();
  const pageCategoryRoutes = usePageCategoryRoutes();
  const pageTagRoutes = usePageTagRoutes();
  const postRoutes = usePostRoutes();
  const postCategoryRoutes = usePostCategoryRoutes();
  const postTagRoutes = usePostTagRoutes();
  const discountCodeRoutes = useDiscountCodeRoutes();
  const shoppingCartRoutes = useShoppingCartRoutes();
  const orderRoutes = useOrderRoutes();
  const appMenuRoutes = useAppMenuRoutes();

  return (
    <FirebackEssentialRouterManager>
      {/* ~ auto:useRouteJsx */}
      {categoryRoutes}
      {productRoutes}
      {productSubmissionRoutes}
      {tagRoutes}
      {orderRoutes}
      {brandRoutes}
      {pageRoutes}
      {appMenuRoutes}
      {pageCategoryRoutes}
      {pageTagRoutes}
      {postRoutes}
      {postCategoryRoutes}
      {postTagRoutes}
      {discountCodeRoutes}
      {shoppingCartRoutes}
    </FirebackEssentialRouterManager>
  );
}
