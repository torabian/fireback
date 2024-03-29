import { AbacModuleMockProvider } from "@/modules/essentials/AbacMockProvider";
import { BrandMockProvider } from "@/modules/shop/brand/BrandMockProvider";
import { CategoryMockProvider } from "@/modules/shop/category/CategoryMockProvider";
import { DiscountCodeMockProvider } from "@/modules/shop/discount-code/DiscountCodeMockProvider";
import { OrderMockProvider } from "@/modules/shop/order/OrderMockProvider";
import { ProductSubmissionMockProvider } from "@/modules/shop/product-submission/ProductSubmissionMockProvider";
import { ProductMockProvider } from "@/modules/shop/product/ProductMockProvider";
import { ShoppingCartMockProvider } from "@/modules/shop/shopping-cart/ShoppingCartMockProvider";
import { TagMockProvider } from "@/modules/shop/tag/TagMockProvider";

// ~ auto:useMockImport

export const FirebackMockServer = [
  new AbacModuleMockProvider(),
  new ShoppingCartMockProvider(),
  new ProductSubmissionMockProvider(),
  new ProductMockProvider(),
  new BrandMockProvider(),
  new CategoryMockProvider(),
  new DiscountCodeMockProvider(),
  new OrderMockProvider(),
  new TagMockProvider(),

  // ~ auto:useMocknew
];
