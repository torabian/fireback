import { withExtended } from "@/fireback/hooks/columns";
import { enTranslations } from "@/translations/en";
import { ShoppingCartEntity } from "src/sdk/fireback/modules/shop/ShoppingCartEntity";

export const columns = (t: typeof enTranslations) =>
  withExtended(t, [
    {
      name: ShoppingCartEntity.Fields.items,
      title: t.shoppingCarts.items,
      width: 100,
    },
  ]);
