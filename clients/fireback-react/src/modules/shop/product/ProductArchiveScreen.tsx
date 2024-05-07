import { useT } from "@/fireback/hooks/useT";
import { CommonArchiveManager } from "@/fireback/components/entity-manager/CommonArchiveManager";
import { ProductList } from "./ProductList";
import { ProductEntity } from "src/sdk/fireback/modules/shop/ProductEntity";
export const ProductArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.products.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(ProductEntity.Navigation.create(locale));
      }}
    >
      <ProductList />
    </CommonArchiveManager>
  );
};
