import { useCommonEntityManager } from "@/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/fireback/components/entity-manager/CommonEntityManager";
import { ProductForm } from "./ProductEditForm";
import { ProductEntity } from "src/sdk/fireback/modules/shop/ProductEntity";
import { useGetProductByUniqueId } from "src/sdk/fireback/modules/shop/useGetProductByUniqueId";
import { usePostProduct } from "src/sdk/fireback/modules/shop/usePostProduct";
import { usePatchProduct } from "src/sdk/fireback/modules/shop/usePatchProduct";
export const ProductEntityManager = ({ data }: DtoEntity<ProductEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<ProductEntity>
  >({
    data,
  });
  const getSingleHook = useGetProductByUniqueId({
    query: { uniqueId, deep: true },
  });
  const postHook = usePostProduct({
    queryClient,
  });
  const patchHook = usePatchProduct({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      customClass=""
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          ProductEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        ProductEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={ProductForm}
      onEditTitle={t.products.editproduct}
      onCreateTitle={t.products.newproduct}
      data={data}
    />
  );
};
