import { useCommonEntityManager } from "@/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/fireback/components/entity-manager/CommonEntityManager";
import { BrandForm } from "./BrandEditForm";
import { BrandEntity } from "src/sdk/fireback/modules/shop/BrandEntity";
import { useGetBrandByUniqueId } from "src/sdk/fireback/modules/shop/useGetBrandByUniqueId";
import { usePostBrand } from "src/sdk/fireback/modules/shop/usePostBrand";
import { usePatchBrand } from "src/sdk/fireback/modules/shop/usePatchBrand";
export const BrandEntityManager = ({ data }: DtoEntity<BrandEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<BrandEntity>
  >({
    data,
  });
  const getSingleHook = useGetBrandByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostBrand({
    queryClient,
  });
  const patchHook = usePatchBrand({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(BrandEntity.Navigation.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        BrandEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={BrandForm}
      onEditTitle={t.brands.editBrand}
      onCreateTitle={t.brands.newBrand}
      data={data}
    />
  );
};
