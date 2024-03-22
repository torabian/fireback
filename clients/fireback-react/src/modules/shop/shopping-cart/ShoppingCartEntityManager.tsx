import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { ShoppingCartForm } from "./ShoppingCartEditForm";
import { ShoppingCartEntity } from "src/sdk/fireback/modules/shop/ShoppingCartEntity";
import { useGetShoppingCartByUniqueId } from "src/sdk/fireback/modules/shop/useGetShoppingCartByUniqueId";
import { usePostShoppingCart } from "src/sdk/fireback/modules/shop/usePostShoppingCart";
import { usePatchShoppingCart } from "src/sdk/fireback/modules/shop/usePatchShoppingCart";
export const ShoppingCartEntityManager = ({
  data,
}: DtoEntity<ShoppingCartEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<ShoppingCartEntity>
  >({
    data,
  });
  const getSingleHook = useGetShoppingCartByUniqueId({
    query: { uniqueId, withPreloads: "Items.Product" },
  });
  const postHook = usePostShoppingCart({
    queryClient,
  });
  const patchHook = usePatchShoppingCart({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          ShoppingCartEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        ShoppingCartEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={ShoppingCartForm}
      onEditTitle={t.shoppingCarts.editShoppingCart}
      onCreateTitle={t.shoppingCarts.newShoppingCart}
      data={data}
    />
  );
};
