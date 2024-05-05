import { useT } from "@/fireback/hooks/useT";
import { CommonArchiveManager } from "@/fireback/components/entity-manager/CommonArchiveManager";
import { ShoppingCartList } from "./ShoppingCartList";
import { ShoppingCartEntity } from "src/sdk/fireback/modules/shop/ShoppingCartEntity";
export const ShoppingCartArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.shoppingCarts.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(ShoppingCartEntity.Navigation.create(locale));
      }}
    >
      <ShoppingCartList />
    </CommonArchiveManager>
  );
};
